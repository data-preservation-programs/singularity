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
// resume is true if the scan will be resumed from the last scanned item, which is useful for resuming a failed scan
func (w *DatasetWorkerThread) scan(ctx context.Context, source model.Source, scanSource bool) error {
	dataset := *source.Dataset
	var remaining = newRemain()
	var remainingParts []model.ItemPart
	err := w.db.Joins("Item").Preload("Item").
		Where("source_id = ? AND chunk_id is null", source.ID).
		Order("item_parts.id asc").
		Find(&remainingParts).Error
	if err != nil {
		return err
	}
	w.logger.With("remaining", len(remainingParts)).Info("remaining items")
	remaining.add(remainingParts)

	if !scanSource {
		for len(remaining.itemParts) > 0 {
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

		item, itemParts, err := datasource.PushItem(ctx, w.db, entry.Info, source, dataset, w.directoryCache)
		if err != nil {
			return errors.Wrap(err, "failed to push item")
		}
		if item == nil {
			w.logger.Infow("item already exists", "path", entry.Info.Remote())
			continue
		}
		err = database.DoRetry(func() error {
			return w.db.Model(&model.Source{}).Where("id = ?", source.ID).
				Update("last_scanned_path", item.Path).Error
		})
		if err != nil {
			return errors.Wrap(err, "failed to update last scanned path")
		}

		remaining.add(itemParts)
		for remaining.carSize >= dataset.MaxSize {
			err = w.chunkOnce(source, dataset, remaining)
			if err != nil {
				return errors.Wrap(err, "failed to save chunking")
			}
		}
	}

	for len(remaining.itemParts) > 0 {
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
	// If everything fit, create a chunk. Usually this is the case for the last chunk
	if remaining.carSize <= dataset.MaxSize {
		w.logger.Debugw("creating chunk", "size", remaining.carSize)
		_, err := datasource.ChunkHandler(w.db, strconv.FormatUint(uint64(source.ID), 10), datasource.ChunkRequest{
			ItemIDs: remaining.itemIDs(),
		})

		if err != nil {
			return errors.Wrap(err, "failed to create chunk")
		}
		remaining.reset()
		return nil
	}
	// size > maxSize, first, find the first item that makes it larger than maxSize
	s := remaining.carSize
	si := len(remaining.itemParts) - 1
	for si >= 0 {
		s -= toCarSize(remaining.itemParts[si].Length)
		if s <= dataset.MaxSize {
			break
		}
		si--
	}

	// In case si == 0, this is the case where a single item is more than sector size for encryption
	// We will allow a single item to be more than sector size and handle it later during packing
	if si == 0 {
		si = 1
		s += toCarSize(remaining.itemParts[0].Length)
	}

	// create a chunk for [0:si)
	w.logger.Debugw("creating chunk", "size", s)

	itemPartIDs := underscore.Map(remaining.itemParts[:si], func(item model.ItemPart) uint64 {
		return item.ID
	})
	_, err := datasource.ChunkHandler(w.db, strconv.FormatUint(uint64(source.ID), 10), datasource.ChunkRequest{
		ItemIDs: itemPartIDs,
	})
	if err != nil {
		return errors.Wrap(err, "failed to create chunk")
	}
	remaining.itemParts = remaining.itemParts[si:]
	remaining.carSize = remaining.carSize - s + carHeaderSize
	return nil
}
