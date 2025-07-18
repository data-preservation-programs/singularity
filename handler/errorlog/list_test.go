package errorlog

import (
	"context"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/service/errorlog"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/require"
)

func TestDefaultHandler_ListErrorLogsHandler(t *testing.T) {
	db := testutil.TestDB(t)

	// Initialize error logging
	errorlog.Init(db)
	defer errorlog.Default.Stop()

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
			EntityType: "preparation",
			EntityID:   "1",
			EventType:  "error",
			Level:      model.ErrorLevelError,
			Message:    "Preparation failed",
			Component:  "onboard",
			Metadata:   model.ConfigMap{"key2": "value2"},
		},
		{
			CreatedAt:  testTime,
			EntityType: "schedule",
			EntityID:   "2",
			EventType:  "creation",
			Level:      model.ErrorLevelInfo,
			Message:    "Schedule created",
			Component:  "deal_schedule",
			Metadata:   model.ConfigMap{"key3": "value3"},
		},
	}

	for _, log := range testLogs {
		err := db.Create(&log).Error
		require.NoError(t, err)
	}

	ctx := context.Background()
	handler := DefaultHandler{}

	// Test: Get all logs
	logs, total, err := handler.ListErrorLogsHandler(ctx, db, errorlog.QueryFilters{
		Limit: 10,
	})
	require.NoError(t, err)
	require.Equal(t, int64(3), total)
	require.Len(t, logs, 3)

	// Test: Filter by entity type
	logs, total, err = handler.ListErrorLogsHandler(ctx, db, errorlog.QueryFilters{
		EntityType: "preparation",
		Limit:      10,
	})
	require.NoError(t, err)
	require.Equal(t, int64(2), total)
	require.Len(t, logs, 2)
	for _, log := range logs {
		require.Equal(t, "preparation", log.EntityType)
	}

	// Test: Filter by entity ID
	logs, total, err = handler.ListErrorLogsHandler(ctx, db, errorlog.QueryFilters{
		EntityID: "1",
		Limit:    10,
	})
	require.NoError(t, err)
	require.Equal(t, int64(2), total)
	require.Len(t, logs, 2)
	for _, log := range logs {
		require.Equal(t, "1", log.EntityID)
	}

	// Test: Filter by component
	logs, total, err = handler.ListErrorLogsHandler(ctx, db, errorlog.QueryFilters{
		Component: "deal_schedule",
		Limit:     10,
	})
	require.NoError(t, err)
	require.Equal(t, int64(1), total)
	require.Len(t, logs, 1)
	require.Equal(t, "deal_schedule", logs[0].Component)

	// Test: Filter by level
	logs, total, err = handler.ListErrorLogsHandler(ctx, db, errorlog.QueryFilters{
		Level: model.ErrorLevelError,
		Limit: 10,
	})
	require.NoError(t, err)
	require.Equal(t, int64(1), total)
	require.Len(t, logs, 1)
	require.Equal(t, model.ErrorLevelError, logs[0].Level)

	// Test: Filter by event type
	logs, total, err = handler.ListErrorLogsHandler(ctx, db, errorlog.QueryFilters{
		EventType: "creation",
		Limit:     10,
	})
	require.NoError(t, err)
	require.Equal(t, int64(2), total)
	require.Len(t, logs, 2)
	for _, log := range logs {
		require.Equal(t, "creation", log.EventType)
	}

	// Test: Time filtering
	startTime := testTime.Add(-30 * time.Minute)
	logs, total, err = handler.ListErrorLogsHandler(ctx, db, errorlog.QueryFilters{
		StartTime: &startTime,
		Limit:     10,
	})
	require.NoError(t, err)
	require.Equal(t, int64(1), total)
	require.Len(t, logs, 1)
	require.Equal(t, "Schedule created", logs[0].Message)

	endTime := testTime.Add(-30 * time.Minute)
	logs, total, err = handler.ListErrorLogsHandler(ctx, db, errorlog.QueryFilters{
		EndTime: &endTime,
		Limit:   10,
	})
	require.NoError(t, err)
	require.Equal(t, int64(2), total)
	require.Len(t, logs, 2)

	// Test: Pagination
	logs, total, err = handler.ListErrorLogsHandler(ctx, db, errorlog.QueryFilters{
		Limit:  1,
		Offset: 0,
	})
	require.NoError(t, err)
	require.Equal(t, int64(3), total)
	require.Len(t, logs, 1)

	logs, total, err = handler.ListErrorLogsHandler(ctx, db, errorlog.QueryFilters{
		Limit:  1,
		Offset: 1,
	})
	require.NoError(t, err)
	require.Equal(t, int64(3), total)
	require.Len(t, logs, 1)

	// Test: Combined filters
	logs, total, err = handler.ListErrorLogsHandler(ctx, db, errorlog.QueryFilters{
		EntityType: "preparation",
		Level:      model.ErrorLevelError,
		Limit:      10,
	})
	require.NoError(t, err)
	require.Equal(t, int64(1), total)
	require.Len(t, logs, 1)
	require.Equal(t, "Preparation failed", logs[0].Message)

	// Test: No results
	logs, total, err = handler.ListErrorLogsHandler(ctx, db, errorlog.QueryFilters{
		EntityType: "nonexistent",
		Limit:      10,
	})
	require.NoError(t, err)
	require.Equal(t, int64(0), total)
	require.Len(t, logs, 0)
}
