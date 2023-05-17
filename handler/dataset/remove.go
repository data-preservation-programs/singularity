package dataset

import (
	"github.com/data-preservation-programs/go-singularity/database"
	"github.com/data-preservation-programs/go-singularity/handler"
	"github.com/ipfs/go-log/v2"
	"gorm.io/gorm"
)

// RemoveHandler godoc
// @Summary Remove a dataset
// @Tags Dataset
// @Param name path string true "Dataset name"
// @Success 204
// @Failure 400 {object} handler.HTTPError
// @Failure 500 {object} handler.HTTPError
// @Router /dataset/{datasetName} [delete]
func RemoveHandler(
	db *gorm.DB,
	name string,
) *handler.Error {
	log.SetAllLoggers(log.LevelInfo)
	dataset, err := database.FindDatasetByName(db, name)
	if err != nil {
		return handler.NewBadRequestString("failed to find dataset: " + err.Error())
	}
	err = db.Delete(&dataset).Error
	if err != nil {
		return handler.NewHandlerError(err)
	}
	return nil
}
