package datasource

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/rclone/rclone/fs"
	"github.com/rclone/rclone/fs/hash"
	"gorm.io/gorm"
)

type FileInfo struct {
	Path string `json:"path"` // Path to the new file, relative to the source
}

func PushFileHandler(
	ctx context.Context,
	db *gorm.DB,
	datasourceHandlerResolver storagesystem.HandlerResolver,
	sourceID uint32,
	fileInfo FileInfo,
) (*model.File, error) {
	return pushFileHandler(ctx, db.WithContext(ctx), datasourceHandlerResolver, sourceID, fileInfo)
}

// @Summary Push a file to be queued
// @Description Tells Singularity that something is ready to be grabbed for data preparation
// @Tags Data Source
// @Accept json
// @Produce json
// @Param id path string true "Source ID"
// @Param file body FileInfo true "File"
// @Success 201 {object} model.File
// @Failure 400 {string} string "Bad Request"
// @Failure 409 {string} string "File already exists"
// @Failure 500 {string} string "Internal Server Error"
// @Router /source/{id}/push [post]
func pushFileHandler(
	ctx context.Context,
	db *gorm.DB,
	datasourceHandlerResolver storagesystem.HandlerResolver,
	sourceID uint32,
	fileInfo FileInfo,
) (*model.File, error) {
	var source model.Source
	err := db.Joins("Preparation").Where("sources.id = ?", sourceID).First(&source).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, handler.NewInvalidParameterErr(fmt.Sprintf("source %d not found.", sourceID))
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	dsHandler, err := datasourceHandlerResolver.Resolve(ctx, source)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	entry, err := dsHandler.Check(ctx, fileInfo.Path)
	if err != nil {
		return nil, handler.InvalidParameterError{Err: err}
	}

	obj, ok := entry.(fs.ObjectInfo)
	if !ok {
		return nil, handler.NewInvalidParameterErr(fmt.Sprintf("%s is not an object", fileInfo.Path))
	}

	file, fileRanges, err := PushFile(ctx, db, obj, source, *source.Dataset, map[string]uint64{})

	if err != nil {
		return nil, errors.WithStack(err)
	}

	if file == nil {
		return nil, handler.NewDuplicateRecordError(fmt.Sprintf("%s already exists", obj.Remote()))
	}

	file.FileRanges = fileRanges

	return file, nil
}

func PushFile(ctx context.Context, db *gorm.DB, obj fs.ObjectInfo,
	source model.Source, dataset model.Preparation,
	directoryCache map[string]uint64) (*model.File, []model.FileRange, error) {
	logger.Debugw("pushed file", "file", obj.Remote(), "source_id", source.ID, "dataset_id", dataset.ID)
	db = db.WithContext(ctx)
	splitSize := MaxSizeToSplitSize(dataset.MaxSize)
	rootID, err := source.RootDirectoryID(db)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to get root directory id")
	}
	size, hashValue, lastModified, lastModifiedReliable := ExtractFromFsObject(ctx, obj)
	existing := int64(0)
	logger.Debugw("finding if the file already exists",
		"source_id", source.ID, "path", obj.Remote(),
		"hash", hashValue, "size", size, "last_modified_reliable", lastModifiedReliable, "last_modified", lastModified.UnixNano())
	err = db.Model(&model.File{}).Where(
		"source_id = ? AND path = ? AND "+
			"(( hash = ? AND hash != '' ) "+
			"OR (size = ? AND (? OR last_modified_timestamp_nano = ?)))",
		source.ID, obj.Remote(),
		hashValue, size, !lastModifiedReliable, lastModified.UnixNano(),
	).Count(&existing).Error
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to check existing file")
	}
	if existing > 0 {
		logger.Debugw("file already exists", "source_id", source.ID, "path", obj.Remote())
		return nil, nil, nil
	}
	file := model.File{
		SourceID:         source.ID,
		Path:             obj.Remote(),
		Size:             size,
		LastModifiedNano: lastModified.UnixNano(),
		Hash:             hashValue,
	}
	logger.Debugw("new file", "file", file)
	err = EnsureParentDirectories(ctx, db, &file, rootID, directoryCache)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to ensure parent directories")
	}
	var fileRanges []model.FileRange
	err = database.DoRetry(ctx, func() error {
		return db.Transaction(func(db *gorm.DB) error {
			err := db.Create(&file).Error
			if err != nil {
				return errors.Wrap(err, "failed to create file")
			}
			if dataset.UseEncryption() {
				fileRanges = append(fileRanges, model.FileRange{
					FileID: file.ID,
					Offset: 0,
					Length: file.Size,
				})
			} else {
				offset := int64(0)
				for {
					length := splitSize
					if file.Size-offset < length {
						length = file.Size - offset
					}
					fileRanges = append(fileRanges, model.FileRange{
						FileID: file.ID,
						Offset: offset,
						Length: length,
					})
					offset += length
					if offset >= file.Size {
						break
					}
				}
			}
			err = db.CreateInBatches(&fileRanges, util.BatchSize).Error
			if err != nil {
				return errors.Wrap(err, "failed to create file parts")
			}
			return nil
		})
	})
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to create file")
	}
	return &file, fileRanges, nil
}

func EnsureParentDirectories(
	ctx context.Context,
	db *gorm.DB,
	file *model.File, rootDirID uint64,
	directoryCache map[string]uint64) error {
	if file.DirectoryID != nil {
		return nil
	}
	last := rootDirID
	segments := strings.Split(file.Path, "/")
	for i, segment := range segments[:len(segments)-1] {
		p := strings.Join(segments[:i+1], "/")
		if dirID, ok := directoryCache[p]; ok {
			last = dirID
			continue
		}
		newDir := model.Directory{
			SourceID: file.SourceID,
			Name:     segment,
			ParentID: &last,
		}
		logger.Debugw("creating directory", "path", p, "dir", newDir)
		err := database.DoRetry(ctx, func() error {
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

	file.DirectoryID = &last
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
