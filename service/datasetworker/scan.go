package datasetworker

import (
	"context"
	"strconv"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/datasource"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/pack/push"
	"github.com/data-preservation-programs/singularity/storagesystem"
	"github.com/rjNemo/underscore"
)

// scan scans the data source and inserts the model.Job back to database
// scanSource is true if the source will be actually scanned in addition to just picking up remaining ones
// resume is true if the scan will be resumed from the last scanned file, which is useful for resuming a failed scan
func (w *Thread) scan(ctx context.Context, job model.Job, scanSource bool) error {
	db := w.dbNoContext.WithContext(ctx)
	directoryCache := make(map[string]uint64)
	var remaining = newRemain()
	var remainingFileRanges []model.FileRange
	err := db.Joins("File").
		Where("attachment_id = ? AND file_ranges.job_id is null", job.AttachmentID).
		Order("file_ranges.id asc").
		Find(&remainingFileRanges).Error
	if err != nil {
		return errors.WithStack(err)
	}
	w.logger.With("remaining", len(remainingFileRanges)).Info("remaining file ranges")
	remaining.add(remainingFileRanges)

	if !scanSource {
		for len(remaining.fileRanges) > 0 {
			err = w.chunkOnce(ctx, *job.Attachment, remaining)
			if err != nil {
				return errors.Wrap(err, "failed to save packJobing")
			}
		}
		return nil
	}

	sourceScanner, err := storagesystem.NewRCloneHandler(ctx, *job.Attachment.Storage)
	if err != nil {
		return errors.WithStack(err)
	}
	entryChan := sourceScanner.Scan(ctx, "", job.Attachment.LastScannedPath)
	for entry := range entryChan {
		if entry.Error != nil {
			w.logger.Errorw("failed to scan", "error", entry.Error)
			continue
		}

		file, fileRanges, err := push.PushFile(ctx, w.dbNoContext, entry.Info, *job.Attachment, directoryCache)
		if err != nil {
			return errors.Wrapf(err, "failed to push file %s", entry.Info.Remote())
		}
		if file == nil {
			w.logger.Infow("file already exists", "path", entry.Info.Remote())
			continue
		}
		err = database.DoRetry(ctx, func() error {
			return db.Model(&model.SourceAttachment{}).Where("id = ?", job.AttachmentID).
				Update("last_scanned_path", file.Path).Error
		})
		if err != nil {
			return errors.Wrapf(err, "failed to update last scanned path to %s", file.Path)
		}

		remaining.add(fileRanges)
		for remaining.carSize >= job.Attachment.Preparation.MaxSize {
			err = w.chunkOnce(ctx, *job.Attachment, remaining)
			if err != nil {
				return errors.WithStack(err)
			}
		}
	}

	for len(remaining.fileRanges) > 0 {
		err = w.chunkOnce(ctx, *job.Attachment, remaining)
		if err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}

func (w *Thread) chunkOnce(
	ctx context.Context,
	attachment model.SourceAttachment,
	remaining *remain,
) error {
	// If everything fit, create a packJob. Usually this is the case for the last packJob
	if remaining.carSize <= attachment.Preparation.MaxSize {
		w.logger.Debugw("creating packJob", "size", remaining.carSize)
		_, err := datasource.CreateJobHandler(ctx, w.dbNoContext, strconv.FormatUint(uint64(attachment.ID), 10), datasource.CreatePackJobRequest{
			FileRangeIDs: remaining.fileRangeIDs(),
		})

		if err != nil {
			return errors.WithStack(err)
		}
		remaining.reset()
		return nil
	}
	// size > maxSize, first, find the first file range that makes it larger than maxSize
	s := remaining.carSize
	si := len(remaining.fileRanges) - 1
	for si >= 0 {
		s -= toCarSize(remaining.fileRanges[si].Length)
		if s <= attachment.Preparation.MaxSize {
			break
		}
		si--
	}

	// In case si == 0, this is the case where a single item is more than sector size for encryption
	// We will allow a single item to be more than sector size and handle it later during packing
	if si == 0 {
		si = 1
		s += toCarSize(remaining.fileRanges[0].Length)
	}

	// create a packJob for [0:si)
	w.logger.Debugw("creating packJob", "size", s)

	fileRangeIDs := underscore.Map(remaining.fileRanges[:si], func(fileRange model.FileRange) uint64 {
		return fileRange.ID
	})
	_, err := push.CreatePackJob(ctx, w.dbNoContext, attachment.ID, fileRangeIDs)
	if err != nil {
		return errors.WithStack(err)
	}
	remaining.fileRanges = remaining.fileRanges[si:]
	remaining.carSize = remaining.carSize - s + carHeaderSize
	return nil
}

type remain struct {
	fileRanges []model.FileRange
	carSize    int64
}

const carHeaderSize = 59

func newRemain() *remain {
	return &remain{
		fileRanges: make([]model.FileRange, 0),
		// Some buffer for header
		carSize: carHeaderSize,
	}
}

func (r *remain) add(fileRanges []model.FileRange) {
	r.fileRanges = append(r.fileRanges, fileRanges...)
	for _, fileRange := range fileRanges {
		r.carSize += toCarSize(fileRange.Length)
	}
}

func (r *remain) reset() {
	r.fileRanges = make([]model.FileRange, 0)
	r.carSize = carHeaderSize
}

func (r *remain) fileRangeIDs() []uint64 {
	return underscore.Map(r.fileRanges, func(fileRange model.FileRange) uint64 {
		return fileRange.ID
	})
}

func toCarSize(size int64) int64 {
	out := size
	nBlocks := size / 1024 / 1024
	if size%(1024*1024) != 0 {
		nBlocks++
	}

	// For each block, we need to add the bytes for the CID as well as varint
	out += nBlocks * (36 + 9)

	// For every 256 blocks, we need to add another block.
	// The block stores up to 256 CIDs and integers, estimate it to be 12kb
	if nBlocks > 1 {
		out += (((nBlocks - 1) / 256) + 1) * 12000
	}

	return out
}
