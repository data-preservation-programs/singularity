package admin

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

// ResetHandler resets the database by dropping all existing tables and then
// recreating them using migrations defined in the model.
//
// It's generally used during testing or for complete system resets, and caution
// should be exercised before invoking it in a production environment.
//
// Parameters:
//   - db: A pointer to the gorm.DB instance representing the database connection.
//
// Returns:
//   - An error, if any occurred during the operation.
func (DefaultHandler) ResetHandler(ctx context.Context, db *gorm.DB) error {
	db = db.WithContext(ctx)
	err := model.DropAll(db)
	if err != nil {
		return errors.WithStack(err)
	}

	err = model.Migrator(db).Migrate()
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
