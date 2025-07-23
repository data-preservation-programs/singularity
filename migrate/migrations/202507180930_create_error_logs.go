package migrations

import (
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

// ErrorLog represents error logging entries for debugging and monitoring
type ErrorLog struct {
	ID         uint      `gorm:"primaryKey"                         json:"id"`
	CreatedAt  time.Time `json:"createdAt"                          table:"format:2006-01-02 15:04:05"`
	EntityType string    `gorm:"index:idx_entity_type_id"           json:"entityType"` // Type of entity (deal, preparation, schedule, etc.)
	EntityID   string    `gorm:"index:idx_entity_type_id;size:255"  json:"entityId"`   // ID of the associated entity
	EventType  string    `gorm:"index"                              json:"eventType"`  // Type of event (creation, processing, error, etc.)
	Level      string    `gorm:"index"                              json:"level"`      // Error level (info, warning, error, critical)
	Message    string    `gorm:"type:text"                          json:"message"`    // Human-readable error message
	StackTrace string    `gorm:"type:text"                          json:"stackTrace"` // Stack trace if available
	Metadata   string    `gorm:"type:json"                          json:"metadata"`   // Additional context as JSON
	Component  string    `gorm:"index"                              json:"component"`  // Component that generated the error (onboard, deal_schedule, etc.)
	UserID     string    `gorm:"index;size:255"                     json:"userId"`     // Optional user identifier
	SessionID  string    `gorm:"index;size:255"                     json:"sessionId"`  // Optional session identifier
}

// Create migration for error logs table
func _202507180930_create_error_logs() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "202507180930",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&ErrorLog{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable(&ErrorLog{})
		},
	}
}
