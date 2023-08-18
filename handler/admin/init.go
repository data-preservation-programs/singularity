package admin

import (
	"context"

	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

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
