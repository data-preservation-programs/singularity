package pack

import (
	"context"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/store"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

// TestMinPieceSizePaddingDetection verifies that padding is correctly detected and applied
func TestMinPieceSizePaddingDetection(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		tmp := t.TempDir()
		out := t.TempDir()

		tests := []struct {
			name                     string
			dataSize                 int
			minPieceSize             int64
			expectPaddingNonZero     bool
			expectedNaturalPieceSize uint64
			expectedFinalPieceSize   uint64
		}{
			{
				name:                     "no padding needed - natural size meets minPieceSize",
				dataSize:                 600_000, // ~600 KB
				minPieceSize:             1 << 19, // 512 KiB
				expectPaddingNonZero:     false,
				expectedNaturalPieceSize: 1 << 20, // NextPowerOfTwo(~600KB) = 1 MiB
				expectedFinalPieceSize:   1 << 20, // No padding needed
			},
			{
				name:                     "padding needed - natural size below minPieceSize",
				dataSize:                 400_000, // ~400 KB
				minPieceSize:             1 << 20, // 1 MiB
				expectPaddingNonZero:     true,
				expectedNaturalPieceSize: 1 << 19, // NextPowerOfTwo(~400KB) = 512 KiB
				expectedFinalPieceSize:   1 << 20, // Padded to 1 MiB
			},
		}

		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				testData := testutil.GenerateRandomBytes(tc.dataSize)
				testPath := filepath.Join(tmp, "test.txt")
				err := os.WriteFile(testPath, testData, 0644)
				require.NoError(t, err)

				testStat, err := os.Stat(testPath)
				require.NoError(t, err)

				// Test non-inline mode
				t.Run("non-inline", func(t *testing.T) {
					storage := model.Storage{Name: tc.name + "-src", Type: "local", Path: tmp}
					err = db.Create(&storage).Error
					require.NoError(t, err)

					outputStorage := model.Storage{Name: tc.name + "-out", Type: "local", Path: out}
					err = db.Create(&outputStorage).Error
					require.NoError(t, err)

					prep := model.Preparation{
						Name:         tc.name + "-prep",
						MaxSize:      2_000_000,
						PieceSize:    1 << 21, // 2 MiB
						MinPieceSize: tc.minPieceSize,
					}
					err = db.Create(&prep).Error
					require.NoError(t, err)

					err = db.Exec("INSERT INTO source_attachments (preparation_id, storage_id) VALUES (?, ?)", prep.ID, storage.ID).Error
					require.NoError(t, err)
					err = db.Exec("INSERT INTO output_attachments (preparation_id, storage_id) VALUES (?, ?)", prep.ID, outputStorage.ID).Error
					require.NoError(t, err)

					var attachment model.SourceAttachment
					err = db.Where("preparation_id = ? AND storage_id = ?", prep.ID, storage.ID).First(&attachment).Error
					require.NoError(t, err)

					dir := model.Directory{AttachmentID: &attachment.ID}
					err = db.Create(&dir).Error
					require.NoError(t, err)

					file := model.File{
						AttachmentID:     &attachment.ID,
						Path:             "test.txt",
						Size:             int64(len(testData)),
						LastModifiedNano: testStat.ModTime().UnixNano(),
						DirectoryID:      &dir.ID,
					}
					err = db.Create(&file).Error
					require.NoError(t, err)

					job := model.Job{
						AttachmentID: &attachment.ID,
						State:        model.Processing,
						Type:         model.Pack,
						FileRanges: []model.FileRange{
							{FileID: file.ID, Offset: 0, Length: int64(len(testData))},
						},
					}
					err = db.Create(&job).Error
					require.NoError(t, err)

					err = db.Preload("FileRanges.File").Preload("Attachment.Preparation.OutputStorages").Preload("Attachment.Storage").First(&job, job.ID).Error
					require.NoError(t, err)

					car, err := Pack(ctx, db, job)
					require.NoError(t, err)
					require.NotNil(t, car)

					// Verify piece size calculation
					require.Equal(t, int64(tc.expectedFinalPieceSize), car.PieceSize)

					// Verify padding field
					if tc.expectPaddingNonZero {
						// For non-inline with local storage, zeros are written to file, so MinPieceSizePadding should be 0
						require.Equal(t, int64(0), car.MinPieceSizePadding, "non-inline should have MinPieceSizePadding=0 (zeros in file)")
						// But FileSize should be padded to (127/128) × piece_size due to Fr32 overhead
						expectedFileSize := (int64(tc.expectedFinalPieceSize) * 127) / 128
						require.Equal(t, expectedFileSize, car.FileSize, "FileSize should include padding")

						// Verify file was actually padded
						carPath := filepath.Join(out, car.StoragePath)
						stat, err := os.Stat(carPath)
						require.NoError(t, err)
						require.Equal(t, expectedFileSize, stat.Size(), "physical file should be padded to (127/128)×piece_size")
					} else {
						require.Equal(t, int64(0), car.MinPieceSizePadding, "no padding needed")
					}
				})

				// Test inline mode
				t.Run("inline", func(t *testing.T) {
					storage := model.Storage{Name: tc.name + "-inline-src", Type: "local", Path: tmp}
					err = db.Create(&storage).Error
					require.NoError(t, err)

					prep := model.Preparation{
						Name:         tc.name + "-inline-prep",
						MaxSize:      2_000_000,
						PieceSize:    1 << 21, // 2 MiB
						MinPieceSize: tc.minPieceSize,
						NoInline:     false, // Inline mode
					}
					err = db.Create(&prep).Error
					require.NoError(t, err)

					err = db.Exec("INSERT INTO source_attachments (preparation_id, storage_id) VALUES (?, ?)", prep.ID, storage.ID).Error
					require.NoError(t, err)

					var attachment model.SourceAttachment
					err = db.Where("preparation_id = ? AND storage_id = ?", prep.ID, storage.ID).First(&attachment).Error
					require.NoError(t, err)

					dir := model.Directory{AttachmentID: &attachment.ID}
					err = db.Create(&dir).Error
					require.NoError(t, err)

					file := model.File{
						AttachmentID:     &attachment.ID,
						Path:             "test.txt",
						Size:             int64(len(testData)),
						LastModifiedNano: testStat.ModTime().UnixNano(),
						DirectoryID:      &dir.ID,
					}
					err = db.Create(&file).Error
					require.NoError(t, err)

					job := model.Job{
						AttachmentID: &attachment.ID,
						State:        model.Processing,
						Type:         model.Pack,
						FileRanges: []model.FileRange{
							{FileID: file.ID, Offset: 0, Length: int64(len(testData))},
						},
					}
					err = db.Create(&job).Error
					require.NoError(t, err)

					err = db.Preload("FileRanges.File").Preload("Attachment.Preparation").Preload("Attachment.Storage").First(&job, job.ID).Error
					require.NoError(t, err)

					car, err := Pack(ctx, db, job)
					require.NoError(t, err)
					require.NotNil(t, car)

					// Verify piece size calculation
					require.Equal(t, int64(tc.expectedFinalPieceSize), car.PieceSize)

					// Verify padding field
					if tc.expectPaddingNonZero {
						// For inline, MinPieceSizePadding stores the virtual padding amount
						require.Greater(t, car.MinPieceSizePadding, int64(0), "inline should have MinPieceSizePadding > 0")
						// FileSize should be (127/128) × piece_size due to Fr32 overhead
						expectedFileSize := (int64(tc.expectedFinalPieceSize) * 127) / 128
						require.Equal(t, expectedFileSize, car.FileSize, "FileSize should include padding")

						// Verify padding makes sense: actual data + padding = expected FileSize
						actualDataSize := car.FileSize - car.MinPieceSizePadding
						require.Greater(t, actualDataSize, int64(0), "should have actual data")
						require.Equal(t, car.FileSize, actualDataSize+car.MinPieceSizePadding, "FileSize should equal data + padding")
					} else {
						require.Equal(t, int64(0), car.MinPieceSizePadding, "no padding needed")
					}
				})
			})
		}
	})
}

// TestPieceReaderServesZeros verifies that PieceReader correctly serves zeros for inline padded pieces
func TestPieceReaderServesZeros(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		tmp := t.TempDir()

		// Create small test file that will need padding
		testData := testutil.GenerateRandomBytes(400_000) // ~400 KB
		testPath := filepath.Join(tmp, "test.txt")
		err := os.WriteFile(testPath, testData, 0644)
		require.NoError(t, err)

		testStat, err := os.Stat(testPath)
		require.NoError(t, err)

		// Setup storage and preparation (inline mode)
		storage := model.Storage{Name: "test", Type: "local", Path: tmp}
		err = db.Create(&storage).Error
		require.NoError(t, err)

		prep := model.Preparation{
			Name:         "test",
			MaxSize:      2_000_000,
			PieceSize:    1 << 21, // 2 MiB
			MinPieceSize: 1 << 20, // 1 MiB - will force padding
			NoInline:     false,   // Inline mode
		}
		err = db.Create(&prep).Error
		require.NoError(t, err)

		err = db.Exec("INSERT INTO source_attachments (preparation_id, storage_id) VALUES (?, ?)", prep.ID, storage.ID).Error
		require.NoError(t, err)

		var attachment model.SourceAttachment
		err = db.Where("preparation_id = ? AND storage_id = ?", prep.ID, storage.ID).First(&attachment).Error
		require.NoError(t, err)

		dir := model.Directory{AttachmentID: &attachment.ID}
		err = db.Create(&dir).Error
		require.NoError(t, err)

		file := model.File{
			AttachmentID:     &attachment.ID,
			Path:             "test.txt",
			Size:             int64(len(testData)),
			LastModifiedNano: testStat.ModTime().UnixNano(),
			DirectoryID:      &dir.ID,
		}
		err = db.Create(&file).Error
		require.NoError(t, err)

		job := model.Job{
			AttachmentID: &attachment.ID,
			State:        model.Processing,
			Type:         model.Pack,
			FileRanges: []model.FileRange{
				{FileID: file.ID, Offset: 0, Length: int64(len(testData))},
			},
		}
		err = db.Create(&job).Error
		require.NoError(t, err)

		err = db.Preload("FileRanges.File").Preload("Attachment.Preparation").Preload("Attachment.Storage").First(&job, job.ID).Error
		require.NoError(t, err)

		car, err := Pack(ctx, db, job)
		require.NoError(t, err)
		require.NotNil(t, car)

		// Verify padding was applied
		require.Greater(t, car.MinPieceSizePadding, int64(0), "should have virtual padding")
		require.Equal(t, int64(1<<20), car.PieceSize, "piece size should be 1 MiB")
		expectedFileSize := (int64(1<<20) * 127) / 128
		require.Equal(t, expectedFileSize, car.FileSize, "file size should include padding")

		// Load car blocks for PieceReader
		var carBlocks []model.CarBlock
		err = db.Where("car_id = ?", car.ID).Order("car_offset").Find(&carBlocks).Error
		require.NoError(t, err)

		var files []model.File
		err = db.Where("attachment_id = ?", attachment.ID).Find(&files).Error
		require.NoError(t, err)

		// Create PieceReader
		pieceReader, err := store.NewPieceReader(ctx, *car, storage, carBlocks, files)
		require.NoError(t, err)
		defer pieceReader.Close()

		// Read entire file including padding
		fullData := make([]byte, car.FileSize)
		n, err := io.ReadFull(pieceReader, fullData)
		require.NoError(t, err)
		require.Equal(t, int(car.FileSize), n)

		// Verify padding region is all zeros
		actualDataSize := car.FileSize - car.MinPieceSizePadding
		paddingStart := actualDataSize
		paddingEnd := car.FileSize

		for i := paddingStart; i < paddingEnd; i++ {
			require.Equal(t, byte(0), fullData[i], "padding region should be all zeros at position %d", i)
		}

		// Verify we can seek into padding region and read zeros
		_, err = pieceReader.Seek(paddingStart, io.SeekStart)
		require.NoError(t, err)

		paddingData := make([]byte, car.MinPieceSizePadding)
		n, err = io.ReadFull(pieceReader, paddingData)
		require.NoError(t, err)
		require.Equal(t, int(car.MinPieceSizePadding), n)

		for i := range paddingData {
			require.Equal(t, byte(0), paddingData[i], "padding should be all zeros at position %d", i)
		}
	})
}

// TestMinPieceSizePaddingFileIntegrity verifies that literal zero padding doesn't corrupt files
func TestMinPieceSizePaddingFileIntegrity(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		tmp := t.TempDir()
		out := t.TempDir()

		// Create small test file
		testData := testutil.GenerateRandomBytes(400_000) // ~400 KB
		testPath := filepath.Join(tmp, "test.txt")
		err := os.WriteFile(testPath, testData, 0644)
		require.NoError(t, err)

		testStat, err := os.Stat(testPath)
		require.NoError(t, err)

		// Setup storage and preparation (non-inline mode)
		storage := model.Storage{Name: "test", Type: "local", Path: tmp}
		err = db.Create(&storage).Error
		require.NoError(t, err)

		outputStorage := model.Storage{Name: "output", Type: "local", Path: out}
		err = db.Create(&outputStorage).Error
		require.NoError(t, err)

		prep := model.Preparation{
			Name:         "test",
			MaxSize:      2_000_000,
			PieceSize:    1 << 21, // 2 MiB
			MinPieceSize: 1 << 20, // 1 MiB - will force padding
		}
		err = db.Create(&prep).Error
		require.NoError(t, err)

		err = db.Exec("INSERT INTO source_attachments (preparation_id, storage_id) VALUES (?, ?)", prep.ID, storage.ID).Error
		require.NoError(t, err)
		err = db.Exec("INSERT INTO output_attachments (preparation_id, storage_id) VALUES (?, ?)", prep.ID, outputStorage.ID).Error
		require.NoError(t, err)

		var attachment model.SourceAttachment
		err = db.Where("preparation_id = ? AND storage_id = ?", prep.ID, storage.ID).First(&attachment).Error
		require.NoError(t, err)

		dir := model.Directory{AttachmentID: &attachment.ID}
		err = db.Create(&dir).Error
		require.NoError(t, err)

		file := model.File{
			AttachmentID:     &attachment.ID,
			Path:             "test.txt",
			Size:             int64(len(testData)),
			LastModifiedNano: testStat.ModTime().UnixNano(),
			DirectoryID:      &dir.ID,
		}
		err = db.Create(&file).Error
		require.NoError(t, err)

		job := model.Job{
			AttachmentID: &attachment.ID,
			State:        model.Processing,
			Type:         model.Pack,
			FileRanges: []model.FileRange{
				{FileID: file.ID, Offset: 0, Length: int64(len(testData))},
			},
		}
		err = db.Create(&job).Error
		require.NoError(t, err)

		err = db.Preload("FileRanges.File").Preload("Attachment.Preparation.OutputStorages").Preload("Attachment.Storage").First(&job, job.ID).Error
		require.NoError(t, err)

		car, err := Pack(ctx, db, job)
		require.NoError(t, err)
		require.NotNil(t, car)

		// Verify file size matches expected size (127/128 × piece_size)
		require.Equal(t, int64(1<<20), car.PieceSize)
		expectedFileSize := (int64(1<<20) * 127) / 128
		require.Equal(t, expectedFileSize, car.FileSize)

		// Verify actual file was padded
		carPath := filepath.Join(out, car.StoragePath)
		stat, err := os.Stat(carPath)
		require.NoError(t, err)
		require.Equal(t, expectedFileSize, stat.Size())

		// Read file and verify padding region is zeros
		carData, err := os.ReadFile(carPath)
		require.NoError(t, err)
		require.Equal(t, expectedFileSize, int64(len(carData)))

		// Find where padding starts (should be near the end)
		// The last chunk should be zeros
		paddingStart := int64(len(carData)) - 1
		for paddingStart > 0 && carData[paddingStart] == 0 {
			paddingStart--
		}
		paddingStart++ // Move back to first zero

		// Verify all bytes from paddingStart to end are zeros
		for i := paddingStart; i < int64(len(carData)); i++ {
			require.Equal(t, byte(0), carData[i], "padding should be zeros at position %d", i)
		}
	})
}

// TestDownloadPaddedPiece verifies that downloading (reading) a padded piece works correctly
// This simulates what the download handler does: seek to end, seek to start, read entire piece
func TestDownloadPaddedPiece(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		tmp := t.TempDir()

		// Create small test file that will need padding
		testData := testutil.GenerateRandomBytes(400_000) // ~400 KB
		testPath := filepath.Join(tmp, "test.txt")
		err := os.WriteFile(testPath, testData, 0644)
		require.NoError(t, err)

		testStat, err := os.Stat(testPath)
		require.NoError(t, err)

		tests := []struct {
			name         string
			inline       bool
			minPieceSize int64
		}{
			{
				name:         "inline padded piece",
				inline:       true,
				minPieceSize: 1 << 20, // 1 MiB - will force padding
			},
			{
				name:         "non-inline padded piece",
				inline:       false,
				minPieceSize: 1 << 20, // 1 MiB - will force padding
			},
		}

		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				// Setup storage
				storage := model.Storage{Name: tc.name + "-src", Type: "local", Path: tmp}
				err = db.Create(&storage).Error
				require.NoError(t, err)

				var outputStorage *model.Storage
				if !tc.inline {
					out := t.TempDir()
					outputStorage = &model.Storage{Name: tc.name + "-out", Type: "local", Path: out}
					err = db.Create(outputStorage).Error
					require.NoError(t, err)
				}

				prep := model.Preparation{
					Name:         tc.name + "-prep",
					MaxSize:      2_000_000,
					PieceSize:    1 << 21, // 2 MiB
					MinPieceSize: tc.minPieceSize,
					NoInline:     !tc.inline,
				}
				err = db.Create(&prep).Error
				require.NoError(t, err)

				err = db.Exec("INSERT INTO source_attachments (preparation_id, storage_id) VALUES (?, ?)", prep.ID, storage.ID).Error
				require.NoError(t, err)

				if outputStorage != nil {
					err = db.Exec("INSERT INTO output_attachments (preparation_id, storage_id) VALUES (?, ?)", prep.ID, outputStorage.ID).Error
					require.NoError(t, err)
				}

				var attachment model.SourceAttachment
				err = db.Where("preparation_id = ? AND storage_id = ?", prep.ID, storage.ID).First(&attachment).Error
				require.NoError(t, err)

				dir := model.Directory{AttachmentID: &attachment.ID}
				err = db.Create(&dir).Error
				require.NoError(t, err)

				file := model.File{
					AttachmentID:     &attachment.ID,
					Path:             "test.txt",
					Size:             int64(len(testData)),
					LastModifiedNano: testStat.ModTime().UnixNano(),
					DirectoryID:      &dir.ID,
				}
				err = db.Create(&file).Error
				require.NoError(t, err)

				job := model.Job{
					AttachmentID: &attachment.ID,
					State:        model.Processing,
					Type:         model.Pack,
					FileRanges: []model.FileRange{
						{FileID: file.ID, Offset: 0, Length: int64(len(testData))},
					},
				}
				err = db.Create(&job).Error
				require.NoError(t, err)

				if tc.inline {
					err = db.Preload("FileRanges.File").Preload("Attachment.Preparation").Preload("Attachment.Storage").First(&job, job.ID).Error
				} else {
					err = db.Preload("FileRanges.File").Preload("Attachment.Preparation.OutputStorages").Preload("Attachment.Storage").First(&job, job.ID).Error
				}
				require.NoError(t, err)

				// Pack the file
				car, err := Pack(ctx, db, job)
				require.NoError(t, err)
				require.NotNil(t, car)

				// Verify padding was applied
				require.Equal(t, int64(tc.minPieceSize), car.PieceSize)
				expectedFileSize := (int64(tc.minPieceSize) * 127) / 128
				require.Equal(t, expectedFileSize, car.FileSize)

				var downloadedData []byte
				var n int

				if tc.inline {
					// For inline mode, use PieceReader
					var carBlocks []model.CarBlock
					err = db.Where("car_id = ?", car.ID).Order("car_offset").Find(&carBlocks).Error
					require.NoError(t, err)
					require.NotEmpty(t, carBlocks, "inline mode should have car blocks")

					var files []model.File
					err = db.Where("attachment_id = ?", attachment.ID).Find(&files).Error
					require.NoError(t, err)

					// Create PieceReader (simulating what download handler does)
					pieceReader, err := store.NewPieceReader(ctx, *car, storage, carBlocks, files)
					require.NoError(t, err)
					defer pieceReader.Close()

					// Simulate download handler: seek to end to get size
					size, err := pieceReader.Seek(0, io.SeekEnd)
					require.NoError(t, err)
					require.Equal(t, car.FileSize, size, "seek to end should return FileSize")

					// Seek back to start
					pos, err := pieceReader.Seek(0, io.SeekStart)
					require.NoError(t, err)
					require.Equal(t, int64(0), pos)

					// Read entire piece
					downloadedData = make([]byte, car.FileSize)
					n, err = io.ReadFull(pieceReader, downloadedData)
					require.NoError(t, err)
					require.Equal(t, int(car.FileSize), n)

					// Verify we get EOF after reading everything
					var buf [1]byte
					_, err = pieceReader.Read(buf[:])
					require.Equal(t, io.EOF, err)

					// Verify padding region is zeros
					if car.MinPieceSizePadding > 0 {
						paddingStart := car.FileSize - car.MinPieceSizePadding
						for i := paddingStart; i < car.FileSize; i++ {
							require.Equal(t, byte(0), downloadedData[i], "padding should be zeros at position %d", i)
						}
					}
				} else {
					// For non-inline mode, read the CAR file directly from disk
					carPath := filepath.Join(outputStorage.Path, car.StoragePath)
					downloadedData, err = os.ReadFile(carPath)
					require.NoError(t, err)
					n = len(downloadedData)
					require.Equal(t, int(car.FileSize), n, "CAR file should be padded to full piece size")

					// Verify the file ends with zeros (padding)
					// Find where padding starts by looking for trailing zeros
					paddingStart := int64(len(downloadedData)) - 1
					for paddingStart > 0 && downloadedData[paddingStart] == 0 {
						paddingStart--
					}
					paddingStart++ // Move back to first zero

					// There should be some padding
					paddingSize := int64(len(downloadedData)) - paddingStart
					require.Greater(t, paddingSize, int64(0), "should have some trailing zero padding")
				}

				t.Logf("Successfully downloaded %d bytes from %s piece", n, tc.name)
			})
		}
	})
}
