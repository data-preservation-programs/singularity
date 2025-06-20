package autodeal

import (
	"context"
	"fmt"
	"sync"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/handler/dataprep"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/ipfs/go-log/v2"
	"github.com/ybbus/jsonrpc/v3"
	"gorm.io/gorm"
)

// AutoDealServiceInterface defines the interface for auto-deal services
type AutoDealServiceInterface interface {
	CheckPreparationReadiness(ctx context.Context, db *gorm.DB, preparationID string) (bool, error)
	CreateAutomaticDealSchedule(ctx context.Context, db *gorm.DB, lotusClient jsonrpc.RPCClient, preparationID string) (*model.Schedule, error)
	ProcessReadyPreparations(ctx context.Context, db *gorm.DB, lotusClient jsonrpc.RPCClient) error
}

var logger = log.Logger("autodeal-trigger")

// TriggerService handles automatic deal creation when preparations complete
type TriggerService struct {
	autoDealService AutoDealServiceInterface
	mutex           sync.RWMutex
	enabled         bool
}

// NewTriggerService creates a new auto-deal trigger service
func NewTriggerService() *TriggerService {
	return &TriggerService{
		autoDealService: dataprep.DefaultAutoDealService,
		enabled:         true,
	}
}

// SetAutoDealService sets the auto-deal service implementation (for testing)
func (s *TriggerService) SetAutoDealService(service AutoDealServiceInterface) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.autoDealService = service
}

// DefaultTriggerService is the default instance
var DefaultTriggerService = NewTriggerService()

// SetEnabled enables or disables the auto-deal trigger service
func (s *TriggerService) SetEnabled(enabled bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.enabled = enabled
	logger.Infof("Auto-deal trigger service enabled: %t", enabled)
}

// IsEnabled returns whether the auto-deal trigger service is enabled
func (s *TriggerService) IsEnabled() bool {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.enabled
}

// TriggerForJobCompletion checks if a job completion should trigger auto-deal creation
// This method is called when any job completes
func (s *TriggerService) TriggerForJobCompletion(
	ctx context.Context,
	db *gorm.DB,
	lotusClient jsonrpc.RPCClient,
	jobID model.JobID,
) error {
	if !s.IsEnabled() {
		return nil
	}

	// Get the job and its preparation
	var job model.Job
	err := db.WithContext(ctx).
		Joins("Attachment").
		Joins("Attachment.Preparation").
		First(&job, jobID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Warnf("Job %d not found during auto-deal trigger check", jobID)
			return nil
		}
		return errors.WithStack(err)
	}

	// Check if preparation has auto-deal enabled
	if !job.Attachment.Preparation.DealConfig.AutoCreateDeals {
		logger.Debugf("Preparation %s does not have auto-deal enabled, skipping trigger",
			job.Attachment.Preparation.Name)
		return nil
	}

	preparationID := fmt.Sprintf("%d", job.Attachment.Preparation.ID)

	logger.Debugf("Job %d completed for preparation %s with auto-deal enabled, checking readiness",
		jobID, job.Attachment.Preparation.Name)

	// Check if all jobs for this preparation are complete
	isReady, err := s.autoDealService.CheckPreparationReadiness(ctx, db, preparationID)
	if err != nil {
		logger.Errorf("Failed to check preparation readiness for %s: %v",
			job.Attachment.Preparation.Name, err)
		return errors.WithStack(err)
	}

	if !isReady {
		logger.Debugf("Preparation %s is not ready yet, other jobs still in progress",
			job.Attachment.Preparation.Name)
		return nil
	}

	// Check if deal schedule already exists
	var existingScheduleCount int64
	err = db.WithContext(ctx).Model(&model.Schedule{}).
		Where("preparation_id = ?", job.Attachment.Preparation.ID).
		Count(&existingScheduleCount).Error
	if err != nil {
		return errors.WithStack(err)
	}

	if existingScheduleCount > 0 {
		logger.Debugf("Preparation %s already has %d deal schedule(s), skipping auto-creation",
			job.Attachment.Preparation.Name, existingScheduleCount)
		return nil
	}

	logger.Infof("Triggering automatic deal creation for preparation %s",
		job.Attachment.Preparation.Name)

	// Create the deal schedule automatically
	schedule, err := s.autoDealService.CreateAutomaticDealSchedule(ctx, db, lotusClient, preparationID)
	if err != nil {
		logger.Errorf("Failed to create automatic deal schedule for preparation %s: %v",
			job.Attachment.Preparation.Name, err)
		return errors.WithStack(err)
	}

	if schedule != nil {
		logger.Infof("Successfully created automatic deal schedule %d for preparation %s",
			schedule.ID, job.Attachment.Preparation.Name)
	}

	return nil
}

// TriggerForPreparation manually triggers auto-deal creation for a specific preparation
func (s *TriggerService) TriggerForPreparation(
	ctx context.Context,
	db *gorm.DB,
	lotusClient jsonrpc.RPCClient,
	preparationID string,
) error {
	if !s.IsEnabled() {
		return errors.New("auto-deal trigger service is disabled")
	}

	logger.Infof("Manual trigger for preparation %s", preparationID)

	schedule, err := s.autoDealService.CreateAutomaticDealSchedule(ctx, db, lotusClient, preparationID)
	if err != nil {
		return errors.WithStack(err)
	}

	if schedule != nil {
		logger.Infof("Successfully created deal schedule %d for preparation %s",
			schedule.ID, preparationID)
	}

	return nil
}

// BatchProcessReadyPreparations processes all preparations that are ready for auto-deal creation
func (s *TriggerService) BatchProcessReadyPreparations(
	ctx context.Context,
	db *gorm.DB,
	lotusClient jsonrpc.RPCClient,
) error {
	if !s.IsEnabled() {
		return errors.New("auto-deal trigger service is disabled")
	}

	logger.Info("Starting batch processing of ready preparations")

	err := s.autoDealService.ProcessReadyPreparations(ctx, db, lotusClient)
	if err != nil {
		return errors.WithStack(err)
	}

	logger.Info("Batch processing completed")
	return nil
}
