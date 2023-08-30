package dataprep

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestAddSourceStorageHandler_StorageNotFound(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&model.Preparation{}).Error
		require.NoError(t, err)

		_, err = Default.AddSourceStorageHandler(ctx, db, 1, "not found")
		require.ErrorIs(t, err, handlererror.ErrNotFound)
		require.ErrorContains(t, err, "source storage")
	})
}

func TestAddSourceStorageHandler_PreparationNotFound(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&model.Preparation{
			SourceStorages: []model.Storage{{
				Name: "source",
			}},
		}).Error
		require.NoError(t, err)

		_, err = Default.AddSourceStorageHandler(ctx, db, 100, "source")
		require.ErrorIs(t, err, handlererror.ErrNotFound)
		require.ErrorContains(t, err, "preparation")
	})
}

func TestAddSourceStorageHandler_AlreadyAttached(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&model.Preparation{
			SourceStorages: []model.Storage{{
				Name: "source",
			}},
		}).Error
		require.NoError(t, err)

		_, err = Default.AddSourceStorageHandler(ctx, db, 1, "source")
		require.ErrorIs(t, err, handlererror.ErrDuplicateRecord)
		require.ErrorContains(t, err, "already")
	})
}

func TestAddSourceStorageHandler_Success(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&model.Preparation{
			SourceStorages: []model.Storage{{
				Name: "source",
			}},
		}).Error
		require.NoError(t, err)

		err = db.Create(&model.Storage{
			Name: "source2",
		}).Error
		require.NoError(t, err)

		preparation, err := Default.AddSourceStorageHandler(ctx, db, 1, "source2")
		require.NoError(t, err)
		require.Len(t, preparation.SourceStorages, 2)
	})
}
