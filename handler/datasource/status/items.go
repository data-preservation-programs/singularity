package status

import (
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/ipfs/go-log/v2"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"strconv"
)

// GetSourceItemsHandler godoc
// @Summary Get all item details of a data source
// @Tags Data Source
// @Accept json
// @Produce json
// @Param id path string true "Source ID"
// @Param chunk_id query string false "Chunk ID"
// @Success 200 {array} model.Item
// @Failure 500 {object} handler.HTTPError
// @Router /source/{id}/items [get]
func GetSourceItemsHandler(
	db *gorm.DB,
	id string,
	chunkID string,
) ([]model.Item, *handler.Error) {
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

	var items []model.Item
	if chunkID == "" {
		err = db.Where("source_id = ?", sourceID).Find(&items).Error
	} else {
		var c int
		c, err = strconv.Atoi(chunkID)
		if err != nil {
			return nil, handler.NewBadRequestString("invalid chunk id")
		}
		err = db.Where("source_id = ? AND chunk_id = ?", sourceID, c).Find(&items).Error
	}

	if err != nil {
		return nil, handler.NewHandlerError(err)
	}

	return items, nil
}
