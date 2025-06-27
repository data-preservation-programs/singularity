package dataprep

import (
	"context"
	"fmt"
	"strconv"
	"time"

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

const (
	// DefaultTransactionTimeout defines the default timeout for database transactions
	DefaultTransactionTimeout = 30 * time.Second
	// DefaultQueryTimeout defines the default timeout for database queries
	DefaultQueryTimeout = 10 * time.Second
)

type AutoDealService struct {
	notificationHandler *notification.Handler
	scheduleHandler     schedule.Handler
	walletValidator     *wallet.BalanceValidator
	spValidator         *storage.SPValidator
}

func NewAutoDealService() *AutoDealService {
	service := &AutoDealService{
		notificationHandler: notification.Default,
		scheduleHandler:     schedule.Default,
		walletValidator:     wallet.DefaultBalanceValidator,
		spValidator:         storage.DefaultSPValidator,
	}

	autoDealLogger.Info("Auto-deal service initialized")
	return service
}

var DefaultAutoDealService = NewAutoDealService()

// RecoverFailedAutoDeal attempts to recover and retry failed auto-deal creation
func (s *AutoDealService) RecoverFailedAutoDeal(
	ctx context.Context,
	db *gorm.DB,
	lotusClient jsonrpc.RPCClient,
	preparationID string,
) error {
	autoDealLogger.Infof("Attempting to recover failed auto-deal for preparation ID: %s", preparationID)

	return db.Transaction(func(tx *gorm.DB) error {
		// First check if a schedule already exists
		var existingScheduleCount int64
		err := tx.WithContext(ctx).Model(&model.Schedule{}).
			Where("preparation_id = ?", preparationID).
			Set("gorm:query_option", "FOR UPDATE").
			Count(&existingScheduleCount).Error
		if err != nil {
			return errors.Wrap(err, "failed to check existing schedules")
		}

		if existingScheduleCount > 0 {
			autoDealLogger.Infof("Preparation %s already has %d schedule(s), no recovery needed", preparationID, existingScheduleCount)
			return nil
		}

		// Attempt to create the schedule
		_, err = s.CreateAutomaticDealSchedule(ctx, tx, lotusClient, preparationID)
		if err != nil {
			autoDealLogger.Errorf("Failed to recover auto-deal for preparation %s: %v", preparationID, err)
			return errors.Wrap(err, "failed to create auto-deal schedule during recovery")
		}

		autoDealLogger.Infof("Successfully recovered auto-deal for preparation %s", preparationID)
		return nil
	})
}

// CreateAutomaticDealSchedule creates deal schedules automatically for preparations with auto-deal enabled
func (s *AutoDealService) CreateAutomaticDealSchedule(
	ctx context.Context,
	db *gorm.DB,
	lotusClient jsonrpc.RPCClient,
	preparationID string,
) (*model.Schedule, error) {
	autoDealLogger.Infof("Starting automatic deal schedule creation for preparation ID: %s", preparationID)

	// Get preparation with auto-deal settings
	var preparation model.Preparation
	err := preparation.FindByIDOrName(db.WithContext(ctx), preparationID, "Wallets")
	if errors.Is(err, gorm.ErrRecordNotFound) {
		autoDealLogger.Errorf("Preparation not found: %s", preparationID)
		return nil, errors.Wrapf(err, "preparation %s not found", preparationID)
	}
	if err != nil {
		autoDealLogger.Errorf("Failed to fetch preparation %s: %v", preparationID, err)
		return nil, errors.Wrap(err, "failed to fetch preparation")
	}

	// Check if auto-deal creation is enabled
	if !preparation.DealConfig.AutoCreateDeals {
		autoDealLogger.Debugf("Auto-deal creation not enabled for preparation %s (ID: %s)", preparation.Name, preparationID)
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

	autoDealLogger.Infof("Starting validation for preparation %s (wallet_validation=%t, sp_validation=%t)",
		preparation.Name, preparation.WalletValidation, preparation.SPValidation)

	if preparation.WalletValidation {
		autoDealLogger.Debug("Performing wallet validation")
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
		autoDealLogger.Debug("Performing storage provider validation")
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
		autoDealLogger.Errorf("Validation failed for preparation %s with %d errors: %v",
			preparation.Name, len(validationErrors), validationErrors)

		s.logError(ctx, db, "Auto-Deal Creation Failed",
			"Auto-deal creation failed due to validation errors",
			model.ConfigMap{
				"preparation_name":  preparation.Name,
				"validation_errors": fmt.Sprintf("%v", validationErrors),
				"error_count":       strconv.Itoa(len(validationErrors)),
			})
		return nil, errors.Errorf("auto-deal creation failed with %d validation errors", len(validationErrors))
	}

	autoDealLogger.Info("All validations passed successfully")

	// Create the deal schedule using collected parameters
	dealRequest := s.buildDealScheduleRequest(&preparation)

	autoDealLogger.Infof("Building deal schedule request for preparation %s with provider %s", preparation.Name, dealRequest.Provider)
	s.logInfo(ctx, db, "Creating Deal Schedule",
		"Creating deal schedule with provider "+dealRequest.Provider,
		model.ConfigMap{
			"preparation_name": preparation.Name,
			"provider":         dealRequest.Provider,
			"verified":         strconv.FormatBool(dealRequest.Verified),
			"price_per_gb":     fmt.Sprintf("%.6f", dealRequest.PricePerGB),
		})

	// Create deal schedule within a transaction
	var dealSchedule *model.Schedule
	err = db.Transaction(func(tx *gorm.DB) error {
		autoDealLogger.Debugf("Creating deal schedule within transaction for preparation %s", preparation.Name)

		schedule, txErr := s.scheduleHandler.CreateHandler(ctx, tx, lotusClient, *dealRequest)
		if txErr != nil {
			autoDealLogger.Errorf("Failed to create deal schedule for preparation %s: %v", preparation.Name, txErr)
			return errors.Wrap(txErr, "failed to create deal schedule")
		}

		dealSchedule = schedule
		autoDealLogger.Infof("Successfully created deal schedule %d within transaction for preparation %s", schedule.ID, preparation.Name)
		return nil
	})

	if err != nil {
		s.logError(ctx, db, "Deal Schedule Creation Failed",
			"Failed to create automatic deal schedule",
			model.ConfigMap{
				"preparation_name": preparation.Name,
				"error":            err.Error(),
			})
		return nil, errors.Wrap(err, "transaction failed for deal schedule creation")
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
	autoDealLogger.Debugf("Checking readiness for preparation ID: %s", preparationID)

	// Check if all jobs for the preparation are complete with timeout
	queryCtx, cancel := context.WithTimeout(ctx, DefaultQueryTimeout)
	defer cancel()

	var incompleteJobCount int64
	err := db.WithContext(queryCtx).Model(&model.Job{}).
		Joins("JOIN source_attachments ON jobs.attachment_id = source_attachments.id").
		Where("source_attachments.preparation_id = ? AND jobs.state != ?", preparationID, model.Complete).
		Count(&incompleteJobCount).Error
	if err != nil {
		autoDealLogger.Errorf("Failed to count incomplete jobs for preparation %s: %v", preparationID, err)
		return false, errors.Wrap(err, "failed to count incomplete jobs")
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
	autoDealLogger.Info("Starting to process preparations ready for auto-deal creation")

	// Find preparations with auto-deal enabled that don't have schedules yet with timeout
	queryCtx, cancel := context.WithTimeout(ctx, DefaultQueryTimeout)
	defer cancel()

	var preparations []model.Preparation
	err := db.WithContext(queryCtx).Preload("Wallets").
		Where("auto_create_deals = ?", true).
		Find(&preparations).Error
	if err != nil {
		autoDealLogger.Errorf("Failed to fetch preparations with auto-deal enabled: %v", err)
		return errors.Wrap(err, "failed to fetch preparations with auto-deal enabled")
	}

	s.logInfo(ctx, db, "Processing Ready Preparations",
		fmt.Sprintf("Found %d preparations with auto-deal enabled", len(preparations)),
		model.ConfigMap{
			"preparation_count": strconv.Itoa(len(preparations)),
		})

	processedCount := 0
	errorCount := 0

	for _, prep := range preparations {
		prepIDStr := fmt.Sprintf("%d", prep.ID)
		autoDealLogger.Debugf("Processing preparation %s (ID: %s)", prep.Name, prepIDStr)

		// Use a transaction for each preparation processing with timeout
		txCtx, cancel := context.WithTimeout(ctx, DefaultTransactionTimeout)
		defer cancel()

		err := db.Transaction(func(tx *gorm.DB) error {
			// Check if preparation already has a deal schedule
			var existingScheduleCount int64
			err := tx.WithContext(txCtx).Model(&model.Schedule{}).
				Where("preparation_id = ?", prep.ID).
				Set("gorm:query_option", "FOR UPDATE"). // Lock for update
				Count(&existingScheduleCount).Error
			if err != nil {
				autoDealLogger.Errorf("Failed to check existing schedules for preparation %s: %v", prep.Name, err)
				return errors.Wrap(err, "failed to check existing schedules")
			}

			if existingScheduleCount > 0 {
				autoDealLogger.Debugf("Preparation %s already has %d schedule(s), skipping", prep.Name, existingScheduleCount)
				return nil // Not an error, just skip
			}

			// Check if preparation is ready
			isReady, err := s.CheckPreparationReadiness(txCtx, tx, prepIDStr)
			if err != nil {
				autoDealLogger.Errorf("Failed to check readiness for preparation %s: %v", prep.Name, err)
				return errors.Wrap(err, "failed to check preparation readiness")
			}

			if !isReady {
				autoDealLogger.Debugf("Preparation %s is not ready for deal creation yet", prep.Name)
				return nil // Not an error, just not ready
			}

			// Create automatic deal schedule
			_, err = s.CreateAutomaticDealSchedule(txCtx, tx, lotusClient, prepIDStr)
			if err != nil {
				autoDealLogger.Errorf("Failed to create auto-deal schedule for preparation %s: %v", prep.Name, err)
				return errors.Wrap(err, "failed to create auto-deal schedule")
			}

			processedCount++
			autoDealLogger.Infof("Successfully processed preparation %s for auto-deal creation", prep.Name)
			return nil
		})

		if err != nil {
			errorCount++
			autoDealLogger.Errorf("Transaction failed for preparation %s: %v", prep.Name, err)
			continue
		}
	}

	autoDealLogger.Infof("Auto-deal processing complete: %d processed, %d errors out of %d total preparations",
		processedCount, errorCount, len(preparations))

	s.logInfo(ctx, db, "Auto-Deal Processing Complete",
		fmt.Sprintf("Processed %d preparations, %d errors", processedCount, errorCount),
		model.ConfigMap{
			"processed_count": strconv.Itoa(processedCount),
			"error_count":     strconv.Itoa(errorCount),
			"total_count":     strconv.Itoa(len(preparations)),
		})

	if errorCount > 0 {
		return errors.Errorf("auto-deal processing completed with %d errors", errorCount)
	}
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

	autoDealLogger.Debugf("Built deal schedule request for preparation %s: provider=%s, verified=%t, price_per_gb=%f",
		preparation.Name, request.Provider, request.Verified, request.PricePerGB)

	// Convert epoch durations to time-based strings
	if preparation.DealConfig.DealStartDelay > 0 {
		// Convert epochs to duration (1 epoch = 30 seconds)
		epochDuration := preparation.DealConfig.DealStartDelay * 30 * time.Second
		request.StartDelay = epochDuration.String()
	} else {
		request.StartDelay = "72h" // Default
	}

	if preparation.DealConfig.DealDuration > 0 {
		// Convert epochs to duration (1 epoch = 30 seconds)
		epochDuration := preparation.DealConfig.DealDuration * 30 * time.Second
		request.Duration = epochDuration.String()
	} else {
		request.Duration = "12840h" // Default (~535 days)
	}

	// If no provider specified, leave empty - the schedule handler will validate and potentially use default
	if request.Provider == "" {
		// The schedule creation will fail if no provider, but we've already validated this in preparation creation
		autoDealLogger.Warnf("No provider specified for preparation %s, deal creation may fail", preparation.Name)
	} else {
		autoDealLogger.Debugf("Using provider %s for preparation %s", request.Provider, preparation.Name)
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
	autoDealLogger.Debugf("Validating wallets for preparation %s", preparation.Name)

	if len(preparation.Wallets) == 0 {
		autoDealLogger.Warnf("No wallets assigned to preparation %s", preparation.Name)
		*validationErrors = append(*validationErrors, "No wallets assigned to preparation")
		return errors.New("no wallets assigned to preparation")
	}

	// For now, just validate that wallets exist and are accessible
	// In a full implementation, you would calculate required balance based on data size
	for _, wallet := range preparation.Wallets {
		autoDealLogger.Debugf("Validating wallet %s for preparation %s", wallet.Address, preparation.Name)

		result, err := s.walletValidator.ValidateWalletExists(ctx, db, lotusClient, wallet.Address, strconv.FormatUint(uint64(preparation.ID), 10))
		if err != nil {
			errorMsg := fmt.Sprintf("Wallet validation error for %s: %v", wallet.Address, err)
			autoDealLogger.Error(errorMsg)
			*validationErrors = append(*validationErrors, errorMsg)
			return errors.Wrapf(err, "failed to validate wallet %s", wallet.Address)
		}
		if !result.IsValid {
			errorMsg := fmt.Sprintf("Wallet %s is not valid: %s", wallet.Address, result.Message)
			autoDealLogger.Warn(errorMsg)
			*validationErrors = append(*validationErrors, errorMsg)
			return errors.Errorf("wallet %s validation failed: %s", wallet.Address, result.Message)
		}

		autoDealLogger.Debugf("Wallet %s validated successfully", wallet.Address)
	}

	autoDealLogger.Infof("All %d wallets validated successfully for preparation %s", len(preparation.Wallets), preparation.Name)
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
	autoDealLogger.Debugf("Validating storage provider for preparation %s", preparation.Name)

	if preparation.DealConfig.DealProvider == "" {
		autoDealLogger.Warnf("No provider specified for preparation %s, attempting to use default", preparation.Name)

		// Try to get a default provider
		defaultSP, err := s.spValidator.GetDefaultStorageProvider(ctx, db, "auto-deal-creation")
		if err != nil {
			errorMsg := "No provider specified and no default available"
			autoDealLogger.Error(errorMsg)
			*validationErrors = append(*validationErrors, errorMsg)
			return errors.Wrap(err, "failed to get default storage provider")
		}
		// Update preparation with default provider for deal creation
		preparation.DealConfig.DealProvider = defaultSP.ProviderID

		autoDealLogger.Infof("Using default provider %s for preparation %s", defaultSP.ProviderID, preparation.Name)
		s.logInfo(ctx, db, "Using Default Provider",
			"No provider specified, using default "+defaultSP.ProviderID,
			model.ConfigMap{
				"preparation_name": preparation.Name,
				"provider_id":      defaultSP.ProviderID,
			})
	}

	// Validate the provider (this will use the default if we just set it)
	autoDealLogger.Debugf("Validating provider %s for preparation %s", preparation.DealConfig.DealProvider, preparation.Name)

	result, err := s.spValidator.ValidateStorageProvider(ctx, db, lotusClient, preparation.DealConfig.DealProvider, strconv.FormatUint(uint64(preparation.ID), 10))
	if err != nil {
		errorMsg := fmt.Sprintf("Provider validation error: %v", err)
		autoDealLogger.Error(errorMsg)
		*validationErrors = append(*validationErrors, errorMsg)
		return errors.Wrapf(err, "failed to validate storage provider %s", preparation.DealConfig.DealProvider)
	}

	if !result.IsValid {
		errorMsg := fmt.Sprintf("Provider %s is not valid: %s", preparation.DealConfig.DealProvider, result.Message)
		autoDealLogger.Warn(errorMsg)
		*validationErrors = append(*validationErrors, errorMsg)
		return errors.Errorf("provider %s validation failed: %s", preparation.DealConfig.DealProvider, result.Message)
	}

	autoDealLogger.Infof("Provider %s validated successfully for preparation %s", preparation.DealConfig.DealProvider, preparation.Name)
	return nil
}

// GetAutoDealStatus returns the status of auto-deal creation for a preparation
func (s *AutoDealService) GetAutoDealStatus(
	ctx context.Context,
	db *gorm.DB,
	preparationID string,
) (map[string]interface{}, error) {
	autoDealLogger.Debugf("Getting auto-deal status for preparation ID: %s", preparationID)

	var preparation model.Preparation
	err := preparation.FindByIDOrName(db.WithContext(ctx), preparationID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find preparation")
	}

	// Check if preparation is ready
	isReady, err := s.CheckPreparationReadiness(ctx, db, preparationID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to check preparation readiness")
	}

	// Check if schedule exists
	var scheduleCount int64
	err = db.WithContext(ctx).Model(&model.Schedule{}).
		Where("preparation_id = ?", preparation.ID).
		Count(&scheduleCount).Error
	if err != nil {
		return nil, errors.Wrap(err, "failed to count schedules")
	}

	status := map[string]interface{}{
		"preparation_id":    preparation.ID,
		"preparation_name":  preparation.Name,
		"auto_deal_enabled": preparation.DealConfig.AutoCreateDeals,
		"is_ready":          isReady,
		"has_schedule":      scheduleCount > 0,
		"schedule_count":    scheduleCount,
		"wallet_validation": preparation.WalletValidation,
		"sp_validation":     preparation.SPValidation,
	}

	autoDealLogger.Infof("Auto-deal status for %s: enabled=%t, ready=%t, has_schedule=%t",
		preparation.Name, preparation.DealConfig.AutoCreateDeals, isReady, scheduleCount > 0)

	return status, nil
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
