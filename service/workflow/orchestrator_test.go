package workflow

import (
	"context"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/handler/job"
	"github.com/data-preservation-programs/singularity/handler/notification"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/service/autodeal"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestDefaultOrchestratorConfig(t *testing.T) {
	config := DefaultOrchestratorConfig()
	assert.True(t, config.EnableJobProgression)
	assert.True(t, config.EnableAutoDeal)
	assert.Equal(t, 10*time.Second, config.CheckInterval)
	assert.True(t, config.ScanToPack)
	assert.True(t, config.PackToDagGen)
	assert.True(t, config.DagGenToDeals)
}

func TestNewWorkflowOrchestrator(t *testing.T) {
	config := DefaultOrchestratorConfig()
	orchestrator := NewWorkflowOrchestrator(config)

	assert.NotNil(t, orchestrator)
	assert.Equal(t, config, orchestrator.config)
	assert.True(t, orchestrator.enabled)
	assert.NotNil(t, orchestrator.notificationHandler)
	assert.NotNil(t, orchestrator.triggerService)
	assert.NotNil(t, orchestrator.jobHandler)
}

func TestWorkflowOrchestrator_SetEnabled(t *testing.T) {
	orchestrator := NewWorkflowOrchestrator(DefaultOrchestratorConfig())

	// Test enabling/disabling
	orchestrator.SetEnabled(false)
	assert.False(t, orchestrator.IsEnabled())

	orchestrator.SetEnabled(true)
	assert.True(t, orchestrator.IsEnabled())
}

func TestWorkflowOrchestrator_HandleJobCompletion_Disabled(t *testing.T) {
	testutil.One(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		orchestrator := NewWorkflowOrchestrator(DefaultOrchestratorConfig())
		orchestrator.SetEnabled(false)

		err := orchestrator.HandleJobCompletion(ctx, db, nil, 1)
		assert.NoError(t, err)
	})
}

func TestWorkflowOrchestrator_HandleJobCompletion_JobNotFound(t *testing.T) {
	testutil.One(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		orchestrator := NewWorkflowOrchestrator(DefaultOrchestratorConfig())

		err := orchestrator.HandleJobCompletion(ctx, db, nil, 99999)
		assert.NoError(t, err) // Should not error for missing job
	})
}

func TestWorkflowOrchestrator_HandleScanCompletion(t *testing.T) {
	testutil.One(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		// Set up test data
		preparation := &model.Preparation{
			Name: "test-prep",
			SourceStorages: []model.Storage{
				{
					Name: "test-storage",
					Type: "local",
					Path: "/tmp/test",
				},
			},
		}
		require.NoError(t, db.Create(preparation).Error)

		// Source attachment is created automatically by GORM when creating preparation with SourceStorages
		var sourceAttachment model.SourceAttachment
		require.NoError(t, db.Where("preparation_id = ? AND storage_id = ?", preparation.ID, preparation.SourceStorages[0].ID).First(&sourceAttachment).Error)

		// Create a completed scan job
		scanJob := &model.Job{
			Type:         model.Scan,
			State:        model.Complete,
			AttachmentID: sourceAttachment.ID,
		}
		require.NoError(t, db.Create(scanJob).Error)

		// Create mock handlers
		orchestrator := NewWorkflowOrchestrator(DefaultOrchestratorConfig())
		orchestrator.jobHandler = &job.DefaultHandler{}
		orchestrator.notificationHandler = notification.Default

		// Test scan completion handling
		err := orchestrator.HandleJobCompletion(ctx, db, nil, scanJob.ID)

		// Should not error (though actual pack job creation may fail due to missing setup)
		assert.NoError(t, err)
	})
}

func TestWorkflowOrchestrator_HandleScanCompletion_IncompleteScanJobs(t *testing.T) {
	testutil.One(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		// Set up test data
		preparation := &model.Preparation{
			Name: "test-prep",
			SourceStorages: []model.Storage{
				{
					Name: "test-storage",
					Type: "local",
					Path: "/tmp/test",
				},
			},
		}
		require.NoError(t, db.Create(preparation).Error)

		// Source attachment is created automatically by GORM when creating preparation with SourceStorages
		var sourceAttachment model.SourceAttachment
		require.NoError(t, db.Where("preparation_id = ? AND storage_id = ?", preparation.ID, preparation.SourceStorages[0].ID).First(&sourceAttachment).Error)

		// Create completed and incomplete scan jobs
		completedScanJob := &model.Job{
			Type:         model.Scan,
			State:        model.Complete,
			AttachmentID: sourceAttachment.ID,
		}
		require.NoError(t, db.Create(completedScanJob).Error)

		incompleteScanJob := &model.Job{
			Type:         model.Scan,
			State:        model.Processing,
			AttachmentID: sourceAttachment.ID,
		}
		require.NoError(t, db.Create(incompleteScanJob).Error)

		orchestrator := NewWorkflowOrchestrator(DefaultOrchestratorConfig())

		// Test that pack jobs are not started when scan jobs are incomplete
		err := orchestrator.handleScanCompletion(ctx, db, nil, preparation)
		assert.NoError(t, err)

		// Verify no pack jobs were created
		var packJobCount int64
		err = db.Model(&model.Job{}).
			Joins("JOIN source_attachments ON jobs.attachment_id = source_attachments.id").
			Where("source_attachments.preparation_id = ? AND jobs.type = ?", preparation.ID, model.Pack).
			Count(&packJobCount).Error
		require.NoError(t, err)
		assert.Equal(t, int64(0), packJobCount)
	})
}

func TestWorkflowOrchestrator_HandlePackCompletion_NoDag(t *testing.T) {
	testutil.One(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		// Set up test data with NoDag enabled
		preparation := &model.Preparation{
			Name:  "test-prep",
			NoDag: true,
			SourceStorages: []model.Storage{
				{
					Name: "test-storage",
					Type: "local",
					Path: "/tmp/test",
				},
			},
		}
		require.NoError(t, db.Create(preparation).Error)

		// Source attachment is created automatically by GORM when creating preparation with SourceStorages
		var sourceAttachment model.SourceAttachment
		require.NoError(t, db.Where("preparation_id = ? AND storage_id = ?", preparation.ID, preparation.SourceStorages[0].ID).First(&sourceAttachment).Error)

		// Create a completed pack job
		packJob := &model.Job{
			Type:         model.Pack,
			State:        model.Complete,
			AttachmentID: sourceAttachment.ID,
		}
		require.NoError(t, db.Create(packJob).Error)

		orchestrator := NewWorkflowOrchestrator(DefaultOrchestratorConfig())
		triggerService := autodeal.NewTriggerService()
		triggerService.SetEnabled(true)
		orchestrator.triggerService = triggerService

		// Test pack completion with NoDag - should skip directly to deal creation
		err := orchestrator.handlePackCompletion(ctx, db, nil, preparation)

		// Should not error (though auto-deal creation may fail due to missing setup)
		assert.NoError(t, err)
	})
}

func TestWorkflowOrchestrator_ProcessPendingWorkflows_Disabled(t *testing.T) {
	testutil.One(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		orchestrator := NewWorkflowOrchestrator(DefaultOrchestratorConfig())
		orchestrator.SetEnabled(false)

		err := orchestrator.ProcessPendingWorkflows(ctx, db, nil)
		assert.NoError(t, err)
	})
}

func TestWorkflowOrchestrator_ProcessPendingWorkflows(t *testing.T) {
	testutil.One(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		// Set up test data
		preparation := &model.Preparation{
			Name: "test-prep",
			SourceStorages: []model.Storage{
				{
					Name: "test-storage",
					Type: "local",
					Path: "/tmp/test",
				},
			},
		}
		require.NoError(t, db.Create(preparation).Error)

		orchestrator := NewWorkflowOrchestrator(DefaultOrchestratorConfig())

		err := orchestrator.ProcessPendingWorkflows(ctx, db, nil)
		assert.NoError(t, err)
	})
}

func TestWorkflowOrchestrator_CheckPreparationWorkflow(t *testing.T) {
	testutil.One(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		// Set up test data
		preparation := &model.Preparation{
			Name: "test-prep",
			SourceStorages: []model.Storage{
				{
					Name: "test-storage",
					Type: "local",
					Path: "/tmp/test",
				},
			},
		}
		require.NoError(t, db.Create(preparation).Error)

		// Source attachment is created automatically by GORM when creating preparation with SourceStorages
		var sourceAttachment model.SourceAttachment
		require.NoError(t, db.Where("preparation_id = ? AND storage_id = ?", preparation.ID, preparation.SourceStorages[0].ID).First(&sourceAttachment).Error)

		// Create a completed scan job
		scanJob := &model.Job{
			Type:         model.Scan,
			State:        model.Complete,
			AttachmentID: sourceAttachment.ID,
		}
		require.NoError(t, db.Create(scanJob).Error)

		orchestrator := NewWorkflowOrchestrator(DefaultOrchestratorConfig())
		orchestrator.jobHandler = &job.DefaultHandler{}
		orchestrator.notificationHandler = notification.Default

		err := orchestrator.checkPreparationWorkflow(ctx, db, nil, preparation)

		// Should not error (though actual pack job creation may fail due to missing setup)
		assert.NoError(t, err)
	})
}

func TestWorkflowOrchestrator_ConfigurationDisabled(t *testing.T) {
	testutil.One(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		// Test with all workflow stages disabled
		config := OrchestratorConfig{
			EnableJobProgression: false,
			EnableAutoDeal:       false,
			ScanToPack:           false,
			PackToDagGen:         false,
			DagGenToDeals:        false,
		}

		orchestrator := NewWorkflowOrchestrator(config)

		// Set up test data
		preparation := &model.Preparation{
			Name: "test-prep",
			SourceStorages: []model.Storage{
				{
					Name: "test-storage",
					Type: "local",
					Path: "/tmp/test",
				},
			},
		}
		require.NoError(t, db.Create(preparation).Error)

		// Source attachment is created automatically by GORM when creating preparation with SourceStorages
		var sourceAttachment model.SourceAttachment
		require.NoError(t, db.Where("preparation_id = ? AND storage_id = ?", preparation.ID, preparation.SourceStorages[0].ID).First(&sourceAttachment).Error)

		scanJob := &model.Job{
			Type:         model.Scan,
			State:        model.Complete,
			AttachmentID: sourceAttachment.ID,
		}
		require.NoError(t, db.Create(scanJob).Error)

		// Should do nothing when workflow stages are disabled
		err := orchestrator.HandleJobCompletion(ctx, db, nil, scanJob.ID)
		assert.NoError(t, err)

		// Verify no pack jobs were created
		var packJobCount int64
		err = db.Model(&model.Job{}).
			Joins("JOIN source_attachments ON jobs.attachment_id = source_attachments.id").
			Where("source_attachments.preparation_id = ? AND jobs.type = ?", preparation.ID, model.Pack).
			Count(&packJobCount).Error
		require.NoError(t, err)
		assert.Equal(t, int64(0), packJobCount)
	})
}
