package admin

import (
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

// InitHandler godoc
// @Summary Initialize the database
// @Tags Admin
// @Success 204
// @Failure 500 {object} handler.HTTPError
// @Router /admin/init [post]
func InitHandler(db *gorm.DB) *handler.Error {
	err := model.AutoMigrate(db)
	if err != nil {
		return handler.NewHandlerError(err)
	}

	return nil
}
