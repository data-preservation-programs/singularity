package storage

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/multiformats/go-multiaddr"
	"github.com/stretchr/testify/require"
	"github.com/ybbus/jsonrpc/v3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// MockLotusClient for testing SP validation
type MockLotusClient struct {
	minerInfos map[string]*MinerInfo
	minerPower map[string]*MinerPower
	errors     map[string]error
}

func NewMockLotusClient() *MockLotusClient {
	return &MockLotusClient{
		minerInfos: make(map[string]*MinerInfo),
		minerPower: make(map[string]*MinerPower),
		errors:     make(map[string]error),
	}
}

func (m *MockLotusClient) SetMinerInfo(minerID string, info *MinerInfo) {
	m.minerInfos[minerID] = info
}

func (m *MockLotusClient) SetMinerPower(minerID string, power *MinerPower) {
	m.minerPower[minerID] = power
}

func (m *MockLotusClient) SetError(minerID string, err error) {
	m.errors[minerID] = err
}

func (m *MockLotusClient) CallFor(ctx context.Context, result interface{}, method string, args ...interface{}) error {
	if len(args) == 0 {
		return errors.New("no arguments provided")
	}

	minerAddr, ok := args[0].(address.Address)
	if !ok {
		return errors.New("invalid miner address")
	}

	minerID := minerAddr.String()

	if err, exists := m.errors[minerID]; exists {
		return err
	}

	switch method {
	case "Filecoin.StateMinerInfo":
		if info, exists := m.minerInfos[minerID]; exists {
			if infoPtr, ok := result.(*MinerInfo); ok {
				*infoPtr = *info
				return nil
			}
		}
		return errors.New("miner not found")
	case "Filecoin.StateMinerPower":
		if power, exists := m.minerPower[minerID]; exists {
			if powerPtr, ok := result.(*MinerPower); ok {
				*powerPtr = *power
				return nil
			}
		}
		return errors.New("miner power not found")
	}

	return errors.New("unsupported method")
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

	err = db.AutoMigrate(&model.Notification{})
	require.NoError(t, err)

	return db
}

func createTestMinerInfo(peerIDStr string, multiaddrs []string) *MinerInfo {
	var peerID *peer.ID
	if peerIDStr != "" {
		pid, _ := peer.Decode(peerIDStr)
		peerID = &pid
	}

	var addrs []multiaddr.Multiaddr
	for _, addr := range multiaddrs {
		ma, _ := multiaddr.NewMultiaddr(addr)
		addrs = append(addrs, ma)
	}

	return &MinerInfo{
		PeerID:     peerID,
		Multiaddrs: addrs,
		SectorSize: abi.SectorSize(34359738368), // 32 GiB
	}
}

func createTestMinerPower(qualityAdjPower int64) *MinerPower {
	return &MinerPower{
		MinerPower: Claim{
			QualityAdjPower: abi.NewStoragePower(qualityAdjPower),
		},
	}
}

func TestValidateStorageProvider(t *testing.T) {
	db := setupTestDB(t)
	validator := NewSPValidator()
	ctx := context.Background()

	t.Run("valid storage provider", func(t *testing.T) {
		mockClient := NewMockLotusClient()
		providerID := "f01000"

		// Setup mock data
		peerID := "12D3KooWExample"
		multiaddrs := []string{"/ip4/192.168.1.100/tcp/24001"}
		mockClient.SetMinerInfo(providerID, createTestMinerInfo(peerID, multiaddrs))
		mockClient.SetMinerPower(providerID, createTestMinerPower(1000000))

		result, err := validator.ValidateStorageProvider(ctx, db, mockClient, providerID, "prep1")
		require.NoError(t, err)
		require.True(t, result.IsValid)
		require.Equal(t, providerID, result.ProviderID)
		require.Equal(t, peerID, result.PeerID)
		require.Equal(t, multiaddrs, result.Multiaddrs)
		require.True(t, result.AcceptingDeals)
		require.Equal(t, "Storage provider is available and accepting deals", result.Message)
	})

	t.Run("invalid provider ID format", func(t *testing.T) {
		mockClient := NewMockLotusClient()

		result, err := validator.ValidateStorageProvider(ctx, db, mockClient, "invalid-id", "prep1")
		require.NoError(t, err)
		require.False(t, result.IsValid)
		require.Equal(t, "Invalid storage provider ID format", result.Message)
	})

	t.Run("provider not found", func(t *testing.T) {
		mockClient := NewMockLotusClient()
		providerID := "f01999"
		mockClient.SetError(providerID, errors.New("miner not found"))

		result, err := validator.ValidateStorageProvider(ctx, db, mockClient, providerID, "prep1")
		require.NoError(t, err)
		require.False(t, result.IsValid)
		require.Equal(t, "Storage provider not found on network", result.Message)
	})

	t.Run("provider with no power (not accepting deals)", func(t *testing.T) {
		mockClient := NewMockLotusClient()
		providerID := "f01001"

		// Setup mock data with zero power
		peerID := "12D3KooWExample2"
		multiaddrs := []string{"/ip4/192.168.1.101/tcp/24001"}
		mockClient.SetMinerInfo(providerID, createTestMinerInfo(peerID, multiaddrs))
		mockClient.SetMinerPower(providerID, createTestMinerPower(0)) // No power

		result, err := validator.ValidateStorageProvider(ctx, db, mockClient, providerID, "prep1")
		require.NoError(t, err)
		require.False(t, result.IsValid)
		require.False(t, result.AcceptingDeals)
		require.Contains(t, result.Message, "not accepting deals")
		require.Contains(t, result.Warnings, "no active storage power")
	})

	t.Run("provider with no peer ID", func(t *testing.T) {
		mockClient := NewMockLotusClient()
		providerID := "f01002"

		// Setup mock data without peer ID
		mockClient.SetMinerInfo(providerID, createTestMinerInfo("", []string{}))
		mockClient.SetMinerPower(providerID, createTestMinerPower(1000000))

		result, err := validator.ValidateStorageProvider(ctx, db, mockClient, providerID, "prep1")
		require.NoError(t, err)
		require.False(t, result.IsValid)
		require.False(t, result.IsOnline)
		require.Contains(t, result.Warnings, "No peer ID available")
	})
}

func TestGetDefaultStorageProviders(t *testing.T) {
	validator := NewSPValidator()

	defaults := validator.GetDefaultStorageProviders()
	require.NotEmpty(t, defaults)

	// Check that default providers have required fields
	for _, sp := range defaults {
		require.NotEmpty(t, sp.ProviderID)
		require.NotEmpty(t, sp.Name)
		require.NotEmpty(t, sp.Description)
		require.NotEmpty(t, sp.RecommendedUse)
		require.NotNil(t, sp.DefaultSettings)
	}
}

func TestGetDefaultStorageProvider(t *testing.T) {
	db := setupTestDB(t)
	validator := NewSPValidator()
	ctx := context.Background()

	defaultSP, err := validator.GetDefaultStorageProvider(ctx, db, "test-criteria")
	require.NoError(t, err)
	require.NotNil(t, defaultSP)
	require.NotEmpty(t, defaultSP.ProviderID)
	require.NotEmpty(t, defaultSP.Name)

	// Check that a notification was created
	var notifications []model.Notification
	err = db.Find(&notifications).Error
	require.NoError(t, err)
	require.Len(t, notifications, 1)
	require.Equal(t, "info", notifications[0].Type)
	require.Equal(t, "sp-validator", notifications[0].Source)
	require.Contains(t, notifications[0].Title, "Default Storage Provider Selected")
}

func TestValidateAndGetDefault(t *testing.T) {
	db := setupTestDB(t)
	validator := NewSPValidator()
	ctx := context.Background()

	t.Run("valid provider specified", func(t *testing.T) {
		mockClient := NewMockLotusClient()
		providerID := "f01000"

		// Setup valid provider
		peerID := "12D3KooWExample"
		multiaddrs := []string{"/ip4/192.168.1.100/tcp/24001"}
		mockClient.SetMinerInfo(providerID, createTestMinerInfo(peerID, multiaddrs))
		mockClient.SetMinerPower(providerID, createTestMinerPower(1000000))

		result, defaultSP, err := validator.ValidateAndGetDefault(ctx, db, mockClient, providerID, "prep1")
		require.NoError(t, err)
		require.NotNil(t, result)
		require.Nil(t, defaultSP) // Should not need default
		require.True(t, result.IsValid)
		require.Equal(t, providerID, result.ProviderID)
	})

	t.Run("invalid provider falls back to default", func(t *testing.T) {
		mockClient := NewMockLotusClient()
		providerID := "f01999" // Invalid provider
		mockClient.SetError(providerID, errors.New("miner not found"))

		// Setup default provider as valid
		defaultProviderID := "f01000"
		peerID := "12D3KooWExample"
		multiaddrs := []string{"/ip4/192.168.1.100/tcp/24001"}
		mockClient.SetMinerInfo(defaultProviderID, createTestMinerInfo(peerID, multiaddrs))
		mockClient.SetMinerPower(defaultProviderID, createTestMinerPower(1000000))

		result, defaultSP, err := validator.ValidateAndGetDefault(ctx, db, mockClient, providerID, "prep1")
		require.NoError(t, err)
		require.NotNil(t, result)
		require.NotNil(t, defaultSP)
		require.Equal(t, defaultProviderID, result.ProviderID)
		require.Equal(t, defaultProviderID, defaultSP.ProviderID)
	})

	t.Run("no provider specified uses default", func(t *testing.T) {
		mockClient := NewMockLotusClient()

		// Setup default provider as valid
		defaultProviderID := "f01000"
		peerID := "12D3KooWExample"
		multiaddrs := []string{"/ip4/192.168.1.100/tcp/24001"}
		mockClient.SetMinerInfo(defaultProviderID, createTestMinerInfo(peerID, multiaddrs))
		mockClient.SetMinerPower(defaultProviderID, createTestMinerPower(1000000))

		result, defaultSP, err := validator.ValidateAndGetDefault(ctx, db, mockClient, "", "prep1")
		require.NoError(t, err)
		require.NotNil(t, result)
		require.NotNil(t, defaultSP)
		require.Equal(t, defaultProviderID, result.ProviderID)
		require.Equal(t, defaultProviderID, defaultSP.ProviderID)
	})
}

func TestCheckPeerConnectivity(t *testing.T) {
	validator := NewSPValidator()
	ctx := context.Background()

	t.Run("no multiaddrs", func(t *testing.T) {
		connected := validator.checkPeerConnectivity(ctx, []string{})
		require.False(t, connected)
	})

	t.Run("invalid multiaddrs", func(t *testing.T) {
		multiaddrs := []string{"invalid", "/invalid/format"}
		connected := validator.checkPeerConnectivity(ctx, multiaddrs)
		require.False(t, connected)
	})

	t.Run("valid format but unreachable", func(t *testing.T) {
		multiaddrs := []string{"/ip4/192.168.255.255/tcp/99999"}
		connected := validator.checkPeerConnectivity(ctx, multiaddrs)
		require.False(t, connected)
	})
}

func TestTestConnection(t *testing.T) {
	validator := NewSPValidator()
	ctx := context.Background()

	t.Run("invalid multiaddr format", func(t *testing.T) {
		result := validator.testConnection(ctx, "invalid")
		require.False(t, result)
	})

	t.Run("incomplete multiaddr", func(t *testing.T) {
		result := validator.testConnection(ctx, "/ip4/192.168.1.1")
		require.False(t, result)
	})

	t.Run("unreachable address", func(t *testing.T) {
		result := validator.testConnection(ctx, "/ip4/192.168.255.255/tcp/99999")
		require.False(t, result)
	})
}

func TestSPValidatorIntegration(t *testing.T) {
	db := setupTestDB(t)
	validator := NewSPValidator()
	ctx := context.Background()

	// Test the complete flow with notifications
	mockClient := NewMockLotusClient()
	providerID := "f01999"
	mockClient.SetError(providerID, errors.New("miner not found"))

	result, err := validator.ValidateStorageProvider(ctx, db, mockClient, providerID, "prep1")
	require.NoError(t, err)
	require.False(t, result.IsValid)

	// Check that a notification was created
	var notifications []model.Notification
	err = db.Find(&notifications).Error
	require.NoError(t, err)
	require.Len(t, notifications, 1)
	require.Equal(t, "error", notifications[0].Type)
	require.Equal(t, "sp-validator", notifications[0].Source)
	require.Contains(t, notifications[0].Title, "Storage Provider Not Found")
}

func TestGetDefaultStorageProvidersStructure(t *testing.T) {
	defaults := getDefaultStorageProviders()
	require.NotEmpty(t, defaults)

	for i, sp := range defaults {
		t.Run(fmt.Sprintf("default_sp_%d", i), func(t *testing.T) {
			require.NotEmpty(t, sp.ProviderID)
			require.NotEmpty(t, sp.Name)
			require.NotEmpty(t, sp.Description)
			require.NotEmpty(t, sp.RecommendedUse)
			require.NotNil(t, sp.DefaultSettings)

			// Check required default settings
			require.Contains(t, sp.DefaultSettings, "price_per_gb_epoch")
			require.Contains(t, sp.DefaultSettings, "verified")
			require.Contains(t, sp.DefaultSettings, "duration")
			require.Contains(t, sp.DefaultSettings, "start_delay")
		})
	}
}
