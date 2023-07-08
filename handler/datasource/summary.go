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

type ItemSummary struct {
	Total    int64 `json:"total"`    // number of items in the source
	Prepared int64 `json:"prepared"` // number of items prepared
}

type SourceStatus struct {
	ChunkSummary []ChunksByState `json:"chunkSummary"` // summary of the chunks
	ItemSummary  ItemSummary     `json:"itemSummary"`  // summary of the items
}

// @Summary Get the data preparation summary of a data source
// @Tags Data Source
// @Accept json
// @Produce json
// @Param id path string true "Source ID"
// @Success 200 {object} ChunksByState
// @Failure 400 {object} handler.HTTPError
// @Failure 500 {object} handler.HTTPError
// @Router /source/{id}/summary [get]
func getSourceStatusHandler(
	db *gorm.DB,
	id string,
) (*SourceStatus, error) {
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

	summary := SourceStatus{}
	err = db.Model(&model.Chunk{}).
		Select("count(*) as count, packing_state as state").
		Where("source_id = ?", sourceID).
		Group("packing_state").Find(&summary.ChunkSummary).Error
	if err != nil {
		return nil, handler.NewHandlerError(err)
	}

	err = db.Model(&model.Item{}).Where("source_id = ?", sourceID).Count(&summary.ItemSummary.Total).Error
	if err != nil {
		return nil, handler.NewHandlerError(err)
	}

	err = db.Model(&model.Item{}).Where("source_id = ? AND cid IS NOT NULL", sourceID).Count(&summary.ItemSummary.Prepared).Error
	if err != nil {
		return nil, handler.NewHandlerError(err)
	}

	return &summary, nil
}

func GetSourceStatusHandler(
	db *gorm.DB,
	id string,
) (*SourceStatus, error) {
	return getSourceStatusHandler(db, id)
}
