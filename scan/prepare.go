package scan

import (
	"context"
	"fmt"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/pack/push"
	"github.com/data-preservation-programs/singularity/util"
	"gorm.io/gorm"
)

func NextAvailablePackJob(
	ctx context.Context,
	db *gorm.DB,
	attachmentID model.SourceAttachmentID,
) (*model.Job, error) {
	var packJob model.Job
	err := database.DoRetry(ctx, func() error {
		return db.Where(model.Job{AttachmentID: attachmentID, State: model.Created}).Preload("FileRanges").FirstOrCreate(&packJob).Error
	})
	return &packJob, errors.WithStack(err)
}

func PrepareToPackFileRanges(
	ctx context.Context,
	db *gorm.DB,
	attachment model.SourceAttachment,
	remainingParts []model.FileRange,
) (int64, error) {
	fileRangeSet := push.NewFileRangeSet()

	for len(remainingParts) > 0 {
		nextPackJob, err := NextAvailablePackJob(ctx, db, attachment.ID)
		if err != nil {
			return 0, fmt.Errorf("finding next available pack job: %w", err)
		}
		fileRangeSet.Reset()
		fileRangeSet.Add(nextPackJob.FileRanges...)
		for len(remainingParts) > 0 {
			if !fileRangeSet.AddIfFits(remainingParts[0], attachment.Preparation.MaxSize) {
				break
			}
			remainingParts = remainingParts[1:]
		}
		if len(fileRangeSet.FileRanges()) == 0 && len(remainingParts) > 0 {
			fileRangeSet.Add(remainingParts[:1]...)
			remainingParts = remainingParts[1:]
		}
		// if we still have remaining parts, we've filled up this chunk
		packJobState := model.Created
		if len(remainingParts) > 0 {
			packJobState = model.Ready
		}
		err = UpdatePackJob(ctx, db, nextPackJob.ID, packJobState, fileRangeSet.FileRangeIDs())
		if err != nil {
			return 0, fmt.Errorf("updating pack job: %w", err)
		}
	}
	return fileRangeSet.CarSize(), nil
}

func UpdatePackJob(
	ctx context.Context,
	db *gorm.DB,
	packJobID model.JobID,
	state model.JobState,
	fileRangeIDs []model.FileRangeID,
) error {
	return database.DoRetry(ctx, func() error {
		return db.Transaction(
			func(db *gorm.DB) error {
				err := db.Model(&model.Job{}).Where("id = ?", packJobID).Update("state", state).Error
				if err != nil {
					return fmt.Errorf("failed to update pack job: %w", err)
				}
				fileRangeIDChunks := util.ChunkSlice(fileRangeIDs, util.BatchSize)
				for _, fileRangeChunk := range fileRangeIDChunks {
					err = db.Model(&model.FileRange{}).
						Where("id IN ?", fileRangeChunk).Update("job_id", packJobID).Error
					if err != nil {
						return fmt.Errorf("failed to update items: %w", err)
					}
				}
				return nil
			},
		)
	})
}

func PrepareSource(ctx context.Context, db *gorm.DB, attachment model.SourceAttachment) error {
	db = db.WithContext(ctx)
	var remainingFileRanges []model.FileRange
	err := db.Joins("File").
		Where("attachment_id = ? AND file_ranges.job_id is null", attachment.ID).
		Order("file_ranges.id asc").
		Find(&remainingFileRanges).Error
	if err != nil {
		return errors.WithStack(err)
	}
	logger.With("remaining", len(remainingFileRanges)).Info("remaining file ranges")
	_, err = PrepareToPackFileRanges(ctx, db, attachment, remainingFileRanges)
	if err != nil {
		return errors.WithStack(err)
	}
	return markPackJobsReady(ctx, db, attachment.ID)
}

func markPackJobsReady(
	ctx context.Context,
	db *gorm.DB,
	attachmentID model.SourceAttachmentID,
) error {
	return database.DoRetry(ctx, func() error {
		return db.Model(&model.Job{}).Where("attachment_id = ? AND state = ?", attachmentID, model.Created).Update("state", model.Ready).Error
	})
}
