package errorlog

import (
	"context"
	"fmt"
	"runtime"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/ipfs/go-log/v2"
	"gorm.io/gorm"
)

var logger = log.Logger("errorlog")

// LogEntry represents a structured error log entry before persistence
type LogEntry struct {
	EntityType string                 `json:"entityType"`
	EntityID   string                 `json:"entityId"`
	EventType  string                 `json:"eventType"`
	Level      model.ErrorLevel       `json:"level"`
	Message    string                 `json:"message"`
	Error      error                  `json:"-"` // Original error for stack trace extraction
	Metadata   map[string]interface{} `json:"metadata"`
	Component  string                 `json:"component"`
	UserID     string                 `json:"userId,omitempty"`
	SessionID  string                 `json:"sessionId,omitempty"`
}

// ErrorLogger provides non-blocking error logging functionality
type ErrorLogger struct {
	db       *gorm.DB
	logChan  chan LogEntry
	stopChan chan struct{}
	doneChan chan struct{}
}

// Global error logger instance
var Default *ErrorLogger

// Init initializes the global error logger with the provided database connection
func Init(db *gorm.DB) {
	Default = New(db)
	go Default.Start()
}

// New creates a new ErrorLogger instance
func New(db *gorm.DB) *ErrorLogger {
	return &ErrorLogger{
		db:       db,
		logChan:  make(chan LogEntry, 1000), // Buffered channel for non-blocking writes
		stopChan: make(chan struct{}),
		doneChan: make(chan struct{}),
	}
}

// Start begins processing log entries in a background goroutine
func (el *ErrorLogger) Start() {
	defer close(el.doneChan)

	for {
		select {
		case entry := <-el.logChan:
			el.persistEntry(entry)
		case <-el.stopChan:
			// Process remaining entries before shutting down
			for {
				select {
				case entry := <-el.logChan:
					el.persistEntry(entry)
				default:
					return
				}
			}
		}
	}
}

// Stop gracefully shuts down the error logger
func (el *ErrorLogger) Stop() {
	close(el.stopChan)
	<-el.doneChan
}

// Log adds an error log entry to the processing queue (non-blocking)
func (el *ErrorLogger) Log(entry LogEntry) {
	select {
	case el.logChan <- entry:
		// Successfully queued
	default:
		// Channel is full, log to system logger as fallback
		logger.Errorw("Error log channel full, dropping entry",
			"entityType", entry.EntityType,
			"entityId", entry.EntityID,
			"component", entry.Component,
			"message", entry.Message)
	}
}

// persistEntry saves a log entry to the database
func (el *ErrorLogger) persistEntry(entry LogEntry) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Convert metadata to JSON
	metadataJSON := model.ConfigMap{}
	if entry.Metadata != nil {
		for k, v := range entry.Metadata {
			metadataJSON[k] = fmt.Sprintf("%v", v)
		}
	}

	// Extract stack trace if error is provided
	stackTrace := ""
	if entry.Error != nil {
		stackTrace = extractStackTrace(entry.Error)
	}

	errorLog := model.ErrorLog{
		CreatedAt:  time.Now(),
		EntityType: entry.EntityType,
		EntityID:   entry.EntityID,
		EventType:  entry.EventType,
		Level:      entry.Level,
		Message:    entry.Message,
		StackTrace: stackTrace,
		Metadata:   metadataJSON,
		Component:  entry.Component,
		UserID:     entry.UserID,
		SessionID:  entry.SessionID,
	}

	err := el.db.WithContext(ctx).Create(&errorLog).Error
	if err != nil {
		logger.Errorw("Failed to persist error log entry", "error", err, "entry", entry)
	}
}

// extractStackTrace extracts stack trace information from an error
func extractStackTrace(err error) string {
	if err == nil {
		return ""
	}

	// Try to get stack trace from cockroachdb/errors
	if stackTracer, ok := err.(interface{ StackTrace() interface{} }); ok {
		return fmt.Sprintf("%+v", stackTracer.StackTrace())
	}

	// Fallback to runtime stack trace
	buf := make([]byte, 4096)
	n := runtime.Stack(buf, false)
	return string(buf[:n])
}

// Convenience functions for common logging scenarios

// LogInfo logs an informational message
func LogInfo(component, entityType, entityID, eventType, message string, metadata map[string]interface{}) {
	if Default != nil {
		Default.Log(LogEntry{
			EntityType: entityType,
			EntityID:   entityID,
			EventType:  eventType,
			Level:      model.ErrorLevelInfo,
			Message:    message,
			Metadata:   metadata,
			Component:  component,
		})
	}
}

// LogWarning logs a warning message
func LogWarning(component, entityType, entityID, eventType, message string, metadata map[string]interface{}) {
	if Default != nil {
		Default.Log(LogEntry{
			EntityType: entityType,
			EntityID:   entityID,
			EventType:  eventType,
			Level:      model.ErrorLevelWarning,
			Message:    message,
			Metadata:   metadata,
			Component:  component,
		})
	}
}

// LogError logs an error message with optional error for stack trace
func LogError(component, entityType, entityID, eventType, message string, err error, metadata map[string]interface{}) {
	if Default != nil {
		Default.Log(LogEntry{
			EntityType: entityType,
			EntityID:   entityID,
			EventType:  eventType,
			Level:      model.ErrorLevelError,
			Message:    message,
			Error:      err,
			Metadata:   metadata,
			Component:  component,
		})
	}
}

// LogCritical logs a critical error message
func LogCritical(component, entityType, entityID, eventType, message string, err error, metadata map[string]interface{}) {
	if Default != nil {
		Default.Log(LogEntry{
			EntityType: entityType,
			EntityID:   entityID,
			EventType:  eventType,
			Level:      model.ErrorLevelCritical,
			Message:    message,
			Error:      err,
			Metadata:   metadata,
			Component:  component,
		})
	}
}

// LogOnboardEvent logs events related to the onboard command
func LogOnboardEvent(level model.ErrorLevel, preparationID, eventType, message string, err error, metadata map[string]interface{}) {
	if Default != nil {
		Default.Log(LogEntry{
			EntityType: "preparation",
			EntityID:   preparationID,
			EventType:  eventType,
			Level:      level,
			Message:    message,
			Error:      err,
			Metadata:   metadata,
			Component:  "onboard",
		})
	}
}

// LogDealScheduleEvent logs events related to deal schedule creation
func LogDealScheduleEvent(level model.ErrorLevel, scheduleID, eventType, message string, err error, metadata map[string]interface{}) {
	if Default != nil {
		Default.Log(LogEntry{
			EntityType: "schedule",
			EntityID:   scheduleID,
			EventType:  eventType,
			Level:      level,
			Message:    message,
			Error:      err,
			Metadata:   metadata,
			Component:  "deal_schedule",
		})
	}
}

// QueryFilters represents filtering options for querying error logs
type QueryFilters struct {
	EntityType string           `form:"entity_type" query:"entity_type"`
	EntityID   string           `form:"entity_id" query:"entity_id"`
	Component  string           `form:"component" query:"component"`
	Level      model.ErrorLevel `form:"level" query:"level"`
	EventType  string           `form:"event_type" query:"event_type"`
	StartTime  *time.Time       `form:"start_time" query:"start_time"`
	EndTime    *time.Time       `form:"end_time" query:"end_time"`
	Limit      int              `form:"limit" query:"limit"`
	Offset     int              `form:"offset" query:"offset"`
}

// QueryErrorLogs queries error logs with filtering and pagination
func QueryErrorLogs(ctx context.Context, db *gorm.DB, filters QueryFilters) ([]model.ErrorLog, int64, error) {
	query := db.WithContext(ctx).Model(&model.ErrorLog{})

	// Apply filters
	if filters.EntityType != "" {
		query = query.Where("entity_type = ?", filters.EntityType)
	}
	if filters.EntityID != "" {
		query = query.Where("entity_id = ?", filters.EntityID)
	}
	if filters.Component != "" {
		query = query.Where("component = ?", filters.Component)
	}
	if filters.Level != "" {
		query = query.Where("level = ?", filters.Level)
	}
	if filters.EventType != "" {
		query = query.Where("event_type = ?", filters.EventType)
	}
	if filters.StartTime != nil {
		query = query.Where("created_at >= ?", *filters.StartTime)
	}
	if filters.EndTime != nil {
		query = query.Where("created_at <= ?", *filters.EndTime)
	}

	// Get total count
	var total int64
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, errors.WithStack(err)
	}

	// Apply pagination and ordering
	if filters.Limit <= 0 {
		filters.Limit = 50 // Default limit
	}
	if filters.Limit > 1000 {
		filters.Limit = 1000 // Max limit
	}

	var errorLogs []model.ErrorLog
	err = query.Order("created_at DESC").
		Limit(filters.Limit).
		Offset(filters.Offset).
		Find(&errorLogs).Error
	if err != nil {
		return nil, 0, errors.WithStack(err)
	}

	return errorLogs, total, nil
}
