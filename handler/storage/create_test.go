package storage

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestCreate(t *testing.T) {
	t.Run("not supported storage type", func(t *testing.T) {
		testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
			_, err := Default.CreateStorageHandler(ctx, db, "not_supported", CreateRequest{
				"", "", "", nil, model.ClientConfig{},
			})
			require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
		})
	})
	t.Run("local path", func(t *testing.T) {
		testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
			tmp := t.TempDir()
			storage, err := Default.CreateStorageHandler(ctx, db, "local", CreateRequest{"", "local_path_test", tmp, nil, model.ClientConfig{}})
			require.NoError(t, err)
			require.Greater(t, storage.ID, uint32(0))
		})
	})
	t.Run("local path with config", func(t *testing.T) {
		testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
			tmp := t.TempDir()
			storage, err := Default.CreateStorageHandler(ctx, db, "local", CreateRequest{"", "local_path_config_test", tmp,
				map[string]string{
					"copy_links": "true",
				}, model.ClientConfig{}})
			require.NoError(t, err)
			require.Greater(t, storage.ID, uint32(0))
		})
	})

	t.Run("local path with invalid config", func(t *testing.T) {
		testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
			tmp := t.TempDir()
			_, err := Default.CreateStorageHandler(ctx, db, "local", CreateRequest{"", "invalid_config_test", tmp,
				map[string]string{
					"copy_links": "invalid",
				}, model.ClientConfig{}})
			require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
		})
	})

	t.Run("local path with inaccessible path", func(t *testing.T) {
		testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
			_, err := Default.CreateStorageHandler(ctx, db, "local", CreateRequest{"", "inaccessible_path_test", "/invalid/path", nil, model.ClientConfig{}})
			require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
		})
	})

	t.Run("invalid provider", func(t *testing.T) {
		tmp := t.TempDir()
		testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
			_, err := Default.CreateStorageHandler(ctx, db, "local", CreateRequest{"invalid", "invalid_provider_test", tmp, nil, model.ClientConfig{}})
			require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
		})
	})

	t.Run("duplicate name", func(t *testing.T) {
		tmp := t.TempDir()
		testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
			_, err := Default.CreateStorageHandler(ctx, db, "local", CreateRequest{"", "duplicate_name_test", tmp, nil, model.ClientConfig{}})
			require.NoError(t, err)
			_, err = Default.CreateStorageHandler(ctx, db, "local", CreateRequest{"", "duplicate_name_test", tmp, nil, model.ClientConfig{}})
			require.ErrorIs(t, err, handlererror.ErrDuplicateRecord)
		})
	})

	t.Run("name is digits", func(t *testing.T) {
		tmp := t.TempDir()
		testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
			_, err := Default.CreateStorageHandler(ctx, db, "local", CreateRequest{"", "123", tmp, nil, model.ClientConfig{}})
			require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
		})
	})
}
