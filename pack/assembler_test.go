package pack

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"testing"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/storagesystem"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/ipfs/go-cid"
	"github.com/ipld/go-car"
	"github.com/rjNemo/underscore"
	"github.com/stretchr/testify/require"
)

func TestAssembler(t *testing.T) {
	tmp := t.TempDir()
	sizes := map[int]struct {
		size        int
		encSize     int
		sizeUnknown bool
	}{
		0:                {96, 297, false},
		1:                {97, 298, false},
		1024:             {1121, 1321, false},
		1024 * 1024:      {1048674, 1049296, false},
		1024 * 1024 * 2:  {2097436, 2098218, false},
		1024*1024*2 + 1:  {2097520, 2098235, false},
		1024*1024*2 - 1:  {2097434, 2098217, false},
		1024 * 1024 * 10: {10486756, 10489586, false},
		1024*1024*10 - 1: {10486755, 10489585, false},
		1024*1024*10 + 1: {10486840, 10489603, false},
		2048:             {2145, 2345, true},
	}

	ctx := context.Background()
	reader, err := storagesystem.NewRCloneHandler(ctx, model.Storage{
		Type: "local",
		Path: tmp,
	})
	require.NoError(t, err)

	var allFileRanges []model.FileRange
	for size, val := range sizes {
		filename := fmt.Sprintf("%d.bin", size)
		err := os.WriteFile(filepath.Join(tmp, filename), testutil.GenerateRandomBytes(size), 0644)
		require.NoError(t, err)
		stat, err := os.Stat(filepath.Join(tmp, filename))
		require.NoError(t, err)
		length := int64(size)
		if val.sizeUnknown {
			length = -1
		}
		allFileRanges = append(allFileRanges, model.FileRange{
			ID:     model.FileRangeID(size),
			Offset: 0,
			Length: length,
			FileID: model.FileID(size),
			File: &model.File{
				ID:               model.FileID(size),
				Path:             filename,
				Size:             length,
				LastModifiedNano: stat.ModTime().UnixNano(),
			}})
	}

	for size, expected := range sizes {
		fileRange, err := underscore.Find(allFileRanges, func(fileRange model.FileRange) bool {
			return fileRange.ID == model.FileRangeID(size)
		})
		require.NoError(t, err)
		t.Run(fmt.Sprintf("single size=%d", size), func(t *testing.T) {
			assembler := NewAssembler(context.Background(), reader, []model.FileRange{fileRange}, false, false)
			defer assembler.Close()
			content, err := io.ReadAll(assembler)
			require.NoError(t, err)
			require.Equal(t, expected.size, len(content))
			validateCarContent(t, content)
			validateAssembler(t, assembler)
			if expected.sizeUnknown {
				require.Greater(t, len(assembler.carBlocks), 0)
				for _, corrected := range assembler.fileLengthCorrection {
					require.Equal(t, corrected, int64(size))
				}
			}
		})
	}

	sort.Slice(allFileRanges, func(i, j int) bool {
		return allFileRanges[i].ID < allFileRanges[j].ID
	})
	t.Run("all", func(t *testing.T) {
		assembler := NewAssembler(context.Background(), reader, allFileRanges, false, false)
		defer assembler.Close()
		content, err := io.ReadAll(assembler)
		require.NoError(t, err)
		require.Equal(t, 38804284, len(content))
		validateCarContent(t, content)
		validateAssembler(t, assembler)
		require.Greater(t, len(assembler.carBlocks), 0)
	})
	t.Run("noinline", func(t *testing.T) {
		assembler := NewAssembler(context.Background(), reader, allFileRanges, true, false)
		defer assembler.Close()
		content, err := io.ReadAll(assembler)
		require.NoError(t, err)
		require.Equal(t, 38804284, len(content))
		validateCarContent(t, content)
		validateAssembler(t, assembler)
		require.Len(t, assembler.carBlocks, 0)
	})
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
