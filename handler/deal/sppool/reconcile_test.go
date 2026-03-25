package sppool

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

// helper: create a wallet and preparation with that wallet attached
func createPrepWithWallet(t *testing.T, db *gorm.DB, name string) model.Preparation {
	t.Helper()
	w := model.Wallet{Address: "f01" + name, KeyPath: "/tmp/key-" + name, KeyStore: "local"}
	require.NoError(t, db.Create(&w).Error)
	prep := model.Preparation{Name: name, WalletID: &w.ID}
	require.NoError(t, db.Create(&prep).Error)
	return prep
}

func TestReconcile_BasicCrossProduct(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		prep1 := createPrepWithWallet(t, db, "prep1")
		prep2 := createPrepWithWallet(t, db, "prep2")

		pool := model.SPPool{Name: "test-pool", State: model.SPPoolActive}
		require.NoError(t, db.Create(&pool).Error)

		prov1 := model.SPPoolProvider{
			PoolID:   pool.ID,
			Provider: "f01000",
			Policy:   model.ReplicationPolicy{model.DealTypeMarket: 1},
		}
		prov2 := model.SPPoolProvider{
			PoolID:   pool.ID,
			Provider: "f02000",
			Policy:   model.ReplicationPolicy{model.DealTypeMarket: 1, model.DealTypePDP: 1},
		}
		require.NoError(t, db.Create(&prov1).Error)
		require.NoError(t, db.Create(&prov2).Error)

		poolPrep1 := model.SPPoolPreparation{PoolID: pool.ID, PreparationID: prep1.ID}
		poolPrep2 := model.SPPoolPreparation{PoolID: pool.ID, PreparationID: prep2.ID}
		require.NoError(t, db.Create(&poolPrep1).Error)
		require.NoError(t, db.Create(&poolPrep2).Error)

		pool.Providers = []model.SPPoolProvider{prov1, prov2}
		pool.Preparations = []model.SPPoolPreparation{poolPrep1, poolPrep2}

		// 2 preps × (1 market + 2 market+pdp) = 2 + 4 = 6 schedules
		require.NoError(t, reconcile(ctx, db, &pool))

		var schedules []model.Schedule
		require.NoError(t, db.Where("sp_pool_provider_id IS NOT NULL").Find(&schedules).Error)
		require.Len(t, schedules, 6)

		// All should be active.
		for _, s := range schedules {
			require.Equal(t, model.ScheduleActive, s.State)
		}
	})
}

func TestReconcile_Idempotent(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		prep := createPrepWithWallet(t, db, "prep")
		pool := model.SPPool{Name: "pool", State: model.SPPoolActive}
		require.NoError(t, db.Create(&pool).Error)

		prov := model.SPPoolProvider{
			PoolID:   pool.ID,
			Provider: "f01000",
			Policy:   model.ReplicationPolicy{model.DealTypeMarket: 1},
		}
		require.NoError(t, db.Create(&prov).Error)

		poolPrep := model.SPPoolPreparation{PoolID: pool.ID, PreparationID: prep.ID}
		require.NoError(t, db.Create(&poolPrep).Error)

		pool.Providers = []model.SPPoolProvider{prov}
		pool.Preparations = []model.SPPoolPreparation{poolPrep}

		require.NoError(t, reconcile(ctx, db, &pool))
		require.NoError(t, reconcile(ctx, db, &pool))

		var count int64
		require.NoError(t, db.Model(&model.Schedule{}).Where("sp_pool_provider_id IS NOT NULL").Count(&count).Error)
		require.Equal(t, int64(1), count)
	})
}

func TestReconcile_PausedPool(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		prep := createPrepWithWallet(t, db, "prep")
		pool := model.SPPool{Name: "pool", State: model.SPPoolPaused}
		require.NoError(t, db.Create(&pool).Error)

		prov := model.SPPoolProvider{
			PoolID:   pool.ID,
			Provider: "f01000",
			Policy:   model.ReplicationPolicy{model.DealTypeMarket: 1},
		}
		require.NoError(t, db.Create(&prov).Error)

		poolPrep := model.SPPoolPreparation{PoolID: pool.ID, PreparationID: prep.ID}
		require.NoError(t, db.Create(&poolPrep).Error)

		pool.Providers = []model.SPPoolProvider{prov}
		pool.Preparations = []model.SPPoolPreparation{poolPrep}

		require.NoError(t, reconcile(ctx, db, &pool))

		var schedules []model.Schedule
		require.NoError(t, db.Where("sp_pool_provider_id IS NOT NULL").Find(&schedules).Error)
		require.Len(t, schedules, 1)
		require.Equal(t, model.SchedulePaused, schedules[0].State)
	})
}

func TestReconcile_PolicyChangeUnlinks(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		prep := createPrepWithWallet(t, db, "prep")
		pool := model.SPPool{Name: "pool", State: model.SPPoolActive}
		require.NoError(t, db.Create(&pool).Error)

		prov := model.SPPoolProvider{
			PoolID:   pool.ID,
			Provider: "f01000",
			Policy:   model.ReplicationPolicy{model.DealTypeMarket: 1, model.DealTypePDP: 1},
		}
		require.NoError(t, db.Create(&prov).Error)

		poolPrep := model.SPPoolPreparation{PoolID: pool.ID, PreparationID: prep.ID}
		require.NoError(t, db.Create(&poolPrep).Error)

		pool.Providers = []model.SPPoolProvider{prov}
		pool.Preparations = []model.SPPoolPreparation{poolPrep}

		// Initial reconcile: 1 market + 1 PDP = 2 schedules.
		require.NoError(t, reconcile(ctx, db, &pool))

		var initialCount int64
		require.NoError(t, db.Model(&model.Schedule{}).Where("sp_pool_provider_id IS NOT NULL").Count(&initialCount).Error)
		require.Equal(t, int64(2), initialCount)

		// Change policy to market-only (remove PDP).
		prov.Policy = model.ReplicationPolicy{model.DealTypeMarket: 1}
		require.NoError(t, db.Save(&prov).Error)
		pool.Providers = []model.SPPoolProvider{prov}

		require.NoError(t, reconcile(ctx, db, &pool))

		// 1 linked (market), 1 unlinked+paused (pdp).
		var linkedCount int64
		require.NoError(t, db.Model(&model.Schedule{}).Where("sp_pool_provider_id IS NOT NULL").Count(&linkedCount).Error)
		require.Equal(t, int64(1), linkedCount)

		var unlinked []model.Schedule
		require.NoError(t, db.Where("sp_pool_provider_id IS NULL AND deal_type = ?", model.DealTypePDP).Find(&unlinked).Error)
		require.Len(t, unlinked, 1)
		require.Equal(t, model.SchedulePaused, unlinked[0].State)
	})
}

func TestReconcile_EmptyPool(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		pool := model.SPPool{Name: "empty", State: model.SPPoolActive}
		require.NoError(t, db.Create(&pool).Error)

		// No providers or preparations — should be a no-op.
		require.NoError(t, reconcile(ctx, db, &pool))

		var count int64
		require.NoError(t, db.Model(&model.Schedule{}).Count(&count).Error)
		require.Equal(t, int64(0), count)
	})
}
