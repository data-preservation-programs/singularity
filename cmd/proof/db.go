package cmd

import (
	"gorm.io/gorm"
)

var db *gorm.DB

// InitDB initializes the database connection
func InitDB(database *gorm.DB) {
	db = database
}
