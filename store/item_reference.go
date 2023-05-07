package store

import (
	"context"
	"github.com/data-preservation-programs/go-singularity/datasource"
	"github.com/data-preservation-programs/go-singularity/model"
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
	err := i.DB.WithContext(ctx).Model(&model.CarBlock{}).Select("cid").Where("cid = ?", cid.String()).Count(&count).Error
	return count > 0, err
}

func (i ItemReferenceBlockStore) Get(ctx context.Context, cid cid.Cid) (blocks.Block, error) {
	var itemBlocks []model.CarBlock
	err := i.DB.WithContext(ctx).Preload("Item.Source").Where("cid = ?", cid.String()).Limit(10).Find(&itemBlocks).Error
	if err != nil {
		return nil, err
	}
	if len(itemBlocks) == 0 {
		return nil, format.ErrNotFound{Cid: cid}
	}
	errors := make([]error, 0)
	for _, itemBlock := range itemBlocks {
		handler, err := i.HandlerResolver.GetHandler(*itemBlock.Item.Source)
		if err != nil {
			errors = append(errors, err)
			continue
		}
		reader, err := handler.Read(
			ctx,
			itemBlock.Item.Path,
			itemBlock.Offset+itemBlock.Item.Offset,
			itemBlock.Length)
		if err != nil {
			errors = append(errors, err)
			continue
		}
		readBytes, err := io.ReadAll(reader)
		if err != nil {
			errors = append(errors, err)
			reader.Close()
			continue
		}
		return blocks.NewBlockWithCid(readBytes, cid)
	}
	return nil, AggregateError{Errors: errors}
}

func (i ItemReferenceBlockStore) GetSize(ctx context.Context, c cid.Cid) (int, error) {
	var itemBlocks model.CarBlock
	err := i.DB.WithContext(ctx).Where("cid = ?", c.String()).First(&itemBlocks).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, format.ErrNotFound{Cid: c}
		}
		return 0, err
	}
	return int(itemBlocks.Length), nil
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
