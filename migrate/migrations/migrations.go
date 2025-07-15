package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
)

// Get collection of all migrations in order
func GetMigrations() []*gormigrate.Migration {
	return []*gormigrate.Migration{
		_202505010830InitialSchema(),
		_202505010840WalletActorID(),
		_202506240815_create_notifications(),
		_202506240816_create_deal_templates(),
		_202507090900_add_missing_deal_template_fields(),
		_202507090915_add_not_null_defaults(),
		_202507091000_add_schedule_fields_to_deal_templates(),
		_20230815091500_add_one_piece_per_upstream(),
	}
}
