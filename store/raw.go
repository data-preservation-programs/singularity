package store

import (
	"context"
	"github.com/data-preservation-programs/go-singularity/model"
	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"
	format "github.com/ipfs/go-ipld-format"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type RawBlockStore struct {
	DB *gorm.DB
}

func (r RawBlockStore) Has(ctx context.Context, cid cid.Cid) (bool, error) {
	var rawBlock model.RawBlock
	err := r.DB.WithContext(ctx).Select("cid").Where("cid = ?", cid.String()).First(&rawBlock).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r RawBlockStore) Get(ctx context.Context, cid cid.Cid) (blocks.Block, error) {
	var rawBlock model.RawBlock
	err := r.DB.WithContext(ctx).Where("cid = ?", cid.String()).First(&rawBlock).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, format.ErrNotFound{Cid: cid}
	}
	if err != nil {
		return nil, err
	}
	return blocks.NewBlockWithCid(rawBlock.Block, cid)
}

func (r RawBlockStore) GetSize(ctx context.Context, cid cid.Cid) (int, error) {
	var rawBlock model.RawBlock
	err := r.DB.WithContext(ctx).Select("length").Where("cid = ?", cid.String()).First(&rawBlock).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, format.ErrNotFound{Cid: cid}
	}
	if err != nil {
		return 0, err
	}
	return int(rawBlock.Length), nil
}

func (r RawBlockStore) Put(ctx context.Context, block blocks.Block) error {
	//TODO implement me
	panic("implement me")
}

func (r RawBlockStore) PutMany(ctx context.Context, blocks []blocks.Block) error {
	//TODO implement me
	panic("implement me")
}

func (r RawBlockStore) AllKeysChan(ctx context.Context) (<-chan cid.Cid, error) {
	//TODO implement me
	panic("implement me")
}

func (r RawBlockStore) HashOnRead(enabled bool) {
	//TODO implement me
	panic("implement me")
}

func (r RawBlockStore) DeleteBlock(ctx context.Context, cid cid.Cid) error {
	//TODO implement me
	panic("implement me")
}
