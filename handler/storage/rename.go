package storage

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util"
	"gorm.io/gorm"
)

type RenameRequest struct {
	Name string `binding:"required" json:"name"`
}

// RenameStorageHandler updates the name of a storage entry in the database.
//
// This handler finds a storage entry by its ID or name and then updates its name with a new name
// provided in the request payload. The new name cannot be entirely numeric or empty.
//
// Parameters:
//   - ctx: The context for managing timeouts and cancellation.
//   - db: The gorm.DB instance for making database queries.
//   - name: The current name or ID of the storage entry to be renamed.
//   - request: A RenameRequest object containing the new name for the storage entry.
//
// Returns:
//   - A pointer to the updated model.Storage entry, reflecting the new name.
//   - An error if any issues occur during the operation, especially if the provided new name is invalid,
//     or if there are database-related errors.
func (DefaultHandler) RenameStorageHandler(
	ctx context.Context,
	db *gorm.DB,
	name string,
	request RenameRequest,
) (*model.Storage, error) {
	db = db.WithContext(ctx)
	if util.IsAllDigits(request.Name) || request.Name == "" {
		return nil, errors.Wrapf(handlererror.ErrInvalidParameter, "storage name %s cannot be all digits or empty", name)
	}

	var storage model.Storage
	err := storage.FindByIDOrName(db, name)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(handlererror.ErrNotFound, "storage %s does not exist", name)
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	storage.Name = request.Name
	err = database.DoRetry(ctx, func() error {
		return db.Model(&model.Storage{}).Where("id = ?", storage.ID).Update("name", storage.Name).Error
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &storage, nil
}

// @ID RenameStorage
// @Summary Rename a storage connection
// @Tags Storage
// @Param name path string true "Storage ID or name"
// @Param request body RenameRequest true "New storage name"
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /storage/{name}/rename [patch]
func _() {}
