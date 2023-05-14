package dataset

import (
	"github.com/data-preservation-programs/go-singularity/database"
	"github.com/data-preservation-programs/go-singularity/handler"
	"github.com/data-preservation-programs/go-singularity/model"
	"github.com/ipfs/go-log/v2"
	"gorm.io/gorm"
)

// ListPieceHandler godoc
// @Summary List all pieces for a dataset
// @Tags Dataset
// @Produce json
// @Accept json
// @Param name path string true "Dataset name"
// @Success 200 {array} model.Car
// @Failure 400 {object} handler.HTTPError
// @Failure 500 {object} handler.HTTPError
// @Router /dataset/{name}/piece [get]
func ListPieceHandler(
	db *gorm.DB,
	name string,
) ([]model.Car, *handler.Error) {
	log.SetAllLoggers(log.LevelInfo)
	if name == "" {
		return nil, handler.NewBadRequestString("dataset name is required")
	}

	dataset, err := database.FindDatasetByName(db, name)
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
