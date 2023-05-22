package status

import (
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/ipfs/go-log/v2"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"strconv"
)

// GetSourceChunksHandler godoc
// @Summary Get all chunk details of a data source
// @Tags Data Source
// @Accept json
// @Produce json
// @Param id path string true "Source ID"
// @Success 200 {array} model.Chunk
// @Failure 500 {object} handler.HTTPError
// @Router /source/{id}/chunks [get]
func GetSourceChunksHandler(
	db *gorm.DB,
	id string,
) ([]model.Chunk, *handler.Error) {
	log.SetAllLoggers(log.LevelInfo)
	sourceID, err := strconv.Atoi(id)
	if err != nil {
		return nil, handler.NewBadRequestString("invalid source id")
	}
	var source model.Source
	err = db.Where("id = ?", sourceID).First(&source).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, handler.NewBadRequestString("source not found")
	}
	if err != nil {
		return nil, handler.NewHandlerError(err)
	}

	var chunks []model.Chunk
	err = db.Preload("Car").Where("source_id = ?", sourceID).Find(&chunks).Error
	if err != nil {
		return nil, handler.NewHandlerError(err)
	}

	return chunks, nil
}
