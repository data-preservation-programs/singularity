package migrations

import (
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

// _202506240815_create_notifications creates the notifications table
func _202506240815_create_notifications() *gormigrate.Migration {
	type ConfigMap map[string]string

	type Notification struct {
		ID           uint      `gorm:"primaryKey"`
		CreatedAt    time.Time
		Type         string    // info, warning, error
		Level        string    // low, medium, high
		Title        string
		Message      string
		Source       string   // Component that generated the notification
		SourceID     string   // Optional ID of the source entity
		Metadata     ConfigMap `gorm:"type:JSON"`
		Acknowledged bool
	}

	return &gormigrate.Migration{
		ID: "202506240815",
		Migrate: func(tx *gorm.DB) error {
			return tx.Migrator().CreateTable(&Notification{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("notifications")
		},
	}
}