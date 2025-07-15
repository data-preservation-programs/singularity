package scan

import (
	"context"
	"fmt"
	"strings"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/pack/push"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/data-preservation-programs/singularity/storagesystem"
	"gorm.io/gorm"
)

func NextAvailablePackJob(
	ctx context.Context,
	db *gorm.DB,
	attachmentID model.SourceAttachmentID,
) (*model.Job, error) {
	var packJob model.Job
	err := database.DoRetry(ctx, func() error {
		return db.Where(model.Job{AttachmentID: attachmentID, State: model.Created, Type: model.Pack}).Preload("FileRanges").FirstOrCreate(&packJob).Error
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
	var totalSize int64

	// Get storage handler to check if this is a union storage
	if attachment.Storage == nil {
		return 0, fmt.Errorf("source attachment has no storage")
	}
	handler, err := storagesystem.NewRCloneHandler(ctx, *attachment.Storage)
	if err == nil {
		// Check if we can cast the fs to a UnionFS
		if _, ok := handler.Fs().(interface{ ListUpstreams() []string }); ok {
			// This is a union storage, and onePiecePerUpstream is set
			if attachment.Preparation.OnePiecePerUpstream {
				// Group parts by upstream folder
				partsByUpstream := make(map[string][]model.FileRange)

				// Scan through all parts and group them by their upstream folder
				for _, part := range remainingParts {
					upstreamPath := getUpstreamPath(part.File.Path)
					partsByUpstream[upstreamPath] = append(partsByUpstream[upstreamPath], part)
				}

				// Create one piece per upstream folder
				for _, parts := range partsByUpstream {
					nextPackJob, err := NextAvailablePackJob(ctx, db, attachment.ID)
					if err != nil {
						return 0, fmt.Errorf("finding next available pack job: %w", err)
					}
					fileRangeSet.Reset()
					fileRangeSet.Add(nextPackJob.FileRanges...)

					// Add all parts from this upstream to the piece
					for _, part := range parts {
						fileRangeSet.Add(part)
					}

					// Update the job
					err = UpdatePackJob(ctx, db, nextPackJob.ID, model.Created, fileRangeSet.FileRangeIDs())
					if err != nil {
						return 0, fmt.Errorf("updating pack job: %w", err)
					}
					totalSize += fileRangeSet.CarSize()
				}
				return totalSize, nil
			}
		}
	}

	// Default behavior - combine parts based on size
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
		totalSize += fileRangeSet.CarSize()
	}
	return totalSize, nil
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

// getUpstreamPath returns the upstream folder path for a given file path.
// In a union storage, files are mounted under their upstream's root folder.
func getUpstreamPath(filePath string) string {
	parts := strings.Split(filePath, "/")
	if len(parts) > 0 {
		return parts[0]
	}
	return ""
}
