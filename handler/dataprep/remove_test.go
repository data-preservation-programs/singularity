package dataprep

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

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

// postgres-only: used to hang on delete due to duplicate cascade paths
// test the handler path and fail fast if it blocks, dialect branch is intentional
func TestRemovePreparationHandler_CascadeCycle_Postgres(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		if db.Dialector.Name() != "postgres" {
			t.Skip("Skip non-Postgres dialect")
			return
		}
		prep := model.Preparation{Name: "pg-prep"}
		require.NoError(t, db.Create(&prep).Error)
		tmpPg := t.TempDir()
		stor := model.Storage{Name: "pg-storage", Type: "local", Path: tmpPg}
		require.NoError(t, db.Create(&stor).Error)
		sa := model.SourceAttachment{PreparationID: prep.ID, StorageID: stor.ID}
		require.NoError(t, db.Create(&sa).Error)
		root := model.Directory{AttachmentID: sa.ID, Name: "", ParentID: nil}
		require.NoError(t, db.Create(&root).Error)
		d1 := model.Directory{AttachmentID: sa.ID, Name: "sub", ParentID: &root.ID}
		require.NoError(t, db.Create(&d1).Error)
		f := model.File{AttachmentID: sa.ID, DirectoryID: &d1.ID, Path: "sub/a.txt", Size: 1}
		require.NoError(t, db.Create(&f).Error)
		fr := model.FileRange{FileID: f.ID, Offset: 0, Length: 1}
		require.NoError(t, db.Create(&fr).Error)

		done := make(chan error, 1)
		go func() {
			done <- Default.RemovePreparationHandler(ctx, db, fmt.Sprintf("%d", prep.ID), RemoveRequest{})
		}()

		select {
		case err := <-done:
			require.NoError(t, err)
		case <-time.After(3 * time.Second):
			t.Fatal("DELETE hung (deadlock) on Postgres")
		}
	})
}

// mysql-only: innodb used to reject the delete with duplicate cascade paths
// test the handler path and expect success, dialect branch is intentional
func TestRemovePreparationHandler_CascadeCycle_MySQL(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		if db.Dialector.Name() != "mysql" {
			t.Skip("Skip non-MySQL dialect")
			return
		}
		prep := model.Preparation{Name: "my-prep"}
		require.NoError(t, db.Create(&prep).Error)
		tmpMy := t.TempDir()
		stor := model.Storage{Name: "my-storage", Type: "local", Path: tmpMy}
		require.NoError(t, db.Create(&stor).Error)
		sa := model.SourceAttachment{PreparationID: prep.ID, StorageID: stor.ID}
		require.NoError(t, db.Create(&sa).Error)
		root := model.Directory{AttachmentID: sa.ID, Name: "", ParentID: nil}
		require.NoError(t, db.Create(&root).Error)
		d1 := model.Directory{AttachmentID: sa.ID, Name: "sub", ParentID: &root.ID}
		require.NoError(t, db.Create(&d1).Error)
		f := model.File{AttachmentID: sa.ID, DirectoryID: &d1.ID, Path: "sub/a.txt", Size: 1}
		require.NoError(t, db.Create(&f).Error)
		fr := model.FileRange{FileID: f.ID, Offset: 0, Length: 1}
		require.NoError(t, db.Create(&fr).Error)

		err := Default.RemovePreparationHandler(ctx, db, fmt.Sprintf("%d", prep.ID), RemoveRequest{})
		require.NoError(t, err)
	})
}
