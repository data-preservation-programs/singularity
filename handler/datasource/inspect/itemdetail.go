package inspect

import (
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"strconv"
)

// GetSourceItemDetailHandler godoc
// @Summary Get details about an item
// @Tags Data Source
// @Accept json
// @Produce json
// @Param id path string true "Item ID"
// @Success 200 {object} model.Item
// @Failure 500 {object} handler.HTTPError
// @Router /item/{id} [get]
func GetSourceItemDetailHandler(
	db *gorm.DB,
	id string,
) (*model.Item, error) {
	itemID, err := strconv.Atoi(id)
	if err != nil {
		return nil, handler.NewBadRequestString("invalid item id")
	}
	var item model.Item
	err = db.Preload("ItemParts").Where("id = ?", itemID).First(&item).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, handler.NewBadRequestString("item not found")
	}
	if err != nil {
		return nil, handler.NewHandlerError(err)
	}

	return &item, nil
}
