package datasource

import (
	"strconv"

	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type PackingManifestsByState struct {
	Count int64           `json:"count"` // number of packing manifests in this state
	State model.WorkState `json:"state"` // the state of the packing manifests
}

type FileSummary struct {
	Total    int64 `json:"total"`    // number of files in the source
	Prepared int64 `json:"prepared"` // number of files prepared
}

type SourceStatus struct {
	PackingManifestSummary []PackingManifestsByState `json:"packingManifestSummary"` // summary of the packing manifests
	FileSummary            FileSummary               `json:"fileSummary"`            // summary of the files
	FailedPackingManifests []model.PackingManifest   `json:"failedPackingManifests"` // failed packing manifests
}

// @Summary Get the data preparation summary of a data source
// @Tags Data Source
// @Accept json
// @Produce json
// @Param id path string true "Source ID"
// @Success 200 {object} PackingManifestsByState
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
	err = db.Model(&model.PackingManifest{}).
		Select("count(*) as count, packing_state as state").
		Where("source_id = ?", sourceID).
		Group("packing_state").Find(&summary.PackingManifestSummary).Error
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

	err = db.Model(&model.PackingManifest{}).Where("source_id = ? AND packing_state = ?", sourceID, model.Error).Find(&summary.FailedPackingManifests).Error
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
