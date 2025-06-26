//go:build !windows

package storagesystem

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/stretchr/testify/require"
)

func TestInAccessibleFiles(t *testing.T) {
	// Skip test if running as root (permissions don't work the same way)
	if os.Getuid() == 0 {
		t.Skip("Skipping file permission test when running as root")
	}

	tmp := t.TempDir()
	// Inaccessible folder
	err := os.MkdirAll(filepath.Join(tmp, "sub"), 0000)
	require.NoError(t, err)

	// Inaccessible file
	err = os.WriteFile(filepath.Join(tmp, "test.txt"), []byte("test"), 0000)
	require.NoError(t, err)

	// Accessible folder and file
	err = os.MkdirAll(filepath.Join(tmp, "sub2"), 0755)
	require.NoError(t, err)
	err = os.WriteFile(filepath.Join(tmp, "test2.txt"), []byte("test"), 0644)
	require.NoError(t, err)

	// Verify that permissions are actually working
	_, err = os.Open(filepath.Join(tmp, "sub"))
	if err == nil {
		t.Skip("File permissions not enforced on this system - cannot test inaccessible file behavior")
	}

	ctx := context.Background()
	handler, err := NewRCloneHandler(ctx, model.Storage{
		Type: "local",
		Path: tmp,
	})
	require.NoError(t, err)
	t.Run("list", func(t *testing.T) {
		entries, err := handler.List(ctx, "")
		require.NoError(t, err)
		require.Len(t, entries, 4)
		// Interesting, listing inaccessible folder does not return error
		entries, err = handler.List(ctx, "sub")
		require.NoError(t, err)
		require.Len(t, entries, 0)
	})

	t.Run("scan", func(t *testing.T) {
		entryChan := handler.Scan(ctx, "")
		require.NotNil(t, entryChan)
		scannedEntries := []Entry{}
		for entry := range entryChan {
			if entry.Info == nil {
				continue
			}
			scannedEntries = append(scannedEntries, entry)
		}
		require.Len(t, scannedEntries, 2)
		// Inaccessible folder does not return error during scanning
		for _, entry := range scannedEntries {
			require.NoError(t, entry.Error)
		}
	})

	t.Run("read", func(t *testing.T) {
		// Inaccessible file will return error during reading
		_, _, err := handler.Read(ctx, "test.txt", 0, 4)
		require.Error(t, err)
	})
}
