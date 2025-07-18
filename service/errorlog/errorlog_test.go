package errorlog

import (
	"context"
	"testing"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestErrorLogger_Basic(t *testing.T) {
	testutil.One(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		logger := New(db)
		go logger.Start()
		defer logger.Stop()

		// Test basic logging
		entry := LogEntry{
			EntityType: "preparation",
			EntityID:   "123",
			EventType:  "creation",
			Level:      model.ErrorLevelInfo,
			Message:    "Test message",
			Component:  "test",
			Metadata:   map[string]interface{}{"key": "value"},
		}

		logger.Log(entry)

		// Wait for processing
		time.Sleep(100 * time.Millisecond)

		// Verify log was persisted
		var count int64
		err := db.Model(&model.ErrorLog{}).Count(&count).Error
		require.NoError(t, err)
		require.Equal(t, int64(1), count)

		// Verify log content
		var errorLog model.ErrorLog
		err = db.First(&errorLog).Error
		require.NoError(t, err)
		require.Equal(t, "preparation", errorLog.EntityType)
		require.Equal(t, "123", errorLog.EntityID)
		require.Equal(t, "creation", errorLog.EventType)
		require.Equal(t, model.ErrorLevelInfo, errorLog.Level)
		require.Equal(t, "Test message", errorLog.Message)
		require.Equal(t, "test", errorLog.Component)
		require.Equal(t, "value", errorLog.Metadata["key"])
	})
}

func TestErrorLogger_WithError(t *testing.T) {
	testutil.One(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		logger := New(db)
		go logger.Start()
		defer logger.Stop()

		// Test logging with error
		testError := errors.New("test error")
		entry := LogEntry{
			EntityType: "deal",
			EntityID:   "456",
			EventType:  "processing_error",
			Level:      model.ErrorLevelError,
			Message:    "Processing failed",
			Error:      testError,
			Component:  "test",
		}

		logger.Log(entry)

		// Wait for processing
		time.Sleep(100 * time.Millisecond)

		// Verify log was persisted with stack trace
		var errorLog model.ErrorLog
		err := db.First(&errorLog).Error
		require.NoError(t, err)
		require.Equal(t, "deal", errorLog.EntityType)
		require.Equal(t, "456", errorLog.EntityID)
		require.Equal(t, model.ErrorLevelError, errorLog.Level)
		require.NotEmpty(t, errorLog.StackTrace)
	})
}

func TestQueryErrorLogs(t *testing.T) {
	testutil.One(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		logger := New(db)
		go logger.Start()
		defer logger.Stop()

		// Create test data
		entries := []LogEntry{
			{
				EntityType: "preparation",
				EntityID:   "1",
				EventType:  "creation",
				Level:      model.ErrorLevelInfo,
				Message:    "Preparation created",
				Component:  "onboard",
			},
			{
				EntityType: "preparation",
				EntityID:   "1",
				EventType:  "error",
				Level:      model.ErrorLevelError,
				Message:    "Preparation failed",
				Component:  "onboard",
			},
			{
				EntityType: "schedule",
				EntityID:   "2",
				EventType:  "creation",
				Level:      model.ErrorLevelInfo,
				Message:    "Schedule created",
				Component:  "deal_schedule",
			},
		}

		for _, entry := range entries {
			logger.Log(entry)
		}

		// Wait for processing
		time.Sleep(200 * time.Millisecond)

		// Test filtering by entity type
		logs, total, err := QueryErrorLogs(ctx, db, QueryFilters{
			EntityType: "preparation",
			Limit:      10,
		})
		require.NoError(t, err)
		require.Equal(t, int64(2), total)
		require.Len(t, logs, 2)

		// Test filtering by level
		logs, total, err = QueryErrorLogs(ctx, db, QueryFilters{
			Level: model.ErrorLevelError,
			Limit: 10,
		})
		require.NoError(t, err)
		require.Equal(t, int64(1), total)
		require.Len(t, logs, 1)
		require.Equal(t, "Preparation failed", logs[0].Message)

		// Test filtering by component
		logs, total, err = QueryErrorLogs(ctx, db, QueryFilters{
			Component: "deal_schedule",
			Limit:     10,
		})
		require.NoError(t, err)
		require.Equal(t, int64(1), total)
		require.Len(t, logs, 1)
		require.Equal(t, "Schedule created", logs[0].Message)

		// Test pagination
		logs, total, err = QueryErrorLogs(ctx, db, QueryFilters{
			Limit:  1,
			Offset: 1,
		})
		require.NoError(t, err)
		require.Equal(t, int64(3), total)
		require.Len(t, logs, 1)
	})
}

func TestConvenienceFunctions(t *testing.T) {
	testutil.One(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		// Initialize global logger
		Init(db)
		defer Default.Stop()

		// Test convenience functions
		LogInfo("test", "entity", "1", "event", "info message", nil)
		LogWarning("test", "entity", "2", "event", "warning message", nil)
		LogError("test", "entity", "3", "event", "error message", errors.New("test"), nil)
		LogCritical("test", "entity", "4", "event", "critical message", errors.New("critical"), nil)

		LogOnboardEvent(model.ErrorLevelInfo, "5", "started", "Onboard started", nil, nil)
		LogDealScheduleEvent(model.ErrorLevelError, "6", "failed", "Schedule failed", errors.New("schedule error"), nil)

		// Wait for processing
		time.Sleep(200 * time.Millisecond)

		// Verify all logs were created
		var count int64
		err := db.Model(&model.ErrorLog{}).Count(&count).Error
		require.NoError(t, err)
		require.Equal(t, int64(6), count)

		// Verify specific logs
		var logs []model.ErrorLog
		err = db.Order("id").Find(&logs).Error
		require.NoError(t, err)

		require.Equal(t, model.ErrorLevelInfo, logs[0].Level)
		require.Equal(t, model.ErrorLevelWarning, logs[1].Level)
		require.Equal(t, model.ErrorLevelError, logs[2].Level)
		require.Equal(t, model.ErrorLevelCritical, logs[3].Level)
		require.Equal(t, "onboard", logs[4].Component)
		require.Equal(t, "deal_schedule", logs[5].Component)
	})
}