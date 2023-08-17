package pack

import (
	"bytes"
	"context"
	"io"
	"testing"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/ipfs/boxo/util"
	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"
	format "github.com/ipfs/go-ipld-format"
	"github.com/rclone/rclone/fs"
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

func TestMind(t *testing.T) {
	require.Equal(t, 1, Min(1, 2))
	require.Equal(t, 1, Min(2, 1))
	require.Equal(t, 1, Min(1, 1))
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

func TestGenerateCarHeader(t *testing.T) {
	header, err := GenerateCarHeader(cid.NewCidV1(cid.Raw, util.Hash([]byte("hello"))))
	require.NoError(t, err)
	require.Len(t, header, 59)
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

func (m *MockReadHandler) Read(ctx context.Context, path string, offset int64, length int64) (io.ReadCloser, fs.Object, error) {
	args := m.Called(ctx, path, offset, length)
	return args.Get(0).(io.ReadCloser), nil, args.Error(2)
}

func TestGetBlockStreamFromFile(t *testing.T) {
	ctx := context.Background()
	handler := new(MockReadHandler)
	handler.On("Read", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(io.NopCloser(bytes.NewReader([]byte("hello"))), nil, nil)
	file := model.FileRange{
		Offset: 0,
		Length: 5,
		File: &model.File{
			Size: 5,
		},
	}
	blockResultChan, _, err := GetBlockStreamFromFile(ctx, handler, file, nil)
	require.NoError(t, err)
	blockResults := make([]BlockResult, 0)
	for r := range blockResultChan {
		blockResults = append(blockResults, r)
	}
	require.Len(t, blockResults, 1)
	require.EqualValues(t, 0, blockResults[0].Offset)
	require.Equal(t, []byte("hello"), blockResults[0].Raw)
	require.Equal(t, "bafkreibm6jg3ux5qumhcn2b3flc3tyu6dmlb4xa7u5bf44yegnrjhc4yeq", blockResults[0].CID.String())
}
