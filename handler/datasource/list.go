package datasource

import (
	"context"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

func ListSourcesByDatasetHandler(
	ctx context.Context,
	db *gorm.DB,
	datasetName string,
) ([]model.Source, error) {
	db = db.WithContext(ctx)
	var sources []model.Source
	if datasetName == "" {
		err := db.Find(&sources).Error
		if err != nil {
			return nil, err
		}
		return sources, nil
	}
	dataset, err := database.FindDatasetByName(db, datasetName)
	if err != nil {
		return nil, handler.NewInvalidParameterErr("failed to find dataset: " + err.Error())
	}
	err = db.Where("dataset_id = ?", dataset.ID).Find(&sources).Error
	if err != nil {
		return nil, err
	}
	return sources, nil
}

// @Summary List all sources for a dataset
// @Tags Data Source
// @Produce json
// @Param dataset query string false "Dataset name"
// @Success 200 {array} model.Source
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /source [get]
func listSourceHandler(
	ctx context.Context,
	db *gorm.DB,
) ([]model.Source, error) {
	return ListSourcesByDatasetHandler(ctx, db, "")
}

func ListSourceHandler(
	ctx context.Context,
	db *gorm.DB,
) ([]model.Source, error) {
	return listSourceHandler(ctx, db.WithContext(ctx))
}
