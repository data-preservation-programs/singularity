package store

import (
	"context"
	"io"

	"github.com/data-preservation-programs/singularity/datasource"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/pack"
	"github.com/pkg/errors"

	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"

	format "github.com/ipfs/go-ipld-format"
	"gorm.io/gorm"
)

var ErrNotImplemented = errors.New("not implemented")

// ItemReferenceBlockStore is a struct that represents a block store backed by item references.
// It uses a GORM database for storage and a HandlerResolver to resolve data source handlers.
//
// Fields:
// DB: The GORM database used for storage. This should be initialized and connected to a database before use.
// HandlerResolver: The HandlerResolver used to resolve data source handlers. This should be initialized with the appropriate handlers before use.
type ItemReferenceBlockStore struct {
	DB              *gorm.DB
	HandlerResolver datasource.HandlerResolver
}

// Has is a method on the ItemReferenceBlockStore struct that checks if a block with the specified CID exists in the store.
// It uses the context for the database operation and returns an error if the operation fails.
//
// Parameters:
// ctx: The context for the database operation. This can be used to cancel the operation or set a deadline.
// cid: The CID of the block to check for.
//
// Returns:
// A boolean indicating whether the block exists in the store, and an error if the operation failed.
func (i *ItemReferenceBlockStore) Has(ctx context.Context, cid cid.Cid) (bool, error) {
	var count int64
	err := i.DB.WithContext(ctx).Model(&model.CarBlock{}).Select("cid").Where("cid = ?", model.CID(cid)).Count(&count).Error
	return count > 0, err
}

func (i *ItemReferenceBlockStore) Get(ctx context.Context, cid cid.Cid) (blocks.Block, error) {
	var carBlock model.CarBlock
	err := i.DB.WithContext(ctx).Joins("Item.Source").Where("car_blocks.cid = ?", model.CID(cid)).First(&carBlock).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, format.ErrNotFound{Cid: cid}
	}
	if err != nil {
		return nil, err
	}
	if carBlock.RawBlock != nil {
		return blocks.NewBlockWithCid(carBlock.RawBlock, cid)
	}

	handler, err := i.HandlerResolver.Resolve(ctx, *carBlock.Item.Source)
	if err != nil {
		return nil, err
	}
	reader, obj, err := handler.Read(
		ctx,
		carBlock.Item.Path,
		carBlock.ItemOffset,
		int64(carBlock.BlockLength()))
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	same, explanation := pack.IsSameEntry(ctx, *carBlock.Item, obj)
	if !same {
		return nil, &FileHasChangedError{Message: "file has changed: " + explanation}
	}
	readBytes, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return blocks.NewBlockWithCid(readBytes, cid)
}

// GetSize is a method on the ItemReferenceBlockStore struct that retrieves the size of a block with the specified CID from the store.
// It uses the context for the database operation and returns an error if the operation fails.
//
// Parameters:
// ctx: The context for the database operation. This can be used to cancel the operation or set a deadline.
// c: The CID of the block whose size is to be retrieved.
//
// Returns:
// The size of the block in bytes, and an error if the operation failed. If the block does not exist in the store, it returns a
func (i *ItemReferenceBlockStore) GetSize(ctx context.Context, c cid.Cid) (int, error) {
	var carBlock model.CarBlock
	err := i.DB.WithContext(ctx).Where("cid = ?", model.CID(c)).First(&carBlock).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, format.ErrNotFound{Cid: c}
		}
		return 0, err
	}
	return int(carBlock.BlockLength()), nil
}

func (i *ItemReferenceBlockStore) Put(ctx context.Context, block blocks.Block) error {
	return ErrNotImplemented
}

func (i *ItemReferenceBlockStore) PutMany(ctx context.Context, i2 []blocks.Block) error {
	return ErrNotImplemented
}

func (i *ItemReferenceBlockStore) AllKeysChan(ctx context.Context) (<-chan cid.Cid, error) {
	return nil, ErrNotImplemented
}

func (i *ItemReferenceBlockStore) HashOnRead(enabled bool) {
}

func (i *ItemReferenceBlockStore) DeleteBlock(ctx context.Context, cid cid.Cid) error {
	return ErrNotImplemented
}
