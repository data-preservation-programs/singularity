package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func _202507090900_add_missing_deal_template_fields() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "202507090900_add_missing_deal_template_fields",
		Migrate: func(db *gorm.DB) error {
			// Add the missing fields to the deal_templates table
			err := db.Exec(`ALTER TABLE deal_templates ADD COLUMN template_deal_notes TEXT DEFAULT ''`).Error
			if err != nil {
				return err
			}

			err = db.Exec(`ALTER TABLE deal_templates ADD COLUMN template_deal_force BOOLEAN DEFAULT false`).Error
			if err != nil {
				return err
			}

			err = db.Exec(`ALTER TABLE deal_templates ADD COLUMN template_deal_allowed_piece_cids JSON DEFAULT '[]'`).Error
			if err != nil {
				return err
			}

			return nil
		},
		Rollback: func(db *gorm.DB) error {
			// Remove the added fields
			err := db.Exec(`ALTER TABLE deal_templates DROP COLUMN template_deal_notes`).Error
			if err != nil {
				return err
			}

			err = db.Exec(`ALTER TABLE deal_templates DROP COLUMN template_deal_force`).Error
			if err != nil {
				return err
			}

			err = db.Exec(`ALTER TABLE deal_templates DROP COLUMN template_deal_allowed_piece_cids`).Error
			if err != nil {
				return err
			}

			return nil
		},
	}
}
