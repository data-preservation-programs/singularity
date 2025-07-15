package scan

import (
	"context"
	"io"
	"path/filepath"
	"testing"
	"time"
	"unsafe"

	"github.com/data-preservation-programs/singularity/storagesystem"
	"github.com/rclone/rclone/fs"
	"github.com/rclone/rclone/fs/hash"
	"github.com/stretchr/testify/require"
)

// A complete fs.Fs implementation that also supports ListUpstreams
type mockUnionFs struct {
	upstreams []string
	root      string
}

func (m *mockUnionFs) ListUpstreams() []string {
	return m.upstreams
}

func (m *mockUnionFs) Root() string {
	return m.root
}

func (m *mockUnionFs) String() string {
	return "mock:" + m.root
}

func (m *mockUnionFs) Name() string {
	return "mock"
}

func (m *mockUnionFs) Features() *fs.Features {
	return &fs.Features{}
}

func (m *mockUnionFs) Precision() time.Duration {
	return time.Second
}

func (m *mockUnionFs) NewObject(context.Context, string) (fs.Object, error) {
	return nil, fs.ErrorObjectNotFound
}

func (m *mockUnionFs) List(context.Context, string) (fs.DirEntries, error) {
	return nil, fs.ErrorDirNotFound
}

func (m *mockUnionFs) Put(ctx context.Context, in io.Reader, src fs.ObjectInfo, options ...fs.OpenOption) (fs.Object, error) {
	return nil, fs.ErrorObjectNotFound
}

func (m *mockUnionFs) Mkdir(context.Context, string) error {
	return fs.ErrorDirNotFound
}

func (m *mockUnionFs) Rmdir(context.Context, string) error {
	return fs.ErrorDirNotFound
}

func (m *mockUnionFs) Purge(context.Context, string) error {
	return fs.ErrorDirNotFound
}

func (m *mockUnionFs) Copy(context.Context, fs.Object, string) (fs.Object, error) {
	return nil, fs.ErrorObjectNotFound
}

func (m *mockUnionFs) Move(context.Context, fs.Object, string) (fs.Object, error) {
	return nil, fs.ErrorObjectNotFound
}

func (m *mockUnionFs) DirMove(context.Context, fs.Fs, string, string) error {
	return fs.ErrorDirNotFound
}

func (m *mockUnionFs) Hashes() hash.Set {
	return hash.NewHashSet()
}

// Mock RCloneHandler that wraps a mockUnionFs
type mockHandler struct {
	name string
	fs   fs.Fs
}

func (m *mockHandler) Name() string {
	return m.name
}

func (m *mockHandler) Fs() fs.Fs {
	return m.fs
}

func TestGetUnionUpstreams(t *testing.T) {
	ctx := context.Background()
	mockFs := &mockUnionFs{
		upstreams: []string{"folder1", "folder2"},
		root:      "/mock/root",
	}
	handler := &mockHandler{name: "mock", fs: mockFs}

	upstreams, err := GetUnionUpstreams(ctx, (*storagesystem.RCloneHandler)(unsafe.Pointer(handler)))
	require.NoError(t, err)
	require.Equal(t, []string{"folder1", "folder2"}, upstreams)
}

func TestGetUpstreamPaths(t *testing.T) {
	mockFs := &mockUnionFs{
		upstreams: []string{"folder1", "folder2"},
		root:      "/mock/root",
	}
	handler := &mockHandler{name: "mock", fs: mockFs}

	upstreams := []string{"folder1", "folder2"}
	paths := GetUpstreamPaths((*storagesystem.RCloneHandler)(unsafe.Pointer(handler)), upstreams)
	expected := map[string]string{
		"folder1": filepath.Join("/mock/root", "folder1"),
		"folder2": filepath.Join("/mock/root", "folder2"),
	}
	require.Equal(t, expected, paths)
}

func TestGetUpstreamPathFromFilePath(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
		want     string
	}{
		{
			name:     "file in root",
			filePath: "file.txt",
			want:     "",
		},
		{
			name:     "file in folder",
			filePath: "folder1/file.txt",
			want:     "folder1",
		},
		{
			name:     "file in nested folder",
			filePath: "folder1/subfolder/file.txt",
			want:     "folder1",
		},
		{
			name:     "hidden folder",
			filePath: ".folder/file.txt",
			want:     ".folder",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getUpstreamPath(tt.filePath)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestUpstreamPathHandling(t *testing.T) {
	mockFs := &mockUnionFs{
		upstreams: []string{"/abs/path1", "rel/path2", "folder3"},
		root:      "/mock/root",
	}
	handler := &mockHandler{name: "mock", fs: mockFs}

	paths := GetUpstreamPaths((*storagesystem.RCloneHandler)(unsafe.Pointer(handler)), mockFs.upstreams)

	require.Equal(t, "/abs/path1", paths["/abs/path1"], "absolute paths should be preserved")
	require.Equal(t, "/mock/root/rel/path2", paths["rel/path2"], "relative paths should be joined with root")
	require.Equal(t, "/mock/root/folder3", paths["folder3"], "simple paths should be joined with root")
}
