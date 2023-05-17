package service

import (
	"bytes"
	"context"
	"fmt"
	"github.com/data-preservation-programs/go-singularity/database"
	"github.com/data-preservation-programs/go-singularity/datasource"
	"github.com/data-preservation-programs/go-singularity/handler/dataset"
	"github.com/data-preservation-programs/go-singularity/model"
	"github.com/google/uuid"
	"github.com/ipfs/go-log/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"io"
	"math/rand"
	"os"
	"syscall"
	"testing"
	"time"
)

type MockScanner struct {
	mock.Mock
}

type MockResolver struct {
	mock.Mock
}

func (m *MockScanner) Scan(ctx context.Context, path string, last string) <-chan datasource.Entry {
	args := m.Called(ctx, path, last)
	return args.Get(0).(chan datasource.Entry)
}

func (m *MockResolver) GetHandler(source model.Source) (datasource.Handler, error) {
	args := m.Called(source)
	return args.Get(0).(datasource.Handler), args.Error(1)
}

func TestScan_EmptyEntries(t *testing.T) {
	assert := assert.New(t)
	db := database.OpenInMemory()
	defer database.DropAll(db)
	_, src, _ := createTestSource(db, 0)
	mockScanner := &MockScanner{}
	mockResolver := &MockResolver{}
	mockResolver.On("GetHandlerBySource", mock.Anything).Return(mockScanner, nil)

	entryChan := make(chan datasource.Entry)
	mockScanner.On("Scan", mock.Anything, mock.Anything, mock.Anything).
		Return(entryChan, nil)
	go func() {
		close(entryChan)
	}()

	thread := createWorkerThread(db)
	thread.datasourceHandlerResolver = mockResolver
	err := thread.scan(context.Background(), src)
	assert.Nil(err)
	var count int64
	err = db.Model(&model.Chunk{}).Count(&count).Error
	assert.Nil(err)
	assert.Equal(int64(0), count)
}

func TestScan_NewEntries(t *testing.T) {
	assert := assert.New(t)
	db := database.OpenInMemory()
	defer database.DropAll(db)
	_, src, _ := createTestSource(db, 0)
	mockScanner := &MockScanner{}
	mockResolver := &MockResolver{}
	mockResolver.On("GetHandlerBySource", mock.Anything).Return(mockScanner, nil)

	entryChan := make(chan datasource.Entry)
	mockScanner.On("Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(entryChan, nil)
	go func() {
		t := time.Now().UTC()
		entryChan <- datasource.Entry{
			Type:         model.File,
			Path:         "/tmp/1.img",
			Size:         1_000_000,
			LastModified: &t,
		}
		entryChan <- datasource.Entry{
			Type:         model.File,
			Path:         "/tmp/2.img",
			Size:         2_000_000,
			LastModified: &t,
		}
		entryChan <- datasource.Entry{
			Type:         model.File,
			Path:         "/tmp/3.img",
			Size:         4_000_000,
			LastModified: &t,
		}
		close(entryChan)
	}()

	thread := createWorkerThread(db)
	thread.datasourceHandlerResolver = mockResolver
	err := thread.scan(context.Background(), src)
	assert.Nil(err)
	var items []model.Item
	err = db.Order("chunk_id asc").Find(&items).Error
	assert.Nil(err)
	assert.Equal(6, len(items))
}

func TestChunk_SplitBigItem(t *testing.T) {
	assert := assert.New(t)
	db := database.OpenInMemory()
	defer database.DropAll(db)
	ds, src, _ := createTestSource(db, 0)
	items := []model.Item{{
		SourceID: src.ID,
		Type:     model.File,
		Path:     "/tmp/relative/path/file1.txt",
		Size:     1_000_000,
		Offset:   0,
		Length:   1_000_000,
	}, {
		SourceID: src.ID,
		Type:     model.File,
		Path:     "/tmp/relative/path/file2.txt",
		Size:     2_000_000,
		Offset:   0,
		Length:   2_000_000,
	}, {
		SourceID: src.ID,
		Type:     model.File,
		Path:     "/tmp/relative/path/file3.txt",
		Size:     4_000_000,
		Offset:   0,
		Length:   4_000_000,
	}, {
		SourceID: src.ID,
		Type:     model.File,
		Path:     "/tmp/relative/path/file4.txt",
		Size:     8_000_000,
		Offset:   0,
		Length:   8_000_000,
	}}
	assert.Nil(db.Create(&items).Error)
	thread := createWorkerThread(db)
	remaining := newRemain()
	for _, item := range items {
		remaining.add(item)
	}
	for len(remaining.items) > 0 {
		var err error
		err = thread.chunkOnce(src, ds, remaining)
		assert.Nil(err)
	}
	assert.Empty(remaining.items)
	err := db.Order("chunk_id asc").Find(&items).Error
	assert.Nil(err)
	assert.Equal(11, len(items))
	assert.Equal("1 /tmp/relative/path/file1.txt 1000000 (0-1000000)", ToString(items[0]))
	assert.Equal("2 /tmp/relative/path/file2.txt 2000000 (0-1999620)", ToString(items[1]))
	assert.Equal("3 /tmp/relative/path/file2.txt 2000000 (1999620-2000000)", ToString(items[2]))
	assert.Equal("3 /tmp/relative/path/file3.txt 4000000 (0-1999195)", ToString(items[3]))
	assert.Equal("4 /tmp/relative/path/file3.txt 4000000 (1999195-3998815)", ToString(items[4]))
	assert.Equal("5 /tmp/relative/path/file3.txt 4000000 (3998815-4000000)", ToString(items[5]))
	assert.Equal("5 /tmp/relative/path/file4.txt 8000000 (0-1998390)", ToString(items[6]))
	assert.Equal("6 /tmp/relative/path/file4.txt 8000000 (1998390-3998010)", ToString(items[7]))
	assert.Equal("7 /tmp/relative/path/file4.txt 8000000 (3998010-5997630)", ToString(items[8]))
	assert.Equal("8 /tmp/relative/path/file4.txt 8000000 (5997630-7997250)", ToString(items[9]))
	assert.Equal("9 /tmp/relative/path/file4.txt 8000000 (7997250-8000000)", ToString(items[10]))
}

func ToString(i model.Item) string {
	return fmt.Sprintf("%d %s %d (%d-%d)", *i.ChunkID, i.Path, i.Size, i.Offset, i.Length+i.Offset)
}

func TestChunk(t *testing.T) {
	assert := assert.New(t)
	db := database.OpenInMemory()
	defer database.DropAll(db)
	ds, src, _ := createTestSource(db, 0)
	item := model.Item{
		SourceID: src.ID,
		Type:     model.File,
		Path:     "/tmp/relative/path/file.txt",
		Size:     1_000_000,
		Offset:   0,
		Length:   1_000_000,
	}
	assert.Nil(db.Create(&item).Error)
	thread := createWorkerThread(db)
	remaining := newRemain()
	remaining.add(item)
	err := thread.chunkOnce(src, ds, remaining)
	assert.Nil(err)
	assert.Empty(remaining.items)
	var items []model.Item
	err = db.Find(&items).Error
	assert.Nil(err)
	assert.Equal(1, len(items))
	assert.Equal("1 /tmp/relative/path/file.txt 1000000 (0-1000000)", ToString(items[0]))
}

func TestEnsureParentDirectories(t *testing.T) {
	assert := assert.New(t)
	db := database.OpenInMemory()
	defer database.DropAll(db)
	_, src, _ := createTestSource(db, 0)
	chunk := model.Chunk{
		SourceID:     src.ID,
		PackingState: model.Ready,
	}
	assert.Nil(db.Create(&chunk).Error)
	var root model.Directory
	assert.Nil(db.Find(&root, src.RootDirectoryID).Error)
	item := model.Item{
		ChunkID:  &chunk.ID,
		SourceID: src.ID,
		Type:     model.File,
		Path:     "/tmp/relative/path/file.txt",
		Size:     1000,
		Offset:   100,
		Length:   100,
	}
	thread := createWorkerThread(db)
	err := db.Create(&item).Error
	assert.Nil(err)
	err = thread.ensureParentDirectories(&item, root)
	assert.Nil(err)
	var dirs []model.Directory
	assert.Nil(db.Find(&dirs).Error)
	assert.Equal(3, len(dirs))
	assert.Equal("/tmp", dirs[0].Name)
	assert.Equal("relative", dirs[1].Name)
	assert.Equal("path", dirs[2].Name)
	assert.Equal(*dirs[2].ParentID, dirs[1].ID)
	assert.Equal(*dirs[1].ParentID, dirs[0].ID)
}

func createWorkerThread(db *gorm.DB) DatasetWorkerThread {
	id := uuid.New()
	return DatasetWorkerThread{
		id:             id,
		db:             db,
		logger:         log.Logger("test").With("workerID", id.String()),
		directoryCache: map[string]model.Directory{},
	}
}

func createTestSource(db *gorm.DB, scanInterval time.Duration) (model.Dataset, model.Source, model.Directory) {
	ds, err := dataset.CreateHandler(db, dataset.CreateRequest{
		Name:       "test",
		MinSizeStr: "1MB",
		MaxSizeStr: "2MB",
	})
	if err != nil {
		panic(err)
	}
	src, err := dataset.AddSourceHandler(db, dataset.AddSourceRequest{
		DatasetName:  "test",
		SourcePath:   "/tmp",
		ScanInterval: scanInterval,
	})
	if err != nil {
		panic(err)
	}
	var dir model.Directory
	err = db.First(&dir, src.RootDirectoryID).Error
	if err != nil {
		panic(err)
	}

	return *ds, *src, dir
}

type MockStreamer struct {
	mock.Mock
}

func (m *MockStreamer) Open(ctx context.Context, path string, offset uint64, length uint64) (io.ReadCloser, error) {
	args := m.Called(path, offset, length)
	return args.Get(0).(io.ReadCloser), args.Error(1)
}

type readCloser struct {
	io.Reader
}

func (rc *readCloser) Close() error {
	return nil
}

func newRandomReadCloser(size int, seed int64) io.ReadCloser {
	rand.Seed(seed)
	randomBytes := make([]byte, size)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}

	reader := bytes.NewReader(randomBytes)
	return &readCloser{Reader: reader}
}

func TestDatasetWorkerThread_Pack(t *testing.T) {
	assert := assert.New(t)
	db := database.OpenInMemory()
	defer database.DropAll(db)
	thread := createWorkerThread(db)
	ds, src, _ := createTestSource(db, 0)
	chunk := model.Chunk{
		SourceID:     src.ID,
		PackingState: model.Ready,
	}
	err := db.Create(&chunk).Error
	assert.Nil(err)
	item := model.Item{
		ChunkID:  &chunk.ID,
		SourceID: src.ID,
		Type:     model.File,
		Path:     "/mnt/test.bin",
		Size:     10_000_000,
		Offset:   0,
		Length:   10_000_000,
	}
	err = db.Create(&item).Error
	assert.Nil(err)
	streamer := &MockStreamer{}
	streamer.On("Read", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(newRandomReadCloser(10_000_000, 0), nil)

	mockResolver := &MockResolver{}
	mockResolver.On("GetHandlerBySource", mock.Anything).Return(streamer, nil)
	thread.datasourceHandlerResolver = mockResolver
	err = thread.pack(context.Background(), chunk.ID, src, []model.Item{item}, ds.OutputDirs, ds.PieceSize)
	assert.Nil(err)
	var cars []model.Car
	assert.Nil(db.Find(&cars).Error)
	assert.Equal(1, len(cars))
	assert.Equal("baga6ea4seaqlrgzusn5d55czmd4yx6qezpws6h2cb7ji4zlfgsjpt54kgij3wmi", cars[0].PieceCID)
	var rawBlocks []model.RawBlock
	assert.Nil(db.Find(&rawBlocks).Error)
	assert.Equal(1, len(rawBlocks))
	var carBlocks []model.CarBlock
	assert.Nil(db.Find(&carBlocks).Error)
	assert.Equal(11, len(carBlocks))
	var itemBlocks []model.ItemBlock
	assert.Nil(db.Find(&itemBlocks).Error)
	assert.Equal(10, len(itemBlocks))
}

func TestDatasetWorkerThread_FindPackWork_NeedPack(t *testing.T) {
	assert := assert.New(t)
	db := database.OpenInMemory()
	defer database.DropAll(db)
	thread := createWorkerThread(db)
	chunk, err := thread.findPackWork()
	assert.Nil(err)
	assert.Nil(chunk)

	_, src, _ := createTestSource(db, 0)
	db.Model(&src).Update("scanning_state", model.Complete)
	chunk = &model.Chunk{
		SourceID:     src.ID,
		PackingState: model.Ready,
	}
	assert.NoError(db.Create(chunk).Error)
	assert.NoError(db.Create(&model.Item{
		ChunkID:  &chunk.ID,
		SourceID: src.ID,
	}).Error)

	chunk, err = thread.findPackWork()
	assert.Nil(err)
	assert.NotNil(chunk)
	assert.NotNil(chunk.Source)
	assert.NotNil(chunk.Source.Dataset)
	assert.NotEmpty(chunk.Items)
}

func TestDatasetWorkerThread_FindScanWork_NeedScan(t *testing.T) {
	assert := assert.New(t)
	db := database.OpenInMemory()
	defer database.DropAll(db)
	thread := createWorkerThread(db)

	source, err := thread.findScanWork()
	assert.Nil(err)
	assert.Nil(source)

	_, _, _ = createTestSource(db, 0)
	source, err = thread.findScanWork()
	assert.Nil(err)
	assert.NotNil(source)
	assert.Equal(model.Processing, source.ScanningState)
	assert.Equal(thread.id.String(), *source.ScanningWorkerID)
}

func TestDatasetWorkerThread_FindScanWork_NeedRescan(t *testing.T) {
	assert := assert.New(t)
	db := database.OpenInMemory()
	defer database.DropAll(db)
	id := uuid.New()
	thread := DatasetWorkerThread{
		id:             id,
		db:             db,
		logger:         log.Logger("test").With("workerID", id.String()),
		directoryCache: map[string]model.Directory{},
	}

	source, err := thread.findScanWork()
	assert.Nil(err)
	assert.Nil(source)

	_, src, _ := createTestSource(db, time.Second)
	db.Model(src).Updates(map[string]interface{}{
		"scanning_state":         model.Complete,
		"scanning_worker_id":     nil,
		"last_scanned_timestamp": time.Now().UTC().Unix(),
	})

	source, err = thread.findScanWork()
	assert.Nil(err)
	assert.Nil(source)

	time.Sleep(2 * time.Second)
	source, err = thread.findScanWork()
	assert.Nil(err)
	assert.NotNil(source)
	assert.Equal(model.Processing, source.ScanningState)
	assert.Equal(id.String(), *source.ScanningWorkerID)
}

func TestDatasetWorker_ReceiveInterruption(t *testing.T) {
	assert := assert.New(t)
	db := database.OpenInMemory()
	defer database.DropAll(db)
	w := NewDatasetWorker(db, 1)
	ctx := context.TODO()
	errChan := make(chan error)
	go func() {
		errChan <- w.Run(ctx)
	}()

	time.Sleep(time.Second)
	var wCount int64
	err := db.Model(&model.Worker{}).Count(&wCount).Error
	assert.Nil(err)
	assert.Equal(int64(1), wCount)

	pid := os.Getpid()
	syscall.Kill(pid, syscall.SIGTERM)
	err = <-errChan
	assert.ErrorContains(err, "received signal")

	err = db.Model(&model.Worker{}).Count(&wCount).Error
	assert.Nil(err)
	assert.Equal(int64(0), wCount)
}
