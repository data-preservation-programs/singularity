package model

import "gorm.io/gorm"

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&Global{},
		&Worker{},
		&Dataset{},
		&Source{},
		&Chunk{},
		&Item{},
		&Directory{},
		&Car{},
		&RawBlock{},
		&ItemBlock{},
		&CarBlock{},
		&Deal{},
		&Schedule{},
		&Wallet{},
		&WalletAssignment{},
	)
}
