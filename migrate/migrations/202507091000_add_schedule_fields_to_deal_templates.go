package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func _202507091000_add_schedule_fields_to_deal_templates() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "202507091000_add_schedule_fields_to_deal_templates",
		Migrate: func(db *gorm.DB) error {
			// Add scheduling fields
			err := db.Exec(`ALTER TABLE deal_templates ADD COLUMN template_schedule_cron VARCHAR(255) DEFAULT ''`).Error
			if err != nil {
				return err
			}

			err = db.Exec(`ALTER TABLE deal_templates ADD COLUMN template_schedule_deal_number INTEGER DEFAULT 0`).Error
			if err != nil {
				return err
			}

			err = db.Exec(`ALTER TABLE deal_templates ADD COLUMN template_schedule_deal_size VARCHAR(255) DEFAULT '0'`).Error
			if err != nil {
				return err
			}

			// Add restriction fields
			err = db.Exec(`ALTER TABLE deal_templates ADD COLUMN template_total_deal_number INTEGER DEFAULT 0`).Error
			if err != nil {
				return err
			}

			err = db.Exec(`ALTER TABLE deal_templates ADD COLUMN template_total_deal_size VARCHAR(255) DEFAULT '0'`).Error
			if err != nil {
				return err
			}

			err = db.Exec(`ALTER TABLE deal_templates ADD COLUMN template_max_pending_deal_number INTEGER DEFAULT 0`).Error
			if err != nil {
				return err
			}

			err = db.Exec(`ALTER TABLE deal_templates ADD COLUMN template_max_pending_deal_size VARCHAR(255) DEFAULT '0'`).Error
			if err != nil {
				return err
			}

			// Add HTTP headers as JSON array (to support []string format)
			err = db.Exec(`ALTER TABLE deal_templates ADD COLUMN template_http_headers JSON DEFAULT '[]'`).Error
			if err != nil {
				return err
			}

			return nil
		},
		Rollback: func(db *gorm.DB) error {
			// Remove the added fields
			fields := []string{
				"template_schedule_cron",
				"template_schedule_deal_number",
				"template_schedule_deal_size",
				"template_total_deal_number",
				"template_total_deal_size",
				"template_max_pending_deal_number",
				"template_max_pending_deal_size",
				"template_http_headers",
			}

			for _, field := range fields {
				err := db.Exec(`ALTER TABLE deal_templates DROP COLUMN IF EXISTS ` + field).Error
				if err != nil {
					return err
				}
			}

			return nil
		},
	}
}
