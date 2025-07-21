package migrations

import (
	"strings"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

// Fix foreign key constraint issue between deals and deal_state_changes tables
func _202507210939_fix_deal_state_changes_fk() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "202507210939",
		Migrate: func(tx *gorm.DB) error {
			// Get the database dialect
			dialect := tx.Dialector.Name()

			// First, drop any incorrect foreign key constraints
			// Check multiple possible constraint names that might have been created
			incorrectConstraints := []string{
				"fk_deal_state_changes_deal",
				"fk_deals_deal_state_change",
				"fk_deals_deal_id",
			}

			for _, constraintName := range incorrectConstraints {
				if tx.Migrator().HasConstraint("deals", constraintName) {
					if err := tx.Migrator().DropConstraint("deals", constraintName); err != nil {
						// Log but continue if we can't drop
						tx.Logger.Info(tx.Statement.Context, "Could not drop constraint", constraintName, err)
					}
				}
			}

			// For MySQL, we need to check information schema for constraints
			if dialect == "mysql" {
				var constraints []struct {
					ConstraintName string `gorm:"column:CONSTRAINT_NAME"`
				}

				// Find any foreign key constraints on deals table that reference deal_state_changes
				tx.Raw(`
					SELECT CONSTRAINT_NAME 
					FROM information_schema.KEY_COLUMN_USAGE 
					WHERE TABLE_NAME = 'deals' 
					AND REFERENCED_TABLE_NAME = 'deal_state_changes'
					AND TABLE_SCHEMA = DATABASE()
				`).Scan(&constraints)

				for _, c := range constraints {
					if err := tx.Exec("ALTER TABLE deals DROP FOREIGN KEY " + c.ConstraintName).Error; err != nil {
						tx.Logger.Info(tx.Statement.Context, "Could not drop MySQL constraint", c.ConstraintName, err)
					}
				}
			}

			// Ensure the correct foreign key exists on deal_state_changes table
			if !tx.Migrator().HasConstraint("deal_state_changes", "fk_deal_state_changes_deal") {
				// Add the foreign key constraint with raw SQL to ensure it's created correctly
				var sql string
				switch dialect {
				case "mysql":
					sql = `
						ALTER TABLE deal_state_changes 
						ADD CONSTRAINT fk_deal_state_changes_deal 
						FOREIGN KEY (deal_id) REFERENCES deals(id) 
						ON DELETE CASCADE ON UPDATE CASCADE
					`
				case "postgres":
					sql = `
						ALTER TABLE deal_state_changes 
						ADD CONSTRAINT fk_deal_state_changes_deal 
						FOREIGN KEY (deal_id) REFERENCES deals(id) 
						ON DELETE CASCADE ON UPDATE CASCADE
					`
				case "sqlite":
					// SQLite doesn't support adding foreign keys after table creation
					// It would need a table recreation which is complex
					return nil
				default:
					sql = `
						ALTER TABLE deal_state_changes 
						ADD CONSTRAINT fk_deal_state_changes_deal 
						FOREIGN KEY (deal_id) REFERENCES deals(id) 
						ON DELETE CASCADE ON UPDATE CASCADE
					`
				}

				if sql != "" {
					if err := tx.Exec(sql).Error; err != nil {
						// Check if it's a duplicate constraint error
						if !strings.Contains(strings.ToLower(err.Error()), "duplicate") &&
							!strings.Contains(strings.ToLower(err.Error()), "already exists") {
							return err
						}
					}
				}
			}

			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			// No rollback needed as we're just fixing constraints
			return nil
		},
	}
}
