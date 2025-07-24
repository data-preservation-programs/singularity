package statetracker

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/errorcategorization"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestTrackStateChange(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		tracker := NewStateChangeTracker(db)

		// Create a test deal
		deal := &model.Deal{
			State:         model.DealProposed,
			Provider:      "f01234",
			ClientActorID: "f01000",
			PieceCID:      model.CID{},
			PieceSize:     1024,
			StartEpoch:    100,
			EndEpoch:      200,
			Price:         "1000",
			Verified:      false,
		}
		err := db.Create(deal).Error
		require.NoError(t, err)

		// Test tracking initial state change
		metadata := &StateChangeMetadata{
			Reason:       "Initial deal proposal",
			StoragePrice: "1000",
		}

		err = tracker.TrackStateChange(ctx, deal, nil, model.DealProposed, metadata)
		require.NoError(t, err)

		// Verify the state change was recorded
		var stateChanges []model.DealStateChange
		err = db.Where("deal_id = ?", deal.ID).Find(&stateChanges).Error
		require.NoError(t, err)
		require.Len(t, stateChanges, 1)

		sc := stateChanges[0]
		require.Equal(t, deal.ID, sc.DealID)
		require.Equal(t, model.DealState(""), sc.PreviousState) // Empty for initial state
		require.Equal(t, model.DealProposed, sc.NewState)
		require.Equal(t, deal.Provider, sc.ProviderID)
		require.Equal(t, deal.ClientActorID, sc.ClientAddress)

		// Verify metadata was serialized correctly
		var savedMetadata StateChangeMetadata
		err = json.Unmarshal([]byte(sc.Metadata), &savedMetadata)
		require.NoError(t, err)
		require.Equal(t, "Initial deal proposal", savedMetadata.Reason)
		require.Equal(t, "1000", savedMetadata.StoragePrice)
	})
}

func TestTrackStateChangeWithPreviousState(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		tracker := NewStateChangeTracker(db)

		// Create a test deal
		deal := &model.Deal{
			State:         model.DealActive,
			Provider:      "f01234",
			ClientActorID: "f01000",
			PieceCID:      model.CID{},
			PieceSize:     1024,
			StartEpoch:    100,
			EndEpoch:      200,
			Price:         "1000",
			Verified:      false,
		}
		err := db.Create(deal).Error
		require.NoError(t, err)

		// Track state change from proposed to active
		previousState := model.DealProposed
		metadata := &StateChangeMetadata{
			Reason:          "Deal activated",
			ActivationEpoch: int32Ptr(150),
		}

		err = tracker.TrackStateChange(ctx, deal, &previousState, model.DealActive, metadata)
		require.NoError(t, err)

		// Verify the state change was recorded
		var stateChanges []model.DealStateChange
		err = db.Where("deal_id = ?", deal.ID).Find(&stateChanges).Error
		require.NoError(t, err)
		require.Len(t, stateChanges, 1)

		sc := stateChanges[0]
		require.Equal(t, deal.ID, sc.DealID)
		require.Equal(t, model.DealProposed, sc.PreviousState)
		require.Equal(t, model.DealActive, sc.NewState)
		require.Equal(t, deal.Provider, sc.ProviderID)
		require.Equal(t, deal.ClientActorID, sc.ClientAddress)

		// Verify metadata
		var savedMetadata StateChangeMetadata
		err = json.Unmarshal([]byte(sc.Metadata), &savedMetadata)
		require.NoError(t, err)
		require.Equal(t, "Deal activated", savedMetadata.Reason)
		require.NotNil(t, savedMetadata.ActivationEpoch)
		require.Equal(t, int32(150), *savedMetadata.ActivationEpoch)
	})
}

func TestGetStateChangeStats(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		tracker := NewStateChangeTracker(db)

		// Create test deal
		deal := &model.Deal{
			State:         model.DealActive,
			Provider:      "f01234",
			ClientActorID: "f01000",
			PieceCID:      model.CID{},
			PieceSize:     1024,
			StartEpoch:    100,
			EndEpoch:      200,
			Price:         "1000",
			Verified:      false,
		}
		err := db.Create(deal).Error
		require.NoError(t, err)

		// Create state changes with different states and timestamps
		baseTime := time.Now()
		stateChanges := []model.DealStateChange{
			{
				DealID:        deal.ID,
				PreviousState: "",
				NewState:      model.DealProposed,
				Timestamp:     baseTime.Add(-2 * time.Hour),
				ProviderID:    "f01234",
				ClientAddress: "f01000",
				Metadata:      `{"reason": "Initial proposal"}`,
			},
			{
				DealID:        deal.ID,
				PreviousState: model.DealProposed,
				NewState:      model.DealActive,
				Timestamp:     baseTime.Add(-1 * time.Hour),
				ProviderID:    "f01234",
				ClientAddress: "f01000",
				Metadata:      `{"reason": "Deal activated"}`,
			},
		}

		for _, sc := range stateChanges {
			err = db.Create(&sc).Error
			require.NoError(t, err)
		}

		// Test stats retrieval
		stats, err := tracker.GetStateChangeStats(ctx)
		require.NoError(t, err)

		// Verify expected stats structure
		require.Contains(t, stats, "totalStateChanges")
		require.Contains(t, stats, "stateDistribution")
		require.Contains(t, stats, "recentStateChanges24h")
		require.Contains(t, stats, "topProvidersByStateChanges")

		// Verify values
		require.Equal(t, int64(2), stats["totalStateChanges"])
		require.Equal(t, int64(2), stats["recentStateChanges24h"])

		// Check top providers
		topProviders := stats["topProvidersByStateChanges"].([]struct {
			ProviderID string `json:"providerId"`
			Count      int64  `json:"count"`
		})
		require.Len(t, topProviders, 1)
		require.Equal(t, "f01234", topProviders[0].ProviderID)
		require.Equal(t, int64(2), topProviders[0].Count)
	})
}

func TestRecoverMissingStateChanges(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		tracker := NewStateChangeTracker(db)

		// Create a deal without any state change records
		deal := &model.Deal{
			State:         model.DealProposed,
			Provider:      "f01234",
			ClientActorID: "f01000",
			PieceCID:      model.CID{},
			PieceSize:     1024,
			StartEpoch:    100,
			EndEpoch:      200,
			Price:         "1000",
			Verified:      false,
		}
		err := db.Create(deal).Error
		require.NoError(t, err)

		// Verify no state changes exist initially
		var initialChanges []model.DealStateChange
		err = db.Where("deal_id = ?", deal.ID).Find(&initialChanges).Error
		require.NoError(t, err)
		require.Len(t, initialChanges, 0)

		// Run recovery
		err = tracker.RecoverMissingStateChanges(ctx)
		require.NoError(t, err)

		// Verify a recovery state change was created
		var recoveredChanges []model.DealStateChange
		err = db.Where("deal_id = ?", deal.ID).Find(&recoveredChanges).Error
		require.NoError(t, err)
		require.Len(t, recoveredChanges, 1)

		rc := recoveredChanges[0]
		require.Equal(t, deal.ID, rc.DealID)
		require.Equal(t, model.DealState(""), rc.PreviousState) // No previous state for initial record
		require.Equal(t, model.DealProposed, rc.NewState)
		require.Equal(t, deal.Provider, rc.ProviderID)
		require.Equal(t, deal.ClientActorID, rc.ClientAddress)

		// Verify recovery metadata
		var metadata StateChangeMetadata
		err = json.Unmarshal([]byte(rc.Metadata), &metadata)
		require.NoError(t, err)
		require.Contains(t, metadata.Reason, "Recovery")
	})
}

// Helper function to create int32 pointer
func int32Ptr(i int32) *int32 {
	return &i
}

// Integration tests for enhanced error categorization
func TestCreateErrorMetadata(t *testing.T) {
	// Test with full categorization
	categorization := &errorcategorization.ErrorCategorization{
		Category:    errorcategorization.NetworkTimeout,
		DealState:   model.DealErrored,
		Description: "Network timeout during deal negotiation",
		Severity:    errorcategorization.SeverityMedium,
		Retryable:   true,
		Metadata: &errorcategorization.ErrorMetadata{
			NetworkEndpoint: "tcp://192.168.1.100:8080",
			NetworkLatency:  func() *int64 { latency := int64(500); return &latency }(),
			ProviderID:      "f01234",
			ProviderVersion: "1.2.3",
			PieceCID:        "bafkqaaa",
			PieceSize:       func() *int64 { size := int64(1024); return &size }(),
			AttemptNumber:   func() *int { attempt := 2; return &attempt }(),
			ClientAddress:   "f3abc123",
			WalletBalance:   "1000",
			CustomFields: map[string]interface{}{
				"region": "us-west-1",
				"test":   true,
			},
		},
	}

	metadata := CreateErrorMetadata(categorization, "Test reason")
	require.NotNil(t, metadata)

	assert.Equal(t, "Test reason", metadata.Reason)
	assert.Equal(t, string(errorcategorization.NetworkTimeout), metadata.ErrorCategory)
	assert.Equal(t, string(errorcategorization.SeverityMedium), metadata.ErrorSeverity)
	assert.Equal(t, true, *metadata.ErrorRetryable)
	assert.Equal(t, "tcp://192.168.1.100:8080", metadata.NetworkEndpoint)
	assert.Equal(t, int64(500), *metadata.NetworkLatency)
	assert.Equal(t, "f01234", metadata.ProviderID)
	assert.Equal(t, "1.2.3", metadata.ProviderVersion)
	assert.Equal(t, "bafkqaaa", metadata.ProposalID)
	assert.Equal(t, 2, *metadata.AttemptNumber)
	assert.Equal(t, "f3abc123", metadata.ClientAddress)
	assert.Equal(t, "1000", metadata.WalletBalance)
	assert.Equal(t, "us-west-1", metadata.AdditionalFields["region"])
	assert.Equal(t, true, metadata.AdditionalFields["test"])
}

func TestCreateErrorMetadataWithNil(t *testing.T) {
	metadata := CreateErrorMetadata(nil, "Test reason")
	require.NotNil(t, metadata)
	assert.Equal(t, "Test reason", metadata.Reason)
	assert.Empty(t, metadata.ErrorCategory)
	assert.Empty(t, metadata.ErrorSeverity)
	assert.Nil(t, metadata.ErrorRetryable)
}

func TestTrackErrorStateChange(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		tracker := NewStateChangeTracker(db)

		// Create a test deal
		deal := &model.Deal{
			ID:            1,
			State:         model.DealProposed,
			Provider:      "f01234",
			ClientActorID: "f3abc123",
			PieceCID:      model.CID{},
			PieceSize:     1024,
			Price:         "100",
			Verified:      true,
		}

		err := db.Create(deal).Error
		require.NoError(t, err)

		// Test error state change tracking
		errorMessage := "connection timeout during negotiation"
		contextMetadata := &errorcategorization.ErrorMetadata{
			NetworkEndpoint: "tcp://192.168.1.100:8080",
			NetworkLatency:  func() *int64 { latency := int64(500); return &latency }(),
			ProviderID:      "f01234",
			ProviderVersion: "1.2.3",
			AttemptNumber:   func() *int { attempt := 1; return &attempt }(),
			LastAttemptTime: func() *time.Time { t := time.Now(); return &t }(),
			ClientAddress:   "f3abc123",
			WalletBalance:   "1000",
			CustomFields: map[string]interface{}{
				"retry_count": 1,
				"endpoint":    "boost",
			},
		}

		previousState := model.DealProposed
		err = tracker.TrackErrorStateChange(ctx, deal, &previousState, errorMessage, contextMetadata)
		require.NoError(t, err)

		// Verify the state change was recorded
		var stateChanges []model.DealStateChange
		err = db.Where("deal_id = ?", deal.ID).Find(&stateChanges).Error
		require.NoError(t, err)
		require.Len(t, stateChanges, 1)

		stateChange := stateChanges[0]
		assert.Equal(t, deal.ID, stateChange.DealID)
		assert.Equal(t, model.DealProposed, stateChange.PreviousState)
		assert.Equal(t, model.DealErrored, stateChange.NewState) // Should be categorized as errored
		assert.Equal(t, "f01234", stateChange.ProviderID)
		assert.Equal(t, "f3abc123", stateChange.ClientAddress)

		// Parse and verify metadata
		var metadata StateChangeMetadata
		err = json.Unmarshal([]byte(stateChange.Metadata), &metadata)
		require.NoError(t, err)

		assert.Equal(t, "Network timeout during deal negotiation", metadata.Reason)
		assert.Equal(t, errorMessage, metadata.Error)
		assert.Equal(t, string(errorcategorization.NetworkTimeout), metadata.ErrorCategory)
		assert.Equal(t, string(errorcategorization.SeverityMedium), metadata.ErrorSeverity)
		assert.Equal(t, true, *metadata.ErrorRetryable)
		assert.Equal(t, "tcp://192.168.1.100:8080", metadata.NetworkEndpoint)
		assert.Equal(t, int64(500), *metadata.NetworkLatency)
		assert.Equal(t, "f01234", metadata.ProviderID)
		assert.Equal(t, "1.2.3", metadata.ProviderVersion)
		assert.Equal(t, 1, *metadata.AttemptNumber)
		assert.Equal(t, "f3abc123", metadata.ClientAddress)
		assert.Equal(t, "1000", metadata.WalletBalance)
		assert.Equal(t, float64(1), metadata.AdditionalFields["retry_count"])
		assert.Equal(t, "boost", metadata.AdditionalFields["endpoint"])
	})
}

func TestTrackErrorStateChangeWithDifferentErrorTypes(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		tracker := NewStateChangeTracker(db)

		testCases := []struct {
			name              string
			errorMessage      string
			expectedCategory  errorcategorization.ErrorCategory
			expectedDealState model.DealState
			expectedSeverity  errorcategorization.ErrorSeverity
			expectedRetryable bool
		}{
			{
				name:              "network timeout",
				errorMessage:      "context deadline exceeded",
				expectedCategory:  errorcategorization.NetworkTimeout,
				expectedDealState: model.DealErrored,
				expectedSeverity:  errorcategorization.SeverityMedium,
				expectedRetryable: true,
			},
			{
				name:              "deal rejection",
				errorMessage:      "deal rejected by storage provider",
				expectedCategory:  errorcategorization.DealRejectedByProvider,
				expectedDealState: model.DealRejected,
				expectedSeverity:  errorcategorization.SeverityHigh,
				expectedRetryable: false,
			},
			{
				name:              "insufficient funds",
				errorMessage:      "insufficient funds in wallet",
				expectedCategory:  errorcategorization.ClientInsufficientFunds,
				expectedDealState: model.DealErrored,
				expectedSeverity:  errorcategorization.SeverityHigh,
				expectedRetryable: true,
			},
			{
				name:              "piece corrupted",
				errorMessage:      "piece data is corrupted",
				expectedCategory:  errorcategorization.PieceCorrupted,
				expectedDealState: model.DealErrored,
				expectedSeverity:  errorcategorization.SeverityCritical,
				expectedRetryable: false,
			},
		}

		for i, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				// Create a test deal for each case
				deal := &model.Deal{
					ID:            model.DealID(i + 10), // Use unique IDs
					State:         model.DealProposed,
					Provider:      "f01234",
					ClientActorID: "f3abc123",
					PieceCID:      model.CID{},
					PieceSize:     1024,
					Price:         "100",
					Verified:      true,
				}

				err := db.Create(deal).Error
				require.NoError(t, err)

				// Track error state change
				contextMetadata := &errorcategorization.ErrorMetadata{
					ProviderID:    "f01234",
					ClientAddress: "f3abc123",
					AttemptNumber: func() *int { attempt := 1; return &attempt }(),
				}

				previousState := model.DealProposed
				err = tracker.TrackErrorStateChange(ctx, deal, &previousState, tc.errorMessage, contextMetadata)
				require.NoError(t, err)

				// Verify the state change was recorded with correct categorization
				var stateChanges []model.DealStateChange
				err = db.Where("deal_id = ?", deal.ID).Find(&stateChanges).Error
				require.NoError(t, err)
				require.Len(t, stateChanges, 1)

				stateChange := stateChanges[0]
				assert.Equal(t, tc.expectedDealState, stateChange.NewState)

				// Parse and verify metadata contains correct categorization
				var metadata StateChangeMetadata
				err = json.Unmarshal([]byte(stateChange.Metadata), &metadata)
				require.NoError(t, err)

				assert.Equal(t, tc.errorMessage, metadata.Error)
				assert.Equal(t, string(tc.expectedCategory), metadata.ErrorCategory)
				assert.Equal(t, string(tc.expectedSeverity), metadata.ErrorSeverity)
				assert.Equal(t, tc.expectedRetryable, *metadata.ErrorRetryable)
			})
		}
	})
}

func TestTrackStateChangeWithEnhancedMetadata(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		tracker := NewStateChangeTracker(db)

		// Create a test deal
		deal := &model.Deal{
			ID:            1,
			State:         model.DealProposed,
			Provider:      "f01234",
			ClientActorID: "f3abc123",
			PieceCID:      model.CID{},
			PieceSize:     1024,
			Price:         "100",
			Verified:      true,
		}

		err := db.Create(deal).Error
		require.NoError(t, err)

		// Create enhanced metadata
		now := time.Now()
		metadata := &StateChangeMetadata{
			Reason:          "Deal successfully negotiated",
			TransactionID:   "bafy123456",
			PublishCID:      "bafy789012",
			StoragePrice:    "100",
			ErrorCategory:   "",
			ErrorSeverity:   "",
			NetworkEndpoint: "tcp://provider.com:8080",
			NetworkLatency:  func() *int64 { latency := int64(200); return &latency }(),
			ProviderVersion: "1.2.3",
			ProviderRegion:  "us-west-1",
			ProposalID:      "prop-123",
			AttemptNumber:   func() *int { attempt := 1; return &attempt }(),
			LastAttemptTime: &now,
			WalletBalance:   "5000",
			SystemLoad:      func() *float64 { load := 0.75; return &load }(),
			MemoryUsage:     func() *int64 { mem := int64(1024 * 1024 * 512); return &mem }(),         // 512MB
			DiskSpaceUsed:   func() *int64 { disk := int64(1024 * 1024 * 1024 * 10); return &disk }(), // 10GB
			DatabaseHealth:  "healthy",
			AdditionalFields: map[string]interface{}{
				"custom_field1": "value1",
				"custom_field2": 42,
				"custom_field3": true,
			},
		}

		// Track the state change
		previousState := model.DealProposed
		err = tracker.TrackStateChange(ctx, deal, &previousState, model.DealPublished, metadata)
		require.NoError(t, err)

		// Verify the state change was recorded
		var stateChanges []model.DealStateChange
		err = db.Where("deal_id = ?", deal.ID).Find(&stateChanges).Error
		require.NoError(t, err)
		require.Len(t, stateChanges, 1)

		stateChange := stateChanges[0]
		assert.Equal(t, deal.ID, stateChange.DealID)
		assert.Equal(t, model.DealProposed, stateChange.PreviousState)
		assert.Equal(t, model.DealPublished, stateChange.NewState)

		// Parse and verify all metadata fields
		var retrievedMetadata StateChangeMetadata
		err = json.Unmarshal([]byte(stateChange.Metadata), &retrievedMetadata)
		require.NoError(t, err)

		assert.Equal(t, "Deal successfully negotiated", retrievedMetadata.Reason)
		assert.Equal(t, "bafy123456", retrievedMetadata.TransactionID)
		assert.Equal(t, "bafy789012", retrievedMetadata.PublishCID)
		assert.Equal(t, "100", retrievedMetadata.StoragePrice)
		assert.Equal(t, "tcp://provider.com:8080", retrievedMetadata.NetworkEndpoint)
		assert.Equal(t, int64(200), *retrievedMetadata.NetworkLatency)
		assert.Equal(t, "1.2.3", retrievedMetadata.ProviderVersion)
		assert.Equal(t, "us-west-1", retrievedMetadata.ProviderRegion)
		assert.Equal(t, "prop-123", retrievedMetadata.ProposalID)
		assert.Equal(t, 1, *retrievedMetadata.AttemptNumber)
		assert.Equal(t, now.Unix(), retrievedMetadata.LastAttemptTime.Unix())
		assert.Equal(t, "5000", retrievedMetadata.WalletBalance)
		assert.Equal(t, 0.75, *retrievedMetadata.SystemLoad)
		assert.Equal(t, int64(1024*1024*512), *retrievedMetadata.MemoryUsage)
		assert.Equal(t, int64(1024*1024*1024*10), *retrievedMetadata.DiskSpaceUsed)
		assert.Equal(t, "healthy", retrievedMetadata.DatabaseHealth)
		assert.Equal(t, "value1", retrievedMetadata.AdditionalFields["custom_field1"])
		assert.Equal(t, float64(42), retrievedMetadata.AdditionalFields["custom_field2"])
		assert.Equal(t, true, retrievedMetadata.AdditionalFields["custom_field3"])
	})
}
