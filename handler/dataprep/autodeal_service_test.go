package dataprep

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/handler/storage"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/ybbus/jsonrpc/v3"
	"gorm.io/gorm"
)

// MockLotusClient for testing
type MockLotusClient struct {
	responses map[string]interface{}
	errors    map[string]error
}

func NewMockLotusClient() *MockLotusClient {
	return &MockLotusClient{
		responses: make(map[string]interface{}),
		errors:    make(map[string]error),
	}
}

func (m *MockLotusClient) Call(ctx context.Context, method string, params ...interface{}) (*jsonrpc.RPCResponse, error) {
	return nil, nil
}

func (m *MockLotusClient) CallFor(ctx context.Context, out interface{}, method string, params ...interface{}) error {
	if err, exists := m.errors[method]; exists {
		return err
	}
	if response, exists := m.responses[method]; exists {
		switch v := out.(type) {
		case *string:
			if str, ok := response.(string); ok {
				*v = str
			}
		}
	}
	return nil
}

func (m *MockLotusClient) CallBatch(ctx context.Context, requests jsonrpc.RPCRequests) (jsonrpc.RPCResponses, error) {
	return nil, nil
}

func (m *MockLotusClient) CallBatchRaw(ctx context.Context, requests jsonrpc.RPCRequests) (jsonrpc.RPCResponses, error) {
	return nil, nil
}

func (m *MockLotusClient) CallRaw(ctx context.Context, request *jsonrpc.RPCRequest) (*jsonrpc.RPCResponse, error) {
	return nil, nil
}

func (m *MockLotusClient) SetResponse(method string, response interface{}) {
	m.responses[method] = response
}

func (m *MockLotusClient) SetError(method string, err error) {
	m.errors[method] = err
}

func TestAutoDealService_CreateAutomaticDealSchedule(t *testing.T) {
	tmp1 := t.TempDir()
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		service := NewAutoDealService()
		mockClient := NewMockLotusClient()

		// Set up mock responses
		mockClient.SetResponse("Filecoin.StateLookupID", "f01000")

		// Create source storage
		sourceStorage, err := storage.Default.CreateStorageHandler(ctx, db, "local", storage.CreateRequest{
			Name: "test-source",
			Path: tmp1,
		})
		require.NoError(t, err)

		// Create wallet
		testWallet := &model.Wallet{
			ID:      "f01234",
			Address: "f1abjxfbp274xpdqcpuaykwkfb43omjotacm2p3za",
		}
		err = db.Create(testWallet).Error
		require.NoError(t, err)

		// Create preparation with auto-deal enabled but validation disabled to avoid complex mocking
		preparation := &model.Preparation{
			Name:                "test-auto-prep",
			AutoCreateDeals:     true,
			DealPricePerGB:      0.1,
			DealPricePerGBEpoch: 0.0000001,
			DealPricePerDeal:    0.01,
			DealDuration:        time.Hour * 24 * 535,
			DealStartDelay:      time.Hour * 72,
			DealVerified:        true,
			DealKeepUnsealed:    true,
			DealAnnounceToIPNI:  true,
			DealProvider:        "f01000",
			DealURLTemplate:     "https://example.com/deals/{id}",
			WalletValidation:    false, // Disable to avoid complex validation
			SPValidation:        false, // Disable to avoid complex validation
			SourceStorages:      []model.Storage{*sourceStorage},
			Wallets:             []model.Wallet{*testWallet},
		}
		err = db.Create(preparation).Error
		require.NoError(t, err)

		// Test auto-deal schedule creation
		schedule, err := service.CreateAutomaticDealSchedule(ctx, db, mockClient, fmt.Sprintf("%d", preparation.ID))
		require.NoError(t, err)
		require.NotNil(t, schedule)

		// Verify schedule properties
		assert.Equal(t, preparation.ID, schedule.PreparationID)
		assert.Equal(t, "f01000", schedule.Provider)
		assert.Equal(t, 0.1, schedule.PricePerGB)
		assert.Equal(t, 0.0000001, schedule.PricePerGBEpoch)
		assert.Equal(t, 0.01, schedule.PricePerDeal)
		assert.True(t, schedule.Verified)
		assert.True(t, schedule.KeepUnsealed)
		assert.True(t, schedule.AnnounceToIPNI)
		assert.Contains(t, schedule.Notes, "Automatically created")

		// Verify notification was created
		var notifications []model.Notification
		err = db.Where("source = ?", "auto-deal-service").Find(&notifications).Error
		require.NoError(t, err)
		assert.Greater(t, len(notifications), 0)
	})
}

func TestAutoDealService_CreateAutomaticDealSchedule_NotEnabled(t *testing.T) {
	tmp1 := t.TempDir()
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		service := NewAutoDealService()
		mockClient := NewMockLotusClient()

		// Create source storage
		sourceStorage, err := storage.Default.CreateStorageHandler(ctx, db, "local", storage.CreateRequest{
			Name: "test-source",
			Path: tmp1,
		})
		require.NoError(t, err)

		// Create preparation without auto-deal enabled
		preparation := &model.Preparation{
			Name:            "test-prep-no-auto",
			AutoCreateDeals: false, // Disabled
			SourceStorages:  []model.Storage{*sourceStorage},
		}
		err = db.Create(preparation).Error
		require.NoError(t, err)

		// Test auto-deal schedule creation
		schedule, err := service.CreateAutomaticDealSchedule(ctx, db, mockClient, fmt.Sprintf("%d", preparation.ID))
		require.NoError(t, err)
		assert.Nil(t, schedule) // Should return nil when auto-deal is disabled

		// Verify appropriate notification was created
		var notifications []model.Notification
		err = db.Where("source = ? AND title = ?", "auto-deal-service", "Auto-Deal Not Enabled").Find(&notifications).Error
		require.NoError(t, err)
		assert.Greater(t, len(notifications), 0)
	})
}

func TestAutoDealService_CheckPreparationReadiness(t *testing.T) {
	tmp1 := t.TempDir()
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		service := NewAutoDealService()

		// Create source storage
		sourceStorage, err := storage.Default.CreateStorageHandler(ctx, db, "local", storage.CreateRequest{
			Name: "test-source",
			Path: tmp1,
		})
		require.NoError(t, err)

		// Create preparation
		preparation := &model.Preparation{
			Name:           "test-prep",
			SourceStorages: []model.Storage{*sourceStorage},
		}
		err = db.Create(preparation).Error
		require.NoError(t, err)

		// Get the source attachment that was created automatically
		var attachment model.SourceAttachment
		err = db.Where("preparation_id = ? AND storage_id = ?", preparation.ID, sourceStorage.ID).First(&attachment).Error
		require.NoError(t, err)

		// Test with no jobs (should be ready)
		isReady, err := service.CheckPreparationReadiness(ctx, db, fmt.Sprintf("%d", preparation.ID))
		require.NoError(t, err)
		assert.True(t, isReady)

		// Add incomplete job
		job := &model.Job{
			AttachmentID: attachment.ID,
			State:        model.Processing,
		}
		err = db.Create(job).Error
		require.NoError(t, err)

		// Test with incomplete job (should not be ready)
		isReady, err = service.CheckPreparationReadiness(ctx, db, fmt.Sprintf("%d", preparation.ID))
		require.NoError(t, err)
		assert.False(t, isReady)

		// Complete the job
		err = db.Model(job).Update("state", model.Complete).Error
		require.NoError(t, err)

		// Test with completed job (should be ready)
		isReady, err = service.CheckPreparationReadiness(ctx, db, fmt.Sprintf("%d", preparation.ID))
		require.NoError(t, err)
		assert.True(t, isReady)
	})
}

func TestAutoDealService_ProcessReadyPreparations(t *testing.T) {
	tmp1 := t.TempDir()
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		service := NewAutoDealService()
		mockClient := NewMockLotusClient()

		// Set up mock responses
		mockClient.SetResponse("Filecoin.StateLookupID", "f01000")
		mockClient.SetResponse("Filecoin.WalletBalance", "1000000000000000000") // 1 FIL

		// Create source storage
		sourceStorage, err := storage.Default.CreateStorageHandler(ctx, db, "local", storage.CreateRequest{
			Name: "test-source",
			Path: tmp1,
		})
		require.NoError(t, err)

		// Create wallet
		testWallet := &model.Wallet{
			ID:      "f01234",
			Address: "f1abjxfbp274xpdqcpuaykwkfb43omjotacm2p3za",
		}
		err = db.Create(testWallet).Error
		require.NoError(t, err)

		// Create preparation with auto-deal enabled (ready for deals)
		preparation1 := &model.Preparation{
			Name:            "test-auto-prep-1",
			AutoCreateDeals: true,
			DealProvider:    "f01000",
			SourceStorages:  []model.Storage{*sourceStorage},
			Wallets:         []model.Wallet{*testWallet},
		}
		err = db.Create(preparation1).Error
		require.NoError(t, err)

		// Create preparation with auto-deal disabled
		preparation2 := &model.Preparation{
			Name:            "test-prep-no-auto",
			AutoCreateDeals: false,
			SourceStorages:  []model.Storage{*sourceStorage},
		}
		err = db.Create(preparation2).Error
		require.NoError(t, err)

		// Process ready preparations
		err = service.ProcessReadyPreparations(ctx, db, mockClient)
		require.NoError(t, err)

		// Verify that a schedule was created for preparation1
		var schedules []model.Schedule
		err = db.Where("preparation_id = ?", preparation1.ID).Find(&schedules).Error
		require.NoError(t, err)
		assert.Len(t, schedules, 1)

		// Verify that no schedule was created for preparation2
		err = db.Where("preparation_id = ?", preparation2.ID).Find(&schedules).Error
		require.NoError(t, err)
		assert.Len(t, schedules, 0)

		// Verify notifications were created
		var notifications []model.Notification
		err = db.Where("source = ?", "auto-deal-service").Find(&notifications).Error
		require.NoError(t, err)
		assert.Greater(t, len(notifications), 0)
	})
}

func TestAutoDealService_buildDealScheduleRequest(t *testing.T) {
	service := NewAutoDealService()

	preparation := &model.Preparation{
		ID:                  1,
		DealProvider:        "f01000",
		DealPricePerGB:      0.1,
		DealPricePerGBEpoch: 0.0000001,
		DealPricePerDeal:    0.01,
		DealDuration:        time.Hour * 24 * 535,
		DealStartDelay:      time.Hour * 72,
		DealVerified:        true,
		DealKeepUnsealed:    true,
		DealAnnounceToIPNI:  true,
		DealURLTemplate:     "https://example.com/deals/{id}",
		DealHTTPHeaders: model.ConfigMap{
			"Authorization": "Bearer token",
			"Content-Type":  "application/json",
		},
	}

	request := service.buildDealScheduleRequest(preparation)

	assert.Equal(t, "1", request.Preparation)
	assert.Equal(t, "f01000", request.Provider)
	assert.Equal(t, 0.1, request.PricePerGB)
	assert.Equal(t, 0.0000001, request.PricePerGBEpoch)
	assert.Equal(t, 0.01, request.PricePerDeal)
	assert.True(t, request.Verified)
	assert.True(t, request.KeepUnsealed)
	assert.True(t, request.IPNI)
	assert.Equal(t, "https://example.com/deals/{id}", request.URLTemplate)
	assert.Contains(t, request.Notes, "Automatically created")
	assert.Contains(t, request.HTTPHeaders, "Authorization=Bearer token")
	assert.Contains(t, request.HTTPHeaders, "Content-Type=application/json")
	assert.Equal(t, (time.Hour * 72).String(), request.StartDelay)
	assert.Equal(t, (time.Hour * 24 * 535).String(), request.Duration)
}

func TestAutoDealService_ValidationErrors(t *testing.T) {
	tmp1 := t.TempDir()
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		service := NewAutoDealService()
		mockClient := NewMockLotusClient()

		// Set up mock responses
		mockClient.SetResponse("Filecoin.StateLookupID", "f01000")

		// Create source storage
		sourceStorage, err := storage.Default.CreateStorageHandler(ctx, db, "local", storage.CreateRequest{
			Name: "test-source",
			Path: tmp1,
		})
		require.NoError(t, err)

		// Create preparation with auto-deal enabled but no wallets (validation will fail)
		preparation := &model.Preparation{
			Name:             "test-auto-prep-fail",
			AutoCreateDeals:  true,
			DealProvider:     "f01000",
			WalletValidation: true,  // Enable validation
			SPValidation:     false, // Disable to avoid complex mocking
			SourceStorages:   []model.Storage{*sourceStorage},
			// No wallets - this should cause validation to fail
		}
		err = db.Create(preparation).Error
		require.NoError(t, err)

		// Test auto-deal schedule creation should fail
		schedule, err := service.CreateAutomaticDealSchedule(ctx, db, mockClient, fmt.Sprintf("%d", preparation.ID))
		assert.Error(t, err)
		assert.Nil(t, schedule)

		// Verify error notification was created
		var notifications []model.Notification
		err = db.Where("source = ? AND type = ?", "auto-deal-service", "error").Find(&notifications).Error
		require.NoError(t, err)
		assert.Greater(t, len(notifications), 0)
	})
}
