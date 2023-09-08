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
		size    int
		encSize int
	}{
		0:                {96, 297},
		1:                {97, 298},
		1024:             {1121, 1321},
		1024 * 1024:      {1048674, 1049296},
		1024 * 1024 * 2:  {2097436, 2098218},
		1024*1024*2 + 1:  {2097520, 2098235},
		1024*1024*2 - 1:  {2097434, 2098217},
		1024 * 1024 * 10: {10486756, 10489586},
		1024*1024*10 - 1: {10486755, 10489585},
		1024*1024*10 + 1: {10486840, 10489603},
	}

	ctx := context.Background()
	reader, err := storagesystem.NewRCloneHandler(ctx, model.Storage{
		Type: "local",
		Path: tmp,
	})
	require.NoError(t, err)

	var allFileRanges []model.FileRange
	for size := range sizes {
		filename := fmt.Sprintf("%d.bin", size)
		err := os.WriteFile(filepath.Join(tmp, filename), testutil.GenerateRandomBytes(size), 0644)
		require.NoError(t, err)
		stat, err := os.Stat(filepath.Join(tmp, filename))
		require.NoError(t, err)
		allFileRanges = append(allFileRanges, model.FileRange{
			ID:     uint64(size),
			Offset: 0,
			Length: int64(size),
			FileID: uint64(size),
			File: &model.File{
				ID:               uint64(size),
				Path:             filename,
				Size:             int64(size),
				LastModifiedNano: stat.ModTime().UnixNano(),
			}})
	}

	for size, expected := range sizes {
		fileRange, err := underscore.Find(allFileRanges, func(fileRange model.FileRange) bool {
			return fileRange.ID == uint64(size)
		})
		require.NoError(t, err)
		t.Run(fmt.Sprintf("single size=%d", size), func(t *testing.T) {
			assembler := NewAssembler(context.Background(), reader, []model.FileRange{fileRange})
			defer assembler.Close()
			content, err := io.ReadAll(assembler)
			require.NoError(t, err)
			require.Equal(t, expected.size, len(content))
			validateCarContent(t, content)
			validateAssembler(t, assembler)
		})
	}

	sort.Slice(allFileRanges, func(i, j int) bool {
		return allFileRanges[i].ID < allFileRanges[j].ID
	})
	t.Run("all", func(t *testing.T) {
		assembler := NewAssembler(context.Background(), reader, allFileRanges)
		defer assembler.Close()
		content, err := io.ReadAll(assembler)
		require.NoError(t, err)
		require.Equal(t, 38802198, len(content))
		validateCarContent(t, content)
		validateAssembler(t, assembler)
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
