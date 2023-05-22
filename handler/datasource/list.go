package datasource

import (
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/ipfs/go-log/v2"
	"gorm.io/gorm"
)

// ListSourceHandler godoc
// @Summary List all sources for a dataset
// @Tags Data Source
// @Accept json
// @Produce json
// @Param dataset query string false "Dataset name"
// @Success 200 {array} model.Source
// @Failure 500 {object} handler.HTTPError
// @Router /sources [get]
func ListSourceHandler(
	db *gorm.DB,
	datasetName string,
) ([]model.Source, *handler.Error) {
	log.SetAllLoggers(log.LevelInfo)
	var sources []model.Source
	if datasetName == "" {
		err := db.Find(&sources).Error
		if err != nil {
			return nil, handler.NewHandlerError(err)
		}
		return sources, nil
	}
	dataset, err := database.FindDatasetByName(db, datasetName)
	if err != nil {
		return nil, handler.NewBadRequestString("failed to find dataset: " + err.Error())
	}
	err = db.Where("dataset_id = ?", dataset.ID).Find(&sources).Error
	if err != nil {
		return nil, handler.NewHandlerError(err)
	}
	return sources, nil
}
