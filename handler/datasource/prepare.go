package datasource

import (
	"context"
	"errors"
	"fmt"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/datasource"
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/rjNemo/underscore"
	"gorm.io/gorm"
)

func PrepareToPackFileHandler(
	ctx context.Context,
	db *gorm.DB,
	fileID uint64,
) (int64, error) {
	return prepareToPackFileHandler(ctx, db, fileID)
}

// @Summary prepare packjobs for a given item
// @Tags Data Source
// @Accept json
// @Produce json
// @Param id path integer true "File ID"
// @Success 201 {object} int64
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /file/{id}/prepare_to_pack [post]
func prepareToPackFileHandler(
	ctx context.Context,
	db *gorm.DB,
	fileID uint64,
) (int64, error) {
	var file model.File
	err := db.Preload("Source.Dataset").Where("id = ?", fileID).First(&file).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, handler.NewInvalidParameterErr("file not found")
	}
	if err != nil {
		return 0, err
	}

	var remainingParts []model.FileRange
	err = db.Where("file_id = ? AND pack_job_id is null", fileID).Order("file_ranges.id asc").
		Find(&remainingParts).Error
	if err != nil {
		return 0, err
	}

	return prepareToPackFileRanges(ctx, db, file.SourceID, file.Source.Dataset.MaxSize, remainingParts)
}

func prepareToPackFileRanges(
	ctx context.Context,
	db *gorm.DB,
	sourceID uint32,
	maxPackJobSize int64,
	remainingParts []model.FileRange,
) (int64, error) {
	fileRangeSet := newFileRangeSet()

	for len(remainingParts) > 0 {
		nextPackJob, err := nextAvailablePackJob(db, sourceID)
		if err != nil {
			return 0, fmt.Errorf("finding next available pack job: %w", err)
		}
		fileRangeSet.reset()
		fileRangeSet.add(nextPackJob.FileRanges)
		for len(remainingParts) > 0 {
			if !fileRangeSet.addIfFits(remainingParts[0], maxPackJobSize) {
				break
			}
			remainingParts = remainingParts[1:]
		}
		if len(fileRangeSet.fileRanges) == 0 && len(remainingParts) > 0 {
			fileRangeSet.add(remainingParts[:1])
			remainingParts = remainingParts[1:]
		}
		// if we still have remaining parts, we've filled up this chunk
		packJobState := model.Created
		if len(remainingParts) > 0 {
			packJobState = model.Ready
		}
		err = updatePackJob(ctx, db, nextPackJob.ID, packJobState, fileRangeSet.fileRangeIDs())
		if err != nil {
			return 0, fmt.Errorf("updating pack job: %w", err)
		}
	}
	return fileRangeSet.carSize, nil
}

func PrepareToPackSourceHandler(
	ctx context.Context,
	db *gorm.DB,
	datasourceHandlerResolver datasource.HandlerResolver,
	sourceID uint32,
) error {
	return prepareToPackSourceHandler(ctx, db, datasourceHandlerResolver, sourceID)
}

// @Summary prepare to pack a data source
// @Tags Data Source
// @Accept json
// @Produce json
// @Param id path integer true "Source ID"
// @Success 201
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /source/{id}/finalize [post]
func prepareToPackSourceHandler(
	ctx context.Context,
	db *gorm.DB,
	datasourceHandlerResolver datasource.HandlerResolver,
	sourceID uint32,
) error {
	var source model.Source
	err := db.Joins("Dataset").Where("sources.id = ?", sourceID).First(&source).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return handler.NewInvalidParameterErr(fmt.Sprintf("source %d not found.", sourceID))
	}
	if err != nil {
		return err
	}

	return PrepareSource(ctx, db, datasourceHandlerResolver, source, false)
}
func markPackJobsReady(
	ctx context.Context,
	db *gorm.DB,
	sourceID uint32,
) error {
	return database.DoRetry(ctx, func() error {
		return db.Model(&model.PackJob{}).Where("source_id = ? AND packing_state = ?", sourceID, model.Created).Update("packing_state", model.Ready).Error
	})
}

func updatePackJob(
	ctx context.Context,
	db *gorm.DB,
	packJobID uint32,
	state model.WorkState,
	fileRangeIDs []uint64,
) error {
	return database.DoRetry(ctx, func() error {
		return db.Transaction(
			func(db *gorm.DB) error {
				err := db.Model(&model.PackJob{}).Where("id = ?", packJobID).Update("packing_state", state).Error
				if err != nil {
					return fmt.Errorf("failed to update pack job: %w", err)
				}
				fileRangeIDChunks := util.ChunkSlice(fileRangeIDs, util.BatchSize)
				for _, fileRangeChunk := range fileRangeIDChunks {
					err = db.Model(&model.FileRange{}).
						Where("id IN ?", fileRangeChunk).Update("pack_job_id", packJobID).Error
					if err != nil {
						return fmt.Errorf("failed to update items: %w", err)
					}
				}
				return nil
			},
		)
	})
}

type fileRangeSet struct {
	fileRanges []model.FileRange
	carSize    int64
}

const carHeaderSize = 59

func newFileRangeSet() *fileRangeSet {
	return &fileRangeSet{
		fileRanges: make([]model.FileRange, 0),
		// Some buffer for header
		carSize: carHeaderSize,
	}
}

func (frs *fileRangeSet) add(fileRanges []model.FileRange) {
	frs.fileRanges = append(frs.fileRanges, fileRanges...)
	for _, fileRanges := range fileRanges {
		frs.carSize += toCarSize(fileRanges.Length)
	}
}

func (frs *fileRangeSet) addIfFits(fileRange model.FileRange, maxSize int64) bool {
	nextSize := toCarSize(fileRange.Length)
	if frs.carSize+nextSize > maxSize {
		return false
	}
	frs.fileRanges = append(frs.fileRanges, fileRange)
	frs.carSize += nextSize
	return true
}

func (frs *fileRangeSet) reset() {
	frs.fileRanges = make([]model.FileRange, 0)
	frs.carSize = carHeaderSize
}

func (frs *fileRangeSet) fileRangeIDs() []uint64 {
	return underscore.Map(frs.fileRanges, func(fileRange model.FileRange) uint64 {
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

func nextAvailablePackJob(
	db *gorm.DB,
	sourceID uint32,
) (*model.PackJob, error) {
	var packJob model.PackJob
	err := db.Where(model.PackJob{SourceID: sourceID, PackingState: model.Created}).Preload("FileRanges").FirstOrCreate(&packJob).Error
	if err != nil {
		return nil, err
	}
	return &packJob, nil
}

func PrepareSource(ctx context.Context, db *gorm.DB, datasourceHandlerResolver datasource.HandlerResolver, source model.Source, scanSource bool) error {
	directoryCache := make(map[string]uint64)
	dataset := *source.Dataset
	var remainingParts []model.FileRange
	err := db.Joins("File").
		Where("source_id = ? AND file_ranges.pack_job_id is null", source.ID).
		Order("file_ranges.id asc").
		Find(&remainingParts).Error
	if err != nil {
		return err
	}
	logger.With("remaining", len(remainingParts)).Info("remaining files")

	if scanSource {
		sourceScanner, err := datasourceHandlerResolver.Resolve(ctx, source)
		if err != nil {
			return fmt.Errorf("failed to get source scanner: %w", err)
		}
		entryChan := sourceScanner.Scan(ctx, "", source.LastScannedPath)
		for entry := range entryChan {
			if entry.Error != nil {
				logger.Errorw("failed to scan", "error", entry.Error)
				continue
			}
			file, fileRanges, err := PushFile(ctx, db, entry.Info, source, dataset, directoryCache)
			if err != nil {
				return fmt.Errorf("failed to push file: %w", err)
			}
			if file == nil {
				logger.Infow("file already exists", "path", entry.Info.Remote())
				continue
			}
			err = database.DoRetry(ctx, func() error {
				return db.Model(&model.Source{}).Where("id = ?", source.ID).
					Update("last_scanned_path", file.Path).Error
			})
			if err != nil {
				return fmt.Errorf("failed to update last scanned path: %w", database.ErrDatabaseNotSupported)
			}

			remainingParts = append(remainingParts, fileRanges...)
		}
	}
	logger.With("remaining", len(remainingParts)).Info("remaining items")
	_, err = prepareToPackFileRanges(ctx, db, source.ID, dataset.MaxSize, remainingParts)
	if err != nil {
		return err
	}
	return markPackJobsReady(ctx, db, source.ID)
}
