package dataset

import (
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

// ListPiecesHandler godoc
// @Summary List all pieces for the dataset that are available for deal making
// @Tags Dataset
// @Produce json
// @Accept json
// @Param datasetName path string true "Dataset name"
// @Success 200 {array} model.Car
// @Failure 400 {object} handler.HTTPError
// @Failure 500 {object} handler.HTTPError
// @Router /dataset/{datasetName}/piece [get]
func ListPiecesHandler(
	db *gorm.DB,
	datasetName string,
) ([]model.Car, error) {
	if datasetName == "" {
		return nil, handler.NewBadRequestString("dataset name is required")
	}

	dataset, err := database.FindDatasetByName(db, datasetName)
	if err != nil {
		return nil, handler.NewBadRequestString("failed to find dataset: " + err.Error())
	}

	var cars []model.Car
	err = db.Where("dataset_id = ?", dataset.ID).Find(&cars).Error
	if err != nil {
		return nil, handler.NewHandlerError(err)
	}

	return cars, nil
}
