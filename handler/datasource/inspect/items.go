package inspect

import (
	"strconv"

	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func GetSourceFilesHandler(
	db *gorm.DB,
	id string,
) ([]model.File, error) {
	return getSourceFilesHandler(db, id)
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
		return nil, handler.NewInvalidParameterErr("invalid source id")
	}
	var source model.Source
	err = db.Where("id = ?", sourceID).First(&source).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, handler.NewInvalidParameterErr("source not found")
	}
	if err != nil {
		return nil, err
	}

	var files []model.File
	err = db.Where("source_id = ?", sourceID).Find(&files).Error
	if err != nil {
		return nil, err
	}

	return files, nil
}
