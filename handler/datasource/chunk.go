package datasource

import (
	"strconv"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type ChunkRequest struct {
	ItemIDs []uint64 `json:"itemIDs" validation:"required"`
}

func ChunkHandler(
	db *gorm.DB,
	sourceID string,
	request ChunkRequest,
) (*model.Chunk, error) {
	return chunkHandler(db, sourceID, request)
}

// @Summary Create a chunk for the specified items
// @Tags Data Source
// @Accept json
// @Produce json
// @Param id path string true "Source ID"
// @Param request body ChunkRequest true "Request body"
// @Success 201 {object} model.Item
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /source/{id}/chunk [post]
func chunkHandler(
	db *gorm.DB,
	sourceID string,
	request ChunkRequest,
) (*model.Chunk, error) {
	sourceIDInt, err := strconv.Atoi(sourceID)
	if err != nil {
		return nil, handler.NewInvalidParameterErr("invalid source id")
	}

	chunk := model.Chunk{
		SourceID:     uint32(sourceIDInt),
		PackingState: model.Ready,
	}

	err = database.DoRetry(func() error {
		return db.Transaction(
			func(db *gorm.DB) error {
				err := db.Create(&chunk).Error
				if err != nil {
					return errors.Wrap(err, "failed to create chunk")
				}
				fileRangeIDChunks := util.ChunkSlice(request.ItemIDs, util.BatchSize)
				for _, fileRangeIDChunks := range fileRangeIDChunks {
					err = db.Model(&model.FileRange{}).
						Where("id IN ?", fileRangeIDChunks).Update("chunk_id", chunk.ID).Error
					if err != nil {
						return errors.Wrap(err, "failed to update items")
					}
				}
				return nil
			},
		)
	})
	if err != nil {
		return nil, err
	}

	return &chunk, nil
}
