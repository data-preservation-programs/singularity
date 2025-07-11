package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

// _202507090915_add_not_null_defaults adds NOT NULL DEFAULT constraints to deal template fields
func _202507090915_add_not_null_defaults() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "202507090915_add_not_null_defaults",
		Migrate: func(db *gorm.DB) error {
			// Add NOT NULL DEFAULT constraints to string fields in deal_templates table
			// Update existing NULL values to empty string first, then add NOT NULL constraint

			// Update NULL values to empty string for string fields
			err := db.Exec(`UPDATE deal_templates SET template_deal_provider = '' WHERE template_deal_provider IS NULL`).Error
			if err != nil {
				return err
			}

			err = db.Exec(`UPDATE deal_templates SET template_deal_template = '' WHERE template_deal_template IS NULL`).Error
			if err != nil {
				return err
			}

			err = db.Exec(`UPDATE deal_templates SET template_deal_http_headers = '' WHERE template_deal_http_headers IS NULL`).Error
			if err != nil {
				return err
			}

			err = db.Exec(`UPDATE deal_templates SET template_deal_url_template = '' WHERE template_deal_url_template IS NULL`).Error
			if err != nil {
				return err
			}

			err = db.Exec(`UPDATE deal_templates SET description = '' WHERE description IS NULL`).Error
			if err != nil {
				return err
			}

			// Add NOT NULL DEFAULT constraints for string fields
			err = db.Exec(`ALTER TABLE deal_templates ALTER COLUMN template_deal_provider SET NOT NULL`).Error
			if err != nil {
				return err
			}

			err = db.Exec(`ALTER TABLE deal_templates ALTER COLUMN template_deal_provider SET DEFAULT ''`).Error
			if err != nil {
				return err
			}

			err = db.Exec(`ALTER TABLE deal_templates ALTER COLUMN template_deal_template SET NOT NULL`).Error
			if err != nil {
				return err
			}

			err = db.Exec(`ALTER TABLE deal_templates ALTER COLUMN template_deal_template SET DEFAULT ''`).Error
			if err != nil {
				return err
			}

			err = db.Exec(`ALTER TABLE deal_templates ALTER COLUMN template_deal_http_headers SET NOT NULL`).Error
			if err != nil {
				return err
			}

			err = db.Exec(`ALTER TABLE deal_templates ALTER COLUMN template_deal_http_headers SET DEFAULT ''`).Error
			if err != nil {
				return err
			}

			err = db.Exec(`ALTER TABLE deal_templates ALTER COLUMN template_deal_url_template SET NOT NULL`).Error
			if err != nil {
				return err
			}

			err = db.Exec(`ALTER TABLE deal_templates ALTER COLUMN template_deal_url_template SET DEFAULT ''`).Error
			if err != nil {
				return err
			}

			err = db.Exec(`ALTER TABLE deal_templates ALTER COLUMN description SET NOT NULL`).Error
			if err != nil {
				return err
			}

			err = db.Exec(`ALTER TABLE deal_templates ALTER COLUMN description SET DEFAULT ''`).Error
			if err != nil {
				return err
			}

			// Add NOT NULL constraint to name field (already unique, but should be explicit)
			err = db.Exec(`ALTER TABLE deal_templates ALTER COLUMN name SET NOT NULL`).Error
			if err != nil {
				return err
			}

			return nil
		},
		Rollback: func(db *gorm.DB) error {
			// Remove NOT NULL constraints (making columns nullable again)
			err := db.Exec(`ALTER TABLE deal_templates ALTER COLUMN template_deal_provider DROP NOT NULL`).Error
			if err != nil {
				return err
			}

			err = db.Exec(`ALTER TABLE deal_templates ALTER COLUMN template_deal_template DROP NOT NULL`).Error
			if err != nil {
				return err
			}

			err = db.Exec(`ALTER TABLE deal_templates ALTER COLUMN template_deal_http_headers DROP NOT NULL`).Error
			if err != nil {
				return err
			}

			err = db.Exec(`ALTER TABLE deal_templates ALTER COLUMN template_deal_url_template DROP NOT NULL`).Error
			if err != nil {
				return err
			}

			err = db.Exec(`ALTER TABLE deal_templates ALTER COLUMN description DROP NOT NULL`).Error
			if err != nil {
				return err
			}

			err = db.Exec(`ALTER TABLE deal_templates ALTER COLUMN name DROP NOT NULL`).Error
			if err != nil {
				return err
			}

			return nil
		},
	}
}
