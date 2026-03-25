package sppool

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

func (DefaultHandler) RemovePreparationHandler(
	ctx context.Context,
	db *gorm.DB,
	id uint32,
	preparationID uint32,
) error {
	db = db.WithContext(ctx)

	// Verify pool exists.
	var pool model.SPPool
	err := db.Preload("Providers").First(&pool, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.Wrapf(handlererror.ErrNotFound, "sp pool %d not found", id)
	}
	if err != nil {
		return errors.WithStack(err)
	}

	// Verify preparation exists in this pool.
	var poolPrep model.SPPoolPreparation
	err = db.Where("id = ? AND pool_id = ?", preparationID, id).First(&poolPrep).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.Wrapf(handlererror.ErrNotFound, "preparation %d not found in pool %d", preparationID, id)
	}
	if err != nil {
		return errors.WithStack(err)
	}

	// Pause and unlink schedules for this preparation across all pool providers.
	providerIDs := providerIDsFromPool(&pool)
	if len(providerIDs) > 0 {
		if err := database.DoRetry(ctx, func() error {
			return db.Model(&model.Schedule{}).
				Where("sp_pool_provider_id IN ? AND preparation_id = ?", providerIDs, poolPrep.PreparationID).
				Updates(map[string]any{
					"state":               model.SchedulePaused,
					"sp_pool_provider_id": nil,
				}).Error
		}); err != nil {
			return errors.WithStack(err)
		}
	}

	// Delete the pool-preparation link.
	if err := database.DoRetry(ctx, func() error {
		return db.Delete(&poolPrep).Error
	}); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// @ID RemoveSPPoolPreparation
// @Summary Remove a preparation from an SP Pool
// @Description Removes the preparation. Associated schedules are paused and unlinked, not deleted.
// @Tags SP Pool
// @Param id path int true "SP Pool ID"
// @Param preparation_id path int true "Pool Preparation ID"
// @Success 204
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /sp-pool/{id}/preparation/{preparation_id} [delete]
func _() {}
