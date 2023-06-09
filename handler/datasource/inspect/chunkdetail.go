package inspect

import (
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"strconv"
)

// GetSourceChunkDetailHandler godoc
// @Summary Get detail of a specific chunk
// @Tags Data Source
// @Accept json
// @Produce json
// @Param id path string true "Chunk ID"
// @Success 200 {object} model.Chunk
// @Failure 500 {object} handler.HTTPError
// @Router /source/{id}/chunks [get]
func GetSourceChunkDetailHandler(
	db *gorm.DB,
	id string,
) (*model.Chunk, *handler.Error) {
	chunkID, err := strconv.Atoi(id)
	if err != nil {
		return nil, handler.NewBadRequestString("invalid chunk id")
	}
	var chunk model.Chunk
	err = db.Preload("Car").Preload("ItemParts.Item").Where("id = ?", chunkID).First(&chunk).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, handler.NewBadRequestString("chunk not found")
	}
	if err != nil {
		return nil, handler.NewHandlerError(err)
	}

	return &chunk, nil
}
