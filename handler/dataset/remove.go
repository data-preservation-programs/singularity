package dataset

import (
	"github.com/data-preservation-programs/go-singularity/handler"
	"github.com/data-preservation-programs/go-singularity/model"
	"github.com/ipfs/go-log/v2"
	"gorm.io/gorm"
)

// RemoveHandler godoc
// @Summary Remove a dataset
func RemoveHandler(
	db *gorm.DB,
	datasetName string,
) *handler.Error {
	log.SetAllLoggers(log.LevelInfo)
	var dataset model.Dataset
	err := db.Where("name = ?", datasetName).First(&dataset).Error
	if err != nil {
		return handler.NewBadRequestError(err)
	}
	err = db.Delete(&dataset).Error
	if err != nil {
		return handler.NewHandlerError(err)
	}
	return nil
}
