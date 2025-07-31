package statetracker

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestCategorizeError(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		expected string
	}{
		{
			name:     "Network error",
			err:      errors.New("network connection failed"),
			expected: ErrorCategoryNetwork,
		},
		{
			name:     "Provider error",
			err:      errors.New("storage provider not responding"),
			expected: ErrorCategoryProvider,
		},
		{
			name:     "Client error",
			err:      errors.New("client wallet insufficient funds"),
			expected: ErrorCategoryClient,
		},
		{
			name:     "Chain error",
			err:      errors.New("chain consensus issue"),
			expected: ErrorCategoryChain,
		},
		{
			name:     "Database error",
			err:      errors.New("database connection timeout"),
			expected: ErrorCategoryDB,
		},
		{
			name:     "Timeout error",
			err:      errors.New("operation timeout exceeded"),
			expected: ErrorCategoryTimeout,
		},
		{
			name:     "Funding error",
			err:      errors.New("insufficient collateral balance"),
			expected: ErrorCategoryFunding,
		},
		{
			name:     "Slashing error",
			err:      errors.New("deal was slashed"),
			expected: ErrorCategorySlashing,
		},
		{
			name:     "Expiry error",
			err:      errors.New("deal expired"),
			expected: ErrorCategoryExpiry,
		},
		{
			name:     "Unknown error",
			err:      errors.New("some unknown error"),
			expected: ErrorCategoryInternal,
		},
		{
			name:     "Nil error",
			err:      nil,
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CategorizeError(tt.err)
			require.Equal(t, tt.expected, result)
		})
	}
}

func TestCreateEnhancedMetadata(t *testing.T) {
	activationEpoch := int32(100)
	expirationEpoch := int32(200)
	slashingEpoch := int32(150)

	dealInfo := &DealInfo{
		StoragePrice:    "1000",
		PieceSize:       1024,
		VerifiedDeal:    true,
		ActivationEpoch: &activationEpoch,
		ExpirationEpoch: &expirationEpoch,
		SlashingEpoch:   &slashingEpoch,
		PublishCID:      "bafy123",
		TransactionID:   "tx123",
		ChainTipSetKey:  "tipset123",
	}

	err := errors.New("test error")
	reason := "Test state change"

	metadata := CreateEnhancedMetadata(reason, err, dealInfo)

	require.Equal(t, reason, metadata.Reason)
	require.Equal(t, err.Error(), metadata.Error)
	require.Equal(t, ErrorCategoryInternal, metadata.ErrorCategory)
	require.Equal(t, "1000", metadata.StoragePrice)
	require.Equal(t, int64(1024), metadata.PieceSize)
	require.True(t, metadata.VerifiedDeal)
	require.Equal(t, int32(100), *metadata.ActivationEpoch)
	require.Equal(t, int32(200), *metadata.ExpirationEpoch)
	require.Equal(t, int32(150), *metadata.SlashingEpoch)
	require.Equal(t, "bafy123", metadata.PublishCID)
	require.Equal(t, "tx123", metadata.TransactionID)
	require.Equal(t, "tipset123", metadata.ChainTipSetKey)
}

func TestCreateEnhancedMetadataWithNil(t *testing.T) {
	metadata := CreateEnhancedMetadata("test", nil, nil)

	require.Equal(t, "test", metadata.Reason)
	require.Empty(t, metadata.Error)
	require.Empty(t, metadata.ErrorCategory)
	require.Empty(t, metadata.StoragePrice)
	require.Equal(t, int64(0), metadata.PieceSize)
	require.False(t, metadata.VerifiedDeal)
}

func TestTrackStateChangeWithError(t *testing.T) {
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

		// Test tracking state change with error
		testErr := errors.New("network connection failed")
		dealInfo := &DealInfo{
			StoragePrice: "1000",
			PieceSize:    1024,
			VerifiedDeal: false,
		}

		previousState := model.DealProposed
		epochHeight := int32(150)
		err = tracker.TrackStateChangeWithError(
			ctx,
			deal.ID,
			&previousState,
			model.DealErrored,
			&epochHeight,
			nil,
			deal.Provider,
			deal.ClientActorID,
			testErr,
			dealInfo,
		)
		require.NoError(t, err)

		// Verify the state change was recorded with enhanced metadata
		var stateChanges []model.DealStateChange
		err = db.Where("deal_id = ?", deal.ID).Find(&stateChanges).Error
		require.NoError(t, err)
		require.Len(t, stateChanges, 1)

		sc := stateChanges[0]
		require.Equal(t, deal.ID, sc.DealID)
		require.Equal(t, model.DealProposed, sc.PreviousState)
		require.Equal(t, model.DealErrored, sc.NewState)
		require.Equal(t, deal.Provider, sc.ProviderID)
		require.Equal(t, deal.ClientActorID, sc.ClientAddress)
		require.Equal(t, int32(150), *sc.EpochHeight)

		// Verify enhanced metadata was serialized correctly
		var savedMetadata StateChangeMetadata
		err = json.Unmarshal([]byte(sc.Metadata), &savedMetadata)
		require.NoError(t, err)
		require.Contains(t, savedMetadata.Reason, "network")
		require.Equal(t, testErr.Error(), savedMetadata.Error)
		require.Equal(t, ErrorCategoryNetwork, savedMetadata.ErrorCategory)
		require.Equal(t, "1000", savedMetadata.StoragePrice)
		require.Equal(t, int64(1024), savedMetadata.PieceSize)
		require.False(t, savedMetadata.VerifiedDeal)
	})
}

func TestEnhancedStateChangeStats(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		tracker := NewStateChangeTracker(db)

		// Create test deals
		deal1 := &model.Deal{
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
		deal2 := &model.Deal{
			State:         model.DealSlashed,
			Provider:      "f01235",
			ClientActorID: "f01001",
			PieceCID:      model.CID{},
			PieceSize:     2048,
			StartEpoch:    100,
			EndEpoch:      200,
			Price:         "2000",
			Verified:      true,
		}
		err := db.Create([]model.Deal{*deal1, *deal2}).Error
		require.NoError(t, err)

		// Create state changes with enhanced metadata
		dealInfo1 := &DealInfo{
			StoragePrice: "1000",
			PieceSize:    1024,
			VerifiedDeal: false,
		}
		dealInfo2 := &DealInfo{
			StoragePrice: "2000",
			PieceSize:    2048,
			VerifiedDeal: true,
		}

		slashErr := errors.New("deal was slashed due to sector fault")
		epochHeight1 := int32(150)
		err = tracker.TrackStateChangeWithError(
			ctx,
			deal1.ID,
			nil,
			model.DealActive,
			&epochHeight1,
			nil,
			deal1.Provider,
			deal1.ClientActorID,
			nil,
			dealInfo1,
		)
		require.NoError(t, err)

		epochHeight2 := int32(175)
		err = tracker.TrackStateChangeWithError(
			ctx,
			deal2.ID,
			nil,
			model.DealSlashed,
			&epochHeight2,
			nil,
			deal2.Provider,
			deal2.ClientActorID,
			slashErr,
			dealInfo2,
		)
		require.NoError(t, err)

		// Test enhanced stats retrieval
		stats, err := tracker.GetStateChangeStats(ctx)
		require.NoError(t, err)

		// Verify basic stats
		require.Contains(t, stats, "totalStateChanges")
		require.Contains(t, stats, "stateDistribution")
		require.Contains(t, stats, "recentStateChanges24h")
		require.Contains(t, stats, "topProvidersByStateChanges")

		require.Equal(t, int64(2), stats["totalStateChanges"])
		require.Equal(t, int64(2), stats["recentStateChanges24h"])

		// Verify state distribution includes both states
		stateDistribution := stats["stateDistribution"].([]struct {
			State string `json:"state"`
			Count int64  `json:"count"`
		})
		require.Len(t, stateDistribution, 2)

		// Verify top providers includes both providers
		topProviders := stats["topProvidersByStateChanges"].([]struct {
			ProviderID string `json:"providerId"`
			Count      int64  `json:"count"`
		})
		require.Len(t, topProviders, 2)
	})
}
