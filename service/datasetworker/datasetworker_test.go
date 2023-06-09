package datasetworker

import (
	"context"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/datasource"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"os"
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

func TestDatasetWorkerThread_pack(t *testing.T) {
	temp := t.TempDir()
	file, err := os.Create(temp + "/test.txt")
	assert.NoError(t, err)
	_, err = file.WriteString("test")
	assert.NoError(t, err)
	err = file.Close()
	file, err = os.Create(temp + "/test2.txt")
	assert.NoError(t, err)
	_, err = file.WriteString("test2")
	assert.NoError(t, err)
	err = file.Close()
	assert.NoError(t, err)
	db := database.OpenInMemory()
	thread := DatasetWorkerThread{
		id:                        uuid.New(),
		db:                        db,
		logger:                    logger.With("key", "value"),
		datasourceHandlerResolver: datasource.DefaultHandlerResolver{},
		directoryCache:            map[string]model.Directory{},
		config: DatasetWorkerConfig{
			Concurrency:    1,
			ExitOnComplete: true,
			EnableScan:     true,
			EnablePack:     true,
			EnableDag:      true,
		},
	}
	dataset := model.Dataset{
		Name:       "test",
		MaxSize:    1024,
		OutputDirs: []string{temp},
	}
	err = db.Create(&dataset).Error
	assert.NoError(t, err)
	root := model.Directory{}
	err = db.Create(&root).Error
	assert.NoError(t, err)
	source := model.Source{
		DatasetID:         dataset.ID,
		RootDirectoryID:   root.ID,
		ScanningState:     model.Complete,
		Type:              "local",
		Path:              temp,
		DeleteAfterExport: true,
	}
	err = db.Create(&source).Error
	assert.NoError(t, err)
	source.Dataset = &dataset
	source.RootDirectory = &root
	stat1, _ := os.Stat(temp + "/test.txt")
	stat2, _ := os.Stat(temp + "/test2.txt")
	item2 := &model.Item{
		SourceID:                  source.ID,
		Path:                      "test2.txt",
		Size:                      5,
		DirectoryID:               &root.ID,
		LastModifiedTimestampNano: stat2.ModTime().UTC().UnixNano(),
	}
	err = db.Create(item2).Error
	assert.NoError(t, err)
	chunk := model.Chunk{
		SourceID: source.ID,
		ItemParts: []model.ItemPart{
			{
				Item: &model.Item{
					SourceID:                  source.ID,
					Path:                      "test.txt",
					Size:                      4,
					DirectoryID:               &root.ID,
					LastModifiedTimestampNano: stat1.ModTime().UTC().UnixNano(),
				},
				Offset: 0,
				Length: 4,
			},
		},
	}
	err = db.Create(&chunk).Error
	assert.NoError(t, err)
	parts := []model.ItemPart{{
		ItemID:  item2.ID,
		Offset:  0,
		Length:  2,
		ChunkID: &chunk.ID,
	}, {
		ItemID:  item2.ID,
		Offset:  2,
		Length:  3,
		ChunkID: &chunk.ID,
	}}
	err = db.Create(&parts).Error
	assert.NoError(t, err)
	chunk.Source = &source
	err = db.Preload("ItemParts.Item").Find(&chunk).Error
	assert.NoError(t, err)
	err = thread.pack(context.TODO(), chunk)
	assert.NoError(t, err)
	var cars []model.Car
	err = db.Find(&cars).Error
	assert.NoError(t, err)
	assert.Greater(t, cars[0].PieceSize, int64(0))
	assert.Equal(t, *cars[0].ChunkID, chunk.ID)
	assert.NotEmpty(t, cars[0].FilePath)
	var chunks []model.Chunk
	err = db.Find(&chunks).Error
	assert.NoError(t, err)
	var items []model.Item
	err = db.Find(&items).Error
	assert.NoError(t, err)
	for _, item := range items {
		assert.NotEqual(t, item.CID.String(), "")
	}
	var itemParts []model.ItemPart
	err = db.Find(&itemParts).Error
	assert.NoError(t, err)
	for _, itemPart := range itemParts {
		assert.NotEqual(t, itemPart.CID.String(), "")
	}
	var dirs []model.Directory
	err = db.Find(&dirs).Error
	assert.NoError(t, err)
	for _, dir := range dirs {
		assert.NotEqual(t, dir.CID.String(), "")
	}
	var carBlocks []model.CarBlock
	err = db.Find(&carBlocks).Error
	assert.NoError(t, err)
	assert.Len(t, carBlocks, 3)

	_, err = os.Stat(temp + "/test.txt")
	assert.True(t, os.IsNotExist(err))
	_, err = os.Stat(temp + "/test2.txt")
	assert.True(t, os.IsNotExist(err))
}

func TestDatasetWorkerThread_scan(t *testing.T) {
	ctx := context.Background()
	db := database.OpenInMemory()
	thread := DatasetWorkerThread{
		id:                        uuid.New(),
		db:                        db,
		logger:                    logger.With("key", "value"),
		datasourceHandlerResolver: datasource.DefaultHandlerResolver{},
		directoryCache:            map[string]model.Directory{},
		config: DatasetWorkerConfig{
			Concurrency:    1,
			ExitOnComplete: true,
			EnableScan:     true,
			EnablePack:     true,
			EnableDag:      true,
		},
	}
	dataset := model.Dataset{
		Name:    "test",
		MaxSize: 1024,
	}
	err := db.Create(&dataset).Error
	assert.NoError(t, err)
	root := model.Directory{}
	err = db.Create(&root).Error
	assert.NoError(t, err)
	cmd, _ := os.Getwd()
	source := model.Source{
		DatasetID:       dataset.ID,
		RootDirectoryID: root.ID,
		ScanningState:   model.Ready,
		Type:            "local",
		Path:            cmd,
	}
	err = db.Create(&source).Error
	assert.NoError(t, err)
	source.Dataset = &dataset
	source.RootDirectory = &root
	err = thread.scan(ctx, source, true)
	assert.NoError(t, err)
	var dirs []model.Directory
	err = db.Find(&dirs).Error
	assert.NoError(t, err)
	assert.Greater(t, len(dirs), 0)
	var items []model.Item
	err = db.Find(&items).Error
	assert.NoError(t, err)
	assert.Greater(t, len(items), 0)
	var itemparts []model.ItemPart
	err = db.Find(&itemparts).Error
	assert.NoError(t, err)
	assert.Greater(t, len(itemparts), 0)
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
			ItemID: items[0].ID,
		},
		{
			ItemID: items[1].ID,
		},
		{
			ItemID: items[1].ID,
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
		err = db.Where("1 = 1").Delete(&model.Chunk{}).Error
		assert.NoError(t, err)
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
			assert.NotNil(t, src.Dataset)
			assert.NotNil(t, src.RootDirectory)
		} else {
			assert.Nil(t, src)
		}
	}
}
