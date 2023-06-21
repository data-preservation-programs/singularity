package pack

import (
	"bytes"
	"context"
	"github.com/data-preservation-programs/singularity/model"
	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"
	util "github.com/ipfs/go-ipfs-util"
	format "github.com/ipfs/go-ipld-format"
	"github.com/rclone/rclone/fs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"testing"
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
	assert.NoError(t, err)
	assert.Equal(t, uint64(10), size)
	assert.NotNil(t, node)
	size, err = node.Size()
	assert.NoError(t, err)
	assert.Equal(t, uint64(108), size)
	assert.Equal(t, "bafybeiejlvvmfokp5c6q2eqgbfjeaokz3nqho5c7yy3ov527vsatgsqfma", node.String())
}

func TestMind(t *testing.T) {
	assert.Equal(t, 1, Min(1, 2))
	assert.Equal(t, 1, Min(2, 1))
	assert.Equal(t, 1, Min(1, 1))
}

func TestAssembleItemFromLinks(t *testing.T) {
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

	blocks, node, err := AssembleItemFromLinks(links)
	assert.NoError(t, err)
	assert.NotNil(t, node)
	size, err := node.Size()
	assert.NoError(t, err)
	assert.Equal(t, uint64(108), size)
	assert.Equal(t, "bafybeiejlvvmfokp5c6q2eqgbfjeaokz3nqho5c7yy3ov527vsatgsqfma", node.String())
	assert.Len(t, blocks, 1)
	assert.Equal(t, "bafybeiejlvvmfokp5c6q2eqgbfjeaokz3nqho5c7yy3ov527vsatgsqfma", blocks[0].Cid().String())
}

func TestAssembleItemFromLinks_LargeFile(t *testing.T) {
	links := []format.Link{}
	for i := 0; i < 2_000_000; i++ {
		links = append(links, format.Link{
			Name: "",
			Size: 5,
			Cid:  cid.NewCidV1(cid.Raw, util.Hash([]byte("hello"))),
		})
	}

	blocks, node, err := AssembleItemFromLinks(links)
	assert.NoError(t, err)
	assert.NotNil(t, node)
	size, err := node.Size()
	assert.NoError(t, err)
	assert.Equal(t, uint64(10000113), size)
	assert.Equal(t, "bafybeidd6nph3lz2c34mvxoqsaxx6er2au7c73j32yfs6rhmiqeh35ko6m", node.String())
	assert.Len(t, blocks, 1957)
	assert.Equal(t, "bafybeidci5xcdskgvefhv3x6kp6zpyxkbiqshqjbpsdgmk2k3mexzaggwi", blocks[0].Cid().String())
}

func TestGenerateCarHeader(t *testing.T) {
	header, err := GenerateCarHeader(cid.NewCidV1(cid.Raw, util.Hash([]byte("hello"))))
	assert.NoError(t, err)
	assert.Len(t, header, 59)
}

func TestWriteCarHeader(t *testing.T) {
	buf := new(bytes.Buffer)
	header, err := WriteCarHeader(buf, cid.NewCidV1(cid.Raw, util.Hash([]byte("hello"))))
	assert.NoError(t, err)
	assert.Len(t, header, 59)
	assert.Equal(t, buf.Bytes(), header)
}

func TestWriteCarBlock(t *testing.T) {
	buf := new(bytes.Buffer)
	block := blocks.NewBlock([]byte("hello"))
	n, err := WriteCarBlock(buf, block)
	assert.NoError(t, err)
	assert.Equal(t, int64(40), n)
	assert.Equal(t, 40, buf.Len())
}

type MockReadHandler struct {
	mock.Mock
}

func (m *MockReadHandler) Read(ctx context.Context, path string, offset int64, length int64) (io.ReadCloser, fs.Object, error) {
	args := m.Called(ctx, path, offset, length)
	return args.Get(0).(io.ReadCloser), nil, args.Error(2)
}

func TestGetBlockStreamFromItem(t *testing.T) {
	ctx := context.Background()
	handler := new(MockReadHandler)
	handler.On("Read", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(io.NopCloser(bytes.NewReader([]byte("hello"))), nil, nil)
	item := model.ItemPart{
		Offset: 0,
		Length: 5,
		Item: &model.Item{
			Size: 5,
		},
	}
	blockResultChan, _, err := GetBlockStreamFromItem(ctx, handler, item, nil)
	assert.NoError(t, err)
	blockResults := make([]BlockResult, 0)
	for r := range blockResultChan {
		blockResults = append(blockResults, r)
	}
	assert.Len(t, blockResults, 1)
	assert.EqualValues(t, 0, blockResults[0].Offset)
	assert.Equal(t, []byte("hello"), blockResults[0].Raw)
	assert.Equal(t, "bafkreibm6jg3ux5qumhcn2b3flc3tyu6dmlb4xa7u5bf44yegnrjhc4yeq", blockResults[0].CID.String())
}
