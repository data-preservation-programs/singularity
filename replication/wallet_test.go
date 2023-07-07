package replication

import (
	"context"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/ybbus/jsonrpc/v3"
)

type MockRPCClient struct {
	mock.Mock
}

func (m *MockRPCClient) Call(ctx context.Context, method string, params ...interface{}) (*jsonrpc.RPCResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockRPCClient) CallRaw(ctx context.Context, request *jsonrpc.RPCRequest) (*jsonrpc.RPCResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockRPCClient) CallFor(ctx context.Context, out interface{}, method string, params ...interface{}) error {
	return m.Called(ctx, out, method, params).Error(0)
}

func (m *MockRPCClient) CallBatch(ctx context.Context, requests jsonrpc.RPCRequests) (jsonrpc.RPCResponses, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockRPCClient) CallBatchRaw(ctx context.Context, requests jsonrpc.RPCRequests) (jsonrpc.RPCResponses, error) {
	//TODO implement me
	panic("implement me")
}

func TestDatacapWalletChooser_Choose(t *testing.T) {
	db := database.OpenInMemory()
	lotusClient := new(MockRPCClient)

	// Set up the test data
	wallets := []model.Wallet{
		{ID: "1", Address: "address1"},
		{ID: "2", Address: "address2"},
		{ID: "3", Address: "address3"},
		{ID: "4", Address: "address4"},
	}

	// Set up expectations for the lotusClient mock
	lotusClient.On("CallFor", mock.Anything, mock.AnythingOfType("*string"), "Filecoin.StateMarketBalance", []interface{}{"address1", nil}).
		Return(nil).Run(func(args mock.Arguments) {
		resultPtr := args.Get(1).(*string)
		*resultPtr = "1000000"
	})

	lotusClient.On("CallFor", mock.Anything, mock.AnythingOfType("*string"), "Filecoin.StateMarketBalance", []interface{}{"address2", nil}).
		Return(errors.New("failed to get datacap"))

	lotusClient.On("CallFor", mock.Anything, mock.AnythingOfType("*string"), "Filecoin.StateMarketBalance", []interface{}{"address3", nil}).
		Return(nil).Run(func(args mock.Arguments) {
		resultPtr := args.Get(1).(*string)
		*resultPtr = "1000000"
	})

	lotusClient.On("CallFor", mock.Anything, mock.AnythingOfType("*string"), "Filecoin.StateMarketBalance", []interface{}{"address4", nil}).
		Return(nil).Run(func(args mock.Arguments) {
		resultPtr := args.Get(1).(*string)
		*resultPtr = "900000"
	})

	chooser := NewDatacapWalletChooser(db, time.Minute, "lotusAPI", "lotusToken", 900001)
	chooser.lotusClient = lotusClient

	err := db.Create(&wallets).Error
	require.NoError(t, err)
	err = db.Create(&model.Deal{
		ClientID:  "3",
		Verified:  true,
		State:     model.DealProposed,
		PieceSize: 500000,
	}).Error
	require.NoError(t, err)

	t.Run("Choose wallet with empty wallet", func(t *testing.T) {
		_, err := chooser.Choose(context.Background(), []model.Wallet{})
		require.ErrorAs(t, err, &ErrNoWallet)
	})

	t.Run("Choose wallet with sufficient datacap", func(t *testing.T) {
		chosenWallet, err := chooser.Choose(context.Background(), []model.Wallet{wallets[0], wallets[1]})
		require.NoError(t, err)
		require.Equal(t, "address1", chosenWallet.Address)
	})

	t.Run("Choose wallet with insufficient datacap", func(t *testing.T) {
		_, err := chooser.Choose(context.Background(), []model.Wallet{wallets[2], wallets[3]})
		require.ErrorAs(t, err, &ErrNoDatacap)
	})
}
