package model

import (
	"crypto/rand"
	"encoding/base64"

	"github.com/cockroachdb/errors"
	"github.com/google/uuid"
	logging "github.com/ipfs/go-log/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var Tables = []any{
	&Worker{},
	&Global{},
	&Preparation{},
	&Storage{},
	&OutputAttachment{},
	&SourceAttachment{},
	&Job{},
	&File{},
	&FileRange{},
	&Directory{},
	&Car{},
	&CarBlock{},
	&Deal{},
	&Schedule{},
	&Wallet{},
}

var logger = logging.Logger("model")

// AutoMigrate attempts to automatically migrate the database schema.
//
// This function performs a few operations:
//  1. Automatically migrates the tables in the database to match the structures defined in the application.
//  2. Creates an instance ID if it doesn't already exist.
//  3. Generates a new encryption salt and stores it in the database if it doesn't already exist.
//
// The purpose of the auto-migration feature is to simplify schema changes and manage
// basic system configurations without manually altering the database. This is especially
// useful during development or when deploying updates that include schema changes.
//
// Parameters:
//   - db: A pointer to a gorm.DB object, which provides database access.
//
// Returns:
//   - An error if any issues arise during the process, otherwise nil.
func AutoMigrate(db *gorm.DB) error {
	logger.Info("Auto migrating tables")
	err := db.AutoMigrate(Tables...)
	if err != nil {
		return errors.Wrap(err, "failed to auto migrate")
	}

	logger.Debug("Creating instance id")
	err = db.Clauses(clause.OnConflict{
		DoNothing: true,
	}).Create(&Global{Key: "instance_id", Value: uuid.NewString()}).Error
	if err != nil {
		return errors.Wrap(err, "failed to create instance id")
	}

	salt := make([]byte, 8)
	_, err = rand.Read(salt)
	if err != nil {
		return errors.Wrap(err, "failed to generate salt")
	}
	encoded := base64.StdEncoding.EncodeToString(salt)
	row := Global{
		Key:   "salt",
		Value: encoded,
	}

	logger.Debug("Creating encryption salt")
	err = db.Clauses(clause.OnConflict{
		DoNothing: true,
	}).Create(row).Error
	if err != nil {
		return errors.Wrap(err, "failed to create salt")
	}

	return nil
}

// DropAll removes all tables specified in the Tables slice from the database.
//
// This function is typically used during development or testing where a clean database
// slate is required. It iterates over the predefined Tables list and drops each table.
// Care should be taken when using this function in production environments as it will
// result in data loss.
//
// Parameters:
//   - db: A pointer to a gorm.DB object, which provides database access.
//
// Returns:
//   - An error if any issues arise during the table drop process, otherwise nil.
func DropAll(db *gorm.DB) error {
	logger.Info("Dropping all tables")
	for _, table := range Tables {
		err := db.Migrator().DropTable(table)
		if err != nil {
			return errors.Wrap(err, "failed to drop table")
		}
	}
	return nil
}
