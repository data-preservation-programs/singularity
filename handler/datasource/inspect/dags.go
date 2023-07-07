package inspect

import (
	"strconv"

	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// GetDagsHandler godoc
// @Summary Get all dag details of a data source
// @Tags Data Source
// @Accept json
// @Produce json
// @Param id path string true "Source ID"
// @Success 200 {array} model.Car
// @Failure 500 {object} handler.HTTPError
// @Router /source/{id}/chunks [get]
func GetDagsHandler(
	db *gorm.DB,
	id string,
) ([]model.Car, error) {
	sourceID, err := strconv.Atoi(id)
	if err != nil {
		return nil, handler.NewBadRequestString("invalid source id")
	}
	var source model.Source
	err = db.Where("id = ?", sourceID).First(&source).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, handler.NewBadRequestString("source not found")
	}
	if err != nil {
		return nil, handler.NewHandlerError(err)
	}

	var cars []model.Car
	err = db.Where("source_id = ? AND chunk_id IS NULL", sourceID).Find(&cars).Error
	if err != nil {
		return nil, handler.NewHandlerError(err)
	}

	return cars, nil
}
