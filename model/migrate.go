package model

import "gorm.io/gorm"

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&Worker{},
		&Dataset{},
		&Source{},
		&Chunk{},
		&Item{},
		&BlockRaw{},
		&BlockReference{},
	)
}
