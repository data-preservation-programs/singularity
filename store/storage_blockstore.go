package store

import (
	"context"
	"io"
	"sync"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/storagesystem"
	"github.com/data-preservation-programs/singularity/util"
	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"
	format "github.com/ipfs/go-ipld-format"
	"gorm.io/gorm"
)

// StorageBlockStore is a blockstore backed by the singularity database and
// rclone storage backends. It pools rclone handlers per storage and holds
// open a streaming file reader to serve sequential block reads efficiently.
//
// DAG nodes (directory structure, file roots) are stored inline in the DB
// and returned without any storage I/O. File-backed leaf blocks are read
// from source files via rclone, with a single-entry reader cache that
// exploits the depth-first traversal pattern: consecutive Get() calls for
// blocks from the same file read from the same held-open stream.
type StorageBlockStore struct {
	DBNoContext *gorm.DB

	mu       sync.Mutex
	handlers map[model.StorageID]*storagesystem.RCloneHandler
	active   *fileReader
}

type fileReader struct {
	fileID model.FileID
	reader io.ReadCloser
	offset int64
	cancel context.CancelFunc
}

func (s *StorageBlockStore) Has(ctx context.Context, c cid.Cid) (bool, error) {
	var count int64
	err := s.DBNoContext.WithContext(ctx).Model(&model.CarBlock{}).
		Select("cid").Where("cid = ?", model.CID(c)).Count(&count).Error
	return count > 0, errors.WithStack(err)
}

func (s *StorageBlockStore) Get(ctx context.Context, c cid.Cid) (blocks.Block, error) {
	var carBlock model.CarBlock
	err := s.DBNoContext.WithContext(ctx).
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
	s.mu.Lock()
	defer s.mu.Unlock()

	if carBlock.File == nil || carBlock.File.Attachment == nil || carBlock.File.Attachment.Storage == nil {
		return nil, errors.Errorf("block %s has no associated storage (orphaned or deleted source file)", c)
	}
	storage := *carBlock.File.Attachment.Storage
	file := *carBlock.File
	blockLen := int64(carBlock.BlockLength())

	handler, err := s.getHandler(ctx, storage)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// try to read from held-open stream
	if s.active != nil && s.active.fileID == file.ID && s.active.offset == carBlock.FileOffset {
		data, err := s.readFromActive(blockLen)
		if err != nil {
			s.closeActive()
			return nil, err
		}
		return blocks.NewBlockWithCid(data, c)
	}

	// different file or non-sequential offset -- close old, open new
	s.closeActive()

	// use a detached context for the cached reader so it outlives the
	// request that created it. closeActive() cancels it on cleanup.
	readerCtx, readerCancel := context.WithCancel(context.Background())
	reader, obj, err := handler.Read(readerCtx, file.Path, carBlock.FileOffset, file.Size-carBlock.FileOffset)
	if err != nil {
		readerCancel()
		return nil, errors.WithStack(err)
	}

	same, explanation := storagesystem.IsSameEntry(ctx, file, obj)
	if !same {
		reader.Close()
		readerCancel()
		return nil, errors.Wrap(ErrFileHasChanged, explanation)
	}

	s.active = &fileReader{
		fileID: file.ID,
		reader: reader,
		offset: carBlock.FileOffset,
		cancel: readerCancel,
	}

	data, err := s.readFromActive(blockLen)
	if err != nil {
		s.closeActive()
		return nil, err
	}
	return blocks.NewBlockWithCid(data, c)
}

func (s *StorageBlockStore) readFromActive(length int64) ([]byte, error) {
	buf := make([]byte, length)
	_, err := io.ReadFull(s.active.reader, buf)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	s.active.offset += length
	return buf, nil
}

func (s *StorageBlockStore) closeActive() {
	if s.active != nil {
		s.active.reader.Close()
		s.active.cancel()
		s.active = nil
	}
}

func (s *StorageBlockStore) getHandler(ctx context.Context, storage model.Storage) (*storagesystem.RCloneHandler, error) {
	if s.handlers == nil {
		s.handlers = make(map[model.StorageID]*storagesystem.RCloneHandler)
	}
	if h, ok := s.handlers[storage.ID]; ok {
		return h, nil
	}
	h, err := storagesystem.NewRCloneHandler(ctx, storage)
	if err != nil {
		return nil, err
	}
	s.handlers[storage.ID] = h
	return h, nil
}

func (s *StorageBlockStore) GetSize(ctx context.Context, c cid.Cid) (int, error) {
	var carBlock model.CarBlock
	err := s.DBNoContext.WithContext(ctx).Where("cid = ?", model.CID(c)).First(&carBlock).Error
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

// Close releases all held resources.
// rclone handler cleanup is not implemented -- backends may hold
// connections (e.g. SFTP) but RCloneHandler doesn't expose Shutdown.
func (s *StorageBlockStore) Close() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.closeActive()
}
