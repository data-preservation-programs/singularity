package prep

import (
	"context"
	"github.com/data-preservation-programs/go-singularity/datasource"
	"github.com/data-preservation-programs/go-singularity/model"
	"github.com/data-preservation-programs/go-singularity/pack"
	"github.com/google/uuid"
	"github.com/ipfs/go-log/v2"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
	"os/signal"
	"strings"
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
	directoryCache map[string]model.Directory
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
	workerIds := make([]string, w.concurrency)
	for i := 0; i < w.concurrency; i++ {
		id := uuid.NewString()
		workerIds[i] = id
		thread := WorkerThread{
			id:             id,
			db:             w.db.WithContext(ctx),
			logger:         log.Logger("prep-worker").With("worker", id),
			directoryCache: map[string]model.Directory{},
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

func (w *WorkerThread) run(ctx context.Context, errChan chan<- error) {
	go model.StartHealthCheck(ctx, w.db, w.id)
	for {
		w.directoryCache = map[string]model.Directory{}
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

		// 2nd, find ipld work
		// 3rd, find packing work
		chunk, err := w.findPackWork(ctx)
		if err != nil {
			w.logger.Errorw("failed to find pack work", "error", err)
			time.Sleep(time.Minute)
			continue
		}
		if chunk != nil {
			err = w.pack(ctx, *chunk)
			if err != nil {
				w.logger.Errorw("failed to pack", "error", err)
			}
		}
		time.Sleep(time.Minute)
	}
}

func (w *WorkerThread) updateInactive() error {
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

func (w *WorkerThread) findPackWork(ctx context.Context) (*model.Chunk, error) {
	chunk := model.Chunk{}
	err := w.db.Model(&model.Chunk{}).
		Where("packing_state = ?", model.Ready).
		Order("id asc").
		Limit(1).
		Updates(map[string]interface{}{
			"packing_state":      model.Processing,
			"scanning_worker_id": w.id,
			"error_message":      "",
		}).First(&chunk).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if err != nil {
		w.logger.With("chunk", chunk).Info("found chunk to pack")
		return &chunk, nil
	}

	return nil, nil
}

func (w *WorkerThread) findScanWork(ctx context.Context) (*model.Source, error) {
	source := model.Source{}
	// Find all ready source (never scanned before)
	err := w.db.Model(&model.Source{}).
		Where("scanning_state = ?", model.Ready).
		Order("id asc").
		Limit(1).
		Updates(map[string]interface{}{
			"scanning_state":     model.Processing,
			"scanning_worker_id": w.id,
			"error_message":      "",
		}).First(&source).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if err != nil {
		w.logger.With("source", source).Info("found source to scan")
		return &source, nil
	}

	// Find all source that is complete but needs rescanning
	err = w.db.Model(&model.Source{}).
		Where("scanning_state = ? AND scan_interval != 0 AND last_scanned + scan_interval < ?",
			model.Complete, time.Now().UTC()).
		Order("id asc").
		Limit(1).
		Updates(map[string]interface{}{
			"scanning_state":     model.Processing,
			"scanning_worker_id": w.id,
			"error_message":      "",
		}).First(&source).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if err != nil {
		w.logger.With("source", source).Info("found source to scan")
		return &source, nil
	}

	return nil, nil
}

func (w *WorkerThread) EnsureParentDirectories(item model.Item) (model.Directory, error) {
	streamer := datasource.Streamers[item.Type]
	dirPath := streamer.GetDir(item.Path)
	last := item.Chunk.Source.RootDirectory
	segments := strings.Split(dirPath, "/")
	for i, segment := range segments {
		if segment == "" {
			continue
		}
		p := strings.Join(segments[:i+1], "/")
		if dir, ok := w.directoryCache[p]; ok {
			last = dir
			continue
		}
		newDir := model.Directory{
			Name:     segment,
			ParentID: &last.ID,
		}
		err := w.db.Where("parent_id = ? AND name = ?", last.ID, segment).
			FirstOrCreate(&newDir).Error
		if err != nil {
			return last, err
		}
		w.directoryCache[p] = newDir
		last = newDir
	}

	return last, nil
}

func (w *WorkerThread) Chunk(
	ctx context.Context,
	source model.Source,
	remainingItems []model.Item,
	size uint64) ([]model.Item, uint64, error) {
	if size <= source.Dataset.MaxSize {
		err := w.db.Transaction(func(tx *gorm.DB) error {
			chunk := model.Chunk{
				SourceID:     source.ID,
				PackingState: model.Ready,
			}
			err := tx.Create(&chunk).Error
			if err != nil {
				return err
			}
			for _, item := range remainingItems {
				last, err := w.EnsureParentDirectories(item)
				if err != nil {
					return errors.Wrap(err, "ensure parent directories")
				}

				err = tx.Model(&item).Updates(map[string]interface{}{
					"chunk_id":     chunk.ID,
					"directory_id": last.ID,
				}).Error
				if err != nil {
					return err
				}
			}
			return nil
		})
		if err != nil {
			return nil, 0, err
		}
		return nil, 0, nil
	}
	// size > maxSize, first, find the first item that makes it larger than maxSize
	s := uint64(0)
	si := 0
	for i, item := range remainingItems {
		s += item.Length
		if s >= source.Dataset.MaxSize {
			si = i
			s -= item.Length
			break
		}
	}

	// then, we need to split the next item
	remainingSize := (source.Dataset.MinSize+source.Dataset.MaxSize)/2 - s
	newRemainingItems := make([]model.Item, 0)
	newSize := uint64(0)
	err := w.db.Transaction(func(tx *gorm.DB) error {
		chunk := model.Chunk{
			SourceID:     source.ID,
			PackingState: model.Ready,
		}
		err := tx.Create(&chunk).Error
		if err != nil {
			return err
		}
		for _, item := range remainingItems[:si] {
			err = tx.Model(&item).Update("chunk_id", chunk.ID).Error
			if err != nil {
				return err
			}
		}
		bigItem := remainingItems[si]
		newItem := model.Item{
			ScannedAt:    bigItem.ScannedAt,
			ChunkID:      chunk.ID,
			Type:         bigItem.Type,
			Path:         bigItem.Path,
			Size:         bigItem.Size,
			Offset:       bigItem.Offset,
			Length:       remainingSize,
			LastModified: bigItem.LastModified,
			Version:      bigItem.Version,
		}
		err = tx.Create(&newItem).Error
		if err != nil {
			return err
		}
		err = tx.Model(&bigItem).Updates(map[string]interface{}{
			"offset": bigItem.Offset + remainingSize,
			"length": bigItem.Length - remainingSize,
		}).Error
		if err != nil {
			return err
		}
		newRemainingItems = append(newRemainingItems, bigItem)
		newSize += bigItem.Length
		for _, item := range remainingItems[si+1:] {
			newRemainingItems = append(newRemainingItems, item)
			newSize += item.Length
		}
		return nil
	})
	if err != nil {
		return nil, 0, err
	}
	return newRemainingItems, newSize, nil
}

func (w *WorkerThread) pack(ctx context.Context, chunk model.Chunk) error {
	items := chunk.Items
	var outDir string
	ourDirs := chunk.Source.Dataset.OutputDirs
	if len(ourDirs) > 0 {
		var err error
		outDir, err = pack.GetPathWithMostSpace(ourDirs)
		if err != nil {
			w.logger.Warnw("failed to get path with most space. using the first one", "error", err)
			outDir = ourDirs[0]
		}
	}
	result, err := pack.PackItems(ctx, items, outDir, chunk.Source.Dataset.PieceSize)
	if err != nil {
		return errors.Wrap(err, "failed to pack items")
	}
	err = w.db.Transaction(func(tx *gorm.DB) error {
		car := model.Car{
			CreatedAt: time.Now().UTC(),
			PieceCID:  result.PieceCID.String(),
			PieceSize: result.PieceSize,
			RootCID:   result.RootCID.String(),
			FileSize:  result.CarFileSize,
			FilePath:  result.CarFilePath,
			ChunkID:   chunk.ID,
			Header:    result.Header,
		}
		err := tx.Create(&car).Error
		if err != nil {
			return errors.Wrap(err, "failed to create car")
		}
		err = tx.Create(&result.RawBlocks).Error
		if err != nil {
			return errors.Wrap(err, "failed to create raw blocks")
		}
		err = tx.Create(&result.ItemBlocks).Error
		if err != nil {
			return errors.Wrap(err, "failed to create item blocks")
		}
		for _, item := range result.CarBlocks {
			item.CarID = car.ID
		}
		err = tx.Create(&result.CarBlocks).Error
		if err != nil {
			return errors.Wrap(err, "failed to create car blocks")
		}
		return nil
	})
	if err != nil {
		return errors.Wrap(err, "failed to save car")
	}
	return nil
}

func (w *WorkerThread) scan(ctx context.Context, source model.Source) error {
	var remainingItems = make([]model.Item, 0)
	result := w.db.Where("source_id = ? AND chunk_id is null", source.ID).Order("scanned_at desc").Find(&remainingItems)
	if result.Error != nil {
		return result.Error
	}
	w.logger.With("remaining", len(remainingItems)).Info("remaining items")

	lastPath := ""
	currentSize := uint64(0)
	for _, item := range remainingItems {
		currentSize += item.Length
		lastPath = item.Path
	}

	sourceHandler := datasource.Scanners[source.Type]
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
		existing := int64(0)
		err = w.db.Where("source_id = ? AND path = ? AND size = ? AND last_modified = ?",
			source.ID, entry.Path, entry.Size, entry.LastModified).Count(&existing).Error
		if err != nil {
			return err
		}
		if existing > 0 {
			continue
		}
		item := model.Item{
			ScannedAt:    entry.ScannedAt,
			Type:         entry.Type,
			Path:         entry.Path,
			Size:         entry.Size,
			Offset:       0,
			Length:       entry.Size,
			LastModified: entry.LastModified,
			Version:      0,
		}
		w.logger.Debugw("found item", "item", item)
		err = w.db.Create(&item).Error
		if err != nil {
			return errors.Wrap(err, "failed to create item")
		}

		remainingItems = append(remainingItems, item)
		currentSize += item.Length
		for currentSize >= source.Dataset.MinSize {
			remainingItems, currentSize, err = w.Chunk(ctx, source, remainingItems, currentSize)
			if err != nil {
				return errors.Wrap(err, "failed to save chunking")
			}
		}
	}

	for currentSize > 0 {
		remainingItems, currentSize, err = w.Chunk(ctx, source, remainingItems, currentSize)
		if err != nil {
			return errors.Wrap(err, "failed to save chunking")
		}
	}

	return nil
}
