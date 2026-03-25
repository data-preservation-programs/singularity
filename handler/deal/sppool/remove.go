package sppool

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

func (DefaultHandler) RemoveHandler(
	ctx context.Context,
	db *gorm.DB,
	id uint32,
) error {
	db = db.WithContext(ctx)
	var pool model.SPPool
	err := db.First(&pool, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.Wrapf(handlererror.ErrNotFound, "sp pool %d not found", id)
	}
	if err != nil {
		return errors.WithStack(err)
	}

	// Unlink all schedules managed by this pool's providers before deletion.
	// The CASCADE on SPPoolProvider will delete the provider rows, but we
	// want to preserve the schedules (SET NULL on sp_pool_provider_id).
	// GORM handles this via the FK constraint, so just delete the pool.
	if err := database.DoRetry(ctx, func() error {
		return db.Delete(&pool).Error
	}); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// @ID RemoveSPPool
// @Summary Remove an SP Pool
// @Description Removes the pool. Generated schedules are preserved but unlinked.
// @Tags SP Pool
// @Param id path int true "SP Pool ID"
// @Success 204
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /sp-pool/{id} [delete]
func _() {}
