package wallet

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/stretchr/testify/require"
)

func TestRemoveHandler(t *testing.T) {
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		err := db.Create(&model.Wallet{
			ID: "test",
		}).Error
		require.NoError(t, err)
		err = RemoveHandler(ctx, db, "test")
		require.NoError(t, err)
		err = RemoveHandler(ctx, db, "test")
		require.ErrorIs(t, err, handlererror.ErrNotFound)
	})
}
