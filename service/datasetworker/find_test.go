package datasetworker

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/service/healthcheck"
	"github.com/google/uuid"
	"github.com/gotidy/ptr"
	"github.com/stretchr/testify/require"
)

func TestFindPackWork(t *testing.T) {
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()
	thread := &Thread{
		dbNoContext: db,
		config: Config{
			EnablePack: true,
		},
		logger: logger.With("test", true),
		id:     uuid.New(),
	}

	ctx := context.Background()
	_, err = healthcheck.Register(ctx, thread.dbNoContext, thread.id, model.DatasetWorker, true)
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
		Attachment: &model.SourceAttachment{
			PreparationID: 1,
			StorageID:     1,
		},
		State: model.Ready,
		Type:  model.Pack,
	}).Error
	require.NoError(t, err)
	err = db.Create(&model.FileRange{
		JobID: ptr.Of(uint64(1)),
		File: &model.File{
			AttachmentID: 1,
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
}
