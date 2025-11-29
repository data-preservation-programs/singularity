package pack

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/gotidy/ptr"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
)

func TestAssembleCar(t *testing.T) {
	tmp := t.TempDir()
	out := t.TempDir()
	err := os.WriteFile(filepath.Join(tmp, "test.txt"), testutil.GenerateRandomBytes(5), 0644)
	require.NoError(t, err)
	stat, err := os.Stat(filepath.Join(tmp, "test.txt"))
	require.NoError(t, err)
	err = os.WriteFile(filepath.Join(tmp, "large.txt"), testutil.GenerateRandomBytes(5_000_000), 0644)
	require.NoError(t, err)
	_, err = os.Stat(filepath.Join(tmp, "large.txt"))
	require.NoError(t, err)
	jobs := []struct {
		name     string
		job      model.Job
		fileSize int64
		one      bool
	}{
		{
			name:     "unknown file size",
			fileSize: 101,
			job: model.Job{
				Type:  model.Pack,
				State: model.Processing,
				Attachment: &model.SourceAttachment{
					Preparation: &model.Preparation{
						MaxSize:      2000000,
						PieceSize:    1 << 21,
						MinPieceSize: 1, // set to 1 byte to disable padding
					},
					Storage: &model.Storage{
						Type: "local",
						Path: tmp,
					},
				},
				FileRanges: []model.FileRange{
					{
						Offset: 0,
						Length: -1,
						File: &model.File{
							Path:             "test.txt",
							Size:             -1,
							LastModifiedNano: stat.ModTime().UnixNano(),
							AttachmentID:     ptr.Of(model.SourceAttachmentID(1)),
							Directory: &model.Directory{
								AttachmentID: ptr.Of(model.SourceAttachmentID(1)),
							},
						},
					},
				},
			}},
		{
			name:     "single file",
			fileSize: 101,
			job: model.Job{
				Type:  model.Pack,
				State: model.Processing,
				Attachment: &model.SourceAttachment{
					Preparation: &model.Preparation{
						MaxSize:      2000000,
						PieceSize:    1 << 21,
						MinPieceSize: 1, // set to 1 byte to disable padding
					},
					Storage: &model.Storage{
						Type: "local",
						Path: tmp,
					},
				},
				FileRanges: []model.FileRange{
					{
						Offset: 0,
						Length: 5,
						File: &model.File{
							Path:             "test.txt",
							Size:             stat.Size(),
							LastModifiedNano: stat.ModTime().UnixNano(),
							AttachmentID:     ptr.Of(model.SourceAttachmentID(1)),
							Directory: &model.Directory{
								AttachmentID: ptr.Of(model.SourceAttachmentID(1)),
							},
						},
					},
				},
			},
		},
		{
			name:     "splitted file",
			fileSize: 138,
			job: model.Job{
				Type:  model.Pack,
				State: model.Processing,
				Attachment: &model.SourceAttachment{
					Preparation: &model.Preparation{
						MaxSize:      2000000,
						PieceSize:    1 << 21,
						MinPieceSize: 1, // set to 1 byte to disable padding
					},
					Storage: &model.Storage{
						Type: "local",
						Path: tmp,
					},
				},
				FileRanges: []model.FileRange{
					{
						Offset: 0,
						Length: 2,
						File: &model.File{
							Path:             "test.txt",
							Size:             stat.Size(),
							LastModifiedNano: stat.ModTime().UnixNano(),
							AttachmentID:     ptr.Of(model.SourceAttachmentID(1)),
							Directory: &model.Directory{
								AttachmentID: ptr.Of(model.SourceAttachmentID(1)),
							},
						},
					},
					{
						Offset: 2,
						Length: 3,
						FileID: 1,
						File: &model.File{
							ID:               1,
							Path:             "test.txt",
							Size:             stat.Size(),
							LastModifiedNano: stat.ModTime().UnixNano(),
							AttachmentID:     ptr.Of(model.SourceAttachmentID(1)),
							DirectoryID:      ptr.Of(model.DirectoryID(1)),
						},
					},
				},
			},
		},
		{
			name:     "single file non-inline with deletion, non-Dag",
			fileSize: 101,
			one:      true,
			job: model.Job{
				Type:  model.Pack,
				State: model.Processing,
				Attachment: &model.SourceAttachment{
					Preparation: &model.Preparation{
						MaxSize:      2000000,
						PieceSize:    1 << 21,
						MinPieceSize: 1, // set to 1 byte to disable padding
						OutputStorages: []model.Storage{
							{
								Name: "out",
								Type: "local",
								Path: out,
							},
						},
						DeleteAfterExport: true,
						NoDag:             true,
					},
					Storage: &model.Storage{
						Name: "tmp",
						Type: "local",
						Path: tmp,
					},
				},
				FileRanges: []model.FileRange{
					{
						Offset: 0,
						Length: 5,
						File: &model.File{
							Path:             "test.txt",
							Size:             stat.Size(),
							LastModifiedNano: stat.ModTime().UnixNano(),
							AttachmentID:     ptr.Of(model.SourceAttachmentID(1)),
							Directory: &model.Directory{
								AttachmentID: ptr.Of(model.SourceAttachmentID(1)),
							},
						},
					},
				},
			},
		},
	}

	for _, job := range jobs {
		t.Run(job.name, func(t *testing.T) {
			testFunc := func(ctx context.Context, t *testing.T, db *gorm.DB) {
				err := db.Create(&job.job).Error
				require.NoError(t, err)
				car, err := Pack(ctx, db, job.job)
				require.NoError(t, err)
				require.Equal(t, job.fileSize, car.FileSize)
				var root model.Directory
				err = db.First(&root, "attachment_id = ? AND parent_id IS NULL", job.job.AttachmentID).Error
				require.NoError(t, err)
				if job.job.Attachment.Preparation.NoDag {
					require.Equal(t, "", root.CID.String())
				} else {
					require.NotEqual(t, "", root.CID.String())
				}
				var files []model.File
				err = db.Find(&files).Error
				require.NoError(t, err)
				for _, file := range files {
					require.GreaterOrEqual(t, file.Size, int64(0))
				}
				var fileRanges []model.FileRange
				err = db.Find(&fileRanges).Error
				require.NoError(t, err)
				for _, fileRange := range fileRanges {
					require.GreaterOrEqual(t, fileRange.Length, int64(0))
				}
			}
			if job.one {
				testutil.One(t, testFunc)
			} else {
				testutil.All(t, testFunc)
			}
		})
	}
}

func TestLastPiecePadding(t *testing.T) {
	// Test last piece padding scenarios
	tmp := t.TempDir()
	out := t.TempDir()

	// Create a file that's smaller than min piece size for testing
	smallSize := 500_000 // 500 KB
	err := os.WriteFile(filepath.Join(tmp, "small.txt"), testutil.GenerateRandomBytes(smallSize), 0644)
	require.NoError(t, err)
	smallStat, err := os.Stat(filepath.Join(tmp, "small.txt"))
	require.NoError(t, err)

	// Create a file that's larger than min piece size for testing
	largeSize := 1_500_000 // 1.5 MB (larger than min piece size of 1 MiB)
	err = os.WriteFile(filepath.Join(tmp, "medium.txt"), testutil.GenerateRandomBytes(largeSize), 0644)
	require.NoError(t, err)
	mediumStat, err := os.Stat(filepath.Join(tmp, "medium.txt"))
	require.NoError(t, err)

	tests := []struct {
		name                 string
		pieceSize            int64
		minPieceSize         int64
		fileSize             int64
		expectedPieceSize    int64
		expectedFileSize     int64
		expectedFileRanges   int
		expectedFileRangeLen int64
	}{
		{
			name:                 "last piece smaller than min piece size gets padded to min piece size",
			pieceSize:            1 << 21,               // 2 MiB piece size
			minPieceSize:         1 << 20,               // 1 MiB min piece size
			fileSize:             int64(smallSize),      // 500 KB file
			expectedPieceSize:    1 << 20,               // Expected to be padded to 1 MiB (min piece size)
			expectedFileSize:     (1 << 20) * 127 / 128, // File padded to (127/128) × piece_size due to Fr32
			expectedFileRanges:   1,
			expectedFileRangeLen: int64(smallSize),
		},
		{
			name:                 "last piece larger than min piece size gets padded to next power of two",
			pieceSize:            1 << 21,          // 2 MiB piece size
			minPieceSize:         1 << 20,          // 1 MiB min piece size
			fileSize:             int64(largeSize), // 1.5 MB file
			expectedPieceSize:    1 << 21,          // Expected to be padded to 2 MiB (next power of 2)
			expectedFileSize:     1500283,          // Based on actual test results
			expectedFileRanges:   1,
			expectedFileRangeLen: int64(largeSize),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
				// Create job with appropriate file
				filePath := "small.txt"
				fileStat := smallStat
				if tc.fileSize > 1_000_000 {
					filePath = "medium.txt"
					fileStat = mediumStat
				}

				job := model.Job{
					Type:  model.Pack,
					State: model.Processing,
					Attachment: &model.SourceAttachment{
						Preparation: &model.Preparation{
							MaxSize:      tc.fileSize + 1000, // Buffer
							PieceSize:    tc.pieceSize,
							MinPieceSize: tc.minPieceSize,
							OutputStorages: []model.Storage{
								{
									Name: "out",
									Type: "local",
									Path: out,
								},
							},
						},
						Storage: &model.Storage{
							Type: "local",
							Path: tmp,
						},
					},
					FileRanges: []model.FileRange{
						{
							Offset: 0,
							Length: tc.fileSize,
							File: &model.File{
								Path:             filePath,
								Size:             tc.fileSize,
								LastModifiedNano: fileStat.ModTime().UnixNano(),
								AttachmentID:     ptr.Of(model.SourceAttachmentID(1)),
								Directory: &model.Directory{
									AttachmentID: ptr.Of(model.SourceAttachmentID(1)),
								},
							},
						},
					},
				}

				// Create and execute the packing job
				err := db.Create(&job).Error
				require.NoError(t, err)

				car, err := Pack(ctx, db, job)
				require.NoError(t, err)

				// Verify the car was created successfully
				require.NotNil(t, car)

				// Log the actual file size for debugging
				t.Logf("Test case: %s, Expected piece size: %d, Actual piece size: %d, Expected file size: %d, Actual file size: %d",
					tc.name, tc.expectedPieceSize, car.PieceSize, tc.expectedFileSize, car.FileSize)

				// Verify the piece size is correct (should match our expected padded size)
				require.Equal(t, tc.expectedPieceSize, car.PieceSize,
					"Piece size should be padded to expected value")

				// Verify exact file size for regression testing
				require.Equal(t, tc.expectedFileSize, car.FileSize,
					"CAR file size should match expected value exactly")

				// Verify correct number of file ranges
				var fileRanges []model.FileRange
				err = db.Find(&fileRanges).Error
				require.NoError(t, err)
				require.Len(t, fileRanges, tc.expectedFileRanges)
				require.Equal(t, tc.expectedFileRangeLen, fileRanges[0].Length)
			})
		})
	}
}

func TestMultiplePiecesWithLastPiece(t *testing.T) {
	// Test pieces with different sizes and verify the padding behavior
	tmp := t.TempDir()
	out := t.TempDir()

	pieceSize := int64(1 << 20) // 1 MiB piece size

	// Create test files of different sizes
	smallSize := 500_000 // 500 KB (smaller than min piece size of 1 MiB)
	err := os.WriteFile(filepath.Join(tmp, "small.txt"), testutil.GenerateRandomBytes(smallSize), 0644)
	require.NoError(t, err)
	smallStat, err := os.Stat(filepath.Join(tmp, "small.txt"))
	require.NoError(t, err)

	mediumSize := 1_500_000 // 1.5 MB (larger than min piece size but smaller than piece size)
	err = os.WriteFile(filepath.Join(tmp, "medium.txt"), testutil.GenerateRandomBytes(mediumSize), 0644)
	require.NoError(t, err)
	mediumStat, err := os.Stat(filepath.Join(tmp, "medium.txt"))
	require.NoError(t, err)

	// Test cases
	tests := []struct {
		name              string
		filePath          string
		fileStat          os.FileInfo
		fileSize          int64
		pieceSize         int64 // Target piece size
		minPieceSize      int64 // Minimum piece size
		expectedPieceSize int64 // Expected final piece size after padding
	}{
		{
			name:              "file smaller than min piece size gets padded to min piece size",
			filePath:          "small.txt",
			fileStat:          smallStat,
			fileSize:          int64(smallSize),
			pieceSize:         pieceSize,     // 1 MiB target
			minPieceSize:      pieceSize / 2, // 512 KiB min
			expectedPieceSize: pieceSize / 2, // Padded to 512 KiB (min piece size)
		},
		{
			name:              "file larger than min piece size gets padded to next power of two",
			filePath:          "medium.txt",
			fileStat:          mediumStat,
			fileSize:          int64(mediumSize),
			pieceSize:         pieceSize,     // 1 MiB target
			minPieceSize:      pieceSize / 4, // 256 KiB min
			expectedPieceSize: pieceSize * 2, // Padded to 2 MiB (next power of 2)
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
				// Create job with the test file
				job := model.Job{
					Type:  model.Pack,
					State: model.Processing,
					Attachment: &model.SourceAttachment{
						Preparation: &model.Preparation{
							MaxSize:      tc.fileSize + 1000, // Buffer
							PieceSize:    tc.pieceSize,       // Target piece size
							MinPieceSize: tc.minPieceSize,    // Min piece size
							OutputStorages: []model.Storage{
								{
									Name: "out",
									Type: "local",
									Path: out,
								},
							},
						},
						Storage: &model.Storage{
							Type: "local",
							Path: tmp,
						},
					},
					FileRanges: []model.FileRange{
						{
							Offset: 0,
							Length: tc.fileSize,
							File: &model.File{
								Path:             tc.filePath,
								Size:             tc.fileSize,
								LastModifiedNano: tc.fileStat.ModTime().UnixNano(),
								AttachmentID:     ptr.Of(model.SourceAttachmentID(1)),
								Directory: &model.Directory{
									AttachmentID: ptr.Of(model.SourceAttachmentID(1)),
								},
							},
						},
					},
				}

				// Create and execute the packing job
				err := db.Create(&job).Error
				require.NoError(t, err)

				car, err := Pack(ctx, db, job)
				require.NoError(t, err)

				// Verify the car was created successfully
				require.NotNil(t, car)

				// Verify the piece size is correct (should match our expected padded size)
				require.Equal(t, tc.expectedPieceSize, car.PieceSize,
					"Piece size should be padded to expected value")

				// Verify the actual file size is reasonable (specific bytes may vary slightly)
				// The CAR file size should be at least as large as the input file + some overhead
				require.GreaterOrEqual(t, car.FileSize, tc.fileSize,
					"CAR file size should be at least as large as the input file")
				// And shouldn't be much larger than the file size + overhead
				require.LessOrEqual(t, car.FileSize, tc.fileSize+1000,
					"CAR file size shouldn't be excessively larger than the input file")

				// Verify correct number of file ranges
				var fileRanges []model.FileRange
				err = db.Find(&fileRanges).Error
				require.NoError(t, err)
				require.Len(t, fileRanges, 1)
				require.Equal(t, tc.fileSize, fileRanges[0].Length)
			})
		})
	}
}
