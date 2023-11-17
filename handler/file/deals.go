package file

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

type DealsForFileRange struct {
	FileRange model.FileRange
	Deals     []model.Deal
}

// GetFileDealsHandler retrieves the deals associated with a given file ID.
//
// The method operates by querying the database using the provided file ID. It starts by selecting
// the relevant file range, joining it with the cars table on the job_id field, and then joining
// with the deals table using the piece_cid field.
//
// Parameters:
//   - ctx: The context for managing timeouts and cancellation.
//   - db: The gorm.DB instance for database operations.
//   - id: The ID of the file for which deals need to be retrieved.
//
// Returns:
//   - A slice of DealsForFileRange containing the deals associated with the provided file ID for each FileRange.
//   - An error if any issues occur during the database operation.
func (DefaultHandler) GetFileDealsHandler(
	ctx context.Context,
	db *gorm.DB,
	id uint64,
) ([]DealsForFileRange, error) {
	db = db.WithContext(ctx)
	var result []DealsForFileRange

	var file model.File
	err := db.Preload("FileRanges").First(&file, id).Error
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get file with id %d", id)
	}

	for _, fileRange := range file.FileRanges {
		var deals []model.Deal
		err = db.Model(&model.Deal{}).Where("id in (?)",
			db.Model(&model.Deal{}).Select("deals.id").
				Joins("JOIN cars on deals.piece_cid = cars.piece_cid").
				Where("cars.job_id = ?", fileRange.JobID)).Find(&deals).Error
		if err != nil {
			return nil, errors.Wrapf(err, "failed to get deals for file range with id %d", fileRange.ID)
		}
		result = append(result, DealsForFileRange{
			FileRange: fileRange,
			Deals:     deals,
		})
	}

	return result, nil
}

// @ID GetFileDeals
// @Summary Get all deals that have been made for a file
// @Tags File
// @Accept json
// @Produce json
// @Param id path int true "File ID"
// @Success 200 {array} DealsForFileRange
// @Failure 500 {object} api.HTTPError
// @Router /file/{id}/deals [get]
func _() {}
