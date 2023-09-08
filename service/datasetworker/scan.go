package datasetworker

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/pack/push"
	"github.com/data-preservation-programs/singularity/storagesystem"
	"github.com/gotidy/ptr"
	"gorm.io/gorm"
)

// scan scans the data source and inserts the model.Job back to database
// resume is true if the scan will be resumed from the last scanned file, which is useful for resuming a failed scan
func (w *Thread) scan(ctx context.Context, attachment model.SourceAttachment) error {
	db := w.dbNoContext.WithContext(ctx)
	directoryCache := make(map[string]uint64)
	var remaining = push.NewFileRangeSet()
	var remainingFileRanges []model.FileRange
	err := db.Joins("File").
		Where("attachment_id = ? AND file_ranges.job_id is null", attachment.ID).
		Order("file_ranges.id asc").
		Find(&remainingFileRanges).Error
	if err != nil {
		return errors.WithStack(err)
	}
	w.logger.With("remaining", len(remainingFileRanges)).Info("remaining file ranges")
	err = addFileRangesAndCreatePackJob(ctx, db, attachment.ID, remaining, attachment.Preparation.MaxSize, remainingFileRanges...)
	if err != nil {
		return errors.WithStack(err)
	}

	sourceScanner, err := storagesystem.NewRCloneHandler(ctx, *attachment.Storage)
	if err != nil {
		return errors.WithStack(err)
	}
	entryChan := sourceScanner.Scan(ctx, "", attachment.LastScannedPath)
	var lastScannedPath *string
	defer func() {
		if lastScannedPath != nil {
			err = database.DoRetry(ctx, func() error {
				return db.Model(&model.SourceAttachment{}).Where("id = ?", attachment.ID).
					Update("last_scanned_path", lastScannedPath).Error
			})
			if err != nil {
				w.logger.Errorw("failed to update last scanned path", "error", err)
			}
		}
	}()
	for entry := range entryChan {
		if entry.Error != nil {
			w.logger.Errorw("failed to scan", "error", entry.Error)
			continue
		}

		file, fileRanges, err := push.PushFile(ctx, w.dbNoContext, entry.Info, attachment, directoryCache)
		if err != nil {
			return errors.Wrapf(err, "failed to push file %s", entry.Info.Remote())
		}
		if file == nil {
			w.logger.Infow("file already exists", "path", entry.Info.Remote())
			continue
		}

		lastScannedPath = &file.Path

		err = addFileRangesAndCreatePackJob(ctx, db, attachment.ID, remaining, attachment.Preparation.MaxSize, fileRanges...)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	lastScannedPath = ptr.Of("")

	if len(remaining.FileRanges()) > 0 {
		err = createPackJob(ctx, db, attachment.ID, remaining)
		if err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}

func createPackJob(
	ctx context.Context,
	db *gorm.DB,
	attachmentID uint32,
	remaining *push.FileRangeSet,
) error {
	job, err := push.CreatePackJob(ctx, db, attachmentID, remaining.FileRangeIDs())
	if err != nil {
		return errors.WithStack(err)
	}
	logger.Info("created pack job %d with %d file ranges", job.ID, len(remaining.FileRanges()))
	remaining.Reset()
	return nil
}

func addFileRangesAndCreatePackJob(
	ctx context.Context,
	db *gorm.DB,
	attachmentID uint32,
	remaining *push.FileRangeSet,
	maxSize int64,
	fileRanges ...model.FileRange) error {
	for _, fileRange := range fileRanges {
		fit := remaining.AddIfFits(fileRange, maxSize)
		if fit {
			continue
		}
		err := createPackJob(ctx, db, attachmentID, remaining)
		if err != nil {
			return errors.WithStack(err)
		}
		remaining.Add(fileRange)
	}

	return nil
}
