package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

// _20230815091500_add_one_piece_per_upstream adds one_piece_per_upstream column to the preparations table
func _20230815091500_add_one_piece_per_upstream() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20230815091500",
		Migrate: func(tx *gorm.DB) error {
			return tx.Exec("ALTER TABLE preparations ADD COLUMN one_piece_per_upstream BOOLEAN NOT NULL DEFAULT FALSE").Error
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Exec("ALTER TABLE preparations DROP COLUMN IF EXISTS one_piece_per_upstream").Error
		},
	}
}
