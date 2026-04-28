package dealpusher

import (
	"context"
	"errors"
	"testing"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/ybbus/jsonrpc/v3"
	"gorm.io/gorm"
)

// rpcClientMock satisfies jsonrpc.RPCClient for the methods we exercise.
type rpcClientMock struct {
	mock.Mock
}

func (m *rpcClientMock) Call(ctx context.Context, method string, params ...any) (*jsonrpc.RPCResponse, error) {
	panic("not implemented")
}
func (m *rpcClientMock) CallRaw(ctx context.Context, request *jsonrpc.RPCRequest) (*jsonrpc.RPCResponse, error) {
	panic("not implemented")
}
func (m *rpcClientMock) CallBatch(ctx context.Context, requests jsonrpc.RPCRequests) (jsonrpc.RPCResponses, error) {
	panic("not implemented")
}
func (m *rpcClientMock) CallBatchRaw(ctx context.Context, requests jsonrpc.RPCRequests) (jsonrpc.RPCResponses, error) {
	panic("not implemented")
}
func (m *rpcClientMock) CallFor(ctx context.Context, out any, method string, params ...any) error {
	return m.Called(ctx, out, method, params).Error(0)
}

// f410 form of 0xE3e842B9D89ed2Ee3976b9b8916827302618c29e on testnet.
const (
	testProviderID   = "t0186503"
	testProviderF410 = "t410f4puefooyt3jo4olwxg4jc2bhgatbrqu6hqc73uy"
	testProviderEVM  = "0xE3e842B9D89ed2Ee3976b9b8916827302618c29e"
)

func TestResolveProviderEVMAddress_CacheHit(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		require.NoError(t, db.Create(&model.Actor{ID: testProviderID, Address: testProviderF410}).Error)

		d := &DealPusher{dbNoContext: db}
		evm, err := d.resolveProviderEVMAddress(ctx, testProviderID)
		require.NoError(t, err)
		require.Equal(t, common.HexToAddress(testProviderEVM), evm)
	})
}

func TestResolveProviderEVMAddress_OnchainFallback(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		rpc := new(rpcClientMock)
		rpc.On("CallFor", mock.Anything, mock.Anything, "Filecoin.StateLookupRobustAddress", mock.Anything).
			Run(func(args mock.Arguments) {
				out := args.Get(1).(*string)
				*out = testProviderF410
			}).
			Return(nil)

		d := &DealPusher{dbNoContext: db, lotusClient: rpc}
		evm, err := d.resolveProviderEVMAddress(ctx, testProviderID)
		require.NoError(t, err)
		require.Equal(t, common.HexToAddress(testProviderEVM), evm)

		// caches the lookup for subsequent calls
		var cached model.Actor
		require.NoError(t, db.Where("id = ?", testProviderID).First(&cached).Error)
		require.Equal(t, testProviderF410, cached.Address)
	})
}

func TestResolveProviderEVMAddress_OnchainFailure(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		rpc := new(rpcClientMock)
		rpc.On("CallFor", mock.Anything, mock.Anything, "Filecoin.StateLookupRobustAddress", mock.Anything).
			Return(errors.New("rpc unavailable"))

		d := &DealPusher{dbNoContext: db, lotusClient: rpc}
		_, err := d.resolveProviderEVMAddress(ctx, testProviderID)
		require.Error(t, err)
		require.ErrorContains(t, err, "failed to resolve robust address")
	})
}

func TestResolveProviderEVMAddress_NonDelegatedFromChain(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		rpc := new(rpcClientMock)
		rpc.On("CallFor", mock.Anything, mock.Anything, "Filecoin.StateLookupRobustAddress", mock.Anything).
			Run(func(args mock.Arguments) {
				out := args.Get(1).(*string)
				*out = testutil.TestWalletAddr
			}).
			Return(nil)

		d := &DealPusher{dbNoContext: db, lotusClient: rpc}
		_, err := d.resolveProviderEVMAddress(ctx, testProviderID)
		require.Error(t, err)
		require.ErrorContains(t, err, "not a delegated")
	})
}

func TestValidatePDPProofSetPieceSize(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name      string
		pieceSize int64
		wantError string
	}{
		{
			name:      "valid",
			pieceSize: 1 << 20,
		},
		{
			name:      "zero",
			pieceSize: 0,
			wantError: "must be greater than 0",
		},
		{
			name:      "negative",
			pieceSize: -1024,
			wantError: "must be greater than 0",
		},
		{
			name:      "non_power_of_two",
			pieceSize: 1536,
			wantError: "must be a power of two",
		},
		{
			name:      "too_large",
			pieceSize: 1 << 31,
			wantError: "exceeds max allowed",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			err := validatePDPProofSetPieceSize(tc.pieceSize)
			if tc.wantError == "" {
				require.NoError(t, err)
				return
			}
			require.Error(t, err)
			require.ErrorContains(t, err, tc.wantError)
		})
	}
}
