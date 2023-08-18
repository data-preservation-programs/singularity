package inspect

import (
	"context"
	"strconv"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

func GetDagsHandler(
	ctx context.Context,
	db *gorm.DB,
	id string,
) ([]model.Car, error) {
	return getDagsHandler(db.WithContext(ctx), id)
}

// @Summary Get all dag details of a data source
// @Tags Data Source
// @Accept json
// @Produce json
// @Param id path string true "Source ID"
// @Success 200 {array} model.Car
// @Failure 500 {object} api.HTTPError
// @Router /source/{id}/packjobs [get]
func getDagsHandler(
	db *gorm.DB,
	id string,
) ([]model.Car, error) {
	sourceID, err := strconv.Atoi(id)
	if err != nil {
		return nil, handler.NewInvalidParameterErr("invalid source id")
	}
	var source model.Source
	err = db.Where("id = ?", sourceID).First(&source).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, handler.NewInvalidParameterErr("source not found")
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var cars []model.Car
	err = db.Where("source_id = ? AND pack_job_id IS NULL", sourceID).Find(&cars).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return cars, nil
}
