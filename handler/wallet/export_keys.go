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

// actor row with the legacy private_key column, used only by export-keys.
// Actor.PrivateKey is gorm:"-" in the model so we need a local type for raw SQL.
type legacyActorRow struct {
	ID         string
	Address    string
	PrivateKey string
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

	var actors []legacyActorRow
	err := db.Raw("SELECT id, address, private_key FROM actors WHERE private_key != '' AND private_key IS NOT NULL").Scan(&actors).Error
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

	// after exporting keys, migrate wallet_assignments if the legacy table exists
	if err := migrateWalletAssignments(db); err != nil {
		result.Errors = append(result.Errors, fmt.Sprintf("wallet_assignments migration: %v", err))
	}

	return result, nil
}

// migrates legacy wallet_assignments (preparation_id, wallet_id=actor_id_string)
// to preparation.wallet_id (uint FK to new wallets table). drops the table after.
func migrateWalletAssignments(db *gorm.DB) error {
	if !db.Migrator().HasTable("wallet_assignments") {
		return nil
	}

	type row struct {
		PreparationID uint
		WalletID      string // old actor ID (f0...)
	}
	var rows []row
	if err := db.Raw("SELECT preparation_id, wallet_id FROM wallet_assignments").Scan(&rows).Error; err != nil {
		return errors.Wrap(err, "failed to read wallet_assignments")
	}

	migrated := 0
	for _, r := range rows {
		// find the new wallet by actor_id
		var wallet model.Wallet
		if err := db.Where("actor_id = ?", r.WalletID).First(&wallet).Error; err != nil {
			logger.Warnw("wallet_assignment: no wallet found for actor, skipping",
				"preparation_id", r.PreparationID, "actor_id", r.WalletID)
			continue
		}
		if err := db.Exec("UPDATE preparations SET wallet_id = ? WHERE id = ? AND wallet_id IS NULL",
			wallet.ID, r.PreparationID).Error; err != nil {
			return errors.Wrapf(err, "failed to set wallet_id for preparation %d", r.PreparationID)
		}
		migrated++
	}

	if err := db.Migrator().DropTable("wallet_assignments"); err != nil {
		return errors.Wrap(err, "failed to drop wallet_assignments")
	}

	if migrated > 0 {
		logger.Infow("migrated wallet_assignments to preparation.wallet_id", "migrated", migrated)
	}
	return nil
}

// exports a single actor's key to keystore, returns (true, "") on success,
// (false, "") on skip, ("", errMsg) on failure
func exportOneKey(db *gorm.DB, ks keystore.KeyStore, actor legacyActorRow) (exported bool, errMsg string) {
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
		// verify the key file exists and contains a valid key for this address
		stored, err := ks.Get(existing.KeyPath)
		if err != nil {
			return false, fmt.Sprintf("actor %s: wallet record exists but key file unreadable at %s: %v", actor.ID, existing.KeyPath, err)
		}
		storedAddr, err := keystore.AddressFromExport(stored)
		if err != nil {
			return false, fmt.Sprintf("actor %s: key file at %s is corrupt: %v", actor.ID, existing.KeyPath, err)
		}
		if storedAddr != addr {
			return false, fmt.Sprintf("actor %s: key file at %s contains wrong address %s (expected %s)", actor.ID, existing.KeyPath, storedAddr, addr)
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
	// can't use db.Migrator().HasColumn(&model.Actor{}, "private_key") because
	// the field is gorm:"-" -- GORM won't resolve the column name. query directly.
	dialect := db.Dialector.Name()
	var count int64
	switch dialect {
	case "sqlite":
		db.Raw("SELECT COUNT(*) FROM pragma_table_info('actors') WHERE name = 'private_key'").Scan(&count)
	case "postgres":
		db.Raw("SELECT COUNT(*) FROM information_schema.columns WHERE table_schema = current_schema() AND table_name = 'actors' AND column_name = 'private_key'").Scan(&count)
	default: // mysql
		db.Raw("SELECT COUNT(*) FROM information_schema.columns WHERE table_schema = DATABASE() AND table_name = 'actors' AND column_name = 'private_key'").Scan(&count)
	}
	return count > 0
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
