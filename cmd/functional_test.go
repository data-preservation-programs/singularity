package cmd

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/storagesystem"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/data-preservation-programs/singularity/util/testutil"
	uio "github.com/ipfs/go-unixfs/io"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

// The bind address needs to be different for different test package so that they don't conflict.
const contentProviderBind = "127.0.0.1:7778"

// TestDeleteAfterExportWithMultipleOutput tests four scenarios:
// 1. The preparation is created with --delete-after-export, which will delete the source storage after export
// 2. The preparation is created with multiple output folder so expect CAR files split into multiple folders.
// 3. The preparation disables inline preparation
// 4. The preparation disables dag
func TestDeleteAfterExportWithMultipleOutput(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		source := t.TempDir()
		output1 := t.TempDir()
		output2 := t.TempDir()
		// create file of different sizes
		sizes := []int{0, 1, 1 << 20, 10 << 20, 30 << 20}
		for _, size := range sizes {
			err := os.WriteFile(filepath.Join(source, fmt.Sprintf("size-%d.txt", size)), testutil.GenerateFixedBytes(size), 0777)
			require.NoError(t, err)
		}

		runner := Runner{mode: Verbose}
		defer runner.Save(t, source, output1, output2)

		// Create prep with both source and output
		_, _, err := runner.Run(ctx, fmt.Sprintf("singularity prep create --no-inline --no-dag --delete-after-export --max-size 3MB --name prep --local-source %s --local-output %s --local-output %s",
			testutil.EscapePath(source), testutil.EscapePath(output1), testutil.EscapePath(output2)))
		require.NoError(t, err)

		// Start scanning
		_, _, err = runner.Run(ctx, "singularity prep start-scan 1 1")
		require.NoError(t, err)

		// run the dataset worker
		_, _, err = runner.Run(ctx, "singularity run dataset-worker --exit-on-complete=true --exit-on-error=true")
		require.NoError(t, err)

		// List Pieces
		_, _, err = runner.Run(ctx, "singularity prep list-pieces 1")
		require.NoError(t, err)

		// Check output1 and output2 to make sure they are have some CAR files
		entries, err := os.ReadDir(output1)
		require.NoError(t, err)
		require.NotEmpty(t, entries)
		entries, err = os.ReadDir(output2)
		require.NoError(t, err)
		require.NotEmpty(t, entries)

		// Check source to make sure all files are removed
		entries, err = os.ReadDir(source)
		require.NoError(t, err)
		require.Empty(t, entries)

		// Check if any carblocks are saved in the database
		var carBlocksCount int64
		err = db.Model(&model.CarBlock{}).Count(&carBlocksCount).Error
		require.NoError(t, err)
		require.Equal(t, int64(0), carBlocksCount)

		// Check if any CID or data has been set for directories
		var dirs []model.Directory
		err = db.Find(&dirs).Error
		require.NoError(t, err)
		for _, dir := range dirs {
			require.Empty(t, dir.CID.String())
			require.Empty(t, dir.Data)
		}

		// Explore should still work even the directory does not have CID
		_, _, err = runner.Run(ctx, "singularity prep explore 1 1")
		require.NoError(t, err)
	})
}

// TestRescan tests the file versioning and new file discovery during a source storage rescan
// 1. create a prep with local source and output
// 2. do data prep
// 3. add new file and override existing file
// 4. do data prep again
// 5. check that both versions of the file
func TestRescan(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		source := t.TempDir()
		output := t.TempDir()
		extract1 := t.TempDir()
		extract2 := t.TempDir()
		err := os.WriteFile(filepath.Join(source, "file1.txt"), []byte("hello file1"), 0777)
		require.NoError(t, err)
		err = os.WriteFile(filepath.Join(source, "file2.txt"), []byte("hello file2"), 0777)
		require.NoError(t, err)
		runner := Runner{mode: Verbose}
		defer runner.Save(t, source, output, extract1, extract2)
		// Create prep with both source and output
		_, _, err = runner.Run(ctx, fmt.Sprintf("singularity storage create local --name source --path %s", testutil.EscapePath(source)))
		require.NoError(t, err)

		_, _, err = runner.Run(ctx, fmt.Sprintf("singularity prep create --source source --local-output %s", testutil.EscapePath(output)))
		require.NoError(t, err)

		// Start scanning
		_, _, err = runner.Run(ctx, "singularity prep start-scan 1 source")
		require.NoError(t, err)

		// run the dataset worker
		_, _, err = runner.Run(ctx, "singularity run dataset-worker --exit-on-complete=true --exit-on-error=true")
		require.NoError(t, err)

		// Explore rootpath
		exploreRootResult, _, err := runner.Run(ctx, "singularity prep explore 1 source")
		require.NoError(t, err)
		rootCID1 := GetFirstCID(exploreRootResult)

		// run the daggen
		_, _, err = runner.Run(ctx, "singularity prep start-daggen 1 source")
		require.NoError(t, err)

		// run the dataset worker again
		_, _, err = runner.Run(ctx, "singularity run dataset-worker --exit-on-complete=true --exit-on-error=true")
		require.NoError(t, err)

		// Make some change to the file system
		err = os.WriteFile(filepath.Join(source, "file2.txt"), []byte("hello file2 modified"), 0777)
		require.NoError(t, err)
		err = os.WriteFile(filepath.Join(source, "file3.txt"), []byte("hello file3"), 0777)
		require.NoError(t, err)

		// Rescan
		_, _, err = runner.Run(ctx, "singularity prep start-scan 1 source")
		require.NoError(t, err)

		// run the dataset worker
		_, _, err = runner.Run(ctx, "singularity run dataset-worker --exit-on-complete=true --exit-on-error=true")
		require.NoError(t, err)

		// Explore rootpath
		exploreRootResult, _, err = runner.Run(ctx, "singularity prep explore 1 source")
		require.NoError(t, err)
		rootCID2 := GetFirstCID(exploreRootResult)

		// run the daggen
		_, _, err = runner.Run(ctx, "singularity prep start-daggen 1 source")
		require.NoError(t, err)

		// run the dataset worker again
		_, _, err = runner.Run(ctx, "singularity run dataset-worker --exit-on-complete=true --exit-on-error=true")
		require.NoError(t, err)

		// Check prepared pieces again
		_, _, err = runner.Run(ctx, "singularity prep list-pieces 1")
		require.NoError(t, err)

		// Extract those pieces to a new folder for the first snapshot
		_, _, err = runner.Run(ctx, fmt.Sprintf("singularity extract-car -i %s -o %s -c %s", testutil.EscapePath(output), testutil.EscapePath(extract1), rootCID1))
		require.NoError(t, err)

		// Extract those pieces to a new folder for the first snapshot
		_, _, err = runner.Run(ctx, fmt.Sprintf("singularity extract-car -i %s -o %s -c %s", testutil.EscapePath(output), testutil.EscapePath(extract2), rootCID2))
		require.NoError(t, err)

		// Check the extract1 folder
		entries, err := os.ReadDir(extract1)
		require.NoError(t, err)
		require.Len(t, entries, 2)
		file2V1, err := os.ReadFile(filepath.Join(extract1, "file2.txt"))
		require.NoError(t, err)
		require.Equal(t, "hello file2", string(file2V1))

		// Check the extract2 folder
		entries, err = os.ReadDir(extract2)
		require.NoError(t, err)
		require.Len(t, entries, 3)
		file2V2, err := os.ReadFile(filepath.Join(extract2, "file2.txt"))
		require.NoError(t, err)
		require.Equal(t, "hello file2 modified", string(file2V2))
	})
}

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
// 6. Download the pieces using the download server
// 7. Download the pieces using the metadata API with Download utility
// 8. Extract into folder and compare with the original source
// 9. Repeat above with different maxSize and inline
func TestDataPrep(t *testing.T) {
	// Prepare local source
	tmp := t.TempDir()
	s3tmp := t.TempDir()

	s3Handler, err := storagesystem.NewRCloneHandler(context.Background(), model.Storage{
		Type: "s3",
		Path: "public-dataset-test",
		Config: map[string]string{d
			"region":     "us-west-2",
			"provider":   "AWS",
			"chunk_size": "5Mi",
			"list_chunk": "1000",
		},
	})
	require.NoError(t, err)
	var entries []storagesystem.Entry
	for entry := range s3Handler.Scan(context.Background(), "") {
		entries = append(entries, entry)
	}
	for _, entry := range entries {
		require.NoError(t, entry.Error)
		if entry.Info == nil {
			continue
		}
		entryPath := entry.Info.Remote()
		destPath := filepath.Join(s3tmp, entryPath)
		err = os.MkdirAll(filepath.Dir(destPath), 0777)
		require.NoError(t, err)
		readCloser, _, err := s3Handler.Read(context.Background(), entryPath, 0, entry.Info.Size())
		require.NoError(t, err)
		defer readCloser.Close()
		content, err := io.ReadAll(readCloser)
		require.NoError(t, err)
		err = os.WriteFile(destPath, content, 0777)
		require.NoError(t, err)
	}

	originalShardingSize := uio.HAMTShardingSize
	uio.HAMTShardingSize = 1024
	defer func() { uio.HAMTShardingSize = originalShardingSize }()

	err = os.MkdirAll(filepath.Join(tmp, "smallfiles"), 0777)
	require.NoError(t, err)
	// create 100 random files
	for i := 0; i < 100; i++ {
		file := filepath.Join(tmp, "smallfiles", fmt.Sprintf("file-%d.txt", i))
		content := testutil.GenerateFixedBytes(i)
		err := os.WriteFile(file, content, 0777)
		require.NoError(t, err)
	}

	// create 10 nested folders
	folderPath := filepath.Join(tmp, "subfolder")
	for i := 0; i < 10; i++ {
		folderPath = filepath.Join(folderPath, fmt.Sprintf("folder-%d", i))
	}
	err = os.MkdirAll(folderPath, 0777)
	require.NoError(t, err)
	err = os.WriteFile(filepath.Join(folderPath, "file.txt"), testutil.GenerateFixedBytes(1000), 0777)
	require.NoError(t, err)

	// create file of different sizes
	sizes := []int{0, 1, 1 << 20, 1<<20 + 1, 20 << 20}
	for _, size := range sizes {
		err = os.WriteFile(filepath.Join(tmp, fmt.Sprintf("size-%d.txt", size)), testutil.GenerateFixedBytes(size), 0777)
		require.NoError(t, err)
	}

	tests := []struct {
		name          string
		maxSize       int
		sourceType    string
		sourcePath    string
		sourceFlags   string
		downloadFlags string
		compare       string
	}{
		{
			name:          "s3-public",
			maxSize:       3 << 20,
			sourceType:    "s3 aws",
			sourcePath:    "public-dataset-test",
			sourceFlags:   "--region us-west-2",
			downloadFlags: "",
			compare:       s3tmp,
		},
		{
			name:        "local-60",
			maxSize:     60 << 20,
			sourceType:  "local",
			sourcePath:  tmp,
			sourceFlags: "",
			compare:     tmp,
		},
		{
			name:        "local-15",
			maxSize:     15 << 20,
			sourceType:  "local",
			sourcePath:  tmp,
			sourceFlags: "",
			compare:     tmp,
		},
	}

	for _, tt := range tests {
		maxSize := tt.maxSize
		pieceSize := util.NextPowerOfTwo(uint64(maxSize))
		t.Run(tt.name, func(t *testing.T) {
			for _, inline := range []bool{false, true} {
				t.Run(fmt.Sprint(inline), func(t *testing.T) {
					testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
						outDir := t.TempDir()
						downloadDir := t.TempDir()
						downloadDir2 := t.TempDir()
						extractDir := t.TempDir()
						runner := Runner{mode: Normal}
						defer runner.Save(t, tmp, outDir, downloadDir, downloadDir2, extractDir)
						// Create source storage
						_, _, err := runner.Run(ctx, fmt.Sprintf("singularity storage create %s %s --name source --path %s", tt.sourceType, tt.sourceFlags, testutil.EscapePath(tt.sourcePath)))
						require.NoError(t, err)

						var outputStorage string
						// Create output storage if not inline
						if !inline {
							_, _, err = runner.Run(ctx, fmt.Sprintf("singularity storage create local --name output --path %s", testutil.EscapePath(outDir)))
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

						// run the dataset worker
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
						_, _, err = runner.Run(ctx, "singularity prep explore 1 source subfolder")
						require.NoError(t, err)

						// run the daggen
						_, _, err = runner.Run(ctx, "singularity prep start-daggen 1 source")
						require.NoError(t, err)

						// run the dataset worker again
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

						// run the content provider
						contentProviderCtx, cancel := context.WithCancel(ctx)
						contentProviderDone := make(chan struct{})
						downloadServerDone := make(chan struct{})
						defer func() { <-contentProviderDone }()
						defer func() { <-downloadServerDone }()
						defer cancel()
						go func() {
							NewRunner().Run(contentProviderCtx, "singularity run content-provider --http-bind "+contentProviderBind)
							close(contentProviderDone)
						}()
						// Wait for content provider to be ready
						err = WaitForServerReady(ctx, fmt.Sprintf("http://%s/health", contentProviderBind))
						require.NoError(t, err)

						go func() {
							NewRunner().Run(contentProviderCtx, "singularity run download-server --metadata-api http://"+contentProviderBind)
							close(downloadServerDone)
						}()
						// Wait for download server to be ready
						err = WaitForServerReady(ctx, fmt.Sprintf("http://%s/health", "127.0.0.1:8888"))
						require.NoError(t, err)

						// Download all pieces from content provider
						for _, pieceCID := range pieceCIDs {
							downloaded, err := Download(ctx, fmt.Sprintf("http://%s/piece/%s", contentProviderBind, pieceCID), 10)
							require.NoError(t, err)
							calculatedPieceCID := CalculateCommp(t, downloaded, pieceSize)
							require.Equal(t, pieceCID, calculatedPieceCID)
							err = os.WriteFile(filepath.Join(downloadDir, pieceCID+".car"), downloaded, 0777)
							require.NoError(t, err)
						}

						// Download all pieces using local download server
						for _, pieceCID := range pieceCIDs {
							downloaded, err := Download(ctx, fmt.Sprintf("http://%s/piece/%s", "127.0.0.1:8888", pieceCID), 10)
							require.NoError(t, err)
							calculatedPieceCID := CalculateCommp(t, downloaded, pieceSize)
							require.Equal(t, pieceCID, calculatedPieceCID)
						}

						// Download all pieces using download CLI
						for _, pieceCID := range pieceCIDs {
							_, _, err = runner.Run(ctx, fmt.Sprintf("singularity download %s --quiet --api http://%s --out-dir %s %s", tt.downloadFlags, contentProviderBind, testutil.EscapePath(downloadDir2), pieceCID))
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
						CompareDirectories(t, tt.compare, extractDir)
					})
				})
			}
		})
	}
}
