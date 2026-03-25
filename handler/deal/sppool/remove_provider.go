package sppool

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

func (DefaultHandler) RemoveProviderHandler(
	ctx context.Context,
	db *gorm.DB,
	id uint32,
	providerID uint32,
) error {
	db = db.WithContext(ctx)

	// Verify pool exists.
	var pool model.SPPool
	err := db.First(&pool, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.Wrapf(handlererror.ErrNotFound, "sp pool %d not found", id)
	}
	if err != nil {
		return errors.WithStack(err)
	}

	// Verify provider exists in this pool.
	var provider model.SPPoolProvider
	err = db.Where("id = ? AND pool_id = ?", providerID, id).First(&provider).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.Wrapf(handlererror.ErrNotFound, "provider %d not found in pool %d", providerID, id)
	}
	if err != nil {
		return errors.WithStack(err)
	}

	// Pause and unlink schedules associated with this provider.
	if err := database.DoRetry(ctx, func() error {
		return db.Model(&model.Schedule{}).
			Where("sp_pool_provider_id = ?", provider.ID).
			Updates(map[string]any{
				"state":               model.SchedulePaused,
				"sp_pool_provider_id": nil,
			}).Error
	}); err != nil {
		return errors.WithStack(err)
	}

	// Delete the provider record.
	if err := database.DoRetry(ctx, func() error {
		return db.Delete(&provider).Error
	}); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// @ID RemoveSPPoolProvider
// @Summary Remove a provider from an SP Pool
// @Description Removes the provider. Associated schedules are paused and unlinked, not deleted.
// @Tags SP Pool
// @Param id path int true "SP Pool ID"
// @Param provider_id path int true "Provider ID"
// @Success 204
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /sp-pool/{id}/provider/{provider_id} [delete]
func _() {}
