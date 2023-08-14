package inspect

import (
	"context"
	"strconv"

	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func GetSourceChunkDetailHandler(
	ctx context.Context,
	db *gorm.DB,
	id string,
) (*model.Chunk, error) {
	return getSourceChunkDetailHandler(db.WithContext(ctx), id)
}

// @Summary Get detail of a specific chunk
// @Tags Data Source
// @Accept json
// @Produce json
// @Param id path string true "Chunk ID"
// @Success 200 {object} model.Chunk
// @Failure 500 {object} api.HTTPError
// @Router /chunk/{id} [get]
func getSourceChunkDetailHandler(
	db *gorm.DB,
	id string,
) (*model.Chunk, error) {
	chunkID, err := strconv.Atoi(id)
	if err != nil {
		return nil, handler.NewInvalidParameterErr("invalid chunk id")
	}
	var chunk model.Chunk
	err = db.Preload("Cars").Preload("ItemParts").Where("id = ?", chunkID).First(&chunk).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, handler.NewInvalidParameterErr("chunk not found")
	}
	if err != nil {
		return nil, err
	}

	itemMap := make(map[uint64]*model.Item)
	for _, part := range chunk.ItemParts {
		itemMap[part.ItemID] = nil
	}

	itemIDChunks := util.ChunkMapKeys(itemMap, util.BatchSize)
	for _, itemIDChunk := range itemIDChunks {
		var items []model.Item
		err = db.Where("id IN ?", itemIDChunk).Find(&items).Error
		if err != nil {
			return nil, err
		}
		for i, item := range items {
			itemMap[item.ID] = &items[i]
		}
	}

	for i, part := range chunk.ItemParts {
		item, ok := itemMap[part.ItemID]
		if ok {
			chunk.ItemParts[i].Item = item
		}
	}

	return &chunk, nil
}
