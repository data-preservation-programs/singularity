package push

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/data-preservation-programs/singularity/pack/packutil"
	blocks "github.com/ipfs/go-block-format"
	format "github.com/ipfs/go-ipld-format"
	"github.com/stretchr/testify/require"
)

func TestToCarSize(t *testing.T) {
	jobs := []struct {
		origin   int64
		estimate int64
		actual   int64
	}{
		{0, 37, 37},
		{1, 40, 38},
		{1001, 1040, 1039},
		{1048576, 1048615, 1048615},
		{1048577, 1048803, 1048795},
		{2097152, 2097378, 2097377},
		{10485760, 10486714, 10486697},
		{104857600, 104866744, 104866548},
	}
	for _, job := range jobs {
		t.Run(fmt.Sprintf("%d", job.origin), func(t *testing.T) {
			require.Equal(t, job.estimate, toCarSize(job.origin))
			require.Equal(t, job.actual, toCarSizeExpensive(t, job.origin))
			require.GreaterOrEqual(t, job.estimate, job.actual)
		})
	}
}

func toCarSizeExpensive(t *testing.T, size int64) int64 {
	t.Helper()
	written := int64(0)
	writer := bytes.NewBuffer(nil)
	if size == 0 {
		blk, err := blocks.NewBlockWithCid([]byte{}, packutil.EmptyFileCid)
		require.NoError(t, err)
		n, err := packutil.WriteCarBlock(writer, blk)
		require.NoError(t, err)
		written += n
		return written
	}

	var links []format.Link
	for size > 0 {
		blkSize := packutil.ChunkSize
		if size < packutil.ChunkSize {
			blkSize = size
		}
		blk, err := blocks.NewBlockWithCid(make([]byte, blkSize), packutil.EmptyFileCid)
		require.NoError(t, err)
		n, err := packutil.WriteCarBlock(writer, blk)
		require.NoError(t, err)
		written += n
		links = append(links, format.Link{
			Size: uint64(blkSize),
			Cid:  blk.Cid(),
		})
		size -= blkSize
	}

	if len(links) > 1 {
		blks, _, err := packutil.AssembleFileFromLinks(links)
		require.NoError(t, err)
		for _, blk := range blks {
			n, err := packutil.WriteCarBlock(writer, blk)
			require.NoError(t, err)
			written += n
		}
	}
	return written
}
