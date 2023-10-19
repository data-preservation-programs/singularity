package dataprep

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/gotidy/ptr"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestRemovePreparationHandler_NotExist(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := Default.RemovePreparationHandler(ctx, db, "name", RemoveRequest{})
		require.ErrorIs(t, err, handlererror.ErrNotFound)
	})
}

func TestRemovePreparationHandler_HasActiveJobs(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		preparation := model.Preparation{
			SourceStorages: []model.Storage{{}},
		}
		err := db.Create(&preparation).Error
		require.NoError(t, err)
		err = db.Create(&model.Job{
			AttachmentID: model.SourceAttachmentID(1),
			State:        model.Processing,
		}).Error
		require.NoError(t, err)
		err = Default.RemovePreparationHandler(ctx, db, "1", RemoveRequest{})
		require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
		require.ErrorContains(t, err, "active jobs")
	})
}

func TestRemovePreparationHandler_Success(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		tmp := t.TempDir()
		err := os.WriteFile(filepath.Join(tmp, "1.car"), []byte("1"), 0o644)
		require.NoError(t, err)
		storages := []model.Storage{{}, {
			Type: "local", Path: tmp, Name: "output",
		}}
		err = db.Create(storages).Error
		require.NoError(t, err)
		preparation := model.Preparation{
			SourceStorages: []model.Storage{storages[0]},
			OutputStorages: []model.Storage{storages[1]},
		}
		err = db.Create(&preparation).Error
		require.NoError(t, err)
		err = db.Create(&model.Job{
			AttachmentID: model.SourceAttachmentID(1),
			State:        model.Complete,
		}).Error
		require.NoError(t, err)
		err = db.Create(&model.Car{
			Storage:       ptr.Of(storages[1]),
			StoragePath:   "1.car",
			PreparationID: 1,
			AttachmentID:  ptr.Of(model.SourceAttachmentID(1)),
			JobID:         ptr.Of(model.JobID(1)),
		}).Error
		require.NoError(t, err)
		err = db.Create(&model.CarBlock{
			CarID: 1,
			File: &model.File{
				AttachmentID: 1,
				Directory: &model.Directory{
					AttachmentID: 1,
				},
				FileRanges: []model.FileRange{
					{},
				},
			},
		}).Error
		require.NoError(t, err)
		err = db.Create(&model.Schedule{
			PreparationID: 1,
		}).Error
		require.NoError(t, err)

		for _, model := range []interface{}{&model.Preparation{}, &model.Job{}, &model.Car{}, &model.CarBlock{},
			&model.Schedule{}, &model.File{}, &model.FileRange{}, &model.Directory{}} {
			var count int64
			err = db.Model(model).Count(&count).Error
			require.NoError(t, err)
			require.Greater(t, count, int64(0))
		}

		err = Default.RemovePreparationHandler(ctx, db, "1", RemoveRequest{RemoveCars: true})
		require.NoError(t, err)

		for _, model := range []interface{}{&model.Preparation{}, &model.Job{}, &model.Car{}, &model.CarBlock{},
			&model.Schedule{}, &model.File{}, &model.FileRange{}, &model.Directory{}} {
			var count int64
			err = db.Model(model).Count(&count).Error
			require.NoError(t, err)
			require.Zero(t, count)
		}

		entries, err := os.ReadDir(tmp)
		require.NoError(t, err)
		require.Len(t, entries, 0)
	})
}
