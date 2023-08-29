package schedule

import (
	"context"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

func ListByDatasetHandler(
	ctx context.Context,
	db *gorm.DB,
	datasetName string,
) ([]model.Schedule, error) {
	return listByDatasetHandler(ctx, db, datasetName)
}

// @Summary List deal making schedules by dataset
// @Tags Deal Schedule
// @Produce json
// @Success 200 {array} model.Schedule
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /schedules [get]
func listByDatasetHandler(
	ctx context.Context,
	db *gorm.DB,
	datasetName string,
) ([]model.Schedule, error) {
	db = db.WithContext(ctx)
	var schedules []model.Schedule
	if datasetName == "" {
		err := db.Find(&schedules).Error
		if err != nil {
			return nil, err
		}
		return schedules, nil
	}
	dataset, err := database.FindDatasetByName(db, datasetName)
	if err != nil {
		return nil, handler.NewInvalidParameterErr("failed to find dataset: " + err.Error())
	}
	err = db.Where("dataset_id = ?", dataset.ID).Find(&schedules).Error
	if err != nil {
		return nil, err
	}
	return schedules, nil
}

// @Summary List all deal making schedules
// @Tags Deal Schedule
// @Produce json
// @Success 200 {array} model.Schedule
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /schedules [get]
func listHandler(
	ctx context.Context,
	db *gorm.DB,
) ([]model.Schedule, error) {
	return ListByDatasetHandler(ctx, db, "")
}

func ListHandler(
	ctx context.Context,
	db *gorm.DB,
) ([]model.Schedule, error) {
	return listHandler(ctx, db)
}
