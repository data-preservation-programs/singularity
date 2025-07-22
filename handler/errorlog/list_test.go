package errorlog

import (
	"context"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/model"
	errorlogservice "github.com/data-preservation-programs/singularity/service/errorlog"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestListErrorLogsHandler(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		handler := DefaultHandler{}

		// Create test error logs
		testTime := time.Now()
		errorLogs := []model.ErrorLog{
			{
				EntityType: "preparation",
				EntityID:   "123",
				Component:  "onboard",
				Level:      model.ErrorLevelError,
				EventType:  "scan_failed",
				Message:    "Failed to scan directory",
				StackTrace: "stack trace here",
				CreatedAt:  testTime.Add(-3 * time.Hour),
			},
			{
				EntityType: "deal",
				EntityID:   "456",
				Component:  "deal_schedule",
				Level:      model.ErrorLevelWarning,
				EventType:  "deal_timeout",
				Message:    "Deal negotiation timeout",
				CreatedAt:  testTime.Add(-2 * time.Hour),
			},
			{
				EntityType: "preparation",
				EntityID:   "789",
				Component:  "pack",
				Level:      model.ErrorLevelCritical,
				EventType:  "pack_failed",
				Message:    "Critical packing error",
				CreatedAt:  testTime.Add(-1 * time.Hour),
			},
			{
				EntityType: "deal",
				EntityID:   "101",
				Component:  "deal_schedule",
				Level:      model.ErrorLevelInfo,
				EventType:  "deal_created",
				Message:    "Deal created successfully",
				CreatedAt:  testTime,
			},
		}

		// Create the test logs in database
		for _, log := range errorLogs {
			err := db.Create(&log).Error
			require.NoError(t, err)
		}

		// Test listing all error logs without filters
		logs, total, err := handler.ListErrorLogsHandler(ctx, db, errorlogservice.QueryFilters{
			Limit:  50,
			Offset: 0,
		})
		require.NoError(t, err)
		require.Equal(t, int64(4), total)
		require.Len(t, logs, 4)

		// Test filtering by entity type
		logs, total, err = handler.ListErrorLogsHandler(ctx, db, errorlogservice.QueryFilters{
			EntityType: "preparation",
			Limit:      50,
			Offset:     0,
		})
		require.NoError(t, err)
		require.Equal(t, int64(2), total)
		require.Len(t, logs, 2)
		for _, log := range logs {
			require.Equal(t, "preparation", log.EntityType)
		}

		// Test filtering by component
		logs, total, err = handler.ListErrorLogsHandler(ctx, db, errorlogservice.QueryFilters{
			Component: "deal_schedule",
			Limit:     50,
			Offset:    0,
		})
		require.NoError(t, err)
		require.Equal(t, int64(2), total)
		require.Len(t, logs, 2)
		for _, log := range logs {
			require.Equal(t, "deal_schedule", log.Component)
		}

		// Test filtering by level
		logs, total, err = handler.ListErrorLogsHandler(ctx, db, errorlogservice.QueryFilters{
			Level:  model.ErrorLevelError,
			Limit:  50,
			Offset: 0,
		})
		require.NoError(t, err)
		require.Equal(t, int64(1), total)
		require.Len(t, logs, 1)
		require.Equal(t, model.ErrorLevelError, logs[0].Level)

		// Test filtering by entity ID
		logs, total, err = handler.ListErrorLogsHandler(ctx, db, errorlogservice.QueryFilters{
			EntityID: "456",
			Limit:    50,
			Offset:   0,
		})
		require.NoError(t, err)
		require.Equal(t, int64(1), total)
		require.Len(t, logs, 1)
		require.Equal(t, "456", logs[0].EntityID)

		// Test filtering by event type
		logs, total, err = handler.ListErrorLogsHandler(ctx, db, errorlogservice.QueryFilters{
			EventType: "scan_failed",
			Limit:     50,
			Offset:    0,
		})
		require.NoError(t, err)
		require.Equal(t, int64(1), total)
		require.Len(t, logs, 1)
		require.Equal(t, "scan_failed", logs[0].EventType)

		// Test pagination with limit
		logs, total, err = handler.ListErrorLogsHandler(ctx, db, errorlogservice.QueryFilters{
			Limit:  2,
			Offset: 0,
		})
		require.NoError(t, err)
		require.Equal(t, int64(4), total)
		require.Len(t, logs, 2)

		// Test pagination with offset
		logs, total, err = handler.ListErrorLogsHandler(ctx, db, errorlogservice.QueryFilters{
			Limit:  2,
			Offset: 2,
		})
		require.NoError(t, err)
		require.Equal(t, int64(4), total)
		require.Len(t, logs, 2)

		// Test time range filtering
		startTime := testTime.Add(-2*time.Hour - 30*time.Minute)
		endTime := testTime.Add(-30 * time.Minute)
		logs, total, err = handler.ListErrorLogsHandler(ctx, db, errorlogservice.QueryFilters{
			StartTime: &startTime,
			EndTime:   &endTime,
			Limit:     50,
			Offset:    0,
		})
		require.NoError(t, err)
		require.Equal(t, int64(2), total) // Should match logs at -2h and -1h
		require.Len(t, logs, 2)

		// Test combination of filters
		logs, total, err = handler.ListErrorLogsHandler(ctx, db, errorlogservice.QueryFilters{
			EntityType: "deal",
			Component:  "deal_schedule",
			Level:      model.ErrorLevelWarning,
			Limit:      50,
			Offset:     0,
		})
		require.NoError(t, err)
		require.Equal(t, int64(1), total)
		require.Len(t, logs, 1)
		require.Equal(t, "deal", logs[0].EntityType)
		require.Equal(t, "deal_schedule", logs[0].Component)
		require.Equal(t, model.ErrorLevelWarning, logs[0].Level)
	})
}

func TestListErrorLogsHandlerEmptyResult(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		handler := DefaultHandler{}

		// Test with empty database
		logs, total, err := handler.ListErrorLogsHandler(ctx, db, errorlogservice.QueryFilters{
			Limit:  50,
			Offset: 0,
		})
		require.NoError(t, err)
		require.Equal(t, int64(0), total)
		require.Len(t, logs, 0)

		// Test with filters that don't match anything
		logs, total, err = handler.ListErrorLogsHandler(ctx, db, errorlogservice.QueryFilters{
			EntityType: "nonexistent",
			Limit:      50,
			Offset:     0,
		})
		require.NoError(t, err)
		require.Equal(t, int64(0), total)
		require.Len(t, logs, 0)
	})
}
