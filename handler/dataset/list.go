package dataset

import (
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

// @Summary List all datasets
// @Tags Dataset
// @Produce json
// @Success 200 {array} model.Dataset
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /dataset [get]
func listHandler(
	db *gorm.DB,
) ([]model.Dataset, error) {
	var datasets []model.Dataset
	err := db.Find(&datasets).Error
	if err != nil {
		return nil, err
	}
	return datasets, nil
}

func ListHandler(
	db *gorm.DB,
) ([]model.Dataset, error) {
	return listHandler(db)
}
