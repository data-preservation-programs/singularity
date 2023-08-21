package util

import (
	"bytes"
	"context"
	"io"
	"testing"
	"time"

	"github.com/ipfs/boxo/util"
	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"
	format "github.com/ipfs/go-ipld-format"
	"github.com/rclone/rclone/fs"
	"github.com/rclone/rclone/fs/hash"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreateParentNode(t *testing.T) {
	links := []format.Link{
		{
			Name: "",
			Size: 5,
			Cid:  cid.NewCidV1(cid.Raw, util.Hash([]byte("hello"))),
		},
		{
			Name: "",
			Size: 5,
			Cid:  cid.NewCidV1(cid.Raw, util.Hash([]byte("world"))),
		},
	}

	node, size, err := createParentNode(links)
	require.NoError(t, err)
	require.Equal(t, uint64(10), size)
	require.NotNil(t, node)
	size, err = node.Size()
	require.NoError(t, err)
	require.Equal(t, uint64(108), size)
	require.Equal(t, "bafybeiejlvvmfokp5c6q2eqgbfjeaokz3nqho5c7yy3ov527vsatgsqfma", node.String())
}

func TestMin(t *testing.T) {
	require.Equal(t, 1, Min(1, 2))
	require.Equal(t, 1, Min(2, 1))
	require.Equal(t, 1, Min(1, 1))
}

func TestAssembleFileFromLinks_SingleLink(t *testing.T) {
	links := []format.Link{
		{
			Name: "",
			Size: 5,
			Cid:  cid.NewCidV1(cid.Raw, util.Hash([]byte("hello"))),
		},
	}
	_, _, err := AssembleFileFromLinks(links)
	require.ErrorIs(t, err, errLinkLessThanTwo)
}

func TestAssembleFileFromLinks(t *testing.T) {
	links := []format.Link{
		{
			Name: "",
			Size: 5,
			Cid:  cid.NewCidV1(cid.Raw, util.Hash([]byte("hello"))),
		},
		{
			Name: "",
			Size: 5,
			Cid:  cid.NewCidV1(cid.Raw, util.Hash([]byte("world"))),
		},
	}

	blocks, node, err := AssembleFileFromLinks(links)
	require.NoError(t, err)
	require.NotNil(t, node)
	size, err := node.Size()
	require.NoError(t, err)
	require.Equal(t, uint64(108), size)
	require.Equal(t, "bafybeiejlvvmfokp5c6q2eqgbfjeaokz3nqho5c7yy3ov527vsatgsqfma", node.String())
	require.Len(t, blocks, 1)
	require.Equal(t, "bafybeiejlvvmfokp5c6q2eqgbfjeaokz3nqho5c7yy3ov527vsatgsqfma", blocks[0].Cid().String())
}

func TestAssembleFileFromLinks_LargeFile(t *testing.T) {
	links := []format.Link{}
	for i := 0; i < 2_000; i++ {
		links = append(links, format.Link{
			Name: "",
			Size: 5,
			Cid:  cid.NewCidV1(cid.Raw, util.Hash([]byte("hello"))),
		})
	}

	blocks, node, err := AssembleFileFromLinks(links)
	require.NoError(t, err)
	require.NotNil(t, node)
	size, err := node.Size()
	require.NoError(t, err)
	require.EqualValues(t, 10103, size)
	require.Equal(t, "bafybeiamlrjlfotypfc5hse7uenmgmav7yq5vcb75yd7bxh2aaavxbi4ou", node.String())
	require.Len(t, blocks, 3)
	require.Equal(t, "bafybeidci5xcdskgvefhv3x6kp6zpyxkbiqshqjbpsdgmk2k3mexzaggwi", blocks[0].Cid().String())
}

func TestWriteCarHeader(t *testing.T) {
	buf := new(bytes.Buffer)
	header, err := WriteCarHeader(buf, cid.NewCidV1(cid.Raw, util.Hash([]byte("hello"))))
	require.NoError(t, err)
	require.Len(t, header, 59)
	require.Equal(t, buf.Bytes(), header)
}

func TestWriteCarBlock(t *testing.T) {
	buf := new(bytes.Buffer)
	block := blocks.NewBlock([]byte("hello"))
	n, err := WriteCarBlock(buf, block)
	require.NoError(t, err)
	require.Equal(t, int64(40), n)
	require.Equal(t, 40, buf.Len())
}

type MockReadHandler struct {
	mock.Mock
}

func (m *MockReadHandler) Name() string {
	return "mock"
}

func (m *MockReadHandler) Read(ctx context.Context, path string, offset int64, length int64) (io.ReadCloser, fs.Object, error) {
	args := m.Called(ctx, path, offset, length)
	if args.Get(1) == nil {
		return args.Get(0).(io.ReadCloser), nil, args.Error(2)
	}
	return args.Get(0).(io.ReadCloser), args.Get(1).(fs.Object), args.Error(2)
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
