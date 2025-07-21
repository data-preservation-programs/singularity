package migrations

import (
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

// DealStateChange represents a state transition in the deal lifecycle
// This table tracks all changes to deal states for comprehensive auditing
type DealStateChange struct {
	ID            uint64    `gorm:"primaryKey;autoIncrement"                  json:"id"`
	DealID        uint64    `gorm:"index;not null"                           json:"dealId"`        // Deal ID (internal singularity ID)
	PreviousState string    `gorm:"size:50;index"                            json:"previousState"` // Previous deal state
	NewState      string    `gorm:"size:50;index;not null"                   json:"newState"`      // New deal state
	Timestamp     time.Time `gorm:"index;not null" json:"timestamp"`                               // When the state change occurred
	EpochHeight   *int32    `gorm:"index"                                    json:"epochHeight"`   // Filecoin epoch when change occurred
	SectorID      *string   `gorm:"size:100;index"                           json:"sectorId"`      // Storage provider sector ID
	ProviderID    string    `gorm:"size:20;index"                            json:"providerId"`    // Storage provider ID
	ClientAddress string    `gorm:"size:86;index"                            json:"clientAddress"` // Client wallet address
	Metadata      string    `gorm:"type:TEXT"                                json:"metadata"`      // Additional metadata as JSON
}

// Create migration for deal_state_changes table
func _202507180900_create_deal_state_changes() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "202507180900",
		Migrate: func(tx *gorm.DB) error {
			// Create the deal_state_changes table
			err := tx.AutoMigrate(&DealStateChange{})
			if err != nil {
				return err
			}

			// Add foreign key constraint from deal_state_changes.deal_id to deals.id
			return tx.Exec("ALTER TABLE deal_state_changes ADD CONSTRAINT fk_deal_state_changes_deal FOREIGN KEY (deal_id) REFERENCES deals(id) ON DELETE CASCADE").Error
		},
		Rollback: func(tx *gorm.DB) error {
			// Drop foreign key constraint first if it exists
			tx.Exec("ALTER TABLE deal_state_changes DROP FOREIGN KEY IF EXISTS fk_deal_state_changes_deal")
			return tx.Migrator().DropTable(&DealStateChange{})
		},
	}
}
