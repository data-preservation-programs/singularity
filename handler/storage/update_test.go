package storage

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/stretchr/testify/require"
)

func TestUpdateStorageHandler(t *testing.T) {
	ctx := context.Background()
	t.Run("no config provided", func(t *testing.T) {
		db, closer, err := database.OpenInMemory()
		require.NoError(t, err)
		defer closer.Close()
		_, err = UpdateStorageHandler(ctx, db, "test", nil)
		require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
	})
	t.Run("storage not found", func(t *testing.T) {
		db, closer, err := database.OpenInMemory()
		require.NoError(t, err)
		defer closer.Close()
		_, err = UpdateStorageHandler(ctx, db, "test", make(map[string]string))
		require.ErrorIs(t, err, handlererror.ErrNotFound)
	})
	t.Run("change local path config", func(t *testing.T) {
		db, closer, err := database.OpenInMemory()
		require.NoError(t, err)
		defer closer.Close()
		tmp := t.TempDir()
		_, err = CreateStorageHandler(ctx, db, "local", "", "name", tmp, nil)
		require.NoError(t, err)
		storage, err := UpdateStorageHandler(ctx, db, "name", map[string]string{
			"copy_links": "true",
		})
		require.NoError(t, err)
		require.Equal(t, "true", storage.Config["copy_links"])
	})
	t.Run("invalid config", func(t *testing.T) {
		db, closer, err := database.OpenInMemory()
		require.NoError(t, err)
		defer closer.Close()
		tmp := t.TempDir()
		_, err = CreateStorageHandler(ctx, db, "local", "", "name", tmp, nil)
		require.NoError(t, err)
		_, err = UpdateStorageHandler(ctx, db, "name", map[string]string{
			"copy_links": "invalid",
		})
		require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
	})
}
