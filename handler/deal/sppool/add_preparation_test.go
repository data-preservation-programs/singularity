package sppool

import (
	"context"
	"testing"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestAddPreparationHandler(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		pool := model.SPPool{Name: "pool", State: model.SPPoolActive}
		require.NoError(t, db.Create(&pool).Error)

		prep := createPrepWithWallet(t, db, "prep")

		t.Run("success", func(t *testing.T) {
			result, err := Default.AddPreparationHandler(ctx, db, uint32(pool.ID), AddPreparationRequest{
				Preparation: prep.Name,
			})
			require.NoError(t, err)
			require.Equal(t, prep.ID, result.PreparationID)
		})

		t.Run("duplicate", func(t *testing.T) {
			_, err := Default.AddPreparationHandler(ctx, db, uint32(pool.ID), AddPreparationRequest{
				Preparation: prep.Name,
			})
			require.True(t, errors.Is(err, handlererror.ErrDuplicateRecord))
		})

		t.Run("not found pool", func(t *testing.T) {
			_, err := Default.AddPreparationHandler(ctx, db, 99999, AddPreparationRequest{
				Preparation: prep.Name,
			})
			require.True(t, errors.Is(err, handlererror.ErrNotFound))
		})

		t.Run("prep without wallet", func(t *testing.T) {
			noWalletPrep := model.Preparation{Name: "no-wallet"}
			require.NoError(t, db.Create(&noWalletPrep).Error)

			_, err := Default.AddPreparationHandler(ctx, db, uint32(pool.ID), AddPreparationRequest{
				Preparation: noWalletPrep.Name,
			})
			require.True(t, errors.Is(err, handlererror.ErrInvalidParameter))
		})
	})
}
