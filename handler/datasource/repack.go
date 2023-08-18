package datasource

import (
	"context"
	"strconv"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

// @Summary Trigger a repack of a pack job or all errored pack jobs of a data source
// @Tags Data Source
// @Produce json
// @Param id path string true "Source ID"
// @Param request body RepackRequest true "Request body"
// @Success 200 {array} model.PackJob
// @Failure 500 {object} api.HTTPError
// @Router /source/{id}/repack [post]
func repackHandler(
	ctx context.Context,
	db *gorm.DB,
	id string,
	request RepackRequest,
) ([]model.PackJob, error) {
	if id == "" && request.PackJobID == nil {
		return nil, handler.NewInvalidParameterErr("either source id or pack job id must be provided")
	}

	var sourceID int
	var err error
	if id != "" {
		sourceID, err = strconv.Atoi(id)
		if err != nil {
			return nil, handler.NewInvalidParameterErr("invalid source id")
		}
	}

	if request.PackJobID != nil {
		packJobID := *request.PackJobID
		var packJob model.PackJob
		statement := db.Where("id = ?", packJobID)
		if sourceID != 0 {
			statement = statement.Where("source_id = ?", sourceID)
		}
		err = statement.First(&packJob).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, handler.NewInvalidParameterErr("pack job not found")
		}
		if err != nil {
			return nil, errors.WithStack(err)
		}
		if packJob.PackingState == model.Error || packJob.PackingState == model.Complete {
			err = database.DoRetry(ctx, func() error {
				return db.Model(&packJob).Updates(map[string]any{
					"packing_state": model.Ready,
					"error_message": "",
				}).Error
			})
			if err != nil {
				return nil, errors.WithStack(err)
			}
		} else {
			return nil, handler.NewInvalidParameterErr("pack job is not in error or complete state")
		}
		return []model.PackJob{packJob}, nil
	}

	var packJobs []model.PackJob
	err = db.Transaction(func(db *gorm.DB) error {
		err = db.Where("source_id = ? and packing_state = ?", sourceID, model.Error).Find(&packJobs).Error
		if err != nil {
			return errors.WithStack(err)
		}
		err = db.Model(&model.PackJob{}).Where("source_id = ? and packing_state = ?", sourceID, model.Error).Updates(map[string]any{
			"packing_state": model.Ready,
			"error_message": "",
		}).Error
		if err != nil {
			return errors.WithStack(err)
		}
		for i := range packJobs {
			packJobs[i].PackingState = model.Ready
			packJobs[i].ErrorMessage = ""
		}
		return nil
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return packJobs, nil
}

type RepackRequest struct {
	PackJobID *uint64 `json:"packJobId"`
}

func RepackHandler(
	ctx context.Context,
	db *gorm.DB,
	id string,
	request RepackRequest,
) ([]model.PackJob, error) {
	return repackHandler(ctx, db.WithContext(ctx), id, request)
}
