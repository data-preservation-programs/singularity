package admin

import (
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

func InitHandler(db *gorm.DB) error {
	return initHandler(db)
}

// @Summary Initialize the database
// @Tags Admin
// @Success 204
// @Failure 500 {object} handler.HTTPError
// @Router /admin/init [post]
func initHandler(db *gorm.DB) error {
	err := model.AutoMigrate(db)
	if err != nil {
		return handler.NewHandlerError(err)
	}

	return nil
}
