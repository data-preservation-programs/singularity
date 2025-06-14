package notification

import (
	"context"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err)

	err = db.AutoMigrate(&model.Notification{})
	require.NoError(t, err)

	return db
}

func TestCreateNotification(t *testing.T) {
	db := setupTestDB(t)
	handler := &Handler{}
	ctx := context.Background()

	request := CreateNotificationRequest{
		Type:    NotificationTypeInfo,
		Level:   NotificationLevelLow,
		Title:   "Test Notification",
		Message: "This is a test notification",
		Source:  "test-handler",
		Metadata: model.ConfigMap{
			"test_key": "test_value",
		},
	}

	notification, err := handler.CreateNotification(ctx, db, request)
	require.NoError(t, err)
	require.NotNil(t, notification)
	require.Equal(t, string(NotificationTypeInfo), notification.Type)
	require.Equal(t, string(NotificationLevelLow), notification.Level)
	require.Equal(t, "Test Notification", notification.Title)
	require.Equal(t, "This is a test notification", notification.Message)
	require.Equal(t, "test-handler", notification.Source)
	require.Equal(t, "test_value", notification.Metadata["test_key"])
	require.False(t, notification.Acknowledged)
	require.NotZero(t, notification.ID)
}

func TestLogWarning(t *testing.T) {
	db := setupTestDB(t)
	handler := &Handler{}
	ctx := context.Background()

	metadata := model.ConfigMap{
		"preparation_id": "123",
		"wallet_id":      "456",
	}

	notification, err := handler.LogWarning(ctx, db, "wallet-validator", "Insufficient Balance", "Wallet does not have enough FIL for deal", metadata)
	require.NoError(t, err)
	require.NotNil(t, notification)
	require.Equal(t, string(NotificationTypeWarning), notification.Type)
	require.Equal(t, string(NotificationLevelMedium), notification.Level)
	require.Equal(t, "Insufficient Balance", notification.Title)
	require.Equal(t, "wallet-validator", notification.Source)
	require.Equal(t, metadata, notification.Metadata)
}

func TestLogError(t *testing.T) {
	db := setupTestDB(t)
	handler := &Handler{}
	ctx := context.Background()

	notification, err := handler.LogError(ctx, db, "sp-validator", "Storage Provider Unreachable", "Failed to connect to storage provider")
	require.NoError(t, err)
	require.NotNil(t, notification)
	require.Equal(t, string(NotificationTypeError), notification.Type)
	require.Equal(t, string(NotificationLevelHigh), notification.Level)
	require.Equal(t, "Storage Provider Unreachable", notification.Title)
	require.Equal(t, "sp-validator", notification.Source)
}

func TestLogInfo(t *testing.T) {
	db := setupTestDB(t)
	handler := &Handler{}
	ctx := context.Background()

	notification, err := handler.LogInfo(ctx, db, "prep-handler", "Preparation Created", "New preparation created successfully")
	require.NoError(t, err)
	require.NotNil(t, notification)
	require.Equal(t, string(NotificationTypeInfo), notification.Type)
	require.Equal(t, string(NotificationLevelLow), notification.Level)
	require.Equal(t, "Preparation Created", notification.Title)
	require.Equal(t, "prep-handler", notification.Source)
}

func TestListNotifications(t *testing.T) {
	db := setupTestDB(t)
	handler := &Handler{}
	ctx := context.Background()

	// Create test notifications
	_, err := handler.LogInfo(ctx, db, "test", "Info 1", "First info message")
	require.NoError(t, err)

	_, err = handler.LogWarning(ctx, db, "test", "Warning 1", "First warning message")
	require.NoError(t, err)

	_, err = handler.LogError(ctx, db, "test", "Error 1", "First error message")
	require.NoError(t, err)

	// Test list all notifications
	notifications, err := handler.ListNotifications(ctx, db, 0, 10, nil, nil)
	require.NoError(t, err)
	require.Len(t, notifications, 3)

	// Test filter by type
	warningType := NotificationTypeWarning
	notifications, err = handler.ListNotifications(ctx, db, 0, 10, &warningType, nil)
	require.NoError(t, err)
	require.Len(t, notifications, 1)
	require.Equal(t, string(NotificationTypeWarning), notifications[0].Type)

	// Test filter by acknowledged status
	acknowledged := false
	notifications, err = handler.ListNotifications(ctx, db, 0, 10, nil, &acknowledged)
	require.NoError(t, err)
	require.Len(t, notifications, 3)
	for _, n := range notifications {
		require.False(t, n.Acknowledged)
	}

	// Test pagination
	notifications, err = handler.ListNotifications(ctx, db, 0, 2, nil, nil)
	require.NoError(t, err)
	require.Len(t, notifications, 2)

	notifications, err = handler.ListNotifications(ctx, db, 2, 10, nil, nil)
	require.NoError(t, err)
	require.Len(t, notifications, 1)
}

func TestAcknowledgeNotification(t *testing.T) {
	db := setupTestDB(t)
	handler := &Handler{}
	ctx := context.Background()

	// Create a notification
	notification, err := handler.LogWarning(ctx, db, "test", "Test Warning", "Test message")
	require.NoError(t, err)
	require.False(t, notification.Acknowledged)

	// Acknowledge it
	err = handler.AcknowledgeNotification(ctx, db, notification.ID)
	require.NoError(t, err)

	// Verify it's acknowledged
	updated, err := handler.GetNotificationByID(ctx, db, notification.ID)
	require.NoError(t, err)
	require.True(t, updated.Acknowledged)
}

func TestGetNotificationByID(t *testing.T) {
	db := setupTestDB(t)
	handler := &Handler{}
	ctx := context.Background()

	// Create a notification
	original, err := handler.LogInfo(ctx, db, "test", "Test Info", "Test message")
	require.NoError(t, err)

	// Retrieve it by ID
	retrieved, err := handler.GetNotificationByID(ctx, db, original.ID)
	require.NoError(t, err)
	require.Equal(t, original.ID, retrieved.ID)
	require.Equal(t, original.Title, retrieved.Title)
	require.Equal(t, original.Message, retrieved.Message)
	require.Equal(t, original.Source, retrieved.Source)
}

func TestDeleteNotification(t *testing.T) {
	db := setupTestDB(t)
	handler := &Handler{}
	ctx := context.Background()

	// Create a notification
	notification, err := handler.LogError(ctx, db, "test", "Test Error", "Test message")
	require.NoError(t, err)

	// Delete it
	err = handler.DeleteNotification(ctx, db, notification.ID)
	require.NoError(t, err)

	// Verify it's gone
	_, err = handler.GetNotificationByID(ctx, db, notification.ID)
	require.Error(t, err)
}

func TestCreateNotificationWithoutMetadata(t *testing.T) {
	db := setupTestDB(t)
	handler := &Handler{}
	ctx := context.Background()

	// Test logging without metadata
	notification, err := handler.LogInfo(ctx, db, "test", "Simple Info", "Simple message")
	require.NoError(t, err)
	require.NotNil(t, notification)
	require.Nil(t, notification.Metadata)
}

func TestNotificationTimestamp(t *testing.T) {
	db := setupTestDB(t)
	handler := &Handler{}
	ctx := context.Background()

	before := time.Now()
	notification, err := handler.LogInfo(ctx, db, "test", "Timestamp Test", "Testing timestamp")
	require.NoError(t, err)
	after := time.Now()

	require.True(t, notification.CreatedAt.After(before) || notification.CreatedAt.Equal(before))
	require.True(t, notification.CreatedAt.Before(after) || notification.CreatedAt.Equal(after))
}
