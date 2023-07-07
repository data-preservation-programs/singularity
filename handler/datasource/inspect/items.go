package inspect

import (
	"strconv"

	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// GetSourceItemsHandler godoc
// @Summary Get all item details of a data source
// @Tags Data Source
// @Accept json
// @Produce json
// @Param id path string true "Source ID"
// @Success 200 {array} model.Item
// @Failure 400 {object} handler.HTTPError
// @Failure 500 {object} handler.HTTPError
// @Router /source/{id}/items [get]
func GetSourceItemsHandler(
	db *gorm.DB,
	id string,
) ([]model.Item, error) {
	sourceID, err := strconv.Atoi(id)
	if err != nil {
		return nil, handler.NewBadRequestString("invalid source id")
	}
	var source model.Source
	err = db.Where("id = ?", sourceID).First(&source).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, handler.NewBadRequestString("source not found")
	}
	if err != nil {
		return nil, handler.NewHandlerError(err)
	}

	var items []model.Item
	err = db.Where("source_id = ?", sourceID).Find(&items).Error
	if err != nil {
		return nil, handler.NewHandlerError(err)
	}

	return items, nil
}
