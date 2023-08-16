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
func (w *Thread) scan(ctx context.Context, source model.Source, scanSource bool) error {
	db := w.dbNoContext.WithContext(ctx)
	directoryCache := make(map[string]uint64)
	dataset := *source.Dataset
	var remaining = newRemain()
	var remainingParts []model.ItemPart
	err := db.Joins("Item").
		Where("source_id = ? AND item_parts.chunk_id is null", source.ID).
		Order("item_parts.id asc").
		Find(&remainingParts).Error
	if err != nil {
		return err
	}
	w.logger.With("remaining", len(remainingParts)).Info("remaining items")
	remaining.add(remainingParts)

	if !scanSource {
		for len(remaining.itemParts) > 0 {
			err = w.chunkOnce(ctx, source, dataset, remaining)
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

		item, itemParts, err := datasource.PushItem(ctx, w.dbNoContext, entry.Info, source, dataset, directoryCache)
		if err != nil {
			return errors.Wrap(err, "failed to push item")
		}
		if item == nil {
			w.logger.Infow("item already exists", "path", entry.Info.Remote())
			continue
		}
		err = database.DoRetry(ctx, func() error {
			return db.Model(&model.Source{}).Where("id = ?", source.ID).
				Update("last_scanned_path", item.Path).Error
		})
		if err != nil {
			return errors.Wrap(err, "failed to update last scanned path")
		}

		remaining.add(itemParts)
		for remaining.carSize >= dataset.MaxSize {
			err = w.chunkOnce(ctx, source, dataset, remaining)
			if err != nil {
				return errors.Wrap(err, "failed to save chunking")
			}
		}
	}

	for len(remaining.itemParts) > 0 {
		err = w.chunkOnce(ctx, source, dataset, remaining)
		if err != nil {
			return errors.Wrap(err, "failed to save chunking")
		}
	}
	return nil
}

func (w *Thread) chunkOnce(
	ctx context.Context,
	source model.Source,
	dataset model.Dataset,
	remaining *remain,
) error {
	// If everything fit, create a chunk. Usually this is the case for the last chunk
	if remaining.carSize <= dataset.MaxSize {
		w.logger.Debugw("creating chunk", "size", remaining.carSize)
		_, err := datasource.ChunkHandler(ctx, w.dbNoContext, strconv.FormatUint(uint64(source.ID), 10), datasource.ChunkRequest{
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
	_, err := datasource.ChunkHandler(ctx, w.dbNoContext, strconv.FormatUint(uint64(source.ID), 10), datasource.ChunkRequest{
		ItemIDs: itemPartIDs,
	})
	if err != nil {
		return errors.Wrap(err, "failed to create chunk")
	}
	remaining.itemParts = remaining.itemParts[si:]
	remaining.carSize = remaining.carSize - s + carHeaderSize
	return nil
}

type remain struct {
	itemParts []model.ItemPart
	carSize   int64
}

const carHeaderSize = 59

func newRemain() *remain {
	return &remain{
		itemParts: make([]model.ItemPart, 0),
		// Some buffer for header
		carSize: carHeaderSize,
	}
}

func (r *remain) add(itemParts []model.ItemPart) {
	r.itemParts = append(r.itemParts, itemParts...)
	for _, itemPart := range itemParts {
		r.carSize += toCarSize(itemPart.Length)
	}
}

func (r *remain) reset() {
	r.itemParts = make([]model.ItemPart, 0)
	r.carSize = carHeaderSize
}

func (r *remain) itemIDs() []uint64 {
	return underscore.Map(r.itemParts, func(itemPart model.ItemPart) uint64 {
		return itemPart.ID
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
