package wallet

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/stretchr/testify/require"
)

func TestDetachHandler(t *testing.T) {
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()
	ctx := context.Background()
	require.NoError(t, err)
	err = db.Create(&model.Preparation{
		Wallets: []model.Wallet{{
			ID: "test",
		}},
	}).Error
	require.NoError(t, err)

	t.Run("preparation not found", func(t *testing.T) {
		_, err := DetachHandler(ctx, db, 2, "test")
		require.ErrorIs(t, err, handlererror.ErrNotFound)
	})

	t.Run("wallet not found", func(t *testing.T) {
		_, err := DetachHandler(ctx, db, 1, "invalid")
		require.ErrorIs(t, err, handlererror.ErrNotFound)
	})

	t.Run("success", func(t *testing.T) {
		preparation, err := DetachHandler(ctx, db, 1, "test")
		require.NoError(t, err)
		require.Len(t, preparation.Wallets, 0)
	})
}
