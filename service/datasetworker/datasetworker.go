package datasetworker

import (
	"context"
	"fmt"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/metrics"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/service"
	"github.com/data-preservation-programs/singularity/service/healthcheck"
	"github.com/google/uuid"
	"github.com/ipfs/go-log/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var logger = log.Logger("datasetworker")

type Worker struct {
	dbNoContext *gorm.DB
	config      Config
}

type Config struct {
	Concurrency    int
	ExitOnComplete bool
	EnableScan     bool
	EnablePack     bool
	EnableDag      bool
	ExitOnError    bool
}

func NewWorker(db *gorm.DB, config Config) *Worker {
	return &Worker{
		dbNoContext: db,
		config:      config,
	}
}

type Thread struct {
	id          uuid.UUID
	dbNoContext *gorm.DB
	logger      *zap.SugaredLogger
	config      Config
}

// Start initializes and starts the execution of a worker thread.
// This function:
// 1. Creates a cancellable context derived from the input context.
// 2. Registers the worker with a health check service, providing a state function for reporting its status.
// 3. Launches separate goroutines to report health status, clean up old health check records, execute the worker's task, and handle cleanup.
// 4. Returns channels that are closed when the health reporting, health check cleanup, worker execution, and worker cleanup are complete.
//
// Parameters:
//
//	ctx : The parent context for this thread, used to propagate cancellations.
//
// Returns:
//
//	[]service.Done : A slice of channels that are closed when respective components of the worker complete their execution.
//	service.Fail   : A channel that receives an error if the worker encounters a failure during its execution.
//	error          : An error is returned if the worker fails to register with the health check service. Otherwise, it returns nil.
func (w *Thread) Start(ctx context.Context) ([]service.Done, service.Fail, error) {
	var cancel context.CancelFunc
	ctx, cancel = context.WithCancel(ctx)

	_, err := healthcheck.Register(ctx, w.dbNoContext, w.id, model.DatasetWorker, true)
	if err != nil {
		cancel()
		return nil, nil, errors.Wrap(err, "failed to register worker")
	}

	healthcheckDone := make(chan struct{})
	go func() {
		defer close(healthcheckDone)
		healthcheck.StartReportHealth(ctx, w.dbNoContext, w.id, model.DatasetWorker)
		w.logger.Info("health report stopped")
	}()

	healthcheckCleanupDone := make(chan struct{})
	go func() {
		defer close(healthcheckCleanupDone)
		healthcheck.StartHealthCheckCleanup(ctx, w.dbNoContext)
		w.logger.Info("healthcheck cleanup stopped")
	}()

	done := make(chan struct{})
	fail := make(chan error)
	go func() {
		defer cancel()
		defer close(done)
		w.run(ctx, fail)
		w.logger.Info("worker thread finished")
	}()

	cleanupDone := make(chan struct{})
	go func() {
		defer close(cleanupDone)
		<-ctx.Done()
		ctx2, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		//nolint:contextcheck
		err := w.cleanup(ctx2)
		if err != nil {
			w.logger.Errorw("failed to cleanup", "error", err)
		} else {
			w.logger.Info("cleanup complete")
		}
	}()

	return []service.Done{done, healthcheckDone, healthcheckCleanupDone, cleanupDone}, fail, nil
}

func (w *Thread) Name() string {
	return "Preparation Worker Thread - " + w.id.String()
}

func (w *Thread) cleanup(ctx context.Context) error {
	return database.DoRetry(ctx, func() error {
		return w.dbNoContext.WithContext(ctx).Where("id = ?", w.id.String()).Delete(&model.Worker{}).Error
	})
}

// Run initializes and starts a set of worker threads based on the Concurrency specified in the configuration.
// This function:
// 1. Creates an array of worker threads, each having a unique identifier.
// 2. Initializes each thread with a shared set of dependencies (e.g., database, logger) and individual configuration.
// 3. Invokes the StartServers function to run all the threads, passing the initialized threads and a logger.
//
// Parameters:
//
//	ctx : The context under which all the worker threads are run, used to propagate cancellations.
//
// Returns:
//
//	error : An error is returned if the StartServers function encounters an issue while starting the threads. Otherwise, it returns nil.
func (w Worker) Run(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	err := metrics.Init(ctx, w.dbNoContext)
	if err != nil {
		return errors.WithStack(err)
	}

	metricsFlushed := make(chan struct{})
	go func() {
		defer close(metricsFlushed)
		metrics.Default.Start(ctx)
		//nolint:contextcheck
		metrics.Default.Flush()
	}()

	threads := make([]service.Server, w.config.Concurrency)
	for i := 0; i < w.config.Concurrency; i++ {
		id := uuid.New()
		thread := &Thread{
			id:          id,
			dbNoContext: w.dbNoContext,
			logger:      logger.With("workerID", id.String()),
			config:      w.config,
		}
		threads[i] = thread
	}
	err = service.StartServers(ctx, logger, threads...)
	cancel()
	<-metricsFlushed
	return errors.WithStack(err)
}

func (w Worker) Name() string {
	return "Preparation Worker Main"
}

func (w *Thread) handleWorkComplete(ctx context.Context, jobID uint64) error {
	return database.DoRetry(ctx, func() error {
		return w.dbNoContext.WithContext(ctx).Model(&model.Job{}).Where("id = ?", jobID).Updates(map[string]any{
			"worker_id":         nil,
			"error_message":     "",
			"error_stack_trace": "",
			"state":             model.Complete,
		}).Error
	})
}

func (w *Thread) handleWorkError(ctx context.Context, jobID uint64, err error) error {
	updates := make(map[string]any)
	updates["worker_id"] = nil
	// Reset the state to ready if the context was canceled
	if errors.Is(err, context.Canceled) {
		updates["error_message"] = ""
		updates["error_stack_trace"] = ""
		updates["state"] = model.Ready
		var cancel context.CancelFunc
		//nolint:contextcheck
		ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
	} else {
		updates["error_message"] = err.Error()
		updates["error_stack_trace"] = fmt.Sprintf("%+v", err)
		updates["state"] = model.Error
	}
	return database.DoRetry(ctx, func() error {
		return w.dbNoContext.WithContext(ctx).Model(&model.Job{}).Where("id = ?", jobID).Updates(updates).Error
	})
}

// run is the core loop that a Thread executes when started.
// It continually looks for work to process, handles errors, and reports updates:
// 1. It attempts to find work to do. The types of work are defined by WorkType enumeration (e.g., Scan, Pack, Dag).
// 2. It processes the found work based on its type, reporting errors if they occur.
// 3. If an error occurs, it either exits or waits for a minute before looking for more work, based on the configuration.
// 4. If no work is found, it either exits or waits for 15 seconds before looking for more work, based on the configuration.
// 5. It gracefully stops if the provided context is cancelled.
//
// Parameters:
//
//	ctx     : The context used for managing the lifecycle of the run loop. If Done, the loop exits cleanly.
//	errChan : A channel for reporting errors that cause the loop to exit when ExitOnError is true.
//
// This function is intended to run as a goroutine.
//
// In case of a panic, the error is recovered and sent to the errChan.
func (w *Thread) run(ctx context.Context, errChan chan error) {
	defer func() {
		if err := recover(); err != nil {
			errChan <- errors.Errorf("panic: %v", err)
		}
	}()

	var jobTypes []model.JobType
	if w.config.EnableDag {
		jobTypes = append(jobTypes, model.DagGen)
	}
	if w.config.EnableScan {
		jobTypes = append(jobTypes, model.Scan)
	}
	if w.config.EnablePack {
		jobTypes = append(jobTypes, model.Pack)
	}
	minInterval := 5 * time.Second
	maxInterval := 160 * time.Second
	interval := minInterval
	for {
		job, err := w.findJob(ctx, jobTypes)
		if err != nil {
			goto errorLoop
		}

		if job == nil {
			if w.config.ExitOnComplete {
				w.logger.Info("no work found, exiting")
				return
			}
			w.logger.Info("no work found")
			goto loop
		}

		switch job.Type {
		case model.Scan:
			err = w.scan(ctx, *job.Attachment)
		case model.Pack:
			err = w.pack(ctx, *job)
		case model.DagGen:
			err = w.ExportDag(ctx, *job)
		}
		if err != nil {
			err2 := w.handleWorkError(ctx, job.ID, err)
			if err2 != nil {
				w.logger.Errorw("failed to update state to error",
					"type", job.Type, "jobID", job.ID, "error", err2)
			}
			goto errorLoop
		} else {
			err2 := w.handleWorkComplete(ctx, job.ID)
			if err2 != nil {
				w.logger.Errorw("failed to update state to complete",
					"type", job.Type, "jobID", job.ID, "error", err2)
				goto errorLoop
			}
			interval = minInterval
			continue
		}
	errorLoop:
		if ctx.Err() != nil {
			w.logger.Info("context cancelled, exiting")
			return
		}
		w.logger.Errorw("error encountered", "error", err)
		if w.config.ExitOnError {
			select {
			case errChan <- err:
			case <-ctx.Done():
			}
			return
		}
	loop:
		w.logger.Infof("sleeping for %s", interval)
		select {
		case <-ctx.Done():
			return
		case <-time.After(interval):
			interval *= 2
			if interval > maxInterval {
				interval = maxInterval
			}
		}
	}
}
