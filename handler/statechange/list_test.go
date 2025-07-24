package statechange

import (
	"context"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestListStateChangesHandler(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		handler := DefaultHandler{}

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
				NewState:      model.DealActive,
				Timestamp:     time.Now().Add(-1 * time.Hour),
				ProviderID:    "f01234",
				ClientAddress: "f01000",
				Metadata:      `{"reason": "Deal activated"}`,
			},
		}

		for _, sc := range stateChanges {
			err = db.Create(&sc).Error
			require.NoError(t, err)
		}

		// Test list all state changes
		query := model.DealStateChangeQuery{}
		response, err := handler.ListStateChangesHandler(ctx, db, query)
		require.NoError(t, err)
		require.Equal(t, int64(2), response.Total)
		require.Len(t, response.StateChanges, 2)

		// Test filtering by deal ID
		query = model.DealStateChangeQuery{
			DealID: &deal.ID,
		}
		response, err = handler.ListStateChangesHandler(ctx, db, query)
		require.NoError(t, err)
		require.Equal(t, int64(2), response.Total)
		require.Len(t, response.StateChanges, 2)

		// Test filtering by state
		state := model.DealActive
		query = model.DealStateChangeQuery{
			State: &state,
		}
		response, err = handler.ListStateChangesHandler(ctx, db, query)
		require.NoError(t, err)
		require.Equal(t, int64(1), response.Total)
		require.Len(t, response.StateChanges, 1)
		require.Equal(t, model.DealActive, response.StateChanges[0].NewState)

		// Test pagination
		limit := 1
		query = model.DealStateChangeQuery{
			Limit: &limit,
		}
		response, err = handler.ListStateChangesHandler(ctx, db, query)
		require.NoError(t, err)
		require.Equal(t, int64(2), response.Total)
		require.Len(t, response.StateChanges, 1)
	})
}

func TestGetDealStateChangesHandler(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		handler := DefaultHandler{}

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
				NewState:      model.DealActive,
				Timestamp:     time.Now().Add(-1 * time.Hour),
				ProviderID:    "f01234",
				ClientAddress: "f01000",
				Metadata:      `{"reason": "Deal activated"}`,
			},
		}

		for _, sc := range stateChanges {
			err = db.Create(&sc).Error
			require.NoError(t, err)
		}

		// Test getting state changes for existing deal
		changes, err := handler.GetDealStateChangesHandler(ctx, db, deal.ID)
		require.NoError(t, err)
		require.Len(t, changes, 2)

		// Verify they are ordered by timestamp ASC (oldest first)
		require.True(t, changes[0].Timestamp.Before(changes[1].Timestamp))
		require.Equal(t, model.DealProposed, changes[0].NewState)
		require.Equal(t, model.DealActive, changes[1].NewState)

		// Test getting state changes for non-existent deal
		_, err = handler.GetDealStateChangesHandler(ctx, db, model.DealID(99999))
		require.Error(t, err)
	})
}

func TestGetStateChangeStatsHandler(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		handler := DefaultHandler{}

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
			err = db.Create(&sc).Error
			require.NoError(t, err)
		}

		// Test getting stats
		stats, err := handler.GetStateChangeStatsHandler(ctx, db)
		require.NoError(t, err)

		// Verify stats structure
		require.Contains(t, stats, "totalStateChanges")
		require.Contains(t, stats, "stateDistribution")
		require.Contains(t, stats, "recentStateChanges24h")
		require.Contains(t, stats, "topProvidersByStateChanges")

		// Verify values
		require.Equal(t, int64(2), stats["totalStateChanges"])
		require.Equal(t, int64(2), stats["recentStateChanges24h"])
	})
}
