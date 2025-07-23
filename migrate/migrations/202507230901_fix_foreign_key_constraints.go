package migrations

import (
	"strings"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

// Fix foreign key constraint issues that prevent deal insertions
func _202507230901_fix_foreign_key_constraints() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "202507230901",
		Migrate: func(tx *gorm.DB) error {
			// Fix the backwards foreign key constraint that's preventing deal insertions
			// This runs after all model AutoMigrate calls, so it fixes any wrong constraints GORM created
			dialect := tx.Dialector.Name()

			if dialect == "sqlite" {
				// Disable foreign keys temporarily for all operations
				tx.Exec("PRAGMA foreign_keys = OFF")
				defer tx.Exec("PRAGMA foreign_keys = ON")

				// Check both tables for problematic constraints
				var dealsInfo []struct {
					SQL string `gorm:"column:sql"`
				}
				tx.Raw("SELECT sql FROM sqlite_master WHERE type='table' AND name='deals'").Scan(&dealsInfo)

				// Fix deals table if it has wrong constraint
				for _, info := range dealsInfo {
					if strings.Contains(info.SQL, "fk_deal_state_changes_deal") ||
						strings.Contains(info.SQL, "REFERENCES deal_state_changes") {
						// Recreate deals table without the incorrect constraint

						// Backup existing data
						tx.Exec("CREATE TEMPORARY TABLE deals_backup AS SELECT * FROM deals")

						// Drop the problematic table
						if err := tx.Migrator().DropTable("deals"); err != nil {
							return err
						}

						// Create clean deals table
						createDealsSQL := `
							CREATE TABLE deals (
								id INTEGER PRIMARY KEY AUTOINCREMENT,
								created_at DATETIME,
								updated_at DATETIME,
								last_verified_at DATETIME,
								deal_id INTEGER UNIQUE,
								state TEXT,
								client_actor_id TEXT,
								provider TEXT,
								proposal_id TEXT,
								label TEXT,
								piece_cid BLOB,
								piece_size INTEGER,
								start_epoch INTEGER,
								end_epoch INTEGER,
								sector_start_epoch INTEGER,
								price TEXT,
								verified BOOLEAN,
								error_message TEXT,
								schedule_id INTEGER,
								client_id INTEGER,
								
								CONSTRAINT fk_deals_schedule FOREIGN KEY (schedule_id) REFERENCES schedules(id) ON DELETE SET NULL,
								CONSTRAINT fk_deals_wallet FOREIGN KEY (client_id) REFERENCES wallets(id) ON DELETE SET NULL
							)
						`

						tx.Exec(createDealsSQL)
						tx.Exec("CREATE INDEX idx_deals_pending ON deals(state, client_id)")
						tx.Exec("CREATE UNIQUE INDEX uni_deals_deal_id ON deals(deal_id)")
						tx.Exec("INSERT INTO deals SELECT * FROM deals_backup")
						tx.Exec("DROP TABLE deals_backup")

						break
					}
				}

				// Always ensure deal_state_changes has the correct constraint
				var stateChangesInfo []struct {
					SQL string `gorm:"column:sql"`
				}
				tx.Raw("SELECT sql FROM sqlite_master WHERE type='table' AND name='deal_state_changes'").Scan(&stateChangesInfo)

				needsConstraint := true
				for _, info := range stateChangesInfo {
					if strings.Contains(info.SQL, "REFERENCES deals") {
						needsConstraint = false
						break
					}
				}

				if needsConstraint {
					// Recreate deal_state_changes with correct constraint
					tx.Exec("CREATE TEMPORARY TABLE deal_state_changes_backup AS SELECT * FROM deal_state_changes")
					if err := tx.Migrator().DropTable("deal_state_changes"); err != nil {
						return err
					}

					createStateChangesSQL := `
						CREATE TABLE deal_state_changes (
							id INTEGER PRIMARY KEY AUTOINCREMENT,
							deal_id INTEGER NOT NULL,
							previous_state TEXT,
							new_state TEXT NOT NULL,
							timestamp DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
							epoch_height INTEGER,
							sector_id TEXT,
							provider_id TEXT NOT NULL,
							client_address TEXT NOT NULL,
							metadata TEXT,
							
							FOREIGN KEY (deal_id) REFERENCES deals(id) ON DELETE CASCADE ON UPDATE CASCADE
						)
					`

					tx.Exec(createStateChangesSQL)

					// Create all the necessary indexes
					tx.Exec("CREATE INDEX idx_deal_state_changes_deal_id ON deal_state_changes(deal_id)")
					tx.Exec("CREATE INDEX idx_deal_state_changes_previous_state ON deal_state_changes(previous_state)")
					tx.Exec("CREATE INDEX idx_deal_state_changes_new_state ON deal_state_changes(new_state)")
					tx.Exec("CREATE INDEX idx_deal_state_changes_timestamp ON deal_state_changes(timestamp)")
					tx.Exec("CREATE INDEX idx_deal_state_changes_epoch_height ON deal_state_changes(epoch_height)")
					tx.Exec("CREATE INDEX idx_deal_state_changes_sector_id ON deal_state_changes(sector_id)")
					tx.Exec("CREATE INDEX idx_deal_state_changes_provider_id ON deal_state_changes(provider_id)")
					tx.Exec("CREATE INDEX idx_deal_state_changes_client_address ON deal_state_changes(client_address)")

					tx.Exec("INSERT INTO deal_state_changes SELECT * FROM deal_state_changes_backup")
					tx.Exec("DROP TABLE deal_state_changes_backup")
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
