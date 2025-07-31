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

	// Step 3: List all state changes
	t.Run("List All State Changes", func(t *testing.T) {
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

	// Step 4: Get specific deal state history
	t.Run("Get Deal 500 State History", func(t *testing.T) {
		args := []string{
			"singularity", "deal", "state", "get", "500",
		}

		app := &cli.App{
			Commands: []*cli.Command{
				{
					Name: "deal",
					Subcommands: []*cli.Command{
						{
							Name: "state",
							Subcommands: []*cli.Command{
								GetCmd,
							},
						},
					},
				},
			},
		}

		err := app.Run(args)
		require.NoError(t, err)
	})

	// Step 5: Get stats by provider
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

	// Step 6: Export state changes to CSV
	t.Run("Export State Changes", func(t *testing.T) {
		tempFile := "integration-export.csv"
		defer os.Remove(tempFile)

		args := []string{
			"singularity", "deal", "state", "list",
			"--provider", "f01234",
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

		// Verify export file exists
		fileInfo, err := os.Stat(tempFile)
		require.NoError(t, err)
		require.Greater(t, fileInfo.Size(), int64(0))
	})

	// Step 7: Reset error deals
	t.Run("Reset Error Deals", func(t *testing.T) {
		args := []string{
			"singularity", "deal", "state", "repair", "reset-error-deals",
			"--provider", "f01234",
			"--reset-to-state", "proposed",
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

		// Verify deal 501 was reset
		var updatedDeal model.Deal
		err = db.First(&updatedDeal, 501).Error
		require.NoError(t, err)
		require.Equal(t, model.DealProposed, updatedDeal.State)
	})

	// Verify the complete state history
	var totalChanges int64
	err = db.Model(&model.DealStateChange{}).Count(&totalChanges).Error
	require.NoError(t, err)
	require.GreaterOrEqual(t, totalChanges, int64(4)) // At least 4 state changes
}
