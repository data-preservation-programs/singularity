package file

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

// GetFileHandler retrieves a file with its associated file ranges from the database using a given file ID.
//
// This function preloads the associated FileRanges for the queried file. If no record is found matching the
// provided ID, it returns an ErrNotFound error.
//
// Parameters:
//   - ctx: The context for managing timeouts and cancellation.
//   - db: The gorm.DB instance for database operations.
//   - id: The ID of the file to be retrieved.
//
// Returns:
//   - A pointer to the retrieved model.File, if found.
//   - An error if any issues occur during the database operation, including when the file is not found.
func (DefaultHandler) GetFileHandler(
	ctx context.Context,
	db *gorm.DB,
	id uint64,
) (*model.File, error) {
	db = db.WithContext(ctx)
	var file model.File
	err := db.Preload("FileRanges").First(&file, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(handlererror.ErrNotFound, "file '%d' does not exist", id)
	}
	return &file, errors.WithStack(err)
}

// @ID GetFile
// @Summary Get details about a file
// @Tags File
// @Accept json
// @Produce json
// @Param id path int true "File ID"
// @Success 200 {object} model.File
// @Failure 500 {object} api.HTTPError
// @Router /file/{id} [get]
func _() {}
