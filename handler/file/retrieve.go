package file

import (
	"context"
	"io"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/storagesystem"
	"gorm.io/gorm"
)

// RetrieveFileHandler retrieves the actual bytes for a file on disk using a given file ID.
//
// # For now, this function only works if the file remains available in its original source storage
//
// Parameters:
// - ctx: The context for managing timeouts and cancellation.
// - db: The gorm.DB instance for database operations.
// - id: The ID of the file to be retrieved.
//
// Returns:
// - A ReadSeekCloser for the given file
// - the name of the file
// - An error if any issues occur during the database operation, including when the file is not found.
func (DefaultHandler) RetrieveFileHandler(
	ctx context.Context,
	db *gorm.DB,
	id uint64,
) (data io.ReadSeekCloser, name string, modTime time.Time, err error) {
	db = db.WithContext(ctx)
	var file model.File
	err = db.Preload("Attachment.Storage").First(&file, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, "", time.Time{}, errors.Wrapf(handlererror.ErrNotFound, "file '%d' does not exist", id)
	}
	if err != nil {
		return nil, "", time.Time{}, errors.WithStack(err)
	}

	rclone, err := storagesystem.NewRCloneHandler(ctx, *file.Attachment.Storage)
	if err != nil {
		return nil, file.FileName(), time.Unix(0, file.LastModifiedNano), errors.WithStack(err)
	}

	// TODO: if we cannot open the file from the file system, use Filecoin retrieval
	seeker, obj, err := storagesystem.Open(rclone, ctx, file.Path)
	if err != nil {
		return nil, file.FileName(), time.Unix(0, file.LastModifiedNano), errors.WithStack(err)
	}

	return seeker, file.FileName(), obj.ModTime(ctx), nil
}
