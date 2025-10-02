package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func _202509171745_files_attachment_no_cascade() *gormigrate.Migration {
	// fixes deadlock that occurs when deleting preparation due to multiple cascade paths
	// we do this instead of making files.attachment_id nullable to keep value types (NULLABLE implies pointer in gorm)
	return &gormigrate.Migration{
		ID: "202509171745",
		Migrate: func(tx *gorm.DB) error {
			if tx.Dialector.Name() == "postgres" {
				//  set null on fk then rely on trigger to clean up orphans
				// postgres is more restrictive than InnoDB so we have to do this manually
				if err := tx.Exec("ALTER TABLE files ALTER COLUMN attachment_id DROP NOT NULL").Error; err != nil {
					return errors.Wrap(err, "make files.attachment_id nullable")
				}
				if err := tx.Exec("ALTER TABLE files DROP CONSTRAINT IF EXISTS fk_files_attachment").Error; err != nil {
					return errors.Wrap(err, "drop fk_files_attachment")
				}
				if err := tx.Exec("ALTER TABLE files ADD CONSTRAINT fk_files_attachment FOREIGN KEY (attachment_id) REFERENCES source_attachments(id) ON DELETE SET NULL").Error; err != nil {
					return errors.Wrap(err, "add fk_files_attachment set null")
				}

				if err := tx.Exec(`
					CREATE OR REPLACE FUNCTION delete_orphan_files() RETURNS trigger AS $$
					BEGIN
						IF NEW.attachment_id IS NULL THEN
							DELETE FROM files WHERE id = NEW.id;
						END IF;
						RETURN NULL;
					END; $$ LANGUAGE plpgsql;
				`).Error; err != nil {
					return errors.Wrap(err, "create delete_orphan_files function")
				}

				if err := tx.Exec("DROP TRIGGER IF EXISTS trg_delete_orphan_files ON files").Error; err != nil {
					return errors.Wrap(err, "drop existing trigger")
				}

				if err := tx.Exec(`
					CREATE TRIGGER trg_delete_orphan_files
					AFTER UPDATE OF attachment_id ON files
					FOR EACH ROW
					WHEN (NEW.attachment_id IS NULL)
					EXECUTE FUNCTION delete_orphan_files();
				`).Error; err != nil {
					return errors.Wrap(err, "create trigger")
				}

				return nil
			}
			if tx.Dialector.Name() == "mysql" {
				// mysql uses restrict to emulate no cascade behavior here
				if err := tx.Exec("ALTER TABLE files DROP FOREIGN KEY fk_files_attachment").Error; err != nil {
				}
				if err := tx.Exec("ALTER TABLE files ADD CONSTRAINT fk_files_attachment FOREIGN KEY (attachment_id) REFERENCES source_attachments(id) ON DELETE RESTRICT").Error; err != nil {
					return errors.Wrap(err, "add fk_files_attachment restrict")
				}
				return nil
			}
			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			if tx.Dialector.Name() == "postgres" {
				// remove trigger and function then restore previous not null and no action fk
				if err := tx.Exec("DROP TRIGGER IF EXISTS trg_delete_orphan_files ON files").Error; err != nil {
					return errors.Wrap(err, "drop trigger (rollback)")
				}
				if err := tx.Exec("DROP FUNCTION IF EXISTS delete_orphan_files()").Error; err != nil {
					return errors.Wrap(err, "drop function (rollback)")
				}
				if err := tx.Exec("ALTER TABLE files DROP CONSTRAINT IF EXISTS fk_files_attachment").Error; err != nil {
					return errors.Wrap(err, "drop fk_files_attachment (rollback)")
				}
				if err := tx.Exec("ALTER TABLE files ALTER COLUMN attachment_id SET NOT NULL").Error; err != nil {
					return errors.Wrap(err, "make files.attachment_id not null (rollback)")
				}
				if err := tx.Exec("ALTER TABLE files ADD CONSTRAINT fk_files_attachment FOREIGN KEY (attachment_id) REFERENCES source_attachments(id) ON DELETE NO ACTION").Error; err != nil {
					return errors.Wrap(err, "add fk_files_attachment no action (rollback)")
				}
				return nil
			}
			if tx.Dialector.Name() == "mysql" {
				// restore mysql cascade to match previous behavior on rollback
				if err := tx.Exec("ALTER TABLE files DROP FOREIGN KEY fk_files_attachment").Error; err != nil {
				}
				if err := tx.Exec("ALTER TABLE files ADD CONSTRAINT fk_files_attachment FOREIGN KEY (attachment_id) REFERENCES source_attachments(id) ON DELETE CASCADE").Error; err != nil {
					return errors.Wrap(err, "add fk_files_attachment cascade (mysql rollback)")
				}
				return nil
			}
			return nil
		},
	}
}
