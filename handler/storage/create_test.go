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
			storage, err := Default.CreateStorageHandler(ctx, db, "local", CreateRequest{"", testutil.GenerateUniqueName("local-path-test"), tmp, nil, model.ClientConfig{}})
			require.NoError(t, err)
			require.Greater(t, storage.ID, uint32(0))
		})
	})
	t.Run("local path with config", func(t *testing.T) {
		testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
			tmp := t.TempDir()
			storage, err := Default.CreateStorageHandler(ctx, db, "local", CreateRequest{"", testutil.GenerateUniqueName("local-path-config-test"), tmp,
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
			_, err := Default.CreateStorageHandler(ctx, db, "local", CreateRequest{"", testutil.GenerateUniqueName("invalid-config-test"), tmp,
				map[string]string{
					"copy_links": "invalid",
				}, model.ClientConfig{}})
			require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
		})
	})

	t.Run("local path with inaccessible path", func(t *testing.T) {
		testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
			_, err := Default.CreateStorageHandler(ctx, db, "local", CreateRequest{"", testutil.GenerateUniqueName("inaccessible-path-test"), "/invalid/path", nil, model.ClientConfig{}})
			require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
		})
	})

	t.Run("invalid provider", func(t *testing.T) {
		tmp := t.TempDir()
		testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
			_, err := Default.CreateStorageHandler(ctx, db, "local", CreateRequest{"invalid", testutil.GenerateUniqueName("invalid-provider-test"), tmp, nil, model.ClientConfig{}})
			require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
		})
	})

	t.Run("duplicate name", func(t *testing.T) {
		tmp := t.TempDir()
		testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
			uniqueName := testutil.GenerateUniqueName("duplicate-test")
			_, err := Default.CreateStorageHandler(ctx, db, "local", CreateRequest{"", uniqueName, tmp, nil, model.ClientConfig{}})
			require.NoError(t, err)
			_, err = Default.CreateStorageHandler(ctx, db, "local", CreateRequest{"", uniqueName, tmp, nil, model.ClientConfig{}})
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
