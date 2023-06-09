package datasetworker

import (
	"context"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/service/healthcheck"
	"github.com/rjNemo/underscore"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/data-preservation-programs/singularity/datasource"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/google/uuid"
	"github.com/ipfs/go-log/v2"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var logger = log.Logger("datasetworker")

type DatasetWorker struct {
	db      *gorm.DB
	threads []DatasetWorkerThread
	config  DatasetWorkerConfig
}

type DatasetWorkerConfig struct {
	Concurrency    int
	ExitOnComplete bool
	EnableScan     bool
	EnablePack     bool
	EnableDag      bool
}

func NewDatasetWorker(db *gorm.DB, config DatasetWorkerConfig) *DatasetWorker {
	return &DatasetWorker{
		db:      db,
		threads: make([]DatasetWorkerThread, config.Concurrency),
		config:  config,
	}
}

type DatasetWorkerThread struct {
	id                        uuid.UUID
	db                        *gorm.DB
	logger                    *zap.SugaredLogger
	directoryCache            map[string]model.Directory
	workType                  model.WorkType
	workingOn                 string
	datasourceHandlerResolver datasource.HandlerResolver
	config                    DatasetWorkerConfig
}

func (w DatasetWorker) cleanup() error {
	workerIDs := make([]string, len(w.threads))
	for i, thread := range w.threads {
		workerIDs[i] = thread.id.String()
	}
	return w.db.Where("id IN ?", workerIDs).Delete(&model.Worker{}).Error
}

func (w DatasetWorker) Run(parent context.Context) error {
	ctx, cancel := context.WithCancel(parent)
	defer cancel()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM, syscall.SIGTRAP)
	errChan := make(chan error)

	go healthcheck.StartHealthCheckCleanup(ctx, w.db)

	var wg sync.WaitGroup
	for i := 0; i < w.config.Concurrency; i++ {
		wg.Add(1)
		id := uuid.New()
		thread := DatasetWorkerThread{
			id:                        id,
			db:                        w.db.WithContext(ctx),
			logger:                    logger.With("workerID", id.String()),
			datasourceHandlerResolver: datasource.DefaultHandlerResolver{},
			config:                    w.config,
		}
		w.threads[i] = thread
		go thread.run(ctx, errChan, &wg)
	}

	done := make(chan struct{})

	go func() {
		wg.Wait()
		done <- struct{}{}
	}()

	select {
	case <-done:
		log.Logger("worker").Info("All work done, cleaning up")
		w.cleanup()
		return nil
	case <-signalChan:
		log.Logger("worker").Info("received signal, cleaning up")
		w.cleanup()
		return cli.Exit("received signal", 130)
	case err := <-errChan:
		log.Logger("worker").Errorw("one of the worker thread encountered unrecoverable error", "error", err)
		w.cleanup()
		return cli.Exit("worker thread failed", 1)
	}
}

func (w *DatasetWorkerThread) getState() healthcheck.State {
	return healthcheck.State{
		WorkType:  w.workType,
		WorkingOn: w.workingOn,
	}
}

func (w *DatasetWorkerThread) run(ctx context.Context, errChan chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()
	defer func() {
		if err := recover(); err != nil {
			errChan <- errors.Errorf("panic: %v", err)
		}
	}()
	healthcheck.HealthCheck(w.db, w.id, w.getState)
	go healthcheck.StartHealthCheck(ctx, w.db, w.id, w.getState)
	for {
		w.directoryCache = map[string]model.Directory{}
		// 1st, find scanning work
		source, err := w.findScanWork()
		if err != nil {
			w.logger.Errorw("failed to scan", "error", err)
			goto errorLoop
		}
		if source != nil {
			err = w.scan(ctx, *source, source.Type != "manual")
			if err != nil {
				w.logger.Errorw("failed to scan", "error", err)
				newState := model.Error
				newErrorMessage := err.Error()
				if errors.Is(err, context.Canceled) {
					newState = model.Ready
					newErrorMessage = ""
					cancelCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
					defer cancel()
					w.db = w.db.WithContext(cancelCtx)
				}
				err = database.DoRetry(func() error {
					return w.db.Model(&model.Source{}).Where("id = ?", source.ID).Updates(
						map[string]interface{}{
							"scanning_state":         newState,
							"scanning_worker_id":     nil,
							"error_message":          newErrorMessage,
							"last_scanned_timestamp": time.Now().UTC().Unix(),
						},
					).Error
				})
				if err != nil {
					w.logger.Errorw("failed to update source with error", "error", err)
				}
				goto errorLoop
			}

			err = database.DoRetry(func() error {
				return w.db.Model(&model.Source{}).Where("id = ?", source.ID).Updates(
					map[string]interface{}{
						"scanning_state":         model.Complete,
						"scanning_worker_id":     nil,
						"last_scanned_timestamp": time.Now().UTC().Unix(),
						"last_scanned_path":      "",
					},
				).Error
			})
			if err != nil {
				w.logger.Errorw("failed to update source to complete", "error", err)
				goto errorLoop
			}
			continue
		}

		// 2nd, find packing work
		{
			chunk, err := w.findPackWork()
			if err != nil {
				w.logger.Errorw("failed to find pack work", "error", err)
				goto errorLoop
			}
			if chunk != nil {
				err = w.pack(
					ctx,
					*chunk,
				)
				if err != nil {
					w.logger.Errorw("failed to pack", "error", err)
					newState := model.Error
					newErrorMessage := err.Error()
					if errors.Is(err, context.Canceled) {
						newState = model.Ready
						newErrorMessage = ""
						cancelCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
						defer cancel()
						w.db = w.db.WithContext(cancelCtx)
					}
					err = database.DoRetry(func() error {
						return w.db.Model(&model.Chunk{}).Where("id = ?", chunk.ID).Updates(
							map[string]interface{}{
								"packing_state":     newState,
								"packing_worker_id": nil,
								"error_message":     newErrorMessage,
							},
						).Error
					})
					if err != nil {
						w.logger.Errorw("failed to update chunk with error", "error", err)
					}
					goto errorLoop
				}
				err = database.DoRetry(func() error {
					return w.db.Model(&model.Chunk{}).Where("id = ?", chunk.ID).Updates(
						map[string]interface{}{
							"packing_state":     model.Complete,
							"packing_worker_id": nil,
						},
					).Error
				})
				if err != nil {
					w.logger.Errorw("failed to update chunk to complete", "error", err)
					goto errorLoop
				}
				continue
			}
		}
		if w.config.ExitOnComplete {
			w.logger.Debug("exiting on complete")
			return
		}
	errorLoop:
		w.logger.Debug("sleeping for a minute")
		select {
		case <-ctx.Done():
			return
		case <-time.After(time.Minute):
		}
	}
}

type remain struct {
	itemParts []model.ItemPart
	carSize   int64
}

const carHeaderSize = 59

func newRemain() *remain {
	return &remain{
		itemParts: make([]model.ItemPart, 0),
		// Some buffer for header
		carSize: carHeaderSize,
	}
}

func (r *remain) add(itemParts []model.ItemPart) {
	r.itemParts = append(r.itemParts, itemParts...)
	for _, itemPart := range itemParts {
		r.carSize += toCarSize(itemPart.Length)
	}
}

func (r *remain) reset() {
	r.itemParts = make([]model.ItemPart, 0)
	r.carSize = carHeaderSize
}

func (r *remain) itemIDs() []uint64 {
	return underscore.Map(r.itemParts, func(itemPart model.ItemPart) uint64 {
		return itemPart.ID
	})
}

func toCarSize(size int64) int64 {
	out := size
	nBlocks := size / 1024 / 1024
	if size%(1024*1024) != 0 {
		nBlocks++
	}

	// For each block, we need to add the bytes for the CID as well as varint
	out += nBlocks * (36 + 9)

	// For every 256 blocks, we need to add another block.
	// The block stores up to 256 CIDs and integers, estimate it to be 12kb
	if nBlocks > 1 {
		out += (((nBlocks - 1) / 256) + 1) * 12000
	}

	return out
}
