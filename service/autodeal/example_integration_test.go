package autodeal

import (
	"context"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

// TestFullIntegrationWorkflow demonstrates the complete auto-deal workflow
// This test shows how the system automatically creates deal schedules when
// preparation jobs complete with auto-deal enabled
func TestFullIntegrationWorkflow(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		// Step 1: Create a wallet
		wallet := model.Wallet{
			ID:      "test1",
			Address: "f3test123456789abcdef",
		}
		require.NoError(t, db.Create(&wallet).Error)

		// Step 2: Create a preparation with auto-deal enabled
		preparation := model.Preparation{
			Name:             "auto-deal-test",
			AutoCreateDeals:  true,
			DealProvider:     "f01234", // Test provider
			DealVerified:     true,
			DealPricePerGB:   0.0000001,
			DealDuration:     time.Hour * 24 * 365, // 1 year
			DealStartDelay:   time.Hour * 72,       // 3 days
			WalletValidation: false,                // Skip validation for test
			SPValidation:     false,                // Skip validation for test
			PieceSize:        34359738368,          // 32GiB
		}
		require.NoError(t, db.Create(&preparation).Error)

		// Attach wallet to preparation using GORM many-to-many
		require.NoError(t, db.Model(&preparation).Association("Wallets").Append(&wallet))

		// Debug: Verify wallet association
		var prepWithWallets model.Preparation
		err := prepWithWallets.FindByIDOrName(db, "1", "Wallets")
		require.NoError(t, err)
		require.Len(t, prepWithWallets.Wallets, 1, "Preparation should have one wallet attached")
		t.Logf("Wallet attached: %s", prepWithWallets.Wallets[0].Address)

		// Step 3: Create storage and attachment
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

		// Step 4: Create all jobs as complete (simplify for debugging)
		jobs := []model.Job{
			{
				Type:         model.Pack,
				State:        model.Complete,
				AttachmentID: attachment.ID,
			},
		}
		for i := range jobs {
			require.NoError(t, db.Create(&jobs[i]).Error)
		}

		// Step 5: Verify preparation is ready
		triggerService := NewTriggerService()
		isReady, err := triggerService.autoDealService.CheckPreparationReadiness(ctx, db, "1")
		require.NoError(t, err)
		assert.True(t, isReady, "Preparation should be ready with all jobs complete")

		// Step 6: Trigger auto-deal creation
		lotusClient := util.NewLotusClient("", "")
		err = triggerService.TriggerForJobCompletion(ctx, db, lotusClient, jobs[0].ID)
		require.NoError(t, err, "Auto-deal creation should succeed")

		// Step 7: Verify deal schedule was created
		var schedules []model.Schedule
		err = db.Where("preparation_id = ?", preparation.ID).Find(&schedules).Error
		require.NoError(t, err)
		require.Len(t, schedules, 1, "One deal schedule should be created")

		schedule := schedules[0]
		assert.Equal(t, preparation.ID, schedule.PreparationID)
		assert.Equal(t, "f01234", schedule.Provider)
		assert.True(t, schedule.Verified)
		assert.Contains(t, schedule.Notes, "Automatically created")
	})
}

// TestMonitorServiceIntegration tests the background monitoring service
func TestMonitorServiceIntegration(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		// Create a wallet
		wallet := model.Wallet{
			ID:      "monitor1",
			Address: "f3monitor123456789abcdef",
		}
		require.NoError(t, db.Create(&wallet).Error)

		// Create a preparation ready for auto-deal
		preparation := model.Preparation{
			Name:             "monitor-test",
			AutoCreateDeals:  true,
			DealProvider:     "f05678",
			DealVerified:     false,
			DealPricePerGB:   0.0000002,
			WalletValidation: false,
			SPValidation:     false,
		}
		require.NoError(t, db.Create(&preparation).Error)

		// Attach wallet to preparation using GORM many-to-many
		require.NoError(t, db.Model(&preparation).Association("Wallets").Append(&wallet))

		storage := model.Storage{
			Name: "monitor-storage",
			Type: "local",
			Path: "/tmp/monitor",
		}
		require.NoError(t, db.Create(&storage).Error)

		attachment := model.SourceAttachment{
			PreparationID: preparation.ID,
			StorageID:     storage.ID,
		}
		require.NoError(t, db.Create(&attachment).Error)

		// All jobs complete - preparation is ready
		job := model.Job{
			Type:         model.Pack,
			State:        model.Complete,
			AttachmentID: attachment.ID,
		}
		require.NoError(t, db.Create(&job).Error)

		// Configure monitor service for quick processing
		config := MonitorConfig{
			CheckInterval:   100 * time.Millisecond,
			EnableBatchMode: true,
			ExitOnComplete:  false,
			ExitOnError:     true,
			MaxRetries:      1,
			RetryInterval:   time.Second,
		}

		lotusClient := util.NewLotusClient("", "")
		monitor := NewMonitorService(db, lotusClient, config)

		// Run monitor service briefly
		ctx, cancel := context.WithTimeout(ctx, time.Second)
		defer cancel()

		// This will process the ready preparation
		err := monitor.Run(ctx)
		// Should timeout, which is expected
		assert.Error(t, err, "Monitor should exit due to context timeout")

		// Verify schedule was created by the monitor
		var scheduleCount int64
		db.Model(&model.Schedule{}).Where("preparation_id = ?", preparation.ID).Count(&scheduleCount)
		assert.Equal(t, int64(1), scheduleCount, "Monitor should have created one schedule")
	})
}

// TestRealWorldScenario demonstrates a realistic scenario with multiple preparations
func TestRealWorldScenario(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		triggerService := NewTriggerService()
		lotusClient := util.NewLotusClient("", "")

		// Create wallets for testing
		wallets := []model.Wallet{
			{ID: "dataset1", Address: "f3dataset1123456789abcdef"},
			{ID: "dataset2", Address: "f3dataset2123456789abcdef"},
			{ID: "dataset3", Address: "f3dataset3123456789abcdef"},
		}
		for i := range wallets {
			require.NoError(t, db.Create(&wallets[i]).Error)
		}

		// Scenario: Multiple preparations, some with auto-deal enabled, some without
		preparations := []model.Preparation{
			{
				Name:             "dataset-1",
				AutoCreateDeals:  true,
				DealProvider:     "f01111",
				DealVerified:     true,
				WalletValidation: false,
				SPValidation:     false,
			},
			{
				Name:             "dataset-2",
				AutoCreateDeals:  false, // Auto-deal disabled
				DealProvider:     "f02222",
				WalletValidation: false,
				SPValidation:     false,
			},
			{
				Name:             "dataset-3",
				AutoCreateDeals:  true,
				DealProvider:     "f03333",
				DealVerified:     false,
				WalletValidation: false,
				SPValidation:     false,
			},
		}

		for i := range preparations {
			require.NoError(t, db.Create(&preparations[i]).Error)

			// Attach wallet to each preparation using GORM many-to-many
			require.NoError(t, db.Model(&preparations[i]).Association("Wallets").Append(&wallets[i]))
		}

		storage := model.Storage{
			Name: "shared-storage",
			Type: "local",
			Path: "/tmp/datasets",
		}
		require.NoError(t, db.Create(&storage).Error)

		// Create attachments and complete jobs for each preparation
		for i, prep := range preparations {
			attachment := model.SourceAttachment{
				PreparationID: prep.ID,
				StorageID:     storage.ID,
			}
			require.NoError(t, db.Create(&attachment).Error)

			// Create completed jobs
			job := model.Job{
				Type:         model.Pack,
				State:        model.Complete,
				AttachmentID: attachment.ID,
			}
			require.NoError(t, db.Create(&job).Error)

			// Trigger auto-deal for each completion
			err := triggerService.TriggerForJobCompletion(ctx, db, lotusClient, job.ID)
			require.NoError(t, err, "Failed to trigger for preparation %d", i+1)
		}

		// Verify results: only preparations with auto-deal enabled should have schedules
		var allSchedules []model.Schedule
		err := db.Find(&allSchedules).Error
		require.NoError(t, err)

		assert.Len(t, allSchedules, 2, "Should have 2 schedules (dataset-1 and dataset-3)")

		// Verify specific schedules
		var dataset1Schedules []model.Schedule
		err = db.Where("preparation_id = ?", preparations[0].ID).Find(&dataset1Schedules).Error
		require.NoError(t, err)
		assert.Len(t, dataset1Schedules, 1, "Dataset-1 should have one schedule")

		var dataset2Schedules []model.Schedule
		err = db.Where("preparation_id = ?", preparations[1].ID).Find(&dataset2Schedules).Error
		require.NoError(t, err)
		assert.Len(t, dataset2Schedules, 0, "Dataset-2 should have no schedules (auto-deal disabled)")

		var dataset3Schedules []model.Schedule
		err = db.Where("preparation_id = ?", preparations[2].ID).Find(&dataset3Schedules).Error
		require.NoError(t, err)
		assert.Len(t, dataset3Schedules, 1, "Dataset-3 should have one schedule")
	})
}

// TestErrorHandling tests error scenarios and recovery
func TestErrorHandling(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		triggerService := NewTriggerService()
		lotusClient := util.NewLotusClient("", "")

		// Test 1: Trigger for non-existent job
		err := triggerService.TriggerForJobCompletion(ctx, db, lotusClient, 99999)
		assert.NoError(t, err, "Should gracefully handle non-existent job")

		// Test 2: Disabled service
		triggerService.SetEnabled(false)
		err = triggerService.TriggerForJobCompletion(ctx, db, lotusClient, 1)
		assert.NoError(t, err, "Disabled service should not error but should not process")

		// Test 3: Batch processing with disabled service
		err = triggerService.BatchProcessReadyPreparations(ctx, db, lotusClient)
		assert.Error(t, err, "Batch processing should error when service is disabled")
		assert.Contains(t, err.Error(), "disabled")

		// Re-enable for manual trigger test
		triggerService.SetEnabled(true)

		// Test 4: Manual trigger for non-existent preparation
		err = triggerService.TriggerForPreparation(ctx, db, lotusClient, "99999")
		assert.Error(t, err, "Should error for non-existent preparation")
	})
}
