package storage

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

// ListStoragesHandler fetches all the storage entries from the database.
//
// Parameters:
//   - ctx: The context for database transactions and other operations.
//   - db: A pointer to the gorm.DB instance representing the database connection.
//
// Returns:
//   - A slice containing all Storage model entries found in the database.
//   - An error, if any occurred during the operation.
func (DefaultHandler) ListStoragesHandler(
	ctx context.Context,
	db *gorm.DB) ([]model.Storage, error) {
	db = db.WithContext(ctx)
	var storages []model.Storage
	if err := db.Preload("PreparationsAsSource").Preload("PreparationsAsOutput").Find(&storages).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return storages, nil
}

// @ID ListStorages
// @Summary List all storages
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {array} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /storage [get]
func _() {}
