package wallet

import (
	"context"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/ybbus/jsonrpc/v3"
	"gorm.io/gorm"
)

// MockLotusClient is a mock implementation of jsonrpc.RPCClient
type MockLotusClient struct {
	mock.Mock
}

func (m *MockLotusClient) Call(ctx context.Context, method string, params ...interface{}) (*jsonrpc.RPCResponse, error) {
	args := m.Called(ctx, method, params)
	return args.Get(0).(*jsonrpc.RPCResponse), args.Error(1)
}

func (m *MockLotusClient) CallFor(ctx context.Context, out interface{}, method string, params ...interface{}) error {
	args := m.Called(ctx, out, method, params)
	return args.Error(0)
}

func (m *MockLotusClient) CallBatch(ctx context.Context, requests jsonrpc.RPCRequests) (jsonrpc.RPCResponses, error) {
	args := m.Called(ctx, requests)
	return args.Get(0).(jsonrpc.RPCResponses), args.Error(1)
}

func (m *MockLotusClient) CallBatchFor(ctx context.Context, out []interface{}, requests jsonrpc.RPCRequests) error {
	args := m.Called(ctx, out, requests)
	return args.Error(0)
}

func (m *MockLotusClient) CallBatchRaw(ctx context.Context, requests jsonrpc.RPCRequests) (jsonrpc.RPCResponses, error) {
	args := m.Called(ctx, requests)
	return args.Get(0).(jsonrpc.RPCResponses), args.Error(1)
}

func (m *MockLotusClient) CallRaw(ctx context.Context, request *jsonrpc.RPCRequest) (*jsonrpc.RPCResponse, error) {
	args := m.Called(ctx, request)
	return args.Get(0).(*jsonrpc.RPCResponse), args.Error(1)
}

func TestGetBalanceHandler(t *testing.T) {
	ctx := context.Background()
	mockClient := new(MockLotusClient)

	// Test valid wallet address with balance and datacap
	t.Run("Valid wallet with balance and datacap", func(t *testing.T) {
		// Mock WalletBalance call - using mock.Anything for variadic parameters
		mockClient.On("CallFor", ctx, mock.AnythingOfType("*string"), "Filecoin.WalletBalance", mock.Anything).
			Return(nil).Run(func(args mock.Arguments) {
			resultPtr := args.Get(1).(*string)
			*resultPtr = "1000000000000000000" // 1 FIL in attoFIL
		})

		// Mock StateVerifiedClientStatus call - using mock.Anything for variadic parameters
		mockClient.On("CallFor", ctx, mock.AnythingOfType("*string"), "Filecoin.StateVerifiedClientStatus", mock.Anything, mock.Anything).
			Return(nil).Run(func(args mock.Arguments) {
			resultPtr := args.Get(1).(*string)
			*resultPtr = "1099511627776" // 1024 GiB in bytes
		})

		handler := DefaultHandler{}
		result, err := handler.GetBalanceHandler(ctx, &gorm.DB{}, mockClient, "f12syf7zd3lfsv43aj2kb454ymaqw7debhumjnbqa")

		require.NoError(t, err)
		assert.Equal(t, "f12syf7zd3lfsv43aj2kb454ymaqw7debhumjnbqa", result.Address)
		assert.Equal(t, "1.000000 FIL", result.Balance)
		assert.Equal(t, "1000000000000000000", result.BalanceAttoFIL)
		assert.Equal(t, "1024.00 GiB", result.DataCap)
		assert.Equal(t, int64(1099511627776), result.DataCapBytes)

		mockClient.AssertExpectations(t)
		mockClient.ExpectedCalls = nil
	})

	// Test wallet with zero balance and no datacap
	t.Run("Wallet with zero balance and no datacap", func(t *testing.T) {
		// Mock WalletBalance call
		mockClient.On("CallFor", ctx, mock.AnythingOfType("*string"), "Filecoin.WalletBalance", mock.Anything).
			Return(nil).Run(func(args mock.Arguments) {
			resultPtr := args.Get(1).(*string)
			*resultPtr = "0" // 0 FIL
		})

		// Mock StateVerifiedClientStatus call returning no datacap
		mockClient.On("CallFor", ctx, mock.AnythingOfType("*string"), "Filecoin.StateVerifiedClientStatus", mock.Anything, mock.Anything).
			Return(nil).Run(func(args mock.Arguments) {
			resultPtr := args.Get(1).(*string)
			*resultPtr = "0" // No datacap
		})

		handler := DefaultHandler{}
		result, err := handler.GetBalanceHandler(ctx, &gorm.DB{}, mockClient, "f12syf7zd3lfsv43aj2kb454ymaqw7debhumjnbqa")

		require.NoError(t, err)
		assert.Equal(t, "f12syf7zd3lfsv43aj2kb454ymaqw7debhumjnbqa", result.Address)
		assert.Equal(t, "0.000000 FIL", result.Balance)
		assert.Equal(t, "0", result.BalanceAttoFIL)
		assert.Equal(t, "0.00 GiB", result.DataCap)
		assert.Equal(t, int64(0), result.DataCapBytes)

		mockClient.AssertExpectations(t)
		mockClient.ExpectedCalls = nil
	})

	// Test invalid wallet address
	t.Run("Invalid wallet address", func(t *testing.T) {
		handler := DefaultHandler{}
		_, err := handler.GetBalanceHandler(ctx, &gorm.DB{}, mockClient, "invalid-address")

		require.Error(t, err)
		assert.Contains(t, err.Error(), "invalid wallet address format")
	})

	// Test API error handling for balance lookup
	t.Run("Balance API error", func(t *testing.T) {
		mockClient.On("CallFor", ctx, mock.AnythingOfType("*string"), "Filecoin.WalletBalance", mock.Anything).
			Return(assert.AnError)

		// Since the implementation still tries to get datacap even when balance fails, we need to mock this too
		mockClient.On("CallFor", ctx, mock.AnythingOfType("*string"), "Filecoin.StateVerifiedClientStatus", mock.Anything, mock.Anything).
			Return(nil).Run(func(args mock.Arguments) {
			resultPtr := args.Get(1).(*string)
			*resultPtr = "null"
		})

		handler := DefaultHandler{}
		result, err := handler.GetBalanceHandler(ctx, &gorm.DB{}, mockClient, "f12syf7zd3lfsv43aj2kb454ymaqw7debhumjnbqa")

		require.NoError(t, err) // Function should not return error, but include error in response
		assert.NotNil(t, result.Error)
		assert.Contains(t, *result.Error, "failed to get wallet balance")

		// Should still have datacap info since that call succeeds
		assert.Equal(t, "0.00 GiB", result.DataCap)
		assert.Equal(t, int64(0), result.DataCapBytes)

		mockClient.AssertExpectations(t)
		mockClient.ExpectedCalls = nil
	})

	// Test API error handling for datacap lookup
	t.Run("Datacap API error", func(t *testing.T) {
		// Mock successful balance call
		mockClient.On("CallFor", ctx, mock.AnythingOfType("*string"), "Filecoin.WalletBalance", mock.Anything).
			Return(nil).Run(func(args mock.Arguments) {
			resultPtr := args.Get(1).(*string)
			*resultPtr = "1000000000000000000"
		})

		// Mock failed datacap call
		mockClient.On("CallFor", ctx, mock.AnythingOfType("*string"), "Filecoin.StateVerifiedClientStatus", mock.Anything, mock.Anything).
			Return(assert.AnError)

		handler := DefaultHandler{}
		result, err := handler.GetBalanceHandler(ctx, &gorm.DB{}, mockClient, "f12syf7zd3lfsv43aj2kb454ymaqw7debhumjnbqa")

		require.NoError(t, err) // Function should not return error, but include error in response
		assert.NotNil(t, result.Error)
		assert.Contains(t, *result.Error, "failed to get datacap balance")

		// Should still have balance info since that call succeeds
		assert.Equal(t, "1.000000 FIL", result.Balance)
		assert.Equal(t, "1000000000000000000", result.BalanceAttoFIL)

		// Datacap should be zero due to error
		assert.Equal(t, "0.00 GiB", result.DataCap)
		assert.Equal(t, int64(0), result.DataCapBytes)

		mockClient.AssertExpectations(t)
		mockClient.ExpectedCalls = nil
	})
}

func TestFormatFILFromAttoFIL(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"Zero balance", "0", "0.000000 FIL"},
		{"One FIL", "1000000000000000000", "1.000000 FIL"},
		{"Half FIL", "500000000000000000", "0.500000 FIL"},
		{"Small amount", "1000000000000000", "0.001000 FIL"},
		{"Large amount", "10000000000000000000", "10.000000 FIL"},
		{"Fractional amount", "1500000000000000000", "1.500000 FIL"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var amount big.Int
			amount.SetString(tc.input, 10)
			result := formatFILFromAttoFIL(&amount)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestFormatDatacap(t *testing.T) {
	testCases := []struct {
		name     string
		input    int64
		expected string
	}{
		{"Zero datacap", 0, "0.00 GiB"},
		{"One GiB", 1073741824, "1.00 GiB"},
		{"Half GiB", 536870912, "0.50 GiB"},
		{"Multiple GiB", 10737418240, "10.00 GiB"},
		{"Large datacap", 1099511627776, "1024.00 GiB"}, // 1 TiB
		{"Fractional GiB", 1610612736, "1.50 GiB"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := formatDatacap(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}
