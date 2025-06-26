package workflow

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/job"
	"github.com/data-preservation-programs/singularity/handler/notification"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/service/autodeal"
	"github.com/ipfs/go-log/v2"
	"github.com/ybbus/jsonrpc/v3"
	"gorm.io/gorm"
)

var logger = log.Logger("workflow-orchestrator")

// WorkflowOrchestrator manages automatic job progression through scan → pack → daggen → deals
type WorkflowOrchestrator struct {
	notificationHandler *notification.Handler
	triggerService      *autodeal.TriggerService
	jobHandler          *job.DefaultHandler
	mutex               sync.RWMutex
	enabled             bool
	config              OrchestratorConfig
	preparationLocks    map[uint]*sync.Mutex // Per-preparation locks for workflow transitions
	locksMutex          sync.RWMutex         // Protects the preparationLocks map
}

// OrchestratorConfig configures the workflow orchestrator
type OrchestratorConfig struct {
	EnableJobProgression bool          `json:"enableJobProgression"` // Enable automatic scan → pack → daggen
	EnableAutoDeal       bool          `json:"enableAutoDeal"`       // Enable automatic deal creation
	CheckInterval        time.Duration `json:"checkInterval"`        // How often to check for ready jobs
	ScanToPack           bool          `json:"scanToPack"`           // Auto-progress scan → pack
	PackToDagGen         bool          `json:"packToDagGen"`         // Auto-progress pack → daggen
	DagGenToDeals        bool          `json:"dagGenToDeals"`        // Auto-progress daggen → deals
	RetryEnabled         bool          `json:"retryEnabled"`         // Enable database retry for job creation
}

// DefaultOrchestratorConfig returns sensible defaults
func DefaultOrchestratorConfig() OrchestratorConfig {
	return OrchestratorConfig{
		EnableJobProgression: true,
		EnableAutoDeal:       true,
		CheckInterval:        10 * time.Second,
		ScanToPack:           true,
		PackToDagGen:         true,
		DagGenToDeals:        true,
		RetryEnabled:         true,
	}
}

// NewWorkflowOrchestrator creates a new workflow orchestrator
func NewWorkflowOrchestrator(config OrchestratorConfig) *WorkflowOrchestrator {
	return &WorkflowOrchestrator{
		notificationHandler: notification.Default,
		triggerService:      autodeal.DefaultTriggerService,
		jobHandler:          &job.DefaultHandler{},
		enabled:             true,
		config:              config,
		preparationLocks:    make(map[uint]*sync.Mutex),
	}
}

var DefaultOrchestrator = NewWorkflowOrchestrator(DefaultOrchestratorConfig())

// SetEnabled enables or disables the workflow orchestrator
func (o *WorkflowOrchestrator) SetEnabled(enabled bool) {
	o.mutex.Lock()
	defer o.mutex.Unlock()
	o.enabled = enabled
	logger.Infof("Workflow orchestrator enabled: %t", enabled)
}

// IsEnabled returns whether the orchestrator is enabled
func (o *WorkflowOrchestrator) IsEnabled() bool {
	o.mutex.RLock()
	defer o.mutex.RUnlock()
	return o.enabled
}

// lockPreparation acquires a lock for a specific preparation to prevent concurrent workflow transitions
func (o *WorkflowOrchestrator) lockPreparation(preparationID uint) {
	o.locksMutex.Lock()
	if _, exists := o.preparationLocks[preparationID]; !exists {
		o.preparationLocks[preparationID] = &sync.Mutex{}
	}
	mutex := o.preparationLocks[preparationID]
	o.locksMutex.Unlock()

	mutex.Lock()
}

// unlockPreparation releases the lock for a specific preparation
func (o *WorkflowOrchestrator) unlockPreparation(preparationID uint) {
	o.locksMutex.RLock()
	mutex := o.preparationLocks[preparationID]
	o.locksMutex.RUnlock()

	if mutex != nil {
		mutex.Unlock()
	}
}

// HandleJobCompletion processes job completion and triggers next stage if appropriate
func (o *WorkflowOrchestrator) HandleJobCompletion(
	ctx context.Context,
	db *gorm.DB,
	lotusClient jsonrpc.RPCClient,
	jobID model.JobID,
) error {
	if !o.IsEnabled() {
		return nil
	}

	// Get the completed job details
	var job model.Job
	err := db.WithContext(ctx).
		Joins("Attachment").
		Joins("Attachment.Preparation").
		First(&job, jobID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Warnf("Job %d not found during workflow orchestration", jobID)
			return nil
		}
		return errors.WithStack(err)
	}

	preparation := job.Attachment.Preparation
	logger.Infof("Processing job completion: JobID=%d, Type=%s, Preparation=%s",
		jobID, job.Type, preparation.Name)

	// Acquire preparation-specific lock to prevent concurrent workflow transitions
	o.lockPreparation(uint(preparation.ID))
	defer o.unlockPreparation(uint(preparation.ID))

	// Handle job progression based on type
	switch job.Type {
	case model.Scan:
		if o.config.ScanToPack {
			return o.handleScanCompletion(ctx, db, lotusClient, preparation)
		}
	case model.Pack:
		if o.config.PackToDagGen {
			return o.handlePackCompletion(ctx, db, lotusClient, preparation)
		}
	case model.DagGen:
		if o.config.DagGenToDeals {
			return o.handleDagGenCompletion(ctx, db, lotusClient, preparation)
		}
	}

	return nil
}

// handleScanCompletion triggers pack jobs after all scan jobs complete
func (o *WorkflowOrchestrator) handleScanCompletion(
	ctx context.Context,
	db *gorm.DB,
	lotusClient jsonrpc.RPCClient,
	preparation *model.Preparation,
) error {
	// Check if all scan jobs for this preparation are complete
	var incompleteScanCount int64
	err := db.WithContext(ctx).Model(&model.Job{}).
		Joins("JOIN source_attachments ON jobs.attachment_id = source_attachments.id").
		Where("source_attachments.preparation_id = ? AND jobs.type = ? AND jobs.state != ?",
			preparation.ID, model.Scan, model.Complete).
		Count(&incompleteScanCount).Error
	if err != nil {
		return errors.WithStack(err)
	}

	if incompleteScanCount > 0 {
		logger.Debugf("Preparation %s still has %d incomplete scan jobs",
			preparation.Name, incompleteScanCount)
		return nil
	}

	logger.Infof("All scan jobs complete for preparation %s, starting pack jobs", preparation.Name)

	// Use a transaction to ensure atomicity when starting pack jobs
	err = db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Re-check scan job completion within transaction to prevent race conditions
		var incompleteScanCount int64
		err := tx.Model(&model.Job{}).
			Joins("JOIN source_attachments ON jobs.attachment_id = source_attachments.id").
			Where("source_attachments.preparation_id = ? AND jobs.type = ? AND jobs.state != ?",
				preparation.ID, model.Scan, model.Complete).
			Count(&incompleteScanCount).Error
		if err != nil {
			return errors.WithStack(err)
		}

		if incompleteScanCount > 0 {
			logger.Debugf("Preparation %s still has %d incomplete scan jobs (double-checked in transaction)",
				preparation.Name, incompleteScanCount)
			return nil // No error, just nothing to do
		}

		// Check if pack jobs have already been started (prevent duplicate creation)
		var existingPackCount int64
		err = tx.Model(&model.Job{}).
			Joins("JOIN source_attachments ON jobs.attachment_id = source_attachments.id").
			Where("source_attachments.preparation_id = ? AND jobs.type = ?",
				preparation.ID, model.Pack).
			Count(&existingPackCount).Error
		if err != nil {
			return errors.WithStack(err)
		}

		if existingPackCount > 0 {
			logger.Debugf("Pack jobs already exist for preparation %s, skipping", preparation.Name)
			return nil
		}

		// Start pack jobs for all source attachments
		var attachments []model.SourceAttachment
		err = tx.Where("preparation_id = ?", preparation.ID).Find(&attachments).Error
		if err != nil {
			return errors.WithStack(err)
		}

		for _, attachment := range attachments {
			err = o.startPackJobs(ctx, tx, uint(attachment.ID))
			if err != nil {
				logger.Errorf("Failed to start pack jobs for attachment %d: %v", attachment.ID, err)
				return errors.WithStack(err) // Fail the transaction on any error
			}
		}

		return nil
	})
	if err != nil {
		return errors.WithStack(err)
	}

	o.logWorkflowProgress(ctx, db, "Scan → Pack Transition",
		fmt.Sprintf("Started pack jobs for preparation %s", preparation.Name),
		model.ConfigMap{
			"preparation_id":   fmt.Sprintf("%d", preparation.ID),
			"preparation_name": preparation.Name,
			"stage":            "scan_to_pack",
		})

	return nil
}

// handlePackCompletion triggers daggen jobs after all pack jobs complete
func (o *WorkflowOrchestrator) handlePackCompletion(
	ctx context.Context,
	db *gorm.DB,
	lotusClient jsonrpc.RPCClient,
	preparation *model.Preparation,
) error {
	// Check if all pack jobs for this preparation are complete
	var incompletePackCount int64
	err := db.WithContext(ctx).Model(&model.Job{}).
		Joins("JOIN source_attachments ON jobs.attachment_id = source_attachments.id").
		Where("source_attachments.preparation_id = ? AND jobs.type = ? AND jobs.state != ?",
			preparation.ID, model.Pack, model.Complete).
		Count(&incompletePackCount).Error
	if err != nil {
		return errors.WithStack(err)
	}

	if incompletePackCount > 0 {
		logger.Debugf("Preparation %s still has %d incomplete pack jobs",
			preparation.Name, incompletePackCount)
		return nil
	}

	// Skip daggen if NoDag is enabled
	if preparation.NoDag {
		logger.Infof("Preparation %s has NoDag enabled, skipping to deal creation", preparation.Name)
		return o.handleDagGenCompletion(ctx, db, lotusClient, preparation)
	}

	logger.Infof("All pack jobs complete for preparation %s, starting daggen jobs", preparation.Name)

	// Use a transaction to ensure atomicity when starting daggen jobs
	err = db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Re-check pack job completion within transaction to prevent race conditions
		var incompletePackCount int64
		err := tx.Model(&model.Job{}).
			Joins("JOIN source_attachments ON jobs.attachment_id = source_attachments.id").
			Where("source_attachments.preparation_id = ? AND jobs.type = ? AND jobs.state != ?",
				preparation.ID, model.Pack, model.Complete).
			Count(&incompletePackCount).Error
		if err != nil {
			return errors.WithStack(err)
		}

		if incompletePackCount > 0 {
			logger.Debugf("Preparation %s still has %d incomplete pack jobs (double-checked in transaction)",
				preparation.Name, incompletePackCount)
			return nil // No error, just nothing to do
		}

		// Check if daggen jobs have already been started (prevent duplicate creation)
		var existingDagGenCount int64
		err = tx.Model(&model.Job{}).
			Joins("JOIN source_attachments ON jobs.attachment_id = source_attachments.id").
			Where("source_attachments.preparation_id = ? AND jobs.type = ?",
				preparation.ID, model.DagGen).
			Count(&existingDagGenCount).Error
		if err != nil {
			return errors.WithStack(err)
		}

		if existingDagGenCount > 0 {
			logger.Debugf("DagGen jobs already exist for preparation %s, skipping", preparation.Name)
			return nil
		}

		// Start daggen jobs for all source attachments
		var attachments []model.SourceAttachment
		err = tx.Where("preparation_id = ?", preparation.ID).Find(&attachments).Error
		if err != nil {
			return errors.WithStack(err)
		}

		for _, attachment := range attachments {
			err = o.startDagGenJobs(ctx, tx, uint(attachment.ID))
			if err != nil {
				logger.Errorf("Failed to start daggen jobs for attachment %d: %v", attachment.ID, err)
				return errors.WithStack(err) // Fail the transaction on any error
			}
		}

		return nil
	})
	if err != nil {
		return errors.WithStack(err)
	}

	o.logWorkflowProgress(ctx, db, "Pack → DagGen Transition",
		fmt.Sprintf("Started daggen jobs for preparation %s", preparation.Name),
		model.ConfigMap{
			"preparation_id":   fmt.Sprintf("%d", preparation.ID),
			"preparation_name": preparation.Name,
			"stage":            "pack_to_daggen",
		})

	return nil
}

// handleDagGenCompletion triggers auto-deal creation after all daggen jobs complete
func (o *WorkflowOrchestrator) handleDagGenCompletion(
	ctx context.Context,
	db *gorm.DB,
	lotusClient jsonrpc.RPCClient,
	preparation *model.Preparation,
) error {
	if !o.config.EnableAutoDeal {
		logger.Debugf("Auto-deal creation disabled for preparation %s", preparation.Name)
		return nil
	}

	// Check if all jobs for this preparation are complete
	var incompleteJobCount int64
	err := db.WithContext(ctx).Model(&model.Job{}).
		Joins("JOIN source_attachments ON jobs.attachment_id = source_attachments.id").
		Where("source_attachments.preparation_id = ? AND jobs.state != ?",
			preparation.ID, model.Complete).
		Count(&incompleteJobCount).Error
	if err != nil {
		return errors.WithStack(err)
	}

	if incompleteJobCount > 0 {
		logger.Debugf("Preparation %s still has %d incomplete jobs",
			preparation.Name, incompleteJobCount)
		return nil
	}

	logger.Infof("All jobs complete for preparation %s, triggering auto-deal creation", preparation.Name)

	// Trigger auto-deal creation using existing service
	err = o.triggerService.TriggerForPreparation(ctx, db, lotusClient, fmt.Sprintf("%d", preparation.ID))
	if err != nil {
		logger.Errorf("Failed to create auto-deal for preparation %s: %v", preparation.Name, err)
		return errors.WithStack(err)
	}

	o.logWorkflowProgress(ctx, db, "DagGen → Deals Transition",
		fmt.Sprintf("Triggered auto-deal creation for preparation %s", preparation.Name),
		model.ConfigMap{
			"preparation_id":   fmt.Sprintf("%d", preparation.ID),
			"preparation_name": preparation.Name,
			"stage":            "daggen_to_deals",
		})

	return nil
}

// startPackJobs starts pack jobs for a source attachment
func (o *WorkflowOrchestrator) startPackJobs(ctx context.Context, db *gorm.DB, attachmentID uint) error {
	// Load the attachment with its associations
	var attachment model.SourceAttachment
	err := db.Preload("Preparation").Preload("Storage").First(&attachment, attachmentID).Error
	if err != nil {
		return errors.Wrapf(err, "failed to load source attachment %d", attachmentID)
	}

	if o.config.RetryEnabled {
		return database.DoRetry(ctx, func() error {
			_, err := o.jobHandler.StartPackHandler(ctx, db, attachment.Preparation.Name, attachment.Storage.Name, 0)
			return err
		})
	} else {
		_, err := o.jobHandler.StartPackHandler(ctx, db, attachment.Preparation.Name, attachment.Storage.Name, 0)
		return errors.WithStack(err)
	}
}

// startDagGenJobs starts daggen jobs for a source attachment
func (o *WorkflowOrchestrator) startDagGenJobs(ctx context.Context, db *gorm.DB, attachmentID uint) error {
	// Load the attachment with its associations
	var attachment model.SourceAttachment
	err := db.Preload("Preparation").Preload("Storage").First(&attachment, attachmentID).Error
	if err != nil {
		return errors.Wrapf(err, "failed to load source attachment %d", attachmentID)
	}

	if o.config.RetryEnabled {
		return database.DoRetry(ctx, func() error {
			_, err := o.jobHandler.StartDagGenHandler(ctx, db, attachment.Preparation.Name, attachment.Storage.Name)
			return err
		})
	} else {
		_, err := o.jobHandler.StartDagGenHandler(ctx, db, attachment.Preparation.Name, attachment.Storage.Name)
		return errors.WithStack(err)
	}
}

// logWorkflowProgress logs workflow progression events
func (o *WorkflowOrchestrator) logWorkflowProgress(ctx context.Context, db *gorm.DB, title, message string, metadata model.ConfigMap) {
	_, err := o.notificationHandler.LogInfo(ctx, db, "workflow-orchestrator", title, message, metadata)
	if err != nil {
		logger.Errorf("Failed to log workflow progress: %v", err)
	}
}

// ProcessPendingWorkflows processes preparations that need workflow progression
func (o *WorkflowOrchestrator) ProcessPendingWorkflows(
	ctx context.Context,
	db *gorm.DB,
	lotusClient jsonrpc.RPCClient,
) error {
	if !o.IsEnabled() {
		return nil
	}

	logger.Debug("Checking for preparations needing workflow progression")

	// Find preparations that might need progression
	var preparations []model.Preparation
	err := db.WithContext(ctx).Find(&preparations).Error
	if err != nil {
		return errors.WithStack(err)
	}

	for _, prep := range preparations {
		err = o.checkPreparationWorkflow(ctx, db, lotusClient, &prep)
		if err != nil {
			logger.Errorf("Failed to check workflow for preparation %s: %v", prep.Name, err)
			continue
		}
	}

	return nil
}

// checkPreparationWorkflow checks if a preparation needs workflow progression
func (o *WorkflowOrchestrator) checkPreparationWorkflow(
	ctx context.Context,
	db *gorm.DB,
	lotusClient jsonrpc.RPCClient,
	preparation *model.Preparation,
) error {
	// Acquire preparation-specific lock to prevent concurrent workflow transitions
	o.lockPreparation(uint(preparation.ID))
	defer o.unlockPreparation(uint(preparation.ID))
	// Get job counts by type and state
	type JobCount struct {
		Type  model.JobType  `json:"type"`
		State model.JobState `json:"state"`
		Count int64          `json:"count"`
	}

	var jobCounts []JobCount
	err := db.WithContext(ctx).Model(&model.Job{}).
		Select("type, state, count(*) as count").
		Joins("JOIN source_attachments ON jobs.attachment_id = source_attachments.id").
		Where("source_attachments.preparation_id = ?", preparation.ID).
		Group("type, state").
		Find(&jobCounts).Error
	if err != nil {
		return errors.WithStack(err)
	}

	// Analyze job state to determine if progression is needed
	scanComplete := true
	packComplete := true
	hasPackJobs := false
	hasDagGenJobs := false

	for _, jc := range jobCounts {
		switch jc.Type {
		case model.Scan:
			if jc.State != model.Complete {
				scanComplete = false
			}
		case model.Pack:
			hasPackJobs = true
			if jc.State != model.Complete {
				packComplete = false
			}
		case model.DagGen:
			hasDagGenJobs = true
		}
	}

	// Trigger appropriate progression
	if scanComplete && !hasPackJobs && o.config.ScanToPack {
		logger.Debugf("Triggering pack jobs for preparation %s", preparation.Name)
		return o.handleScanCompletion(ctx, db, lotusClient, preparation)
	}

	if packComplete && hasPackJobs && !hasDagGenJobs && o.config.PackToDagGen {
		logger.Debugf("Triggering daggen jobs for preparation %s", preparation.Name)
		return o.handlePackCompletion(ctx, db, lotusClient, preparation)
	}

	return nil
}
