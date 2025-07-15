package cliutil


import (
	"context"
	"flag"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"encoding/json"

	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
)

func TestOnePiecePerUpstreamFlag(t *testing.T) {
	app := cli.NewApp()
	app.Flags = CommonDealFlags // This includes our one-piece-per-upstream flag

	t.Run("flag disabled", func(t *testing.T) {
		set := flag.NewFlagSet("test", flag.ContinueOnError)
		for _, f := range app.Flags {
			f.Apply(set)
		}
		ctx := cli.NewContext(app, set, nil)
		require.False(t, ctx.Bool("one-piece-per-upstream"))
	})

	t.Run("flag enabled", func(t *testing.T) {
		set := flag.NewFlagSet("test", flag.ContinueOnError)
		for _, f := range app.Flags {
			f.Apply(set)
		}
		require.NoError(t, set.Parse([]string{"--one-piece-per-upstream"}))
		ctx := cli.NewContext(app, set, nil)
		require.True(t, ctx.Bool("one-piece-per-upstream"))
	})
}

func TestOnePiecePerUpstreamCLI(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		// Set up test directories
		tempDir := t.TempDir()
		folder1 := filepath.Join(tempDir, "folder1")
		folder2 := filepath.Join(tempDir, "folder2")
		require.NoError(t, os.Mkdir(folder1, 0755))
		require.NoError(t, os.Mkdir(folder2, 0755))

		// Create test files in each folder (16MB each to ensure they don't get combined)
		testFile1 := filepath.Join(folder1, "test1.txt")
		testFile2 := filepath.Join(folder2, "test2.txt")
		data := make([]byte, 16*1024*1024) // 16MB
		require.NoError(t, os.WriteFile(testFile1, data, 0644))
		require.NoError(t, os.WriteFile(testFile2, data, 0644))

		// Set up the database connection environment variable
		t.Setenv("DATABASE_CONNECTION_STRING", "sqlite::memory:")

		// Get the path to the singularity binary
			   singularityPath := "./singularity"
			   if _, err := os.Stat(singularityPath); err != nil {
					   t.Skip("local singularity binary not found; please build before running test")
			   }

		t.Run("with one piece per upstream", func(t *testing.T) {
			outputDir := filepath.Join(t.TempDir(), "output1")
			require.NoError(t, os.Mkdir(outputDir, 0755))

			cmd := exec.Command(singularityPath, "prep", "create",
				"--name", "test-prep-1",
				"--local-source", folder1,
				"--local-source", folder2,
				"--local-output", outputDir,
				"--one-piece-per-upstream",
			)
			output, err := cmd.CombinedOutput()
			require.NoError(t, err, "command failed: %s", string(output))
			require.Contains(t, string(output), "created successfully")

			// Verify that we have two separate CAR files (one for each folder)
			files, err := os.ReadDir(outputDir)
			require.NoError(t, err)
			carFiles := 0
			for _, f := range files {
				if filepath.Ext(f.Name()) == ".car" {
					carFiles++
				}
			}
			require.Equal(t, 2, carFiles, "Expected one CAR file per folder")
			
			// Verify that each CAR file has a unique piece CID using prep piece list-pieces
			cmd = exec.Command(singularityPath, "prep", "piece", "list-pieces", "test-prep-1", "--json")
			output, err = cmd.CombinedOutput()
			require.NoError(t, err, "list-pieces failed: %s", string(output))

			// Parse the JSON output
			type Car struct {
				PieceCID string `json:"pieceCid"`
			}
			type PieceList struct {
				Pieces []Car `json:"pieces"`
			}
			var pieceLists []PieceList
			require.NoError(t, json.Unmarshal(output, &pieceLists), "failed to parse list-pieces output: %s", string(output))

			// Collect all piece CIDs
			pieceCIDs := make(map[string]struct{})
			for _, pl := range pieceLists {
				for _, car := range pl.Pieces {
					pieceCIDs[car.PieceCID] = struct{}{}
				}
			}
			require.Equal(t, 2, len(pieceCIDs), "Expected unique piece CID per upstream/source folder")
		})

		t.Run("default behavior", func(t *testing.T) {
			outputDir := filepath.Join(t.TempDir(), "output2")
			require.NoError(t, os.Mkdir(outputDir, 0755))

			cmd := exec.Command(singularityPath, "prep", "create",
				"--name", "test-prep-2",
				"--local-source", folder1,
				"--local-source", folder2,
				"--local-output", outputDir,
			)
			output, err := cmd.CombinedOutput()
			require.NoError(t, err, "command failed: %s", string(output))
			require.Contains(t, string(output), "created successfully")

			// Verify that files were combined (should have 1 or 2 CAR files depending on size)
			files, err := os.ReadDir(outputDir)
			require.NoError(t, err)
			carFiles := 0
			for _, f := range files {
				if filepath.Ext(f.Name()) == ".car" {
					carFiles++
				}
			}
			require.LessOrEqual(t, carFiles, 2, "Expected files to be combined when possible")
		})
	})
}
