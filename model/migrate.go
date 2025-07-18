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
	&Notification{},
	&DealTemplate{},
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
	&ErrorLog{},
}

var logger = logging.Logger("model")

// Options for gormigrate instance
var options = &gormigrate.Options{
	TableName:                 "migrations",
	IDColumnName:              "id",
	IDColumnSize:              255,
	UseTransaction:            false,
	ValidateUnknownMigrations: false,
}

// NOTE: this NEEDS to match the values in MigrationOptions above
//
//	type <TableName (singular)> struct {
//	  ID string `gorm:"primaryKey;column:<IDColumnName>;size:<IDColumnSize>"`
//	}
type migration struct {
	ID string `gorm:"primaryKey;column:id;size:255"`
}

// Handle initializing any database if no migrations are found
//
// In the case of existing database:
//  1. Migrations table is created and first migration is inserted, which should match the existing, if outdated, data
//  2. Any remaining migrations are run
//
// In the case of a new database:
//  1. Automatically migrates the tables in the database to match the current structures defined in the application.
//  2. Creates an instance ID if it doesn't already exist.
//  3. Generates a new encryption salt and stores it in the database if it doesn't already exist.
func _init(db *gorm.DB) error {
	logger.Info("Initializing database")

	// If this is an existing database before versioned migration strategy was implemented
	if db.Migrator().HasTable("wallets") && !db.Migrator().HasColumn("wallets", "actor_id") {
		// NOTE: We're going to have to recreate some internals of Gormigrate. It would be cleaner
		// to use them directly but they're private methods. The general idea is to run all the
		// migration functions _except_ the first which is hopefully the state of the database
		// when they were on the older automigrate strategy.
		logger.Info("Manually creating versioned migration table in existing database")

		// Create migrations table
		err := db.Table(options.TableName).AutoMigrate(&migration{})
		if err != nil {
			return errors.Wrap(err, "failed to create migrations table on init")
		}

		logger.Info("Manually running missing migrations")
		// Skip first migration, run the rest to get current
		for _, m := range migrations.GetMigrations()[1:] {
			err = m.Migrate(db)
			if err != nil {
				return errors.Wrap(err, "failed to run migration with ID: "+m.ID)
			}
		}
	} else {
		logger.Info("Auto migrating tables in clean database")
		// This is a brand new database, run automigrate script on current schema
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
	}

	return nil
}

type migrator struct {
	gormigrate.Gormigrate
	db      *gorm.DB
	Options gormigrate.Options
}

// Drop all current database tables
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

// Get all migrations run
func (m *migrator) GetMigrationsRun() ([]migration, error) {
	var migrations []migration
	err := m.db.Find(&migrations).Error
	if err != nil {
		return nil, err
	}
	return migrations, nil
}

// Get ID of last migration ran
func (m *migrator) GetLastMigration() (string, error) {
	migrations, err := m.GetMigrationsRun()
	if len(migrations) == 0 || err != nil {
		return "", err
	}
	return migrations[len(migrations)-1].ID, nil
}

// Has migration ID ran
func (m *migrator) HasRunMigration(id string) (bool, error) {
	var count int64
	err := m.db.Table(m.Options.TableName).Where(m.Options.IDColumnName+" = ?", id).Count(&count).Error
	return count > 0, err
}

// Setup new Gormigrate instance
//
// Parameters:
//   - db: A pointer to a gorm.DB object, which provides database access.
//
// Returns:
//   - A migration interface
func GetMigrator(db *gorm.DB) *migrator {
	g := gormigrate.New(db, options, migrations.GetMigrations())

	// Initialize database with current schema if no previous migrations are found
	g.InitSchema(_init)

	return &migrator{
		*g,
		db,
		*options,
	}
}
