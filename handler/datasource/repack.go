package datasource

import (
	"strconv"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// @Summary Trigger a repack of a chunk or all errored chunks of a data source
// @Tags Data Source
// @Produce json
// @Param id path string true "Source ID"
// @Param request body RepackRequest true "Request body"
// @Success 200 {array} model.Chunk
// @Failure 500 {object} api.HTTPError
// @Router /source/{id}/repack [post]
func repackHandler(
	db *gorm.DB,
	id string,
	request RepackRequest,
) ([]model.Chunk, error) {
	if id == "" && request.ChunkID == nil {
		return nil, handler.NewInvalidParameterErr("either source id or chunk id must be provided")
	}

	var sourceID int
	var err error
	if id != "" {
		sourceID, err = strconv.Atoi(id)
		if err != nil {
			return nil, handler.NewInvalidParameterErr("invalid source id")
		}
	}

	if request.ChunkID != nil {
		chunkID := *request.ChunkID
		var chunk model.Chunk
		statement := db.Where("id = ?", chunkID)
		if sourceID != 0 {
			statement = statement.Where("source_id = ?", sourceID)
		}
		err = statement.First(&chunk).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, handler.NewInvalidParameterErr("chunk not found")
		}
		if err != nil {
			return nil, err
		}
		if chunk.PackingState == model.Error || chunk.PackingState == model.Complete {
			err = database.DoRetry(func() error {
				return db.Model(&chunk).Updates(map[string]any{
					"packing_state": model.Ready,
					"error_message": "",
				}).Error
			})
			if err != nil {
				return nil, err
			}
		} else {
			return nil, handler.NewInvalidParameterErr("chunk is not in error or complete state")
		}
		return []model.Chunk{chunk}, nil
	}

	var chunks []model.Chunk
	err = db.Transaction(func(db *gorm.DB) error {
		err = db.Where("source_id = ? and packing_state = ?", sourceID, model.Error).Find(&chunks).Error
		if err != nil {
			return err
		}
		err = db.Model(&model.Chunk{}).Where("source_id = ? and packing_state = ?", sourceID, model.Error).Updates(map[string]any{
			"packing_state": model.Ready,
			"error_message": "",
		}).Error
		if err != nil {
			return err
		}
		for i := range chunks {
			chunks[i].PackingState = model.Ready
			chunks[i].ErrorMessage = ""
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return chunks, nil
}

type RepackRequest struct {
	ChunkID *uint64 `json:"chunkId"`
}

func RepackHandler(
	db *gorm.DB,
	id string,
	request RepackRequest,
) ([]model.Chunk, error) {
	return repackHandler(db, id, request)
}
