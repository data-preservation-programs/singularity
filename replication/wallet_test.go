package replication

import (
	"context"
	"testing"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/ybbus/jsonrpc/v3"
	"gorm.io/gorm"
)

type MockRPCClient struct {
	mock.Mock
}

func (m *MockRPCClient) Call(ctx context.Context, method string, params ...any) (*jsonrpc.RPCResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockRPCClient) CallRaw(ctx context.Context, request *jsonrpc.RPCRequest) (*jsonrpc.RPCResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockRPCClient) CallFor(ctx context.Context, out any, method string, params ...any) error {
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
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		lotusClient := new(MockRPCClient)

		// Set up the test data
		wallets := []model.Wallet{
			{ActorID: "1", Address: "address1"},
			{ActorID: "2", Address: "address2"},
			{ActorID: "3", Address: "address3"},
			{ActorID: "4", Address: "address4"},
		}

		// Set up expectations for the lotusClient mock
		lotusClient.On("CallFor", mock.Anything, mock.AnythingOfType("*string"), "Filecoin.StateMarketBalance", []any{"address1", nil}).
			Return(nil).Run(func(args mock.Arguments) {
			resultPtr := args.Get(1).(*string)
			*resultPtr = "1000000"
		})

		lotusClient.On("CallFor", mock.Anything, mock.AnythingOfType("*string"), "Filecoin.StateMarketBalance", []any{"address2", nil}).
			Return(errors.New("failed to get datacap"))

		lotusClient.On("CallFor", mock.Anything, mock.AnythingOfType("*string"), "Filecoin.StateMarketBalance", []any{"address3", nil}).
			Return(nil).Run(func(args mock.Arguments) {
			resultPtr := args.Get(1).(*string)
			*resultPtr = "1000000"
		})

		lotusClient.On("CallFor", mock.Anything, mock.AnythingOfType("*string"), "Filecoin.StateMarketBalance", []any{"address4", nil}).
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
	})
}

func TestRandomWalletChooser(t *testing.T) {
	chooser := &RandomWalletChooser{}
	ctx := context.Background()
	wallet, err := chooser.Choose(ctx, []model.Wallet{
		{ActorID: "1", Address: "address1"},
		{ActorID: "2", Address: "address2"},
	})
	require.NoError(t, err)
	require.Contains(t, wallet.Address, "address")
}
