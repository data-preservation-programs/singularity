package statechange

import (
	"context"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestGetDealStateChangesHandler(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		handler := DefaultHandler{}

		// Test with non-existent deal
		_, err := handler.GetDealStateChangesHandler(ctx, db, 999)
		require.ErrorIs(t, err, handlererror.ErrNotFound)

		// Create a test deal
		deal := model.Deal{
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
		err = db.Create(&deal).Error
		require.NoError(t, err)

		// Create test state changes
		stateChanges := []model.DealStateChange{
			{
				DealID:        deal.ID,
				PreviousState: "",
				NewState:      model.DealProposed,
				Timestamp:     time.Now().Add(-2 * time.Hour),
				ProviderID:    "f01234",
				ClientAddress: "f01000",
				Metadata:      `{"reason": "Initial proposal"}`,
			},
			{
				DealID:        deal.ID,
				PreviousState: model.DealProposed,
				NewState:      model.DealPublished,
				Timestamp:     time.Now().Add(-1 * time.Hour),
				ProviderID:    "f01234",
				ClientAddress: "f01000",
				Metadata:      `{"reason": "Deal published"}`,
			},
		}

		for _, sc := range stateChanges {
			err = db.Create(&sc).Error
			require.NoError(t, err)
		}

		// Test successful retrieval
		result, err := handler.GetDealStateChangesHandler(ctx, db, deal.ID)
		require.NoError(t, err)
		require.Len(t, result, 2)
		
		// Verify ordering (should be ascending by timestamp)
		require.Equal(t, model.DealProposed, result[0].NewState)
		require.Equal(t, model.DealPublished, result[1].NewState)
	})
}

func TestListStateChangesHandler(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *testutil.TestDB) {
		handler := DefaultHandler{}

		// Create test deals
		deal1 := model.Deal{
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
		deal2 := model.Deal{
			State:         model.DealExpired,
			Provider:      "f05678",
			ClientActorID: "f01001",
			PieceCID:      model.CID{},
			PieceSize:     2048,
			StartEpoch:    150,
			EndEpoch:      250,
			Price:         "2000",
			Verified:      true,
		}
		
		err := db.DB.Create(&deal1).Error
		require.NoError(t, err)
		err = db.DB.Create(&deal2).Error
		require.NoError(t, err)

		// Create test state changes
		baseTime := time.Now().Add(-24 * time.Hour)
		stateChanges := []model.DealStateChange{
			{
				DealID:        deal1.ID,
				PreviousState: "",
				NewState:      model.DealProposed,
				Timestamp:     baseTime,
				ProviderID:    "f01234",
				ClientAddress: "f01000",
				Metadata:      `{"reason": "Initial proposal"}`,
			},
			{
				DealID:        deal1.ID,
				PreviousState: model.DealProposed,
				NewState:      model.DealActive,
				Timestamp:     baseTime.Add(1 * time.Hour),
				ProviderID:    "f01234",
				ClientAddress: "f01000",
				Metadata:      `{"reason": "Deal activated"}`,
			},
			{
				DealID:        deal2.ID,
				PreviousState: "",
				NewState:      model.DealProposed,
				Timestamp:     baseTime.Add(2 * time.Hour),
				ProviderID:    "f05678",
				ClientAddress: "f01001",
				Metadata:      `{"reason": "Initial proposal"}`,
			},
			{
				DealID:        deal2.ID,
				PreviousState: model.DealProposed,
				NewState:      model.DealExpired,
				Timestamp:     baseTime.Add(3 * time.Hour),
				ProviderID:    "f05678",
				ClientAddress: "f01001",
				Metadata:      `{"reason": "Deal expired"}`,
			},
		}

		for _, sc := range stateChanges {
			err = db.DB.Create(&sc).Error
			require.NoError(t, err)
		}

		// Test listing all state changes
		query := model.DealStateChangeQuery{}
		result, err := handler.ListStateChangesHandler(ctx, db.DB, query)
		require.NoError(t, err)
		require.Equal(t, int64(4), result.Total)
		require.Len(t, result.StateChanges, 4)

		// Test filtering by deal ID
		query = model.DealStateChangeQuery{
			DealID: &deal1.ID,
		}
		result, err = handler.ListStateChangesHandler(ctx, db.DB, query)
		require.NoError(t, err)
		require.Equal(t, int64(2), result.Total)
		require.Len(t, result.StateChanges, 2)

		// Test filtering by state
		state := model.DealExpired
		query = model.DealStateChangeQuery{
			State: &state,
		}
		result, err = handler.ListStateChangesHandler(ctx, db.DB, query)
		require.NoError(t, err)
		require.Equal(t, int64(1), result.Total)
		require.Len(t, result.StateChanges, 1)
		require.Equal(t, model.DealExpired, result.StateChanges[0].NewState)

		// Test filtering by provider
		provider := "f01234"
		query = model.DealStateChangeQuery{
			ProviderID: &provider,
		}
		result, err = handler.ListStateChangesHandler(ctx, db.DB, query)
		require.NoError(t, err)
		require.Equal(t, int64(2), result.Total)
		require.Len(t, result.StateChanges, 2)

		// Test pagination
		limit := 2
		query = model.DealStateChangeQuery{
			Limit: &limit,
		}
		result, err = handler.ListStateChangesHandler(ctx, db.DB, query)
		require.NoError(t, err)
		require.Equal(t, int64(4), result.Total)
		require.Len(t, result.StateChanges, 2)

		// Test time range filtering
		startTime := baseTime.Add(1 * time.Hour)
		endTime := baseTime.Add(2 * time.Hour + 30*time.Minute)
		query = model.DealStateChangeQuery{
			StartTime: &startTime,
			EndTime:   &endTime,
		}
		result, err = handler.ListStateChangesHandler(ctx, db.DB, query)
		require.NoError(t, err)
		require.Equal(t, int64(2), result.Total)
		require.Len(t, result.StateChanges, 2)
	})
}

func TestGetStateChangeStatsHandler(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *testutil.TestDB) {
		handler := DefaultHandler{}

		// Create test deal
		deal := model.Deal{
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
		err := db.DB.Create(&deal).Error
		require.NoError(t, err)

		// Create test state changes
		stateChanges := []model.DealStateChange{
			{
				DealID:        deal.ID,
				PreviousState: "",
				NewState:      model.DealProposed,
				Timestamp:     time.Now().Add(-1 * time.Hour),
				ProviderID:    "f01234",
				ClientAddress: "f01000",
				Metadata:      `{"reason": "Initial proposal"}`,
			},
			{
				DealID:        deal.ID,
				PreviousState: model.DealProposed,
				NewState:      model.DealActive,
				Timestamp:     time.Now().Add(-30 * time.Minute),
				ProviderID:    "f01234",
				ClientAddress: "f01000",
				Metadata:      `{"reason": "Deal activated"}`,
			},
		}

		for _, sc := range stateChanges {
			err = db.DB.Create(&sc).Error
			require.NoError(t, err)
		}

		// Test stats retrieval
		stats, err := handler.GetStateChangeStatsHandler(ctx, db.DB)
		require.NoError(t, err)
		
		// Check that we get expected stats structure
		require.Contains(t, stats, "totalStateChanges")
		require.Contains(t, stats, "stateDistribution")
		require.Contains(t, stats, "recentStateChanges24h")
		require.Contains(t, stats, "topProvidersByStateChanges")
		
		// Verify basic counts
		require.Equal(t, int64(2), stats["totalStateChanges"])
		require.Equal(t, int64(2), stats["recentStateChanges24h"])
	})
}