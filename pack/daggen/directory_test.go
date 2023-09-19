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

func benchmarkMarshal(i int, size int, b *testing.B) {
	b.Helper()
	dirData := NewDirectoryData()
	for j := 0; j < i; j++ {
		err := dirData.AddFile(context.Background(), strconv.Itoa(j), cid.NewCidV1(cid.Raw, util.Hash([]byte(strconv.Itoa(j)))), 4)
		require.NoError(b, err)
	}
	for n := 0; n < b.N; n++ {
		out, err := dirData.MarshalBinary(context.Background())
		require.NoError(b, err)
		require.LessOrEqual(b, len(out), size)
		err = dirData.UnmarshalBinary(context.Background(), out)
		require.NoError(b, err)
	}
}

func BenchmarkMarshal_0(b *testing.B) {
	benchmarkMarshal(0, 69, b)
}

func BenchmarkMarshal_1(b *testing.B) {
	benchmarkMarshal(1, 120, b)
}

func BenchmarkMarshal_100(b *testing.B) {
	benchmarkMarshal(100, 4200, b)
}

func BenchmarkMarshal_10000(b *testing.B) {
	benchmarkMarshal(10000, 500000, b)
}

func BenchmarkMarshal_1000000(b *testing.B) {
	benchmarkMarshal(1000000, 85000000, b)
}

func TestDataSize_HAMTDirectory(t *testing.T) {
	ctx := context.Background()
	oldShardingSize := io.HAMTShardingSize
	defer func() {
		io.HAMTShardingSize = oldShardingSize
	}()
	io.HAMTShardingSize = int(units.KiB)

	initial := []byte{}
	dirData := &DirectoryData{}
	err := dirData.UnmarshalBinary(ctx, initial)
	require.NoError(t, err)
	for i := 0; i < 100; i++ {
		err = dirData.AddFile(ctx, strconv.Itoa(i), cid.NewCidV1(cid.Raw, util.Hash([]byte(strconv.Itoa(i)))), 4)
		require.NoError(t, err)
	}
	initial, err = dirData.MarshalBinary(ctx)
	require.NoError(t, err)
	require.LessOrEqual(t, len(initial), 5500)

	for i := 0; i < 10; i++ {
		dirData := &DirectoryData{}
		err := dirData.UnmarshalBinary(ctx, initial)
		require.NoError(t, err)
		initial, err = dirData.MarshalBinary(ctx)
		require.NoError(t, err)
	}
	require.LessOrEqual(t, len(initial), 5500)
}

func TestDataSize_BasicDirectory(t *testing.T) {
	ctx := context.Background()

	var initial []byte
	for i := 0; i < 100; i++ {
		dirData := &DirectoryData{}
		err := dirData.UnmarshalBinary(ctx, initial)
		require.NoError(t, err)
		err = dirData.AddFile(ctx, strconv.Itoa(i), cid.NewCidV1(cid.Raw, util.Hash([]byte(strconv.Itoa(i)))), 4)
		require.NoError(t, err)
		initial, err = dirData.MarshalBinary(ctx)
		require.NoError(t, err)
	}

	require.LessOrEqual(t, len(initial), 4200)

	initial = []byte{}
	dirData := &DirectoryData{}
	err := dirData.UnmarshalBinary(ctx, initial)
	require.NoError(t, err)
	for i := 0; i < 100; i++ {
		err = dirData.AddFile(ctx, strconv.Itoa(i), cid.NewCidV1(cid.Raw, util.Hash([]byte(strconv.Itoa(i)))), 4)
		require.NoError(t, err)
	}
	initial, err = dirData.MarshalBinary(ctx)
	require.NoError(t, err)
	require.LessOrEqual(t, len(initial), 9000)

	for i := 0; i < 100; i++ {
		dirData := &DirectoryData{}
		err := dirData.UnmarshalBinary(ctx, initial)
		require.NoError(t, err)
		initial, err = dirData.MarshalBinary(ctx)
		require.NoError(t, err)
	}
	require.LessOrEqual(t, len(initial), 9000)
}

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
		err = dirData.AddFile(ctx, strconv.Itoa(i), cid.NewCidV1(cid.Raw, util.Hash([]byte(strconv.Itoa(i)))), 4)
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
	err = d.AddFile(ctx, "test", cid.NewCidV1(cid.Raw, util.Hash([]byte("test"))), 4)
	require.NoError(t, err)
	_, err = d.AddFileFromLinks(ctx, "test2", []format.Link{
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
	err = d.AddFile(ctx, "test4", cid.NewCidV1(cid.Raw, util.Hash([]byte("test4"))), 5)
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
	cache := map[model.DirectoryID]*DirectoryDetail{
		1: {
			Dir:  &model.Directory{ID: 1},
			Data: &root,
		},
		2: {
			Dir:  &model.Directory{ID: 2, Name: "name", ParentID: ptr.Of(model.DirectoryID(1))},
			Data: &dir,
		},
	}
	children := map[model.DirectoryID][]model.DirectoryID{
		1: {2},
	}
	err := root.AddFile(ctx, "test", cid.NewCidV1(cid.Raw, util.Hash([]byte("test"))), 4)
	require.NoError(t, err)
	err = dir.AddFile(ctx, "test2", cid.NewCidV1(cid.Raw, util.Hash([]byte("test2"))), 5)
	require.NoError(t, err)
	_, err = DirectoryTree{cache: cache, childrenCache: children}.Resolve(ctx, 1)
	require.NoError(t, err)
	node, err := root.Node()
	require.NoError(t, err)
	require.Equal(t, 2, len(node.Links()))
	require.Equal(t, "name", node.Links()[0].Name)
	require.Equal(t, "test", node.Links()[1].Name)
}
