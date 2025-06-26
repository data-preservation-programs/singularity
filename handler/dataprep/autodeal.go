package dataprep

import (
	"context"
	"fmt"
	"strconv"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/handler/deal/schedule"
	"github.com/data-preservation-programs/singularity/handler/notification"
	"github.com/data-preservation-programs/singularity/handler/storage"
	"github.com/data-preservation-programs/singularity/handler/wallet"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/ipfs/go-log/v2"
	"github.com/ybbus/jsonrpc/v3"
	"gorm.io/gorm"
)

var autoDealLogger = log.Logger("auto-deal")

type AutoDealService struct {
	notificationHandler *notification.Handler
	scheduleHandler     schedule.Handler
	walletValidator     *wallet.BalanceValidator
	spValidator         *storage.SPValidator
}

func NewAutoDealService() *AutoDealService {
	return &AutoDealService{
		notificationHandler: notification.Default,
		scheduleHandler:     schedule.Default,
		walletValidator:     wallet.DefaultBalanceValidator,
		spValidator:         storage.DefaultSPValidator,
	}
}

var DefaultAutoDealService = NewAutoDealService()

// CreateAutomaticDealSchedule creates deal schedules automatically for preparations with auto-deal enabled
func (s *AutoDealService) CreateAutomaticDealSchedule(
	ctx context.Context,
	db *gorm.DB,
	lotusClient jsonrpc.RPCClient,
	preparationID string,
) (*model.Schedule, error) {
	// Get preparation with auto-deal settings
	var preparation model.Preparation
	err := preparation.FindByIDOrName(db.WithContext(ctx), preparationID, "Wallets")
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(err, "preparation %s not found", preparationID)
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// Check if auto-deal creation is enabled
	if !preparation.DealConfig.AutoCreateDeals {
		s.logInfo(ctx, db, "Auto-Deal Not Enabled",
			"Preparation "+preparation.Name+" does not have auto-deal creation enabled",
			model.ConfigMap{
				"preparation_id":   preparationID,
				"preparation_name": preparation.Name,
			})
		return nil, nil
	}

	s.logInfo(ctx, db, "Starting Auto-Deal Schedule Creation",
		"Creating automatic deal schedule for preparation "+preparation.Name,
		model.ConfigMap{
			"preparation_id":   preparationID,
			"preparation_name": preparation.Name,
		})

	// Perform final validation before creating deals
	validationPassed := true
	validationErrors := []string{}

	if preparation.WalletValidation {
		err = s.validateWalletsForDealCreation(ctx, db, lotusClient, &preparation, &validationErrors)
		if err != nil {
			validationPassed = false
			s.logWarning(ctx, db, "Wallet Validation Failed",
				"Wallet validation failed during auto-deal creation",
				model.ConfigMap{
					"preparation_name": preparation.Name,
					"error":            err.Error(),
				})
		}
	}

	if preparation.SPValidation {
		err = s.validateProviderForDealCreation(ctx, db, lotusClient, &preparation, &validationErrors)
		if err != nil {
			validationPassed = false
			s.logWarning(ctx, db, "Provider Validation Failed",
				"Storage provider validation failed during auto-deal creation",
				model.ConfigMap{
					"preparation_name": preparation.Name,
					"error":            err.Error(),
				})
		}
	}

	// If validation failed, log and return
	if !validationPassed {
		s.logError(ctx, db, "Auto-Deal Creation Failed",
			"Auto-deal creation failed due to validation errors",
			model.ConfigMap{
				"preparation_name":  preparation.Name,
				"validation_errors": fmt.Sprintf("%v", validationErrors),
			})
		return nil, errors.New("auto-deal creation failed validation")
	}

	// Create the deal schedule using collected parameters
	dealRequest := s.buildDealScheduleRequest(&preparation)

	s.logInfo(ctx, db, "Creating Deal Schedule",
		"Creating deal schedule with provider "+dealRequest.Provider,
		model.ConfigMap{
			"preparation_name": preparation.Name,
			"provider":         dealRequest.Provider,
			"verified":         strconv.FormatBool(dealRequest.Verified),
			"price_per_gb":     fmt.Sprintf("%.6f", dealRequest.PricePerGB),
		})

	dealSchedule, err := s.scheduleHandler.CreateHandler(ctx, db, lotusClient, *dealRequest)
	if err != nil {
		s.logError(ctx, db, "Deal Schedule Creation Failed",
			"Failed to create automatic deal schedule",
			model.ConfigMap{
				"preparation_name": preparation.Name,
				"error":            err.Error(),
			})
		return nil, errors.WithStack(err)
	}

	s.logInfo(ctx, db, "Auto-Deal Schedule Created Successfully",
		fmt.Sprintf("Successfully created deal schedule %d for preparation %s", dealSchedule.ID, preparation.Name),
		model.ConfigMap{
			"preparation_name": preparation.Name,
			"schedule_id":      strconv.FormatUint(uint64(dealSchedule.ID), 10),
			"provider":         dealSchedule.Provider,
		})

	return dealSchedule, nil
}

// CheckPreparationReadiness checks if a preparation is ready for auto-deal creation
func (s *AutoDealService) CheckPreparationReadiness(
	ctx context.Context,
	db *gorm.DB,
	preparationID string,
) (bool, error) {
	// Check if all jobs for the preparation are complete
	var incompleteJobCount int64
	err := db.WithContext(ctx).Model(&model.Job{}).
		Joins("JOIN source_attachments ON jobs.attachment_id = source_attachments.id").
		Where("source_attachments.preparation_id = ? AND jobs.state != ?", preparationID, model.Complete).
		Count(&incompleteJobCount).Error
	if err != nil {
		return false, errors.WithStack(err)
	}

	isReady := incompleteJobCount == 0

	s.logInfo(ctx, db, "Preparation Readiness Check",
		fmt.Sprintf("Preparation %s readiness: %t (incomplete jobs: %d)", preparationID, isReady, incompleteJobCount),
		model.ConfigMap{
			"preparation_id":  preparationID,
			"is_ready":        strconv.FormatBool(isReady),
			"incomplete_jobs": strconv.FormatInt(incompleteJobCount, 10),
		})

	return isReady, nil
}

// ProcessReadyPreparations finds and processes all preparations ready for auto-deal creation
func (s *AutoDealService) ProcessReadyPreparations(
	ctx context.Context,
	db *gorm.DB,
	lotusClient jsonrpc.RPCClient,
) error {
	// Find preparations with auto-deal enabled that don't have schedules yet
	var preparations []model.Preparation
	err := db.WithContext(ctx).Preload("Wallets").
		Where("auto_create_deals = ?", true).
		Find(&preparations).Error
	if err != nil {
		return errors.WithStack(err)
	}

	s.logInfo(ctx, db, "Processing Ready Preparations",
		fmt.Sprintf("Found %d preparations with auto-deal enabled", len(preparations)),
		model.ConfigMap{
			"preparation_count": strconv.Itoa(len(preparations)),
		})

	processedCount := 0
	errorCount := 0

	for _, prep := range preparations {
		// Check if preparation already has a deal schedule
		var existingScheduleCount int64
		err = db.WithContext(ctx).Model(&model.Schedule{}).
			Where("preparation_id = ?", prep.ID).Count(&existingScheduleCount).Error
		if err != nil {
			autoDealLogger.Errorf("Failed to check existing schedules for preparation %s: %v", prep.Name, err)
			errorCount++
			continue
		}

		if existingScheduleCount > 0 {
			autoDealLogger.Debugf("Preparation %s already has %d schedule(s), skipping", prep.Name, existingScheduleCount)
			continue
		}

		// Check if preparation is ready
		isReady, err := s.CheckPreparationReadiness(ctx, db, fmt.Sprintf("%d", prep.ID))
		if err != nil {
			autoDealLogger.Errorf("Failed to check readiness for preparation %s: %v", prep.Name, err)
			errorCount++
			continue
		}

		if !isReady {
			autoDealLogger.Debugf("Preparation %s is not ready for deal creation yet", prep.Name)
			continue
		}

		// Create automatic deal schedule
		_, err = s.CreateAutomaticDealSchedule(ctx, db, lotusClient, fmt.Sprintf("%d", prep.ID))
		if err != nil {
			autoDealLogger.Errorf("Failed to create auto-deal schedule for preparation %s: %v", prep.Name, err)
			errorCount++
			continue
		}

		processedCount++
	}

	s.logInfo(ctx, db, "Auto-Deal Processing Complete",
		fmt.Sprintf("Processed %d preparations, %d errors", processedCount, errorCount),
		model.ConfigMap{
			"processed_count": strconv.Itoa(processedCount),
			"error_count":     strconv.Itoa(errorCount),
		})

	return nil
}

// buildDealScheduleRequest constructs a deal schedule create request from preparation parameters
func (s *AutoDealService) buildDealScheduleRequest(preparation *model.Preparation) *schedule.CreateRequest {
	request := &schedule.CreateRequest{
		Preparation:     strconv.FormatUint(uint64(preparation.ID), 10),
		Provider:        preparation.DealConfig.DealProvider,
		PricePerGBEpoch: preparation.DealConfig.DealPricePerGbEpoch,
		PricePerGB:      preparation.DealConfig.DealPricePerGb,
		PricePerDeal:    preparation.DealConfig.DealPricePerDeal,
		Verified:        preparation.DealConfig.DealVerified,
		IPNI:            preparation.DealConfig.DealAnnounceToIpni,
		KeepUnsealed:    preparation.DealConfig.DealKeepUnsealed,
		URLTemplate:     preparation.DealConfig.DealURLTemplate,
		Notes:           "Automatically created by auto-deal system",
	}

	// Convert HTTP headers from ConfigMap to []string
	var httpHeaders []string
	for key, value := range preparation.DealConfig.DealHTTPHeaders {
		httpHeaders = append(httpHeaders, key+"="+value)
	}
	request.HTTPHeaders = httpHeaders

	// Convert durations to strings
	if preparation.DealConfig.DealStartDelay > 0 {
		request.StartDelay = preparation.DealConfig.DealStartDelay.String()
	} else {
		request.StartDelay = "72h" // Default
	}

	if preparation.DealConfig.DealDuration > 0 {
		request.Duration = preparation.DealConfig.DealDuration.String()
	} else {
		request.Duration = "12840h" // Default (~535 days)
	}

	// If no provider specified, leave empty - the schedule handler will validate and potentially use default
	if request.Provider == "" {
		// The schedule creation will fail if no provider, but we've already validated this in preparation creation
		autoDealLogger.Warnf("No provider specified for preparation %s, deal creation may fail", preparation.Name)
	}

	return request
}

// validateWalletsForDealCreation performs wallet validation for deal creation
func (s *AutoDealService) validateWalletsForDealCreation(
	ctx context.Context,
	db *gorm.DB,
	lotusClient jsonrpc.RPCClient,
	preparation *model.Preparation,
	validationErrors *[]string,
) error {
	if len(preparation.Wallets) == 0 {
		*validationErrors = append(*validationErrors, "No wallets assigned to preparation")
		return errors.New("no wallets assigned")
	}

	// For now, just validate that wallets exist and are accessible
	// In a full implementation, you would calculate required balance based on data size
	for _, wallet := range preparation.Wallets {
		result, err := s.walletValidator.ValidateWalletExists(ctx, db, lotusClient, wallet.Address, strconv.FormatUint(uint64(preparation.ID), 10))
		if err != nil {
			*validationErrors = append(*validationErrors, fmt.Sprintf("Wallet validation error for %s: %v", wallet.Address, err))
			return err
		}
		if !result.IsValid {
			*validationErrors = append(*validationErrors, fmt.Sprintf("Wallet %s is not valid: %s", wallet.Address, result.Message))
			return errors.New("wallet validation failed")
		}
	}

	return nil
}

// validateProviderForDealCreation performs storage provider validation for deal creation
func (s *AutoDealService) validateProviderForDealCreation(
	ctx context.Context,
	db *gorm.DB,
	lotusClient jsonrpc.RPCClient,
	preparation *model.Preparation,
	validationErrors *[]string,
) error {
	if preparation.DealConfig.DealProvider == "" {
		// Try to get a default provider
		defaultSP, err := s.spValidator.GetDefaultStorageProvider(ctx, db, "auto-deal-creation")
		if err != nil {
			*validationErrors = append(*validationErrors, "No provider specified and no default available")
			return err
		}
		// Update preparation with default provider for deal creation
		preparation.DealConfig.DealProvider = defaultSP.ProviderID

		s.logInfo(ctx, db, "Using Default Provider",
			"No provider specified, using default "+defaultSP.ProviderID,
			model.ConfigMap{
				"preparation_name": preparation.Name,
				"provider_id":      defaultSP.ProviderID,
			})
	}

	// Validate the provider (this will use the default if we just set it)
	result, err := s.spValidator.ValidateStorageProvider(ctx, db, lotusClient, preparation.DealConfig.DealProvider, strconv.FormatUint(uint64(preparation.ID), 10))
	if err != nil {
		*validationErrors = append(*validationErrors, fmt.Sprintf("Provider validation error: %v", err))
		return err
	}

	if !result.IsValid {
		*validationErrors = append(*validationErrors, fmt.Sprintf("Provider %s is not valid: %s", preparation.DealConfig.DealProvider, result.Message))
		return errors.New("provider validation failed")
	}

	return nil
}

// Helper methods for logging
func (s *AutoDealService) logError(ctx context.Context, db *gorm.DB, title, message string, metadata model.ConfigMap) {
	_, err := s.notificationHandler.LogError(ctx, db, "auto-deal-service", title, message, metadata)
	if err != nil {
		autoDealLogger.Errorf("Failed to log error notification: %v", err)
	}
}

func (s *AutoDealService) logWarning(ctx context.Context, db *gorm.DB, title, message string, metadata model.ConfigMap) {
	_, err := s.notificationHandler.LogWarning(ctx, db, "auto-deal-service", title, message, metadata)
	if err != nil {
		autoDealLogger.Errorf("Failed to log warning notification: %v", err)
	}
}

func (s *AutoDealService) logInfo(ctx context.Context, db *gorm.DB, title, message string, metadata model.ConfigMap) {
	_, err := s.notificationHandler.LogInfo(ctx, db, "auto-deal-service", title, message, metadata)
	if err != nil {
		autoDealLogger.Errorf("Failed to log info notification: %v", err)
	}
}
