package daggen

import (
	"context"
	"strconv"
	"testing"

	"github.com/alecthomas/units"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/gotidy/ptr"
	"github.com/ipfs/boxo/util"
	"github.com/ipfs/go-cid"
	format "github.com/ipfs/go-ipld-format"
	"github.com/ipfs/go-unixfs/io"
	"github.com/stretchr/testify/require"
)

func TestMarshalling(t *testing.T) {
	ctx := context.Background()
	oldShardingSize := io.HAMTShardingSize
	defer func() {
		io.HAMTShardingSize = oldShardingSize
	}()
	io.HAMTShardingSize = int(units.KiB)
	var initial []byte
	dirData := &DirectoryData{}
	err := dirData.UnmarshalBinary(ctx, initial)
	require.NoError(t, err)
	for i := 0; i < 100; i++ {
		err = dirData.AddItem(ctx, strconv.Itoa(i), cid.NewCidV1(cid.Raw, util.Hash([]byte(strconv.Itoa(i)))), 4)
		require.NoError(t, err)
	}
	root, err := dirData.Node()
	require.NoError(t, err)
	initial, err = dirData.MarshalBinary(ctx)
	require.NoError(t, err)
	require.NotEmpty(t, initial)
	err = dirData.UnmarshalBinary(ctx, initial)
	require.NoError(t, err)
	newRoot, err := dirData.Node()
	require.NoError(t, err)
	links, err := dirData.dir.Links(ctx)
	require.NoError(t, err)
	require.Len(t, links, 100)
	require.Equal(t, root.Cid(), newRoot.Cid())
}

func TestDirectoryData(t *testing.T) {
	ctx := context.Background()
	d := NewDirectoryData()
	binary, err := d.MarshalBinary(ctx)
	require.NoError(t, err)
	err = d.UnmarshalBinary(ctx, binary)
	require.NoError(t, err)
	err = d.AddItem(ctx, "test", cid.NewCidV1(cid.Raw, util.Hash([]byte("test"))), 4)
	require.NoError(t, err)
	_, err = d.AddItemFromLinks(ctx, "test2", []format.Link{
		{
			Cid:  cid.NewCidV1(cid.Raw, util.Hash([]byte("test2"))),
			Size: 5,
		},
		{
			Cid:  cid.NewCidV1(cid.Raw, util.Hash([]byte("test3"))),
			Size: 5,
		},
	})
	require.NoError(t, err)
	binary, err = d.MarshalBinary(ctx)
	require.NoError(t, err)
	err = d.UnmarshalBinary(ctx, binary)
	require.NoError(t, err)
	err = d.AddItem(ctx, "test4", cid.NewCidV1(cid.Raw, util.Hash([]byte("test4"))), 5)
	require.NoError(t, err)
	_, err = d.MarshalBinary(ctx)
	require.NoError(t, err)
	err = d.UnmarshalBinary(ctx, binary)
	require.NoError(t, err)
	root, err := d.Node()
	require.NoError(t, err)
	require.Len(t, root.Links(), 2)
	size, err := root.Size()
	require.NoError(t, err)
	require.EqualValues(t, 213, size)
}

func TestResolveDirectoryTree(t *testing.T) {
	ctx := context.Background()
	root := NewDirectoryData()
	dir := NewDirectoryData()
	cache := map[uint64]*DirectoryDetail{
		1: {
			Dir:  &model.Directory{ID: 1},
			Data: &root,
		},
		2: {
			Dir:  &model.Directory{ID: 2, Name: "name", ParentID: ptr.Of(uint64(1))},
			Data: &dir,
		},
	}
	children := map[uint64][]uint64{
		1: {2},
	}
	err := root.AddItem(ctx, "test", cid.NewCidV1(cid.Raw, util.Hash([]byte("test"))), 4)
	require.NoError(t, err)
	err = dir.AddItem(ctx, "test2", cid.NewCidV1(cid.Raw, util.Hash([]byte("test2"))), 5)
	require.NoError(t, err)
	_, err = DirectoryTree{cache: cache, childrenCache: children}.Resolve(ctx, 1)
	require.NoError(t, err)
	node, err := root.Node()
	require.NoError(t, err)
	require.Equal(t, 2, len(node.Links()))
	require.Equal(t, "name", node.Links()[0].Name)
	require.Equal(t, "test", node.Links()[1].Name)
}
