package cmd

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/handler/wallet"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func swapWalletHandler(mockHandler wallet.Handler) func() {
	actual := wallet.Default
	wallet.Default = mockHandler
	return func() {
		wallet.Default = actual
	}
}

func TestWalletCreate(t *testing.T) {
	testutil.OneWithoutReset(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		runner := NewRunner()
		defer runner.Save(t)
		mockHandler := new(wallet.MockWallet)
		defer swapWalletHandler(mockHandler)()
		mockHandler.On("CreateHandler", mock.Anything, mock.Anything, mock.Anything).Return(&model.Wallet{
			ActorID:    "id",
			Address:    "address",
			PrivateKey: "private",
		}, nil)

		_, _, err := runner.Run(ctx, "singularity wallet create")
		require.NoError(t, err)

		_, _, err = runner.Run(ctx, "singularity --verbose wallet create")
		require.NoError(t, err)

		_, _, err = runner.Run(ctx, "singularity wallet create secp256k1")
		require.NoError(t, err)

		_, _, err = runner.Run(ctx, "singularity wallet create bls")
		require.NoError(t, err)
	})
}

func TestWalletCreate_BadType(t *testing.T) {
	testutil.OneWithoutReset(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		runner := NewRunner()
		defer runner.Save(t)
		mockHandler := new(wallet.MockWallet)
		defer swapWalletHandler(mockHandler)()
		mockHandler.On("CreateHandler", mock.Anything, mock.Anything, mock.Anything).Return((*model.Wallet)(nil), errors.New("unsupported key type: not-a-real-type"))

		_, _, err := runner.Run(ctx, "singularity wallet create not-a-real-type")
		require.Error(t, err)
	})
}

func TestWalletImport(t *testing.T) {
	testutil.OneWithoutReset(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		tmp := t.TempDir()
		err := os.WriteFile(filepath.Join(tmp, "private"), []byte("private"), 0644)
		require.NoError(t, err)

		runner := NewRunner()
		defer runner.Save(t)
		mockHandler := new(wallet.MockWallet)
		defer swapWalletHandler(mockHandler)()
		mockHandler.On("ImportHandler", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&model.Wallet{
			ActorID:    "id",
			Address:    "address",
			PrivateKey: "private",
		}, nil)

		_, _, err = runner.Run(ctx, "singularity wallet import "+testutil.EscapePath(filepath.Join(tmp, "private")))
		require.NoError(t, err)

		_, _, err = runner.Run(ctx, "singularity --verbose wallet import "+testutil.EscapePath(filepath.Join(tmp, "private")))
		require.NoError(t, err)
	})
}

func TestWalletInit(t *testing.T) {
	testutil.OneWithoutReset(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		runner := NewRunner()
		defer runner.Save(t)
		mockHandler := new(wallet.MockWallet)
		defer swapWalletHandler(mockHandler)()
		mockHandler.On("InitHandler", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&model.Wallet{
			ActorID:    "id",
			Address:    "address",
			PrivateKey: "private",
		}, nil)

		_, _, err := runner.Run(ctx, "singularity wallet init xxx")
		require.NoError(t, err)
	})
}

func TestWalletList(t *testing.T) {
	testutil.OneWithoutReset(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		runner := NewRunner()
		defer runner.Save(t)
		mockHandler := new(wallet.MockWallet)
		defer swapWalletHandler(mockHandler)()
		mockHandler.On("ListHandler", mock.Anything, mock.Anything).Return([]model.Wallet{{
			ActorID:    "id1",
			Address:    "address1",
			PrivateKey: "private1",
		}, {
			ActorID:    "id2",
			Address:    "address2",
			PrivateKey: "private2",
		}}, nil)

		_, _, err := runner.Run(ctx, "singularity wallet list")
		require.NoError(t, err)

		_, _, err = runner.Run(ctx, "singularity --verbose wallet list")
		require.NoError(t, err)
	})
}

func TestWalletRemove(t *testing.T) {
	testutil.OneWithoutReset(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		runner := NewRunner()
		defer runner.Save(t)
		mockHandler := new(wallet.MockWallet)
		defer swapWalletHandler(mockHandler)()
		mockHandler.On("RemoveHandler", mock.Anything, mock.Anything, mock.Anything).Return(nil)

		_, _, err := runner.Run(ctx, "singularity wallet remove --really-do-it xxx")
		require.NoError(t, err)
	})
}

func TestWalletRemove_NoReallyDoIt(t *testing.T) {
	testutil.OneWithoutReset(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		runner := NewRunner()
		defer runner.Save(t)
		mockHandler := new(wallet.MockWallet)
		defer swapWalletHandler(mockHandler)()
		mockHandler.On("RemoveHandler", mock.Anything, mock.Anything, mock.Anything).Return(nil)

		_, _, err := runner.Run(ctx, "singularity wallet remove xxx")
		require.ErrorIs(t, err, cliutil.ErrReallyDoIt)
	})
}

func TestWalletUpdate(t *testing.T) {
	testutil.OneWithoutReset(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		runner := NewRunner()
		defer runner.Save(t)
		mockHandler := new(wallet.MockWallet)
		defer swapWalletHandler(mockHandler)()
		mockHandler.On("UpdateHandler", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&model.Wallet{
			ActorID:     "id",
			ActorName:   "Updated Name",
			Address:     "address",
			ContactInfo: "test@example.com",
			Location:    "US-East",
			WalletType:  model.SPWallet,
		}, nil)

		_, _, err := runner.Run(ctx, "singularity wallet update --name Updated --contact test@example.com --location US-East address")
		require.NoError(t, err)

		_, _, err = runner.Run(ctx, "singularity --verbose wallet update --name Updated address")
		require.NoError(t, err)
	})
}

func TestWalletUpdate_NoAddress(t *testing.T) {
	testutil.OneWithoutReset(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		runner := NewRunner()
		defer runner.Save(t)

		_, _, err := runner.Run(ctx, "singularity wallet update --name Test")
		require.Error(t, err)
		require.Contains(t, err.Error(), "incorrect number of arguments")
	})
}

func TestWalletUpdate_NoFields(t *testing.T) {
	testutil.OneWithoutReset(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		runner := NewRunner()
		defer runner.Save(t)

		_, _, err := runner.Run(ctx, "singularity wallet update address")
		require.Error(t, err)
		require.Contains(t, err.Error(), "at least one field must be provided for update")
	})
}
