package inspect

import (
	"context"

	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

func GetFileDealsHandler(
	ctx context.Context,
	db *gorm.DB,
	id uint64,
) ([]model.Deal, error) {
	return getFileDealsHandler(ctx, db, id)
}

// @Summary Get all deals that have been made for a file
// @Tags Data Source
// @Accept json
// @Produce json
// @Param id path string true "Source ID"
// @Success 200 {array} model.Deal
// @Failure 500 {object} api.HTTPError
// @Router /file/{id}/deals [get]
func getFileDealsHandler(
	ctx context.Context,
	db *gorm.DB,
	id uint64,
) ([]model.Deal, error) {
	var deals []model.Deal
	query := db.
		Model(&model.FileRange{}).
		Select("deals.*").
		Joins("JOIN cars ON file_ranges.pack_job_id = cars.pack_job_id").
		Joins("JOIN deals ON cars.piece_cid = deals.piece_cid").
		Where("file_ranges.file_id = ?", id)
	if err := query.Find(&deals).Error; err != nil {
		return nil, err
	}

	return deals, nil
}
