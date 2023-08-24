package dataprep

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/stretchr/testify/require"
)

func TestStartPackHandler_StorageNotFound(t *testing.T) {
	ctx := context.Background()
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()

	err = db.Create(&model.Preparation{}).Error
	require.NoError(t, err)
	_, err = StartPackHandler(ctx, db, 1, "not found", 1)
	require.ErrorIs(t, err, handlererror.ErrNotFound)
}

func TestStartPackHandler_PreparationNotFound(t *testing.T) {
	ctx := context.Background()
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()

	err = db.Create(&model.Preparation{
		SourceStorages: []model.Storage{{
			Name: "source",
		}},
	}).Error
	require.NoError(t, err)
	_, err = StartPackHandler(ctx, db, 2, "source", 1)
	require.ErrorIs(t, err, handlererror.ErrNotFound)
}

func TestStartPackHandler_JobNotFound(t *testing.T) {
	ctx := context.Background()
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()

	err = db.Create(&model.Preparation{
		SourceStorages: []model.Storage{{
			Name: "source",
		}},
	}).Error
	require.NoError(t, err)
	_, err = StartPackHandler(ctx, db, 1, "source", 1)
	require.ErrorIs(t, err, handlererror.ErrNotFound)
}

func TestStartPackHandler_StartExisting(t *testing.T) {
	ctx := context.Background()
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()

	err = db.Create(&model.Preparation{
		SourceStorages: []model.Storage{{
			Name: "source",
		}},
	}).Error
	require.NoError(t, err)
	err = db.Create(&model.Job{
		AttachmentID: 1,
		State:        model.Error,
		Type:         model.Pack,
	}).Error
	jobs, err := StartPackHandler(ctx, db, 1, "source", 1)
	require.NoError(t, err)
	require.Len(t, jobs, 1)
	require.Equal(t, model.Ready, jobs[0].State)
	require.Equal(t, model.Pack, jobs[0].Type)
}

func TestStartPackHandler_AlreadyProcessing(t *testing.T) {
	ctx := context.Background()
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()

	err = db.Create(&model.Preparation{
		SourceStorages: []model.Storage{{
			Name: "source",
		}},
	}).Error
	require.NoError(t, err)
	err = db.Create(&model.Job{
		AttachmentID: 1,
		State:        model.Processing,
		Type:         model.Pack,
	}).Error
	_, err = StartPackHandler(ctx, db, 1, "source", 1)
	require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
}

func TestStartPackHandler_All(t *testing.T) {
	ctx := context.Background()
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()

	err = db.Create(&model.Preparation{
		SourceStorages: []model.Storage{{
			Name: "source",
		}},
	}).Error
	require.NoError(t, err)
	err = db.Create(&model.Job{
		AttachmentID: 1,
		State:        model.Error,
		Type:         model.Pack,
	}).Error
	jobs, err := StartPackHandler(ctx, db, 1, "source", 0)
	require.NoError(t, err)
	require.Len(t, jobs, 1)
	require.Equal(t, model.Ready, jobs[0].State)
	require.Equal(t, model.Pack, jobs[0].Type)
}

func TestPausePackHandler_All(t *testing.T) {
	ctx := context.Background()
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()

	err = db.Create(&model.Preparation{
		SourceStorages: []model.Storage{{
			Name: "source",
		}},
	}).Error
	require.NoError(t, err)
	err = db.Create(&model.Job{
		AttachmentID: 1,
		State:        model.Ready,
		Type:         model.Pack,
	}).Error
	jobs, err := PausePackHandler(ctx, db, 1, "source", 0)
	require.NoError(t, err)
	require.Len(t, jobs, 1)
	require.Equal(t, model.Paused, jobs[0].State)
	require.Equal(t, model.Pack, jobs[0].Type)
}

func TestPausePackHandler_Existing(t *testing.T) {
	ctx := context.Background()
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()

	err = db.Create(&model.Preparation{
		SourceStorages: []model.Storage{{
			Name: "source",
		}},
	}).Error
	require.NoError(t, err)
	err = db.Create(&model.Job{
		AttachmentID: 1,
		State:        model.Ready,
		Type:         model.Pack,
	}).Error
	jobs, err := PausePackHandler(ctx, db, 1, "source", 1)
	require.NoError(t, err)
	require.Len(t, jobs, 1)
	require.Equal(t, model.Paused, jobs[0].State)
	require.Equal(t, model.Pack, jobs[0].Type)
}

func TestPausePackHandler_JobNotExist(t *testing.T) {
	ctx := context.Background()
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()

	err = db.Create(&model.Preparation{
		SourceStorages: []model.Storage{{
			Name: "source",
		}},
	}).Error
	require.NoError(t, err)
	err = db.Create(&model.Job{
		AttachmentID: 1,
		State:        model.Ready,
		Type:         model.Pack,
	}).Error
	_, err = PausePackHandler(ctx, db, 1, "source", 2)
	require.ErrorIs(t, err, handlererror.ErrNotFound)
}

func TestPausePackHandler_AlreadyPaused(t *testing.T) {
	ctx := context.Background()
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()

	err = db.Create(&model.Preparation{
		SourceStorages: []model.Storage{{
			Name: "source",
		}},
	}).Error
	require.NoError(t, err)
	err = db.Create(&model.Job{
		AttachmentID: 1,
		State:        model.Paused,
		Type:         model.Pack,
	}).Error
	_, err = PausePackHandler(ctx, db, 1, "source", 1)
	require.ErrorIs(t, err, handlererror.ErrInvalidParameter)
}
