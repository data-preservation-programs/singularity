package storage

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	t.Run("not supported storage type", func(t *testing.T) {
		ctx := context.Background()
		db, closer, err := database.OpenInMemory()
		require.NoError(t, err)
		defer closer.Close()
		_, err = CreateStorageHandler(ctx, db, "not_supported", "", "", "", nil)
		require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
	})
	t.Run("local path", func(t *testing.T) {
		ctx := context.Background()
		db, closer, err := database.OpenInMemory()
		require.NoError(t, err)
		defer closer.Close()
		tmp := t.TempDir()
		storage, err := CreateStorageHandler(ctx, db, "local", "", "name", tmp, nil)
		require.NoError(t, err)
		require.Greater(t, storage.ID, uint32(0))
	})
	t.Run("local path with config", func(t *testing.T) {
		ctx := context.Background()
		db, closer, err := database.OpenInMemory()
		require.NoError(t, err)
		defer closer.Close()
		tmp := t.TempDir()
		storage, err := CreateStorageHandler(ctx, db, "local", "", "name", tmp,
			map[string]string{
				"copy_links": "true",
			})
		require.NoError(t, err)
		require.Greater(t, storage.ID, uint32(0))
	})
	t.Run("local path with invalid config", func(t *testing.T) {
		ctx := context.Background()
		db, closer, err := database.OpenInMemory()
		require.NoError(t, err)
		defer closer.Close()
		tmp := t.TempDir()
		_, err = CreateStorageHandler(ctx, db, "local", "", "name", tmp,
			map[string]string{
				"copy_links": "invalid",
			})
		require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
	})

	t.Run("local path with inaccessible path", func(t *testing.T) {
		ctx := context.Background()
		db, closer, err := database.OpenInMemory()
		require.NoError(t, err)
		defer closer.Close()
		_, err = CreateStorageHandler(ctx, db, "local", "", "name", "/invalid/path", nil)
		require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
	})

	t.Run("invalid provider", func(t *testing.T) {
		ctx := context.Background()
		tmp := t.TempDir()
		db, closer, err := database.OpenInMemory()
		require.NoError(t, err)
		defer closer.Close()
		_, err = CreateStorageHandler(ctx, db, "local", "invalid", "name", tmp, nil)
		require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
	})

	t.Run("duplicate name", func(t *testing.T) {
		ctx := context.Background()
		tmp := t.TempDir()
		db, closer, err := database.OpenInMemory()
		require.NoError(t, err)
		defer closer.Close()
		_, err = CreateStorageHandler(ctx, db, "local", "", "name", tmp, nil)
		require.NoError(t, err)
		_, err = CreateStorageHandler(ctx, db, "local", "", "name", tmp, nil)
		require.ErrorIs(t, err, handlererror.ErrDuplicateRecord)
	})
}
