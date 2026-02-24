package wallet

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/util/keystore"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestImportKeystoreHandler_Success(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		ks, err := keystore.NewLocalKeyStore(t.TempDir())
		require.NoError(t, err)

		h := DefaultHandler{}
		w, err := h.ImportKeystoreHandler(ctx, db, ks, ImportKeystoreRequest{
			PrivateKey: testutil.TestPrivateKeyHex,
			Name:       "test-wallet",
		})
		require.NoError(t, err)
		require.NotNil(t, w)
		require.Equal(t, testutil.TestWalletAddr, w.Address)
		require.Equal(t, "test-wallet", w.Name)
		require.Equal(t, "local", w.KeyStore)
		require.NotEmpty(t, w.KeyPath)
		require.Nil(t, w.ActorID)
		require.True(t, ks.Has(w.KeyPath))
	})
}

func TestImportKeystoreHandler_NoName(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		ks, err := keystore.NewLocalKeyStore(t.TempDir())
		require.NoError(t, err)

		h := DefaultHandler{}
		w, err := h.ImportKeystoreHandler(ctx, db, ks, ImportKeystoreRequest{
			PrivateKey: testutil.TestPrivateKeyHex,
		})
		require.NoError(t, err)
		require.Equal(t, "", w.Name)
		require.Equal(t, testutil.TestWalletAddr, w.Address)
	})
}

func TestImportKeystoreHandler_Duplicate(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		ks, err := keystore.NewLocalKeyStore(t.TempDir())
		require.NoError(t, err)

		h := DefaultHandler{}
		_, err = h.ImportKeystoreHandler(ctx, db, ks, ImportKeystoreRequest{
			PrivateKey: testutil.TestPrivateKeyHex,
		})
		require.NoError(t, err)

		// second import of same key should fail
		_, err = h.ImportKeystoreHandler(ctx, db, ks, ImportKeystoreRequest{
			PrivateKey: testutil.TestPrivateKeyHex,
		})
		require.ErrorIs(t, err, handlererror.ErrDuplicateRecord)
	})
}

func TestImportKeystoreHandler_InvalidKey(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		ks, err := keystore.NewLocalKeyStore(t.TempDir())
		require.NoError(t, err)

		h := DefaultHandler{}
		_, err = h.ImportKeystoreHandler(ctx, db, ks, ImportKeystoreRequest{
			PrivateKey: "not-a-valid-key",
		})
		require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
	})
}

func TestImportKeystoreHandler_EmptyKey(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		ks, err := keystore.NewLocalKeyStore(t.TempDir())
		require.NoError(t, err)

		h := DefaultHandler{}
		_, err = h.ImportKeystoreHandler(ctx, db, ks, ImportKeystoreRequest{
			PrivateKey: "",
		})
		require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
	})
}
