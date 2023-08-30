package daggen

import (
	"testing"

	"github.com/ipfs/boxo/util"
	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/require"
)

func TestDummyNode(t *testing.T) {
	cidValue := cid.NewCidV1(cid.Raw, util.Hash([]byte("test")))
	node := NewDummyNode(4, cidValue)
	require.Nil(t, node.RawData())
	require.Equal(t, cidValue, node.Cid())
	require.Contains(t, node.String(), "DummyNode - ")
	require.Nil(t, node.Loggable())
	_, _, err := node.Resolve(nil)
	require.ErrorIs(t, err, ErrDummyNode)
	require.Nil(t, node.Tree("", 0))
	_, _, err = node.ResolveLink(nil)
	require.ErrorIs(t, err, ErrDummyNode)
	copied := node.Copy()
	require.Equal(t, *node, *copied.(*DummyNode))
	require.Nil(t, node.Links())
	_, err = node.Stat()
	require.ErrorIs(t, err, ErrDummyNode)
	size, err := node.Size()
	require.NoError(t, err)
	require.Equal(t, uint64(4), size)
}
