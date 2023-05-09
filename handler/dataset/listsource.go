package dataset

import (
	"github.com/data-preservation-programs/go-singularity/handler"
	"github.com/data-preservation-programs/go-singularity/model"
	"github.com/ipfs/go-log/v2"
	"gorm.io/gorm"
)

func ListSourceHandler(
	db *gorm.DB,
	datasetName string,
) ([]model.Source, *handler.Error) {
	log.SetAllLoggers(log.LevelInfo)
	var sources []model.Source
	err := db.Where("dataset_id in (?)", db.Table("datasets").Select("id").Where("name = ?", datasetName)).Find(&sources).Error
	if err != nil {
		return nil, handler.NewHandlerError(err)
	}
	return sources, nil
}
