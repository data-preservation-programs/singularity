package cmd

import (
	"context"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestErrorLogListCommand(t *testing.T) {
	testutil.One(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		runner := NewRunner()
		defer runner.Save(t)

		// Create some test error logs in the database
		testTime := time.Now()
		testLogs := []model.ErrorLog{
			{
				EntityType: "preparation",
				EntityID:   "123",
				Component:  "onboard",
				Level:      model.ErrorLevelError,
				EventType:  "scan_failed",
				Message:    "Failed to scan directory",
				CreatedAt:  testTime,
			},
			{
				EntityType: "deal",
				EntityID:   "456",
				Component:  "deal_schedule",
				Level:      model.ErrorLevelWarning,
				EventType:  "deal_timeout",
				Message:    "Deal negotiation timeout",
				CreatedAt:  testTime.Add(-1 * time.Hour),
			},
		}

		// Create the test logs in database
		for _, log := range testLogs {
			err := db.Create(&log).Error
			require.NoError(t, err)
		}

		// Test basic list without filters
		_, _, err := runner.Run(ctx, "singularity error log list")
		require.NoError(t, err)

		// Test with entity type filter
		_, _, err = runner.Run(ctx, "singularity error log list --entity-type preparation")
		require.NoError(t, err)

		// Test with multiple filters
		_, _, err = runner.Run(ctx, "singularity error log list --entity-type deal --component deal_schedule --level warning --limit 10 --offset 0")
		require.NoError(t, err)

		// Test with time filters
		startTime := testTime.Add(-2 * time.Hour)
		endTime := testTime
		_, _, err = runner.Run(ctx, "singularity error log list --start-time "+startTime.Format(time.RFC3339)+" --end-time "+endTime.Format(time.RFC3339))
		require.NoError(t, err)

		// Test verbose mode
		_, _, err = runner.Run(ctx, "singularity --verbose error log list")
		require.NoError(t, err)

		// Test JSON mode
		_, _, err = runner.Run(ctx, "singularity --json error log list")
		require.NoError(t, err)
	})
}

func TestErrorLogListValidation(t *testing.T) {
	testutil.OneWithoutReset(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		runner := NewRunner()
		defer runner.Save(t)

		// Test invalid error level
		_, _, err := runner.Run(ctx, "singularity error log list --level invalid")
		require.Error(t, err)
		require.Contains(t, err.Error(), "invalid error level")

		// Test invalid start time format
		_, _, err = runner.Run(ctx, "singularity error log list --start-time invalid-time")
		require.Error(t, err)
		require.Contains(t, err.Error(), "invalid start-time format")

		// Test invalid end time format
		_, _, err = runner.Run(ctx, "singularity error log list --end-time invalid-time")
		require.Error(t, err)
		require.Contains(t, err.Error(), "invalid end-time format")
	})
}

func TestErrorLogListEmptyResult(t *testing.T) {
	testutil.One(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		runner := NewRunner()
		defer runner.Save(t)

		// Test empty result (no error logs in database)
		_, _, err := runner.Run(ctx, "singularity error log list")
		require.NoError(t, err)

		// Test with filters on empty database
		_, _, err = runner.Run(ctx, "singularity error log list --entity-type preparation --level error")
		require.NoError(t, err)
	})
}
