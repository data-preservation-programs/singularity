package store

import (
	"context"
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/storagesystem"
	"github.com/data-preservation-programs/singularity/util"
	lru "github.com/hashicorp/golang-lru/v2"
	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"
	format "github.com/ipfs/go-ipld-format"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

// SpanConfig tunes the span-read behavior of StorageBlockStore.
// Blocks are ~1MiB (pack chunk size), so block counts are roughly MiB.
type SpanConfig struct {
	// SpanBlocks is the number of consecutive blocks fetched per backend read.
	SpanBlocks int
	// PrefetchSpans is the number of spans prefetched ahead once access to a
	// file looks sequential.
	PrefetchSpans int
	// MaxBackendReads caps concurrent backend reads across all requests.
	MaxBackendReads int
	// CacheBlocks is the block cache capacity in blocks. When <= 0 it is
	// derived from the other settings.
	CacheBlocks int
	// ReadTimeout bounds a single backend span read once it holds a
	// connection slot, so a hung backend cannot pin slots indefinitely.
	ReadTimeout time.Duration
}

func (c SpanConfig) withDefaults() SpanConfig {
	if c.SpanBlocks <= 0 {
		c.SpanBlocks = 8
	}
	if c.PrefetchSpans <= 0 {
		c.PrefetchSpans = 2
	}
	if c.MaxBackendReads <= 0 {
		c.MaxBackendReads = 64
	}
	// an active sequential client keeps 1+PrefetchSpans spans warm; size the
	// cache so MaxBackendReads concurrent streams don't evict each other's
	// unconsumed spans and re-fetch them
	if c.CacheBlocks <= 0 {
		c.CacheBlocks = c.MaxBackendReads * c.SpanBlocks * (1 + c.PrefetchSpans)
	}
	if c.ReadTimeout <= 0 {
		c.ReadTimeout = 2 * time.Minute
	}
	return c
}

// StorageBlockStore is a blockstore backed by the singularity database and
// rclone storage backends.
//
// DAG nodes (directory structure, file roots) are stored inline in the DB
// and returned without any storage I/O. File-backed leaf blocks are read
// from source files via ranged reads covering a span of consecutive blocks;
// spans land in a bounded LRU cache so the sequential Gets that follow are
// served from memory. Sequential access additionally prefetches the next
// spans in parallel, which is what provides per-client throughput beyond
// the backend's per-connection rate. There are no held-open streams and no
// global lock: concurrency is bounded only by the backend-read semaphore.
type StorageBlockStore struct {
	dbNoContext *gorm.DB
	cfg         SpanConfig

	handlersMu sync.RWMutex
	handlers   map[model.StorageID]*storagesystem.RCloneHandler

	cache    *lru.Cache[string, []byte]
	inflight singleflight.Group
	sem      chan struct{}

	// seq tracks, per file, the offset window in which the next cache miss
	// counts as a sequential continuation. Entries are disposable hints.
	seqMu sync.Mutex
	seq   map[model.FileID]seqWindow

	// bg outlives individual requests: span results are shared via
	// singleflight and cached, so a fetch must not die with the request
	// that happened to trigger it. Close cancels it.
	bg     context.Context
	cancel context.CancelFunc
}

type seqWindow struct {
	lo, hi int64
}

func NewStorageBlockStore(db *gorm.DB, cfg SpanConfig) *StorageBlockStore {
	cfg = cfg.withDefaults()
	cache, err := lru.New[string, []byte](cfg.CacheBlocks)
	if err != nil {
		panic(err)
	}
	bg, cancel := context.WithCancel(context.Background())
	return &StorageBlockStore{
		dbNoContext: db,
		cfg:         cfg,
		handlers:    make(map[model.StorageID]*storagesystem.RCloneHandler),
		cache:       cache,
		sem:         make(chan struct{}, cfg.MaxBackendReads),
		seq:         make(map[model.FileID]seqWindow),
		bg:          bg,
		cancel:      cancel,
	}
}

func (s *StorageBlockStore) Has(ctx context.Context, c cid.Cid) (bool, error) {
	if s.cache.Contains(c.KeyString()) {
		return true, nil
	}
	var count int64
	err := s.dbNoContext.WithContext(ctx).Model(&model.CarBlock{}).
		Select("cid").Where("cid = ?", model.CID(c)).Count(&count).Error
	return count > 0, errors.WithStack(err)
}

func (s *StorageBlockStore) Get(ctx context.Context, c cid.Cid) (blocks.Block, error) {
	if data, ok := s.cache.Get(c.KeyString()); ok {
		return blocks.NewBlockWithCid(data, c)
	}

	var carBlock model.CarBlock
	err := s.dbNoContext.WithContext(ctx).
		Joins("File.Attachment.Storage").
		Where("car_blocks.cid = ?", model.CID(c)).
		First(&carBlock).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, format.ErrNotFound{Cid: c}
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// inline block -- DAG nodes, small files
	if carBlock.RawBlock != nil {
		return blocks.NewBlockWithCid(carBlock.RawBlock, c)
	}

	return s.readFileBlock(ctx, carBlock, c)
}

func (s *StorageBlockStore) readFileBlock(ctx context.Context, carBlock model.CarBlock, c cid.Cid) (blocks.Block, error) {
	if carBlock.File == nil || carBlock.File.Attachment == nil || carBlock.File.Attachment.Storage == nil {
		return nil, errors.Errorf("block %s has no associated storage (orphaned or deleted source file)", c)
	}
	file := *carBlock.File
	storage := *carBlock.File.Attachment.Storage

	blks, end, err := s.fetchSpan(ctx, file, storage, carBlock.FileOffset)
	if err != nil {
		return nil, err
	}

	s.maybePrefetch(file, storage, carBlock.FileOffset, end)

	for _, blk := range blks {
		if blk.Cid().Equals(c) {
			return blk, nil
		}
	}
	return nil, errors.Errorf("block %s missing from fetched span (car_blocks changed underneath)", c)
}

type spanResult struct {
	blks []blocks.Block
	end  int64
}

// fetchSpan reads one span of consecutive blocks starting at offset and
// populates the cache. Concurrent fetches of the same span are collapsed.
// The fetch itself runs on the store's background context: the result is
// shared, so it must not be poisoned by one caller's cancellation. Only the
// wait is bound to ctx -- a caller whose client disconnected detaches while
// the fetch continues and still lands in the cache.
func (s *StorageBlockStore) fetchSpan(ctx context.Context, file model.File, storage model.Storage, offset int64) ([]blocks.Block, int64, error) {
	key := fmt.Sprintf("%d:%d", file.ID, offset)
	ch := s.inflight.DoChan(key, func() (any, error) {
		rows, err := s.spanRows(file.ID, offset, s.cfg.SpanBlocks)
		if err != nil {
			return nil, err
		}
		if len(rows) == 0 {
			return &spanResult{}, nil
		}
		blks, err := s.readRows(file, storage, rows)
		if err != nil {
			return nil, err
		}
		last := rows[len(rows)-1]
		return &spanResult{blks: blks, end: last.FileOffset + int64(last.BlockLength())}, nil
	})
	select {
	case res := <-ch:
		if res.Err != nil {
			return nil, 0, res.Err
		}
		r, _ := res.Val.(*spanResult)
		return r.blks, r.end, nil
	case <-ctx.Done():
		return nil, 0, errors.WithStack(ctx.Err())
	}
}

// spanRows returns up to n car_blocks rows for a file starting at offset,
// trimmed to the contiguous run so a single ranged read covers them exactly.
func (s *StorageBlockStore) spanRows(fileID model.FileID, offset int64, n int) ([]model.CarBlock, error) {
	var rows []model.CarBlock
	err := s.dbNoContext.WithContext(s.bg).
		Where("file_id = ? AND file_offset >= ?", fileID, offset).
		Order("file_offset").Limit(n).Find(&rows).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}
	end := int64(-1)
	for i := range rows {
		if i > 0 && rows[i].FileOffset != end {
			rows = rows[:i]
			break
		}
		end = rows[i].FileOffset + int64(rows[i].BlockLength())
	}
	return rows, nil
}

// readRows performs the ranged backend read covering rows and caches the
// resulting blocks. Each block owns its own allocation so a retained block
// never pins the rest of the span.
func (s *StorageBlockStore) readRows(file model.File, storage model.Storage, rows []model.CarBlock) ([]blocks.Block, error) {
	blks := make([]blocks.Block, len(rows))

	// a prefetch may race a sync read of an overlapping span after a window
	// reset -- if everything is already cached, skip the backend read
	allCached := true
	for i := range rows {
		data, ok := s.cache.Get(cid.Cid(rows[i].CID).KeyString())
		if !ok {
			allCached = false
			break
		}
		blk, err := blocks.NewBlockWithCid(data, cid.Cid(rows[i].CID))
		if err != nil {
			return nil, errors.WithStack(err)
		}
		blks[i] = blk
	}
	if allCached {
		return blks, nil
	}

	handler, err := s.getHandler(storage)
	if err != nil {
		return nil, err
	}

	var total int64
	for i := range rows {
		total += int64(rows[i].BlockLength())
	}

	select {
	case s.sem <- struct{}{}:
	case <-s.bg.Done():
		return nil, errors.WithStack(s.bg.Err())
	}
	defer func() { <-s.sem }()

	// bound slot occupancy -- a hung backend read must not pin a
	// connection slot until shutdown
	rctx, rcancel := context.WithTimeout(s.bg, s.cfg.ReadTimeout)
	defer rcancel()

	reader, obj, err := handler.Read(rctx, file.Path, rows[0].FileOffset, total)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer reader.Close()

	same, explanation := storagesystem.IsSameEntry(rctx, file, obj)
	if !same {
		return nil, errors.Wrap(ErrFileHasChanged, explanation)
	}

	for i := range rows {
		data := make([]byte, rows[i].BlockLength())
		if _, err := io.ReadFull(reader, data); err != nil {
			return nil, errors.WithStack(err)
		}
		blk, err := blocks.NewBlockWithCid(data, cid.Cid(rows[i].CID))
		if err != nil {
			return nil, errors.WithStack(err)
		}
		blks[i] = blk
		s.cache.Add(cid.Cid(rows[i].CID).KeyString(), data)
	}
	return blks, nil
}

// maybePrefetch schedules background reads of the next spans when access
// looks sequential. The first span of a file only opens the window; the
// second consecutive one starts prefetching, so random seeks and small
// files never trigger it.
func (s *StorageBlockStore) maybePrefetch(file model.File, storage model.Storage, offset, end int64) {
	if end <= 0 {
		return
	}

	s.seqMu.Lock()
	w, ok := s.seq[file.ID]
	sequential := ok && offset > w.lo && offset <= w.hi
	// hints are disposable -- reset rather than track LRU
	if len(s.seq) > 4096 {
		s.seq = make(map[model.FileID]seqWindow)
	}
	s.seq[file.ID] = seqWindow{lo: offset, hi: end}
	s.seqMu.Unlock()

	if !sequential {
		return
	}

	go func() {
		rows, err := s.spanRows(file.ID, end, s.cfg.SpanBlocks*s.cfg.PrefetchSpans)
		if err != nil || len(rows) == 0 {
			return
		}
		// spans fetch in parallel -- this is the multipart read
		for i := 0; i < len(rows); i += s.cfg.SpanBlocks {
			go func(off int64) {
				_, _, _ = s.fetchSpan(s.bg, file, storage, off)
			}(rows[i].FileOffset)
		}
		// extend the window past the prefetched region so the miss that
		// follows its consumption still counts as sequential
		last := rows[len(rows)-1]
		hi := last.FileOffset + int64(last.BlockLength())
		s.seqMu.Lock()
		if w, ok := s.seq[file.ID]; ok && hi > w.hi {
			w.hi = hi
			s.seq[file.ID] = w
		}
		s.seqMu.Unlock()
	}()
}

func (s *StorageBlockStore) getHandler(storage model.Storage) (*storagesystem.RCloneHandler, error) {
	s.handlersMu.RLock()
	h, ok := s.handlers[storage.ID]
	s.handlersMu.RUnlock()
	if ok {
		return h, nil
	}

	s.handlersMu.Lock()
	defer s.handlersMu.Unlock()
	if h, ok := s.handlers[storage.ID]; ok {
		return h, nil
	}
	h, err := storagesystem.NewRCloneHandler(s.bg, storage)
	if err != nil {
		return nil, err
	}
	s.handlers[storage.ID] = h
	return h, nil
}

func (s *StorageBlockStore) GetSize(ctx context.Context, c cid.Cid) (int, error) {
	if data, ok := s.cache.Get(c.KeyString()); ok {
		return len(data), nil
	}
	var carBlock model.CarBlock
	err := s.dbNoContext.WithContext(ctx).Where("cid = ?", model.CID(c)).First(&carBlock).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, format.ErrNotFound{Cid: c}
		}
		return 0, errors.WithStack(err)
	}
	return int(carBlock.BlockLength()), nil
}

func (s *StorageBlockStore) Put(ctx context.Context, block blocks.Block) error {
	return util.ErrNotImplemented
}

func (s *StorageBlockStore) PutMany(ctx context.Context, blks []blocks.Block) error {
	return util.ErrNotImplemented
}

func (s *StorageBlockStore) AllKeysChan(ctx context.Context) (<-chan cid.Cid, error) {
	return nil, util.ErrNotImplemented
}

func (s *StorageBlockStore) HashOnRead(enabled bool) {}

func (s *StorageBlockStore) DeleteBlock(ctx context.Context, c cid.Cid) error {
	return util.ErrNotImplemented
}

// Close cancels in-flight background reads; they exit on context
// cancellation. rclone handler cleanup is not implemented -- backends may
// hold connections (e.g. SFTP) but RCloneHandler doesn't expose Shutdown.
func (s *StorageBlockStore) Close() {
	s.cancel()
	s.cache.Purge()
}
