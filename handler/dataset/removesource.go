package dataset

import (
	"github.com/data-preservation-programs/go-singularity/handler"
	"github.com/data-preservation-programs/go-singularity/model"
	"github.com/ipfs/go-log/v2"
	"gorm.io/gorm"
)

func RemoveSourceHandler(
	db *gorm.DB,
	datasetName string,
	sourcePath string,
)  *handler.Error {
	log.SetAllLoggers(log.LevelInfo)
	var source model.Source
	err := db.Where("dataset_id in (?) AND path = ?", db.Table("datasets").Select("id").Where("name = ?", datasetName), sourcePath).First(&source).Error
	if err != nil {
		return handler.NewBadRequestError(err)
	}
	err = db.Delete(&source).Error
	if err != nil {
		return handler.NewHandlerError(err)
	}
	return nil
}
