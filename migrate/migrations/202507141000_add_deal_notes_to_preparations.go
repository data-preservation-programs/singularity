package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

// _202507141000_add_deal_notes_to_preparations adds the missing deal_config_deal_notes column to the preparations table
func _202507141000_add_deal_notes_to_preparations() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "202507141000_add_deal_notes_to_preparations",
		Migrate: func(db *gorm.DB) error {
			// Add the missing deal_config_deal_notes column to the preparations table
			err := db.Exec(`ALTER TABLE preparations ADD COLUMN deal_config_deal_notes TEXT DEFAULT ''`).Error
			if err != nil {
				return err
			}

			// Add the missing deal_config_deal_force column to the preparations table
			err = db.Exec(`ALTER TABLE preparations ADD COLUMN deal_config_deal_force BOOLEAN DEFAULT false`).Error
			if err != nil {
				return err
			}

			// Add the missing deal_config_deal_allowed_piece_cids column to the preparations table
			err = db.Exec(`ALTER TABLE preparations ADD COLUMN deal_config_deal_allowed_piece_cids JSON DEFAULT '[]'`).Error
			if err != nil {
				return err
			}

			// Add the missing deal_config_schedule_cron column to the preparations table
			err = db.Exec(`ALTER TABLE preparations ADD COLUMN deal_config_schedule_cron VARCHAR(255) DEFAULT ''`).Error
			if err != nil {
				return err
			}

			// Add the missing deal_config_schedule_deal_number column to the preparations table
			err = db.Exec(`ALTER TABLE preparations ADD COLUMN deal_config_schedule_deal_number INTEGER DEFAULT 0`).Error
			if err != nil {
				return err
			}

			// Add the missing deal_config_schedule_deal_size column to the preparations table
			err = db.Exec(`ALTER TABLE preparations ADD COLUMN deal_config_schedule_deal_size VARCHAR(255) DEFAULT '0'`).Error
			if err != nil {
				return err
			}

			// Add the missing deal_config_total_deal_number column to the preparations table
			err = db.Exec(`ALTER TABLE preparations ADD COLUMN deal_config_total_deal_number INTEGER DEFAULT 0`).Error
			if err != nil {
				return err
			}

			// Add the missing deal_config_total_deal_size column to the preparations table
			err = db.Exec(`ALTER TABLE preparations ADD COLUMN deal_config_total_deal_size VARCHAR(255) DEFAULT '0'`).Error
			if err != nil {
				return err
			}

			// Add the missing deal_config_max_pending_deal_number column to the preparations table
			err = db.Exec(`ALTER TABLE preparations ADD COLUMN deal_config_max_pending_deal_number INTEGER DEFAULT 0`).Error
			if err != nil {
				return err
			}

			// Add the missing deal_config_max_pending_deal_size column to the preparations table
			err = db.Exec(`ALTER TABLE preparations ADD COLUMN deal_config_max_pending_deal_size VARCHAR(255) DEFAULT '0'`).Error
			if err != nil {
				return err
			}

			// Add the missing deal_config_http_headers column to the preparations table
			err = db.Exec(`ALTER TABLE preparations ADD COLUMN deal_config_http_headers JSON DEFAULT '[]'`).Error
			if err != nil {
				return err
			}

			return nil
		},
		Rollback: func(db *gorm.DB) error {
			// Remove the added columns
			err := db.Exec(`ALTER TABLE preparations DROP COLUMN deal_config_deal_notes`).Error
			if err != nil {
				return err
			}

			err = db.Exec(`ALTER TABLE preparations DROP COLUMN deal_config_deal_force`).Error
			if err != nil {
				return err
			}

			err = db.Exec(`ALTER TABLE preparations DROP COLUMN deal_config_deal_allowed_piece_cids`).Error
			if err != nil {
				return err
			}

			err = db.Exec(`ALTER TABLE preparations DROP COLUMN deal_config_schedule_cron`).Error
			if err != nil {
				return err
			}

			err = db.Exec(`ALTER TABLE preparations DROP COLUMN deal_config_schedule_deal_number`).Error
			if err != nil {
				return err
			}

			err = db.Exec(`ALTER TABLE preparations DROP COLUMN deal_config_schedule_deal_size`).Error
			if err != nil {
				return err
			}

			err = db.Exec(`ALTER TABLE preparations DROP COLUMN deal_config_total_deal_number`).Error
			if err != nil {
				return err
			}

			err = db.Exec(`ALTER TABLE preparations DROP COLUMN deal_config_total_deal_size`).Error
			if err != nil {
				return err
			}

			err = db.Exec(`ALTER TABLE preparations DROP COLUMN deal_config_max_pending_deal_number`).Error
			if err != nil {
				return err
			}

			err = db.Exec(`ALTER TABLE preparations DROP COLUMN deal_config_max_pending_deal_size`).Error
			if err != nil {
				return err
			}

			err = db.Exec(`ALTER TABLE preparations DROP COLUMN deal_config_http_headers`).Error
			if err != nil {
				return err
			}

			return nil
		},
	}
}
