package datasource

import (
	"strconv"

	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type PackJobsByState struct {
	Count int64           `json:"count"` // number of pack jobs in this state
	State model.WorkState `json:"state"` // the state of the pack jobs
}

type FileSummary struct {
	Total    int64 `json:"total"`    // number of files in the source
	Prepared int64 `json:"prepared"` // number of files prepared
}

type SourceStatus struct {
	PackJobSummary []PackJobsByState `json:"packJobSummary"` // summary of the pack jobs
	FileSummary    FileSummary       `json:"fileSummary"`    // summary of the files
	FailedPackJobs []model.PackJob   `json:"failedPackJobs"` // failed pack jobs
}

// @Summary Get the data preparation summary of a data source
// @Tags Data Source
// @Accept json
// @Produce json
// @Param id path string true "Source ID"
// @Success 200 {object} PackJobsByState
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
	err = db.Model(&model.PackJob{}).
		Select("count(*) as count, packing_state as state").
		Where("source_id = ?", sourceID).
		Group("packing_state").Find(&summary.PackJobSummary).Error
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

	err = db.Model(&model.PackJob{}).Where("source_id = ? AND packing_state = ?", sourceID, model.Error).Find(&summary.FailedPackJobs).Error
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
