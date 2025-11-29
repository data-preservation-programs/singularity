package push

import (
	"context"
	"strings"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util"
	logging "github.com/ipfs/go-log/v2"
	"github.com/rclone/rclone/fs"
	"github.com/rclone/rclone/fs/hash"
	"gorm.io/gorm"
)

var logger = logging.Logger("pushfile")

func MaxSizeToSplitSize(m int64) int64 {
	r := util.NextPowerOfTwo(uint64(m)) / 4
	if r > 1<<30 {
		r = 1 << 30
	}

	return int64(r)
}

func ExtractFromFsObject(ctx context.Context, info fs.ObjectInfo) (size int64, hashValue string, lastModified time.Time) {
	lastModified = info.ModTime(ctx)
	size = info.Size()
	// For local file system, rclone is actually hashing the file stream, which is inefficient.
	// So we skip hashing for local file system.
	// For some of the remote storage, there may not have any supported hash type.
	if info.Fs().Features().SlowHash {
		return
	}
	supportedHash := info.Fs().Hashes().GetOne()
	if supportedHash != hash.None {
		var err error
		hashValue, err = info.Hash(ctx, supportedHash)
		if err != nil {
			logger.Errorw("failed to hash", "error", err)
		}
	}
	return
}

func PushFile(
	ctx context.Context,
	db *gorm.DB,
	obj fs.ObjectInfo,
	attachment model.SourceAttachment,
	directoryCache map[string]model.DirectoryID,
) (*model.File, []model.FileRange, error) {
	logger.Debugw("pushing file", "file", obj.Remote(), "preparation", attachment.PreparationID, "storage", attachment.StorageID)
	db = db.WithContext(ctx)
	splitSize := MaxSizeToSplitSize(attachment.Preparation.MaxSize)
	rootID, err := attachment.RootDirectoryID(ctx, db)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "failed to get root directory for attachment %d", attachment.ID)
	}

	size, hashValue, lastModified := ExtractFromFsObject(ctx, obj)
	existing := int64(0)

	logger.Debugw("finding if the file already exists", "path", obj.Remote(),
		"hash", hashValue, "size", size, "last_modified", lastModified.UnixNano())
	query := "attachment_id = ? AND path = ? AND last_modified_nano = ?"
	args := []interface{}{attachment.ID, obj.Remote(), lastModified.UnixNano()}
	if hashValue != "" {
		query += " AND hash = ?"
		args = append(args, hashValue)
	}
	// An edge case is the size is not available from the data source, which results in size = -1.
	if size >= 0 {
		query += " AND size = ?"
		args = append(args, size)
	} else {
		logger.Warnw("size is not available, this may overflow a sector if the actual size of the file is too large", "path", obj.Remote())
	}
	err = db.Model(&model.File{}).Where(query, args...).Count(&existing).Error
	if err != nil {
		return nil, nil, errors.Wrapf(err, "failed to find existing file %s", obj.Remote())
	}
	if existing > 0 {
		logger.Debugw("file already exists", "path", obj.Remote())
		return nil, nil, nil
	}

	file := model.File{
		AttachmentID:     &attachment.ID,
		Path:             obj.Remote(),
		Size:             size,
		LastModifiedNano: lastModified.UnixNano(),
		Hash:             hashValue,
	}

	logger.Infow("new file", "file", file)
	err = EnsureParentDirectories(ctx, db, &file, rootID, directoryCache)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	var fileRanges []model.FileRange
	err = database.DoRetry(ctx, func() error {
		return db.Transaction(func(db *gorm.DB) error {
			err := db.Create(&file).Error
			if err != nil {
				return errors.WithStack(err)
			}
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
			err = db.CreateInBatches(&fileRanges, util.BatchSize).Error
			if err != nil {
				return errors.WithStack(err)
			}
			return nil
		})
	})
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	return &file, fileRanges, nil
}

func EnsureParentDirectories(
	ctx context.Context,
	db *gorm.DB,
	file *model.File, rootDirID model.DirectoryID,
	directoryCache map[string]model.DirectoryID,
) error {
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
			AttachmentID: file.AttachmentID,
			Name:         segment,
			ParentID:     &last,
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
			return errors.WithStack(err)
		}
		directoryCache[p] = newDir.ID
		last = newDir.ID
	}

	file.DirectoryID = &last
	return nil
}

func CreatePackJob(
	ctx context.Context,
	db *gorm.DB,
	attachmentID model.SourceAttachmentID,
	fileRangeIDs []model.FileRangeID,
) (*model.Job, error) {
	db = db.WithContext(ctx)
	job := model.Job{
		Type:         model.Pack,
		AttachmentID: &attachmentID,
		State:        model.Ready,
	}

	err := database.DoRetry(ctx, func() error {
		return db.Transaction(
			func(db *gorm.DB) error {
				err := db.Create(&job).Error
				if err != nil {
					return errors.WithStack(err)
				}
				fileRangeIDChunks := util.ChunkSlice(fileRangeIDs, util.BatchSize)
				for _, fileRangeIDChunks := range fileRangeIDChunks {
					err = db.Model(&model.FileRange{}).
						Where("id IN ?", fileRangeIDChunks).Update("job_id", job.ID).Error
					if err != nil {
						return errors.WithStack(err)
					}
				}
				return nil
			},
		)
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &job, nil
}
