package inspect

import (
	"strconv"

	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

func GetFileDealsHandler(
	db *gorm.DB,
	id string,
) ([]model.Deal, error) {
	return getFileDealsHandler(db, id)
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
	db *gorm.DB,
	id string,
) ([]model.Deal, error) {
	fileID, err := strconv.Atoi(id)
	if err != nil {
		return nil, handler.NewInvalidParameterErr("invalid file id")
	}
	var deals []model.Deal
	query := db.
		Model(&model.FileRange{}).
		Select("deals.*").
		Joins("JOIN cars ON file_ranges.pack_job_id = cars.pack_job_id").
		Joins("JOIN deals ON cars.piece_cid = deals.piece_cid").
		Where("file_ranges.file_id = ?", fileID)
	if err := query.Find(&deals).Error; err != nil {
		return nil, err
	}

	return deals, nil
}
