package model

import (
	"gorm.io/gorm"
)

// ListPieces returns all Car records for a given PreparationID.
func ListPieces(db *gorm.DB, prepID PreparationID) ([]Car, error) {
	var cars []Car
	err := db.Where("preparation_id = ?", prepID).Find(&cars).Error
	return cars, err
}
