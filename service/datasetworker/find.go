package datasetworker

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

func (w *Thread) findJob(ctx context.Context, typesOrdered []model.JobType) (*model.Job, error) {
	db := w.dbNoContext.WithContext(ctx)
	if !w.config.EnableDag {
		return nil, nil
	}

	var job model.Job
	for _, jobType := range typesOrdered {
		err := database.DoRetry(ctx, func() error {
			return db.Transaction(func(db *gorm.DB) error {
				err := db.Preload("Attachment.Preparation.OutputStorages").Preload("Attachment.Storage").
					Where("type = ? AND state = ? OR (state = ? AND worker_id is null)", jobType, model.Ready, model.Processing).
					First(&job).Error
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return nil
				}

				if err != nil {
					return errors.WithStack(err)
				}

				return db.Model(&job).
					Updates(map[string]any{
						"state":         model.Processing,
						"worker_id":     w.id,
						"error_message": "",
					}).Error
			})
		})
		if err != nil {
			return nil, errors.WithStack(err)
		}
		if job.ID == 0 {
			continue
		}
	}

	if job.ID == 0 {
		//nolint: nilnil
		return nil, nil
	}

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
