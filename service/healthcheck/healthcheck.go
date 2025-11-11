package healthcheck

import (
	"context"
	"os"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/google/uuid"
	"github.com/ipfs/go-log/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var (
	staleThreshold = time.Minute * 5
	reportInterval = time.Minute
)

var cleanupInterval = time.Minute * 5

type State struct {
	JobType   model.JobType
	WorkingOn string
}

var logger = log.Logger("healthcheck")

// StartHealthCheckCleanup continuously runs the HealthCheckCleanup function
// at intervals specified by the cleanupInterval. The function cleans up resources
// or updates the health status of various components in your application, as
// implemented by HealthCheckCleanup.
//
// It is designed to be run as a background task and will continue to run
// until the passed context is cancelled.
//
// Parameters:
//   - ctx context.Context: The context that controls cancellations and timeouts.
//     When the context is cancelled (e.g. during shutdown), the function returns,
//     stopping its background cleaning task.
//   - db *gorm.DB: The database connection object used by HealthCheckCleanup to interact with
//     the database.
func StartHealthCheckCleanup(ctx context.Context, db *gorm.DB) {
	timer := time.NewTimer(cleanupInterval)
	defer timer.Stop()
	for {
		HealthCheckCleanup(ctx, db)
		select {
		case <-ctx.Done():
			return
		case <-timer.C:
			timer.Reset(cleanupInterval)
			continue
		}
	}
}

// HealthCheckCleanup is a function that cleans up stale workers and work files in the database.
//
// It first finds workers that haven't sent a heartbeat for a certain threshold (staleThreshold).
// Then it explicitly updates jobs owned by those workers to prevent CASCADE lock contention.
// Finally, it removes the stale workers.
//
// All database operations are retried on failure using the DoRetry function.
//
// Parameters:
//   - db: The Gorm DBNoContext connection to use for database queries.
func HealthCheckCleanup(ctx context.Context, db *gorm.DB) {
	db = db.WithContext(ctx)
	logger.Debugw("running healthcheck cleanup")

	// Clean up stale workers and their jobs in a single transaction to ensure consistency
	// and avoid FK CASCADE deadlocks by explicitly updating jobs first with SKIP LOCKED.
	err := database.DoRetry(ctx, func() error {
		return db.Transaction(func(tx *gorm.DB) error {
			// Find stale workers
			var staleWorkers []model.Worker
			err := tx.Where("last_heartbeat < ?", time.Now().UTC().Add(-staleThreshold)).Find(&staleWorkers).Error
			if err != nil {
				return errors.WithStack(err)
			}

			if len(staleWorkers) == 0 {
				return nil
			}

			staleWorkerIDs := make([]string, len(staleWorkers))
			for i, w := range staleWorkers {
				staleWorkerIDs[i] = w.ID
			}

			// Explicitly update jobs owned by stale workers with SKIP LOCKED to avoid deadlock
			// with concurrent bulk job updates. We use a two-step approach (SELECT then UPDATE)
			// to be compatible with both PostgreSQL and MySQL/MariaDB.
			var jobsToUpdate []model.Job
			err = tx.Clauses(clause.Locking{
				Strength: "UPDATE",
				Options:  "SKIP LOCKED",
			}).Select("id").
				Where("worker_id IN ? AND state = ?", staleWorkerIDs, model.Processing).
				Find(&jobsToUpdate).Error
			if err != nil {
				return errors.WithStack(err)
			}

			// Update the jobs we successfully locked
			if len(jobsToUpdate) > 0 {
				jobIDsToUpdate := make([]model.JobID, len(jobsToUpdate))
				for i, job := range jobsToUpdate {
					jobIDsToUpdate[i] = job.ID
				}

				err = tx.Model(&model.Job{}).
					Where("id IN ?", jobIDsToUpdate).
					Updates(map[string]any{
						"worker_id": nil,
						"state":     model.Ready,
					}).Error
				if err != nil {
					return errors.WithStack(err)
				}
			}

			// Only delete workers that have no remaining jobs with worker_id set
			// This avoids deadlock from FK CASCADE trying to lock jobs that are locked by other transactions
			for _, workerID := range staleWorkerIDs {
				var remainingJobCount int64
				err = tx.Model(&model.Job{}).
					Where("worker_id = ?", workerID).
					Count(&remainingJobCount).Error
				if err != nil {
					return errors.WithStack(err)
				}

				// Only delete if no jobs reference this worker anymore
				if remainingJobCount == 0 {
					err = tx.Where("id = ?", workerID).Delete(&model.Worker{}).Error
					if err != nil {
						return errors.WithStack(err)
					}
				}
				// If remainingJobCount > 0, skip this worker and let next cleanup attempt handle it
			}
			return nil
		})
	})
	if err != nil && !errors.Is(err, context.Canceled) {
		logger.Errorw("failed to cleanup dead workers", "error", err)
	}

	// Safety check: clean up any orphaned jobs that might have been missed
	// (e.g., jobs we couldn't lock due to SKIP LOCKED, or jobs left in inconsistent state)
	err = database.DoRetry(ctx, func() error {
		return db.Model(&model.Job{}).Where("worker_id NOT IN (?) AND state = ?",
			db.Table("workers").Select("id"), model.Processing).
			Updates(map[string]any{
				"worker_id": nil,
				"state":     model.Ready,
			}).Error
	})
	if err != nil && !errors.Is(err, context.Canceled) {
		logger.Errorw("failed to cleanup orphaned jobs", "error", err)
	}
}

// Register registers a new worker in the database. It uses the provided context and database connection.
// The workerID is used to uniquely identify the worker. The workerType is the type of the worker.
// If allowDuplicate is set to true, it allows the registration of duplicate workers.
//
// The function returns two values:
//   - alreadyRunning: A boolean indicating if the worker is already running.
//   - err: An error that will be nil if no errors occurred.
//
// The function first gets the hostname of the machine where it's running. If it fails to get the hostname, it returns an error.
// Then it gets the current state of the worker using the getState function.
// It then creates a new worker model with the provided workerID, the current time as the last heartbeat, the hostname, and the work type and working on values from the state.
//
// If allowDuplicate is set to false, the function checks if there are any active workers with the same work type and whose last heartbeat is not stale.
// If there are such workers, it sets alreadyRunning to true and returns.
//
// Finally, it tries to create the worker in the database. If it fails, it returns an error.
func Register(ctx context.Context, db *gorm.DB, workerID uuid.UUID, workerType model.WorkerType, allowDuplicate bool) (alreadyRunning bool, err error) {
	hostname, err := os.Hostname()
	if err != nil {
		return false, errors.WithStack(err)
	}
	worker := model.Worker{
		ID:            workerID.String(),
		LastHeartbeat: time.Now().UTC(),
		Hostname:      hostname,
		Type:          workerType,
	}
	logger.Debugw("registering worker", "worker", worker)
	err = database.DoRetry(ctx, func() error {
		if !allowDuplicate {
			var activeWorkerCount int64
			err := db.WithContext(ctx).Model(&model.Worker{}).Where("type = ? AND last_heartbeat > ?", workerType, time.Now().UTC().Add(-staleThreshold)).
				Count(&activeWorkerCount).Error
			if err != nil {
				return errors.Wrap(err, "failed to count active workers")
			}
			if activeWorkerCount > 0 {
				alreadyRunning = true
				return nil
			}
		}

		return db.WithContext(ctx).Create(&worker).Error
	})

	return alreadyRunning, errors.WithStack(err)
}

// ReportHealth reports the health of a worker to the database. It uses the provided context and database connection.
// The workerID is used to uniquely identify the worker.
//
// The function first gets the hostname of the machine where it's running. If it fails to get the hostname, it logs an error and returns.
// Then it gets the current state of the worker using the getState function.
// It then creates a new worker model with the provided workerID, the current time as the last heartbeat, the hostname, and the work type and working on values from the state.
//
// The function then tries to create the worker in the database or update the existing worker if one with the same ID already exists.
// The update will set the last heartbeat, work type, working on, and hostname fields to the values from the worker model.
// If the database operation fails, it logs an error.
func ReportHealth(ctx context.Context, db *gorm.DB, workerID uuid.UUID, workerType model.WorkerType) {
	hostname, err := os.Hostname()
	if err != nil {
		logger.Errorw("failed to get hostname", "error", err)
		return
	}
	worker := model.Worker{
		ID:            workerID.String(),
		LastHeartbeat: time.Now().UTC(),
		Hostname:      hostname,
		Type:          workerType,
	}
	logger.Debugw("sending heartbeat", "worker", worker)
	err = database.DoRetry(ctx, func() error {
		return db.WithContext(ctx).Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns([]string{"last_heartbeat", "type", "hostname"}),
		}).Create(&worker).Error
	})
	if err != nil {
		logger.Errorw("failed to send heartbeat", "error", err)
	}
}

// StartReportHealth continuously runs the ReportHealth function at intervals
// specified by the reportInterval. This function is responsible for reporting
// the health status of a worker to a centralized store (e.g., a database).
//
// The health status of the worker is determined by calling the provided getState
// function, which should return the current state of the worker.
//
// This function is designed to be run as a background task and will continue to
// run until the passed context is cancelled.
//
// Parameters:
//   - ctx context.Context: The context that controls cancellations and timeouts.
//     When the context is cancelled (e.g. during shutdown), the function returns,
//     stopping its background health reporting task.
//   - db *gorm.DB: The database connection object used by ReportHealth to interact with
//     the database.
//   - workerID uuid.UUID: The unique identifier for the worker whose health is being reported.
func StartReportHealth(ctx context.Context, db *gorm.DB, workerID uuid.UUID, workerType model.WorkerType) {
	timer := time.NewTimer(reportInterval)
	defer timer.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-timer.C:
			ReportHealth(ctx, db, workerID, workerType)
			timer.Reset(reportInterval)
		}
	}
}
