package datasetworker

import (
	"context"
	"database/sql"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

// findJob searches for a Job from the database based on the ordered list of job types provided.
// It iterates through the typesOrdered list, and for each type, it attempts to find a Job of that type which is
// either Ready or is marked as Processing but hasn't been claimed by any worker yet. Once a suitable Job is found,
// it marks that Job as being processed by the current worker thread.
//
// Parameters:
//   - ctx: The context which controls the lifetime of the operation.
//   - typesOrdered: A slice of model.JobType values representing the job types to search for in order of preference.
//
// Returns:
//   - A pointer to the found model.Job instance or nil if no suitable Job was found.
//   - An error, if any occurred during the operation.
func (w *Thread) findJob(ctx context.Context, typesOrdered []model.JobType) (*model.Job, error) {
	db := w.dbNoContext.WithContext(ctx)

	txOpts := &sql.TxOptions{
		Isolation: sql.LevelSerializable,
	}
	var job model.Job
	for _, jobType := range typesOrdered {
		err := database.DoRetry(ctx, func() error {
			return db.Transaction(func(db *gorm.DB) error {
				err := db.Preload("Attachment.Preparation.OutputStorages").Preload("Attachment.Storage").
					Where("type = ? AND (state = ? OR (state = ? AND worker_id IS NULL))", jobType, model.Ready, model.Processing).
					First(&job).Error
				if err != nil {
					if errors.Is(err, gorm.ErrRecordNotFound) {
						job.ID = 0
						return nil
					}
					return errors.WithStack(err)
				}

				return db.Model(&job).
					Updates(map[string]any{
						"state":         model.Processing,
						"worker_id":     w.id,
						"error_message": "",
					}).Error
			}, txOpts)
		})
		if err != nil {
			return nil, errors.WithStack(err)
		}
		if job.ID != 0 {
			break
		}
	}

	if job.ID == 0 {
		//nolint: nilnil
		return nil, nil
	}

	w.logger.Debugw("found job", "jobID", job.ID, "jobType", job.Type, "workerID", w.id)

	if job.Type == model.Pack {
		var fileRanges []model.FileRange
		err := db.Joins("File").Where("file_ranges.job_id = ?", job.ID).Order("file_ranges.id asc").Find(&fileRanges).Error
		if err != nil {
			return nil, errors.WithStack(err)
		}
		job.FileRanges = fileRanges
	}
	return &job, nil
}
