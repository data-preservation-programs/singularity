package datasetworker

import (
	"context"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/datasource"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDatasetWorkerRun(t *testing.T) {
	db := database.OpenInMemory()
	worker := NewDatasetWorker(db, DatasetWorkerConfig{
		Concurrency:    1,
		ExitOnComplete: true,
	})
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	err := worker.Run(ctx)
	assert.NoError(t, err)
}

func TestDatasetWorkerThread_findPackWork(t *testing.T) {
	db := database.OpenInMemory()
	thread := DatasetWorkerThread{
		id:                        uuid.New(),
		db:                        db,
		logger:                    logger.With("key", "value"),
		datasourceHandlerResolver: datasource.DefaultHandlerResolver{},
		config: DatasetWorkerConfig{
			Concurrency:    1,
			ExitOnComplete: true,
			EnableScan:     true,
			EnablePack:     true,
			EnableDag:      true,
		},
	}
	worker := model.Worker{
		ID: thread.id.String(),
	}
	err := db.Create(&worker).Error
	assert.NoError(t, err)
	dataset := model.Dataset{
		Name: "test",
	}
	err = db.Create(&dataset).Error
	assert.NoError(t, err)
	root := model.Directory{}
	err = db.Create(&root).Error
	assert.NoError(t, err)
	source := model.Source{
		DatasetID:       dataset.ID,
		RootDirectoryID: root.ID,
		ScanningState:   model.Complete,
	}
	err = db.Create(&source).Error
	assert.NoError(t, err)
	items := []model.Item{
		{
			SourceID: source.ID,
		},
		{
			SourceID: source.ID,
		},
	}
	err = db.Create(&items).Error
	assert.NoError(t, err)
	itemParts := []model.ItemPart{
		{
			ItemID:   items[0].ID,
			SourceID: source.ID,
		},
		{
			ItemID:   items[1].ID,
			SourceID: source.ID,
		},
		{
			ItemID:   items[1].ID,
			SourceID: source.ID,
		},
	}
	err = db.Create(&itemParts).Error
	assert.NoError(t, err)
	chunks := map[*model.Chunk]bool{
		&model.Chunk{
			PackingState: model.Ready,
		}: true,
		&model.Chunk{
			PackingState: model.Processing,
		}: true,
		&model.Chunk{
			PackingState: model.Error,
		}: false,
		&model.Chunk{
			PackingState:    model.Processing,
			PackingWorkerID: &worker.ID,
		}: false,
	}
	for chunk, shouldBeFound := range chunks {
		err := db.Where("1 = 1").Delete(&model.Chunk{}).Error
		assert.NoError(t, err)
		chunk.SourceID = source.ID
		err = db.Create(chunk).Error
		for _, itemPart := range itemParts {
			itemPart.ChunkID = &chunk.ID
			err = db.Save(&itemPart).Error
			assert.NoError(t, err)
		}
		assert.NoError(t, err)
		ck, err := thread.findPackWork()
		assert.NoError(t, err)
		if shouldBeFound {
			assert.NotNil(t, ck)
			assert.NotNil(t, ck.Source)
			assert.NotNil(t, ck.Source.Dataset)
			assert.NotNil(t, ck.ItemParts)
			assert.NotNil(t, ck.ItemParts[0].Item)
		} else {
			assert.Nil(t, ck)
		}
	}
}

func TestDatasetWorkerThread_findScanWork(t *testing.T) {
	db := database.OpenInMemory()
	thread := DatasetWorkerThread{
		id:                        uuid.New(),
		db:                        db,
		logger:                    logger.With("key", "value"),
		datasourceHandlerResolver: datasource.DefaultHandlerResolver{},
		config: DatasetWorkerConfig{
			Concurrency:    1,
			ExitOnComplete: true,
			EnableScan:     true,
			EnablePack:     true,
			EnableDag:      true,
		},
	}
	worker := model.Worker{
		ID: thread.id.String(),
	}
	err := db.Create(&worker).Error
	assert.NoError(t, err)
	dataset := model.Dataset{
		Name: "test",
	}
	err = db.Create(&dataset).Error
	assert.NoError(t, err)
	root := model.Directory{}
	err = db.Create(&root).Error
	assert.NoError(t, err)
	sources := map[*model.Source]bool{
		// data source that is ready to be scanned
		&model.Source{
			ScanningState: model.Ready,
		}: true,
		// data source that is being scanned but does not have a worker id
		&model.Source{
			ScanningState:    model.Processing,
			ScanningWorkerID: nil,
		}: true,
		// data source that has completed scanning and should be scanned again
		&model.Source{
			ScanningState:        model.Complete,
			LastScannedTimestamp: 0,
			ScanIntervalSeconds:  100,
		}: true,
		// data source that has completed scanning and should not be scanned again
		&model.Source{
			ScanningState:        model.Complete,
			LastScannedTimestamp: 0,
			ScanIntervalSeconds:  0,
		}: false,
		// data source that has completed scanning and should not be scanned again
		&model.Source{
			ScanningState:        model.Complete,
			LastScannedTimestamp: time.Now().Unix(),
			ScanIntervalSeconds:  100,
		}: false,
		// data source that has errored
		&model.Source{
			ScanningState: model.Error,
		}: false,
		// data source that has worker working on it
		&model.Source{
			ScanningState:    model.Processing,
			ScanningWorkerID: &worker.ID,
		}: false,
	}
	for source, shouldBeFound := range sources {
		err := db.Where("1 = 1").Delete(&model.Source{}).Error
		assert.NoError(t, err)
		source.DatasetID = dataset.ID
		source.RootDirectoryID = root.ID
		err = db.Create(source).Error
		assert.NoError(t, err)
		src, err := thread.findScanWork()
		assert.NoError(t, err)
		if shouldBeFound {
			assert.NotNil(t, src)
		} else {
			assert.Nil(t, src)
		}
	}
}
