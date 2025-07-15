package cmd

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/data-preservation-programs/singularity/handler/dataprep"
	"github.com/data-preservation-programs/singularity/handler/storage"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestUnionStorageDataPreparation(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		// Set up test directories
		sourceDir := t.TempDir()
		outputDir := t.TempDir()

		// Create test folders and files
		folder1 := filepath.Join(sourceDir, "folder1")
		folder2 := filepath.Join(sourceDir, "folder2")
		require.NoError(t, os.Mkdir(folder1, 0755))
		require.NoError(t, os.Mkdir(folder2, 0755))
		require.NoError(t, os.WriteFile(filepath.Join(folder1, "test1.txt"), []byte("test1"), 0644))
		require.NoError(t, os.WriteFile(filepath.Join(folder2, "test2.txt"), []byte("test2"), 0644))

		// Create storage handlers
		_, err := storage.Default.CreateStorageHandler(ctx, db, "local", storage.CreateRequest{
			Name: "source",
			Path: sourceDir,
		})
		require.NoError(t, err)

		_, err = storage.Default.CreateStorageHandler(ctx, db, "local", storage.CreateRequest{
			Name: "output",
			Path: outputDir,
		})
		require.NoError(t, err)

		// Test with one-piece-per-upstream enabled
		t.Run("one piece per upstream", func(t *testing.T) {
			prep, err := dataprep.Default.CreatePreparationHandler(ctx, db, dataprep.CreateRequest{
				Name:                "test_prep_1",
				SourceStorages:      []string{"source"},
				OutputStorages:      []string{"output"},
				MaxSizeStr:         "31GiB",
				OnePiecePerUpstream: true,
			})
			require.NoError(t, err)
			require.NotNil(t, prep)

			// Verify that we got one piece per upstream folder
			pieces, err := model.ListPieces(db, prep.ID)
			require.NoError(t, err)
			require.Len(t, pieces, 2) // Should have one piece per folder

			// Verify each piece corresponds to a folder
			piecesByPath := make(map[string]bool)
			for _, piece := range pieces {
			   piecesByPath[piece.StoragePath] = true
			}
			require.True(t, piecesByPath[folder1])
			require.True(t, piecesByPath[folder2])
		})

		// Test with one-piece-per-upstream disabled (default behavior)
		t.Run("default behavior", func(t *testing.T) {
			prep, err := dataprep.Default.CreatePreparationHandler(ctx, db, dataprep.CreateRequest{
				Name:                "test_prep_2",
				SourceStorages:      []string{"source"},
				OutputStorages:      []string{"output"},
				MaxSizeStr:         "31GiB",
				OnePiecePerUpstream: false,
			})
			require.NoError(t, err)
			require.NotNil(t, prep)

			// Verify default behavior
			pieces, err := model.ListPieces(db, prep.ID)
			require.NoError(t, err)
			require.True(t, len(pieces) <= 2) // Should have at most 2 pieces (might be combined into 1)
		})
	})
}
