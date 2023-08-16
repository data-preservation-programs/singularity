package datasetworker

import (
	"context"
	"strconv"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/datasource"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/pkg/errors"
	"github.com/rjNemo/underscore"
)

// scan scans the data source and inserts the chunking strategy back to database
// scanSource is true if the source will be actually scanned in addition to just picking up remaining ones
// resume is true if the scan will be resumed from the last scanned file, which is useful for resuming a failed scan
func (w *DatasetWorkerThread) scan(ctx context.Context, source model.Source, scanSource bool) error {
	dataset := *source.Dataset
	var remaining = newRemain()
	var remainingParts []model.FileRange
	err := w.db.Joins("File").
		Where("source_id = ? AND file_ranges.packing_manifest_id is null", source.ID).
		Order("file_ranges.id asc").
		Find(&remainingParts).Error
	if err != nil {
		return err
	}
	w.logger.With("remaining", len(remainingParts)).Info("remaining files")
	remaining.add(remainingParts)

	if !scanSource {
		for len(remaining.fileRanges) > 0 {
			err = w.chunkOnce(source, dataset, remaining)
			if err != nil {
				return errors.Wrap(err, "failed to save chunking")
			}
		}
		return nil
	}

	sourceScanner, err := w.datasourceHandlerResolver.Resolve(ctx, source)
	if err != nil {
		return errors.Wrap(err, "failed to get source scanner")
	}
	entryChan := sourceScanner.Scan(ctx, "", source.LastScannedPath)
	for entry := range entryChan {
		if entry.Error != nil {
			w.logger.Errorw("failed to scan", "error", entry.Error)
			continue
		}

		file, fileRanges, err := datasource.PushFile(ctx, w.db, entry.Info, source, dataset, w.directoryCache)
		if err != nil {
			return errors.Wrap(err, "failed to push file")
		}
		if file == nil {
			w.logger.Infow("file already exists", "path", entry.Info.Remote())
			continue
		}
		err = database.DoRetry(func() error {
			return w.db.Model(&model.Source{}).Where("id = ?", source.ID).
				Update("last_scanned_path", file.Path).Error
		})
		if err != nil {
			return errors.Wrap(err, "failed to update last scanned path")
		}

		remaining.add(fileRanges)
		for remaining.carSize >= dataset.MaxSize {
			err = w.chunkOnce(source, dataset, remaining)
			if err != nil {
				return errors.Wrap(err, "failed to save chunking")
			}
		}
	}

	for len(remaining.fileRanges) > 0 {
		err = w.chunkOnce(source, dataset, remaining)
		if err != nil {
			return errors.Wrap(err, "failed to save chunking")
		}
	}
	return nil
}

func (w *DatasetWorkerThread) chunkOnce(
	source model.Source,
	dataset model.Dataset,
	remaining *remain,
) error {
	// If everything fit, create a packing manifest. Usually this is the case for the last packing manifest
	if remaining.carSize <= dataset.MaxSize {
		w.logger.Debugw("creating packing manifest", "size", remaining.carSize)
		_, err := datasource.CreatePackingManifestHandler(w.db, strconv.FormatUint(uint64(source.ID), 10), datasource.PackingManifestRequest{
			FileIDs: remaining.itemIDs(),
		})

		if err != nil {
			return errors.Wrap(err, "failed to create packing manifest")
		}
		remaining.reset()
		return nil
	}
	// size > maxSize, first, find the first file that makes it larger than maxSize
	s := remaining.carSize
	si := len(remaining.fileRanges) - 1
	for si >= 0 {
		s -= toCarSize(remaining.fileRanges[si].Length)
		if s <= dataset.MaxSize {
			break
		}
		si--
	}

	// In case si == 0, this is the case where a single file is more than sector size for encryption
	// We will allow a single file to be more than sector size and handle it later during packing
	if si == 0 {
		si = 1
		s += toCarSize(remaining.fileRanges[0].Length)
	}

	// create a packing manifest for [0:si)
	w.logger.Debugw("creating packing manifest", "size", s)

	fileRangeIDs := underscore.Map(remaining.fileRanges[:si], func(file model.FileRange) uint64 {
		return file.ID
	})
	_, err := datasource.CreatePackingManifestHandler(w.db, strconv.FormatUint(uint64(source.ID), 10), datasource.PackingManifestRequest{
		FileIDs: fileRangeIDs,
	})
	if err != nil {
		return errors.Wrap(err, "failed to create packing manifest")
	}
	remaining.fileRanges = remaining.fileRanges[si:]
	remaining.carSize = remaining.carSize - s + carHeaderSize
	return nil
}
