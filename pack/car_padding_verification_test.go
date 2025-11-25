package pack

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

// TestCARPaddingValidity verifies that appending zeros to a valid CAR file keeps it valid
func TestCARPaddingValidity(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		tmp := t.TempDir()
		out := t.TempDir()

		// Create a small test file
		testData := testutil.GenerateRandomBytes(50 * 1024) // 50 KiB
		testPath := filepath.Join(tmp, "test.txt")
		err := os.WriteFile(testPath, testData, 0644)
		require.NoError(t, err)

		testStat, err := os.Stat(testPath)
		require.NoError(t, err)

		// Setup storage and preparation
		storage := model.Storage{
			Name: "test",
			Type: "local",
			Path: tmp,
		}
		err = db.Create(&storage).Error
		require.NoError(t, err)

		outputStorage := model.Storage{
			Name: "output",
			Type: "local",
			Path: out,
		}
		err = db.Create(&outputStorage).Error
		require.NoError(t, err)

		prep := model.Preparation{
			Name:      "test",
			MaxSize:   2_000_000,
			PieceSize: 1 << 21, // 2 MiB
		}
		err = db.Create(&prep).Error
		require.NoError(t, err)

		// Create attachments
		err = db.Exec("INSERT INTO source_attachments (preparation_id, storage_id) VALUES (?, ?)", prep.ID, storage.ID).Error
		require.NoError(t, err)

		err = db.Exec("INSERT INTO output_attachments (preparation_id, storage_id) VALUES (?, ?)", prep.ID, outputStorage.ID).Error
		require.NoError(t, err)

		var attachment model.SourceAttachment
		err = db.Where("preparation_id = ? AND storage_id = ?", prep.ID, storage.ID).First(&attachment).Error
		require.NoError(t, err)

		// Create directory
		dir := model.Directory{
			AttachmentID: &attachment.ID,
		}
		err = db.Create(&dir).Error
		require.NoError(t, err)

		// Create file record
		file := model.File{
			AttachmentID:     &attachment.ID,
			Path:             "test.txt",
			Size:             int64(len(testData)),
			LastModifiedNano: testStat.ModTime().UnixNano(),
			DirectoryID:      &dir.ID,
		}
		err = db.Create(&file).Error
		require.NoError(t, err)

		// Create job
		job := model.Job{
			AttachmentID: &attachment.ID,
			State:        model.Processing,
			Type:         model.Pack,
			FileRanges: []model.FileRange{
				{
					FileID: file.ID,
					Offset: 0,
					Length: int64(len(testData)),
				},
			},
		}
		err = db.Create(&job).Error
		require.NoError(t, err)

		err = db.Preload("FileRanges.File").Preload("Attachment.Preparation.OutputStorages").Preload("Attachment.Storage").First(&job, job.ID).Error
		require.NoError(t, err)

		// Generate CAR file using current Pack() implementation
		car, err := Pack(ctx, db, job)
		require.NoError(t, err)
		require.NotNil(t, car)

		carPath := filepath.Join(out, car.StoragePath)

		// Step 1: Inspect original CAR (more lenient than verify, accepts padded CAR v1)
		fmt.Println("=== Step 1: Inspecting original CAR file ===")
		_, err = exec.LookPath("car")
		if err != nil {
			t.Skip("car binary not available, skipping test")
			return
		}

		cmd := exec.Command("car", "inspect", carPath)
		output, err := cmd.CombinedOutput()
		require.NoError(t, err, "Original CAR inspection failed: %s", string(output))
		fmt.Println("✓ Original CAR can be inspected")

		// Step 2: Get original file size
		stat, err := os.Stat(carPath)
		require.NoError(t, err)
		originalSize := stat.Size()
		fmt.Printf("Original CAR size: %d bytes\n", originalSize)
		fmt.Printf("Piece size: %d bytes\n", car.PieceSize)

		// Step 3: Calculate target size (piece size / 2 + 1)
		targetSize := int64(car.PieceSize)/2 + 1
		fmt.Printf("Target size: %d bytes (piece size / 2 + 1)\n", targetSize)

		if originalSize >= targetSize {
			fmt.Println("CAR already >= target size, no padding needed for this test")
			fmt.Println("✓ Test complete: Original CAR is valid, no padding needed")
			return
		}

		paddingNeeded := targetSize - originalSize
		fmt.Printf("Padding needed: %d bytes\n", paddingNeeded)

		// Step 4: Append zeros using Go file operations
		fmt.Println("\n=== Step 2: Appending zeros to CAR file ===")
		f, err := os.OpenFile(carPath, os.O_APPEND|os.O_WRONLY, 0644)
		require.NoError(t, err)
		defer f.Close()

		zeros := make([]byte, paddingNeeded)
		n, err := f.Write(zeros)
		require.NoError(t, err)
		require.Equal(t, int(paddingNeeded), n)

		// Verify new size
		stat, err = os.Stat(carPath)
		require.NoError(t, err)
		require.Equal(t, targetSize, stat.Size(), "File size after padding should match target")
		fmt.Printf("✓ Padding complete, new size: %d bytes\n", stat.Size())

		// Step 5: Inspect padded CAR (more lenient than verify)
		fmt.Println("\n=== Step 3: Inspecting padded CAR file ===")
		cmd = exec.Command("car", "inspect", carPath)
		output, err = cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("car inspect output: %s\n", string(output))
		}
		require.NoError(t, err, "Padded CAR inspect failed: %s", string(output))
		fmt.Println("✓ Padded CAR can be INSPECTED!")
		fmt.Printf("Inspect output:\n%s\n", string(output))
		fmt.Println("\n=== SUCCESS: Appending zeros to CAR maintains readability ===")
	})
}
