package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func CreateDealProofTrackings202507210010() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "202507210010",
		Migrate: func(tx *gorm.DB) error {
			return tx.Exec(`
				CREATE TABLE IF NOT EXISTS deal_proof_trackings (
					deal_id INTEGER PRIMARY KEY,
					provider TEXT NOT NULL,
					sector_id INTEGER NOT NULL,
					sector_start_epoch INTEGER NOT NULL,
					current_deadline_index INTEGER NOT NULL,
					period_start_epoch INTEGER NOT NULL,
					estimated_next_proof_time DATETIME NOT NULL,
					faults INTEGER NOT NULL,
					recoveries INTEGER NOT NULL,
					last_updated_at DATETIME NOT NULL
				)
			`).Error
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Exec(`DROP TABLE IF EXISTS deal_proof_trackings`).Error
		},
	}
}
