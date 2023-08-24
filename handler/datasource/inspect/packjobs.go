package inspect

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

type GetSourcePackJobsRequest struct {
	State model.WorkState `json:"state"`
}

func GetSourcePackJobsHandler(
	ctx context.Context,
	db *gorm.DB,
	id uint32,
	request GetSourcePackJobsRequest,
) ([]model.PackJob, error) {
	return getSourcePackJobsHandler(db.WithContext(ctx), id, request)
}

// @Summary Get all pack job details of a data source
// @Tags Data Source
// @Accept json
// @Produce json
// @Param id path string true "Source ID"
// @Param request body GetSourcePackJobsRequest true "GetSourcePackJobsRequest"
// @Success 200 {array} model.PackJob
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /source/{id}/packjobs [get]
func getSourcePackJobsHandler(
	db *gorm.DB,
	sourceID uint32,
	request GetSourcePackJobsRequest,
) ([]model.PackJob, error) {
	var source model.Source
	err := db.Where("id = ?", sourceID).First(&source).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, handlererror.NewInvalidParameterErr("source not found")
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var packJobs []model.PackJob
	if request.State == "" {
		err = db.Where("source_id = ?", sourceID).Find(&packJobs).Error
	} else {
		err = db.Where("source_id = ? AND packing_state = ?", sourceID, request.State).Find(&packJobs).Error
	}

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return packJobs, nil
}
