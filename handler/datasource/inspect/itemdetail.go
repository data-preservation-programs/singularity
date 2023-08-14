package inspect

import (
	"strconv"

	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func GetSourceItemDetailHandler(
	db *gorm.DB,
	id string,
) (*model.Item, error) {
	return getSourceItemDetailHandler(db, id)
}

// @Summary Get details about an item
// @Tags Data Source
// @Accept json
// @Produce json
// @Param id path string true "Item ID"
// @Success 200 {object} model.Item
// @Failure 500 {object} api.HTTPError
// @Router /item/{id} [get]
func getSourceItemDetailHandler(
	db *gorm.DB,
	id string,
) (*model.Item, error) {
	itemID, err := strconv.Atoi(id)
	if err != nil {
		return nil, handler.NewInvalidParameterErr("invalid item id")
	}
	var item model.Item
	err = db.Preload("FileRanges").Where("id = ?", itemID).First(&item).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, handler.NewInvalidParameterErr("item not found")
	}
	if err != nil {
		return nil, err
	}

	return &item, nil
}
