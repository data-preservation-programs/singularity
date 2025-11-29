package cmd

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

// downloadTestBindBase is used to generate unique ports for each test
const downloadTestBindBase = 7779

// TestDownloadPaddedPieces tests downloading padded pieces in three scenarios:
// 1. non-inline + enable-http-piece: content-provider serves CAR file directly from disk
// 2. inline + enable-http-piece: content-provider assembles piece on-the-fly
// 3. inline + metadata-only: downloader assembles piece locally from metadata
func TestDownloadPaddedPieces(t *testing.T) {
	tests := []struct {
		name                string
		inline              bool
		enableHTTPPiece     bool
		minPieceSize        int64
		expectedPaddingInfo string
	}{
		{
			name:                "non-inline with padding",
			inline:              false,
			enableHTTPPiece:     true,
			minPieceSize:        1 << 20, // 1 MiB - will force padding
			expectedPaddingInfo: "literal zeros in CAR file",
		},
		{
			name:                "inline with piece assembly",
			inline:              true,
			enableHTTPPiece:     true,
			minPieceSize:        1 << 20, // 1 MiB - will force padding
			expectedPaddingInfo: "PieceReader serves zeros virtually",
		},
		{
			name:                "inline metadata-only",
			inline:              true,
			enableHTTPPiece:     false,
			minPieceSize:        1 << 20, // 1 MiB - will force padding
			expectedPaddingInfo: "downloader assembles with PieceReader",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
				source := t.TempDir()
				// Create a test file that's smaller than minPieceSize to force padding
				err := os.WriteFile(filepath.Join(source, "test.txt"), testutil.GenerateFixedBytes(100<<10), 0644) // 100 KiB file
				require.NoError(t, err)

				output := t.TempDir()
				downloadDir := t.TempDir()

				// Use unique port for each test
				downloadTestBind := fmt.Sprintf("127.0.0.1:%d", downloadTestBindBase+i)

				runner := Runner{mode: Normal}
				defer runner.Save(t, source, output, downloadDir)

				// Create source storage
				_, _, err = runner.Run(ctx, fmt.Sprintf("singularity storage create local --name source --path %s", testutil.EscapePath(source)))
				require.NoError(t, err)

				var outputStorageFlag string
				if !tc.inline {
					// For non-inline mode, use --local-output
					outputStorageFlag = fmt.Sprintf(" --local-output %s", testutil.EscapePath(output))
				}

				// Create preparation with minPieceSize to force padding
				prepName := "test"
				inlineFlag := ""
				if !tc.inline {
					// Only add flag for non-inline mode (inline is default)
					inlineFlag = " --no-inline"
				}
				_, _, err = runner.Run(ctx, fmt.Sprintf("singularity prep create --name %s --max-size 2MB --min-piece-size %d%s --source source%s",
					prepName, tc.minPieceSize, inlineFlag, outputStorageFlag))
				require.NoError(t, err)

				// Start scanning
				_, _, err = runner.Run(ctx, fmt.Sprintf("singularity prep start-scan %s source", prepName))
				require.NoError(t, err)

				// Run dataset worker to pack
				_, _, err = runner.Run(ctx, "singularity run dataset-worker --exit-on-complete=true --exit-on-error=true")
				require.NoError(t, err)

				// List pieces
				listPiecesOut, _, err := runner.Run(ctx, fmt.Sprintf("singularity prep list-pieces %s", prepName))
				require.NoError(t, err)
				pieceCIDs := GetAllPieceCIDs(listPiecesOut)
				require.NotEmpty(t, pieceCIDs, "should have at least one piece")

				t.Logf("Created %d piece(s) with %s", len(pieceCIDs), tc.expectedPaddingInfo)

				// Verify the piece is padded by checking database
				var cars []model.Car
				err = db.Find(&cars).Error
				require.NoError(t, err)
				require.NotEmpty(t, cars)

				// Find the car that matches our expected criteria
				// FileSize should be (127/128) × PieceSize due to Fr32 padding overhead
				expectedFileSize := (tc.minPieceSize * 127) / 128
				var paddedCar *model.Car
				for _, car := range cars {
					if car.PieceSize == tc.minPieceSize && car.FileSize == expectedFileSize {
						carCopy := car
						paddedCar = &carCopy
						t.Logf("Found padded CAR: PieceCID=%s, PieceSize=%d, FileSize=%d (expected %d), MinPieceSizePadding=%d",
							car.PieceCID.String(), car.PieceSize, car.FileSize, expectedFileSize, car.MinPieceSizePadding)

						if tc.inline {
							require.Greater(t, car.MinPieceSizePadding, int64(0), "inline mode should have virtual padding")
						} else {
							require.Equal(t, int64(0), car.MinPieceSizePadding, "non-inline mode should have zeros in file")
						}
						break
					}
				}
				require.NotNil(t, paddedCar, "should have found a padded CAR")

				// For non-inline mode, verify the CAR file on disk has correct size
				if !tc.inline {
					carPath := filepath.Join(output, paddedCar.PieceCID.String()+".car")
					fileInfo, err := os.Stat(carPath)
					require.NoError(t, err, "CAR file should exist on disk")
					t.Logf("CAR file on disk: path=%s, size=%d (should equal piece size %d for Curio TreeD)",
						carPath, fileInfo.Size(), paddedCar.PieceSize)
					require.Equal(t, paddedCar.FileSize, fileInfo.Size(), "CAR file should be padded to piece size")

					// CommP was calculated before padding, so we don't verify it here
					// The important verification is that downloaded piece matches (tested below)
				}

				// Start content-provider with appropriate flags
				contentProviderCtx, cancel := context.WithCancel(ctx)
				defer cancel()
				contentProviderDone := make(chan struct{})
				defer func() { <-contentProviderDone }()

				httpPieceFlag := ""
				if !tc.enableHTTPPiece {
					httpPieceFlag = " --enable-http-piece=false"
				}

				go func() {
					NewRunner().Run(contentProviderCtx, fmt.Sprintf("singularity run content-provider --http-bind %s%s", downloadTestBind, httpPieceFlag))
					close(contentProviderDone)
				}()

				// Wait for content-provider to be ready
				err = WaitForServerReady(ctx, fmt.Sprintf("http://%s/health", downloadTestBind))
				require.NoError(t, err)

				// Download all pieces
				for _, pieceCID := range pieceCIDs {
					t.Logf("Downloading %s...", pieceCID)

					var downloaded []byte
					if tc.enableHTTPPiece {
						// For scenarios with enable-http-piece, download directly from /piece endpoint
						// (content-provider serves the piece directly)
						downloaded, err = Download(ctx, fmt.Sprintf("http://%s/piece/%s", downloadTestBind, pieceCID), 10)
						require.NoError(t, err)
					} else {
						// For metadata-only mode, use download command (downloader assembles locally)
						_, _, err = runner.Run(ctx, fmt.Sprintf("singularity download --quiet --api http://%s --out-dir %s %s",
							downloadTestBind, testutil.EscapePath(downloadDir), pieceCID))
						require.NoError(t, err)

						// Read downloaded file
						downloaded, err = os.ReadFile(filepath.Join(downloadDir, pieceCID+".car"))
						require.NoError(t, err)
					}

					require.NotEmpty(t, downloaded)

					// Verify the downloaded piece is the expected size (127/128 × pieceSize)
					require.Equal(t, int(expectedFileSize), len(downloaded), "downloaded piece should be padded to (127/128)×pieceSize")

					// Verify CommP matches - use the actual piece size from database
					calculatedPieceCID := CalculateCommp(t, downloaded, uint64(paddedCar.PieceSize))
					if pieceCID != calculatedPieceCID {
						t.Logf("CommP mismatch: expected=%s, calculated=%s, downloaded_size=%d, piece_size=%d",
							pieceCID, calculatedPieceCID, len(downloaded), paddedCar.PieceSize)
					}
					require.Equal(t, pieceCID, calculatedPieceCID, "CommP should match")

					// Verify file ends with zeros (padding)
					foundNonZero := false
					paddingStart := len(downloaded) - 1
					for paddingStart > 0 && downloaded[paddingStart] == 0 {
						paddingStart--
					}
					paddingStart++ // Move to first zero
					require.Greater(t, len(downloaded)-paddingStart, 0, "should have trailing zero padding")

					// Verify everything before padding region is not all zeros
					for i := 0; i < paddingStart; i++ {
						if downloaded[i] != 0 {
							foundNonZero = true
							break
						}
					}
					require.True(t, foundNonZero, "should have non-zero data before padding")

					t.Logf("Successfully downloaded and verified %s: %d bytes with %d bytes of padding",
						pieceCID, len(downloaded), len(downloaded)-paddingStart)
				}

				cancel() // Stop content-provider
			})
		})
	}
}
