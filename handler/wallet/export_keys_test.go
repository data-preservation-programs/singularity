package wallet

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/keystore"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestExportKeysHandler_ExportsActorKey(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		ks, err := keystore.NewLocalKeyStore(t.TempDir())
		require.NoError(t, err)

		// create an actor with a legacy private key
		actor := model.Actor{
			ID:         "f01234",
			Address:    testutil.TestWalletAddr,
			PrivateKey: testutil.TestPrivateKeyHex,
		}
		require.NoError(t, db.Create(&actor).Error)

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

		// create actor with legacy key
		actor := model.Actor{
			ID:         "f01234",
			Address:    testutil.TestWalletAddr,
			PrivateKey: testutil.TestPrivateKeyHex,
		}
		require.NoError(t, db.Create(&actor).Error)

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

		// create actor with legacy key
		actor := model.Actor{
			ID:         "f01234",
			Address:    testutil.TestWalletAddr,
			PrivateKey: testutil.TestPrivateKeyHex,
		}
		require.NoError(t, db.Create(&actor).Error)

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

		actor := model.Actor{
			ID:         "f01234",
			Address:    testutil.TestWalletAddr,
			PrivateKey: testutil.TestPrivateKeyHex,
		}
		require.NoError(t, db.Create(&actor).Error)

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

		// actor with no private key
		actor := model.Actor{
			ID:      "f09999",
			Address: "f1abc",
		}
		require.NoError(t, db.Create(&actor).Error)

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

		// actor with garbage key
		actor := model.Actor{
			ID:         "f05555",
			Address:    "f1bad",
			PrivateKey: "not-a-valid-key",
		}
		require.NoError(t, db.Create(&actor).Error)

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

		// create actor with legacy key
		actor := model.Actor{
			ID:         "f01234",
			Address:    testutil.TestWalletAddr,
			PrivateKey: testutil.TestPrivateKeyHex,
		}
		require.NoError(t, db.Create(&actor).Error)

		// create wallet record pointing to a nonexistent key file
		require.NoError(t, db.Create(&model.Wallet{
			KeyPath:  "/nonexistent/key",
			KeyStore: "local",
			Address:  testutil.TestWalletAddr,
		}).Error)

		result, err := ExportKeysHandler(ctx, db, ks)
		require.NoError(t, err)
		require.Equal(t, 0, result.Exported)
		require.Equal(t, 0, result.Skipped)
		require.Len(t, result.Errors, 1)
		require.Contains(t, result.Errors[0], "key file missing")
	})
}

func TestDropPrivateKeyColumn(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		// column exists after AutoMigrate
		require.True(t, db.Migrator().HasColumn(&model.Actor{}, "private_key"))

		// drop it
		require.NoError(t, DropPrivateKeyColumn(db))
		require.False(t, db.Migrator().HasColumn(&model.Actor{}, "private_key"))

		// idempotent -- second call is a no-op
		require.NoError(t, DropPrivateKeyColumn(db))
	})
}

func TestDropPrivateKeyColumn_ExportThenDrop(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		ks, err := keystore.NewLocalKeyStore(t.TempDir())
		require.NoError(t, err)

		// create actor with legacy key
		actor := model.Actor{
			ID:         "f01234",
			Address:    testutil.TestWalletAddr,
			PrivateKey: testutil.TestPrivateKeyHex,
		}
		require.NoError(t, db.Create(&actor).Error)

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
