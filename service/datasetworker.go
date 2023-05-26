package service

import (
	"context"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/pack/daggen"
	"github.com/fxamacker/cbor"
	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-blockservice"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-datastore"
	blockstore "github.com/ipfs/go-ipfs-blockstore"
	format "github.com/ipfs/go-ipld-format"
	"github.com/ipfs/go-merkledag"
	uio "github.com/ipfs/go-unixfs/io"
	"github.com/rjNemo/underscore"
	"gorm.io/gorm/clause"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/data-preservation-programs/singularity/datasource"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/pack"
	"github.com/google/uuid"
	"github.com/ipfs/go-log/v2"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type DatasetWorker struct {
	db             *gorm.DB
	concurrency    int
	threads        []DatasetWorkerThread
	exitOnComplete bool
	enableScan     bool
	enablePack     bool
	enableDag      bool
}

func NewDatasetWorker(db *gorm.DB, concurrency int, exitOnComplete bool, enableScan bool, enablePack bool, enableDag bool) *DatasetWorker {
	return &DatasetWorker{
		db:             db,
		concurrency:    concurrency,
		threads:        make([]DatasetWorkerThread, concurrency),
		exitOnComplete: exitOnComplete,
		enableScan:     enableScan,
		enablePack:     enablePack,
		enableDag:      enableDag,
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
	exitOnComplete            bool
	enableScan                bool
	enablePack                bool
	enableDag                 bool
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
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM, syscall.SIGTRAP, os.Kill, os.Interrupt)
	errChan := make(chan error)

	if w.db.Dialector.Name() == "sqlite" {
		healthCheckCleanup(w.db)
	} else {
		go StartHealthCheckCleanup(ctx, w.db)
	}

	var wg sync.WaitGroup
	for i := 0; i < w.concurrency; i++ {
		wg.Add(1)
		id := uuid.New()
		thread := DatasetWorkerThread{
			id:                        id,
			db:                        w.db.WithContext(ctx),
			logger:                    log.Logger("worker").With("workerID", id.String()),
			directoryCache:            map[string]model.Directory{},
			datasourceHandlerResolver: datasource.DefaultHandlerResolver{},
			exitOnComplete:            w.exitOnComplete,
			enableScan:                w.enableScan,
			enablePack:                w.enablePack,
			enableDag:                 w.enableDag,
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

func (w *DatasetWorkerThread) getState() State {
	return State{
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
	HealthCheck(w.db, w.id, w.getState)
	if w.db.Dialector.Name() != "sqlite" {
		go StartHealthCheck(ctx, w.db, w.id, w.getState)
	}
	for {
		w.directoryCache = map[string]model.Directory{}
		// 1st, find scanning work
		var source *model.Source
		err := database.DoRetry(func() error {
			var err error
			source, err = w.findScanWork()
			return err
		})
		if err != nil {
			w.logger.Errorw("failed to scan", "error", err)
			goto nextLoop
		}
		if source != nil {
			err = w.scan(ctx, *source)
			if err != nil {
				w.logger.Errorw("failed to scan", "error", err)
				newState := model.Error
				newErrorMessage := err.Error()
				if errors.Is(err, context.Canceled) {
					newState = model.Ready
					newErrorMessage = ""
					ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
					defer cancel()
					w.db = w.db.WithContext(ctx)
				}
				err = w.db.Model(source).Updates(
					map[string]interface{}{
						"scanning_state":         newState,
						"scanning_worker_id":     nil,
						"error_message":          newErrorMessage,
						"last_scanned_timestamp": time.Now().UTC().Unix(),
					},
				).Error
				if err != nil {
					w.logger.Errorw("failed to update source", "error", err)
				}
				goto nextLoop
			}

			err = w.db.Model(source).Updates(
				map[string]interface{}{
					"scanning_state":         model.Complete,
					"scanning_worker_id":     nil,
					"last_scanned_timestamp": time.Now().UTC().Unix(),
				},
			).Error
			if err != nil {
				w.logger.Errorw("failed to update source", "error", err)
			}
			continue
		}

		// 2nd, find ipld work
		// 3rd, find packing work
		{
			var chunk *model.Chunk
			err = database.DoRetry(func() error {
				var err error
				chunk, err = w.findPackWork()
				return err
			})
			if err != nil {
				w.logger.Errorw("failed to find pack work", "error", err)
				goto nextLoop
			}
			if chunk != nil {
				err = w.pack(
					ctx,
					chunk.ID,
					*chunk.Source,
					chunk.Items,
					chunk.Source.Dataset.OutputDirs,
					chunk.Source.Dataset.PieceSize,
				)
				if err != nil {
					w.logger.Errorw("failed to pack", "error", err)
					newState := model.Error
					newErrorMessage := err.Error()
					if errors.Is(err, context.Canceled) {
						newState = model.Ready
						newErrorMessage = ""
						ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
						defer cancel()
						w.db = w.db.WithContext(ctx)
					}
					err = w.db.Model(&model.Chunk{}).Where("id = ?", chunk.ID).Updates(
						map[string]interface{}{
							"packing_state":     newState,
							"packing_worker_id": nil,
							"error_message":     newErrorMessage,
						},
					).Error
					if err != nil {
						w.logger.Errorw("failed to update chunk", "error", err)
					}
					goto nextLoop
				} else {
					err = w.db.Model(&model.Chunk{}).Where("id = ?", chunk.ID).Updates(
						map[string]interface{}{
							"packing_state":     model.Complete,
							"packing_worker_id": nil,
						},
					).Error
					if err != nil {
						w.logger.Errorw("failed to update chunk", "error", err)
						goto nextLoop
					}
					continue
				}
			}
		}
	nextLoop:
		if w.exitOnComplete {
			w.logger.Debug("exiting on complete")
			return
		}
		w.logger.Debug("sleeping for a minute")
		select {
		case <-ctx.Done():
			return
		case <-time.After(time.Minute):
		}
	}
}

func (w *DatasetWorkerThread) updateInactive() error {
	result := w.db.Model(&model.Source{}).
		Where("scanning_state = ?", model.Processing).
		Where("scanning_worker_id = ?", w.id).
		Updates(
			map[string]interface{}{
				"scanning_state": model.Ready,
			},
		)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (w *DatasetWorkerThread) findPackWork() (*model.Chunk, error) {
	if !w.enablePack {
		return nil, nil
	}
	chunk := model.Chunk{}
	err := w.db.Transaction(
		func(db *gorm.DB) error {
			err := db.Preload("Source.Dataset").Preload("Items").
				Set("gorm:query_option", "FOR UPDATE").
				Where(
					"packing_state = ? OR (packing_state = ? AND packing_worker_id is null)",
					model.Ready,
					model.Processing,
				).
				Order("id asc").
				First(&chunk).Error
			if err != nil {
				return err
			}
			err = db.Model(&model.Chunk{}).Where("id = ?", chunk.ID).Updates(
				map[string]interface{}{
					"packing_state":     model.Processing,
					"packing_worker_id": w.id,
					"error_message":     "",
				},
			).Error
			if err != nil {
				return err
			}
			return nil
		},
	)
	if err == nil {
		w.logger.With("chunk_id", chunk.ID).Info("found chunk to pack")
		return &chunk, nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		//nolint: nilnil
		return nil, nil
	}

	return nil, err
}

func (w *DatasetWorkerThread) findScanWork() (*model.Source, error) {
	if !w.enableScan {
		return nil, nil
	}
	var source model.Source
	// Find all ready sources or sources that is being processed but does not have a worker id
	err := w.db.Transaction(
		func(db *gorm.DB) error {
			err := db.Where(
				"scanning_state = ? OR (scanning_state = ? AND scanning_worker_id is null)",
				model.Ready,
				model.Processing,
			).Clauses(clause.Locking{Strength: "UPDATE"}).
				Order("id asc").
				First(&source).Error
			if err != nil {
				return err
			}
			err = db.Model(&source).Updates(
				map[string]interface{}{
					"scanning_state":     model.Processing,
					"scanning_worker_id": w.id,
					"error_message":      "",
				},
			).Error
			if err != nil {
				return err
			}
			return nil
		},
	)
	if err == nil {
		w.logger.With("source", source).Info("found source to scan")
		return &source, nil
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// Find all source that is complete but needs rescanning
	err = w.db.Transaction(
		func(db *gorm.DB) error {
			err := db.Where(
				"scanning_state = ? AND scan_interval_seconds > 0 AND last_scanned_timestamp + scan_interval_seconds < ?",
				model.Complete, time.Now().UTC().Unix(),
			).
				Clauses(clause.Locking{Strength: "UPDATE"}).
				Order("id asc").
				First(&source).Error
			if err != nil {
				return err
			}
			err = db.Model(&source).Updates(
				map[string]interface{}{
					"scanning_state":     model.Processing,
					"scanning_worker_id": w.id,
					"error_message":      "",
				},
			).Error
			if err != nil {
				return err
			}
			return nil
		},
	)
	if err == nil {
		w.logger.With("source", source).Info("found source to rescan")
		return &source, nil
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	//nolint: nilnil
	return nil, nil
}

func (w *DatasetWorkerThread) ensureParentDirectories(item *model.Item, root model.Directory) error {
	if item.DirectoryID != nil {
		return nil
	}
	last := root
	relativePath := strings.TrimPrefix(item.Path, root.Name+"/")
	segments := strings.Split(relativePath, "/")
	for i, segment := range segments[:len(segments)-1] {
		p := strings.Join(segments[:i+1], "/")
		if dir, ok := w.directoryCache[p]; ok {
			last = dir
			continue
		}
		newDir := model.Directory{
			Name:     segment,
			ParentID: &last.ID,
		}
		err := database.DoRetry(func() error {
			return w.db.Select("id", "cid", "name", "parent_id").
				Where("parent_id = ? AND name = ?", last.ID, segment).
				FirstOrCreate(&newDir).Error
		})
		if err != nil {
			return errors.Wrap(err, "failed to create directory")
		}
		w.directoryCache[p] = newDir
		last = newDir
	}

	item.DirectoryID = &last.ID
	return nil
}

type remain struct {
	items   []model.Item
	carSize int64
}

const carHeaderSize = 200

func newRemain() *remain {
	return &remain{
		items: make([]model.Item, 0),
		// Some buffer for header
		carSize: carHeaderSize,
	}
}

func (r *remain) add(item model.Item) {
	r.items = append(r.items, item)
	r.carSize += toCarSize(item.Length)
}

func (r *remain) reset() {
	r.items = make([]model.Item, 0)
	r.carSize = carHeaderSize
}

func (r *remain) itemIDs() []uint64 {
	out := make([]uint64, len(r.items))
	for i, item := range r.items {
		out[i] = item.ID
	}
	return out
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

func fromCarSize(size int64) int64 {
	// Rough estimate for how much can we store for remaining space of car file
	return size - (size/1024/1024+1)*90
}

func (w *DatasetWorkerThread) chunkOnce(
	source model.Source,
	dataset model.Dataset,
	remaining *remain,
) error {
	// If everything fit, create a chunk. Usually this is the case for the last chunk
	if remaining.carSize <= dataset.MaxSize {
		err := database.DoRetry(func() error {
			return w.db.Transaction(
				func(db *gorm.DB) error {
					chunk := model.Chunk{
						SourceID:     source.ID,
						PackingState: model.Ready,
					}
					err := w.db.Create(&chunk).Error
					if err != nil {
						return errors.Wrap(err, "failed to create chunk")
					}
					err = w.db.Model(&model.Item{}).
						Where("id IN (?)", remaining.itemIDs()).Update("chunk_id", chunk.ID).Error
					if err != nil {
						return errors.Wrap(err, "failed to update items")
					}
					return nil
				},
			)
		})
		if err != nil {
			return errors.Wrap(err, "failed to create chunk")
		}
		remaining.items = remaining.items[:0]
		remaining.carSize = carHeaderSize
		return nil
	}
	// size > maxSize, first, find the first item that makes it larger than maxSize
	s := remaining.carSize
	si := len(remaining.items) - 1
	for si >= 0 {
		s -= toCarSize(remaining.items[si].Length)
		if s <= dataset.MaxSize {
			break
		}
		si--
	}

	remainingSize := fromCarSize(dataset.MaxSize - s)
	// s is the size of the chunk for [0:si)
	if s >= dataset.MinSize || remainingSize <= 0 {
		// we found a chunk that is between minSize and maxSize
		err := database.DoRetry(func() error {
			return w.db.Transaction(
				func(db *gorm.DB) error {
					chunk := model.Chunk{
						SourceID:     source.ID,
						PackingState: model.Ready,
					}
					err := w.db.Create(&chunk).Error
					if err != nil {
						return errors.Wrap(err, "failed to create chunk")
					}
					itemIDs := make([]uint64, len(remaining.items[:si]))
					for i, item := range remaining.items[:si] {
						itemIDs[i] = item.ID
					}
					err = w.db.Model(&model.Item{}).Where("id IN (?)", itemIDs).Update("chunk_id", chunk.ID).Error
					if err != nil {
						return errors.Wrap(err, "failed to update items")
					}
					return nil
				},
			)
		})
		if err != nil {
			return errors.Wrap(err, "failed to create chunk")
		}
		remaining.items = remaining.items[si:]
		remaining.carSize = remaining.carSize - s + carHeaderSize
		return nil
	}

	// then, we need to split the next item
	newRemaining := newRemain()
	err := database.DoRetry(func() error {
		return w.db.Transaction(
			func(db *gorm.DB) error {
				chunk := model.Chunk{
					SourceID:     source.ID,
					PackingState: model.Ready,
				}
				err := db.Create(&chunk).Error
				if err != nil {
					return errors.Wrap(err, "failed to create chunk")
				}
				itemIDs := make([]uint64, len(remaining.items[:si]))
				for i, item := range remaining.items[:si] {
					itemIDs[i] = item.ID
				}
				err = db.Model(&model.Item{}).Where("id IN (?)", itemIDs).Update("chunk_id", chunk.ID).Error
				if err != nil {
					return errors.Wrap(err, "failed to update items")
				}
				bigItem := remaining.items[si]
				newItem := model.Item{
					ScannedAt:    bigItem.ScannedAt,
					ChunkID:      &chunk.ID,
					Path:         bigItem.Path,
					Size:         bigItem.Size,
					Offset:       bigItem.Offset,
					Length:       remainingSize,
					LastModified: bigItem.LastModified,
					SourceID:     source.ID,
					DirectoryID:  bigItem.DirectoryID,
				}
				err = db.Create(&newItem).Error
				if err != nil {
					return errors.Wrap(err, "failed to create item during chunking")
				}
				err = db.Model(&bigItem).Updates(
					map[string]interface{}{
						"offset": bigItem.Offset + remainingSize,
						"length": bigItem.Length - remainingSize,
					},
				).Error
				if err != nil {
					return errors.Wrap(err, "failed to update item")
				}
				newRemaining.add(bigItem)
				for _, item := range remaining.items[si+1:] {
					newRemaining.add(item)
				}
				return nil
			},
		)
	})
	if err != nil {
		return errors.Wrap(err, "failed to create chunk")
	}
	*remaining = *newRemaining
	return nil
}

func (w *DatasetWorkerThread) pack(
	ctx context.Context, chunkID uint32,
	source model.Source, items []model.Item, ourDirs []string, pieceSize int64,
) error {
	var outDir string
	if len(ourDirs) > 0 {
		var err error
		outDir, err = pack.GetPathWithMostSpace(ourDirs)
		if err != nil {
			w.logger.Warnw("failed to get path with most space. using the first one", "error", err)
			outDir = ourDirs[0]
		}
	}
	handler, err := w.datasourceHandlerResolver.Resolve(ctx, source)
	if err != nil {
		return errors.Wrap(err, "failed to get datasource handler")
	}
	result, err := pack.ProcessItems(ctx, handler, items, outDir, pieceSize, source.Dataset.EncryptionRecipients)
	if err != nil {
		return errors.Wrap(err, "failed to pack items")
	}

	for itemID, itemCID := range result.ItemCIDs {
		err = database.DoRetry(func() error {
			return w.db.Model(&model.Item{}).Where("id = ?", itemID).Update("cid", model.CID(itemCID)).Error
		})
		if err != nil {
			return errors.Wrap(err, "failed to update cid of item")
		}
	}


	err = database.DoRetry(func() error {
		return w.db.Transaction(
			func(db *gorm.DB) error {
				car := model.Car{
					CreatedAt: time.Now().UTC(),
					PieceCID:  model.CID(result.PieceCID),
					PieceSize: result.PieceSize,
					RootCID:   model.CID(result.RootCID),
					FileSize:  result.CarFileSize,
					FilePath:  result.CarFilePath,
					ChunkID:   &chunkID,
					DatasetID: source.DatasetID,
					Header:    result.Header,
				}
				err := db.Create(&car).Error
				if err != nil {
					return errors.Wrap(err, "failed to create car")
				}
				for i, _ := range result.CarBlocks {
					result.CarBlocks[i].CarID = car.ID
				}
				err = db.CreateInBatches(&result.CarBlocks, 1000).Error
				if err != nil {
					return errors.Wrap(err, "failed to create car blocks")
				}

				type DirectoryData struct {
					Directory  uio.Directory
					Blockstore blockstore.Blockstore
				}
				dirs := map[uint64]DirectoryData{}
				for _, item := range items {
					item.CID = model.CID(result.ItemCIDs[item.ID])
					dirID := item.DirectoryID
					for dirID != nil {
						if _, ok := dirs[*dirID]; !ok {
							var dir model.Directory
							err = db.Clauses(clause.Locking{Strength: "UPDATE"}).
								Where("id = ?", *dirID).First(&dir).Error
							if err != nil {
								return errors.Wrap(err, "failed to get directory")
							}
							if len(dir.Data) == 0 {
								ds := datastore.NewMapDatastore()
								bs := blockstore.NewBlockstore(ds)
								dagServ := merkledag.NewDAGService(blockservice.New(bs, nil))
								dirs[dir.ID] = DirectoryData{
									Directory:  uio.NewDirectory(dagServ),
									Blockstore: bs,
								}
								dirs[dir.ID].Directory.SetCidBuilder(merkledag.V1CidPrefix())
							} else {
								ds := datastore.NewMapDatastore()
								bs := blockstore.NewBlockstore(ds)
								var m map[string][]byte
								err = cbor.Unmarshal(dir.Data, &m)
								if err != nil {
									return errors.Wrap(err, "failed to unmarshal directory bytes")
								}
								for k, v := range m {
									blk, _ := blocks.NewBlockWithCid(v, cid.MustParse(k))
									_ = bs.Put(ctx, blk)
								}
								dagServ := merkledag.NewDAGService(blockservice.New(bs, nil))
								dirNode, err := dagServ.Get(ctx, cid.Cid(dir.CID))
								if err != nil {
									return errors.Wrap(err, "failed to get directory node")
								}
								directory, err := uio.NewDirectoryFromNode(dagServ, dirNode)
								directory.SetCidBuilder(merkledag.V1CidPrefix())
								if err != nil {
									return errors.Wrap(err, "failed to get directory")
								}
								dirs[dir.ID] = DirectoryData{
									Directory:  directory,
									Blockstore: bs,
								}
							}
							dirID = dir.ParentID
						} else {
							break
						}
					}

					dirID = item.DirectoryID
					segments := strings.Split(item.Path, "/")
					name := segments[len(segments)-1]
					dirData := dirs[*dirID]
					if item.Length == item.Size {
						err = dirData.Directory.AddChild(ctx, name, daggen.NewDummyNode(uint64(item.Length), cid.Cid(item.CID)))
						if err != nil {
							return errors.Wrap(err, "failed to add child for full item")
						}
					} else {
						// Find all parts of this file
						var parts []model.Item
						err = db.Where("scanned_at = ? AND source_id = ? AND path = ?", item.ScannedAt, item.SourceID, item.Path).Order("offset asc").Find(&parts).Error
						if err != nil {
							return errors.Wrap(err, "failed to find parts")
						}
						if !underscore.All(parts, func(part model.Item) bool {
							return cid.Cid(part.CID) != cid.Undef
						}) {
							// This item is not ready to be added to the directory because some parts are not yet ready
							break
						}
						// Sanitize all items are consecutive
						var offset int64
						for _, part := range parts {
							if part.Offset != offset {
								return errors.New("parts are not consecutive")
							}
							offset += part.Length
						}
						if offset != item.Size {
							return errors.New("parts are not consecutive")
						}
						var links []pack.Link
						for _, part := range parts {
							links = append(links, pack.Link{
								Link: format.Link{
									Name: "",
									Size: uint64(part.Length),
									Cid:  cid.Cid(part.CID),
								},
								ChunkSize: uint64(part.Length),
							})
						}
						blks, rootLink, _ := pack.AssembleItem(links)
						err = dirData.Directory.AddChild(ctx, name, rootLink)
						if err != nil {
							return errors.Wrap(err, "failed to add child for item parts")
						}
						err = dirData.Blockstore.PutMany(ctx, blks)
						if err != nil {
							return errors.Wrap(err, "failed to put blocks")
						}
					}
				}
				for dirID, dirData := range dirs {
					rootNode, err := dirData.Directory.GetNode()
					if err != nil {
						return errors.Wrap(err, "failed to get directory node")
					}
					keyChan, err := dirData.Blockstore.AllKeysChan(ctx)
					if err != nil {
						return errors.Wrap(err, "failed to get keys")
					}
					m := map[string][]byte{}
					for key := range keyChan {
						blk, err := dirData.Blockstore.Get(ctx, key)
						if err != nil {
							return errors.Wrap(err, "failed to get block")
						}
						m[key.String()] = blk.RawData()
					}
					m[rootNode.Cid().String()] = rootNode.RawData()
					bytes, err := cbor.Marshal(&m, cbor.CanonicalEncOptions())
					if err != nil {
						return errors.Wrap(err, "failed to marshal directory")
					}
					err = db.Model(&model.Directory{}).Where("id = ?", dirID).Updates(map[string]interface{}{
						"cid":  model.CID(rootNode.Cid()),
						"data": bytes,
					}).Error
					if err != nil {
						return errors.Wrap(err, "failed to update directory")
					}
				}
				return nil
			},
		)
	})
	if err != nil {
		return errors.Wrap(err, "failed to save car")
	}

	w.logger.With("source", source).Info("finished packing")
	if source.DeleteAfterExport && result.CarFilePath != "" {
		w.logger.Info("Deleting original data source")
		for i, item := range items {
			object := result.Objects[i]
			if item.Offset == 0 && item.Length == item.Size {
				err = object.Remove(ctx)
				if err != nil {
					w.logger.Warnw("failed to remove object", "error", err)
				}
				continue
			}
			// Make sure all parts of this file has been exported before deleting
			var allItems []model.Item
			err = w.db.Model(&model.Item{}).Where("source_id = ? AND path = ? AND size = ? AND last_modified = ?",
				item.SourceID,
				item.Path,
				item.Size,
				item.LastModified).Find(&allItems).Error
			if err != nil {
				w.logger.Warnw("failed to get all items for file", "error", err)
				continue
			}
			if underscore.Any(allItems, func(i model.Item) bool {
				return cid.Cid(i.CID) == cid.Undef || i.ChunkID == nil
			}) {
				w.logger.Info("not all items have been exported yet, skipping delete")
				continue
			}
			chunkIDs := underscore.Map(allItems, func(i model.Item) uint32 { return *i.ChunkID })
			var exported int64
			w.db.Model(&model.Car{}).Where("chunk_id in (?) AND file_path != ?", chunkIDs, "").Count(&exported)
			if exported != int64(len(chunkIDs)) {
				w.logger.Infow("not all items have been exported yet, skipping delete", "exported", exported)
				continue
			}
			err = object.Remove(ctx)
			if err != nil {
				w.logger.Warnw("failed to remove object", "error", err)
			}
		}
	}
	return nil
}

func (w *DatasetWorkerThread) scan(ctx context.Context, source model.Source) error {
	var dataset model.Dataset
	err := w.db.Model(&model.Dataset{}).Where("id = ?", source.DatasetID).First(&dataset).Error
	if err != nil {
		return errors.Wrap(err, "failed to get dataset")
	}
	var root model.Directory
	err = w.db.Model(&model.Directory{}).Where("id = ?", source.RootDirectoryID).First(&root).Error
	if err != nil {
		return errors.Wrap(err, "failed to get root directory")
	}

	var remaining = newRemain()
	err = w.db.Where("source_id = ? AND chunk_id is null", source.ID).
		Order("scanned_at asc").
		Find(&remaining.items).Error
	if err != nil {
		return err
	}
	w.logger.With("remaining", len(remaining.items)).Info("remaining items")

	lastPath := ""
	for _, item := range remaining.items {
		remaining.carSize += toCarSize(item.Length)
		lastPath = item.Path
	}

	if !source.PushOnly {
		sourceScanner, err := w.datasourceHandlerResolver.Resolve(ctx, source)
		if err != nil {
			return errors.Wrap(err, "failed to get source scanner")
		}
		entryChan := sourceScanner.Scan(ctx, "", lastPath)
		for entry := range entryChan {
			if entry.Error != nil {
				w.logger.Errorw("failed to scan", "error", entry.Error)
				continue
			}
			existing := int64(0)
			err = w.db.Model(&model.Item{}).Where(
				"source_id = ? AND path = ? AND size = ? AND last_modified = ?",
				source.ID, entry.Info.Remote(), entry.Info.Size(), entry.Info.ModTime(ctx),
			).Count(&existing).Error
			if err != nil {
				return err
			}
			if existing > 0 {
				continue
			}
			item := model.Item{
				SourceID:     source.ID,
				ScannedAt:    entry.ScannedAt,
				Path:         entry.Info.Remote(),
				Size:         entry.Info.Size(),
				Offset:       0,
				Length:       entry.Info.Size(),
				LastModified: entry.Info.ModTime(ctx),
			}
			w.logger.Debugw("found item", "item", item)
			err = w.ensureParentDirectories(&item, root)
			if err != nil {
				return errors.Wrap(err, "failed to ensure parent directories")
			}
			err = database.DoRetry(func() error {
				return w.db.Create(&item).Error
			})
			if err != nil {
				return errors.Wrap(err, "failed to create item during scanning")
			}

			remaining.add(item)
			for remaining.carSize >= dataset.MinSize {
				err = w.chunkOnce(source, dataset, remaining)
				if err != nil {
					return errors.Wrap(err, "failed to save chunking")
				}
			}
		}
	}

	for len(remaining.items) > 0 {
		err = w.chunkOnce(source, dataset, remaining)
		if err != nil {
			return errors.Wrap(err, "failed to save chunking")
		}
	}

	w.logger.With("source", source).Info("finished scanning")
	return nil
}
