package wallet

import (
	"context"
	"fmt"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/keystore"
	"gorm.io/gorm"
)

// ExportKeysResult reports what export-keys did for each actor
type ExportKeysResult struct {
	Exported int      // actors whose keys were saved to keystore + wallet record created
	Skipped  int      // actors whose wallet already exists (by address match)
	Errors   []string // actors that failed (logged but don't abort)
}

// exports private keys from the legacy Actor.PrivateKey column into the
// filesystem keystore, creating Wallet records where needed.
// idempotent -- skips actors whose address already has a Wallet record.
// does not delete Actor.PrivateKey; caller decides when to drop the column.
func ExportKeysHandler(
	ctx context.Context,
	db *gorm.DB,
	ks keystore.KeyStore,
) (*ExportKeysResult, error) {
	db = db.WithContext(ctx)

	var actors []model.Actor
	err := db.Where("private_key != '' AND private_key IS NOT NULL").Find(&actors).Error
	if err != nil {
		return nil, errors.Wrap(err, "failed to query actors with private keys")
	}

	result := &ExportKeysResult{}

	for _, actor := range actors {
		exported, errMsg := exportOneKey(db, ks, actor)
		if errMsg != "" {
			result.Errors = append(result.Errors, errMsg)
			continue
		}
		if exported {
			result.Exported++
		} else {
			result.Skipped++
		}
	}

	return result, nil
}

// exports a single actor's key to keystore, returns (true, "") on success,
// (false, "") on skip, ("", errMsg) on failure
func exportOneKey(db *gorm.DB, ks keystore.KeyStore, actor model.Actor) (exported bool, errMsg string) {
	// derive address to check for existing wallet
	addr, err := keystore.AddressFromExport(actor.PrivateKey)
	if err != nil {
		return false, fmt.Sprintf("actor %s: invalid key format: %v", actor.ID, err)
	}

	// check if wallet already exists for this address
	var existing model.Wallet
	err = db.Where("address = ?", addr.String()).First(&existing).Error
	if err == nil {
		// wallet exists, link actor if not already linked
		if existing.ActorID == nil {
			existing.ActorID = &actor.ID
			db.Save(&existing)
		}
		// verify the key file actually exists on disk
		if !ks.Has(existing.KeyPath) {
			return false, fmt.Sprintf("actor %s: wallet record exists but key file missing at %s", actor.ID, existing.KeyPath)
		}
		logger.Debugw("wallet already exists for actor", "actorID", actor.ID, "address", addr.String())
		return false, ""
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return false, fmt.Sprintf("actor %s: db query failed: %v", actor.ID, err)
	}

	// save key to keystore
	keyPath, _, err := ks.Put(actor.PrivateKey)
	if err != nil {
		return false, fmt.Sprintf("actor %s: keystore write failed: %v", actor.ID, err)
	}

	// create wallet record
	w := model.Wallet{
		KeyPath:  keyPath,
		KeyStore: "local",
		Address:  addr.String(),
		ActorID:  &actor.ID,
	}
	if err := db.Create(&w).Error; err != nil {
		// cleanup keystore file on db failure
		ks.Delete(keyPath)
		return false, fmt.Sprintf("actor %s: wallet create failed: %v", actor.ID, err)
	}

	logger.Infow("exported actor key to keystore",
		"actorID", actor.ID, "address", addr.String(), "walletID", w.ID)
	return true, ""
}

func HasPrivateKeyColumn(db *gorm.DB) bool {
	return db.Migrator().HasColumn(&model.Actor{}, "private_key")
}

// counts actors that still have a non-empty private_key in the database.
// returns 0 if the column has been dropped.
func CountLegacyKeys(db *gorm.DB) int64 {
	if !HasPrivateKeyColumn(db) {
		return 0
	}
	var count int64
	db.Raw("SELECT COUNT(*) FROM actors WHERE private_key != '' AND private_key IS NOT NULL").Scan(&count)
	return count
}

// drops the private_key column from the actors table.
// caller is responsible for confirming this is desired.
func DropPrivateKeyColumn(db *gorm.DB) error {
	if !HasPrivateKeyColumn(db) {
		return nil // already dropped
	}
	return db.Exec("ALTER TABLE actors DROP COLUMN private_key").Error
}
