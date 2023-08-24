package admin

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

// InitHandler initializes the database by running migrations on the provided db instance.
// It ensures the database schema matches the expected schema defined in the model.
//
// Parameters:
// - db: A pointer to the gorm.DB instance representing the database connection.
//
// Returns:
// - An error, if any occurred during the operation.
func InitHandler(ctx context.Context, db *gorm.DB) error {
	return initHandler(db.WithContext(ctx))
}

// @Summary Initialize the database
// @Tags Admin
// @Success 204
// @Failure 500 {object} api.HTTPError
// @Router /admin/init [post]
func initHandler(db *gorm.DB) error {
	err := model.AutoMigrate(db)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
