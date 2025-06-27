//go:build !windows

package pack

import (
	"context"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/storagesystem"
	"github.com/stretchr/testify/require"
)

func TestAssembler_InaccessibleFile(t *testing.T) {
	tmp := t.TempDir()
	ctx := context.Background()
	err := os.WriteFile(filepath.Join(tmp, "test.txt"), []byte("test"), 0000)
	require.NoError(t, err)

	stat, err := os.Stat(filepath.Join(tmp, "test.txt"))
	require.NoError(t, err)

	reader, err := storagesystem.NewRCloneHandler(ctx, model.Storage{
		Type: "local",
		Path: tmp,
	})
	require.NoError(t, err)

	assembler := NewAssembler(context.Background(), reader, []model.FileRange{
		{
			Offset: 0,
			Length: 4,
			File: &model.File{
				Path:             "test.txt",
				Size:             4,
				LastModifiedNano: stat.ModTime().UnixNano(),
			},
		},
	}, false, false)
	defer func() { _ = assembler.Close() }()

	_, err = io.ReadAll(assembler)
	require.Error(t, err)

	assembler2 := NewAssembler(context.Background(), reader, []model.FileRange{
		{
			Offset: 0,
			Length: 4,
			File: &model.File{
				Path:             "test.txt",
				Size:             4,
				LastModifiedNano: stat.ModTime().UnixNano(),
			},
		},
	}, false, true)
	defer func() { _ = assembler2.Close() }()

	_, err = io.ReadAll(assembler2)
	require.NoError(t, err)
}
