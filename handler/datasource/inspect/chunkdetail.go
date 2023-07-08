package inspect

import (
	"strconv"

	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func GetSourceChunkDetailHandler(
	db *gorm.DB,
	id string,
) (*model.Chunk, error) {
	return getSourceChunkDetailHandler(db, id)
}

// @Summary Get detail of a specific chunk
// @Tags Data Source
// @Accept json
// @Produce json
// @Param id path string true "Chunk ID"
// @Success 200 {object} model.Chunk
// @Failure 500 {object} handler.HTTPError
// @Router /chunk/{id} [get]
func getSourceChunkDetailHandler(
	db *gorm.DB,
	id string,
) (*model.Chunk, error) {
	chunkID, err := strconv.Atoi(id)
	if err != nil {
		return nil, handler.NewBadRequestString("invalid chunk id")
	}
	var chunk model.Chunk
	err = db.Preload("Cars").Preload("ItemParts.Item").Where("id = ?", chunkID).First(&chunk).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, handler.NewBadRequestString("chunk not found")
	}
	if err != nil {
		return nil, handler.NewHandlerError(err)
	}

	return &chunk, nil
}
