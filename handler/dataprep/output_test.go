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

func TestAddOutputStorageHandler_StorageNotFound(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&model.Preparation{}).Error
		require.NoError(t, err)

		_, err = Default.AddOutputStorageHandler(ctx, db, "1", "not found")
		require.ErrorIs(t, err, handlererror.ErrNotFound)
		require.ErrorContains(t, err, "output storage")
	})
}

func TestAddOutputStorageHandler_PreparationNotFound(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&model.Preparation{
			OutputStorages: []model.Storage{{
				Name: "output",
			}},
		}).Error
		require.NoError(t, err)

		_, err = Default.AddOutputStorageHandler(ctx, db, "100", "output")
		require.ErrorIs(t, err, handlererror.ErrNotFound)
		require.ErrorContains(t, err, "preparation")
	})
}

func TestAddOutputStorageHandler_AlreadyAttached(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		// Create preparation with output storage already attached
		prep := model.Preparation{
			OutputStorages: []model.Storage{
				{
					Name: "output-already-attached",
					Type: "local",
					Path: "/tmp",
				},
			},
		}
		err := db.Create(&prep).Error
		require.NoError(t, err)

		// Try to attach the same storage again - this should fail
		prepIDStr := strconv.Itoa(int(prep.ID))
		_, err = Default.AddOutputStorageHandler(ctx, db, prepIDStr, "output-already-attached")
		require.ErrorIs(t, err, handlererror.ErrDuplicateRecord)
		require.ErrorContains(t, err, "already")
	})
}

func TestAddOutputStorageHandler_Success(t *testing.T) {
	for _, name := range []string{"1", "name"} {
		t.Run(name, func(t *testing.T) {
			testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
				err := db.Create(&model.Preparation{
					Name: "name",
					OutputStorages: []model.Storage{{
						Name: "output",
					}},
				}).Error
				require.NoError(t, err)

				err = db.Create(&model.Storage{
					Name: "output2",
				}).Error
				require.NoError(t, err)

				preparation, err := Default.AddOutputStorageHandler(ctx, db, name, "output2")
				require.NoError(t, err)
				require.Len(t, preparation.OutputStorages, 2)
			})
		})
	}
}

func TestRemoveOutputStorageHandler_StorageNotFound(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&model.Preparation{}).Error
		require.NoError(t, err)

		_, err = Default.RemoveOutputStorageHandler(ctx, db, "1", "not found")
		require.ErrorIs(t, err, handlererror.ErrNotFound)
		require.ErrorContains(t, err, "output storage")
	})
}

func TestRemoveOutputStorageHandler_PreparationNotFound(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&model.Preparation{
			OutputStorages: []model.Storage{{
				Name: "output",
			}},
		}).Error
		require.NoError(t, err)

		_, err = Default.RemoveOutputStorageHandler(ctx, db, "100", "output")
		require.ErrorIs(t, err, handlererror.ErrNotFound)
		require.ErrorContains(t, err, "preparation")
	})
}

func TestRemoveOutputStorageHandler_NotAttached(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&model.Preparation{}).Error
		require.NoError(t, err)

		err = db.Create(&model.Storage{
			Name: "output",
		}).Error
		require.NoError(t, err)

		_, err = Default.RemoveOutputStorageHandler(ctx, db, "1", "output")
		require.ErrorIs(t, err, handlererror.ErrNotFound)
		require.ErrorContains(t, err, "not attached")
	})
}

func TestRemoveOutputStorageHandler_DeleteAfterExportTrue(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&model.Preparation{
			DeleteAfterExport: true,
			OutputStorages: []model.Storage{{
				Name: "output",
			}},
		}).Error
		require.NoError(t, err)

		_, err = Default.RemoveOutputStorageHandler(ctx, db, "1", "output")
		require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
		require.ErrorContains(t, err, "cannot remove the only output storage from a preparation with deleteAfterExport enabled")
	})
}

func TestRemoveOutputStorageHandler_NoInlineWithoutOutput(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&model.Preparation{
			NoInline: true,
			OutputStorages: []model.Storage{{
				Name: "output",
			}},
		}).Error
		require.NoError(t, err)

		_, err = Default.RemoveOutputStorageHandler(ctx, db, "1", "output")
		require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
		require.ErrorContains(t, err, "cannot remove the only output storage from a preparation with inline preparation disabled")
	})
}

func TestRemoveOutputStorageHandler_Success(t *testing.T) {
	for _, name := range []string{"1", "name"} {
		t.Run(name, func(t *testing.T) {
			testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
				err := db.Create(&model.Preparation{
					Name: "name",
					OutputStorages: []model.Storage{{
						Name: "output",
					}},
				}).Error
				require.NoError(t, err)

				preparation, err := Default.RemoveOutputStorageHandler(ctx, db, name, "output")
				require.NoError(t, err)
				require.Len(t, preparation.OutputStorages, 0)
			})
		})
	}
}
