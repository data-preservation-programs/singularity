package dataprep

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"golang.org/x/exp/slices"
	"gorm.io/gorm"
)

var pausableStatesForScan = []model.JobState{model.Processing, model.Ready}

var startableStatesForScan = []model.JobState{model.Paused, model.Created, model.Error, model.Complete}

func validateSourceStorage(ctx context.Context, db *gorm.DB, id uint32, name string) (*model.Storage, error) {
	db = db.WithContext(ctx)
	var storage model.Storage
	err := db.Where("name = ?", name).First(&storage).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(handlererror.ErrNotFound, "storage '%s' does not exist", name)
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var source model.SourceAttachment
	err = db.Where("preparation_id = ? AND storage_id = ?", id, storage.ID).First(&source).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(handlererror.ErrNotFound, "source '%s' is not attached to preparation %d", name, id)
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &storage, nil
}

// StartJobHandler initializes or restarts a job for a given source storage.
//
// The function checks if there's an existing job of the given type for the source. If the job
// doesn't exist, it creates one. If the job exists and is in a startable state, it will reset
// the job to be ready to start again. If the job is already running, it returns an error.
//
// Parameters:
// - ctx: The context for database transactions and other operations.
// - db: A pointer to the gorm.DB instance representing the database connection.
// - id: The unique identifier for the desired Preparation record.
// - name: The name of the source storage.
// - jobType: The type of the job (e.g., Scan, Upload).
//
// Returns:
//   - A pointer to the model.Job record that was created or updated.
//   - An error, if any occurred during the database transaction or if the source storage doesn't exist,
//     or if there's already a running job of the specified type for the source.
//
// Note:
// The function ensures the job is either newly created or reset, and is ready to be executed by a worker.
func StartJobHandler(
	ctx context.Context,
	db *gorm.DB,
	id uint32,
	name string,
	jobType model.JobType) (*model.Job, error) {
	db = db.WithContext(ctx)
	source, err := validateSourceStorage(ctx, db, id, name)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var job model.Job
	err = database.DoRetry(ctx, func() error {
		return db.Transaction(func(db *gorm.DB) error {
			err := db.Where("type = ? AND attachment_id = ?", jobType, source.ID).First(&job).Error
			if errors.Is(err, gorm.ErrRecordNotFound) {
				job = model.Job{
					State:        model.Ready,
					Type:         jobType,
					AttachmentID: source.ID,
				}
				return errors.WithStack(db.Create(&job).Error)
			}
			if slices.Contains(startableStatesForScan, job.State) {
				return errors.WithStack(db.Model(&job).Updates(map[string]any{
					"state":             model.Ready,
					"error_message":     "",
					"error_stack_trace": "",
				}).Error)
			}

			return errors.Wrapf(handlererror.ErrInvalidParameter, "%s job for source '%s' is already running", jobType, name)
		})
	})

	return &job, errors.WithStack(err)
}

func StartScanHandler(
	ctx context.Context,
	db *gorm.DB,
	id uint32,
	name string) (*model.Job, error) {
	return StartJobHandler(ctx, db, id, name, model.Scan)
}

// @Summary Start a new scanning job
// @Tags Job
// @Accept json
// @Produce json
// @Param id path int true "Preparation ID"
// @Param name path string true "Storage name"
// @Success 200 {object} Job
// @Failure 400 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /preparation/{id}/source/{name}/start-scan [post]
func _() {}

// PauseJobHandler attempts to pause a running job for a given source storage.
//
// This function checks if there's an existing job of the given type for the source. If the job
// exists and is in a pausable state, it updates the job's state to 'Paused'. If the job doesn't
// exist or is not in a pausable state, it returns an appropriate error.
//
// Parameters:
// - ctx: The context for database transactions and other operations.
// - db: A pointer to the gorm.DB instance representing the database connection.
// - id: The unique identifier for the desired Preparation record.
// - name: The name of the source storage.
// - jobType: The type of the job (e.g., Scan, Upload).
//
// Returns:
//   - A pointer to the model.Job record that was paused.
//   - An error, if any occurred during the database transaction or if the job doesn't exist,
//     or if the job is not in a pausable state.
func PauseJobHandler(
	ctx context.Context,
	db *gorm.DB,
	id uint32,
	name string,
	jobType model.JobType) (*model.Job, error) {
	db = db.WithContext(ctx)
	source, err := validateSourceStorage(ctx, db, id, name)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var job model.Job
	err = database.DoRetry(ctx, func() error {
		return db.Transaction(func(db *gorm.DB) error {
			err := db.Where("type = ? AND attachment_id = ?", jobType, source.ID).First(&job).Error
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.Wrapf(handlererror.ErrNotFound, "%s job for source '%s' does not exist", jobType, name)
			}
			if !slices.Contains(pausableStatesForScan, job.State) {
				return errors.Wrapf(handlererror.ErrInvalidParameter, "%s job for source '%s' is not running", jobType, name)
			}

			return errors.WithStack(db.Model(&job).Update("state", model.Paused).Error)
		})
	})

	return &job, errors.WithStack(err)
}

func PauseScanHandler(
	ctx context.Context,
	db *gorm.DB,
	id uint32,
	name string) (*model.Job, error) {
	return PauseJobHandler(ctx, db, id, name, model.Scan)
}

// @Summary Pause an ongoing scanning job
// @Tags Job
// @Accept json
// @Produce json
// @Param id path int true "Preparation ID"
// @Param name path string true "Storage name"
// @Success 200 {object} Job
// @Failure 400 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /preparation/{id}/source/{name}/pause-scan [post]
func _() {}
