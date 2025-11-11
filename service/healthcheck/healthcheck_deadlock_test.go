package healthcheck

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

// TestHealthCheckCleanupNoDeadlock verifies that concurrent worker cleanup and bulk job
// updates do not deadlock. Worker cleanup triggers FK CASCADE to jobs while bulk updates
// modify jobs directly. The fix uses FOR UPDATE SKIP LOCKED.
func TestHealthCheckCleanupNoDeadlock(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		req := require.New(t)

		// Create test preparation and storage for jobs
		preparation := model.Preparation{
			Name:      "test-prep-" + uuid.New().String(),
			MaxSize:   1 << 30,
			PieceSize: 1 << 30,
		}
		err := db.Create(&preparation).Error
		req.NoError(err)

		storage := model.Storage{
			Name: "test-storage-" + uuid.New().String(),
			Type: "local",
			Path: "/tmp/test",
		}
		err = db.Create(&storage).Error
		req.NoError(err)

		attachment := model.SourceAttachment{
			PreparationID: preparation.ID,
			StorageID:     storage.ID,
		}
		err = db.Create(&attachment).Error
		req.NoError(err)

		// Create workers and jobs assigned to them
		const numWorkers = 5
		const numJobsPerWorker = 10
		workerIDs := make([]string, numWorkers)
		jobIDs := make([]model.JobID, 0, numWorkers*numJobsPerWorker)

		for i := range numWorkers {
			workerID := uuid.New().String()
			workerIDs[i] = workerID

			// Create worker with old heartbeat
			worker := model.Worker{
				ID:            workerID,
				LastHeartbeat: time.Now().UTC().Add(-10 * time.Minute),
				Hostname:      "test-host",
				Type:          model.DatasetWorker,
			}
			err = db.Create(&worker).Error
			req.NoError(err)

			// Create jobs assigned to this worker
			for range numJobsPerWorker {
				job := model.Job{
					Type:         model.Pack,
					State:        model.Processing,
					WorkerID:     &workerID,
					AttachmentID: attachment.ID,
				}
				err = db.Create(&job).Error
				req.NoError(err)
				jobIDs = append(jobIDs, job.ID)
			}
		}

		const numIterations = 20
		var wg sync.WaitGroup
		errChan := make(chan error, numIterations*2)

		for i := range numIterations {
			wg.Add(2)

			// Goroutine 1: Worker cleanup
			go func(iteration int) {
				defer wg.Done()
				time.Sleep(time.Duration(iteration%3) * time.Millisecond)

				cleanupCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
				defer cancel()

				HealthCheckCleanup(cleanupCtx, db)
			}(i)

			// Goroutine 2: Bulk job update
			go func(iteration int) {
				defer wg.Done()
				time.Sleep(time.Duration(iteration%5) * time.Millisecond)

				updateCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
				defer cancel()

				startIdx := (iteration * 5) % len(jobIDs)
				endIdx := min(startIdx+10, len(jobIDs))
				batchJobIDs := jobIDs[startIdx:endIdx]

				err := db.WithContext(updateCtx).Transaction(func(tx *gorm.DB) error {
					for _, jobID := range batchJobIDs {
						err := tx.Model(&model.Job{}).
							Where("id = ?", jobID).
							Updates(map[string]any{
								"state":             model.Ready,
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
		case <-time.After(30 * time.Second):
			req.Fail("Test timed out - likely deadlock occurred")
		}

		// Verify workers deleted
		var remainingWorkers []model.Worker
		err = db.Where("id IN ?", workerIDs).Find(&remainingWorkers).Error
		req.NoError(err)
		req.Empty(remainingWorkers, "All stale workers should be deleted")

		// Verify jobs reset
		var remainingJobs []model.Job
		err = db.Where("id IN ?", jobIDs).Find(&remainingJobs).Error
		req.NoError(err)

		for _, job := range remainingJobs {
			req.Nil(job.WorkerID, "Job %d should have NULL worker_id", job.ID)
			req.Contains([]model.JobState{model.Ready, model.Complete}, job.State,
				"Job %d should be in Ready or Complete state", job.ID)
		}
	})
}
