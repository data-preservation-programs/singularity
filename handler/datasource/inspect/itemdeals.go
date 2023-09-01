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
	query := db.Where("deals.id IN (?)", db.Table("deals").
		Joins("JOIN cars ON deals.piece_cid = cars.piece_cid").
		Joins("JOIN file_ranges ON cars.pack_job_id = file_ranges.pack_job_id").
		Where("file_ranges.file_id = ?", id).
		Distinct("deals.id"))
	if err := query.Find(&deals).Error; err != nil {
		return nil, err
	}

	return deals, nil
}
