package inspect

import (
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type GetSourceChunksRequest struct {
	State model.WorkState `json:"state"`
}

func GetSourceChunksHandler(
	db *gorm.DB,
	id uint32,
	request GetSourceChunksRequest,
) ([]model.Chunk, error) {
	return getSourceChunksHandler(db, id, request)
}

// @Summary Get all chunk details of a data source
// @Tags Data Source
// @Accept json
// @Produce json
// @Param id path string true "Source ID"
// @Param request body GetSourceChunksRequest true "GetSourceChunksRequest"
// @Success 200 {array} model.Chunk
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /source/{id}/chunks [get]
func getSourceChunksHandler(
	db *gorm.DB,
	sourceID uint32,
	request GetSourceChunksRequest,
) ([]model.Chunk, error) {
	var source model.Source
	err := db.Where("id = ?", sourceID).First(&source).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, handler.NewInvalidParameterErr("source not found")
	}
	if err != nil {
		return nil, err
	}

	var chunks []model.Chunk
	if request.State == "" {
		err = db.Preload("Cars").Where("source_id = ?", sourceID).Find(&chunks).Error
	} else {
		err = db.Preload("Cars").Where("source_id = ? AND packing_state = ?", sourceID, request.State).Find(&chunks).Error
	}

	if err != nil {
		return nil, err
	}

	return chunks, nil
}
