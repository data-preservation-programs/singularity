package dataset

import (
	"context"

	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

// @Summary List all datasets
// @Tags Preparation
// @Produce json
// @Success 200 {array} model.Preparation
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /dataset [get]
func listHandler(
	db *gorm.DB,
) ([]model.Preparation, error) {
	var datasets []model.Preparation
	err := db.Find(&datasets).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return datasets, nil
}

func ListHandler(
	ctx context.Context,
	db *gorm.DB,
) ([]model.Preparation, error) {
	return listHandler(db.WithContext(ctx))
}
