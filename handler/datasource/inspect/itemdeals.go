package inspect

import (
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

func GetItemDealsHandler(
	db *gorm.DB,
	id string,
) ([]model.Deal, error) {
	return getItemDealsHandler(db, id)
}

// @Summary Get all deals that have been made for an item
// @Tags Data Source
// @Accept json
// @Produce json
// @Param id path string true "Source ID"
// @Success 200 {array} model.Deal
// @Failure 500 {object} api.HTTPError
// @Router /item/{id}/deals [get]
func getItemDealsHandler(
	db *gorm.DB,
	id string,
) ([]model.Deal, error) {
	var deals []model.Deal
	query := db.
		Model(&model.Item{}).
		Select("deals.*").
		Joins("JOIN item_parts ON items.id = item_parts.item_id").
		Joins("JOIN chunks ON item_parts.chunk_id = chunks.id").
		Joins("JOIN cars ON chunks.id = cars.chunk_id").
		Joins("JOIN deals ON cars.piece_cid = deals.piece_cid").
		Where("id = ?", id)
	if err := query.Find(&deals).Error; err != nil {
		return nil, err
	}

	return deals, nil
}
