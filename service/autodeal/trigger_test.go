package autodeal

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/ybbus/jsonrpc/v3"
	"gorm.io/gorm"
)

type MockAutoDealer struct {
	mock.Mock
}

func (m *MockAutoDealer) CheckPreparationReadiness(ctx context.Context, db *gorm.DB, preparationID string) (bool, error) {
	args := m.Called(ctx, db, preparationID)
	return args.Bool(0), args.Error(1)
}

func (m *MockAutoDealer) CreateAutomaticDealSchedule(ctx context.Context, db *gorm.DB, lotusClient jsonrpc.RPCClient, preparationID string) (*model.Schedule, error) {
	args := m.Called(ctx, db, lotusClient, preparationID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Schedule), args.Error(1)
}

func (m *MockAutoDealer) ProcessReadyPreparations(ctx context.Context, db *gorm.DB, lotusClient jsonrpc.RPCClient) error {
	args := m.Called(ctx, db, lotusClient)
	return args.Error(0)
}

var _ AutoDealServiceInterface = (*MockAutoDealer)(nil)

func TestTriggerService_SetEnabled(t *testing.T) {
	service := NewTriggerService()

	// Test initial state
	assert.True(t, service.IsEnabled())

	// Test disable
	service.SetEnabled(false)
	assert.False(t, service.IsEnabled())

	// Test enable
	service.SetEnabled(true)
	assert.True(t, service.IsEnabled())
}

func TestTriggerService_TriggerForJobCompletion_Disabled(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		service := NewTriggerService()
		service.SetEnabled(false)

		err := service.TriggerForJobCompletion(ctx, db, nil, 1)

		assert.NoError(t, err)
	})
}

func TestTriggerService_TriggerForJobCompletion_AutoDealDisabled(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		service := NewTriggerService()

		// Create test data
		preparation := model.Preparation{
			Name: "test-prep",
			DealConfig: model.DealConfig{
				DealConfig: model.DealConfig{
					AutoCreateDeals: false,
				},
			},
		}
		db.Create(&preparation)

		storage := model.Storage{
			Name: "test-storage",
			Type: "local",
		}
		db.Create(&storage)

		attachment := model.SourceAttachment{
			PreparationID: preparation.ID,
			StorageID:     storage.ID,
		}
		db.Create(&attachment)

		job := model.Job{
			Type:         model.Pack,
			State:        model.Complete,
			AttachmentID: attachment.ID,
		}
		db.Create(&job)

		err := service.TriggerForJobCompletion(ctx, db, nil, job.ID)

		assert.NoError(t, err)
	})
}

func TestTriggerService_TriggerForJobCompletion_NotReady(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		service := NewTriggerService()

		// Mock the auto-deal service
		mockAutoDealer := &MockAutoDealer{}
		service.SetAutoDealService(mockAutoDealer)

		// Create test data
		preparation := model.Preparation{
			Name: "test-prep",
			DealConfig: model.DealConfig{
				AutoCreateDeals: true,
			},
		}
		db.Create(&preparation)

		storage := model.Storage{
			Name: "test-storage",
			Type: "local",
		}
		db.Create(&storage)

		attachment := model.SourceAttachment{
			PreparationID: preparation.ID,
			StorageID:     storage.ID,
		}
		db.Create(&attachment)

		job := model.Job{
			Type:         model.Pack,
			State:        model.Complete,
			AttachmentID: attachment.ID,
		}
		db.Create(&job)

		// Mock that preparation is not ready
		mockAutoDealer.On("CheckPreparationReadiness", mock.Anything, mock.Anything, "1").Return(false, nil)

		err := service.TriggerForJobCompletion(ctx, db, nil, job.ID)

		assert.NoError(t, err)
		mockAutoDealer.AssertExpectations(t)
	})
}

func TestTriggerService_TriggerForJobCompletion_Success(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		service := NewTriggerService()

		// Mock the auto-deal service
		mockAutoDealer := &MockAutoDealer{}
		service.SetAutoDealService(mockAutoDealer)

		// Create test data
		preparation := model.Preparation{
			Name: "test-prep",
			DealConfig: model.DealConfig{
				AutoCreateDeals: true,
			},
		}
		db.Create(&preparation)

		storage := model.Storage{
			Name: "test-storage",
			Type: "local",
		}
		db.Create(&storage)

		attachment := model.SourceAttachment{
			PreparationID: preparation.ID,
			StorageID:     storage.ID,
		}
		db.Create(&attachment)

		job := model.Job{
			Type:         model.Pack,
			State:        model.Complete,
			AttachmentID: attachment.ID,
		}
		db.Create(&job)

		expectedSchedule := &model.Schedule{
			ID:            1,
			PreparationID: preparation.ID,
		}

		// Mock successful flow
		mockAutoDealer.On("CheckPreparationReadiness", mock.Anything, mock.Anything, "1").Return(true, nil)
		mockAutoDealer.On("CreateAutomaticDealSchedule", mock.Anything, mock.Anything, mock.Anything, "1").Return(expectedSchedule, nil)

		err := service.TriggerForJobCompletion(ctx, db, nil, job.ID)

		assert.NoError(t, err)
		mockAutoDealer.AssertExpectations(t)
	})
}

func TestTriggerService_TriggerForJobCompletion_ExistingSchedule(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		service := NewTriggerService()

		// Mock the auto-deal service
		mockAutoDealer := &MockAutoDealer{}
		service.SetAutoDealService(mockAutoDealer)

		// Create test data
		preparation := model.Preparation{
			Name: "test-prep",
			DealConfig: model.DealConfig{
				AutoCreateDeals: true,
			},
		}
		db.Create(&preparation)

		storage := model.Storage{
			Name: "test-storage",
			Type: "local",
		}
		db.Create(&storage)

		attachment := model.SourceAttachment{
			PreparationID: preparation.ID,
			StorageID:     storage.ID,
		}
		db.Create(&attachment)

		job := model.Job{
			Type:         model.Pack,
			State:        model.Complete,
			AttachmentID: attachment.ID,
		}
		db.Create(&job)

		// Create existing schedule
		existingSchedule := model.Schedule{
			PreparationID: preparation.ID,
			Provider:      "f01234",
		}
		db.Create(&existingSchedule)

		// Mock that preparation is ready but should skip due to existing schedule
		mockAutoDealer.On("CheckPreparationReadiness", mock.Anything, mock.Anything, "1").Return(true, nil)

		err := service.TriggerForJobCompletion(ctx, db, nil, job.ID)

		assert.NoError(t, err)
		mockAutoDealer.AssertExpectations(t)
		// CreateAutomaticDealSchedule should NOT be called due to existing schedule
		mockAutoDealer.AssertNotCalled(t, "CreateAutomaticDealSchedule")
	})
}

func TestTriggerService_TriggerForPreparation_Disabled(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		service := NewTriggerService()
		service.SetEnabled(false)

		err := service.TriggerForPreparation(ctx, nil, nil, "1")

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "disabled")
	})
}

func TestTriggerService_TriggerForPreparation_Success(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		service := NewTriggerService()

		// Mock the auto-deal service
		mockAutoDealer := &MockAutoDealer{}
		service.SetAutoDealService(mockAutoDealer)

		expectedSchedule := &model.Schedule{
			ID:            1,
			PreparationID: 1,
		}

		mockAutoDealer.On("CreateAutomaticDealSchedule", mock.Anything, mock.Anything, mock.Anything, "1").Return(expectedSchedule, nil)

		err := service.TriggerForPreparation(ctx, nil, nil, "1")

		assert.NoError(t, err)
		mockAutoDealer.AssertExpectations(t)
	})
}

func TestTriggerService_BatchProcessReadyPreparations_Disabled(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		service := NewTriggerService()
		service.SetEnabled(false)

		err := service.BatchProcessReadyPreparations(ctx, nil, nil)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "disabled")
	})
}

func TestTriggerService_BatchProcessReadyPreparations_Success(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		service := NewTriggerService()

		// Mock the auto-deal service
		mockAutoDealer := &MockAutoDealer{}
		service.SetAutoDealService(mockAutoDealer)

		mockAutoDealer.On("ProcessReadyPreparations", mock.Anything, mock.Anything, mock.Anything).Return(nil)

		err := service.BatchProcessReadyPreparations(ctx, nil, nil)

		assert.NoError(t, err)
		mockAutoDealer.AssertExpectations(t)
	})
}
