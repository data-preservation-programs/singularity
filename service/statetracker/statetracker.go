package statetracker

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/ipfs/go-log/v2"
	"gorm.io/gorm"
)

var Logger = log.Logger("statetracker")

// Error categories for better error handling and analytics
const (
	ErrorCategoryNetwork   = "network"     // Network connectivity issues
	ErrorCategoryProvider  = "provider"    // Storage provider related issues
	ErrorCategoryClient    = "client"      // Client related issues
	ErrorCategoryChain     = "chain"       // Blockchain/consensus issues
	ErrorCategoryDB        = "database"    // Database operation issues
	ErrorCategoryInternal  = "internal"    // Internal processing errors
	ErrorCategoryTimeout   = "timeout"     // Timeout related errors
	ErrorCategoryFunding   = "funding"     // Insufficient funds or collateral
	ErrorCategorySlashing  = "slashing"    // Deal slashing related
	ErrorCategoryExpiry    = "expiry"      // Deal expiration related
)

// StateChangeTracker is responsible for tracking deal state changes
type StateChangeTracker struct {
	db *gorm.DB
}

// StateChangeMetadata represents additional metadata that can be stored with a state change
type StateChangeMetadata struct {
	Reason           string            `json:"reason,omitempty"`           // Reason for the state change
	Error            string            `json:"error,omitempty"`            // Error message if applicable
	ErrorCategory    string            `json:"errorCategory,omitempty"`    // Category of error (network, provider, client, etc.)
	TransactionID    string            `json:"transactionId,omitempty"`    // On-chain transaction ID
	PublishCID       string            `json:"publishCid,omitempty"`       // Message CID for deal publication
	ActivationEpoch  *int32            `json:"activationEpoch,omitempty"`  // Epoch when deal was activated
	ExpirationEpoch  *int32            `json:"expirationEpoch,omitempty"`  // Epoch when deal expires
	SlashingEpoch    *int32            `json:"slashingEpoch,omitempty"`    // Epoch when deal was slashed
	StoragePrice     string            `json:"storagePrice,omitempty"`     // Storage price per epoch
	PieceSize        int64             `json:"pieceSize,omitempty"`        // Size of the piece
	VerifiedDeal     bool              `json:"verifiedDeal,omitempty"`     // Whether this is a verified deal
	RetryCount       int               `json:"retryCount,omitempty"`       // Number of retry attempts
	ProcessingTime   int64             `json:"processingTime,omitempty"`   // Time taken to process in milliseconds
	ChainTipSetKey   string            `json:"chainTipSetKey,omitempty"`   // Chain tipset key when event occurred
	AdditionalFields map[string]string `json:"additionalFields,omitempty"` // Any additional custom fields
}

// NewStateChangeTracker creates a new instance of StateChangeTracker
func NewStateChangeTracker(db *gorm.DB) *StateChangeTracker {
	return &StateChangeTracker{
		db: db,
	}
}

// TrackStateChange records a deal state change with comprehensive metadata
func (t *StateChangeTracker) TrackStateChange(ctx context.Context, deal *model.Deal, previousState *model.DealState, newState model.DealState, metadata *StateChangeMetadata) error {
	return t.TrackStateChangeWithDetails(ctx, deal.ID, previousState, newState, nil, nil, deal.Provider, deal.ClientActorID, metadata)
}

// TrackStateChangeWithDetails records a state change with all available details
func (t *StateChangeTracker) TrackStateChangeWithDetails(
	ctx context.Context,
	dealID model.DealID,
	previousState *model.DealState,
	newState model.DealState,
	epochHeight *int32,
	sectorID *string,
	providerID string,
	clientAddress string,
	metadata *StateChangeMetadata,
) error {
	// Serialize metadata to JSON
	var metadataJSON string
	if metadata != nil {
		metadataBytes, err := json.Marshal(metadata)
		if err != nil {
			Logger.Warnw("Failed to serialize metadata", "dealId", dealID, "error", err)
			metadataJSON = "{}"
		} else {
			metadataJSON = string(metadataBytes)
		}
	} else {
		metadataJSON = "{}"
	}

	// Log the state change
	if previousState != nil {
		Logger.Infow("Deal state change tracked",
			"dealId", dealID,
			"previousState", *previousState,
			"newState", newState,
			"provider", providerID,
			"client", clientAddress,
			"epochHeight", epochHeight,
			"sectorId", sectorID,
		)
	} else {
		Logger.Infow("Initial deal state tracked",
			"dealId", dealID,
			"newState", newState,
			"provider", providerID,
			"client", clientAddress,
			"epochHeight", epochHeight,
			"sectorId", sectorID,
		)
	}

	// Record the state change using database retry mechanism
	return database.DoRetry(ctx, func() error {
		return model.CreateDealStateChange(
			t.db.WithContext(ctx),
			dealID,
			previousState,
			newState,
			epochHeight,
			sectorID,
			providerID,
			clientAddress,
			metadataJSON,
		)
	})
}

// RecoverMissingStateChanges attempts to recover state changes that may have been missed
// during service restarts by comparing current deal states with the last recorded state changes
func (t *StateChangeTracker) RecoverMissingStateChanges(ctx context.Context) error {
	Logger.Info("Starting recovery of missing state changes")

	db := t.db.WithContext(ctx)

	// Get all deals that don't have any state change records
	var dealsWithoutChanges []model.Deal
	err := db.Raw(`
		SELECT d.* FROM deals d 
		LEFT JOIN deal_state_changes dsc ON d.id = dsc.deal_id 
		WHERE dsc.deal_id IS NULL
	`).Scan(&dealsWithoutChanges).Error
	if err != nil {
		return errors.Wrap(err, "failed to find deals without state changes")
	}

	Logger.Infow("Found deals without state change records", "count", len(dealsWithoutChanges))

	// Create initial state change records for deals without any
	for _, deal := range dealsWithoutChanges {
		metadata := &StateChangeMetadata{
			Reason: "Recovery - initial state recorded during system restart",
		}

		err = t.TrackStateChangeWithDetails(
			ctx,
			deal.ID,
			nil, // No previous state for initial record
			deal.State,
			nil, // No epoch info available for recovery
			nil, // No sector ID available for recovery
			deal.Provider,
			deal.ClientActorID,
			metadata,
		)
		if err != nil {
			Logger.Errorw("Failed to create recovery state change", "dealId", deal.ID, "error", err)
			// Continue with other deals even if one fails
			continue
		}
	}

	// Get deals where the current state doesn't match the latest recorded state change
	var inconsistentDeals []struct {
		DealID            model.DealID    `json:"dealId"`
		CurrentState      model.DealState `json:"currentState"`
		LastRecordedState model.DealState `json:"lastRecordedState"`
		Provider          string          `json:"provider"`
		ClientActorID     string          `json:"clientActorId"`
	}

	err = db.Raw(`
		SELECT 
			d.id as deal_id,
			d.state as current_state,
			dsc.new_state as last_recorded_state,
			d.provider,
			d.client_actor_id
		FROM deals d
		INNER JOIN (
			SELECT deal_id, new_state, 
				ROW_NUMBER() OVER (PARTITION BY deal_id ORDER BY timestamp DESC) as rn
			FROM deal_state_changes
		) dsc ON d.id = dsc.deal_id AND dsc.rn = 1
		WHERE d.state != dsc.new_state
	`).Scan(&inconsistentDeals).Error
	if err != nil {
		return errors.Wrap(err, "failed to find deals with inconsistent states")
	}

	Logger.Infow("Found deals with inconsistent state records", "count", len(inconsistentDeals))

	// Create state change records for inconsistent deals
	for _, deal := range inconsistentDeals {
		metadata := &StateChangeMetadata{
			Reason: fmt.Sprintf("Recovery - state changed from %s to %s during system downtime",
				deal.LastRecordedState, deal.CurrentState),
		}

		err = t.TrackStateChangeWithDetails(
			ctx,
			deal.DealID,
			&deal.LastRecordedState,
			deal.CurrentState,
			nil, // No epoch info available for recovery
			nil, // No sector ID available for recovery
			deal.Provider,
			deal.ClientActorID,
			metadata,
		)
		if err != nil {
			Logger.Errorw("Failed to create recovery state change for inconsistent deal",
				"dealId", deal.DealID, "error", err)
			// Continue with other deals even if one fails
			continue
		}
	}

	Logger.Infow("Completed recovery of missing state changes",
		"newRecords", len(dealsWithoutChanges),
		"inconsistentFixed", len(inconsistentDeals))

	return nil
}

// GetStateChangesForDeal retrieves all state changes for a specific deal
func (t *StateChangeTracker) GetStateChangesForDeal(ctx context.Context, dealID model.DealID) ([]model.DealStateChange, error) {
	return model.GetDealStateChangesByDealID(t.db.WithContext(ctx), dealID)
}

// GetStateChanges retrieves state changes with filtering and pagination
func (t *StateChangeTracker) GetStateChanges(ctx context.Context, query model.DealStateChangeQuery) ([]model.DealStateChange, int64, error) {
	return model.GetDealStateChanges(t.db.WithContext(ctx), query)
}

// GetStateChangeStats returns statistics about state changes
func (t *StateChangeTracker) GetStateChangeStats(ctx context.Context) (map[string]interface{}, error) {
	db := t.db.WithContext(ctx)
	stats := make(map[string]interface{})

	// Total state changes
	var totalChanges int64
	err := db.Model(&model.DealStateChange{}).Count(&totalChanges).Error
	if err != nil {
		return nil, errors.Wrap(err, "failed to count total state changes")
	}
	stats["totalStateChanges"] = totalChanges

	// State changes by new state
	var stateDistribution []struct {
		State string `json:"state"`
		Count int64  `json:"count"`
	}
	err = db.Model(&model.DealStateChange{}).
		Select("new_state as state, COUNT(*) as count").
		Group("new_state").
		Scan(&stateDistribution).Error
	if err != nil {
		return nil, errors.Wrap(err, "failed to get state distribution")
	}
	stats["stateDistribution"] = stateDistribution

	// State changes in the last 24 hours
	var recentChanges int64
	since := time.Now().Add(-24 * time.Hour)
	err = db.Model(&model.DealStateChange{}).
		Where("timestamp >= ?", since).
		Count(&recentChanges).Error
	if err != nil {
		return nil, errors.Wrap(err, "failed to count recent state changes")
	}
	stats["recentStateChanges24h"] = recentChanges

	// Most active providers by state changes
	var activeProviders []struct {
		ProviderID string `json:"providerId"`
		Count      int64  `json:"count"`
	}
	err = db.Model(&model.DealStateChange{}).
		Select("provider_id, COUNT(*) as count").
		Group("provider_id").
		Order("count DESC").
		Limit(10).
		Scan(&activeProviders).Error
	if err != nil {
		return nil, errors.Wrap(err, "failed to get active providers")
	}
	stats["topProvidersByStateChanges"] = activeProviders

	return stats, nil
}

// CategorizeError determines the error category based on the error content
func CategorizeError(err error) string {
	if err == nil {
		return ""
	}

	errStr := strings.ToLower(err.Error())
	
	switch {
	case strings.Contains(errStr, "database") || strings.Contains(errStr, "sql") || strings.Contains(errStr, "gorm"):
		return ErrorCategoryDB
	case strings.Contains(errStr, "timeout") || strings.Contains(errStr, "deadline exceeded"):
		return ErrorCategoryTimeout
	case strings.Contains(errStr, "network") || strings.Contains(errStr, "connection") || strings.Contains(errStr, "tcp"):
		return ErrorCategoryNetwork
	case strings.Contains(errStr, "provider") || strings.Contains(errStr, "storage provider") || strings.Contains(errStr, "miner"):
		return ErrorCategoryProvider
	case strings.Contains(errStr, "client") || strings.Contains(errStr, "wallet"):
		return ErrorCategoryClient
	case strings.Contains(errStr, "chain") || strings.Contains(errStr, "consensus") || strings.Contains(errStr, "tipset"):
		return ErrorCategoryChain
	case strings.Contains(errStr, "fund") || strings.Contains(errStr, "balance") || strings.Contains(errStr, "collateral"):
		return ErrorCategoryFunding
	case strings.Contains(errStr, "slash"):
		return ErrorCategorySlashing
	case strings.Contains(errStr, "expir"):
		return ErrorCategoryExpiry
	default:
		return ErrorCategoryInternal
	}
}

// CreateEnhancedMetadata creates metadata with enhanced information for state changes
func CreateEnhancedMetadata(reason string, err error, dealInfo *DealInfo) *StateChangeMetadata {
	metadata := &StateChangeMetadata{
		Reason: reason,
	}

	if err != nil {
		metadata.Error = err.Error()
		metadata.ErrorCategory = CategorizeError(err)
	}

	if dealInfo != nil {
		metadata.StoragePrice = dealInfo.StoragePrice
		metadata.PieceSize = dealInfo.PieceSize
		metadata.VerifiedDeal = dealInfo.VerifiedDeal
		if dealInfo.ActivationEpoch != nil {
			metadata.ActivationEpoch = dealInfo.ActivationEpoch
		}
		if dealInfo.ExpirationEpoch != nil {
			metadata.ExpirationEpoch = dealInfo.ExpirationEpoch
		}
		if dealInfo.SlashingEpoch != nil {
			metadata.SlashingEpoch = dealInfo.SlashingEpoch
		}
		if dealInfo.PublishCID != "" {
			metadata.PublishCID = dealInfo.PublishCID
		}
		if dealInfo.TransactionID != "" {
			metadata.TransactionID = dealInfo.TransactionID
		}
		if dealInfo.ChainTipSetKey != "" {
			metadata.ChainTipSetKey = dealInfo.ChainTipSetKey
		}
	}

	return metadata
}

// DealInfo represents additional deal information for enhanced metadata
type DealInfo struct {
	StoragePrice     string
	PieceSize        int64
	VerifiedDeal     bool
	ActivationEpoch  *int32
	ExpirationEpoch  *int32
	SlashingEpoch    *int32
	PublishCID       string
	TransactionID    string
	ChainTipSetKey   string
}

// TrackStateChangeWithError tracks a state change that resulted from an error
func (t *StateChangeTracker) TrackStateChangeWithError(
	ctx context.Context,
	dealID model.DealID,
	previousState *model.DealState,
	newState model.DealState,
	epochHeight *int32,
	sectorID *string,
	providerID string,
	clientAddress string,
	err error,
	dealInfo *DealInfo,
) error {
	reason := "Deal state change"
	if err != nil {
		reason = fmt.Sprintf("Deal state change due to error: %s", CategorizeError(err))
	}

	metadata := CreateEnhancedMetadata(reason, err, dealInfo)
	
	return t.TrackStateChangeWithDetails(
		ctx,
		dealID,
		previousState,
		newState,
		epochHeight,
		sectorID,
		providerID,
		clientAddress,
		metadata,
	)
}
