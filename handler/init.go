package handler

import (
	"github.com/data-preservation-programs/go-singularity/model"
	"gorm.io/gorm"
)

// InitHandler godoc
// @Summary Initialize the database
// @Tags Global
// @Success 204
// @Failure 400 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /init [post]
func InitHandler(db *gorm.DB) *Error {
	err := model.AutoMigrate(db)
	if err != nil {
		return NewHandlerError(err)
	}

	return nil
}
