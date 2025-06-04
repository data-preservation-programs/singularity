package notification

import (
	"context"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/ipfs/go-log/v2"
	"gorm.io/gorm"
)

var logger = log.Logger("notification")

type NotificationType string

const (
	NotificationTypeInfo    NotificationType = "info"
	NotificationTypeWarning NotificationType = "warning"
	NotificationTypeError   NotificationType = "error"
)

type NotificationLevel string

const (
	NotificationLevelLow    NotificationLevel = "low"
	NotificationLevelMedium NotificationLevel = "medium"
	NotificationLevelHigh   NotificationLevel = "high"
)

type Handler struct{}

var Default = &Handler{}

type CreateNotificationRequest struct {
	Type        NotificationType  `json:"type"`
	Level       NotificationLevel `json:"level"`
	Title       string            `json:"title"`
	Message     string            `json:"message"`
	Source      string            `json:"source"`
	SourceID    string            `json:"sourceId,omitempty"`
	Metadata    model.ConfigMap   `json:"metadata,omitempty"`
	Acknowledged bool             `json:"acknowledged"`
}

// CreateNotification creates a new notification and saves it to the database
func (h *Handler) CreateNotification(ctx context.Context, db *gorm.DB, request CreateNotificationRequest) (*model.Notification, error) {
	notification := &model.Notification{
		Type:         string(request.Type),
		Level:        string(request.Level),
		Title:        request.Title,
		Message:      request.Message,
		Source:       request.Source,
		SourceID:     request.SourceID,
		Metadata:     request.Metadata,
		Acknowledged: request.Acknowledged,
		CreatedAt:    time.Now(),
	}

	if err := db.WithContext(ctx).Create(notification).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	// Log the notification for immediate visibility
	h.logNotification(notification)

	return notification, nil
}

// LogWarning creates and logs a warning notification
func (h *Handler) LogWarning(ctx context.Context, db *gorm.DB, source, title, message string, metadata ...model.ConfigMap) (*model.Notification, error) {
	var meta model.ConfigMap
	if len(metadata) > 0 {
		meta = metadata[0]
	}

	return h.CreateNotification(ctx, db, CreateNotificationRequest{
		Type:     NotificationTypeWarning,
		Level:    NotificationLevelMedium,
		Title:    title,
		Message:  message,
		Source:   source,
		Metadata: meta,
	})
}

// LogError creates and logs an error notification
func (h *Handler) LogError(ctx context.Context, db *gorm.DB, source, title, message string, metadata ...model.ConfigMap) (*model.Notification, error) {
	var meta model.ConfigMap
	if len(metadata) > 0 {
		meta = metadata[0]
	}

	return h.CreateNotification(ctx, db, CreateNotificationRequest{
		Type:     NotificationTypeError,
		Level:    NotificationLevelHigh,
		Title:    title,
		Message:  message,
		Source:   source,
		Metadata: meta,
	})
}

// LogInfo creates and logs an info notification
func (h *Handler) LogInfo(ctx context.Context, db *gorm.DB, source, title, message string, metadata ...model.ConfigMap) (*model.Notification, error) {
	var meta model.ConfigMap
	if len(metadata) > 0 {
		meta = metadata[0]
	}

	return h.CreateNotification(ctx, db, CreateNotificationRequest{
		Type:     NotificationTypeInfo,
		Level:    NotificationLevelLow,
		Title:    title,
		Message:  message,
		Source:   source,
		Metadata: meta,
	})
}

// ListNotifications retrieves notifications with pagination and filtering
func (h *Handler) ListNotifications(ctx context.Context, db *gorm.DB, offset, limit int, notificationType *NotificationType, acknowledged *bool) ([]*model.Notification, error) {
	var notifications []*model.Notification
	
	query := db.WithContext(ctx).Model(&model.Notification{})
	
	if notificationType != nil {
		query = query.Where("type = ?", string(*notificationType))
	}
	
	if acknowledged != nil {
		query = query.Where("acknowledged = ?", *acknowledged)
	}
	
	if err := query.Order("created_at DESC").Offset(offset).Limit(limit).Find(&notifications).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	
	return notifications, nil
}

// AcknowledgeNotification marks a notification as acknowledged
func (h *Handler) AcknowledgeNotification(ctx context.Context, db *gorm.DB, id uint) error {
	if err := db.WithContext(ctx).Model(&model.Notification{}).Where("id = ?", id).Update("acknowledged", true).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// GetNotificationByID retrieves a specific notification by ID
func (h *Handler) GetNotificationByID(ctx context.Context, db *gorm.DB, id uint) (*model.Notification, error) {
	var notification model.Notification
	if err := db.WithContext(ctx).First(&notification, id).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return &notification, nil
}

// DeleteNotification removes a notification from the database
func (h *Handler) DeleteNotification(ctx context.Context, db *gorm.DB, id uint) error {
	if err := db.WithContext(ctx).Delete(&model.Notification{}, id).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// logNotification logs the notification to the system logger
func (h *Handler) logNotification(notification *model.Notification) {
	logMsg := logger.With("source", notification.Source, "title", notification.Title)
	
	switch notification.Type {
	case string(NotificationTypeError):
		logMsg.Errorf("[%s] %s: %s", notification.Source, notification.Title, notification.Message)
	case string(NotificationTypeWarning):
		logMsg.Warnf("[%s] %s: %s", notification.Source, notification.Title, notification.Message)
	case string(NotificationTypeInfo):
		logMsg.Infof("[%s] %s: %s", notification.Source, notification.Title, notification.Message)
	default:
		logMsg.Infof("[%s] %s: %s", notification.Source, notification.Title, notification.Message)
	}
}