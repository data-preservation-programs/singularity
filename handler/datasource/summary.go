package datasource

import (
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"strconv"
)

type ChunksByState struct {
	Count int64           `json:"count"` // number of chunks in this state
	State model.WorkState `json:"state"` // the state of the chunks
}

type ItemSummary struct {
	NumOfItems int64 `json:"numOfItems"` // number of items in the source
}

type SourceStatusSummary struct {
	ChunkSummary []ChunksByState `json:"chunkSummary"` // summary of the chunks
	ItemSummary  ItemSummary     `json:"itemSummary"`  // summary of the items
	Source       model.Source    `json:"source"`       // the source
}

// GetSourceSummaryHandler godoc
// @Summary Get the data preparation summary of a data source
// @Tags Data Source
// @Accept json
// @Produce json
// @Param id path string true "Source ID"
// @Success 200 {object} ChunksByState
// @Failure 500 {object} handler.HTTPError
// @Router /source/{id}/summary [get]
func GetSourceSummaryHandler(
	db *gorm.DB,
	id string,
) (*SourceStatusSummary, *handler.Error) {
	sourceID, err := strconv.Atoi(id)
	if err != nil {
		return nil, handler.NewBadRequestString("invalid source id")
	}
	var source model.Source
	err = db.Preload("RootDirectory").Where("id = ?", sourceID).First(&source).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, handler.NewBadRequestString("source not found")
	}
	if err != nil {
		return nil, handler.NewHandlerError(err)
	}

	summary := SourceStatusSummary{
		Source: source,
	}
	err = db.Model(&model.Chunk{}).
		Select("count(*) as count, packing_state as state").
		Where("source_id = ?", sourceID).
		Group("packing_state").Find(&summary.ChunkSummary).Error
	if err != nil {
		return nil, handler.NewHandlerError(err)
	}

	err = db.Model(&model.Item{}).Where("source_id = ?", sourceID).Count(&summary.ItemSummary.NumOfItems).Error
	if err != nil {
		return nil, handler.NewHandlerError(err)
	}

	return &summary, nil
}
