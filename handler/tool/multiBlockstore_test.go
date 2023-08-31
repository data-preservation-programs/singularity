package tool

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/singularity/util"
	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-datastore"
	blockstore "github.com/ipfs/go-ipfs-blockstore"
	ipld "github.com/ipfs/go-ipld-format"
	"github.com/stretchr/testify/require"
)

func TestMultiBlockStore(t *testing.T) {
	ctx := context.Background()
	bstore1 := blockstore.NewBlockstore(datastore.NewMapDatastore())
	bstore2 := blockstore.NewBlockstore(datastore.NewMapDatastore())
	mbstore := &multiBlockstore{bss: []blockstore.Blockstore{bstore1, bstore2}}
	block1 := blocks.NewBlock([]byte("test1"))
	block2 := blocks.NewBlock([]byte("test2"))
	cid1 := block1.Cid()
	cid2 := block2.Cid()

	err := bstore1.Put(ctx, block1)
	require.NoError(t, err)
	err = bstore2.Put(ctx, block2)
	require.NoError(t, err)

	err = mbstore.DeleteBlock(ctx, cid.Undef)
	require.ErrorIs(t, err, util.ErrNotImplemented)

	has, err := mbstore.Has(ctx, cid1)
	require.NoError(t, err)
	require.True(t, has)

	has, err = mbstore.Has(ctx, cid2)
	require.NoError(t, err)
	require.True(t, has)

	has, err = mbstore.Has(ctx, cid.Undef)
	require.NoError(t, err)
	require.False(t, has)

	blk1, err := mbstore.Get(ctx, cid1)
	require.NoError(t, err)
	require.Equal(t, block1, blk1)

	blk2, err := mbstore.Get(ctx, cid2)
	require.NoError(t, err)
	require.Equal(t, block2, blk2)

	_, err = mbstore.Get(ctx, cid.Undef)
	require.ErrorIs(t, err, ipld.ErrNotFound{Cid: cid.Undef})

	size, err := mbstore.GetSize(ctx, cid1)
	require.NoError(t, err)
	require.Equal(t, len(block1.RawData()), size)

	size, err = mbstore.GetSize(ctx, cid2)
	require.NoError(t, err)
	require.Equal(t, len(block2.RawData()), size)

	err = mbstore.Put(ctx, block1)
	require.ErrorIs(t, err, util.ErrNotImplemented)

	err = mbstore.PutMany(ctx, nil)
	require.ErrorIs(t, err, util.ErrNotImplemented)

	_, err = mbstore.AllKeysChan(ctx)
	require.ErrorIs(t, err, util.ErrNotImplemented)

	mbstore.HashOnRead(true)
}
