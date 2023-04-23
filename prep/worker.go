package prep

import (
	"context"
	"github.com/data-preservation-programs/go-singularity/model"
	"github.com/data-preservation-programs/go-singularity/source"
	"github.com/data-preservation-programs/go-singularity/source/fssource"
	"github.com/google/uuid"
	"github.com/ipfs/go-log/v2"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Worker struct {
	db          *gorm.DB
	concurrency int
}

func NewWorker(db *gorm.DB, concurrency int) *Worker {
	return &Worker{db: db, concurrency: concurrency}
}

type WorkerThread struct {
	id             string
	db             *gorm.DB
	logger         *zap.SugaredLogger
	sourceHandlers map[model.SourceType]source.DataSource
}

func (w Worker) Cleanup(workerIds []string) error {
	return w.db.Where("id IN ?", workerIds).Delete(&model.Worker{}).Error
}

func (w Worker) Run(parent context.Context) error {
	ctx, cancel := context.WithCancel(parent)
	defer cancel()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	errChan := make(chan error)
	sourceHandlers := map[model.SourceType]source.DataSource{
		model.Dir: fssource.DirSource{},
	}
	workerIds := make([]string, w.concurrency)
	for i := 0; i < w.concurrency; i++ {
		id := uuid.NewString()
		workerIds[i] = id
		thread := WorkerThread{
			id:             id,
			db:             w.db.WithContext(ctx),
			logger:         log.Logger("prep-worker").With("worker", id),
			sourceHandlers: sourceHandlers,
		}
		go thread.run(ctx, errChan)
	}

	select {
	case <-signalChan:
		log.Logger("prep-worker").Info("received signal, cleaning up")
		w.Cleanup(workerIds)
		return cli.Exit("received signal", 130)
	case err := <-errChan:
		log.Logger("prep-worker").Errorw("one of the worker thread failed", "error", err)
		w.Cleanup(workerIds)
		return cli.Exit("worker thread failed", 1)
	}
}

func (w WorkerThread) run(ctx context.Context, errChan chan<- error) {
	go model.StartHealthCheck(ctx, w.db, w.id)
	for {
		// 1st, find scanning work
		source, err := w.findScanWork(ctx)
		if err != nil {
			w.logger.Errorw("failed to scan", "error", err)
			time.Sleep(time.Minute)
			continue
		}
		if source != nil {
			err = w.scan(ctx, *source)
			if err != nil {
				w.logger.Errorw("failed to scan", "error", err)
				time.Sleep(time.Minute)
				continue
			}
			continue
		}

		// 3rd, find chunking work
	}
}

func (w WorkerThread) updateInactive() error {
	result := w.db.Model(&model.Source{}).
		Where("scanning_state = ?", model.Processing).
		Where("scanning_worker_id = ?", w.id).
		Updates(map[string]interface{}{
			"scanning_state": model.Ready,
		})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (w WorkerThread) findScanWork(ctx context.Context) (*model.Source, error) {
	source := model.Source{}
	// Find all ready source (never scanned before)
	result := w.db.Model(&model.Source{}).
		Where("scanning_state = ?", model.Ready).
		Order("id asc").
		Limit(1).
		Updates(map[string]interface{}{
			"scanning_state":     model.Processing,
			"scanning_worker_id": w.id,
			"error_message":      "",
		}).First(&source)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	if result.Error != nil {
		w.logger.With("source", source).Info("found source to scan")
		return &source, nil
	}

	// Find all source that is complete but needs rescanning
	result = w.db.Model(&model.Source{}).
		Where("scanning_state = ? AND scan_interval != 0 AND last_scanned + scan_interval < ?",
			model.Complete, time.Now().UTC()).
		Order("id asc").
		Limit(1).
		Updates(map[string]interface{}{
			"scanning_state":     model.Processing,
			"scanning_worker_id": w.id,
			"error_message":      "",
		}).First(&source)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	if result.Error != nil {
		w.logger.With("source", source).Info("found source to scan")
		return &source, nil
	}

	return nil, nil
}

func (w WorkerThread) scan(ctx context.Context, source model.Source) error {
	lastItem := model.Item{}
	result := w.db.Where("source_id = ?", source.ID).Order("scanned_at desc").Take(&lastItem)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}
	w.logger.With("lastItem", lastItem).Info("last item scanned")

	lastPath := lastItem.Path
	sourceHandler := w.sourceHandlers[source.Type]
	entryChan, err := sourceHandler.Scan(ctx, source.Path, lastPath)
	if err != nil {
		w.logger.Errorw("failed to scan source", "error", err)
		source.ScanningState = model.Error
		source.ScanningWorkerID = ""
		source.ErrorMessage = err.Error()
		source.LastScanned = time.Now().UTC()
		err = w.db.Save(&source).Error
		if err != nil {
			return errors.Wrap(err, "failed to update source with error")
		}
		return errors.Wrap(err, "failed to scan source")
	}

	for entry := range entryChan {
		item := model.Item{
			ScannedAt:    entry.ScannedAt,
			DatasetID:    source.DatasetID,
			SourceID:     source.ID,
			Type:         entry.Type,
			Path:         entry.Path,
			Size:         entry.Size,
			LastModified: entry.LastModified,
			Version:      0,
		}
		w.logger.Debugw("found item", "item", item)
		err = w.db.Create(&item).Error
		if err != nil {
			return errors.Wrap(err, "failed to create item")
		}
	}

	return nil
}
