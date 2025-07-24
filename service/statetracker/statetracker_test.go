package statetracker

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
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
