package datasource

import (
	"context"
	"fmt"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/rclone/rclone/fs"
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
