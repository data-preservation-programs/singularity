package wallet

import (
	"context"
	"math/big"
	"strings"
	"testing"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/require"
	"github.com/ybbus/jsonrpc/v3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// MockLotusClient for testing
type MockLotusClient struct {
	balances map[string]string
	errors   map[string]error
}

func NewMockLotusClient() *MockLotusClient {
	return &MockLotusClient{
		balances: make(map[string]string),
		errors:   make(map[string]error),
	}
}

func (m *MockLotusClient) SetBalance(address string, balance string) {
	m.balances[address] = balance
	// Also set for the other address format (f1 <-> t1)
	if strings.HasPrefix(address, "f1") {
		tAddr := "t1" + address[2:]
		m.balances[tAddr] = balance
	} else if strings.HasPrefix(address, "t1") {
		fAddr := "f1" + address[2:]
		m.balances[fAddr] = balance
	}
}

func (m *MockLotusClient) SetError(address string, err error) {
	m.errors[address] = err
	// Also set for the other address format (f1 <-> t1)
	if strings.HasPrefix(address, "f1") {
		tAddr := "t1" + address[2:]
		m.errors[tAddr] = err
	} else if strings.HasPrefix(address, "t1") {
		fAddr := "f1" + address[2:]
		m.errors[fAddr] = err
	}
}

func (m *MockLotusClient) CallFor(ctx context.Context, result interface{}, method string, args ...interface{}) error {
	if method == "Filecoin.WalletBalance" && len(args) > 0 {
		if addr, ok := args[0].(address.Address); ok {
			addrStr := addr.String()
			
			// Check errors first
			if err, exists := m.errors[addrStr]; exists {
				return err
			}
			// Check balances
			if balance, exists := m.balances[addrStr]; exists {
				if strPtr, ok := result.(*string); ok {
					*strPtr = balance
					return nil
				}
			}
			// Return error for unset addresses to simulate wallet not found
			return errors.New("wallet not found")
		}
	}
	return errors.New("method not supported")
}

func (m *MockLotusClient) Call(ctx context.Context, method string, args ...interface{}) (*jsonrpc.RPCResponse, error) {
	return nil, nil
}

func (m *MockLotusClient) CallRaw(ctx context.Context, request *jsonrpc.RPCRequest) (*jsonrpc.RPCResponse, error) {
	return nil, nil
}

func (m *MockLotusClient) CallBatch(ctx context.Context, requests jsonrpc.RPCRequests) (jsonrpc.RPCResponses, error) {
	return nil, nil
}

func (m *MockLotusClient) CallBatchRaw(ctx context.Context, requests jsonrpc.RPCRequests) (jsonrpc.RPCResponses, error) {
	return nil, nil
}

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err)
	
	err = db.AutoMigrate(&model.Notification{}, &model.Deal{})
	require.NoError(t, err)
	
	return db
}

func TestCalculateRequiredBalance(t *testing.T) {
	validator := NewBalanceValidator()

	tests := []struct {
		name              string
		pricePerGBEpoch   float64
		pricePerGB        float64
		pricePerDeal      float64
		totalSizeBytes    int64
		durationEpochs    int64
		numberOfDeals     int
		expectedAttoFIL   string
	}{
		{
			name:              "only price per GB epoch",
			pricePerGBEpoch:   0.1,
			pricePerGB:        0,
			pricePerDeal:      0,
			totalSizeBytes:    1073741824, // 1 GB
			durationEpochs:    2880,       // ~1 day
			numberOfDeals:     1,
			expectedAttoFIL:   "288000000000000000000", // 0.1 * 1 * 2880 * 10^18
		},
		{
			name:              "only price per GB",
			pricePerGBEpoch:   0,
			pricePerGB:        1.0,
			pricePerDeal:      0,
			totalSizeBytes:    2147483648, // 2 GB
			durationEpochs:    0,
			numberOfDeals:     1,
			expectedAttoFIL:   "2000000000000000000", // 1.0 * 2 * 10^18
		},
		{
			name:              "only price per deal",
			pricePerGBEpoch:   0,
			pricePerGB:        0,
			pricePerDeal:      0.5,
			totalSizeBytes:    0,
			durationEpochs:    0,
			numberOfDeals:     3,
			expectedAttoFIL:   "1500000000000000000", // 0.5 * 3 * 10^18
		},
		{
			name:              "combined pricing",
			pricePerGBEpoch:   0.01,
			pricePerGB:        0.1,
			pricePerDeal:      0.05,
			totalSizeBytes:    1073741824, // 1 GB
			durationEpochs:    100,
			numberOfDeals:     2,
			expectedAttoFIL:   "1200000000000000000", // (0.01*1*100 + 0.1*1 + 0.05*2) * 10^18 = 1.2 * 10^18
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := validator.CalculateRequiredBalance(
				tt.pricePerGBEpoch,
				tt.pricePerGB,
				tt.pricePerDeal,
				tt.totalSizeBytes,
				tt.durationEpochs,
				tt.numberOfDeals,
			)
			expected, _ := new(big.Int).SetString(tt.expectedAttoFIL, 10)
			// Allow for small floating point precision errors (less than 1000 attoFIL)
			diff := new(big.Int).Sub(result, expected)
			diff.Abs(diff)
			maxDiff := big.NewInt(1000)
			require.True(t, diff.Cmp(maxDiff) <= 0, "Expected %s, got %s, diff: %s", tt.expectedAttoFIL, result.String(), diff.String())
		})
	}
}

func TestValidateWalletExists(t *testing.T) {
	db := setupTestDB(t)
	validator := NewBalanceValidator()
	ctx := context.Background()

	t.Run("valid wallet with balance", func(t *testing.T) {
		mockClient := NewMockLotusClient()
		walletAddr := "f1fib3pv7jua2ockdugtz7viz3cyy6lkhh7rfx3sa"
		mockClient.SetBalance(walletAddr, "1000000000000000000") // 1 FIL

		result, err := validator.ValidateWalletExists(ctx, db, mockClient, walletAddr, "prep1")
		require.NoError(t, err)
		require.True(t, result.IsValid)
		require.Equal(t, walletAddr, result.WalletAddress)
		require.Equal(t, "1 FIL", result.CurrentBalance)
		require.Equal(t, "Wallet exists and is accessible", result.Message)
	})

	t.Run("invalid wallet address format", func(t *testing.T) {
		mockClient := NewMockLotusClient()
		
		result, err := validator.ValidateWalletExists(ctx, db, mockClient, "invalid-address", "prep1")
		require.NoError(t, err)
		require.False(t, result.IsValid)
		require.Equal(t, "Invalid wallet address format", result.Message)
	})

	t.Run("wallet not found", func(t *testing.T) {
		mockClient := NewMockLotusClient()
		walletAddr := "f1abjxfbp274xpdqcpuaykwkfb43omjotacm2p3za"
		// Don't set balance or error - should trigger "wallet not found" in mock

		result, err := validator.ValidateWalletExists(ctx, db, mockClient, walletAddr, "prep1")
		require.NoError(t, err)
		require.False(t, result.IsValid)
		require.Equal(t, "Wallet not found or not accessible", result.Message)
	})
}

func TestValidateWalletBalance(t *testing.T) {
	db := setupTestDB(t)
	validator := NewBalanceValidator()
	ctx := context.Background()

	t.Run("sufficient balance", func(t *testing.T) {
		mockClient := NewMockLotusClient()
		walletAddr := "f1fib3pv7jua2ockdugtz7viz3cyy6lkhh7rfx3sa"
		mockClient.SetBalance(walletAddr, "2000000000000000000") // 2 FIL
		
		requiredAmount := big.NewInt(1000000000000000000) // 1 FIL

		result, err := validator.ValidateWalletBalance(ctx, db, mockClient, walletAddr, requiredAmount, "prep1")
		require.NoError(t, err)
		require.True(t, result.IsValid)
		require.Equal(t, "2 FIL", result.CurrentBalance)
		require.Equal(t, "1 FIL", result.RequiredBalance)
		require.Equal(t, "Wallet has sufficient balance for deal", result.Message)
	})

	t.Run("insufficient balance", func(t *testing.T) {
		mockClient := NewMockLotusClient()
		walletAddr := "f1fib3pv7jua2ockdugtz7viz3cyy6lkhh7rfx3sa"
		mockClient.SetBalance(walletAddr, "500000000000000000") // 0.5 FIL
		
		requiredAmount := big.NewInt(1000000000000000000) // 1 FIL

		result, err := validator.ValidateWalletBalance(ctx, db, mockClient, walletAddr, requiredAmount, "prep1")
		require.NoError(t, err)
		require.False(t, result.IsValid)
		require.Equal(t, "0.5 FIL", result.CurrentBalance)
		require.Equal(t, "1 FIL", result.RequiredBalance)
		require.Contains(t, result.Message, "Insufficient wallet balance")
	})

	t.Run("with pending deals", func(t *testing.T) {
		// Create some pending deals
		// Create valid CIDs
		cid1, _ := cid.Parse("bafybeigdyrzt5sfp7udm7hu76uh7y26nf3efuylqabf3oclgtqy55fbzdi")
		cid2, _ := cid.Parse("bafybeihdwdcefgh4dqkjv67uzcmw7ojee6xedzdetojuzjevtenxquvyku")
		
		deal1 := &model.Deal{
			ClientID:  "f1fib3pv7jua2ockdugtz7viz3cyy6lkhh7rfx3sa",
			State:     model.DealProposed,
			Price:     "200000000000000000", // 0.2 FIL
			Provider:  "f01000",
			PieceCID:  model.CID(cid1),
			PieceSize: 1024,
		}
		deal2 := &model.Deal{
			ClientID:  "f1fib3pv7jua2ockdugtz7viz3cyy6lkhh7rfx3sa",
			State:     model.DealPublished,
			Price:     "300000000000000000", // 0.3 FIL
			Provider:  "f01001",
			PieceCID:  model.CID(cid2),
			PieceSize: 2048,
		}
		
		err := db.Create(deal1).Error
		require.NoError(t, err)
		err = db.Create(deal2).Error
		require.NoError(t, err)

		mockClient := NewMockLotusClient()
		walletAddr := "f1fib3pv7jua2ockdugtz7viz3cyy6lkhh7rfx3sa"
		mockClient.SetBalance(walletAddr, "1000000000000000000") // 1 FIL
		
		requiredAmount := big.NewInt(600000000000000000) // 0.6 FIL

		result, err := validator.ValidateWalletBalance(ctx, db, mockClient, walletAddr, requiredAmount, "prep1")
		require.NoError(t, err)
		require.False(t, result.IsValid) // 1 FIL - 0.5 FIL (pending) = 0.5 FIL < 0.6 FIL required
		require.Equal(t, "1 FIL", result.CurrentBalance)
		require.Equal(t, "0.6 FIL", result.RequiredBalance)
		require.Equal(t, "0.5 FIL", result.AvailableBalance)
		require.Contains(t, result.Message, "Insufficient wallet balance")
	})

	t.Run("invalid wallet address", func(t *testing.T) {
		mockClient := NewMockLotusClient()
		requiredAmount := big.NewInt(1000000000000000000)

		result, err := validator.ValidateWalletBalance(ctx, db, mockClient, "invalid-address", requiredAmount, "prep1")
		require.NoError(t, err)
		require.False(t, result.IsValid)
		require.Equal(t, "Invalid wallet address format", result.Message)
	})

	t.Run("wallet balance query error", func(t *testing.T) {
		mockClient := NewMockLotusClient()
		walletAddr := "f1fib3pv7jua2ockdugtz7viz3cyy6lkhh7rfx3sa"
		mockClient.SetError(walletAddr, errors.New("connection failed"))
		
		requiredAmount := big.NewInt(1000000000000000000)

		result, err := validator.ValidateWalletBalance(ctx, db, mockClient, walletAddr, requiredAmount, "prep1")
		require.Error(t, err)
		require.False(t, result.IsValid)
		require.Equal(t, "Failed to retrieve wallet balance", result.Message)
	})
}

func TestGetPendingDealsAmount(t *testing.T) {
	db := setupTestDB(t)
	validator := NewBalanceValidator()
	ctx := context.Background()

	// Create valid CIDs
	cid1, _ := cid.Parse("bafybeigdyrzt5sfp7udm7hu76uh7y26nf3efuylqabf3oclgtqy55fbzdi")
	cid2, _ := cid.Parse("bafybeihdwdcefgh4dqkjv67uzcmw7ojee6xedzdetojuzjevtenxquvyku")
	cid3, _ := cid.Parse("bafybeifx7yeb55armcsxwwitkymga5xf53dxiarykms3ygqic223w5sk3m")
	cid4, _ := cid.Parse("bafybeibxm2nsadl3fnxv2sxcxmxaco2jl53wpeorjdzidjwf5aqdg7wa6u")

	// Create test deals
	deals := []*model.Deal{
		{
			ClientID:  "f1fib3pv7jua2ockdugtz7viz3cyy6lkhh7rfx3sa",
			State:     model.DealProposed,
			Price:     "100000000000000000", // 0.1 FIL
			Provider:  "f01000",
			PieceCID:  model.CID(cid1),
			PieceSize: 1024,
		},
		{
			ClientID:  "f1fib3pv7jua2ockdugtz7viz3cyy6lkhh7rfx3sa",
			State:     model.DealPublished,
			Price:     "200000000000000000", // 0.2 FIL
			Provider:  "f01001",
			PieceCID:  model.CID(cid2),
			PieceSize: 2048,
		},
		{
			ClientID:  "f1fib3pv7jua2ockdugtz7viz3cyy6lkhh7rfx3sa",
			State:     model.DealActive, // This should not be counted
			Price:     "300000000000000000", // 0.3 FIL
			Provider:  "f01002",
			PieceCID:  model.CID(cid3),
			PieceSize: 3072,
		},
		{
			ClientID:  "f1different", // Different wallet
			State:     model.DealProposed,
			Price:     "400000000000000000", // 0.4 FIL
			Provider:  "f01003",
			PieceCID:  model.CID(cid4),
			PieceSize: 4096,
		},
	}

	for _, deal := range deals {
		err := db.Create(deal).Error
		require.NoError(t, err)
	}

	amount, err := validator.getPendingDealsAmount(ctx, db, "f1fib3pv7jua2ockdugtz7viz3cyy6lkhh7rfx3sa")
	require.NoError(t, err)
	
	// Should be 0.1 + 0.2 = 0.3 FIL (only proposed and published deals for f1test123)
	expected := big.NewInt(300000000000000000)
	require.Equal(t, expected.String(), amount.String())
}

func TestWalletValidatorIntegration(t *testing.T) {
	db := setupTestDB(t)
	validator := NewBalanceValidator()
	ctx := context.Background()

	// Test the complete flow with notifications
	mockClient := NewMockLotusClient()
	walletAddr := "f1fib3pv7jua2ockdugtz7viz3cyy6lkhh7rfx3sa"
	mockClient.SetBalance(walletAddr, "500000000000000000") // 0.5 FIL
	
	requiredAmount := big.NewInt(1000000000000000000) // 1 FIL

	result, err := validator.ValidateWalletBalance(ctx, db, mockClient, walletAddr, requiredAmount, "prep1")
	require.NoError(t, err)
	require.False(t, result.IsValid)

	// Check that a notification was created
	var notifications []model.Notification
	err = db.Find(&notifications).Error
	require.NoError(t, err)
	require.Len(t, notifications, 1)
	require.Equal(t, "warning", notifications[0].Type)
	require.Equal(t, "wallet-validator", notifications[0].Source)
	require.Contains(t, notifications[0].Title, "Insufficient Wallet Balance")
}