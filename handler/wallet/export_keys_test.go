package wallet

import (
	"context"
	"os"
	"testing"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/keystore"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

// model without -:migration, used to simulate a legacy database where
// AutoMigrate created the private_key column
type legacyActor struct {
	ID         string `gorm:"primaryKey;size:15"`
	Address    string `gorm:"index"`
	PrivateKey string
}

func (legacyActor) TableName() string { return "actors" }

// simulates a legacy database by running AutoMigrate with the old model
// that includes private_key as a normal column
func addLegacyColumn(t *testing.T, db *gorm.DB) {
	t.Helper()
	require.NoError(t, db.AutoMigrate(&legacyActor{}))
}

// inserts an actor with a legacy private_key via the legacy model
func createLegacyActor(t *testing.T, db *gorm.DB, id, address, privateKey string) {
	t.Helper()
	require.NoError(t, db.Create(&legacyActor{
		ID: id, Address: address, PrivateKey: privateKey,
	}).Error)
}

func TestExportKeysHandler_ExportsActorKey(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		ks, err := keystore.NewLocalKeyStore(t.TempDir())
		require.NoError(t, err)

		addLegacyColumn(t, db)
		createLegacyActor(t, db, "f01234", testutil.TestWalletAddr, testutil.TestPrivateKeyHex)

		result, err := ExportKeysHandler(ctx, db, ks)
		require.NoError(t, err)
		require.Equal(t, 1, result.Exported)
		require.Equal(t, 0, result.Skipped)
		require.Empty(t, result.Errors)

		// verify wallet was created and linked to actor
		var w model.Wallet
		require.NoError(t, db.Where("address = ?", testutil.TestWalletAddr).First(&w).Error)
		require.Equal(t, "local", w.KeyStore)
		require.NotNil(t, w.ActorID)
		require.Equal(t, "f01234", *w.ActorID)

		// verify key is readable from keystore
		key, err := ks.Get(w.KeyPath)
		require.NoError(t, err)
		require.Equal(t, testutil.TestPrivateKeyHex, key)
	})
}

func TestExportKeysHandler_SkipsExistingWallet(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		ks, err := keystore.NewLocalKeyStore(t.TempDir())
		require.NoError(t, err)

		addLegacyColumn(t, db)
		createLegacyActor(t, db, "f01234", testutil.TestWalletAddr, testutil.TestPrivateKeyHex)

		// pre-import the wallet via normal import path
		h := DefaultHandler{}
		_, err = h.ImportKeystoreHandler(ctx, db, ks, ImportKeystoreRequest{
			PrivateKey: testutil.TestPrivateKeyHex,
		})
		require.NoError(t, err)

		result, err := ExportKeysHandler(ctx, db, ks)
		require.NoError(t, err)
		require.Equal(t, 0, result.Exported)
		require.Equal(t, 1, result.Skipped)
		require.Empty(t, result.Errors)

		// verify still only one wallet record
		var count int64
		db.Model(&model.Wallet{}).Count(&count)
		require.Equal(t, int64(1), count)
	})
}

func TestExportKeysHandler_SkipsExistingWalletLinksActor(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		ks, err := keystore.NewLocalKeyStore(t.TempDir())
		require.NoError(t, err)

		addLegacyColumn(t, db)
		createLegacyActor(t, db, "f01234", testutil.TestWalletAddr, testutil.TestPrivateKeyHex)

		// pre-import the wallet WITHOUT actor linkage
		h := DefaultHandler{}
		_, err = h.ImportKeystoreHandler(ctx, db, ks, ImportKeystoreRequest{
			PrivateKey: testutil.TestPrivateKeyHex,
		})
		require.NoError(t, err)

		result, err := ExportKeysHandler(ctx, db, ks)
		require.NoError(t, err)
		require.Equal(t, 0, result.Exported)
		require.Equal(t, 1, result.Skipped)

		// verify actor was linked
		var w model.Wallet
		require.NoError(t, db.Where("address = ?", testutil.TestWalletAddr).First(&w).Error)
		require.NotNil(t, w.ActorID)
		require.Equal(t, "f01234", *w.ActorID)
	})
}

func TestExportKeysHandler_Idempotent(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		ks, err := keystore.NewLocalKeyStore(t.TempDir())
		require.NoError(t, err)

		addLegacyColumn(t, db)
		createLegacyActor(t, db, "f01234", testutil.TestWalletAddr, testutil.TestPrivateKeyHex)

		// first run exports
		r1, err := ExportKeysHandler(ctx, db, ks)
		require.NoError(t, err)
		require.Equal(t, 1, r1.Exported)

		// second run skips
		r2, err := ExportKeysHandler(ctx, db, ks)
		require.NoError(t, err)
		require.Equal(t, 0, r2.Exported)
		require.Equal(t, 1, r2.Skipped)

		// still only one wallet
		var count int64
		db.Model(&model.Wallet{}).Count(&count)
		require.Equal(t, int64(1), count)
	})
}

func TestExportKeysHandler_NoActorsWithKeys(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		ks, err := keystore.NewLocalKeyStore(t.TempDir())
		require.NoError(t, err)

		addLegacyColumn(t, db)
		// actor with no private key
		require.NoError(t, db.Create(&model.Actor{
			ID:      "f09999",
			Address: "f1abc",
		}).Error)

		result, err := ExportKeysHandler(ctx, db, ks)
		require.NoError(t, err)
		require.Equal(t, 0, result.Exported)
		require.Equal(t, 0, result.Skipped)
		require.Empty(t, result.Errors)
	})
}

func TestExportKeysHandler_InvalidKeyRecordsError(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		ks, err := keystore.NewLocalKeyStore(t.TempDir())
		require.NoError(t, err)

		addLegacyColumn(t, db)
		createLegacyActor(t, db, "f05555", "f1bad", "not-a-valid-key")

		result, err := ExportKeysHandler(ctx, db, ks)
		require.NoError(t, err) // overall operation succeeds
		require.Equal(t, 0, result.Exported)
		require.Equal(t, 0, result.Skipped)
		require.Len(t, result.Errors, 1)
		require.Contains(t, result.Errors[0], "f05555")
		require.Contains(t, result.Errors[0], "invalid key format")
	})
}

func TestExportKeysHandler_MissingKeyFile(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		ks, err := keystore.NewLocalKeyStore(t.TempDir())
		require.NoError(t, err)

		addLegacyColumn(t, db)
		createLegacyActor(t, db, "f01234", testutil.TestWalletAddr, testutil.TestPrivateKeyHex)

		// create wallet record pointing to a nonexistent key file
		require.NoError(t, db.Create(&model.Wallet{
			KeyPath:  "nonexistent",
			KeyStore: "local",
			Address:  testutil.TestWalletAddr,
		}).Error)

		result, err := ExportKeysHandler(ctx, db, ks)
		require.NoError(t, err)
		require.Equal(t, 0, result.Exported)
		require.Equal(t, 0, result.Skipped)
		require.Len(t, result.Errors, 1)
		require.Contains(t, result.Errors[0], "key file unreadable")
	})
}

func TestExportKeysHandler_CorruptKeyFile(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		dir := t.TempDir()
		ks, err := keystore.NewLocalKeyStore(dir)
		require.NoError(t, err)

		addLegacyColumn(t, db)
		createLegacyActor(t, db, "f01234", testutil.TestWalletAddr, testutil.TestPrivateKeyHex)

		// write garbage as a key file inside the keystore dir
		require.NoError(t, os.WriteFile(dir+"/corrupt", []byte("garbage"), 0600))

		require.NoError(t, db.Create(&model.Wallet{
			KeyPath:  "corrupt",
			KeyStore: "local",
			Address:  testutil.TestWalletAddr,
		}).Error)

		result, err := ExportKeysHandler(ctx, db, ks)
		require.NoError(t, err)
		require.Equal(t, 0, result.Exported)
		require.Equal(t, 0, result.Skipped)
		require.Len(t, result.Errors, 1)
		require.Contains(t, result.Errors[0], "corrupt")
	})
}

func TestExportKeysHandler_AutoMigrateDoesNotRecreateColumn(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		// AutoMigrate already ran (via testutil.All) -- column should NOT exist
		require.False(t, db.Migrator().HasColumn(&model.Actor{}, "private_key"),
			"AutoMigrate must not create private_key column (gorm:-:migration)")

		// simulate legacy db, export, drop
		addLegacyColumn(t, db)
		require.True(t, HasPrivateKeyColumn(db))
		require.NoError(t, DropPrivateKeyColumn(db))
		require.False(t, HasPrivateKeyColumn(db))

		// re-run AutoMigrate -- must NOT re-add the column
		require.NoError(t, model.AutoMigrate(db))
		require.False(t, HasPrivateKeyColumn(db),
			"AutoMigrate must not re-add private_key column after drop")
	})
}

func TestDropPrivateKeyColumn_ExportThenDrop(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		ks, err := keystore.NewLocalKeyStore(t.TempDir())
		require.NoError(t, err)

		addLegacyColumn(t, db)
		createLegacyActor(t, db, "f01234", testutil.TestWalletAddr, testutil.TestPrivateKeyHex)

		// export
		result, err := ExportKeysHandler(ctx, db, ks)
		require.NoError(t, err)
		require.Equal(t, 1, result.Exported)
		require.Empty(t, result.Errors)

		// drop column
		require.NoError(t, DropPrivateKeyColumn(db))

		// key is still in keystore
		var w model.Wallet
		require.NoError(t, db.Where("address = ?", testutil.TestWalletAddr).First(&w).Error)
		key, err := ks.Get(w.KeyPath)
		require.NoError(t, err)
		require.Equal(t, testutil.TestPrivateKeyHex, key)

		// actor record still exists, just without the key column
		var a model.Actor
		require.NoError(t, db.First(&a, "id = ?", "f01234").Error)
		require.Equal(t, testutil.TestWalletAddr, a.Address)
	})
}
