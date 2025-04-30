package file

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/scan"
	"gorm.io/gorm"
)

func (DefaultHandler) PrepareToPackFileHandler(
	ctx context.Context,
	db *gorm.DB,
	fileID uint64,
) (int64, error) {
	db = db.WithContext(ctx)
	var file model.File
	err := db.Preload("Attachment.Preparation").Where("id = ?", fileID).First(&file).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, errors.Wrap(handlererror.ErrNotFound, "file not found")
	}
	if err != nil {
		return 0, errors.WithStack(err)
	}

	var remainingParts []model.FileRange
	err = db.Where("file_id = ? AND job_id is null", fileID).Order("file_ranges.id asc").
		Find(&remainingParts).Error
	if err != nil {
		return 0, errors.WithStack(err)
	}

	return scan.PrepareToPackFileRanges(ctx, db, *file.Attachment, remainingParts)
}

// @ID PrepareToPackFile
// @Summary prepare job for a given item
// @Tags File
// @Accept json
// @Produce json
// @Param id path integer true "File ID"
// @Success 200 {object} int64
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /file/{id}/prepare_to_pack [post]
func _() {}
