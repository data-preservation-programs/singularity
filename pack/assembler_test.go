package pack

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/storagesystem"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/ipfs/go-cid"
	"github.com/ipld/go-car"
	"github.com/stretchr/testify/require"
)

func TestAssembler_NoEncryption(t *testing.T) {
	tmp := t.TempDir()
	sizes := map[int]int{
		0:                96,
		1:                97,
		1024:             1121,
		1024 * 1024:      1048674,
		1024 * 1024 * 2:  2097436,
		1024*1024*2 + 1:  2097520,
		1024*1024*2 - 1:  2097434,
		1024 * 1024 * 10: 10486756,
		1024*1024*10 - 1: 10486755,
		1024*1024*10 + 1: 10486840,
	}

	ctx := context.Background()
	reader, err := storagesystem.NewRCloneHandler(ctx, model.Storage{
		Type: "local",
		Path: tmp,
	})
	require.NoError(t, err)

	var allFileRanges []model.FileRange
	for size, expected := range sizes {
		t.Run(fmt.Sprintf("size=%d", size), func(t *testing.T) {
			filename := fmt.Sprintf("%d.bin", size)
			err := os.WriteFile(filepath.Join(tmp, filename), testutil.GenerateRandomBytes(size), 0644)
			require.NoError(t, err)
			stat, err := os.Stat(filepath.Join(tmp, filename))
			require.NoError(t, err)
			fileRanges := []model.FileRange{{
				ID:     uint64(size),
				Offset: 0,
				Length: int64(size),
				FileID: uint64(size),
				File: &model.File{
					ID:               uint64(size),
					Path:             filename,
					Size:             int64(size),
					LastModifiedNano: stat.ModTime().UnixNano(),
				},
			}}
			allFileRanges = append(allFileRanges, fileRanges...)
			assembler := NewAssembler(context.Background(), reader, nil, fileRanges, 30*1024*1024)
			defer assembler.Close()
			content, err := io.ReadAll(assembler)
			require.NoError(t, err)
			require.Equal(t, expected, len(content))
			validateCarContent(t, content)
			validateAssembler(t, assembler)
		})
	}

	t.Run("all", func(t *testing.T) {
		assembler := NewAssembler(context.Background(), reader, nil, allFileRanges, 100*1024*1024)
		defer assembler.Close()
		content, err := io.ReadAll(assembler)
		require.NoError(t, err)
		require.Equal(t, 38802198, len(content))
		validateCarContent(t, content)
		validateAssembler(t, assembler)
	})

	maxSizes := map[int64][]int{
		3000000: []int{10487722, 10486756, 10486755, 7341142},
	}
	for maxSize, contentSizes := range maxSizes {
		t.Run(fmt.Sprintf("maxSize=%d", maxSize), func(t *testing.T) {
			assembler := NewAssembler(context.Background(), reader, nil, allFileRanges, maxSize)
			defer assembler.Close()
			var actualSizes []int
			for assembler.Next() {
				content, err := io.ReadAll(assembler)
				require.NoError(t, err)
				actualSizes = append(actualSizes, len(content))
				validateCarContent(t, content)
			}
			require.EqualValues(t, contentSizes, actualSizes)
			validateAssembler(t, assembler)
		})
	}
}

func validateCarContent(t *testing.T, content []byte) {
	reader, err := car.NewCarReader(bytes.NewReader(content))
	require.NoError(t, err)
	require.NotNil(t, reader.Header)
	for {
		_, err := reader.Next()
		if err == io.EOF {
			break
		}
		require.NoError(t, err)
	}
}

func validateAssembler(t *testing.T, assembler *Assembler) {
	for _, carBlock := range assembler.carBlocks {
		require.True(t, carBlock.FileID != nil || carBlock.RawBlock != nil)
	}
	for _, fileRange := range assembler.fileRanges {
		require.NotEqual(t, cid.Undef, cid.Cid(fileRange.CID))
	}
	require.Nil(t, assembler.buffer)
}
