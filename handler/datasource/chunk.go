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
) error {
	return chunkHandler(db, sourceID, request)
}

func chunkHandler(
	db *gorm.DB,
	sourceID string,
	request ChunkRequest,
) error {
	return database.DoRetry(func() error {
		sourceIDInt, err := strconv.Atoi(sourceID)
		if err != nil {
			return handler.NewInvalidParameterErr("invalid source id")
		}

		return db.Transaction(
			func(db *gorm.DB) error {
				chunk := model.Chunk{
					SourceID:     uint32(sourceIDInt),
					PackingState: model.Ready,
				}
				err := db.Create(&chunk).Error
				if err != nil {
					return errors.Wrap(err, "failed to create chunk")
				}
				itemPartIDChunks := util.ChunkSlice(request.ItemIDs, util.BatchSize)
				for _, itemPartIDChunks := range itemPartIDChunks {
					err = db.Model(&model.ItemPart{}).
						Where("id IN ?", itemPartIDChunks).Update("chunk_id", chunk.ID).Error
					if err != nil {
						return errors.Wrap(err, "failed to update items")
					}
				}
				return nil
				return nil
			},
		)
	})
}
