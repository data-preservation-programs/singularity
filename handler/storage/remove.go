package storage

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

// RemoveHandler deletes the storage entry with the specified name from the database.
// Before deletion, it checks if any attachments are still using the storage,
// and if so, returns an error.
//
// Parameters:
//   - ctx: The context for database transactions and other operations.
//   - db: A pointer to the gorm.DB instance representing the database connection.
//   - name: The ID or name of the storage entry to be deleted.
//
// Returns:
//   - An error, if any occurred during the operation. Returns a specific error
//     if the storage is still in use or if the storage does not exist.
func (DefaultHandler) RemoveHandler(
	ctx context.Context,
	db *gorm.DB,
	name string,
) error {
	db = db.WithContext(ctx)
	err := database.DoRetry(ctx, func() error {
		return db.Transaction(func(db *gorm.DB) error {
			var storage model.Storage
			err := storage.FindByIDOrName(db, name)
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.Wrapf(handlererror.ErrNotFound, "storage %s does not exist", name)
			}
			if err != nil {
				return errors.WithStack(err)
			}

			var sourceCount int64
			var outputCount int64
			err = db.Model(&model.SourceAttachment{}).Where("storage_id = ?", storage.ID).Count(&sourceCount).Error
			if err != nil {
				return errors.WithStack(err)
			}
			err = db.Model(&model.OutputAttachment{}).Where("storage_id = ?", storage.ID).Count(&outputCount).Error
			if err != nil {
				return errors.WithStack(err)
			}

			if sourceCount > 0 || outputCount > 0 {
				return errors.Wrapf(handlererror.ErrInvalidParameter, "storage %s is still in use", name)
			}

			err = db.Delete(&storage).Error
			if err != nil {
				return errors.WithStack(err)
			}
			return nil
		})
	})
	return errors.WithStack(err)
}

// @ID RemoveStorage
// @Summary Remove a storage
// @Tags Storage
// @Param name path string true "Storage ID or name"
// @Success 204
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /storage/{name} [delete]
func _() {}
