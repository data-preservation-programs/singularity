package autodeal

import (
	"context"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

// TestTriggerIntegration tests the triggering logic without external dependencies
func TestTriggerIntegration(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		// Create a mock auto-deal service to avoid external dependencies
		mockService := &MockAutoDealer{}
		triggerService := NewTriggerService()
		triggerService.SetAutoDealService(mockService)

		// Create test data
		preparation := model.Preparation{
			Name:            "trigger-test",
			AutoCreateDeals: true,
		}
		require.NoError(t, db.Create(&preparation).Error)

		storage := model.Storage{
			Name: "test-storage",
			Type: "local",
			Path: "/tmp/test",
		}
		require.NoError(t, db.Create(&storage).Error)

		attachment := model.SourceAttachment{
			PreparationID: preparation.ID,
			StorageID:     storage.ID,
		}
		require.NoError(t, db.Create(&attachment).Error)

		job := model.Job{
			Type:         model.Pack,
			State:        model.Complete,
			AttachmentID: attachment.ID,
		}
		require.NoError(t, db.Create(&job).Error)

		// Mock expectations
		mockService.On("CheckPreparationReadiness", mock.Anything, mock.Anything, "1").Return(true, nil)
		mockService.On("CreateAutomaticDealSchedule", mock.Anything, mock.Anything, mock.Anything, "1").Return(&model.Schedule{ID: 1, PreparationID: preparation.ID}, nil)

		// Test triggering
		err := triggerService.TriggerForJobCompletion(ctx, db, nil, job.ID)
		require.NoError(t, err)

		// Verify mock was called
		mockService.AssertExpectations(t)
	})
}

// TestTriggerFlow tests the complete flow with readiness checking
func TestTriggerFlow(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		mockService := &MockAutoDealer{}
		triggerService := NewTriggerService()
		triggerService.SetAutoDealService(mockService)

		// Setup preparation
		preparation := model.Preparation{
			Name:            "flow-test",
			AutoCreateDeals: true,
		}
		require.NoError(t, db.Create(&preparation).Error)

		storage := model.Storage{
			Name: "flow-storage",
			Type: "local",
		}
		require.NoError(t, db.Create(&storage).Error)

		attachment := model.SourceAttachment{
			PreparationID: preparation.ID,
			StorageID:     storage.ID,
		}
		require.NoError(t, db.Create(&attachment).Error)

		// Create jobs - some incomplete
		jobs := []model.Job{
			{Type: model.Scan, State: model.Complete, AttachmentID: attachment.ID},
			{Type: model.Pack, State: model.Processing, AttachmentID: attachment.ID},
		}
		for i := range jobs {
			require.NoError(t, db.Create(&jobs[i]).Error)
		}

		// First trigger - should not create schedule (preparation not ready)
		mockService.On("CheckPreparationReadiness", mock.Anything, mock.Anything, "1").Return(false, nil).Once()

		err := triggerService.TriggerForJobCompletion(ctx, db, nil, jobs[0].ID)
		require.NoError(t, err)

		// Complete remaining job
		require.NoError(t, db.Model(&jobs[1]).Update("state", model.Complete).Error)

		// Second trigger - should create schedule (preparation ready)
		mockService.On("CheckPreparationReadiness", mock.Anything, mock.Anything, "1").Return(true, nil).Once()
		mockService.On("CreateAutomaticDealSchedule", mock.Anything, mock.Anything, mock.Anything, "1").Return(&model.Schedule{ID: 1, PreparationID: preparation.ID}, nil).Once()

		err = triggerService.TriggerForJobCompletion(ctx, db, nil, jobs[1].ID)
		require.NoError(t, err)

		mockService.AssertExpectations(t)
	})
}

// TestTriggerWithExistingSchedule tests that duplicate schedules are not created
func TestTriggerWithExistingSchedule(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		mockService := &MockAutoDealer{}
		triggerService := NewTriggerService()
		triggerService.SetAutoDealService(mockService)

		// Setup preparation
		preparation := model.Preparation{
			Name:            "existing-test",
			AutoCreateDeals: true,
		}
		require.NoError(t, db.Create(&preparation).Error)

		storage := model.Storage{
			Name: "existing-storage",
			Type: "local",
		}
		require.NoError(t, db.Create(&storage).Error)

		attachment := model.SourceAttachment{
			PreparationID: preparation.ID,
			StorageID:     storage.ID,
		}
		require.NoError(t, db.Create(&attachment).Error)

		job := model.Job{
			Type:         model.Pack,
			State:        model.Complete,
			AttachmentID: attachment.ID,
		}
		require.NoError(t, db.Create(&job).Error)

		// Create existing schedule
		existingSchedule := model.Schedule{
			PreparationID: preparation.ID,
			Provider:      "f01234",
		}
		require.NoError(t, db.Create(&existingSchedule).Error)

		// Mock readiness check (should still be called)
		mockService.On("CheckPreparationReadiness", mock.Anything, mock.Anything, "1").Return(true, nil)

		// Trigger - should NOT create new schedule
		err := triggerService.TriggerForJobCompletion(ctx, db, nil, job.ID)
		require.NoError(t, err)

		// Verify CreateAutomaticDealSchedule was NOT called
		mockService.AssertNotCalled(t, "CreateAutomaticDealSchedule")
		mockService.AssertExpectations(t)
	})
}

// TestTriggerDisabled tests that disabled service doesn't process triggers
func TestTriggerDisabled(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		triggerService := NewTriggerService()
		triggerService.SetEnabled(false)

		// Should return without error but not process anything
		err := triggerService.TriggerForJobCompletion(ctx, db, nil, 1)
		assert.NoError(t, err)
	})
}

// TestBatchProcessing tests the batch processing functionality
func TestBatchProcessing(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		mockService := &MockAutoDealer{}
		triggerService := NewTriggerService()
		triggerService.SetAutoDealService(mockService)

		// Mock batch processing
		mockService.On("ProcessReadyPreparations", mock.Anything, mock.Anything, mock.Anything).Return(nil)

		err := triggerService.BatchProcessReadyPreparations(ctx, db, nil)
		require.NoError(t, err)

		mockService.AssertExpectations(t)
	})
}

// TestMonitorService tests the monitor service functionality
func TestMonitorService(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		// Create mock service
		mockService := &MockAutoDealer{}
		triggerService := NewTriggerService()
		triggerService.SetAutoDealService(mockService)

		// Configure monitor with short interval
		config := MonitorConfig{
			CheckInterval:   50 * time.Millisecond,
			EnableBatchMode: true,
			ExitOnComplete:  false,
			ExitOnError:     false,
			MaxRetries:      1,
			RetryInterval:   time.Second,
		}

		monitor := NewMonitorService(db, nil, config)

		// Mock the batch processing call
		mockService.On("ProcessReadyPreparations", mock.Anything, mock.Anything, mock.Anything).Return(nil).Maybe()

		// Run monitor for a short time
		ctx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
		defer cancel()

		err := monitor.Run(ctx)
		// Should timeout (expected)
		assert.Error(t, err)

		// The mock may or may not be called depending on timing
		// Just verify service ran without panicking
	})
}

// TestErrorScenarios tests various error conditions
func TestErrorScenarios(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		triggerService := NewTriggerService()

		// Test 1: Non-existent job
		err := triggerService.TriggerForJobCompletion(ctx, db, nil, 99999)
		assert.NoError(t, err, "Should handle non-existent job gracefully")

		// Test 2: Disabled service batch processing
		triggerService.SetEnabled(false)
		err = triggerService.BatchProcessReadyPreparations(ctx, db, nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "disabled")

		// Test 3: Disabled service manual trigger
		err = triggerService.TriggerForPreparation(ctx, db, nil, "1")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "disabled")
	})
}
