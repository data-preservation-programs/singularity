package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func AddProofTables202507210000() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "202507210000",
		Migrate: func(tx *gorm.DB) error {
			if err := tx.Exec(`
				CREATE TABLE IF NOT EXISTS deal_proofs (
					id INTEGER PRIMARY KEY AUTOINCREMENT,
					deal_id INTEGER NOT NULL UNIQUE,
					piece_c_id TEXT NOT NULL,
					sector_id INTEGER NOT NULL,
					proof_bytes BLOB NOT NULL,
					proof_status TEXT NOT NULL DEFAULT 'pending',
					client_address TEXT NOT NULL,
					provider_address TEXT NOT NULL,
					created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
					verification_time DATETIME,
					updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
				)
			`).Error; err != nil {
				return err
			}

			if err := tx.Exec(`
				CREATE TABLE IF NOT EXISTS proof_verifications (
					id INTEGER PRIMARY KEY AUTOINCREMENT,
					deal_id INTEGER NOT NULL,
					verified_by TEXT NOT NULL,
					verification_result INTEGER NOT NULL,
					verification_time DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
					created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
					updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
					FOREIGN KEY (deal_id) REFERENCES deal_proofs(deal_id)
				)
			`).Error; err != nil {
				return err
			}

			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			if err := tx.Exec("DROP TABLE IF EXISTS proof_verifications").Error; err != nil {
				return err
			}
			if err := tx.Exec("DROP TABLE IF EXISTS deal_proofs").Error; err != nil {
				return err
			}
			return nil
		},
	}
}
