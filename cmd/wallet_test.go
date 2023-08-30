package cmd

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/handler/wallet"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func swapWalletHandler(mockHandler wallet.Handler) func() {
	actual := wallet.Default
	wallet.Default = mockHandler
	return func() {
		wallet.Default = actual
	}
}

func TestWalletImport(t *testing.T) {
	mockHandler := new(MockWallet)
	defer swapWalletHandler(mockHandler)()
	mockHandler.On("ImportHandler", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&model.Wallet{
		ID:         "id",
		Address:    "address",
		PrivateKey: "private",
	}, nil)
	out, _, err := Run(context.Background(), "singularity wallet import xxx")
	require.NoError(t, err)
	Save(t, out, "wallet_import.txt")
	out, _, err = Run(context.Background(), "singularity --verbose wallet import xxx")
	require.NoError(t, err)
	Save(t, out, "wallet_import_verbose.txt")
}

func TestWalletList(t *testing.T) {
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
	out, _, err := Run(context.Background(), "singularity wallet list")
	require.NoError(t, err)
	Save(t, out, "wallet_list.txt")
	out, _, err = Run(context.Background(), "singularity --verbose wallet list")
	require.NoError(t, err)
	Save(t, out, "wallet_list_verbose.txt")
}

func TestWalletRemove(t *testing.T) {
	mockHandler := new(MockWallet)
	defer swapWalletHandler(mockHandler)()
	mockHandler.On("RemoveHandler", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	_, _, err := Run(context.Background(), "singularity wallet remove --really-do-it xxx")
	require.NoError(t, err)
}

func TestWalletRemove_NoReallyDoIt(t *testing.T) {
	mockHandler := new(MockWallet)
	defer swapWalletHandler(mockHandler)()
	mockHandler.On("RemoveHandler", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	_, _, err := Run(context.Background(), "singularity wallet remove xxx")
	require.ErrorIs(t, err, cliutil.ErrReallyDoIt)
}
