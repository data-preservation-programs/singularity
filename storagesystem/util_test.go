package storagesystem

import (
	"context"
	"io"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/rclone/rclone/backend/s3"
	"github.com/rclone/rclone/fs"
	"github.com/rclone/rclone/fs/hash"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestIsSameEntry(t *testing.T) {
	ctx := context.Background()
	mockObject := new(MockObject)
	mockObject.On("Size").Return(int64(5))
	mockObject.On("Fs").Return(&s3.Fs{})
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
