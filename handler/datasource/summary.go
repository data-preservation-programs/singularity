package datasource

import (
	"strconv"

	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type ChunksByState struct {
	Count int64           `json:"count"` // number of chunks in this state
	State model.WorkState `json:"state"` // the state of the chunks
}

type FileSummary struct {
	Total    int64 `json:"total"`    // number of files in the source
	Prepared int64 `json:"prepared"` // number of files prepared
}

type SourceStatus struct {
	ChunkSummary []ChunksByState `json:"chunkSummary"` // summary of the chunks
	FileSummary  FileSummary     `json:"fileSummary"`  // summary of the files
	FailedChunks []model.Chunk   `json:"failedChunks"` // failed chunks
}

// @Summary Get the data preparation summary of a data source
// @Tags Data Source
// @Accept json
// @Produce json
// @Param id path string true "Source ID"
// @Success 200 {object} ChunksByState
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /source/{id}/summary [get]
func getSourceStatusHandler(
	db *gorm.DB,
	id string,
) (*SourceStatus, error) {
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

	summary := SourceStatus{}
	err = db.Model(&model.Chunk{}).
		Select("count(*) as count, packing_state as state").
		Where("source_id = ?", sourceID).
		Group("packing_state").Find(&summary.ChunkSummary).Error
	if err != nil {
		return nil, err
	}

	err = db.Model(&model.File{}).Where("source_id = ?", sourceID).Count(&summary.FileSummary.Total).Error
	if err != nil {
		return nil, err
	}

	err = db.Model(&model.File{}).Where("source_id = ? AND cid IS NOT NULL", sourceID).Count(&summary.FileSummary.Prepared).Error
	if err != nil {
		return nil, err
	}

	err = db.Model(&model.Chunk{}).Where("source_id = ? AND packing_state = ?", sourceID, model.Error).Find(&summary.FailedChunks).Error
	if err != nil {
		return nil, err
	}

	return &summary, nil
}

func GetSourceStatusHandler(
	db *gorm.DB,
	id string,
) (*SourceStatus, error) {
	return getSourceStatusHandler(db, id)
}
