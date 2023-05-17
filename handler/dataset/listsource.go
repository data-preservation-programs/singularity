package dataset

import (
	"github.com/data-preservation-programs/go-singularity/database"
	"github.com/data-preservation-programs/go-singularity/handler"
	"github.com/data-preservation-programs/go-singularity/model"
	"github.com/ipfs/go-log/v2"
	"gorm.io/gorm"
)

// ListSourceHandler godoc
// @Summary List all sources for a dataset
// @Tags Dataset
// @Param name path string true "Dataset name"
// @Accept json
// @Produce json
// @Success 200 {array} model.Source
// @Failure 500 {object} handler.HTTPError
// @Router /dataset/{name}/sources [get]
func ListSourceHandler(
	db *gorm.DB,
	name string,
) ([]model.Source, *handler.Error) {
	log.SetAllLoggers(log.LevelInfo)
	dataset, err := database.FindDatasetByName(db, name)
	if err != nil {
		return nil, handler.NewBadRequestString("failed to find dataset: " + err.Error())
	}
	var sources []model.Source
	err = db.Where("dataset_id = ?", dataset.ID).Find(&sources).Error
	if err != nil {
		return nil, handler.NewHandlerError(err)
	}
	return sources, nil
}
