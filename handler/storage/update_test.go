package storage

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestUpdateStorageHandler(t *testing.T) {
	t.Run("no config provided", func(t *testing.T) {
		testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
			_, err := Default.UpdateStorageHandler(ctx, db, "test", nil)
			require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
		})
	})
	t.Run("storage not found", func(t *testing.T) {
		testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
			_, err := Default.UpdateStorageHandler(ctx, db, "test", make(map[string]string))
			require.ErrorIs(t, err, handlererror.ErrNotFound)
		})
	})
	t.Run("change local path config", func(t *testing.T) {
		for _, name := range []string{"1", "name"} {
			t.Run(name, func(t *testing.T) {
				testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
					tmp := t.TempDir()
					_, err := Default.CreateStorageHandler(ctx, db, "local", CreateRequest{"", "name", tmp, nil})
					require.NoError(t, err)
					storage, err := Default.UpdateStorageHandler(ctx, db, name, map[string]string{
						"copy_links": "true",
					})
					require.NoError(t, err)
					require.Equal(t, "true", storage.Config["copy_links"])
				})
			})
		}
	})
	t.Run("invalid config", func(t *testing.T) {
		testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
			tmp := t.TempDir()
			_, err := Default.CreateStorageHandler(ctx, db, "local", CreateRequest{"", "name", tmp, nil})
			require.NoError(t, err)
			_, err = Default.UpdateStorageHandler(ctx, db, "name", map[string]string{
				"copy_links": "invalid",
			})
			require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
		})
	})
}
