package dataprep

import (
	"context"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/database"
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
		testutil.EnableDeadlockLogging(t, db)

		// Create test data: multiple preparations sharing one storage
		const numPreparations = 3
		const numFilesPerPrep = 50
		const numDirsPerPrep = 10
		const numJobsPerPrep = 10
		const numCarsPerPrep = 5
		const numBlocksPerCar = 20

		// Create shared storage
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

			// Attach to shared storage
			attachment := model.SourceAttachment{
				PreparationID: prep.ID,
				StorageID:     sharedStorage.ID,
			}
			err = db.Create(&attachment).Error
			req.NoError(err)
			attachments[i] = attachment

			// Create directories
			directories := make([]model.Directory, numDirsPerPrep)
			attachmentID := attachment.ID
			for j := range numDirsPerPrep {
				dir := model.Directory{
					Name:         "dir-" + uuid.New().String()[:8],
					AttachmentID: &attachmentID,
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
					AttachmentID:     &attachmentID,
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
					AttachmentID: &attachmentID,
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
			prepID := prep.ID
			for range numCarsPerPrep {
				car := model.Car{
					PreparationID: &prepID,
					AttachmentID:  &attachmentID,
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
				carID := car.ID
				for k := range numBlocksPerCar {
					block := model.CarBlock{
						CarID:          &carID,
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

		const numIterations = 15
		var wg sync.WaitGroup
		errChan := make(chan error, numIterations*4)

		for i := range numIterations {
			wg.Add(4)

			// Goroutine 1: Delete preparations
			go func(iteration int) {
				defer wg.Done()
				time.Sleep(time.Duration(iteration%3) * time.Millisecond)

				deleteCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
				defer cancel()

				prepIdx := iteration % numPreparations
				handler := DefaultHandler{}
				err := handler.RemovePreparationHandler(deleteCtx, db, preparations[prepIdx].Name, RemoveRequest{})
				if err != nil && deleteCtx.Err() == nil {
					// Ignore expected errors from concurrent operations
					if !isNotFoundError(err) && !isActiveJobsError(err) && !isJobsInUseError(err) {
						errChan <- err
					}
				}
			}(i)

			// Goroutine 2: Bulk job updates
			go func(iteration int) {
				defer wg.Done()
				time.Sleep(time.Duration(iteration%5) * time.Millisecond)

				updateCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
				defer cancel()

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
					// If it's a deadlock error, get InnoDB status
					if strings.Contains(err.Error(), "Deadlock") {
						if deadlockInfo := database.PrintDeadlockInfo(db); deadlockInfo != "" {
							t.Logf("\n%s", deadlockInfo)
						}
					}
					errChan <- err
				}
			}(i)

			// Goroutine 3: File updates
			go func(iteration int) {
				defer wg.Done()
				time.Sleep(time.Duration(iteration%7) * time.Millisecond)

				updateCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
				defer cancel()

				fileIdx := iteration % len(allFileIDs)
				err := db.WithContext(updateCtx).Model(&model.File{}).
					Where("id = ?", allFileIDs[fileIdx]).
					Update("size", 2048).Error

				if err != nil && updateCtx.Err() == nil {
					errChan <- err
				}
			}(i)

			// Goroutine 4: Car block creation
			go func(iteration int) {
				defer wg.Done()
				time.Sleep(time.Duration(iteration%4) * time.Millisecond)

				createCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
				defer cancel()

				carIdx := iteration % len(allCarIDs)
				carID := allCarIDs[carIdx]
				block := model.CarBlock{
					CarID:          &carID,
					CID:            model.CID(testutil.TestCid),
					CarOffset:      int64(1000 + iteration),
					CarBlockLength: 512,
					Varint:         []byte{0x02},
					RawBlock:       []byte("new-block"),
					FileOffset:     int64(1000 + iteration),
				}
				err := db.WithContext(createCtx).Create(&block).Error

				// Ignore FK errors
				if err != nil && createCtx.Err() == nil && !isFKError(err) {
					errChan <- err
				}
			}(i)
		}

		done := make(chan struct{})
		go func() {
			wg.Wait()
			close(done)
		}()

		select {
		case <-done:
		case err := <-errChan:
			req.NoError(err, "Unexpected error during concurrent operations")
		case <-time.After(45 * time.Second):
			req.Fail("Test timed out - likely deadlock occurred")
		}

		// We don't verify complete cleanup because some deletions may be blocked by active jobs.
		t.Logf("Test completed without deadlock")
	})
}

func isNotFoundError(err error) bool {
	return err != nil && (strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "does not exist"))
}

func isActiveJobsError(err error) bool {
	return err != nil && strings.Contains(err.Error(), "active jobs")
}

func isJobsInUseError(err error) bool {
	return err != nil && strings.Contains(err.Error(), "jobs in use")
}

func isFKError(err error) bool {
	if err == nil {
		return false
	}
	s := err.Error()
	return strings.Contains(s, "foreign key") || strings.Contains(s, "FOREIGN KEY") ||
		strings.Contains(s, "violates") || strings.Contains(s, "constraint")
}
