package storagesystem

import (
	"context"
	"io"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/rclone/rclone/backend/s3"
	"github.com/rclone/rclone/fs"
	"github.com/rclone/rclone/fs/config/configmap"
	"github.com/rclone/rclone/fs/hash"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestIsSameEntry(t *testing.T) {
	ctx := context.Background()
	mockObject := new(MockObject)
	mockObject.On("Size").Return(int64(5))
	s3fs, err := s3.NewFs(ctx, "s3", "commoncrawl", configmap.Simple{"chunk_size": "5Mi"})
	require.NoError(t, err)
	mockObject.On("Fs").Return(s3fs)
	mockObject.On("Hash", mock.Anything, mock.Anything).Return("hash", nil)
	tm := time.Now()
	mockObject.On("ModTime", mock.Anything).Return(tm)
	t.Run("size mismatch", func(t *testing.T) {
		same, detail := IsSameEntry(ctx, model.File{
			Size: 4,
		}, mockObject)
		require.False(t, same)
		require.Contains(t, detail, "size mismatch")
	})
	t.Run("size unknown", func(t *testing.T) {
		same, _ := IsSameEntry(ctx, model.File{
			Size:             -1,
			Hash:             "hash",
			LastModifiedNano: tm.UnixNano(),
		}, mockObject)
		require.True(t, same)
	})
	t.Run("hash mismatch", func(t *testing.T) {
		same, detail := IsSameEntry(ctx, model.File{
			Size: 5,
			Hash: "hash2",
		}, mockObject)
		require.False(t, same)
		require.Contains(t, detail, "hash mismatch")
	})
	t.Run("last modified mismatch", func(t *testing.T) {
		same, detail := IsSameEntry(ctx, model.File{
			Size:             5,
			Hash:             "hash",
			LastModifiedNano: 100,
		}, mockObject)
		require.False(t, same)
		require.Contains(t, detail, "last modified mismatch")
	})
	t.Run("all match", func(t *testing.T) {
		same, _ := IsSameEntry(ctx, model.File{
			Size:             5,
			Hash:             "hash",
			LastModifiedNano: tm.UnixNano(),
		}, mockObject)
		require.True(t, same)
	})
	t.Run("all match, ignoring empty hash", func(t *testing.T) {
		same, _ := IsSameEntry(ctx, model.File{
			Size:             5,
			LastModifiedNano: tm.UnixNano(),
		}, mockObject)
		require.True(t, same)
	})
}

type MockObject struct {
	mock.Mock
}

func (m *MockObject) Remote() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockObject) ModTime(ctx context.Context) time.Time {
	args := m.Called(ctx)
	return args.Get(0).(time.Time)
}

func (m *MockObject) Size() int64 {
	args := m.Called()
	return args.Get(0).(int64)
}

func (m *MockObject) Fs() fs.Info {
	args := m.Called()
	return args.Get(0).(fs.Info)
}

func (m *MockObject) Hash(ctx context.Context, ty hash.Type) (string, error) {
	args := m.Called(ctx, ty)
	return args.String(0), args.Error(1)
}

func (m *MockObject) Storable() bool {
	args := m.Called()
	return args.Bool(0)
}

func (m *MockObject) SetModTime(ctx context.Context, t time.Time) error {
	args := m.Called(ctx, t)
	return args.Error(0)
}

func (m *MockObject) Open(ctx context.Context, options ...fs.OpenOption) (io.ReadCloser, error) {
	args := m.Called(ctx, options)
	return args.Get(0).(io.ReadCloser), args.Error(1)
}

func (m *MockObject) Update(ctx context.Context, in io.Reader, src fs.ObjectInfo, options ...fs.OpenOption) error {
	args := m.Called(ctx, in, src, options)
	return args.Error(0)
}

func (m *MockObject) Remove(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func TestGetRandomOutputWriter(t *testing.T) {
	ctx := context.Background()
	s1 := model.Storage{
		ID:   1,
		Type: "local",
		Path: t.TempDir(),
	}
	s2 := model.Storage{
		ID:   1,
		Type: "local",
		Path: t.TempDir(),
	}
	s3 := model.Storage{
		ID:     3,
		Type:   "s3",
		Path:   "commoncrawl",
		Config: map[string]string{"chunk_size": "5Mi"},
	}
	t.Run("no storages", func(t *testing.T) {
		id, writer, err := GetRandomOutputWriter(ctx, []model.Storage{})
		require.NoError(t, err)
		require.Nil(t, id)
		require.Nil(t, writer)
	})
	t.Run("single local storage", func(t *testing.T) {
		id, writer, err := GetRandomOutputWriter(ctx, []model.Storage{s1})
		require.NoError(t, err)
		require.EqualValues(t, 1, *id)
		require.NotNil(t, writer)
	})
	t.Run("single s3 storage", func(t *testing.T) {
		id, writer, err := GetRandomOutputWriter(ctx, []model.Storage{s3})
		require.NoError(t, err)
		require.EqualValues(t, 3, *id)
		require.NotNil(t, writer)
	})
	t.Run("all storage", func(t *testing.T) {
		_, writer, err := GetRandomOutputWriter(ctx, []model.Storage{s1, s2, s3})
		require.NoError(t, err)
		require.NotNil(t, writer)
	})
	t.Run("space warning", func(t *testing.T) {
		current := freeSpaceWarningThreshold
		freeSpaceWarningThreshold = 1.0
		defer func() {
			freeSpaceWarningThreshold = current
		}()

		id, writer, err := GetRandomOutputWriter(ctx, []model.Storage{s1})
		require.NoError(t, err)
		require.EqualValues(t, 1, *id)
		require.NotNil(t, writer)
	})
	t.Run("space error", func(t *testing.T) {
		current := freeSpaceErrorThreshold
		freeSpaceErrorThreshold = 1.0
		defer func() {
			freeSpaceErrorThreshold = current
		}()

		_, _, err := GetRandomOutputWriter(ctx, []model.Storage{s1})
		require.ErrorIs(t, err, ErrStorageNotAvailable)
	})
}
