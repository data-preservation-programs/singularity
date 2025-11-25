package datasetworker

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/service/healthcheck"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/google/uuid"
	"github.com/gotidy/ptr"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestFindPackWork(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		thread := &Thread{
			dbNoContext: db,
			config: Config{
				EnablePack: true,
			},
			logger: logger.With("test", true),
			id:     uuid.New(),
		}

		_, err := healthcheck.Register(ctx, thread.dbNoContext, thread.id, model.DatasetWorker, true)
		require.NoError(t, err)
		found, err := thread.findJob(ctx, []model.JobType{model.Pack})
		require.NoError(t, err)
		require.Nil(t, found)

		err = db.Create(&model.Preparation{
			SourceStorages: []model.Storage{{
				Name: "source",
			}},
			OutputStorages: []model.Storage{{
				Name: "output",
			}},
		}).Error
		require.NoError(t, err)

		err = db.Create(&model.Job{
			AttachmentID: ptr.Of(model.SourceAttachmentID(1)),
			State:        model.Ready,
			Type:         model.Pack,
		}).Error
		require.NoError(t, err)
		err = db.Create(&model.FileRange{
			JobID: ptr.Of(model.JobID(1)),
			File: &model.File{
				AttachmentID: ptr.Of(model.SourceAttachmentID(1)),
			},
		}).Error
		require.NoError(t, err)

		found, err = thread.findJob(ctx, []model.JobType{model.Pack})
		require.NoError(t, err)
		require.NotNil(t, found)
		require.Len(t, found.FileRanges, 1)
		require.NotNil(t, found.FileRanges[0].File)

		var existing model.Job
		err = db.First(&existing, found.ID).Error
		require.NoError(t, err)
		require.Equal(t, model.Processing, existing.State)
		require.Equal(t, thread.id.String(), *existing.WorkerID)
	})
}
