package storage

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/stretchr/testify/require"
)

func TestRemoveStorageHandler(t *testing.T) {
	t.Run("remove storage", func(t *testing.T) {
		ctx := context.Background()
		db, closer, err := database.OpenInMemory()
		require.NoError(t, err)
		defer closer.Close()
		tmp := t.TempDir()
		_, err = CreateStorageHandler(ctx, db, "local", "", "name", tmp, nil)
		require.NoError(t, err)
		err = RemoveHandler(ctx, db, "name")
		require.NoError(t, err)
	})
	t.Run("remove storage that does not exist", func(t *testing.T) {
		ctx := context.Background()
		db, closer, err := database.OpenInMemory()
		require.NoError(t, err)
		defer closer.Close()
		err = RemoveHandler(ctx, db, "name")
		require.ErrorIs(t, err, handlererror.ErrNotFound)
	})
	t.Run("remove storage that is still in use", func(t *testing.T) {
		ctx := context.Background()
		db, closer, err := database.OpenInMemory()
		require.NoError(t, err)
		defer closer.Close()
		tmp := t.TempDir()
		_, err = CreateStorageHandler(ctx, db, "local", "", "name", tmp, nil)
		require.NoError(t, err)
		source := model.SourceAttachment{
			StorageID:   1,
			Preparation: &model.Preparation{},
		}
		err = db.Create(&source).Error
		require.NoError(t, err)
		err = RemoveHandler(ctx, db, "name")
		require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
	})
}
