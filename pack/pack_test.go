package pack

import (
	"bytes"
	"context"
	"github.com/data-preservation-programs/go-singularity/datasource"
	"github.com/data-preservation-programs/go-singularity/model"
	blocks "github.com/ipfs/go-block-format"
	format "github.com/ipfs/go-ipld-format"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"math/rand"
	"os"
	"testing"
	"time"
)

type MockHandler struct {
	mock.Mock
}

func (m *MockHandler) Open(ctx context.Context, path string, offset uint64, length uint64) (io.ReadCloser, error) {
	args := m.Called(path, offset, length)
	return args.Get(0).(io.ReadCloser), args.Error(1)
}

func (m *MockHandler) Scan(ctx context.Context, path string, last string) <-chan datasource.Entry {
	args := m.Called(path, last)
	return args.Get(0).(chan datasource.Entry)
}

func (m *MockHandler) CheckItem(ctx context.Context, path string) (uint64, *time.Time, error) {
	args := m.Called(path)
	return args.Get(0).(uint64), args.Get(1).(*time.Time), args.Error(2)
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

func TestStreamItem(t *testing.T) {
	assert := assert.New(t)
	streamer := &MockHandler{}
	streamer.On("Open", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(newRandomReadCloser(1024 * 1024 * 256 + 1, 0), nil)

	ctx := context.TODO()
	item := model.Item{
		Type:         model.File,
		Path:         "/test",
		Size:         1024 * 1024 * 256 + 1,
		Offset:       0,
		Length:       1024 * 1024 * 256 + 1,
	}
	resultChan, err := streamItem(ctx, streamer, item)
	assert.NoError(err)
	results := make([]BlockResult, 0)
	for result := range resultChan {
		assert.NoError(result.Error)
		results = append(results, result)
	}

	assert.Equal(257, len(results))
	assert.EqualValues(1024 * 1024, results[1].Length)
	assert.EqualValues(1024 * 1024, results[1].Offset)
	assert.Equal("bafkreibiwr2p5y6amnlumocf7rl5tc6zctdfcjioy5eigbx3cvipwxbl4e", results[1].CID.String())
	assert.NotEmpty(results[1].Raw)
	assert.EqualValues(1, results[256].Length)
	assert.EqualValues(1024 * 1024 * 256, results[256].Offset)
	assert.Equal("bafkreicfpzefjbr6p35kamtgvv4bqixmzzu56mnfcf4gcggba5y2q7tj6i", results[256].CID.String())
	assert.NotEmpty(results[256].Raw)
}

func TestWriteCarBlock(t *testing.T) {
	assert := assert.New(t)
	block := blocks.NewBlock([]byte("test"))
	writer := bytes.NewBuffer(nil)
	written, err := writeCarBlock(writer, *block)
	assert.NoError(err)
	assert.EqualValues(39, written)
	expected := []byte{0x26, 0x12, 0x20, 0x9f, 0x86, 0xd0, 0x81, 0x88, 0x4c, 0x7d, 0x65, 0x9a, 0x2f, 0xea, 0xa0, 0xc5, 0x5a, 0xd0, 0x15, 0xa3, 0xbf, 0x4f, 0x1b, 0x2b, 0xb, 0x82, 0x2c, 0xd1, 0x5d, 0x6c, 0x15, 0xb0, 0xf0, 0xa, 0x8, 0x74, 0x65, 0x73, 0x74}
	assert.Equal(expected, writer.Bytes())
	assert.EqualValues(written, len(writer.Bytes()))
}

func TestCreateParentNode(t *testing.T) {
	assert := assert.New(t)
	links := []Link {
		{
			Link: format.Link{
				Name: "",
				Size: 5,
				Cid:  blocks.NewBlock([]byte("hello")).Cid(),
			},
			ChunkSize: 5,
		},
		{
			Link: format.Link{
				Name: "",
				Size: 5,
				Cid:  blocks.NewBlock([]byte("world")).Cid(),
			},
			ChunkSize: 5,
		},
	}
	node, total, err := createParentNode(links)
	assert.NoError(err)
	assert.EqualValues(10, total)
	assert.EqualValues(2, len(node.Links()))
	assert.Equal("bafybeibobk2jo534bx3dyrnqefl5sy7iqivyldk7kwdxra6fu62znkg2ym", node.Cid().String())
	size, err := node.Size()
	assert.NoError(err)
	assert.EqualValues(104, size)
	expected := []byte{0x12, 0x28, 0xa, 0x22, 0x12, 0x20, 0x2c, 0xf2, 0x4d, 0xba, 0x5f, 0xb0, 0xa3, 0xe, 0x26, 0xe8, 0x3b, 0x2a, 0xc5, 0xb9, 0xe2, 0x9e, 0x1b, 0x16, 0x1e, 0x5c, 0x1f, 0xa7, 0x42, 0x5e, 0x73, 0x4, 0x33, 0x62, 0x93, 0x8b, 0x98, 0x24, 0x12, 0x0, 0x18, 0x5, 0x12, 0x28, 0xa, 0x22, 0x12, 0x20, 0x48, 0x6e, 0xa4, 0x62, 0x24, 0xd1, 0xbb, 0x4f, 0xb6, 0x80, 0xf3, 0x4f, 0x7c, 0x9a, 0xd9, 0x6a, 0x8f, 0x24, 0xec, 0x88, 0xbe, 0x73, 0xea, 0x8e, 0x5a, 0x6c, 0x65, 0x26, 0xe, 0x9c, 0xb8, 0xa7, 0x12, 0x0, 0x18, 0x5, 0xa, 0x8, 0x8, 0x2, 0x18, 0xa, 0x20, 0x5, 0x20, 0x5}
	assert.Equal(expected, node.RawData())
	assert.EqualValues(94, len(node.RawData()))
}

func TestPackItems(t *testing.T) {
	assert := assert.New(t)
	streamer := &MockHandler{}
	streamer.On("Open", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(newRandomReadCloser(10_000_000, 0), nil)
	item := model.Item{
		Type:         model.File,
		Path:         "/tmp/file1.img",
		Size:         10_000_000,
		Offset:       0,
		Length:       10_000_000,
	}
	ctx := context.Background()
	outDir := os.TempDir()
	pieceSize := uint64(1 << 35)
	result, err := PackItems(ctx, streamer, []model.Item{item}, outDir, pieceSize)
	assert.NoError(err)
	assert.Equal("/tmp/baga6ea4seaqln473vpjjacg4b3tfcgdemopmgjpq4uwxm6drmr2tawj7lvqrgki.car", result.CarFilePath)
	assert.EqualValues(10000972, result.CarFileSize)
	assert.Equal("baga6ea4seaqln473vpjjacg4b3tfcgdemopmgjpq4uwxm6drmr2tawj7lvqrgki", result.PieceCID.String())
	assert.EqualValues(1 << 35, result.PieceSize)
	assert.Equal("bafkreib5gcnatcliv7ubb3qwpygfwic66nqqqkng2nha6c5ezjthk3dply", result.RootCID.String())
	assert.Equal([]byte{0xa2, 0x65, 0x72, 0x6f, 0x6f, 0x74, 0x73, 0x81, 0xd8, 0x2a, 0x58, 0x25, 0x0, 0x1, 0x55, 0x12, 0x20, 0x3d, 0x30, 0x9a, 0x9, 0x89, 0x68, 0xaf, 0xe8, 0x10, 0xee, 0x16, 0x7e, 0xc, 0x5b, 0x20, 0x5e, 0xf3, 0x61, 0x8, 0x29, 0xa6, 0xd3, 0x4e, 0xf, 0xb, 0xa4, 0xca, 0x66, 0x75, 0x6c, 0x6f, 0x5e, 0x67, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x0}, result.Header)
	assert.Equal(1, len(result.RawBlocks))
	assert.Equal(11, len(result.CarBlocks))
	assert.Equal(10, len(result.ItemBlocks))
}

func TestPackItemsWithoutOutdir(t *testing.T) {
	assert := assert.New(t)
	streamer := &MockHandler{}
	streamer.On("Open", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(newRandomReadCloser(10_000_000, 0), nil)
	item := model.Item{
		Type:         model.File,
		Path:         "/tmp/file1.img",
		Size:         10_000_000,
		Offset:       0,
		Length:       10_000_000,
	}
	ctx := context.Background()
	pieceSize := uint64(1 << 35)
	result, err := PackItems(ctx, streamer, []model.Item{item}, "", pieceSize)
	assert.NoError(err)
	assert.Equal("", result.CarFilePath)
	assert.EqualValues(10000972, result.CarFileSize)
	assert.Equal("baga6ea4seaqln473vpjjacg4b3tfcgdemopmgjpq4uwxm6drmr2tawj7lvqrgki", result.PieceCID.String())
	assert.EqualValues(1 << 35, result.PieceSize)
	assert.Equal("bafkreib5gcnatcliv7ubb3qwpygfwic66nqqqkng2nha6c5ezjthk3dply", result.RootCID.String())
	assert.Equal([]byte{0xa2, 0x65, 0x72, 0x6f, 0x6f, 0x74, 0x73, 0x81, 0xd8, 0x2a, 0x58, 0x25, 0x0, 0x1, 0x55, 0x12, 0x20, 0x3d, 0x30, 0x9a, 0x9, 0x89, 0x68, 0xaf, 0xe8, 0x10, 0xee, 0x16, 0x7e, 0xc, 0x5b, 0x20, 0x5e, 0xf3, 0x61, 0x8, 0x29, 0xa6, 0xd3, 0x4e, 0xf, 0xb, 0xa4, 0xca, 0x66, 0x75, 0x6c, 0x6f, 0x5e, 0x67, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x0}, result.Header)
	assert.Equal(1, len(result.RawBlocks))
	assert.Equal(11, len(result.CarBlocks))
	assert.Equal(10, len(result.ItemBlocks))
}
