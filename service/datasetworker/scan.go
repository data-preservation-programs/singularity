package datasetworker

import (
	"context"
	"strings"
	"time"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/pkg/errors"
	"github.com/rclone/rclone/fs"
	"github.com/rclone/rclone/fs/hash"
	"github.com/rjNemo/underscore"
	"gorm.io/gorm"
)

func MaxSizeToSplitSize(m int64) int64 {
	r := util.NextPowerOfTwo(uint64(m)) / 4
	if r > 1<<30 {
		r = 1 << 30
	}

	return int64(r)
}

func ExtractFromFsObject(ctx context.Context, info fs.ObjectInfo) (size int64, hashValue string, lastModified time.Time, lastModifiedReliable bool) {
	// last modified can be time.Now() if fetch failed so it may not be reliable.
	// This usually won't happen for most cloud provider i.e. S3
	// Because during scanning, the modified time is already fetched.
	lastModified = info.ModTime(ctx)
	// If last modified is not reliable, we will skip using it as a way to determine if the file has already scanned
	lastModifiedReliable = !lastModified.IsZero() && lastModified.Before(time.Now().Add(-time.Millisecond))
	supportedHash := info.Fs().Hashes().GetOne()
	// For local file system, rclone is actually hashing the file stream which is not efficient.
	// So we skip hashing for local file system.
	// For some of the remote storage, there may not have any supported hash type.
	if supportedHash != hash.None && info.Fs().Name() != "local" {
		var err error
		hashValue, err = info.Hash(ctx, supportedHash)
		if err != nil {
			logger.Errorw("failed to hash", "error", err)
		}
	}
	size = info.Size()
	return
}

func PushItem(ctx context.Context, db *gorm.DB, obj fs.ObjectInfo,
	source model.Source, dataset model.Dataset,
	directoryCache map[string]uint64) (*model.Item, []model.ItemPart, error) {
	db = db.WithContext(ctx)
	splitSize := MaxSizeToSplitSize(dataset.MaxSize)
	rootID, err := source.RootDirectoryID(db)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to get root directory id")
	}
	size, hashValue, lastModified, lastModifiedReliable := ExtractFromFsObject(ctx, obj)
	existing := int64(0)
	err = db.Model(&model.Item{}).Where(
		"source_id = ? AND path = ? AND "+
			"(( hash = ? AND hash != '' ) "+
			"OR (size = ? AND (? OR last_modified_timestamp_nano = ?)))",
		source.ID, obj.Remote(),
		hashValue, size, !lastModifiedReliable, lastModified.UnixNano(),
	).Count(&existing).Error
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to check existing item")
	}
	if existing > 0 {
		return nil, nil, nil
	}
	item := model.Item{
		SourceID:                  source.ID,
		Path:                      obj.Remote(),
		Size:                      size,
		LastModifiedTimestampNano: lastModified.UnixNano(),
		Hash:                      hashValue,
	}
	logger.Debugw("found item", "item", item)
	err = EnsureParentDirectories(db, &item, rootID, directoryCache)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to ensure parent directories")
	}
	var itemParts []model.ItemPart
	err = database.DoRetry(func() error {
		return db.Transaction(func(db *gorm.DB) error {
			err := db.Create(&item).Error
			if err != nil {
				return errors.Wrap(err, "failed to create item")
			}
			if dataset.UseEncryption() {
				itemParts = append(itemParts, model.ItemPart{
					ItemID: item.ID,
					Offset: 0,
					Length: item.Size,
				})
			} else {
				offset := int64(0)
				for {
					length := splitSize
					if item.Size-offset < length {
						length = item.Size - offset
					}
					itemParts = append(itemParts, model.ItemPart{
						ItemID: item.ID,
						Offset: offset,
						Length: length,
					})
					offset += length
					if offset >= item.Size {
						break
					}
				}
			}
			err = db.Create(&itemParts).Error
			if err != nil {
				return errors.Wrap(err, "failed to create item parts")
			}
			return nil
		})
	})
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to create item")
	}
	return &item, itemParts, nil
}

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

		item, itemParts, err := PushItem(ctx, w.db, entry.Info, source, dataset, w.directoryCache)
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
		err := database.DoRetry(func() error {
			return w.db.Transaction(
				func(db *gorm.DB) error {
					chunk := model.Chunk{
						SourceID:     source.ID,
						PackingState: model.Ready,
					}
					err := w.db.Create(&chunk).Error
					if err != nil {
						return errors.Wrap(err, "failed to create chunk")
					}
					err = w.db.Model(&model.ItemPart{}).
						Where("id IN ?", remaining.itemIDs()).Update("chunk_id", chunk.ID).Error
					if err != nil {
						return errors.Wrap(err, "failed to update items")
					}
					return nil
				},
			)
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
	err := database.DoRetry(func() error {
		return w.db.Transaction(
			func(db *gorm.DB) error {
				chunk := model.Chunk{
					SourceID:     source.ID,
					PackingState: model.Ready,
				}
				err := w.db.Create(&chunk).Error
				if err != nil {
					return errors.Wrap(err, "failed to create chunk")
				}
				itemPartIDs := underscore.Map(remaining.itemParts[:si], func(item model.ItemPart) uint64 {
					return item.ID
				})
				err = w.db.Model(&model.ItemPart{}).Where("id IN ?", itemPartIDs).Update("chunk_id", chunk.ID).Error
				if err != nil {
					return errors.Wrap(err, "failed to update items")
				}
				return nil
			},
		)
	})
	if err != nil {
		return errors.Wrap(err, "failed to create chunk")
	}
	remaining.itemParts = remaining.itemParts[si:]
	remaining.carSize = remaining.carSize - s + carHeaderSize
	return nil
}

func EnsureParentDirectories(db *gorm.DB,
	item *model.Item, rootDirID uint64,
	directoryCache map[string]uint64) error {
	if item.DirectoryID != nil {
		return nil
	}
	last := rootDirID
	segments := strings.Split(item.Path, "/")
	for i, segment := range segments[:len(segments)-1] {
		p := strings.Join(segments[:i+1], "/")
		if dirID, ok := directoryCache[p]; ok {
			last = dirID
			continue
		}
		newDir := model.Directory{
			SourceID: item.SourceID,
			Name:     segment,
			ParentID: &last,
		}
		err := database.DoRetry(func() error {
			return db.Transaction(func(db *gorm.DB) error {
				return db.
					Where("parent_id = ? AND name = ?", last, segment).
					FirstOrCreate(&newDir).Error
			})
		})
		if err != nil {
			return errors.Wrap(err, "failed to create directory")
		}
		directoryCache[p] = newDir.ID
		last = newDir.ID
	}

	item.DirectoryID = &last
	return nil
}
