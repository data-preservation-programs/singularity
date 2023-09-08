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

func TestRemoveStorageHandler(t *testing.T) {
	t.Run("remove storage", func(t *testing.T) {
		for _, name := range []string{"1", "name"} {
			t.Run(name, func(t *testing.T) {
				testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
					tmp := t.TempDir()
					_, err := Default.CreateStorageHandler(ctx, db, "local", CreateRequest{"", "name", tmp, nil})
					require.NoError(t, err)
					err = Default.RemoveHandler(ctx, db, name)
					require.NoError(t, err)
				})
			})
		}
	})
	t.Run("remove storage that does not exist", func(t *testing.T) {
		testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
			err := Default.RemoveHandler(ctx, db, "name")
			require.ErrorIs(t, err, handlererror.ErrNotFound)
		})
	})
	t.Run("remove storage that is still in use", func(t *testing.T) {
		testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
			tmp := t.TempDir()
			_, err := Default.CreateStorageHandler(ctx, db, "local", CreateRequest{"", "name", tmp, nil})
			require.NoError(t, err)
			source := model.SourceAttachment{
				StorageID:   1,
				Preparation: &model.Preparation{},
			}
			err = db.Create(&source).Error
			require.NoError(t, err)
			err = Default.RemoveHandler(ctx, db, "name")
			require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
		})
	})
}
