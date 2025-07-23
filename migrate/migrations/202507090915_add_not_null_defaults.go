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
			// Update existing NULL values to empty string first
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

			// Handle database-specific syntax for adding NOT NULL constraints
			dialectName := db.Dialector.Name()

			switch dialectName {
			case "sqlite":
				// SQLite doesn't support ALTER COLUMN directly, so we skip the NOT NULL constraints
				// The GORM model tags will enforce this in the application layer
				return nil
			case "postgres":
				// PostgreSQL syntax
				queries := []string{
					`ALTER TABLE deal_templates ALTER COLUMN template_deal_provider SET NOT NULL`,
					`ALTER TABLE deal_templates ALTER COLUMN template_deal_provider SET DEFAULT ''`,
					`ALTER TABLE deal_templates ALTER COLUMN template_deal_template SET NOT NULL`,
					`ALTER TABLE deal_templates ALTER COLUMN template_deal_template SET DEFAULT ''`,
					`ALTER TABLE deal_templates ALTER COLUMN template_deal_http_headers SET NOT NULL`,
					`ALTER TABLE deal_templates ALTER COLUMN template_deal_http_headers SET DEFAULT ''`,
					`ALTER TABLE deal_templates ALTER COLUMN template_deal_url_template SET NOT NULL`,
					`ALTER TABLE deal_templates ALTER COLUMN template_deal_url_template SET DEFAULT ''`,
					`ALTER TABLE deal_templates ALTER COLUMN description SET NOT NULL`,
					`ALTER TABLE deal_templates ALTER COLUMN description SET DEFAULT ''`,
					`ALTER TABLE deal_templates ALTER COLUMN name SET NOT NULL`,
				}
				for _, query := range queries {
					if err := db.Exec(query).Error; err != nil {
						return err
					}
				}
			case "mysql":
				// MySQL syntax
				queries := []string{
					`ALTER TABLE deal_templates MODIFY template_deal_provider VARCHAR(255) NOT NULL DEFAULT ''`,
					`ALTER TABLE deal_templates MODIFY template_deal_template VARCHAR(255) NOT NULL DEFAULT ''`,
					`ALTER TABLE deal_templates MODIFY template_deal_http_headers TEXT NOT NULL DEFAULT ''`,
					`ALTER TABLE deal_templates MODIFY template_deal_url_template TEXT NOT NULL DEFAULT ''`,
					`ALTER TABLE deal_templates MODIFY description TEXT NOT NULL DEFAULT ''`,
					`ALTER TABLE deal_templates MODIFY name VARCHAR(255) NOT NULL`,
				}
				for _, query := range queries {
					if err := db.Exec(query).Error; err != nil {
						return err
					}
				}
			}

			return nil
		},
		Rollback: func(db *gorm.DB) error {
			// Handle database-specific syntax for removing NOT NULL constraints
			dialectName := db.Dialector.Name()

			switch dialectName {
			case "sqlite":
				// SQLite doesn't support ALTER COLUMN directly, so nothing to rollback
				return nil
			case "postgres":
				// PostgreSQL syntax
				queries := []string{
					`ALTER TABLE deal_templates ALTER COLUMN template_deal_provider DROP NOT NULL`,
					`ALTER TABLE deal_templates ALTER COLUMN template_deal_template DROP NOT NULL`,
					`ALTER TABLE deal_templates ALTER COLUMN template_deal_http_headers DROP NOT NULL`,
					`ALTER TABLE deal_templates ALTER COLUMN template_deal_url_template DROP NOT NULL`,
					`ALTER TABLE deal_templates ALTER COLUMN description DROP NOT NULL`,
					`ALTER TABLE deal_templates ALTER COLUMN name DROP NOT NULL`,
				}
				for _, query := range queries {
					if err := db.Exec(query).Error; err != nil {
						return err
					}
				}
			case "mysql":
				// MySQL syntax - make columns nullable again
				queries := []string{
					`ALTER TABLE deal_templates MODIFY template_deal_provider VARCHAR(255) NULL`,
					`ALTER TABLE deal_templates MODIFY template_deal_template VARCHAR(255) NULL`,
					`ALTER TABLE deal_templates MODIFY template_deal_http_headers TEXT NULL`,
					`ALTER TABLE deal_templates MODIFY template_deal_url_template TEXT NULL`,
					`ALTER TABLE deal_templates MODIFY description TEXT NULL`,
					`ALTER TABLE deal_templates MODIFY name VARCHAR(255) NULL`,
				}
				for _, query := range queries {
					if err := db.Exec(query).Error; err != nil {
						return err
					}
				}
			}

			return nil
		},
	}
}
