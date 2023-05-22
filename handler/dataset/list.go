package dataset

import (
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/ipfs/go-log/v2"
	"gorm.io/gorm"
)


// ListHandler godoc
// @Summary List all datasets
// @Tags Dataset
// @Produce json
// @Success 200 {array} model.Dataset
// @Failure 400 {object} handler.HTTPError
// @Failure 500 {object} handler.HTTPError
// @Router /datasets [get]
func ListHandler(
	db *gorm.DB,
) ([]model.Dataset, *handler.Error) {
	log.SetAllLoggers(log.LevelInfo)
	var datasets []model.Dataset
	err := db.Find(&datasets).Error
	if err != nil {
		return nil, handler.NewHandlerError(err)
	}
	return datasets, nil
}
