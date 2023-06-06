package daggen

import (
	"github.com/ipfs/go-cid"
	util "github.com/ipfs/go-ipfs-util"
	format "github.com/ipfs/go-ipld-format"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDirectoryData(t *testing.T) {
	d := NewDirectoryData()
	binary, err := d.MarshalBinary()
	assert.NoError(t, err)
	assert.Len(t, binary, 96)
	err = d.UnmarshallBinary(binary)
	assert.NoError(t, err)
	err = d.AddItem("test", cid.NewCidV1(cid.Raw, util.Hash([]byte("test"))), 4)
	assert.NoError(t, err)
	err = d.AddItemFromLinks("test2", []format.Link{
		{
			Cid: cid.NewCidV1(cid.Raw, util.Hash([]byte("test2"))),
			Size: 5,
		},
		{
			Cid: cid.NewCidV1(cid.Raw, util.Hash([]byte("test3"))),
			Size: 5,
		},
	})
	binary, err = d.MarshalBinary()
	assert.NoError(t, err)
	assert.Len(t, binary, 377)
	err = d.UnmarshallBinary(binary)
	assert.NoError(t, err)
	err = d.AddItem("test4", cid.NewCidV1(cid.Raw, util.Hash([]byte("test4"))), 5)
	binary, err = d.MarshalBinary()
	assert.NoError(t, err)
	assert.Len(t, binary, 568)
}

func TestResolveDirectoryTree(t *testing.T) {
	dirCache := make(map[uint64]*DirectoryData)
	childrenCache := make(map[uint64][]uint64)
	root := NewDirectoryData()
	root.Directory.ID = 1
	err := root.AddItem("test", cid.NewCidV1(cid.Raw, util.Hash([]byte("test"))), 4)
	assert.NoError(t, err)
	dir := NewDirectoryData()
	err = dir.AddItem("test2", cid.NewCidV1(cid.Raw, util.Hash([]byte("test2"))), 5)
	dir.Directory.ID = 2
	dir.Directory.Name = "name"
	parentID := uint64(1)
	dir.Directory.ParentID = &parentID
	dirCache[2] = &dir
	dirCache[1] = &root
	childrenCache[1] = []uint64{2}
	_, err = ResolveDirectoryTree(1, dirCache, childrenCache)
	assert.Equal(t, 2, len(root.Node.Links()))
	assert.Equal(t, "name", root.Node.Links()[0].Name)
	assert.Equal(t, "test", root.Node.Links()[1].Name)
}
