package admin

import (
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

// @Summary Reset the database
// @Description This will drop all tables and recreate them.
// @Tags Admin
// @Success 204
// @Failure 500 {object} handler.HTTPError
// @Router /admin/reset [post]
func resetHandler(db *gorm.DB) error {
	err := model.DropAll(db)
	if err != nil {
		return handler.NewHandlerError(err)
	}

	err = model.AutoMigrate(db)
	if err != nil {
		return handler.NewHandlerError(err)
	}

	return nil
}

func ResetHandler(db *gorm.DB) error {
	return resetHandler(db)
}
