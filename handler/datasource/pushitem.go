package datasource

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/datasource"
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/pkg/errors"
	"github.com/rclone/rclone/fs"
	"github.com/rclone/rclone/fs/hash"
	"gorm.io/gorm"
)

type ItemInfo struct {
	Path string `json:"path"` // Path to the new item, relative to the source
}

func PushItemHandler(
	db *gorm.DB,
	ctx context.Context,
	datasourceHandlerResolver datasource.HandlerResolver,
	sourceID uint32,
	itemInfo ItemInfo,
) (*model.Item, error) {
	return pushItemHandler(db, ctx, datasourceHandlerResolver, sourceID, itemInfo)
}

// @Summary Push an item to be queued
// @Description Tells Singularity that something is ready to be grabbed for data preparation
// @Tags Data Source
// @Accept json
// @Produce json
// @Param id path string true "Source ID"
// @Param item body ItemInfo true "Item"
// @Success 201 {object} model.Item
// @Failure 400 {string} string "Bad Request"
// @Failure 409 {string} string "Item already exists"
// @Failure 500 {string} string "Internal Server Error"
// @Router /source/{id}/push [post]
func pushItemHandler(
	db *gorm.DB,
	ctx context.Context,
	datasourceHandlerResolver datasource.HandlerResolver,
	sourceID uint32,
	itemInfo ItemInfo,
) (*model.Item, error) {
	var source model.Source
	err := db.Preload("Dataset").Where("id = ?", sourceID).First(&source).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, handler.NewInvalidParameterErr(fmt.Sprintf("source %d not found.", sourceID))
	}
	if err != nil {
		return nil, err
	}

	dsHandler, err := datasourceHandlerResolver.Resolve(ctx, source)
	if err != nil {
		return nil, err
	}

	entry, err := dsHandler.Check(ctx, itemInfo.Path)
	if err != nil {
		return nil, handler.InvalidParameterError{Err: err}
	}

	obj, ok := entry.(fs.ObjectInfo)
	if !ok {
		return nil, handler.NewInvalidParameterErr(fmt.Sprintf("%s is not an object", itemInfo.Path))
	}

	item, itemParts, err := PushItem(ctx, db, obj, source, *source.Dataset, map[string]uint64{})

	if err != nil {
		return nil, err
	}

	if item == nil {
		return nil, handler.NewDuplicateRecordError(fmt.Sprintf("%s already exists", obj.Remote()))
	}

	item.ItemParts = itemParts

	return item, nil
}

func PushItem(ctx context.Context, db *gorm.DB, obj fs.ObjectInfo,
	source model.Source, dataset model.Dataset,
	directoryCache map[string]uint64) (*model.Item, []model.ItemPart, error) {
	logger.Debugw("pushed item", "item", obj.Remote(), "source_id", source.ID, "dataset_id", dataset.ID)
	db = db.WithContext(ctx)
	splitSize := MaxSizeToSplitSize(dataset.MaxSize)
	rootID, err := source.RootDirectoryID(db)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to get root directory id")
	}
	size, hashValue, lastModified, lastModifiedReliable := ExtractFromFsObject(ctx, obj)
	existing := int64(0)
	logger.Debugw("finding if the item already exists",
		"source_id", source.ID, "path", obj.Remote(),
		"hash", hashValue, "size", size, "last_modified_reliable", lastModifiedReliable, "last_modified", lastModified.UnixNano())
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
		logger.Debugw("item already exists", "source_id", source.ID, "path", obj.Remote())
		return nil, nil, nil
	}
	item := model.Item{
		SourceID:                  source.ID,
		Path:                      obj.Remote(),
		Size:                      size,
		LastModifiedTimestampNano: lastModified.UnixNano(),
		Hash:                      hashValue,
	}
	logger.Debugw("new item", "item", item)
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
		logger.Debugw("creating directory", "path", p, "dir", newDir)
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
