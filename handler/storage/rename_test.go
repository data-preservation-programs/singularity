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

func TestRenameStorageHandler(t *testing.T) {
	t.Run("Storage not found", func(t *testing.T) {
		testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
			_, err := Default.RenameStorageHandler(ctx, db, "name", RenameRequest{"new"})
			require.ErrorIs(t, err, handlererror.ErrNotFound)
		})
	})

	t.Run("success", func(t *testing.T) {
		testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
			err := db.Create(&model.Storage{
				Name: "old",
			}).Error
			require.NoError(t, err)
			new, err := Default.RenameStorageHandler(ctx, db, "old", RenameRequest{"new"})
			require.NoError(t, err)
			require.Equal(t, "new", new.Name)
		})
	})
}
