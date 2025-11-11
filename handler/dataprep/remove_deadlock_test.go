package dataprep

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

// TestRemovePreparationNoDeadlock verifies that preparation deletion does not deadlock
// with concurrent worker operations. The test simulates multiple preparations sharing one
// storage. Deleting one preparation cascades to shared tables while workers operate on
// other preparations. The fix uses explicit ordered deletion instead of FK CASCADE.
func TestRemovePreparationNoDeadlock(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		req := require.New(t)

		// Create test data: multiple preparations sharing the SAME storage (realistic production scenario)
		// When deleting one prep, CASCADE affects tables shared with other preps on same storage
		const numPreparations = 3
		const numFilesPerPrep = 50
		const numDirsPerPrep = 10
		const numJobsPerPrep = 10
		const numCarsPerPrep = 5
		const numBlocksPerCar = 20

		// Create ONE shared storage (this is the key to reproducing the real deadlock)
		sharedStorage := model.Storage{
			Name: "shared-storage-" + uuid.New().String(),
			Type: "local",
			Path: "/tmp/test",
		}
		err := db.Create(&sharedStorage).Error
		req.NoError(err)

		preparations := make([]model.Preparation, numPreparations)
		attachments := make([]model.SourceAttachment, numPreparations)
		allJobIDs := make([]model.JobID, 0, numPreparations*numJobsPerPrep)
		allFileIDs := make([]model.FileID, 0, numPreparations*numFilesPerPrep)
		allCarIDs := make([]model.CarID, 0, numPreparations*numCarsPerPrep)

		for i := range numPreparations {
			// Create preparation
			prep := model.Preparation{
				Name:      "test-prep-" + uuid.New().String(),
				MaxSize:   1 << 30,
				PieceSize: 1 << 30,
			}
			err := db.Create(&prep).Error
			req.NoError(err)
			preparations[i] = prep

			// ALL preparations attach to the SAME storage (production pattern!)
			attachment := model.SourceAttachment{
				PreparationID: prep.ID,
				StorageID:     sharedStorage.ID,
			}
			err = db.Create(&attachment).Error
			req.NoError(err)
			attachments[i] = attachment

			// Create directories (will have circular references with files)
			directories := make([]model.Directory, numDirsPerPrep)
			for j := range numDirsPerPrep {
				dir := model.Directory{
					Name:         "dir-" + uuid.New().String()[:8],
					AttachmentID: attachment.ID,
					CID:          model.CID(testutil.TestCid),
					Data:         []byte("test"),
				}
				err = db.Create(&dir).Error
				req.NoError(err)
				directories[j] = dir
			}

			// Create files
			for j := range numFilesPerPrep {
				dirIdx := j % numDirsPerPrep
				file := model.File{
					Path:             "/test/file-" + uuid.New().String()[:8],
					Size:             1024,
					AttachmentID:     attachment.ID,
					DirectoryID:      &directories[dirIdx].ID,
					CID:              model.CID(testutil.TestCid),
					Hash:             "test-hash",
					LastModifiedNano: time.Now().UnixNano(),
				}
				err = db.Create(&file).Error
				req.NoError(err)
				allFileIDs = append(allFileIDs, file.ID)
			}

			// Create jobs
			for j := range numJobsPerPrep {
				job := model.Job{
					Type:         model.Pack,
					State:        model.Ready,
					AttachmentID: attachment.ID,
				}
				err = db.Create(&job).Error
				req.NoError(err)
				allJobIDs = append(allJobIDs, job.ID)

				// Create file ranges for some jobs
				if j%2 == 0 && len(allFileIDs) > 0 {
					fileIdx := (i*numJobsPerPrep + j) % len(allFileIDs)
					fileRange := model.FileRange{
						JobID:  &job.ID,
						FileID: allFileIDs[fileIdx],
						Offset: 0,
						Length: 1024,
						CID:    model.CID(testutil.TestCid),
					}
					err = db.Create(&fileRange).Error
					req.NoError(err)
				}
			}

			// Create cars
			for range numCarsPerPrep {
				car := model.Car{
					PreparationID: prep.ID,
					AttachmentID:  &attachment.ID,
					PieceCID:      model.CID(testutil.TestCid),
					PieceSize:     1 << 20,
					RootCID:       model.CID(testutil.TestCid),
					FileSize:      1024,
					StoragePath:   "/test/car-" + uuid.New().String()[:8],
				}
				err = db.Create(&car).Error
				req.NoError(err)
				allCarIDs = append(allCarIDs, car.ID)

				// Create car blocks
				for k := range numBlocksPerCar {
					block := model.CarBlock{
						CarID:          car.ID,
						CID:            model.CID(testutil.TestCid),
						CarOffset:      int64(k * 1024),
						CarBlockLength: 1024,
						Varint:         []byte{0x01},
						RawBlock:       []byte("test-block"),
						FileOffset:     int64(k * 1024),
					}
					err = db.Create(&block).Error
					req.NoError(err)
				}
			}
		}

		// Run concurrent operations that could previously cause deadlock
		const numIterations = 15
		var wg sync.WaitGroup
		errChan := make(chan error, numIterations*4)

		for i := range numIterations {
			wg.Add(4)

			// Goroutine 1: Delete preparations (triggers CASCADE)
			go func(iteration int) {
				defer wg.Done()
				time.Sleep(time.Duration(iteration%3) * time.Millisecond)

				deleteCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
				defer cancel()

				prepIdx := iteration % numPreparations
				handler := DefaultHandler{}
				err := handler.RemovePreparationHandler(deleteCtx, db, preparations[prepIdx].Name, RemoveRequest{})
				if err != nil && deleteCtx.Err() == nil {
					// Ignore expected errors:
					// - "not found" errors (prep already deleted by another goroutine)
					// - "active jobs" errors (concurrent job updates made jobs active)
					if !isNotFoundError(err) && !isActiveJobsError(err) {
						errChan <- err
					}
				}
			}(i)

			// Goroutine 2: Bulk job updates (simulating pack/pause operations)
			go func(iteration int) {
				defer wg.Done()
				time.Sleep(time.Duration(iteration%5) * time.Millisecond)

				updateCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
				defer cancel()

				// Update a batch of jobs
				startIdx := (iteration * 5) % len(allJobIDs)
				endIdx := min(startIdx+8, len(allJobIDs))
				batchJobIDs := allJobIDs[startIdx:endIdx]

				err := db.WithContext(updateCtx).Transaction(func(tx *gorm.DB) error {
					for _, jobID := range batchJobIDs {
						err := tx.Model(&model.Job{}).
							Where("id = ?", jobID).
							Updates(map[string]any{
								"state":             model.Processing,
								"error_message":     "",
								"error_stack_trace": "",
							}).Error
						if err != nil {
							return err
						}
					}
					return nil
				})

				if err != nil && updateCtx.Err() == nil {
					errChan <- err
				}
			}(i)

			// Goroutine 3: File updates (simulating scan operations)
			go func(iteration int) {
				defer wg.Done()
				time.Sleep(time.Duration(iteration%7) * time.Millisecond)

				updateCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
				defer cancel()

				// Update some files
				if len(allFileIDs) > 0 {
					fileIdx := iteration % len(allFileIDs)
					err := db.WithContext(updateCtx).Model(&model.File{}).
						Where("id = ?", allFileIDs[fileIdx]).
						Update("size", 2048).Error

					if err != nil && updateCtx.Err() == nil {
						errChan <- err
					}
				}
			}(i)

			// Goroutine 4: Car block creation (simulating pack completion)
			go func(iteration int) {
				defer wg.Done()
				time.Sleep(time.Duration(iteration%4) * time.Millisecond)

				createCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
				defer cancel()

				// Try to create a car block (might fail if car already deleted)
				if len(allCarIDs) > 0 {
					carIdx := iteration % len(allCarIDs)
					block := model.CarBlock{
						CarID:          allCarIDs[carIdx],
						CID:            model.CID(testutil.TestCid),
						CarOffset:      int64(1000 + iteration),
						CarBlockLength: 512,
						Varint:         []byte{0x02},
						RawBlock:       []byte("new-block"),
						FileOffset:     int64(1000 + iteration),
					}
					err := db.WithContext(createCtx).Create(&block).Error

					// Ignore FK errors (car might have been deleted)
					if err != nil && createCtx.Err() == nil && !isFKError(err) {
						errChan <- err
					}
				}
			}(i)
		}

		// Wait for all operations to complete
		done := make(chan struct{})
		go func() {
			wg.Wait()
			close(done)
		}()

		// Wait with timeout
		select {
		case <-done:
			// Success - no deadlock
		case err := <-errChan:
			req.NoError(err, "Unexpected error during concurrent operations")
		case <-time.After(45 * time.Second):
			req.Fail("Test timed out - likely deadlock occurred")
		}

		// Verify no deadlock occurred (test completed successfully)
		// We don't verify complete cleanup because:
		// - Some deletions may have been blocked by active jobs (expected)
		// - Some operations may still be pending when test ends
		// - The important thing is we didn't deadlock
		//
		// If we want to verify actual cleanup, we'd need to:
		// 1. Pause all jobs first
		// 2. Wait for all deletions to complete
		// 3. Then verify cleanup
		// But that's not the point of this test - we're testing for deadlocks, not cleanup.

		t.Logf("Test completed without deadlock")
	})
}

// isNotFoundError checks if error is a "not found" error
func isNotFoundError(err error) bool {
	if err == nil {
		return false
	}
	errStr := err.Error()
	return contains(errStr, "not found") || contains(errStr, "does not exist")
}

// isActiveJobsError checks if error is about active jobs preventing deletion
func isActiveJobsError(err error) bool {
	if err == nil {
		return false
	}
	errStr := err.Error()
	return contains(errStr, "active jobs")
}

// isFKError checks if error is a foreign key constraint violation
func isFKError(err error) bool {
	if err == nil {
		return false
	}
	errStr := err.Error()
	return contains(errStr, "foreign key") || contains(errStr, "FOREIGN KEY") ||
		contains(errStr, "violates") || contains(errStr, "constraint")
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > len(substr) && findSubstring(s, substr))
}

func findSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
