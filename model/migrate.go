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
	&Actor{},
	&Wallet{},
	&PDPProofSet{},
}

var logger = logging.Logger("model")

// fkMigration defines a FK constraint that should be SET NULL instead of CASCADE.
type fkMigration struct {
	table      string
	constraint string
	column     string
	refTable   string
}

// fkMigrations lists FK constraints that need to be changed to SET NULL for fast prep deletion.
var fkMigrations = []fkMigration{
	{"car_blocks", "fk_car_blocks_car", "car_id", "cars"},
	{"car_blocks", "fk_car_blocks_file", "file_id", "files"},
	{"cars", "fk_cars_preparation", "preparation_id", "preparations"},
	{"cars", "fk_cars_attachment", "attachment_id", "source_attachments"},
	{"files", "fk_files_attachment", "attachment_id", "source_attachments"},
	{"directories", "fk_directories_attachment", "attachment_id", "source_attachments"},
	{"jobs", "fk_jobs_attachment", "attachment_id", "source_attachments"},
}

// AutoMigrate attempts to automatically migrate the database schema.
//
// This function performs a few operations:
//  1. Automatically migrates the tables in the database to match the structures defined in the application.
//  2. Migrates FK constraints that changed from CASCADE to SET NULL (for fast prep deletion).
//  3. Creates an instance ID if it doesn't already exist.
//  4. Generates a new encryption salt and stores it in the database if it doesn't already exist.
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

	// Migrate FK constraints from CASCADE to SET NULL for fast preparation deletion.
	// GORM doesn't update existing constraints, so we do it manually.
	if err := migrateFKConstraints(db); err != nil {
		return errors.Wrap(err, "failed to migrate FK constraints")
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

	// Fix postgres sequences if they're out of sync (e.g., after data import)
	if err := fixPostgresSequences(db); err != nil {
		return errors.Wrap(err, "failed to fix sequences")
	}

	// Infer piece_type for cars that predate the column
	if err := inferPieceTypes(db); err != nil {
		return errors.Wrap(err, "failed to infer piece types")
	}

	// Set deal_type for existing deals that predate the column
	if err := inferDealTypes(db); err != nil {
		return errors.Wrap(err, "failed to infer deal types")
	}

	// Set deal_type for existing schedules that predate the column
	if err := inferScheduleDealTypes(db); err != nil {
		return errors.Wrap(err, "failed to infer schedule deal types")
	}

	// Drop legacy fk_deals_actor constraint (Deal.ClientID no longer FKs to Actor)
	if err := dropDealActorFK(db); err != nil {
		return errors.Wrap(err, "failed to drop deal-actor FK")
	}

	// Backfill wallet_id for deals that predate the column
	if err := backfillDealWalletID(db); err != nil {
		return errors.Wrap(err, "failed to backfill deal wallet IDs")
	}

	// Migrate wallet_assignments many2many to preparation.wallet_id 1:1
	if err := migrateWalletAssignments(db); err != nil {
		return errors.Wrap(err, "failed to migrate wallet assignments")
	}

	return nil
}

// migrateFKConstraints updates FK constraints from CASCADE to SET NULL where needed.
// This is idempotent - it checks the current constraint before modifying.
func migrateFKConstraints(db *gorm.DB) error {
	dialect := db.Dialector.Name()
	if dialect == "sqlite" {
		// SQLite doesn't support ALTER CONSTRAINT, and FK enforcement is optional anyway
		return nil
	}

	for _, fk := range fkMigrations {
		// Check if constraint exists and uses CASCADE
		var deleteRule string
		var err error

		if dialect == "postgres" {
			err = db.Raw(`
				SELECT rc.delete_rule
				FROM information_schema.referential_constraints rc
				JOIN information_schema.table_constraints tc ON rc.constraint_name = tc.constraint_name
				WHERE tc.table_name = ? AND tc.constraint_name = ?
			`, fk.table, fk.constraint).Scan(&deleteRule).Error
		} else if dialect == "mysql" {
			err = db.Raw(`
				SELECT DELETE_RULE
				FROM information_schema.REFERENTIAL_CONSTRAINTS
				WHERE TABLE_NAME = ? AND CONSTRAINT_NAME = ?
			`, fk.table, fk.constraint).Scan(&deleteRule).Error
		}

		if err != nil {
			// Constraint might not exist yet (new install), skip
			logger.Debugw("constraint check skipped", "table", fk.table, "constraint", fk.constraint, "error", err)
			continue
		}

		if deleteRule == "" {
			// Constraint doesn't exist, will be created correctly by AutoMigrate
			continue
		}

		if deleteRule == "SET NULL" {
			// Already migrated
			continue
		}

		logger.Infow("migrating FK constraint to SET NULL", "table", fk.table, "constraint", fk.constraint)

		// Drop and recreate with SET NULL
		if dialect == "postgres" {
			// Postgres DDL is transactional - wrap DROP+ADD so failure rolls back both
			// Use NOT VALID to skip row validation - existing rows were valid under CASCADE,
			// so they're still valid under SET NULL.
			err = db.Transaction(func(tx *gorm.DB) error {
				if err := tx.Exec(`ALTER TABLE ` + fk.table + ` DROP CONSTRAINT ` + fk.constraint).Error; err != nil {
					return err
				}
				return tx.Exec(`ALTER TABLE ` + fk.table + ` ADD CONSTRAINT ` + fk.constraint +
					` FOREIGN KEY (` + fk.column + `) REFERENCES ` + fk.refTable + `(id) ON DELETE SET NULL NOT VALID`).Error
			})
			if err != nil {
				return errors.Wrapf(err, "failed to migrate constraint %s", fk.constraint)
			}
		} else if dialect == "mysql" {
			// MySQL DDL causes implicit commit, so no transaction benefit here
			err = db.Exec(`ALTER TABLE ` + fk.table + ` DROP FOREIGN KEY ` + fk.constraint).Error
			if err != nil {
				return errors.Wrapf(err, "failed to drop constraint %s", fk.constraint)
			}
			err = db.Exec(`ALTER TABLE ` + fk.table + ` ADD CONSTRAINT ` + fk.constraint +
				` FOREIGN KEY (` + fk.column + `) REFERENCES ` + fk.refTable + `(id) ON DELETE SET NULL`).Error
			if err != nil {
				return errors.Wrapf(err, "failed to create constraint %s", fk.constraint)
			}
		}
	}

	return nil
}

// sequenceTable maps table names to their primary key column for sequence fixing.
// Only tables with numeric auto-increment PKs are included.
var sequenceTables = []string{
	"preparations",
	"storages",
	"output_attachments",
	"source_attachments",
	"jobs",
	"files",
	"file_ranges",
	"directories",
	"cars",
	"car_blocks",
	"deals",
	"schedules",
	"wallets",
}

// fixPostgresSequences detects and fixes out-of-sync sequences.
// This can happen when data is imported with explicit IDs (e.g., from MySQL).
// PostgreSQL sequences don't auto-update on INSERT with explicit ID values.
func fixPostgresSequences(db *gorm.DB) error {
	if db.Dialector.Name() != "postgres" {
		return nil
	}

	for _, table := range sequenceTables {
		var maxID, lastValue int64

		// get max id from table
		err := db.Raw(`SELECT COALESCE(MAX(id), 0) FROM ` + table).Scan(&maxID).Error
		if err != nil {
			// table might not exist yet
			logger.Debugw("skipping sequence check", "table", table, "error", err)
			continue
		}

		// get sequence name and current value
		seqName := table + "_id_seq"
		err = db.Raw(`SELECT last_value FROM ` + seqName).Scan(&lastValue).Error
		if err != nil {
			logger.Debugw("skipping sequence check", "sequence", seqName, "error", err)
			continue
		}

		// if max(id) >= sequence value, sequence is stale
		if maxID >= lastValue {
			logger.Infow("fixing stale sequence", "table", table, "maxID", maxID, "lastValue", lastValue)
			err = db.Exec(`SELECT setval(?, ?, true)`, seqName, maxID).Error
			if err != nil {
				return errors.Wrapf(err, "failed to fix sequence %s", seqName)
			}
		}
	}

	return nil
}

// inferPieceTypes sets piece_type for cars that predate the column.
// for inline preps, a car is "data" if any of its car_blocks reference files.
// for non-inline preps, car_blocks don't reference files (data is on disk),
// so we fall back to num_of_files > 0 which is only set by the packer.
// everything else is "dag" (directory metadata only).
// idempotent - only updates rows where piece_type is NULL or empty.
func inferPieceTypes(db *gorm.DB) error {
	dialect := db.Dialector.Name()

	// check if any cars need updating
	var count int64
	err := db.Raw(`SELECT COUNT(*) FROM cars WHERE piece_type IS NULL OR piece_type = ''`).Scan(&count).Error
	if err != nil {
		// table might not exist or column missing
		logger.Debugw("skipping piece type inference", "error", err)
		return nil
	}

	if count == 0 {
		return nil
	}

	logger.Infow("inferring piece types for legacy cars", "count", count)

	// dialect-specific UPDATE with subquery
	var query string
	if dialect == "sqlite" {
		query = `
			UPDATE cars SET piece_type = (
				CASE WHEN num_of_files > 0 OR EXISTS (
					SELECT 1 FROM car_blocks WHERE car_blocks.car_id = cars.id AND car_blocks.file_id IS NOT NULL
				) THEN 'data' ELSE 'dag' END
			) WHERE piece_type IS NULL OR piece_type = ''`
	} else {
		query = `
			UPDATE cars c SET piece_type = CASE
				WHEN c.num_of_files > 0 OR EXISTS (
					SELECT 1 FROM car_blocks cb WHERE cb.car_id = c.id AND cb.file_id IS NOT NULL
				) THEN 'data' ELSE 'dag'
			END WHERE c.piece_type IS NULL OR c.piece_type = ''`
	}

	result := db.Exec(query)
	if result.Error != nil {
		return errors.Wrap(result.Error, "failed to infer piece types")
	}

	logger.Infow("inferred piece types", "updated", result.RowsAffected)
	return nil
}

// inferDealTypes sets deal_type for deals that predate the column.
// All existing deals are assumed to be legacy market deals (f05).
// This is idempotent - only updates rows where deal_type is NULL or empty.
func inferDealTypes(db *gorm.DB) error {
	// check if any deals need updating
	var count int64
	err := db.Raw(`SELECT COUNT(*) FROM deals WHERE deal_type IS NULL OR deal_type = ''`).Scan(&count).Error
	if err != nil {
		// table might not exist or column missing
		logger.Debugw("skipping deal type inference", "error", err)
		return nil
	}

	if count == 0 {
		return nil
	}

	logger.Infow("setting deal type for legacy deals", "count", count)

	// All existing deals are legacy market deals
	result := db.Exec(`UPDATE deals SET deal_type = 'market' WHERE deal_type IS NULL OR deal_type = ''`)
	if result.Error != nil {
		return errors.Wrap(result.Error, "failed to set deal types")
	}

	logger.Infow("set deal types", "updated", result.RowsAffected)
	return nil
}

// inferScheduleDealTypes sets deal_type for schedules that predate the column.
// same logic as inferDealTypes -- all pre-existing schedules are market.
func inferScheduleDealTypes(db *gorm.DB) error {
	var count int64
	err := db.Raw(`SELECT COUNT(*) FROM schedules WHERE deal_type IS NULL OR deal_type = ''`).Scan(&count).Error
	if err != nil {
		logger.Debugw("skipping schedule deal type inference", "error", err)
		return nil
	}
	if count == 0 {
		return nil
	}

	logger.Infow("setting deal type for legacy schedules", "count", count)
	result := db.Exec(`UPDATE schedules SET deal_type = 'market' WHERE deal_type IS NULL OR deal_type = ''`)
	if result.Error != nil {
		return errors.Wrap(result.Error, "failed to set schedule deal types")
	}
	logger.Infow("set schedule deal types", "updated", result.RowsAffected)
	return nil
}

// dropDealActorFK removes the legacy fk_deals_actor constraint if it exists.
// Deal.ClientID is now a plain string (no FK to actors table).
func dropDealActorFK(db *gorm.DB) error {
	dialect := db.Dialector.Name()
	if dialect == "sqlite" {
		return nil
	}

	constraint := "fk_deals_actor"
	var exists bool

	if dialect == "postgres" {
		err := db.Raw(`
			SELECT EXISTS (
				SELECT 1 FROM information_schema.table_constraints
				WHERE table_name = 'deals' AND constraint_name = ?
			)`, constraint).Scan(&exists).Error
		if err != nil {
			return errors.Wrapf(err, "failed to check constraint %s", constraint)
		}
	} else if dialect == "mysql" {
		err := db.Raw(`
			SELECT COUNT(*) > 0 FROM information_schema.TABLE_CONSTRAINTS
			WHERE TABLE_NAME = 'deals' AND CONSTRAINT_NAME = ?
		`, constraint).Scan(&exists).Error
		if err != nil {
			return errors.Wrapf(err, "failed to check constraint %s", constraint)
		}
	}

	if !exists {
		return nil
	}

	logger.Infow("dropping legacy deal-actor FK constraint", "constraint", constraint)
	if dialect == "postgres" {
		return db.Exec(`ALTER TABLE deals DROP CONSTRAINT ` + constraint).Error
	}
	// mysql
	return db.Exec(`ALTER TABLE deals DROP FOREIGN KEY ` + constraint).Error
}

// backfillDealWalletID sets wallet_id for existing deals that have a client_id
// matching an actor linked to a wallet. idempotent — only touches NULL wallet_id rows.
func backfillDealWalletID(db *gorm.DB) error {
	var count int64
	err := db.Raw(`SELECT COUNT(*) FROM deals WHERE wallet_id IS NULL AND client_id != ''`).Scan(&count).Error
	if err != nil {
		logger.Debugw("skipping wallet_id backfill", "error", err)
		return nil
	}
	if count == 0 {
		return nil
	}

	logger.Infow("backfilling deal wallet_id from client_id → actor → wallet", "count", count)

	dialect := db.Dialector.Name()
	var query string
	if dialect == "sqlite" {
		query = `
			UPDATE deals SET wallet_id = (
				SELECT w.id FROM wallets w
				WHERE w.actor_id = deals.client_id
				LIMIT 1
			) WHERE wallet_id IS NULL AND client_id != ''
			AND EXISTS (SELECT 1 FROM wallets w WHERE w.actor_id = deals.client_id)`
	} else {
		query = `
			UPDATE deals d
			SET wallet_id = (
				SELECT w.id FROM wallets w
				WHERE w.actor_id = d.client_id
				LIMIT 1
			) WHERE d.wallet_id IS NULL AND d.client_id != ''
			AND EXISTS (SELECT 1 FROM wallets w WHERE w.actor_id = d.client_id)`
	}

	result := db.Exec(query)
	if result.Error != nil {
		return errors.Wrap(result.Error, "failed to backfill wallet_id")
	}
	logger.Infow("backfilled deal wallet_id", "updated", result.RowsAffected)
	return nil
}

// migrateWalletAssignments moves from many2many wallet_assignments to
// preparation.wallet_id 1:1. idempotent — skips if table doesn't exist.
func migrateWalletAssignments(db *gorm.DB) error {
	if !db.Migrator().HasTable("wallet_assignments") {
		return nil
	}

	type row struct {
		PreparationID uint
		WalletID      uint
	}
	var rows []row
	err := db.Raw(`SELECT preparation_id, wallet_id FROM wallet_assignments ORDER BY preparation_id, wallet_id`).Scan(&rows).Error
	if err != nil {
		return errors.Wrap(err, "failed to read wallet_assignments")
	}

	// group by preparation
	byPrep := make(map[uint][]uint)
	for _, r := range rows {
		byPrep[r.PreparationID] = append(byPrep[r.PreparationID], r.WalletID)
	}

	for prepID, walletIDs := range byPrep {
		if len(walletIDs) > 1 {
			logger.Warnw("preparation has multiple wallets, picking lowest ID",
				"preparation_id", prepID, "wallet_ids", walletIDs)
		}
		// walletIDs are ordered by wallet_id (from ORDER BY in query above)
		err := db.Exec(`UPDATE preparations SET wallet_id = ? WHERE id = ? AND wallet_id IS NULL`, walletIDs[0], prepID).Error
		if err != nil {
			return errors.Wrapf(err, "failed to set wallet_id for preparation %d", prepID)
		}
	}

	if err := db.Migrator().DropTable("wallet_assignments"); err != nil {
		return errors.Wrap(err, "failed to drop wallet_assignments")
	}
	logger.Infow("migrated wallet_assignments to preparation.wallet_id", "preparations", len(byPrep))
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
