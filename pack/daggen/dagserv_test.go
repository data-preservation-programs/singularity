package daggen

import (
	"context"
	"testing"

	"github.com/ipfs/boxo/ipld/merkledag"
	"github.com/ipfs/boxo/util"
	"github.com/ipfs/go-cid"
	format "github.com/ipfs/go-ipld-format"
	"github.com/stretchr/testify/require"
)

func TestRecordedDagService(t *testing.T) {
	s := NewRecordedDagService()
	c1 := cid.NewCidV1(cid.Raw, util.Hash([]byte("")))
	c2 := cid.NewCidV1(cid.Raw, util.Hash([]byte("123")))
	c3 := cid.NewCidV1(cid.Raw, util.Hash([]byte("hello")))
	c4 := cid.NewCidV1(cid.Raw, util.Hash([]byte("not exist")))
	dummy1 := NewDummyNode(0, c1)
	dummy2 := NewDummyNode(3, c2)
	node3 := merkledag.NewRawNode([]byte("hello"))
	err := s.Add(context.TODO(), dummy1)
	require.NoError(t, err)
	err = s.Add(context.TODO(), dummy2)
	require.NoError(t, err)
	err = s.Add(context.TODO(), node3)
	require.NoError(t, err)

	node1, err := s.Get(context.TODO(), c1)
	require.NoError(t, err)
	require.Equal(t, dummy1, node1)

	node2, err := s.Get(context.TODO(), c2)
	require.NoError(t, err)
	require.Equal(t, dummy2, node2)

	node4, err := s.Get(context.TODO(), c3)
	require.NoError(t, err)
	require.Equal(t, node3.RawData(), node4.RawData())

	_, err = s.Get(context.TODO(), c4)
	require.ErrorIs(t, err, format.ErrNotFound{})

	for _, v := range s.blockstore {
		require.True(t, v.visited)
	}

	s.ResetVisited()

	for _, v := range s.blockstore {
		require.False(t, v.visited)
	}

	for c := range s.blockstore {
		s.Visit(context.TODO(), c)
	}

	for _, v := range s.blockstore {
		require.True(t, v.visited)
	}
}
