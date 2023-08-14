package datasetworker

import (
	"context"
	"time"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/datasource"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/service"
	"github.com/data-preservation-programs/singularity/service/healthcheck"
	"github.com/google/uuid"
	"github.com/ipfs/go-log/v2"
	"github.com/pkg/errors"
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
	id                        uuid.UUID
	dbNoContext               *gorm.DB
	logger                    *zap.SugaredLogger
	workType                  model.WorkType
	workingOn                 string
	datasourceHandlerResolver datasource.HandlerResolver
	config                    Config
}

func (w *Thread) Start(ctx context.Context) ([]service.Done, service.Fail, error) {
	var cancel context.CancelFunc
	ctx, cancel = context.WithCancel(ctx)
	getState := func() healthcheck.State {
		return healthcheck.State{
			WorkType:  w.workType,
			WorkingOn: w.workingOn,
		}
	}

	_, err := healthcheck.Register(ctx, w.dbNoContext, w.id, getState, true)
	if err != nil {
		cancel()
		return nil, nil, errors.Wrap(err, "failed to register worker")
	}

	healthcheckDone := make(chan struct{})
	go func() {
		defer close(healthcheckDone)
		healthcheck.StartReportHealth(ctx, w.dbNoContext, w.id, getState)
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
	return "Dataset Worker Thread - " + w.id.String()
}

func (w *Thread) cleanup(ctx context.Context) error {
	return database.DoRetry(ctx, func() error {
		return w.dbNoContext.WithContext(ctx).Where("id = ?", w.id.String()).Delete(&model.Worker{}).Error
	})
}

func (w Worker) Run(ctx context.Context) error {
	threads := make([]service.Server, w.config.Concurrency)
	for i := 0; i < w.config.Concurrency; i++ {
		id := uuid.New()
		thread := &Thread{
			id:                        id,
			dbNoContext:               w.dbNoContext,
			logger:                    logger.With("workerID", id.String()),
			datasourceHandlerResolver: datasource.DefaultHandlerResolver{},
			config:                    w.config,
		}
		threads[i] = thread
	}
	return service.StartServers(ctx, logger, threads...)
}

type WorkType string

const (
	WorkTypeNone WorkType = ""
	WorkTypeScan WorkType = "scan"
	WorkTypePack WorkType = "pack"
	WorkTypeDag  WorkType = "dag"
)

var WorkStateKey = map[WorkType]string{
	WorkTypeScan: "scanning_state",
	WorkTypePack: "packing_state",
	WorkTypeDag:  "dag_gen_state",
}

var WorkerIDKey = map[WorkType]string{
	WorkTypeScan: "scanning_worker_id",
	WorkTypePack: "packing_worker_id",
	WorkTypeDag:  "dag_gen_worker_id",
}

var ErrorMessageKey = map[WorkType]string{
	WorkTypeScan: "error_message",
	WorkTypePack: "error_message",
	WorkTypeDag:  "dag_gen_error_message",
}

var WorkModel = map[WorkType]func() any{
	WorkTypeScan: func() any { return &model.Source{} },
	WorkTypePack: func() any { return &model.Chunk{} },
	WorkTypeDag:  func() any { return &model.Source{} },
}

func (w *Thread) handleWorkComplete(ctx context.Context, workType WorkType, id uint64, updates map[string]any) error {
	if workType == WorkTypeNone {
		return nil
	}
	w.logger.Infow("finished "+string(workType), "id", id)
	updates[WorkerIDKey[workType]] = nil
	updates[ErrorMessageKey[workType]] = ""
	updates[WorkStateKey[workType]] = model.Complete
	return database.DoRetry(ctx, func() error {
		return w.dbNoContext.WithContext(ctx).Model(WorkModel[workType]()).Where("id = ?", id).Updates(updates).Error
	})
}

func (w *Thread) handleWorkError(ctx context.Context, workType WorkType, id uint64, err error) error {
	if err == nil || workType == WorkTypeNone {
		return nil
	}
	w.logger.Errorw("failed to "+string(workType), "id", id, "error", err)
	updates := make(map[string]any)
	updates[WorkerIDKey[workType]] = nil
	// reset the state to ready if the context was canceled
	if errors.Is(err, context.Canceled) {
		updates[ErrorMessageKey[workType]] = ""
		updates[WorkStateKey[workType]] = model.Ready
		var cancel context.CancelFunc
		//nolint:contextcheck
		ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
	} else {
		updates[ErrorMessageKey[workType]] = err.Error()
		updates[WorkStateKey[workType]] = model.Error
	}
	return database.DoRetry(ctx, func() error {
		return w.dbNoContext.WithContext(ctx).Model(WorkModel[workType]()).Where("id = ?", id).Updates(updates).Error
	})
}

func (w *Thread) findWork(ctx context.Context) (WorkType, *model.Source, *model.Chunk, error) {
	source, err := w.findDagWork(ctx)
	if err != nil {
		return "", nil, nil, errors.Wrap(err, "failed to find dag work")
	}
	if source != nil {
		return WorkTypeDag, source, nil, nil
	}

	source, err = w.findScanWork(ctx)
	if err != nil {
		return "", nil, nil, errors.Wrap(err, "failed to find scan work")
	}
	if source != nil {
		return WorkTypeScan, source, nil, nil
	}

	chunk, err := w.findPackWork(ctx)
	if err != nil {
		return "", nil, nil, errors.Wrap(err, "failed to find pack work")
	}
	if chunk != nil {
		return WorkTypePack, nil, chunk, nil
	}

	return WorkTypeNone, nil, nil, nil
}

func (w *Thread) run(ctx context.Context, errChan chan error) {
	defer func() {
		if err := recover(); err != nil {
			errChan <- errors.Errorf("panic: %v", err)
		}
	}()
	for {
		var id uint64
		workType, source, chunk, err := w.findWork(ctx)
		if err != nil {
			goto errorLoop
		}

		switch workType {
		case WorkTypeNone:
			if w.config.ExitOnComplete {
				w.logger.Info("no work found, exiting")
				return
			}
			select {
			case <-ctx.Done():
				return
			case <-time.After(15 * time.Second):
				continue
			}
		case WorkTypeScan:
			err = w.scan(ctx, *source, source.Type != "manual")
			id = uint64(source.ID)
		case WorkTypePack:
			err = w.pack(ctx, *chunk)
			id = uint64(chunk.ID)
		case WorkTypeDag:
			err = w.dag(ctx, *source)
			id = uint64(source.ID)
		}
		if err != nil {
			err = w.handleWorkError(ctx, workType, id, err)
			if err != nil {
				w.logger.Errorw("failed to update state to error",
					"type", workType, "id", id, "error", err)
			}
			goto errorLoop
		} else {
			updates := make(map[string]any)
			if workType == WorkTypeScan {
				updates["last_scanned_timestamp"] = time.Now().UTC().Unix()
				updates["last_scanned_path"] = ""
			}
			err = w.handleWorkComplete(ctx, workType, id, updates)
			if err != nil {
				w.logger.Errorw("failed to update state to complete",
					"type", workType, "id", id, "error", err)
				goto errorLoop
			}
			continue
		}
	errorLoop:
		if w.config.ExitOnError {
			errChan <- err
			return
		}
		w.logger.Info("sleeping for a minute")
		select {
		case <-ctx.Done():
			return
		case <-time.After(time.Minute):
		}
	}
}
