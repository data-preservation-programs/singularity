package sppool

import (
	"context"
	"fmt"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

// scheduleKey uniquely identifies a desired schedule within a pool.
type scheduleKey struct {
	SPPoolProviderID model.SPPoolProviderID
	PreparationID    model.PreparationID
	DealType         model.DealType
}

// reconcile ensures that the set of schedules in the database matches the
// desired state defined by the pool's providers (with their policies) and
// preparations. It is idempotent and safe to call after any pool mutation.
func reconcile(ctx context.Context, db *gorm.DB, pool *model.SPPool) error {
	db = db.WithContext(ctx)

	// Ensure associations are loaded.
	if pool.Providers == nil {
		if err := db.Model(pool).Association("Providers").Find(&pool.Providers); err != nil {
			return errors.WithStack(err)
		}
	}
	if pool.Preparations == nil {
		if err := db.Model(pool).Association("Preparations").Find(&pool.Preparations); err != nil {
			return errors.WithStack(err)
		}
	}

	// Build the desired set of schedules from providers × preparations × policy.
	desired := make(map[scheduleKey]model.SPPoolProvider)
	for _, prov := range pool.Providers {
		for _, prep := range pool.Preparations {
			for dt := range prov.Policy {
				key := scheduleKey{
					SPPoolProviderID: prov.ID,
					PreparationID:    prep.PreparationID,
					DealType:         dt,
				}
				desired[key] = prov
			}
		}
	}

	// Find all existing pool-managed schedules.
	providerIDs := make([]model.SPPoolProviderID, len(pool.Providers))
	for i, p := range pool.Providers {
		providerIDs[i] = p.ID
	}

	var existing []model.Schedule
	if len(providerIDs) > 0 {
		if err := db.Where("sp_pool_provider_id IN ?", providerIDs).Find(&existing).Error; err != nil {
			return errors.WithStack(err)
		}
	}

	// Index existing schedules by key.
	existingByKey := make(map[scheduleKey]*model.Schedule, len(existing))
	for i := range existing {
		s := &existing[i]
		if s.SPPoolProviderID == nil {
			continue
		}
		key := scheduleKey{
			SPPoolProviderID: *s.SPPoolProviderID,
			PreparationID:    s.PreparationID,
			DealType:         s.DealType,
		}
		existingByKey[key] = s
	}

	// Create missing schedules and ensure existing ones match desired state.
	for key, prov := range desired {
		if s, ok := existingByKey[key]; ok {
			// Schedule exists — ensure correct state.
			desiredState := model.ScheduleActive
			if pool.State == model.SPPoolPaused {
				desiredState = model.SchedulePaused
			}
			if s.State != desiredState {
				if err := database.DoRetry(ctx, func() error {
					return db.Model(s).Update("state", desiredState).Error
				}); err != nil {
					return errors.WithStack(err)
				}
			}
			delete(existingByKey, key)
		} else {
			// Need to create a new schedule.
			urlTemplate := pool.URLTemplate
			if prov.URLTemplate != nil {
				urlTemplate = *prov.URLTemplate
			}

			provID := prov.ID
			state := model.ScheduleActive
			if pool.State == model.SPPoolPaused {
				state = model.SchedulePaused
			}

			schedule := model.Schedule{
				PreparationID:         key.PreparationID,
				Provider:              prov.Provider,
				DealType:              key.DealType,
				State:                 state,
				URLTemplate:           urlTemplate,
				HTTPHeaders:           pool.HTTPHeaders,
				PricePerGBEpoch:       pool.PricePerGBEpoch,
				PricePerGB:            pool.PricePerGB,
				PricePerDeal:          pool.PricePerDeal,
				Verified:              pool.Verified,
				KeepUnsealed:          pool.KeepUnsealed,
				AnnounceToIPNI:        pool.AnnounceToIPNI,
				StartDelay:            pool.StartDelay,
				Duration:              pool.Duration,
				ScheduleCron:          pool.ScheduleCron,
				ScheduleCronPerpetual: pool.ScheduleCronPerpetual,
				ScheduleDealNumber:    pool.ScheduleDealNumber,
				ScheduleDealSize:      pool.ScheduleDealSize,
				MaxPendingDealNumber:  pool.MaxPendingDealNumber,
				MaxPendingDealSize:    pool.MaxPendingDealSize,
				Force:                 pool.Force,
				Notes:                 fmt.Sprintf("auto:sp-pool-%d", pool.ID),
				SPPoolProviderID:      &provID,
			}

			if err := database.DoRetry(ctx, func() error {
				return db.Create(&schedule).Error
			}); err != nil {
				return errors.WithStack(err)
			}
		}
	}

	// Any remaining existing schedules are no longer desired — pause and unlink.
	for _, s := range existingByKey {
		if err := database.DoRetry(ctx, func() error {
			return db.Model(s).Updates(map[string]any{
				"state":               model.SchedulePaused,
				"sp_pool_provider_id": nil,
			}).Error
		}); err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}
