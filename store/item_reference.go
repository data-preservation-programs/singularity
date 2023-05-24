package store

import (
	"context"
	"github.com/data-preservation-programs/singularity/datasource"
	"github.com/data-preservation-programs/singularity/model"
	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"
	format "github.com/ipfs/go-ipld-format"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"io"
)

type ItemReferenceBlockStore struct {
	DB              *gorm.DB
	HandlerResolver datasource.HandlerResolver
}

func (i ItemReferenceBlockStore) Has(ctx context.Context, cid cid.Cid) (bool, error) {
	var count int64
	err := i.DB.WithContext(ctx).Model(&model.CarBlock{}).Select("Cid").Where("Cid = ?", cid.String()).Count(&count).Error
	return count > 0, err
}

func (i ItemReferenceBlockStore) Get(ctx context.Context, cid cid.Cid) (blocks.Block, error) {
	var carBlock model.CarBlock
	err := i.DB.WithContext(ctx).Preload("Item.Source").Where("Cid = ?", cid.String()).First(&carBlock).Error
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
	reader, _, err := handler.Read(
		ctx,
		carBlock.Item.Path,
		carBlock.CarOffset+carBlock.Item.Offset,
		int64(carBlock.CarBlockLength))
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	readBytes, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return blocks.NewBlockWithCid(readBytes, cid)
}

func (i ItemReferenceBlockStore) GetSize(ctx context.Context, c cid.Cid) (int, error) {
	var itemBlocks model.CarBlock
	err := i.DB.WithContext(ctx).Where("Cid = ?", c.String()).First(&itemBlocks).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, format.ErrNotFound{Cid: c}
		}
		return 0, err
	}
	return int(itemBlocks.CarBlockLength), nil
}

func (i ItemReferenceBlockStore) Put(ctx context.Context, block blocks.Block) error {
	panic("implement me")
}

func (i ItemReferenceBlockStore) PutMany(ctx context.Context, i2 []blocks.Block) error {
	panic("implement me")
}

func (i ItemReferenceBlockStore) AllKeysChan(ctx context.Context) (<-chan cid.Cid, error) {
	panic("implement me")
}

func (i ItemReferenceBlockStore) HashOnRead(enabled bool) {
	panic("implement me")
}

func (i ItemReferenceBlockStore) DeleteBlock(ctx context.Context, cid cid.Cid) error {
	panic("implement me")
}
