package datasource

import (
	"strconv"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// @Summary Remove a source
// @Tags Data Source
// @Param id path string true "Source ID"
// @Success 204
// @Failure 400 {object} handler.HTTPError
// @Failure 500 {object} handler.HTTPError
// @Router /source/{id} [delete]
func removeSourceHandler(
	db *gorm.DB,
	id string,
) error {
	var source model.Source
	sourceID, err := strconv.Atoi(id)
	if err != nil {
		return handler.NewBadRequestString("invalid source id")
	}
	err = db.Where("id = ?", sourceID).First(&source).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return handler.NewBadRequestString("source not found")
	}
	if err != nil {
		return handler.NewHandlerError(err)
	}
	err = database.DoRetry(func() error { return db.Delete(&source).Error })
	if err != nil {
		return handler.NewHandlerError(err)
	}
	return nil
}

func RemoveSourceHandler(
	db *gorm.DB,
	id string,
) error {
	return removeSourceHandler(db, id)
}
