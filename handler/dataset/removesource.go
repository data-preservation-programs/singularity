package dataset

import (
	"github.com/data-preservation-programs/go-singularity/database"
	"github.com/data-preservation-programs/go-singularity/handler"
	"github.com/data-preservation-programs/go-singularity/model"
	"github.com/ipfs/go-log/v2"
	"gorm.io/gorm"
)

// RemoveSourceHandler godoc
// @Summary Remove a source from a dataset
// @Tags Dataset
// @Param name path string true "Dataset name"
// @Param sourcePath path string true "Source path"
// @Success 204
// @Failure 400 {object} handler.HTTPError
// @Failure 500 {object} handler.HTTPError
// @Router /dataset/{name}/source/{sourcePath} [delete]
func RemoveSourceHandler(
	db *gorm.DB,
	name string,
	sourcePath string,
) *handler.Error {
	log.SetAllLoggers(log.LevelInfo)
	dataset, err := database.FindDatasetByName(db, name)
	if err != nil {
		return handler.NewBadRequestString("failed to find dataset: " + err.Error())
	}
	var source model.Source
	err = db.Where("dataset_id = ? AND path = ?", dataset.ID, sourcePath).First(&source).Error
	if err != nil {
		return handler.NewBadRequestError(err)
	}
	err = db.Delete(&source).Error
	if err != nil {
		return handler.NewHandlerError(err)
	}
	return nil
}
