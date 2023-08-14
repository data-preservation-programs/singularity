package inspect

import (
	"context"
	"strconv"

	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func GetSourceItemsHandler(
	ctx context.Context,
	db *gorm.DB,
	id string,
) ([]model.Item, error) {
	return getSourceItemsHandler(db.WithContext(ctx), id)
}

// @Summary Get all item details of a data source
// @Tags Data Source
// @Accept json
// @Produce json
// @Param id path string true "Source ID"
// @Success 200 {array} model.Item
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /source/{id}/items [get]
func getSourceItemsHandler(
	db *gorm.DB,
	id string,
) ([]model.Item, error) {
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

	var items []model.Item
	err = db.Where("source_id = ?", sourceID).Find(&items).Error
	if err != nil {
		return nil, err
	}

	return items, nil
}
