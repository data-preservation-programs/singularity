package datasetworker

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/service/healthcheck"
	"github.com/google/uuid"
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
	getState := func() healthcheck.State {
		return healthcheck.State{
			WorkType:  model.Packing,
			WorkingOn: "something",
		}
	}
	_, err = healthcheck.Register(ctx, thread.dbNoContext, thread.id, getState, true)
	require.NoError(t, err)
	found, err := thread.findPackWork(ctx)
	require.NoError(t, err)
	require.Nil(t, found)

	err = db.Create(&model.Chunk{
		Source: &model.Source{
			Dataset: &model.Dataset{},
		},
		PackingState: model.Ready,
	}).Error
	require.NoError(t, err)

	found, err = thread.findPackWork(ctx)
	require.NoError(t, err)
	require.NotNil(t, found)

	var existing model.Chunk
	err = db.First(&existing, found.ID).Error
	require.NoError(t, err)
	require.Equal(t, model.Processing, existing.PackingState)
	require.Equal(t, thread.id.String(), *existing.PackingWorkerID)
}

func TestFindDagWork(t *testing.T) {
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()
	thread := &Thread{
		dbNoContext: db,
		config: Config{
			EnableDag: true,
		},
		logger: logger.With("test", true),
		id:     uuid.New(),
	}

	ctx := context.Background()
	getState := func() healthcheck.State {
		return healthcheck.State{
			WorkType:  model.Packing,
			WorkingOn: "something",
		}
	}
	_, err = healthcheck.Register(ctx, thread.dbNoContext, thread.id, getState, true)
	require.NoError(t, err)
	found, err := thread.findDagWork(ctx)
	require.NoError(t, err)
	require.Nil(t, found)

	err = db.Create(&model.Source{
		Dataset:     &model.Dataset{},
		DagGenState: model.Ready,
	}).Error
	require.NoError(t, err)

	found, err = thread.findDagWork(ctx)
	require.NoError(t, err)
	require.NotNil(t, found)

	var existing model.Source
	err = db.First(&existing, found.ID).Error
	require.NoError(t, err)
	require.Equal(t, model.Processing, existing.DagGenState)
	require.Equal(t, thread.id.String(), *existing.DagGenWorkerID)
}

func TestFindScanWork(t *testing.T) {
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()
	thread := &Thread{
		dbNoContext: db,
		config: Config{
			EnableScan: true,
		},
		logger: logger.With("test", true),
		id:     uuid.New(),
	}

	ctx := context.Background()
	getState := func() healthcheck.State {
		return healthcheck.State{
			WorkType:  model.Packing,
			WorkingOn: "something",
		}
	}
	_, err = healthcheck.Register(ctx, thread.dbNoContext, thread.id, getState, true)
	require.NoError(t, err)
	found, err := thread.findScanWork(ctx)
	require.NoError(t, err)
	require.Nil(t, found)

	err = db.Create(&model.Source{
		Dataset:       &model.Dataset{},
		ScanningState: model.Ready,
	}).Error
	require.NoError(t, err)

	found, err = thread.findScanWork(ctx)
	require.NoError(t, err)
	require.NotNil(t, found)

	var existing model.Source
	err = db.First(&existing, found.ID).Error
	require.NoError(t, err)
	require.Equal(t, model.Processing, existing.ScanningState)
	require.Equal(t, thread.id.String(), *existing.ScanningWorkerID)
}
