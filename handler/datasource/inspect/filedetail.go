package inspect

import (
	"context"
	"strconv"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

func GetSourceFileDetailHandler(
	ctx context.Context,
	db *gorm.DB,
	id string,
) (*model.File, error) {
	return getSourceFileDetailHandler(db.WithContext(ctx), id)
}

// @Summary Get details about an file
// @Tags Data Source
// @Accept json
// @Produce json
// @Param id path string true "File ID"
// @Success 200 {object} model.File
// @Failure 500 {object} api.HTTPError
// @Router /file/{id} [get]
func getSourceFileDetailHandler(
	db *gorm.DB,
	id string,
) (*model.File, error) {
	fileID, err := strconv.Atoi(id)
	if err != nil {
		return nil, handler.NewInvalidParameterErr("invalid file id")
	}
	var file model.File
	err = db.Preload("FileRanges").Where("id = ?", fileID).First(&file).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, handler.NewInvalidParameterErr("file not found")
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &file, nil
}
