package model

import (
	"crypto/rand"
	"encoding/base64"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/migrate/migrations"
	"github.com/go-gormigrate/gormigrate/v2"
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

// Create new Gormigrate instance
//
// If no migrations are found, an init function performs a few operations:
//  1. Automatically migrates the tables in the database to match the current structures defined in the application.
//  2. Creates an instance ID if it doesn't already exist.
//  3. Generates a new encryption salt and stores it in the database if it doesn't already exist.
//
// Parameters:
//   - db: A pointer to a gorm.DB object, which provides database access.
//
// Returns:
//   - A migration interface
func Migrator(db *gorm.DB) *gormigrate.Gormigrate {
	m := gormigrate.New(db, gormigrate.DefaultOptions, migrations.GetMigrations())

	// Initialize database with current schema if no previous migrations are found
	m.InitSchema(func(tx *gorm.DB) error {
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
	})

	return m
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
