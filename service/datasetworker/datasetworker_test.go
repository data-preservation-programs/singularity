package datasetworker

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestDatasetWorkerRun(t *testing.T) {
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()
	worker := NewWorker(db, Config{
		Concurrency:    1,
		ExitOnComplete: true,
	})
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	err = worker.Run(ctx)
	require.NoError(t, err)
}

func TestDatasetWorker_HandleScanWork_Failure(t *testing.T) {
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()
	worker := NewWorker(db, Config{
		Concurrency:    1,
		ExitOnComplete: true,
		ExitOnError:    true,
		EnableScan:     true,
	})
	source := &model.Source{
		Dataset:       &model.Preparation{},
		Type:          "invalid",
		Path:          "",
		ScanningState: model.Ready,
	}
	err = db.Create(source).Error
	require.NoError(t, err)
	dir := &model.Directory{SourceID: source.ID, Name: ""}
	err = db.Create(dir).Error
	require.NoError(t, err)
	err = worker.Run(context.Background())
	require.ErrorContains(t, err, "didn't find backend")
}

func TestDatasetWorker_HandleScanWork_Success(t *testing.T) {
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()
	tmp := t.TempDir()
	worker := NewWorker(db, Config{
		Concurrency:    1,
		ExitOnComplete: true,
		ExitOnError:    true,
		EnableScan:     true,
	})
	source := &model.Source{
		Dataset:       &model.Preparation{},
		Type:          "local",
		Path:          tmp,
		ScanningState: model.Ready,
	}
	err = db.Create(source).Error
	require.NoError(t, err)
	dir := &model.Directory{SourceID: source.ID, Name: tmp}
	err = db.Create(dir).Error
	require.NoError(t, err)
	err = worker.Run(context.Background())
	require.NoError(t, err)
	var found model.Source
	err = db.Model(&model.Source{}).First(&found).Error
	require.NoError(t, err)
	require.Equal(t, model.Complete, found.ScanningState)
	require.Nil(t, found.ScanningWorkerID)
	require.Equal(t, "", found.ErrorMessage)
	require.Equal(t, "", found.LastScannedPath)
	require.Greater(t, found.LastScannedTimestamp, int64(0))
}

func TestDatasetWorkerThread_pack(t *testing.T) {
	temp := t.TempDir()
	file, err := os.Create(temp + "/test.txt")
	require.NoError(t, err)
	_, err = file.WriteString("test")
	require.NoError(t, err)
	require.NoError(t, file.Close())
	file, err = os.Create(temp + "/test2.txt")
	require.NoError(t, err)
	_, err = file.WriteString("test2")
	require.NoError(t, err)
	err = file.Close()
	require.NoError(t, err)
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()
	thread := Thread{
		id:                        uuid.New(),
		dbNoContext:               db,
		logger:                    logger.With("key", "value"),
		datasourceHandlerResolver: storagesystem.DefaultHandlerResolver{},
		config: Config{
			Concurrency:    1,
			ExitOnComplete: true,
			EnableScan:     true,
			EnablePack:     true,
			EnableDag:      true,
			ExitOnError:    true,
		},
	}
	dataset := model.Preparation{
		Name:       "test",
		MaxSize:    1024,
		OutputDirs: []string{temp},
	}
	err = db.Create(&dataset).Error
	require.NoError(t, err)
	source := model.Source{
		DatasetID:         dataset.ID,
		ScanningState:     model.Complete,
		Type:              "local",
		Path:              temp,
		DeleteAfterExport: true,
	}
	err = db.Create(&source).Error
	require.NoError(t, err)
	source.Dataset = &dataset
	root := model.Directory{
		SourceID: source.ID,
	}
	err = db.Create(&root).Error
	require.NoError(t, err)
	stat1, _ := os.Stat(temp + "/test.txt")
	stat2, _ := os.Stat(temp + "/test2.txt")
	file2 := &model.File{
		SourceID:         source.ID,
		Path:             "test2.txt",
		Size:             5,
		DirectoryID:      &root.ID,
		LastModifiedNano: stat2.ModTime().UTC().UnixNano(),
	}
	err = db.Create(file2).Error
	require.NoError(t, err)
	packJob := model.PackJob{
		SourceID: source.ID,
		FileRanges: []model.FileRange{
			{
				File: &model.File{
					SourceID:         source.ID,
					Path:             "test.txt",
					Size:             4,
					DirectoryID:      &root.ID,
					LastModifiedNano: stat1.ModTime().UTC().UnixNano(),
				},
				Offset: 0,
				Length: 4,
			},
		},
	}
	err = db.Create(&packJob).Error
	require.NoError(t, err)
	parts := []model.FileRange{{
		FileID:    file2.ID,
		Offset:    0,
		Length:    2,
		PackJobID: &packJob.ID,
	}, {
		FileID:    file2.ID,
		Offset:    2,
		Length:    3,
		PackJobID: &packJob.ID,
	}}
	err = db.Create(&parts).Error
	require.NoError(t, err)
	packJob.Source = &source
	err = db.Preload("FileRanges.File").Find(&packJob).Error
	require.NoError(t, err)
	err = thread.pack(context.TODO(), packJob)
	require.NoError(t, err)
	var cars []model.Car
	err = db.Find(&cars).Error
	require.NoError(t, err)
	require.Greater(t, cars[0].PieceSize, int64(0))
	require.Equal(t, *cars[0].PackJobID, packJob.ID)
	require.NotEmpty(t, cars[0].StoragePath)
	var packJobs []model.PackJob
	err = db.Find(&packJobs).Error
	require.NoError(t, err)
	var files []model.File
	err = db.Find(&files).Error
	require.NoError(t, err)
	for _, file := range files {
		require.NotEqual(t, file.CID.String(), "")
	}
	var fileRanges []model.FileRange
	err = db.Find(&fileRanges).Error
	require.NoError(t, err)
	for _, fileRange := range fileRanges {
		require.NotEqual(t, fileRange.CID.String(), "")
	}
	var dirs []model.Directory
	err = db.Find(&dirs).Error
	require.NoError(t, err)
	for _, dir := range dirs {
		require.NotEqual(t, dir.CID.String(), "")
	}
	var carBlocks []model.CarBlock
	err = db.Find(&carBlocks).Error
	require.NoError(t, err)
	require.Len(t, carBlocks, 3)

	_, err = os.Stat(temp + "/test.txt")
	require.True(t, os.IsNotExist(err))
	_, err = os.Stat(temp + "/test2.txt")
	require.True(t, os.IsNotExist(err))
}

func TestDatasetWorkerThread_scan(t *testing.T) {
	ctx := context.Background()
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()
	thread := Thread{
		id:                        uuid.New(),
		dbNoContext:               db,
		logger:                    logger.With("key", "value"),
		datasourceHandlerResolver: storagesystem.DefaultHandlerResolver{},
		config: Config{
			Concurrency:    1,
			ExitOnComplete: true,
			EnableScan:     true,
			EnablePack:     true,
			EnableDag:      true,
			ExitOnError:    true,
		},
	}
	dataset := model.Preparation{
		Name:    "test",
		MaxSize: 1024,
	}
	err = db.Create(&dataset).Error
	require.NoError(t, err)
	cmd, _ := os.Getwd()
	source := model.Source{
		DatasetID:     dataset.ID,
		ScanningState: model.Ready,
		Type:          "local",
		Path:          cmd,
	}
	err = db.Create(&source).Error
	require.NoError(t, err)
	source.Dataset = &dataset
	root := model.Directory{
		SourceID: source.ID,
	}
	err = db.Create(&root).Error
	require.NoError(t, err)
	err = thread.scan(ctx, source, true)
	require.NoError(t, err)
	var dirs []model.Directory
	err = db.Find(&dirs).Error
	require.NoError(t, err)
	require.Greater(t, len(dirs), 0)
	var files []model.File
	err = db.Find(&files).Error
	require.NoError(t, err)
	require.Greater(t, len(files), 0)
	var fileRanges []model.FileRange
	err = db.Find(&fileRanges).Error
	require.NoError(t, err)
	require.Greater(t, len(fileRanges), 0)
}

func TestDatasetWorkerThread_findPackWork(t *testing.T) {
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()
	thread := Thread{
		id:                        uuid.New(),
		dbNoContext:               db,
		logger:                    logger.With("key", "value"),
		datasourceHandlerResolver: storagesystem.DefaultHandlerResolver{},
		config: Config{
			Concurrency:    1,
			ExitOnComplete: true,
			EnableScan:     true,
			EnablePack:     true,
			EnableDag:      true,
			ExitOnError:    true,
		},
	}
	worker := model.Worker{
		ID:            thread.id.String(),
		LastHeartbeat: time.Now(),
	}
	err = db.Create(&worker).Error
	require.NoError(t, err)
	dataset := model.Preparation{
		Name: "test",
	}
	err = db.Create(&dataset).Error
	require.NoError(t, err)
	source := model.Source{
		DatasetID:     dataset.ID,
		ScanningState: model.Complete,
	}
	err = db.Create(&source).Error
	require.NoError(t, err)
	root := model.Directory{
		SourceID: source.ID,
	}
	err = db.Create(&root).Error
	require.NoError(t, err)
	files := []model.File{
		{
			SourceID: source.ID,
		},
		{
			SourceID: source.ID,
		},
	}
	err = db.Create(&files).Error
	require.NoError(t, err)
	fileRanges := []model.FileRange{
		{
			FileID: files[0].ID,
		},
		{
			FileID: files[1].ID,
		},
		{
			FileID: files[1].ID,
		},
	}
	err = db.Create(&fileRanges).Error
	require.NoError(t, err)
	packJobs := map[*model.PackJob]bool{
		{
			PackingState: model.Ready,
		}: true,
		{
			PackingState: model.Processing,
		}: true,
		{
			PackingState: model.Error,
		}: false,
		{
			PackingState:    model.Processing,
			PackingWorkerID: &worker.ID,
		}: false,
	}
	for packJob, shouldBeFound := range packJobs {
		err := db.Where("1 = 1").Delete(&model.PackJob{}).Error
		require.NoError(t, err)
		packJob.SourceID = source.ID
		err = db.Create(packJob).Error
		for _, fileRanges := range fileRanges {
			fileRanges.PackJobID = &packJob.ID
			err = db.Save(&fileRanges).Error
			require.NoError(t, err)
		}
		require.NoError(t, err)
		ck, err := thread.findPackWork(context.Background())
		require.NoError(t, err)
		if shouldBeFound {
			require.NotNil(t, ck)
			require.NotNil(t, ck.Source)
			require.NotNil(t, ck.Source.Dataset)
			require.NotNil(t, ck.FileRanges)
			require.NotNil(t, ck.FileRanges[0].File)
		} else {
			require.Nil(t, ck)
		}
		err = db.Where("1 = 1").Delete(&model.PackJob{}).Error
		require.NoError(t, err)
	}
}

func TestDatasetWorkerThread_findScanWork(t *testing.T) {
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()
	thread := Thread{
		id:                        uuid.New(),
		dbNoContext:               db,
		logger:                    logger.With("key", "value"),
		datasourceHandlerResolver: storagesystem.DefaultHandlerResolver{},
		config: Config{
			Concurrency:    1,
			ExitOnComplete: true,
			EnableScan:     true,
			EnablePack:     true,
			EnableDag:      true,
			ExitOnError:    true,
		},
	}
	worker := model.Worker{
		ID:            thread.id.String(),
		LastHeartbeat: time.Now(),
	}
	err = db.Create(&worker).Error
	require.NoError(t, err)
	dataset := model.Preparation{
		Name: "test",
	}
	err = db.Create(&dataset).Error
	require.NoError(t, err)
	sources := map[*model.Source]bool{
		// data source that is ready to be scanned
		{
			ScanningState: model.Ready,
		}: true,
		// data source that is being scanned but does not have a worker id
		{
			ScanningState:    model.Processing,
			ScanningWorkerID: nil,
		}: true,
		// data source that has completed scanning and should be scanned again
		{
			ScanningState:        model.Complete,
			LastScannedTimestamp: 0,
			ScanIntervalSeconds:  100,
		}: true,
		// data source that has completed scanning and should not be scanned again
		{
			ScanningState:        model.Complete,
			LastScannedTimestamp: 0,
			ScanIntervalSeconds:  0,
		}: false,
		// data source that has completed scanning and should not be scanned again
		{
			ScanningState:        model.Complete,
			LastScannedTimestamp: time.Now().Unix(),
			ScanIntervalSeconds:  100,
		}: false,
		// data source that has errored
		{
			ScanningState: model.Error,
		}: false,
		// data source that has worker working on it
		{
			ScanningState:    model.Processing,
			ScanningWorkerID: &worker.ID,
		}: false,
	}
	for source, shouldBeFound := range sources {
		err := db.Where("1 = 1").Delete(&model.Source{}).Error
		require.NoError(t, err)
		source.DatasetID = dataset.ID
		err = db.Create(source).Error
		require.NoError(t, err)
		root := model.Directory{
			SourceID: source.ID,
		}
		err = db.Create(&root).Error
		require.NoError(t, err)
		src, err := thread.findScanWork(context.Background())
		require.NoError(t, err)
		if shouldBeFound {
			require.NotNil(t, src)
			require.NotNil(t, src.Dataset)
			require.NotNil(t, src.RootDirectory)
		} else {
			require.Nil(t, src)
		}
	}
}
