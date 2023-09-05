package file

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

// GetFileDealsHandler retrieves the deals associated with a given file ID.
//
// The method operates by querying the database using the provided file ID. It starts by selecting
// the relevant file range, joining it with the cars table on the job_id field, and then joining
// with the deals table using the piece_cid field.
//
// Parameters:
// - ctx: The context for managing timeouts and cancellation.
// - db: The gorm.DB instance for database operations.
// - id: The ID of the file for which deals need to be retrieved.
//
// Returns:
// - A slice of model.Deal containing the deals associated with the provided file ID.
// - An error if any issues occur during the database operation.
func (DefaultHandler) GetFileDealsHandler(
	ctx context.Context,
	db *gorm.DB,
	id uint64,
) ([]model.Deal, error) {
	db = db.WithContext(ctx)

	var deals []model.Deal
	query := db.Where("deals.id IN (?)", db.Table("deals").
		Joins("JOIN cars ON deals.piece_cid = cars.piece_cid").
		Joins("JOIN file_ranges ON cars.job_id = file_ranges.job_id").
		Where("file_ranges.file_id = ?", id).
		Distinct("deals.id"))
	if err := query.Find(&deals).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	return deals, nil
}

// @Summary Get all deals that have been made for a file
// @Tags File
// @Accept json
// @Produce json
// @Param id path int true "File ID"
// @Success 200 {array} model.Deal
// @Failure 500 {object} api.HTTPError
// @Router /file/{id}/deals [get]
func _() {}
