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
// It first removes all workers that haven't sent a heartbeat for a certain threshold (staleThreshold).
// If there's an error removing the workers, it logs the error and continues.
//
// Then, it resets the state of any jobs that are marked as being processed by a worker that no longer exists.
// If there's an error updating the sources, it logs the error and continues.
//
// All database operations are retried on failure using the DoRetry function.
//
// Parameters:
//   - db: The Gorm DBNoContext connection to use for database queries.
func HealthCheckCleanup(ctx context.Context, db *gorm.DB) {
	db = db.WithContext(ctx)
	logger.Debugw("running healthcheck cleanup")
	// Remove all workers that haven't sent heartbeat for 5 minutes.
	err := database.DoRetry(ctx, func() error {
		return db.Where("last_heartbeat < ?", time.Now().UTC().Add(-staleThreshold)).Delete(&model.Worker{}).Error
	})
	if err != nil && !errors.Is(err, context.Canceled) {
		log.Logger("healthcheck").Errorw("failed to remove dead workers", "error", err)
	}

	// In case there are some works that have stale foreign key referenced to dead workers, we need to remove them
	err = database.DoRetry(ctx, func() error {
		return db.Model(&model.Job{}).Where("(worker_id NOT IN (?) OR worker_id IS NULL) AND state = ?",
			db.Table("workers").Select("id"), model.Processing).
			Updates(map[string]any{
				"worker_id": nil,
				"state":     model.Ready,
			}).Error
	})
	if err != nil && !errors.Is(err, context.Canceled) {
		log.Logger("healthcheck").Errorw("failed to remove stale workers", "error", err)
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
