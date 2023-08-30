package wallet

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestListAttachedHandler(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&model.Preparation{
			Wallets: []model.Wallet{{
				ID: "test",
			}},
		}).Error
		require.NoError(t, err)

		t.Run("preparation not found", func(t *testing.T) {
			_, err := Default.ListAttachedHandler(ctx, db, 2)
			require.ErrorIs(t, err, handlererror.ErrNotFound)
		})

		t.Run("success", func(t *testing.T) {
			wallets, err := Default.ListAttachedHandler(ctx, db, 1)
			require.NoError(t, err)
			require.Len(t, wallets, 1)
		})
	})
}
