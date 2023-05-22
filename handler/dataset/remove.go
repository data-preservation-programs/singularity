package dataset

import (
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/ipfs/go-log/v2"
	"gorm.io/gorm"
)

// RemoveHandler godoc
// @Summary Remove a specific dataset. This will not remove the CAR files.
// @Description Important! If the dataset is large, this command will take some time to remove all relevant data.
// @Tags Dataset
// @Param datasetName path string true "Dataset name"
// @Success 204
// @Failure 400 {object} handler.HTTPError
// @Failure 500 {object} handler.HTTPError
// @Router /dataset/{datasetName} [delete]
func RemoveHandler(
	db *gorm.DB,
	datasetName string,
) *handler.Error {
	log.SetAllLoggers(log.LevelInfo)
	dataset, err := database.FindDatasetByName(db, datasetName)
	if err != nil {
		return handler.NewBadRequestString("failed to find dataset: " + err.Error())
	}
	err = db.Delete(&dataset).Error
	if err != nil {
		return handler.NewHandlerError(err)
	}
	return nil
}
