package pack

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ipfs/go-cid"
	"github.com/ipld/go-car/v2"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/scan"
	"github.com/data-preservation-programs/singularity/util/testutil"
)

func TestLastPieceBehaviorE2ENoInline(t *testing.T) {
	// This is an end-to-end test that verifies the last piece behavior by:
	// 1. Creating a dataset with a file that will be split across multiple pieces
	// 2. Using scan to automatically create pack jobs
	// 3. Running those pack jobs
	// 4. Verifying the resulting pieces have the expected properties

	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		// Setup: Create temporary directories for source and output
		sourceDir := t.TempDir()
		outputDir := t.TempDir()

		// 1. Create test files with known sizes
		// Create a large random file that will be split into multiple pieces
		// Each piece will exercise different padding behavior
		testFileSize := 4_200_000 // ~4.2 MB - will be split into multiple pieces

		// Create the test file with random data
		err := os.WriteFile(filepath.Join(sourceDir, "large_file.bin"),
			testutil.GenerateRandomBytes(testFileSize), 0644)
		require.NoError(t, err)

		// 2. Create a preparation with specific settings
		pieceSize := int64(2 * 1024 * 1024)    // 2 MiB target piece size
		minPieceSize := int64(1 * 1024 * 1024) // 1 MiB min piece size
		maxSize := pieceSize / 3               // Set max size to ensure we get multiple pieces from our file

		prep := model.Preparation{
			Name:         "test-preparation",
			MaxSize:      maxSize,      // Each job will have at most maxSize bytes (forcing splitting)
			PieceSize:    pieceSize,    // Target piece size
			MinPieceSize: minPieceSize, // Minimum piece size
			NoInline:     true,         // Force writing CAR files to disk instead of using inline preparation
			SourceStorages: []model.Storage{
				{
					Name: "source-storage",
					Type: "local",
					Path: sourceDir,
				},
			},
			OutputStorages: []model.Storage{
				{
					Name: "output-storage",
					Type: "local",
					Path: outputDir,
				},
			},
		}

		// Save the preparation
		err = db.Create(&prep).Error
		require.NoError(t, err)

		// 3. Create the source attachment
		var sourceAttachment model.SourceAttachment
		err = db.Preload("Storage").Preload("Preparation").
			Where("preparation_id = ? AND storage_id = ?", prep.ID, prep.SourceStorages[0].ID).
			First(&sourceAttachment).Error
		require.NoError(t, err)

		// 4. Run the scan job to discover files and create pack jobs
		err = db.Create(&model.Directory{
			AttachmentID: &sourceAttachment.ID,
			Name:         "", // Root directory has empty name
			ParentID:     nil,
		}).Error
		require.NoError(t, err)

		// Run the scan
		t.Logf("Running scan job")
		err = scan.Scan(ctx, db, sourceAttachment)
		require.NoError(t, err)

		// 5. Verify scan created appropriate jobs
		var packJobs []model.Job
		err = db.Where("type = ? AND state = ?", model.Pack, model.Ready).Find(&packJobs).Error
		require.NoError(t, err)

		// We should have multiple pack jobs due to the file size and max size setting
		require.Greater(t, len(packJobs), 2, "Scan should have created multiple pack jobs")

		for i := range packJobs {
			t.Logf("Pack job %d created", i+1)
		}

		// 6. Run all pack jobs and collect CAR files for verification
		carSizes := make(map[int64]int64)

		for _, job := range packJobs {
			// Load the full job with attachments - important to preload OutputStorages
			err = db.Preload("Attachment.Preparation.OutputStorages").Preload("Attachment.Storage").
				Preload("FileRanges.File").Where("id = ?", job.ID).First(&job).Error
			require.NoError(t, err)

			// Execute the pack job
			car, err := Pack(ctx, db, job)
			require.NoError(t, err)

			// Log job and car details
			fileRangeInfo := ""
			if len(job.FileRanges) > 0 {
				fileRangeInfo = fmt.Sprintf(", range length: %d", job.FileRanges[0].Length)
			}
			t.Logf("Packed job ID %d, created car with piece size: %d, file size: %d%s",
				job.ID, car.PieceSize, car.FileSize, fileRangeInfo)

			// Record car sizes for later verification
			carSizes[car.PieceSize] = car.FileSize

			// Update job state
			err = db.Model(&model.Job{}).Where("id = ?", job.ID).Update("state", model.Complete).Error
			require.NoError(t, err)
		}

		// 7. Verify the resulting Cars
		var cars []model.Car
		err = db.Find(&cars).Error
		require.NoError(t, err)

		// Find all CAR files in the output directory
		outputDirFiles, err := os.ReadDir(outputDir)
		require.NoError(t, err)

		// Collect CAR file paths for verification
		var carFilePaths []string
		for _, file := range outputDirFiles {
			if !file.IsDir() && strings.HasSuffix(file.Name(), ".car") {
				carFilePaths = append(carFilePaths, filepath.Join(outputDir, file.Name()))
			}
		}

		require.NotEmpty(t, carFilePaths, "Should have CAR files in the output directory")
		t.Logf("Found %d CAR files in the output directory", len(carFilePaths))

		// Verify we have the expected number of cars matching our jobs
		require.Equal(t, len(packJobs), len(cars), "Should have one car per pack job")
		require.Equal(t, len(packJobs), len(carFilePaths), "Should have one CAR file per pack job")

		// Count cars by piece size
		fullSizePieceCount := 0  // 2 MiB or 4 MiB
		halfSizePieceCount := 0  // 1 MiB
		otherSizePieceCount := 0 // Anything else

		for _, car := range cars {
			t.Logf("Car has piece size: %d, file size: %d", car.PieceSize, car.FileSize)

			if car.PieceSize == pieceSize || car.PieceSize == pieceSize*2 {
				// Full-sized piece (2 MiB or 4 MiB)
				fullSizePieceCount++
				require.Greater(t, car.FileSize, int64(0), "Car file size should be greater than 0")
			} else if car.PieceSize == minPieceSize {
				// Piece padded to min piece size (1 MiB)
				halfSizePieceCount++
				require.Greater(t, car.FileSize, int64(0), "Car file size should be greater than 0")
			} else {
				t.Logf("Found car with unexpected piece size: %d", car.PieceSize)
				otherSizePieceCount++
			}
		}

		// Verify we have the expected types of pieces
		require.Equal(t, 0, otherSizePieceCount, "Should not have any cars with unexpected piece sizes")
		require.Equal(t, fullSizePieceCount+halfSizePieceCount, len(packJobs), "Should have exactly one car per pack job")

		// At least one piece should be padded to min piece size (last piece)
		require.GreaterOrEqual(t, halfSizePieceCount, 1, "Should have at least 1 car padded to min piece size")

		// 8. Verify that file ranges have valid CIDs
		var fileRanges []model.FileRange
		err = db.Find(&fileRanges).Error
		require.NoError(t, err)
		require.Greater(t, len(fileRanges), 0, "Should have at least one file range")

		// Verify that all file ranges have CIDs
		for _, fileRange := range fileRanges {
			require.NotEqual(t, cid.Undef, cid.Cid(fileRange.CID), "File range should have a valid CID")
		}

		// 9. Verify CAR file format using go-car's verification
		for _, carFilePath := range carFilePaths {
			// Verify the CAR file format
			reader, err := car.OpenReader(carFilePath)
			require.NoError(t, err, "Should be able to open CAR file %s", carFilePath)
			defer reader.Close()

			// Verify the CAR has roots
			roots, err := reader.Roots()
			require.NoError(t, err, "Should be able to read CAR roots")
			require.NotEmpty(t, roots, "CAR file should have at least one root")

			// Read all blocks to verify integrity
			rd, err := os.Open(carFilePath)
			require.NoError(t, err)
			defer rd.Close()

			blockReader, err := car.NewBlockReader(rd)
			require.NoError(t, err, "Should be able to create block reader")

			blockCount := 0
			for {
				block, err := blockReader.Next()
				if err == io.EOF {
					break
				}
				require.NoError(t, err, "Should be able to read all blocks")
				require.NotNil(t, block, "Block should not be nil")
				require.NotEqual(t, cid.Undef, block.Cid(), "Block should have valid CID")
				blockCount++
			}

			require.Greater(t, blockCount, 0, "CAR file should contain at least one block")
			t.Logf("Verified CAR file %s: found %d blocks", filepath.Base(carFilePath), blockCount)
		}
	})
}

func TestLastPieceBehaviorE2EInline(t *testing.T) {
	// This is an end-to-end test that verifies the last piece behavior with inline CARs by:
	// 1. Creating a dataset with a file that will be split across multiple pieces
	// 2. Using scan to automatically create pack jobs
	// 3. Running those pack jobs with NoInline:false (so CAR data is stored in database)
	// 4. Verifying the resulting pieces have the expected properties

	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		// Setup: Create temporary directories for source and output
		sourceDir := t.TempDir()
		outputDir := t.TempDir()

		// 1. Create test files with known sizes
		// Create a large random file that will be split into multiple pieces
		// Each piece will exercise different padding behavior
		testFileSize := 4_200_000 // ~4.2 MB - will be split into multiple pieces

		// Create the test file with random data
		err := os.WriteFile(filepath.Join(sourceDir, "large_file.bin"),
			testutil.GenerateRandomBytes(testFileSize), 0644)
		require.NoError(t, err)

		// 2. Create a preparation with specific settings
		pieceSize := int64(2 * 1024 * 1024)    // 2 MiB target piece size
		minPieceSize := int64(1 * 1024 * 1024) // 1 MiB min piece size
		maxSize := pieceSize / 3               // Set max size to ensure we get multiple pieces from our file

		prep := model.Preparation{
			Name:         "test-preparation",
			MaxSize:      maxSize,      // Each job will have at most maxSize bytes (forcing splitting)
			PieceSize:    pieceSize,    // Target piece size
			MinPieceSize: minPieceSize, // Minimum piece size
			NoInline:     false,        // Use inline preparation (CAR data stored in database)
			SourceStorages: []model.Storage{
				{
					Name: "source-storage",
					Type: "local",
					Path: sourceDir,
				},
			},
			OutputStorages: []model.Storage{
				{
					Name: "output-storage",
					Type: "local",
					Path: outputDir,
				},
			},
		}

		// Save the preparation
		err = db.Create(&prep).Error
		require.NoError(t, err)

		// 3. Create the source attachment
		var sourceAttachment model.SourceAttachment
		err = db.Preload("Storage").Preload("Preparation").
			Where("preparation_id = ? AND storage_id = ?", prep.ID, prep.SourceStorages[0].ID).
			First(&sourceAttachment).Error
		require.NoError(t, err)

		// 4. Run the scan job to discover files and create pack jobs
		err = db.Create(&model.Directory{
			AttachmentID: &sourceAttachment.ID,
			Name:         "", // Root directory has empty name
			ParentID:     nil,
		}).Error
		require.NoError(t, err)

		// Run the scan
		t.Logf("Running scan job")
		err = scan.Scan(ctx, db, sourceAttachment)
		require.NoError(t, err)

		// 5. Verify scan created appropriate jobs
		var packJobs []model.Job
		err = db.Where("type = ? AND state = ?", model.Pack, model.Ready).Find(&packJobs).Error
		require.NoError(t, err)

		// We should have multiple pack jobs due to the file size and max size setting
		require.Greater(t, len(packJobs), 2, "Scan should have created multiple pack jobs")

		for i := range packJobs {
			t.Logf("Pack job %d created", i+1)
		}

		// 6. Run all pack jobs and collect CAR files for verification
		carSizes := make(map[int64]int64)

		for _, job := range packJobs {
			// Load the full job with attachments
			err = db.Preload("Attachment.Preparation").Preload("Attachment.Storage").
				Preload("FileRanges.File").Where("id = ?", job.ID).First(&job).Error
			require.NoError(t, err)

			// Execute the pack job
			car, err := Pack(ctx, db, job)
			require.NoError(t, err)

			// Log job and car details
			fileRangeInfo := ""
			if len(job.FileRanges) > 0 {
				fileRangeInfo = fmt.Sprintf(", range length: %d", job.FileRanges[0].Length)
			}
			t.Logf("Packed job ID %d, created car with piece size: %d, file size: %d%s",
				job.ID, car.PieceSize, car.FileSize, fileRangeInfo)

			// Record car sizes for later verification
			carSizes[car.PieceSize] = car.FileSize

			// Update job state
			err = db.Model(&model.Job{}).Where("id = ?", job.ID).Update("state", model.Complete).Error
			require.NoError(t, err)
		}

		// 7. Verify the resulting Cars
		var cars []model.Car
		err = db.Find(&cars).Error
		require.NoError(t, err)

		// For inline preparation, no CAR files should be in the output directory
		outputDirFiles, err := os.ReadDir(outputDir)
		require.NoError(t, err)

		carFileCount := 0
		for _, file := range outputDirFiles {
			if !file.IsDir() && strings.HasSuffix(file.Name(), ".car") {
				carFileCount++
			}
		}

		require.Equal(t, 0, carFileCount, "Should not have CAR files on disk for inline preparation")

		// Count cars by piece size
		fullSizePieceCount := 0  // 2 MiB or 4 MiB
		halfSizePieceCount := 0  // 1 MiB
		otherSizePieceCount := 0 // Anything else

		for _, car := range cars {
			t.Logf("Car has piece size: %d, file size: %d", car.PieceSize, car.FileSize)

			if car.PieceSize == pieceSize || car.PieceSize == pieceSize*2 {
				// Full-sized piece (2 MiB or 4 MiB)
				fullSizePieceCount++
				require.Greater(t, car.FileSize, int64(0), "Car file size should be greater than 0")
				// For inline preparation, cars should exist in database but not have file paths
				require.Empty(t, car.StoragePath, "Car storage path should be empty for inline preparation")
			} else if car.PieceSize == minPieceSize {
				// Piece padded to min piece size (1 MiB)
				halfSizePieceCount++
				require.Greater(t, car.FileSize, int64(0), "Car file size should be greater than 0")
				require.Empty(t, car.StoragePath, "Car storage path should be empty for inline preparation")
			} else {
				t.Logf("Found car with unexpected piece size: %d", car.PieceSize)
				otherSizePieceCount++
			}
		}

		// Verify we have the expected types of pieces
		require.Equal(t, 0, otherSizePieceCount, "Should not have any cars with unexpected piece sizes")
		require.Equal(t, fullSizePieceCount+halfSizePieceCount, len(packJobs), "Should have exactly one car per pack job")

		// At least one piece should be padded to min piece size (last piece)
		require.GreaterOrEqual(t, halfSizePieceCount, 1, "Should have at least 1 car padded to min piece size")

		// 8. Verify that file ranges have valid CIDs
		var fileRanges []model.FileRange
		err = db.Find(&fileRanges).Error
		require.NoError(t, err)
		require.Greater(t, len(fileRanges), 0, "Should have at least one file range")

		// Verify that all file ranges have CIDs
		for _, fileRange := range fileRanges {
			require.NotEqual(t, cid.Undef, cid.Cid(fileRange.CID), "File range should have a valid CID")
		}
	})
}
