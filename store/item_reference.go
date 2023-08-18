package store

import (
	"context"
	"io"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/storagesystem"
	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"

	format "github.com/ipfs/go-ipld-format"
	"gorm.io/gorm"
)

var ErrNotImplemented = errors.New("not implemented")

// FileReferenceBlockStore is a struct that represents a block store backed by file references.
// It uses a GORM database.
//
// Fields:
// DBNoContext: The GORM database used for storage. This should be initialized and connected to a database before use.
type FileReferenceBlockStore struct {
	DBNoContext *gorm.DB
}

// Has is a method on the FileReferenceBlockStore struct that checks if a block with the specified CID exists in the store.
// It uses the context for the database operation and returns an error if the operation fails.
//
// Parameters:
// ctx: The context for the database operation. This can be used to cancel the operation or set a deadline.
// cid: The CID of the block to check for.
//
// Returns:
// A boolean indicating whether the block exists in the store, and an error if the operation failed.
func (i *FileReferenceBlockStore) Has(ctx context.Context, cid cid.Cid) (bool, error) {
	var count int64
	err := i.DBNoContext.WithContext(ctx).Model(&model.CarBlock{}).Select("cid").Where("cid = ?", model.CID(cid)).Count(&count).Error
	return count > 0, errors.WithStack(err)
}

func (i *FileReferenceBlockStore) Get(ctx context.Context, cid cid.Cid) (blocks.Block, error) {
	var carBlock model.CarBlock
	err := i.DBNoContext.WithContext(ctx).Joins("File.Attachment.Storage").Where("car_blocks.cid = ?", model.CID(cid)).First(&carBlock).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, format.ErrNotFound{Cid: cid}
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if carBlock.RawBlock != nil {
		return blocks.NewBlockWithCid(carBlock.RawBlock, cid)
	}

	// TODO: Performance can be improved by caching the handler
	handler, err := storagesystem.NewRCloneHandler(ctx, *carBlock.File.Attachment.Storage)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	reader, obj, err := handler.Read(
		ctx,
		carBlock.File.Path,
		carBlock.FileOffset,
		int64(carBlock.BlockLength()))
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer reader.Close()
	same, explanation := storagesystem.IsSameEntry(ctx, *carBlock.File, obj)
	if !same {
		return nil, errors.Wrap(ErrFileHasChanged, explanation)
	}
	readBytes, err := io.ReadAll(reader)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return blocks.NewBlockWithCid(readBytes, cid)
}

// GetSize is a method on the FileReferenceBlockStore struct that retrieves the size of a block with the specified CID from the store.
// It uses the context for the database operation and returns an error if the operation fails.
//
// Parameters:
// ctx: The context for the database operation. This can be used to cancel the operation or set a deadline.
// c: The CID of the block whose size is to be retrieved.
//
// Returns:
// The size of the block in bytes, and an error if the operation failed. If the block does not exist in the store, it returns a
func (i *FileReferenceBlockStore) GetSize(ctx context.Context, c cid.Cid) (int, error) {
	var carBlock model.CarBlock
	err := i.DBNoContext.WithContext(ctx).Where("cid = ?", model.CID(c)).First(&carBlock).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, format.ErrNotFound{Cid: c}
		}
		return 0, errors.WithStack(err)
	}
	return int(carBlock.BlockLength()), nil
}

func (i *FileReferenceBlockStore) Put(ctx context.Context, block blocks.Block) error {
	return ErrNotImplemented
}

func (i *FileReferenceBlockStore) PutMany(ctx context.Context, i2 []blocks.Block) error {
	return ErrNotImplemented
}

func (i *FileReferenceBlockStore) AllKeysChan(ctx context.Context) (<-chan cid.Cid, error) {
	return nil, ErrNotImplemented
}

func (i *FileReferenceBlockStore) HashOnRead(enabled bool) {
}

func (i *FileReferenceBlockStore) DeleteBlock(ctx context.Context, cid cid.Cid) error {
	return ErrNotImplemented
}
