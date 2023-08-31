package dataprep

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util"
	"golang.org/x/exp/slices"
	"gorm.io/gorm"
)

var startableStatesForPack = []model.JobState{model.Paused, model.Created, model.Error}
var pausableStatesForPack = []model.JobState{model.Processing, model.Ready}

// StartPackHandler initiates pack jobs for a given source storage.
//
// If jobID is provided, this function will attempt to start a specific pack job. If not,
// it will search for all pack jobs in startable states associated with the source and attempt
// to start them. The state of the job will be updated to 'Ready'.
//
// Parameters:
//   - ctx: The context for database transactions and other operations.
//   - db: A pointer to the gorm.DB instance representing the database connection.
//   - id: The unique identifier for the desired Preparation record.
//   - name: The name of the source storage.
//   - jobID: The unique identifier of the job to be started. If set to 0, all eligible jobs
//     for the source will be initiated.
//
// Returns:
//   - A slice of model.Job records that were started.
//   - An error, if any occurred during the database transaction or if the job doesn't exist,
//     or if the job is not in a startable state.
func (DefaultHandler) StartPackHandler(
	ctx context.Context,
	db *gorm.DB,
	id uint32,
	name string,
	jobID int64) ([]model.Job, error) {
	db = db.WithContext(ctx)
	sourceAttachment, err := validateSourceStorage(ctx, db, id, name)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var jobs []model.Job
	if jobID == 0 {
		err = database.DoRetry(ctx, func() error {
			err := db.Where("type = ? AND state in ? AND attachment_id = ?", model.Pack, startableStatesForPack, sourceAttachment.ID).Find(&jobs).Error
			if err != nil {
				return errors.WithStack(err)
			}
			var jobIDs []uint64
			for i, job := range jobs {
				jobIDs = append(jobIDs, job.ID)
				jobs[i].State = model.Ready
			}
			jobIDChunks := util.ChunkSlice(jobIDs, util.BatchSize)
			for _, jobIDs := range jobIDChunks {
				err = db.Model(&model.Job{}).Where("id IN ?", jobIDs).Updates(map[string]any{
					"state":             model.Ready,
					"error_message":     "",
					"error_stack_trace": "",
				}).Error
				if err != nil {
					return errors.WithStack(err)
				}
			}
			return nil
		})
		if err != nil {
			return nil, errors.WithStack(err)
		}
		return jobs, nil
	}

	var job model.Job
	err = database.DoRetry(ctx, func() error {
		return db.Transaction(func(db *gorm.DB) error {
			err := db.Where("id = ? AND type = ? AND attachment_id = ?", jobID, model.Pack, sourceAttachment.ID).First(&job).Error
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.Wrapf(handlererror.ErrNotFound, "pack job %d for sourceAttachment '%s' does not exist", jobID, name)
			}
			if slices.Contains(startableStatesForPack, job.State) {
				return errors.WithStack(db.Model(&job).Update("state", model.Ready).Error)
			}

			return errors.Wrapf(handlererror.ErrInvalidParameter, "pack job %d for sourceAttachment '%s' is running or complete", jobID, name)
		})
	})

	return []model.Job{job}, errors.WithStack(err)
}

// @Summary Start or restart a specific packing job
// @Tags Job
// @Accept json
// @Produce json
// @Param id path int true "Preparation ID"
// @Param name path string true "Storage name"
// @Param job_id path int true "Pack Job ID"
// @Success 200 {object} model.Job
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /preparation/{id}/source/{name}/start-pack/{job_id} [post]
func _() {}

// @Summary Start or restart all packing job
// @Tags Job
// @Accept json
// @Produce json
// @Param id path int true "Preparation ID"
// @Param name path string true "Storage name"
// @Success 200 {object} model.Job
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /preparation/{id}/source/{name}/start-pack [post]
func _() {}

// @Summary Pause a specific packing job
// @Tags Job
// @Accept json
// @Produce json
// @Param id path int true "Preparation ID"
// @Param name path string true "Storage name"
// @Param job_id path int true "Pack Job ID"
// @Success 200 {object} model.Job
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /preparation/{id}/source/{name}/pause-pack/{job_id} [post]
func _() {}

// @Summary Pause all packing job
// @Tags Job
// @Accept json
// @Produce json
// @Param id path int true "Preparation ID"
// @Param name path string true "Storage name"
// @Success 200 {object} model.Job
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /preparation/{id}/source/{name}/pause-pack [post]
func _() {}

// PausePackHandler attempts to pause pack jobs for a given source storage.
//
// If jobID is provided, the function will attempt to pause a specific pack job. If not,
// it will search for all pack jobs in pausable states associated with the source and attempt
// to pause them. The state of the job will be updated to 'Paused'.
//
// Parameters:
//   - ctx: The context for database transactions and other operations.
//   - db: A pointer to the gorm.DB instance representing the database connection.
//   - id: The unique identifier for the desired Preparation record.
//   - name: The name of the source storage.
//   - jobID: The unique identifier of the job to be paused. If set to 0, all eligible jobs
//     for the source will be paused.
//
// Returns:
//   - A slice of model.Job records that were paused.
//   - An error, if any occurred during the database transaction, if the job doesn't exist,
//     or if the job is not in a pausable state.
func (DefaultHandler) PausePackHandler(
	ctx context.Context,
	db *gorm.DB,
	id uint32,
	name string,
	jobID int64) ([]model.Job, error) {
	db = db.WithContext(ctx)
	sourceAttachment, err := validateSourceStorage(ctx, db, id, name)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var jobs []model.Job
	if jobID == 0 {
		err = database.DoRetry(ctx, func() error {
			err := db.Where("type = ? AND state in ? AND attachment_id = ?", model.Pack, pausableStatesForPack, sourceAttachment.ID).Find(&jobs).Error
			if err != nil {
				return errors.WithStack(err)
			}
			var jobIDs []uint64
			for i, job := range jobs {
				jobIDs = append(jobIDs, job.ID)
				jobs[i].State = model.Paused
			}
			jobIDChunks := util.ChunkSlice(jobIDs, util.BatchSize)
			for _, jobIDs := range jobIDChunks {
				err = db.Model(&model.Job{}).Where("id IN ?", jobIDs).Update("state", model.Paused).Error
				if err != nil {
					return errors.WithStack(err)
				}
			}
			return nil
		})
		if err != nil {
			return nil, errors.WithStack(err)
		}
		return jobs, nil
	}

	var job model.Job
	err = database.DoRetry(ctx, func() error {
		return db.Transaction(func(db *gorm.DB) error {
			err := db.Where("id = ? AND type = ? AND attachment_id = ?", jobID, model.Pack, sourceAttachment.ID).First(&job).Error
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.Wrapf(handlererror.ErrNotFound, "pack job %d for sourceAttachment '%s' does not exist", jobID, name)
			}
			if slices.Contains(pausableStatesForPack, job.State) {
				return errors.WithStack(db.Model(&job).Update("state", model.Paused).Error)
			}

			return errors.Wrapf(handlererror.ErrInvalidParameter, "pack job %d for sourceAttachment '%s' is not running", jobID, name)
		})
	})

	return []model.Job{job}, errors.WithStack(err)
}
