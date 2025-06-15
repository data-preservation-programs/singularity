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

// Migrator creates a new Gormigrate instance with optional schema initialization.
func Migrator(db *gorm.DB) *gormigrate.Gormigrate {
	m := gormigrate.New(db, gormigrate.DefaultOptions, migrations.GetMigrations())

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

type migrator struct {
	gormigrate.Gormigrate
	db      *gorm.DB
	Options gormigrate.Options
}

// GetMigrator returns a customized migrator with extended helper methods.
func GetMigrator(db *gorm.DB) *migrator {
	options := gormigrate.DefaultOptions
	return &migrator{
		Gormigrate: *gormigrate.New(db, options, migrations.GetMigrations()),
		db:         db,
		Options:    options,
	}
}

// DropAll removes all tables in the database.
func (m *migrator) DropAll() error {
	tables, err := m.db.Migrator().GetTables()
	if err != nil {
		return errors.Wrap(err, "Failed to get tables")
	}
	for _, t := range tables {
		err = m.db.Migrator().DropTable(t)
		if err != nil {
			return errors.Wrap(err, "Failed to drop all tables")
		}
	}
	return nil
}

// GetMigrationsRun returns a list of all applied migrations.
func (m *migrator) GetMigrationsRun() ([]migration, error) {
	var migrations []migration
	err := m.db.Find(&migrations).Error
	if err != nil {
		return nil, err
	}
	return migrations, nil
}

// GetLastMigration returns the ID of the last migration applied.
func (m *migrator) GetLastMigration() (string, error) {
	migrations, err := m.GetMigrationsRun()
	if err != nil || len(migrations) == 0 {
		return "", err
	}
	return migrations[len(migrations)-1].ID, nil
}

// HasRunMigration checks if a migration by ID has already run.
func (m *migrator) HasRunMigration(id string) (bool, error) {
	var count int64
	err := m.db.Table(m.Options.TableName).Where(m.Options.IDColumnName+" = ?", id).Count(&count).Error
	return count > 0, err
}

// DropAll is also exposed globally for convenience.
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
