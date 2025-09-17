package migrations

import (
    "github.com/go-gormigrate/gormigrate/v2"
    "github.com/pkg/errors"
    "gorm.io/gorm"
)

func _202509171745_files_attachment_no_cascade() *gormigrate.Migration {
    return &gormigrate.Migration{
        ID: "202509171745",
        Migrate: func(tx *gorm.DB) error {
            if tx.Dialector.Name() == "postgres" {
                if err := tx.Exec("ALTER TABLE files DROP CONSTRAINT IF EXISTS fk_files_attachment").Error; err != nil {
                    return errors.Wrap(err, "drop fk_files_attachment")
                }
                if err := tx.Exec("ALTER TABLE files ADD CONSTRAINT fk_files_attachment FOREIGN KEY (attachment_id) REFERENCES source_attachments(id) ON DELETE NO ACTION").Error; err != nil {
                    return errors.Wrap(err, "add fk_files_attachment no action")
                }
                return nil
            }
            if tx.Dialector.Name() == "mysql" {
                // MySQL treats NO ACTION as RESTRICT; omit ON DELETE to use default (RESTRICT)
                if err := tx.Exec("ALTER TABLE files DROP FOREIGN KEY fk_files_attachment").Error; err != nil {
                }
                if err := tx.Exec("ALTER TABLE files ADD CONSTRAINT fk_files_attachment FOREIGN KEY (attachment_id) REFERENCES source_attachments(id)").Error; err != nil {
                    return errors.Wrap(err, "add fk_files_attachment restrict (mysql)")
                }
                return nil
            }
            return nil
        },
        Rollback: func(tx *gorm.DB) error {
            if tx.Dialector.Name() == "postgres" {
                if err := tx.Exec("ALTER TABLE files DROP CONSTRAINT IF EXISTS fk_files_attachment").Error; err != nil {
                    return errors.Wrap(err, "drop fk_files_attachment (rollback)")
                }
                if err := tx.Exec("ALTER TABLE files ADD CONSTRAINT fk_files_attachment FOREIGN KEY (attachment_id) REFERENCES source_attachments(id) ON DELETE CASCADE").Error; err != nil {
                    return errors.Wrap(err, "add fk_files_attachment cascade (rollback)")
                }
                return nil
            }
            if tx.Dialector.Name() == "mysql" {
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


