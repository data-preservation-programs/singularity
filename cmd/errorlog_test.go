package cmd

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

func TestListCmd_Basic(t *testing.T) {
	testutil.One(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		// Create test data
		testTime := time.Now()
		testLogs := []model.ErrorLog{
			{
				CreatedAt:  testTime.Add(-2 * time.Hour),
				EntityType: "preparation",
				EntityID:   "1",
				EventType:  "creation",
				Level:      model.ErrorLevelInfo,
				Message:    "Preparation created",
				Component:  "onboard",
				Metadata:   model.ConfigMap{"key1": "value1"},
			},
			{
				CreatedAt:  testTime.Add(-1 * time.Hour),
				EntityType: "schedule",
				EntityID:   "2",
				EventType:  "error",
				Level:      model.ErrorLevelError,
				Message:    "Schedule failed",
				Component:  "deal_schedule",
				Metadata:   model.ConfigMap{"key2": "value2"},
			},
		}

		for _, log := range testLogs {
			err := db.Create(&log).Error
			require.NoError(t, err)
		}

		// Test JSON output
		runner := NewRunner().WithMode(JSON)
		defer runner.Save(t)

		output, _, err := runner.Run(ctx, "singularity error log list")
		require.NoError(t, err)

		// Parse JSON output
		var result struct {
			ErrorLogs []model.ErrorLog `json:"errorLogs"`
			Total     int64            `json:"total"`
			Limit     int              `json:"limit"`
			Offset    int              `json:"offset"`
			HasMore   bool             `json:"hasMore"`
		}
		err = json.Unmarshal([]byte(output), &result)
		require.NoError(t, err)
		require.Equal(t, int64(2), result.Total)
		require.Len(t, result.ErrorLogs, 2)
		require.False(t, result.HasMore)
	})
}

func TestListCmd_Filtering(t *testing.T) {
	testutil.One(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		// Create test data
		testLogs := []model.ErrorLog{
			{
				CreatedAt:  time.Now(),
				EntityType: "preparation",
				EntityID:   "1",
				EventType:  "creation",
				Level:      model.ErrorLevelInfo,
				Message:    "Preparation created",
				Component:  "onboard",
			},
			{
				CreatedAt:  time.Now(),
				EntityType: "schedule",
				EntityID:   "2",
				EventType:  "error",
				Level:      model.ErrorLevelError,
				Message:    "Schedule failed",
				Component:  "deal_schedule",
			},
		}

		for _, log := range testLogs {
			err := db.Create(&log).Error
			require.NoError(t, err)
		}

		// Test filtering by entity type
		runner := NewRunner().WithMode(JSON)
		defer runner.Save(t)

		output, _, err := runner.Run(ctx, "singularity error log list --entity-type=preparation")
		require.NoError(t, err)

		var result struct {
			ErrorLogs []model.ErrorLog `json:"errorLogs"`
			Total     int64            `json:"total"`
		}
		err = json.Unmarshal([]byte(output), &result)
		require.NoError(t, err)
		require.Equal(t, int64(1), result.Total)
		require.Len(t, result.ErrorLogs, 1)
		require.Equal(t, "preparation", result.ErrorLogs[0].EntityType)
	})
}

func TestListCmd_Validation(t *testing.T) {
	runner := NewRunner()
	defer runner.Save(t)

	// Test invalid error level
	_, _, err := runner.Run(context.Background(), "singularity error log list --level=invalid")
	require.Error(t, err)
	require.Contains(t, err.Error(), "invalid error level")

	// Test invalid time format
	_, _, err = runner.Run(context.Background(), "singularity error log list --start-time=invalid")
	require.Error(t, err)
	require.Contains(t, err.Error(), "invalid start-time format")
}
