package state

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/service/statetracker"
	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v2"
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

	t.Run("List All Changes", func(t *testing.T) {
		args := []string{
			"singularity", "deal", "state", "list",
		}

		app := &cli.App{
			Commands: []*cli.Command{
				{
					Name: "deal",
					Subcommands: []*cli.Command{
						{
							Name: "state",
							Subcommands: []*cli.Command{
								ListCmd,
							},
						},
					},
				},
			},
		}

		err := app.Run(args)
		require.NoError(t, err)
	})

	t.Run("List With Provider Filter", func(t *testing.T) {
		args := []string{
			"singularity", "deal", "state", "list",
			"--provider", "f01234",
		}

		app := &cli.App{
			Commands: []*cli.Command{
				{
					Name: "deal",
					Subcommands: []*cli.Command{
						{
							Name: "state",
							Subcommands: []*cli.Command{
								ListCmd,
							},
						},
					},
				},
			},
		}

		err := app.Run(args)
		require.NoError(t, err)
	})

	t.Run("List With State Filter", func(t *testing.T) {
		args := []string{
			"singularity", "deal", "state", "list",
			"--state", "active",
		}

		app := &cli.App{
			Commands: []*cli.Command{
				{
					Name: "deal",
					Subcommands: []*cli.Command{
						{
							Name: "state",
							Subcommands: []*cli.Command{
								ListCmd,
							},
						},
					},
				},
			},
		}

		err := app.Run(args)
		require.NoError(t, err)
	})
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

	t.Run("Get Overall Stats", func(t *testing.T) {
		args := []string{
			"singularity", "deal", "state", "stats",
		}

		app := &cli.App{
			Commands: []*cli.Command{
				{
					Name: "deal",
					Subcommands: []*cli.Command{
						{
							Name: "state",
							Subcommands: []*cli.Command{
								StatsCmd,
							},
						},
					},
				},
			},
		}

		err := app.Run(args)
		require.NoError(t, err)
	})

	t.Run("Get Stats By Provider", func(t *testing.T) {
		args := []string{
			"singularity", "deal", "state", "stats",
			"--by", "provider",
		}

		app := &cli.App{
			Commands: []*cli.Command{
				{
					Name: "deal",
					Subcommands: []*cli.Command{
						{
							Name: "state",
							Subcommands: []*cli.Command{
								StatsCmd,
							},
						},
					},
				},
			},
		}

		err := app.Run(args)
		require.NoError(t, err)
	})
}

func TestExportFunctionality(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	ctx := context.Background()
	tracker := statetracker.NewStateChangeTracker(db)

	// Create test data
	deal := createTestDeal(t, db, 250, "f01234")
	metadata := &statetracker.StateChangeMetadata{
		Reason:       "Export test",
		StoragePrice: "1000000000000000000",
	}

	prevState := deal.State
	err := tracker.TrackStateChange(ctx, deal, &prevState, model.DealActive, metadata)
	require.NoError(t, err)

	t.Run("Export to CSV via CLI", func(t *testing.T) {
		tempFile := "test-export.csv"
		defer os.Remove(tempFile)

		args := []string{
			"singularity", "deal", "state", "list",
			"--export", "csv",
			"--output", tempFile,
		}

		app := &cli.App{
			Commands: []*cli.Command{
				{
					Name: "deal",
					Subcommands: []*cli.Command{
						{
							Name: "state",
							Subcommands: []*cli.Command{
								ListCmd,
							},
						},
					},
				},
			},
		}

		err := app.Run(args)
		require.NoError(t, err)

		// Verify file was created
		_, err = os.Stat(tempFile)
		require.NoError(t, err)
	})

	t.Run("Export to JSON via CLI", func(t *testing.T) {
		tempFile := "test-export.json"
		defer os.Remove(tempFile)

		args := []string{
			"singularity", "deal", "state", "list",
			"--export", "json",
			"--output", tempFile,
		}

		app := &cli.App{
			Commands: []*cli.Command{
				{
					Name: "deal",
					Subcommands: []*cli.Command{
						{
							Name: "state",
							Subcommands: []*cli.Command{
								ListCmd,
							},
						},
					},
				},
			},
		}

		err := app.Run(args)
		require.NoError(t, err)

		// Verify file was created
		_, err = os.Stat(tempFile)
		require.NoError(t, err)
	})

	t.Run("Test Export Helper Functions", func(t *testing.T) {
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

		tempFile := "test-helper.csv"
		defer os.Remove(tempFile)

		err := exportStateChanges(stateChanges, "csv", tempFile)
		require.NoError(t, err)

		err = exportStateChanges(stateChanges, "unsupported", "test.txt")
		require.Error(t, err)
		require.Contains(t, err.Error(), "unsupported export format")
	})
}

func TestFormatOptionalHelpers(t *testing.T) {
	require.Equal(t, "", formatOptionalInt32(nil))
	value := int32(123)
	require.Equal(t, "123", formatOptionalInt32(&value))

	require.Equal(t, "", formatOptionalString(nil))
	str := "test"
	require.Equal(t, "test", formatOptionalString(&str))
}

func TestRepairForceTransitionCommand(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	// Create a test deal
	deal := createTestDeal(t, db, 456, "f01234")

	t.Run("Dry Run", func(t *testing.T) {
		args := []string{
			"singularity", "deal", "state", "repair", "force-transition",
			"456", "active",
			"--dry-run",
			"--reason", "Test force transition",
		}

		app := &cli.App{
			Commands: []*cli.Command{
				{
					Name: "deal",
					Subcommands: []*cli.Command{
						{
							Name: "state",
							Subcommands: []*cli.Command{
								RepairCmd,
							},
						},
					},
				},
			},
		}

		err := app.Run(args)
		require.NoError(t, err)

		// Verify deal state hasn't changed
		var updatedDeal model.Deal
		err = db.First(&updatedDeal, deal.ID).Error
		require.NoError(t, err)
		require.Equal(t, model.DealProposed, updatedDeal.State)
	})

	t.Run("Actual Force Transition", func(t *testing.T) {
		args := []string{
			"singularity", "deal", "state", "repair", "force-transition",
			"456", "active",
			"--reason", "Test force transition",
			"--epoch", "12345",
			"--sector-id", "s-f01234-100",
		}

		app := &cli.App{
			Commands: []*cli.Command{
				{
					Name: "deal",
					Subcommands: []*cli.Command{
						{
							Name: "state",
							Subcommands: []*cli.Command{
								RepairCmd,
							},
						},
					},
				},
			},
		}

		err := app.Run(args)
		require.NoError(t, err)

		// Verify deal state has changed
		var updatedDeal model.Deal
		err = db.First(&updatedDeal, deal.ID).Error
		require.NoError(t, err)
		require.Equal(t, model.DealActive, updatedDeal.State)

		// Verify state change was tracked
		var stateChange model.DealStateChange
		err = db.Where("deal_id = ? AND new_state = ?", deal.ID, model.DealActive).First(&stateChange).Error
		require.NoError(t, err)
		require.Equal(t, model.DealProposed, stateChange.PreviousState)
	})

	t.Run("Invalid State", func(t *testing.T) {
		args := []string{
			"singularity", "deal", "state", "repair", "force-transition",
			"456", "invalid_state",
		}

		app := &cli.App{
			Commands: []*cli.Command{
				{
					Name: "deal",
					Subcommands: []*cli.Command{
						{
							Name: "state",
							Subcommands: []*cli.Command{
								RepairCmd,
							},
						},
					},
				},
			},
		}

		err := app.Run(args)
		require.Error(t, err)
		require.Contains(t, err.Error(), "invalid state")
	})
}

func TestRepairResetErrorDealsCommand(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	// Create test deals with different states
	deal1 := createTestDeal(t, db, 300, "f01234")
	deal1.State = model.DealErrored
	err := db.Save(&deal1).Error
	require.NoError(t, err)

	deal2 := createTestDeal(t, db, 301, "f01234")
	deal2.State = model.DealErrored
	err = db.Save(&deal2).Error
	require.NoError(t, err)

	deal3 := createTestDeal(t, db, 302, "f01235")
	deal3.State = model.DealActive
	err = db.Save(&deal3).Error
	require.NoError(t, err)

	t.Run("Reset By Provider - Dry Run", func(t *testing.T) {
		args := []string{
			"singularity", "deal", "state", "repair", "reset-error-deals",
			"--provider", "f01234",
			"--dry-run",
		}

		app := &cli.App{
			Commands: []*cli.Command{
				{
					Name: "deal",
					Subcommands: []*cli.Command{
						{
							Name: "state",
							Subcommands: []*cli.Command{
								RepairCmd,
							},
						},
					},
				},
			},
		}

		err := app.Run(args)
		require.NoError(t, err)

		// Verify deals haven't changed
		var errorCount int64
		err = db.Model(&model.Deal{}).Where("state = ?", model.DealErrored).Count(&errorCount).Error
		require.NoError(t, err)
		require.Equal(t, int64(2), errorCount)
	})

	t.Run("Reset Specific Deals", func(t *testing.T) {
		args := []string{
			"singularity", "deal", "state", "repair", "reset-error-deals",
			"--deal-id", "300",
			"--reset-to-state", "published",
		}

		app := &cli.App{
			Commands: []*cli.Command{
				{
					Name: "deal",
					Subcommands: []*cli.Command{
						{
							Name: "state",
							Subcommands: []*cli.Command{
								RepairCmd,
							},
						},
					},
				},
			},
		}

		err := app.Run(args)
		require.NoError(t, err)

		// Verify specific deal was reset
		var updatedDeal model.Deal
		err = db.First(&updatedDeal, 300).Error
		require.NoError(t, err)
		require.Equal(t, model.DealPublished, updatedDeal.State)

		// Verify state change was tracked
		var stateChange model.DealStateChange
		err = db.Where("deal_id = ? AND previous_state = ? AND new_state = ?",
			300, model.DealErrored, model.DealPublished).First(&stateChange).Error
		require.NoError(t, err)
	})
}

func TestRepairCleanupOrphanedChangesCommand(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	ctx := context.Background()
	tracker := statetracker.NewStateChangeTracker(db)

	// Create a deal and track state changes
	deal := createTestDeal(t, db, 400, "f01234")

	metadata := &statetracker.StateChangeMetadata{
		Reason: "Test orphaned cleanup",
	}

	prevState := deal.State
	err := tracker.TrackStateChange(ctx, deal, &prevState, model.DealActive, metadata)
	require.NoError(t, err)

	// Create an orphaned state change (no corresponding deal)
	orphanedChange := &model.DealStateChange{
		DealID:        model.DealID(999),
		PreviousState: model.DealProposed,
		NewState:      model.DealActive,
		Timestamp:     time.Now(),
		ProviderID:    "f01999",
		ClientAddress: "f01000",
	}
	err = db.Create(orphanedChange).Error
	require.NoError(t, err)

	t.Run("Cleanup Orphaned Changes - Dry Run", func(t *testing.T) {
		args := []string{
			"singularity", "deal", "state", "repair", "cleanup-orphaned-changes",
			"--dry-run",
		}

		app := &cli.App{
			Commands: []*cli.Command{
				{
					Name: "deal",
					Subcommands: []*cli.Command{
						{
							Name: "state",
							Subcommands: []*cli.Command{
								RepairCmd,
							},
						},
					},
				},
			},
		}

		err := app.Run(args)
		require.NoError(t, err)

		// Verify orphaned change still exists
		var count int64
		err = db.Model(&model.DealStateChange{}).Where("deal_id = ?", 999).Count(&count).Error
		require.NoError(t, err)
		require.Equal(t, int64(1), count)
	})

	t.Run("Cleanup Orphaned Changes - Actual", func(t *testing.T) {
		args := []string{
			"singularity", "deal", "state", "repair", "cleanup-orphaned-changes",
		}

		app := &cli.App{
			Commands: []*cli.Command{
				{
					Name: "deal",
					Subcommands: []*cli.Command{
						{
							Name: "state",
							Subcommands: []*cli.Command{
								RepairCmd,
							},
						},
					},
				},
			},
		}

		err := app.Run(args)
		require.NoError(t, err)

		// Verify orphaned change was deleted
		var count int64
		err = db.Model(&model.DealStateChange{}).Where("deal_id = ?", 999).Count(&count).Error
		require.NoError(t, err)
		require.Equal(t, int64(0), count)

		// Verify valid state change still exists
		err = db.Model(&model.DealStateChange{}).Where("deal_id = ?", 400).Count(&count).Error
		require.NoError(t, err)
		require.Equal(t, int64(1), count)
	})
}

func TestEndToEndStateManagementIntegration(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	ctx := context.Background()
	tracker := statetracker.NewStateChangeTracker(db)

	// Step 1: Create multiple deals
	deals := []struct {
		id       uint64
		provider string
	}{
		{500, "f01234"},
		{501, "f01234"},
		{502, "f01235"},
		{503, "f01235"},
	}

	for _, d := range deals {
		createTestDeal(t, db, d.id, d.provider)
	}

	// Step 2: Simulate state transitions
	metadata := &statetracker.StateChangeMetadata{
		Reason:       "Integration test state changes",
		StoragePrice: "1000000000000000000",
	}

	// Deal 500: proposed -> published -> active
	deal500 := &model.Deal{}
	err := db.First(deal500, 500).Error
	require.NoError(t, err)

	prevState := deal500.State
	err = tracker.TrackStateChange(ctx, deal500, &prevState, model.DealPublished, metadata)
	require.NoError(t, err)
	deal500.State = model.DealPublished
	err = db.Save(deal500).Error
	require.NoError(t, err)

	prevState = deal500.State
	err = tracker.TrackStateChange(ctx, deal500, &prevState, model.DealActive, metadata)
	require.NoError(t, err)
	deal500.State = model.DealActive
	err = db.Save(deal500).Error
	require.NoError(t, err)

	// Deal 501: proposed -> error
	deal501 := &model.Deal{}
	err = db.First(deal501, 501).Error
	require.NoError(t, err)

	prevState = deal501.State
	err = tracker.TrackStateChange(ctx, deal501, &prevState, model.DealErrored, metadata)
	require.NoError(t, err)
	deal501.State = model.DealErrored
	err = db.Save(deal501).Error
	require.NoError(t, err)

	// Deal 502: proposed -> published (1 transition)
	deal502 := &model.Deal{}
	err = db.First(deal502, 502).Error
	require.NoError(t, err)

	prevState = deal502.State
	err = tracker.TrackStateChange(ctx, deal502, &prevState, model.DealPublished, metadata)
	require.NoError(t, err)
	deal502.State = model.DealPublished
	err = db.Save(deal502).Error
	require.NoError(t, err)

	// Step 3: Verify state tracking functionality
	t.Run("Verify State Change Tracking", func(t *testing.T) {
		var totalChanges int64
		err = db.Model(&model.DealStateChange{}).Count(&totalChanges).Error
		require.NoError(t, err)
		require.Equal(t, int64(4), totalChanges) // Exactly 4 state changes

		// Verify provider-specific changes
		var f01234Changes int64
		err = db.Model(&model.DealStateChange{}).Where("provider_id = ?", "f01234").Count(&f01234Changes).Error
		require.NoError(t, err)
		require.Equal(t, int64(3), f01234Changes) // 2 for deal 500, 1 for deal 501

		// Verify state-specific changes
		var activeChanges int64
		err = db.Model(&model.DealStateChange{}).Where("new_state = ?", model.DealActive).Count(&activeChanges).Error
		require.NoError(t, err)
		require.Equal(t, int64(1), activeChanges) // Only deal 500 went active
	})

	// Step 4: Test export functionality at unit level
	t.Run("Export State Changes to Files", func(t *testing.T) {
		var stateChanges []model.DealStateChange
		err = db.Find(&stateChanges).Error
		require.NoError(t, err)
		require.Len(t, stateChanges, 4)

		// Test CSV export
		csvFile := "integration-test.csv"
		defer os.Remove(csvFile)
		err = exportStateChanges(stateChanges, "csv", csvFile)
		require.NoError(t, err)

		fileInfo, err := os.Stat(csvFile)
		require.NoError(t, err)
		require.Greater(t, fileInfo.Size(), int64(0))

		// Test JSON export
		jsonFile := "integration-test.json"
		defer os.Remove(jsonFile)
		err = exportStateChanges(stateChanges, "json", jsonFile)
		require.NoError(t, err)

		fileInfo, err = os.Stat(jsonFile)
		require.NoError(t, err)
		require.Greater(t, fileInfo.Size(), int64(0))
	})

	// Step 5: Test repair functionality at database level
	t.Run("Error Deal Reset Simulation", func(t *testing.T) {
		// Find error deals for a specific provider
		var errorDeals []model.Deal
		err = db.Where("state = ? AND provider = ?", model.DealErrored, "f01234").Find(&errorDeals).Error
		require.NoError(t, err)
		require.Len(t, errorDeals, 1) // Deal 501

		// Simulate reset operation
		for _, deal := range errorDeals {
			metadata := &statetracker.StateChangeMetadata{
				Reason: "Manual error state reset",
				AdditionalFields: map[string]interface{}{
					"operationType": "error_state_reset",
					"operator":      "test",
				},
			}

			previousState := &deal.State
			err = tracker.TrackStateChangeWithDetails(
				ctx,
				deal.ID,
				previousState,
				model.DealProposed,
				nil,
				nil,
				deal.Provider,
				deal.ClientActorID,
				metadata,
			)
			require.NoError(t, err)

			// Update the deal state
			err = db.Model(&deal).Update("state", model.DealProposed).Error
			require.NoError(t, err)
		}

		// Verify reset worked
		var updatedDeal model.Deal
		err = db.First(&updatedDeal, 501).Error
		require.NoError(t, err)
		require.Equal(t, model.DealProposed, updatedDeal.State)

		// Verify state change was tracked (now 5 total changes)
		var totalChanges int64
		err = db.Model(&model.DealStateChange{}).Count(&totalChanges).Error
		require.NoError(t, err)
		require.Equal(t, int64(5), totalChanges)
	})

	// Step 6: Test comprehensive statistics
	t.Run("State Change Statistics", func(t *testing.T) {
		// Test overall stats
		var stats struct {
			TotalChanges    int64
			UniqueDeals     int64
			UniqueProviders int64
		}

		err = db.Model(&model.DealStateChange{}).Count(&stats.TotalChanges).Error
		require.NoError(t, err)

		err = db.Model(&model.DealStateChange{}).Distinct("deal_id").Count(&stats.UniqueDeals).Error
		require.NoError(t, err)

		err = db.Model(&model.DealStateChange{}).Distinct("provider_id").Count(&stats.UniqueProviders).Error
		require.NoError(t, err)

		require.Equal(t, int64(5), stats.TotalChanges)
		require.Equal(t, int64(3), stats.UniqueDeals)     // deals 500, 501, 502
		require.Equal(t, int64(2), stats.UniqueProviders) // f01234, f01235

		// Test state distribution
		stateDistribution := make(map[string]int64)
		rows, err := db.Model(&model.DealStateChange{}).
			Select("new_state, COUNT(*) as count").
			Group("new_state").
			Rows()
		require.NoError(t, err)
		defer rows.Close()

		for rows.Next() {
			var state string
			var count int64
			err = rows.Scan(&state, &count)
			require.NoError(t, err)
			stateDistribution[state] = count
		}

		require.Equal(t, int64(2), stateDistribution["published"]) // deals 500, 502
		require.Equal(t, int64(1), stateDistribution["active"])    // deal 500
		require.Equal(t, int64(1), stateDistribution["error"])     // deal 501 (original)
		require.Equal(t, int64(1), stateDistribution["proposed"])  // deal 501 (reset)
	})
}
