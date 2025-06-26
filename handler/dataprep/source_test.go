package dataprep

import (
	"context"
	"strconv"
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

		_, err = Default.AddSourceStorageHandler(ctx, db, "1", "not found")
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

		_, err = Default.AddSourceStorageHandler(ctx, db, "100", "source")
		require.ErrorIs(t, err, handlererror.ErrNotFound)
		require.ErrorContains(t, err, "preparation")
	})
}

func TestAddSourceStorageHandler_AlreadyAttached(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		// Create preparation with source storage already attached
		prep := model.Preparation{
			SourceStorages: []model.Storage{
				{
					Name: "source-already-attached",
					Type: "local",
					Path: "/tmp",
				},
			},
		}
		err := db.Create(&prep).Error
		require.NoError(t, err)

		// Try to attach the same storage again - this should fail
		prepIDStr := strconv.Itoa(int(prep.ID))
		_, err = Default.AddSourceStorageHandler(ctx, db, prepIDStr, "source-already-attached")
		require.ErrorIs(t, err, handlererror.ErrDuplicateRecord)
		require.ErrorContains(t, err, "already")
	})
}

func TestAddSourceStorageHandler_Success(t *testing.T) {
	for _, name := range []string{"1", "name"} {
		t.Run(name, func(t *testing.T) {
			testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
				err := db.Create(&model.Preparation{
					Name: "name",
					SourceStorages: []model.Storage{{
						Name: "source",
					}},
				}).Error
				require.NoError(t, err)

				err = db.Create(&model.Storage{
					Name: "source2",
				}).Error
				require.NoError(t, err)

				preparation, err := Default.AddSourceStorageHandler(ctx, db, name, "source2")
				require.NoError(t, err)
				require.Len(t, preparation.SourceStorages, 2)
			})
		})
	}
}
