package statechange

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/service/statetracker"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) (*gorm.DB, func()) {
	connStr := "sqlite:" + t.TempDir() + "/test.db"
	db, closer, err := database.OpenWithLogger(connStr)
	require.NoError(t, err)

	err = model.GetMigrator(db).Migrate()
	require.NoError(t, err)

	return db, func() { _ = closer.Close() }
}

func createTestDeal(t *testing.T, db *gorm.DB, dealID uint64, provider string) *model.Deal {
	deal := &model.Deal{
		ID:            model.DealID(dealID),
		State:         model.DealProposed,
		Provider:      provider,
		ClientActorID: "f01000",
		PieceCID:      model.CID{},
		PieceSize:     1024,
		Price:         "1000000000000000000",
		Verified:      false,
	}
	err := db.Create(deal).Error
	require.NoError(t, err)
	return deal
}

func TestGetCommandWithStateChanges(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	ctx := context.Background()
	tracker := statetracker.NewStateChangeTracker(db)

	deal := createTestDeal(t, db, 123, "f01234")

	metadata := &statetracker.StateChangeMetadata{
		Reason:       "Test state change",
		StoragePrice: deal.Price,
	}

	prevState := deal.State
	err := tracker.TrackStateChange(ctx, deal, &prevState, model.DealPublished, metadata)
	require.NoError(t, err)

	var stateChanges []model.DealStateChange
	err = db.Where("deal_id = ?", deal.ID).Find(&stateChanges).Error
	require.NoError(t, err)
	require.Len(t, stateChanges, 1)
	require.Equal(t, model.DealPublished, stateChanges[0].NewState)
}

func TestListCommandWithFiltering(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	ctx := context.Background()
	tracker := statetracker.NewStateChangeTracker(db)

	deal1 := createTestDeal(t, db, 100, "f01234")
	deal2 := createTestDeal(t, db, 101, "f01235")

	metadata := &statetracker.StateChangeMetadata{
		Reason:       "Test filtering",
		StoragePrice: "1000000000000000000",
	}

	prevState1 := deal1.State
	err := tracker.TrackStateChange(ctx, deal1, &prevState1, model.DealActive, metadata)
	require.NoError(t, err)

	prevState2 := deal2.State
	err = tracker.TrackStateChange(ctx, deal2, &prevState2, model.DealPublished, metadata)
	require.NoError(t, err)

	var stateChanges []model.DealStateChange
	err = db.Find(&stateChanges).Error
	require.NoError(t, err)
	require.Len(t, stateChanges, 2)

	var activeChanges []model.DealStateChange
	err = db.Where("new_state = ?", model.DealActive).Find(&activeChanges).Error
	require.NoError(t, err)
	require.Len(t, activeChanges, 1)
	require.Equal(t, model.DealID(100), activeChanges[0].DealID)
}

func TestStatsCommandWithData(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	ctx := context.Background()
	tracker := statetracker.NewStateChangeTracker(db)

	deal1 := createTestDeal(t, db, 200, "f01234")
	deal2 := createTestDeal(t, db, 201, "f01234")
	deal3 := createTestDeal(t, db, 202, "f01235")

	metadata := &statetracker.StateChangeMetadata{
		Reason:       "Stats test",
		StoragePrice: "1000000000000000000",
	}

	prevState := model.DealProposed
	err := tracker.TrackStateChange(ctx, deal1, &prevState, model.DealActive, metadata)
	require.NoError(t, err)

	err = tracker.TrackStateChange(ctx, deal2, &prevState, model.DealActive, metadata)
	require.NoError(t, err)

	err = tracker.TrackStateChange(ctx, deal3, &prevState, model.DealPublished, metadata)
	require.NoError(t, err)

	var totalCount int64
	err = db.Model(&model.DealStateChange{}).Count(&totalCount).Error
	require.NoError(t, err)
	require.Equal(t, int64(3), totalCount)

	var activeCount int64
	err = db.Model(&model.DealStateChange{}).Where("new_state = ?", model.DealActive).Count(&activeCount).Error
	require.NoError(t, err)
	require.Equal(t, int64(2), activeCount)
}

func TestExportFunctionality(t *testing.T) {
	stateChanges := []model.DealStateChange{
		{
			ID:            1,
			DealID:        model.DealID(123),
			PreviousState: model.DealProposed,
			NewState:      model.DealActive,
			Timestamp:     time.Now(),
			ProviderID:    "f01234",
			ClientAddress: "f01000",
		},
	}

	tempFile := "test.csv"
	defer os.Remove(tempFile)

	err := exportStateChanges(stateChanges, "csv", tempFile)
	require.NoError(t, err)

	_, err = os.Stat(tempFile)
	require.NoError(t, err)

	err = exportStateChanges(stateChanges, "unsupported", "test.txt")
	require.Error(t, err)
	require.Contains(t, err.Error(), "unsupported export format")
}

func TestFormatOptionalHelpers(t *testing.T) {
	require.Equal(t, "", formatOptionalInt32(nil))
	value := int32(123)
	require.Equal(t, "123", formatOptionalInt32(&value))

	require.Equal(t, "", formatOptionalString(nil))
	str := "test"
	require.Equal(t, "test", formatOptionalString(&str))
}
