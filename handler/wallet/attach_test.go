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

func TestAttachHandler(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&model.Wallet{
			ID: "test",
		}).Error
		require.NoError(t, err)
		err = db.Create(&model.Preparation{}).Error
		require.NoError(t, err)

		t.Run("preparation not found", func(t *testing.T) {
			_, err := Default.AttachHandler(ctx, db, "2", "test")
			require.ErrorIs(t, err, handlererror.ErrNotFound)
		})

		t.Run("wallet not found", func(t *testing.T) {
			_, err := Default.AttachHandler(ctx, db, "1", "invalid")
			require.ErrorIs(t, err, handlererror.ErrNotFound)
		})

		t.Run("success", func(t *testing.T) {
			preparation, err := Default.AttachHandler(ctx, db, "1", "test")
			require.NoError(t, err)
			require.Len(t, preparation.Wallets, 1)
		})
	})
}
