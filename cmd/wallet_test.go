package cmd

import (
	"context"
	"testing"

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

func TestWalletImport(t *testing.T) {
	testutil.OneWithoutReset(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		runner := NewRunner()
		defer runner.Save(t)
		mockHandler := new(MockWallet)
		defer swapWalletHandler(mockHandler)()
		mockHandler.On("ImportHandler", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&model.Wallet{
			ID:         "id",
			Address:    "address",
			PrivateKey: "private",
		}, nil)
		_, _, err := runner.Run(ctx, "singularity wallet import xxx")
		require.NoError(t, err)
		_, _, err = runner.Run(ctx, "singularity --verbose wallet import xxx")
		require.NoError(t, err)
	})
}

func TestWalletList(t *testing.T) {
	testutil.OneWithoutReset(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		runner := NewRunner()
		defer runner.Save(t)
		mockHandler := new(MockWallet)
		defer swapWalletHandler(mockHandler)()
		mockHandler.On("ListHandler", mock.Anything, mock.Anything).Return([]model.Wallet{{
			ID:         "id1",
			Address:    "address1",
			PrivateKey: "private1",
		}, {
			ID:         "id2",
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
		mockHandler := new(MockWallet)
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
		mockHandler := new(MockWallet)
		defer swapWalletHandler(mockHandler)()
		mockHandler.On("RemoveHandler", mock.Anything, mock.Anything, mock.Anything).Return(nil)
		_, _, err := runner.Run(ctx, "singularity wallet remove xxx")
		require.ErrorIs(t, err, cliutil.ErrReallyDoIt)
	})
}
