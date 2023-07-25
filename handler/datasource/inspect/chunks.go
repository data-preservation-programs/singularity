package inspect

import (
	"strconv"

	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func GetSourceChunksHandler(
	db *gorm.DB,
	id string,
) ([]model.Chunk, error) {
	return getSourceChunksHandler(db, id)
}

// @Summary Get all chunk details of a data source
// @Tags Data Source
// @Accept json
// @Produce json
// @Param id path string true "Source ID"
// @Success 200 {array} model.Chunk
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /source/{id}/chunks [get]
func getSourceChunksHandler(
	db *gorm.DB,
	id string,
) ([]model.Chunk, error) {
	sourceID, err := strconv.Atoi(id)
	if err != nil {
		return nil, handler.NewInvalidParameterErr("invalid source id")
	}
	var source model.Source
	err = db.Where("id = ?", sourceID).First(&source).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, handler.NewInvalidParameterErr("source not found")
	}
	if err != nil {
		return nil, err
	}

	var chunks []model.Chunk
	err = db.Preload("Cars").Where("source_id = ?", sourceID).Find(&chunks).Error
	if err != nil {
		return nil, err
	}

	return chunks, nil
}
