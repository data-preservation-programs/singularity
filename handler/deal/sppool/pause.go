package sppool

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

func (DefaultHandler) PauseHandler(
	ctx context.Context,
	db *gorm.DB,
	id uint32,
) (*model.SPPool, error) {
	db = db.WithContext(ctx)
	var pool model.SPPool
	err := db.Preload("Providers").First(&pool, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(handlererror.ErrNotFound, "sp pool %d not found", id)
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	pool.State = model.SPPoolPaused
	if err := database.DoRetry(ctx, func() error {
		return db.Save(&pool).Error
	}); err != nil {
		return nil, errors.WithStack(err)
	}

	// Pause all generated schedules.
	providerIDs := providerIDsFromPool(&pool)
	if len(providerIDs) > 0 {
		if err := database.DoRetry(ctx, func() error {
			return db.Model(&model.Schedule{}).
				Where("sp_pool_provider_id IN ? AND state = ?", providerIDs, model.ScheduleActive).
				Update("state", model.SchedulePaused).Error
		}); err != nil {
			return nil, errors.WithStack(err)
		}
	}

	return &pool, nil
}

func providerIDsFromPool(pool *model.SPPool) []model.SPPoolProviderID {
	ids := make([]model.SPPoolProviderID, len(pool.Providers))
	for i, p := range pool.Providers {
		ids[i] = p.ID
	}
	return ids
}

// @ID PauseSPPool
// @Summary Pause an SP Pool
// @Description Pauses the pool and all its generated schedules.
// @Tags SP Pool
// @Produce json
// @Param id path int true "SP Pool ID"
// @Success 200 {object} model.SPPool
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /sp-pool/{id}/pause [post]
func _() {}
