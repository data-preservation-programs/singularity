package admin

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

// @Summary Reset the database
// @Description This will drop all tables and recreate them.
// @Tags Admin
// @Success 204
// @Failure 500 {object} api.HTTPError
// @Router /admin/reset [post]
func resetHandler(db *gorm.DB) error {
	err := model.DropAll(db)
	if err != nil {
		return errors.WithStack(err)
	}

	err = model.AutoMigrate(db)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// ResetHandler resets the database by dropping all existing tables and then
// recreating them using migrations defined in the model.
// It's generally used during testing or for complete system resets, and caution
// should be exercised before invoking it in a production environment.
//
// Parameters:
// - db: A pointer to the gorm.DB instance representing the database connection.
//
// Returns:
// - An error, if any occurred during the operation.
func ResetHandler(ctx context.Context, db *gorm.DB) error {
	return resetHandler(db.WithContext(ctx))
}
