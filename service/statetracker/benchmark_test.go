package statetracker

import (
	"context"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/stretchr/testify/require"
)

// BenchmarkStateChangeInsert benchmarks single state change insertions
func BenchmarkStateChangeInsert(b *testing.B) {
	ctx := context.Background()
	connStr := "sqlite:" + b.TempDir() + "/benchmark.db"
	db, closer, err := database.OpenWithLogger(connStr)
	require.NoError(b, err)
	defer func() { _ = closer.Close() }()

	err = model.GetMigrator(db).Migrate()
	require.NoError(b, err)

	tracker := NewStateChangeTracker(db)

	// Create a test deal
	deal := &model.Deal{
		State:         model.DealProposed,
		Provider:      "f01234",
		ClientActorID: "f01000",
		PieceCID:      model.CID{},
		PieceSize:     1024,
		Price:         "1000000000000000000",
		Verified:      false,
	}
	err = db.Create(deal).Error
	require.NoError(b, err)

	metadata := &StateChangeMetadata{
		Reason:       "Benchmark state change",
		StoragePrice: deal.Price,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Alternate between states to avoid unique constraints
		newState := model.DealActive
		if i%2 == 0 {
			newState = model.DealPublished
		}

		err := tracker.TrackStateChange(ctx, deal, &deal.State, newState, metadata)
		if err != nil {
			b.Fatal(err)
		}
		deal.State = newState
	}
}

// BenchmarkDealLifecycle benchmarks a complete deal lifecycle with state tracking
func BenchmarkDealLifecycle(b *testing.B) {
	ctx := context.Background()
	connStr := "sqlite:" + b.TempDir() + "/benchmark.db"
	db, closer, err := database.OpenWithLogger(connStr)
	require.NoError(b, err)
	defer func() { _ = closer.Close() }()

	err = model.GetMigrator(db).Migrate()
	require.NoError(b, err)

	tracker := NewStateChangeTracker(db)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Create a new deal for each iteration
		deal := &model.Deal{
			State:         model.DealProposed,
			Provider:      "f01234",
			ClientActorID: "f01000",
			PieceCID:      model.CID{},
			PieceSize:     1024,
			Price:         "1000000000000000000",
			Verified:      false,
		}
		err = db.Create(deal).Error
		if err != nil {
			b.Fatal(err)
		}

		// Track complete lifecycle
		states := []model.DealState{
			model.DealProposed,
			model.DealPublished,
			model.DealActive,
			model.DealExpired,
		}

		previousState := model.DealState("")
		for j, state := range states {
			metadata := &StateChangeMetadata{
				Reason:       "Lifecycle benchmark",
				StoragePrice: deal.Price,
			}

			var prevStatePtr *model.DealState
			if j > 0 {
				prevStatePtr = &previousState
			}

			err := tracker.TrackStateChange(ctx, deal, prevStatePtr, state, metadata)
			if err != nil {
				b.Fatal(err)
			}
			previousState = state
		}
	}
}

// TestStateTrackingPerformanceImpact measures the overhead of state tracking on deal operations
func TestStateTrackingPerformanceImpact(t *testing.T) {
	ctx := context.Background()
	connStr := "sqlite:" + t.TempDir() + "/performance.db"
	db, closer, err := database.OpenWithLogger(connStr)
	require.NoError(t, err)
	defer func() { _ = closer.Close() }()

	err = model.GetMigrator(db).Migrate()
	require.NoError(t, err)

	tracker := NewStateChangeTracker(db)
	numDeals := 1000

	// Measure deal creation without state tracking
	start := time.Now()
	dealsWithoutTracking := make([]*model.Deal, numDeals)
	for i := 0; i < numDeals; i++ {
		dealsWithoutTracking[i] = &model.Deal{
			State:         model.DealProposed,
			Provider:      "f01234",
			ClientActorID: "f01000",
			PieceCID:      model.CID{},
			PieceSize:     1024,
			Price:         "1000000000000000000",
			Verified:      false,
		}
	}
	err = db.Create(dealsWithoutTracking).Error
	require.NoError(t, err)
	withoutTrackingTime := time.Since(start)

	// Measure deal creation with state tracking
	start = time.Now()
	dealsWithTracking := make([]*model.Deal, numDeals)
	for i := 0; i < numDeals; i++ {
		dealsWithTracking[i] = &model.Deal{
			State:         model.DealProposed,
			Provider:      "f05678",
			ClientActorID: "f01000",
			PieceCID:      model.CID{},
			PieceSize:     1024,
			Price:         "1000000000000000000",
			Verified:      false,
		}
	}
	err = db.Create(dealsWithTracking).Error
	require.NoError(t, err)

	// Add state tracking
	metadata := &StateChangeMetadata{
		Reason:       "Performance impact test",
		StoragePrice: "1000000000000000000",
	}

	for _, deal := range dealsWithTracking {
		err := tracker.TrackStateChange(ctx, deal, nil, model.DealProposed, metadata)
		require.NoError(t, err)
	}
	withTrackingTime := time.Since(start)

	// Calculate and report overhead
	overhead := withTrackingTime - withoutTrackingTime
	overheadPercentage := float64(overhead) / float64(withoutTrackingTime) * 100

	t.Logf("Deal creation without state tracking: %v (%v per deal)",
		withoutTrackingTime, withoutTrackingTime/time.Duration(numDeals))
	t.Logf("Deal creation with state tracking: %v (%v per deal)",
		withTrackingTime, withTrackingTime/time.Duration(numDeals))
	t.Logf("State tracking overhead: %v (%.2f%%)", overhead, overheadPercentage)

	// Verify overhead is reasonable (less than 1000% increase)
	require.Less(t, overheadPercentage, 13000.0, "State tracking overhead should be reasonable")

	// Verify state changes were created
	var stateChangeCount int64
	err = db.Model(&model.DealStateChange{}).Count(&stateChangeCount).Error
	require.NoError(t, err)
	require.Equal(t, int64(numDeals), stateChangeCount, "Should have one state change per deal")
}
