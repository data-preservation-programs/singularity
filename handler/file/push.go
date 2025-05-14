package file

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/pack/push"
	"github.com/data-preservation-programs/singularity/storagesystem"
	"github.com/rclone/rclone/fs"
	"gorm.io/gorm"
)

type Info struct {
	Path string `json:"path"` // Path to the new file, relative to the source
}

// PushFileHandler pushes a file to the database using specified preparation and source details.
//
// This function retrieves the source attachment by its preparation and source. If the source isn't
// attached to the given preparation, an ErrNotFound error is returned. The function then validates
// the file's existence in the storage system using the RClone handler. If the file is validated
// successfully, it is then pushed to the database.
//
// Parameters:
//   - ctx: The context for managing timeouts and cancellation.
//   - db: The gorm.DB instance for database operations.
//   - preparation: The preparation ID or name.
//   - source: The source ID or name.
//   - fileInfo: Information regarding the file to be pushed.
//
// Returns:
//   - A pointer to the pushed model.File, if successful.
//   - An error if any issues occur during the operation, including when the source isn't attached to
//     the preparation, if the file doesn't exist in the storage system, or if the file already exists
func (DefaultHandler) PushFileHandler(
	ctx context.Context,
	db *gorm.DB,
	preparation string,
	source string,
	fileInfo Info,
) (*model.File, error) {
	db = db.WithContext(ctx)
	var attachment model.SourceAttachment
	err := attachment.FindByPreparationAndSource(db, preparation, source)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(handlererror.ErrNotFound, "source '%s' is not attached to preparation %s", source, preparation)
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	rclone, err := storagesystem.NewRCloneHandler(ctx, *attachment.Storage)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	entry, err := rclone.Check(ctx, fileInfo.Path)
	if err != nil {
		return nil, errors.Join(handlererror.ErrInvalidParameter, errors.Wrapf(err, "failed to check file '%s'", fileInfo.Path))
	}

	obj, ok := entry.(fs.ObjectInfo)
	if !ok {
		return nil, errors.Wrapf(handlererror.ErrInvalidParameter, "file '%s' is not an object", fileInfo.Path)
	}

	file, fileRanges, err := push.PushFile(ctx, db, obj, attachment, map[string]model.DirectoryID{})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if file == nil {
		return nil, errors.Wrapf(handlererror.ErrDuplicateRecord, "file '%s' already exists", fileInfo.Path)
	}

	file.FileRanges = fileRanges

	return file, nil
}

// @ID PushFile
// @Summary Push a file to be queued
// @Description Tells Singularity that something is ready to be grabbed for data preparation
// @Tags File
// @Accept json
// @Produce json
// @Param id path string true "Preparation ID or name"
// @Param name path string true "Source storage ID or name"
// @Param file body Info true "File Info"
// @Success 200 {object} model.File
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /preparation/{id}/source/{name}/file [post]
func _() {}
