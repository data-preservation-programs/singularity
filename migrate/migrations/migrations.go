package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
)

// Get collection of all migrations in order
func GetMigrations() []*gormigrate.Migration {
	return []*gormigrate.Migration{
		_202505010830_initial_schema(),
		_202505010840_wallet_actor_id(),
		_202506240815_create_notifications(),
		_202506240816_create_deal_templates(),
	}
}
