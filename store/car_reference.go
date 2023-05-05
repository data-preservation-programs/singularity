package store

import (
	"context"
	"github.com/data-preservation-programs/go-singularity/model"
	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"
	format "github.com/ipfs/go-ipld-format"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"io"
	"os"
)

type CarReferenceBlockStore struct {
	DB *gorm.DB
}

func (c CarReferenceBlockStore) Has(ctx context.Context, cid cid.Cid) (bool, error) {
	var carBlock model.CarBlock
	err := c.DB.WithContext(ctx).Select("cid").Where("cid = ?", cid.String()).First(&carBlock).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func (c CarReferenceBlockStore) Get(ctx context.Context, cid cid.Cid) (blocks.Block, error) {
	var carBlocks []model.CarBlock
	err := c.DB.WithContext(ctx).Preload("Car").Where("cid = ?", cid.String()).Limit(10).Find(&carBlocks).Error
	if err != nil {
		return nil, err
	}
	if len(carBlocks) == 0 {
		return nil, format.ErrNotFound{Cid: cid}
	}
	errors := make([]error, 0)
	for _, carBlock := range carBlocks {
		path := carBlock.Car.FilePath
		reader, err := os.OpenFile(path, os.O_RDONLY, 0)
		if err != nil {
			errors = append(errors, err)
			continue
		}
		_, err = reader.Seek(int64(carBlock.Offset), io.SeekStart)
		if err != nil {
			errors = append(errors, err)
			reader.Close()
			continue
		}
		readBytes := make([]byte, carBlock.Length)
		_, err = io.ReadFull(reader, readBytes)
		if err != nil {
			errors = append(errors, err)
			reader.Close()
			continue
		}
		reader.Close()
		return blocks.NewBlockWithCid(readBytes, cid)
	}
	return nil, AggregateError{Errors: errors}
}

func (c CarReferenceBlockStore) GetSize(ctx context.Context, cid cid.Cid) (int, error) {
	var carBlock model.CarBlock
	err := c.DB.WithContext(ctx).Where("cid = ?", cid.String()).First(&carBlock).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, format.ErrNotFound{Cid: cid}
		}
		return 0, err
	}
	return int(carBlock.Length), nil
}

func (c CarReferenceBlockStore) Put(ctx context.Context, block blocks.Block) error {
	panic("implement me")
}

func (c CarReferenceBlockStore) PutMany(ctx context.Context, i []blocks.Block) error {
	panic("implement me")
}

func (c CarReferenceBlockStore) AllKeysChan(ctx context.Context) (<-chan cid.Cid, error) {
	panic("implement me")
}

func (c CarReferenceBlockStore) HashOnRead(enabled bool) {
	panic("implement me")
}


func (c CarReferenceBlockStore) DeleteBlock(ctx context.Context, cid cid.Cid) error {
	panic("implement me")
}

