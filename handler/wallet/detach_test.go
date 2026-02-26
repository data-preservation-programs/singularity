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

func TestDetachHandler(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		w := model.Wallet{Address: "f0test", KeyPath: "/tmp/key", KeyStore: "local"}
		require.NoError(t, db.Create(&w).Error)
		require.NoError(t, db.Create(&model.Preparation{WalletID: &w.ID, Wallet: &w}).Error)

		t.Run("preparation not found", func(t *testing.T) {
			_, err := Default.DetachHandler(ctx, db, "2", "f0test")
			require.ErrorIs(t, err, handlererror.ErrNotFound)
		})

		t.Run("wallet not found", func(t *testing.T) {
			_, err := Default.DetachHandler(ctx, db, "1", "invalid")
			require.ErrorIs(t, err, handlererror.ErrNotFound)
		})

		t.Run("success", func(t *testing.T) {
			preparation, err := Default.DetachHandler(ctx, db, "1", "f0test")
			require.NoError(t, err)
			require.Nil(t, preparation.WalletID)
		})
	})
}
