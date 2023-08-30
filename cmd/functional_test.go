package cmd

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/util"
	"github.com/data-preservation-programs/singularity/util/testutil"
	uio "github.com/ipfs/go-unixfs/io"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

// The bind address needs to be different for different test package so that they don't conflict
const contentProviderBind = "127.0.0.1:7778"

// TestPrepCreateWithLocalSource tests the following scenario:
// 1. Create a local source with a few files
//   - file of different sizes
//   - nested folders
//   - folder containing lots of files
//
// 2. Create a local output
// 3. Create a preparation with the local source and output
// 4. Start scanning, packing and daggen
// 5. Download the pieces using the piece API
// 6. Download the pieces using the metadata API with Download utility
// 7. Extract into folder and compare with the original source
// 8. Repeat above with different maxSize and inline
func TestDataPrepWithLocalSource(t *testing.T) {
	// Prepare local source
	tmp := t.TempDir()
	originalShardingSize := uio.HAMTShardingSize
	uio.HAMTShardingSize = 1024
	defer func() { uio.HAMTShardingSize = originalShardingSize }()

	// create 100 random files
	for i := 0; i < 100; i++ {
		file := filepath.Join(tmp, fmt.Sprintf("file-%d.txt", i))
		content := testutil.GenerateFixedBytes(i)
		err := os.WriteFile(file, content, 0777)
		require.NoError(t, err)
	}

	// create 10 nested folders
	folderPath := tmp
	for i := 0; i < 10; i++ {
		folderPath = filepath.Join(folderPath, fmt.Sprintf("folder-%d", i))
	}
	err := os.MkdirAll(folderPath, 0777)
	require.NoError(t, err)
	err = os.WriteFile(filepath.Join(folderPath, "file.txt"), testutil.GenerateFixedBytes(1000), 0777)

	// create file of different sizes
	sizes := []int{0, 1, 1 << 20, 1<<20 + 1, 20 << 20}
	for _, size := range sizes {
		err = os.WriteFile(filepath.Join(tmp, fmt.Sprintf("size-%d.txt", size)), testutil.GenerateFixedBytes(size), 0777)
	}

	// maxSizes := []int{60 << 20, 15 << 20, 3 << 20}
	maxSizes := []int{60 << 20}
	for _, maxSize := range maxSizes {
		pieceSize := util.NextPowerOfTwo(uint64(maxSize))
		t.Run(fmt.Sprint(maxSize), func(t *testing.T) {
			for _, inline := range []bool{true, false} {
				t.Run(fmt.Sprint(inline), func(t *testing.T) {
					for _, mode := range []RunnerMode{Normal, Verbose, JSON} {
						t.Run(fmt.Sprint(mode), func(t *testing.T) {
							testutil.One(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
								outDir := t.TempDir()
								downloadDir := t.TempDir()
								downloadDir2 := t.TempDir()
								extractDir := t.TempDir()
								runner := Runner{mode: mode}
								defer runner.Save(t, tmp, outDir, downloadDir, downloadDir2, extractDir)
								// Create source storage
								_, _, err := runner.Run(ctx, fmt.Sprintf("singularity storage create local source %s", testutil.EscapePath(tmp)))
								require.NoError(t, err)

								var outputStorage string
								// Create output storage if not inline
								if !inline {
									_, _, err = runner.Run(ctx, fmt.Sprintf("singularity storage create local output %s", testutil.EscapePath(outDir)))
									require.NoError(t, err)
									outputStorage = " --output output"
								}

								// Create preparation
								_, _, err = runner.Run(ctx, fmt.Sprintf("singularity prep create --max-size %d --source source%s", maxSize, outputStorage))
								require.NoError(t, err)

								// List all preparations
								_, _, err = runner.Run(ctx, "singularity prep list")
								require.NoError(t, err)

								// Enable scanning
								_, _, err = runner.Run(ctx, "singularity prep start-scan 1 source")
								require.NoError(t, err)

								// Run the dataset worker
								_, _, err = runner.Run(ctx, "singularity run dataset-worker --exit-on-complete=true --exit-on-error=true")
								require.NoError(t, err)

								// Check preparation status
								_, _, err = runner.Run(ctx, "singularity prep status 1")
								require.NoError(t, err)

								// Check prepared pieces
								_, _, err = runner.Run(ctx, "singularity prep list-pieces 1")
								require.NoError(t, err)

								// Explore rootpath
								exploreRootResult, _, err := runner.Run(ctx, "singularity prep explore 1 source")
								require.NoError(t, err)
								rootCID := GetFirstCID(exploreRootResult)

								// Explore subpath
								_, _, err = runner.Run(ctx, "singularity prep explore 1 source folder-0")
								require.NoError(t, err)

								// Run the daggen
								_, _, err = runner.Run(ctx, "singularity prep start-daggen 1 source")
								require.NoError(t, err)

								// Run the dataset worker again
								_, _, err = runner.Run(ctx, "singularity run dataset-worker --exit-on-complete=true --exit-on-error=true")
								require.NoError(t, err)

								// Check preparation status again
								_, _, err = runner.Run(ctx, "singularity prep status 1")
								require.NoError(t, err)

								// Check prepared pieces again
								listPiecesStdout, _, err := runner.Run(ctx, "singularity prep list-pieces 1")
								require.NoError(t, err)

								// Get all PieceCIDs from stdout
								pieceCIDs := GetAllPieceCIDs(listPiecesStdout)

								// Run the content provider
								contentProviderCtx, cancel := context.WithTimeout(ctx, time.Minute)
								contentProviderDone := make(chan struct{})
								defer func() { <-contentProviderDone }()
								defer cancel()
								go func() {
									Run(contentProviderCtx, "singularity run content-provider --http-bind "+contentProviderBind)
									close(contentProviderDone)
								}()

								// Wait for content provider to be ready
								err = WaitForServerReady(ctx, fmt.Sprintf("http://%s/health", contentProviderBind))
								require.NoError(t, err)

								// Download all pieces from content provider
								for _, pieceCID := range pieceCIDs {
									downloaded, err := Download(ctx, fmt.Sprintf("http://%s/piece/%s", contentProviderBind, pieceCID), 4)
									require.NoError(t, err)
									calculatedPieceCID := CalculateCommp(t, downloaded, pieceSize)
									require.Equal(t, pieceCID, calculatedPieceCID)
									err = os.WriteFile(filepath.Join(downloadDir, pieceCID+".car"), downloaded, 0777)
									require.NoError(t, err)
								}

								// Download all pieces using download CLI
								for _, pieceCID := range pieceCIDs {
									_, _, err = runner.Run(ctx, fmt.Sprintf("singularity download --quiet --api http://%s --out-dir %s %s", contentProviderBind, testutil.EscapePath(downloadDir2), pieceCID))
									require.NoError(t, err)
									downloaded, err := os.ReadFile(filepath.Join(downloadDir2, pieceCID+".car"))
									require.NoError(t, err)
									calculatedPieceCID := CalculateCommp(t, downloaded, pieceSize)
									require.Equal(t, pieceCID, calculatedPieceCID)
								}

								// Extract those pieces to a new folder
								_, _, err = runner.Run(ctx, fmt.Sprintf("singularity extract-car -i %s -o %s -c %s", testutil.EscapePath(downloadDir), testutil.EscapePath(extractDir), rootCID))
								require.NoError(t, err)

								// Compare the extracted folder with the original folder
								CompareDirectories(t, tmp, extractDir)
							})
						})
					}
				})
			}
		})
	}
}
