package inspect

import (
	"context"
	"strconv"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

func GetSourceFilesHandler(
	ctx context.Context,
	db *gorm.DB,
	id string,
) ([]model.File, error) {
	return getSourceFilesHandler(db.WithContext(ctx), id)
}

// @Summary Get all file details of a data source
// @Tags Data Source
// @Accept json
// @Produce json
// @Param id path string true "Source ID"
// @Success 200 {array} model.File
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /source/{id}/files [get]
func getSourceFilesHandler(
	db *gorm.DB,
	id string,
) ([]model.File, error) {
	sourceID, err := strconv.Atoi(id)
	if err != nil {
		return nil, handlererror.NewInvalidParameterErr("invalid source id")
	}
	var source model.Source
	err = db.Where("id = ?", sourceID).First(&source).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, handlererror.NewInvalidParameterErr("source not found")
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var files []model.File
	err = db.Where("source_id = ?", sourceID).Find(&files).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return files, nil
}
