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
	"os"
	"strings"
)

type AggregateError struct {
	Errors []error
}

func (a AggregateError) Error() string {
	errors := make([]string, len(a.Errors))
	for i, err := range a.Errors {
		errors[i] = err.Error()
	}

	return strings.Join(errors, ", ")
}

type ItemReferenceBlockstore struct {
	DB *gorm.DB
}

type CarReferenceBlockstore struct {
	DB *gorm.DB
}

func (c CarReferenceBlockstore) DeleteBlock(ctx context.Context, cid cid.Cid) error {
	return c.DB.WithContext(ctx).Delete(&model.CarBlock{}, "cid = ?", cid.String()).Error
}

func (c CarReferenceBlockstore) Has(ctx context.Context, cid cid.Cid) (bool, error) {
	var count int64
	err := c.DB.WithContext(ctx).Model(&model.CarBlock{}).Where("cid = ?", cid.String()).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (c CarReferenceBlockstore) Get(ctx context.Context, cid cid.Cid) (blocks.Block, error) {
	var carBlocks []model.CarBlock
	err := c.DB.WithContext(ctx).Where("cid = ?", cid.String()).Find(&carBlocks).Error
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
		_, err = reader.Read(readBytes)
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

func (c CarReferenceBlockstore) GetSize(ctx context.Context, cid cid.Cid) (int, error) {
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

func (c CarReferenceBlockstore) Put(ctx context.Context, block blocks.Block) error {
	panic("implement me")
}

func (c CarReferenceBlockstore) PutMany(ctx context.Context, i []blocks.Block) error {
	panic("implement me")
}

func (c CarReferenceBlockstore) AllKeysChan(ctx context.Context) (<-chan cid.Cid, error) {
	panic("implement me")
}

func (c CarReferenceBlockstore) HashOnRead(enabled bool) {
	panic("implement me")
}

func (i ItemReferenceBlockstore) DeleteBlock(ctx context.Context, cid cid.Cid) error {
	return i.DB.WithContext(ctx).Delete(&model.ItemBlock{}, "cid = ?", cid.String()).Error
}

func (i ItemReferenceBlockstore) Has(ctx context.Context, cid cid.Cid) (bool, error) {
	var count int64
	err := i.DB.WithContext(ctx).Model(&model.ItemBlock{}).Where("cid = ?", cid.String()).Count(&count).Error
	return count > 0, err
}

func (i ItemReferenceBlockstore) Get(ctx context.Context, cid cid.Cid) (blocks.Block, error) {
	var itemBlocks []model.ItemBlock
	err := i.DB.WithContext(ctx).Where("cid = ?", cid.String()).Find(&itemBlocks).Error
	if err != nil {
		return nil, err
	}
	if len(itemBlocks) == 0 {
		return nil, format.ErrNotFound{Cid: cid}
	}
	errors := make([]error, 0)
	for _, itemBlock := range itemBlocks {
		reader, err := datasource.Streamers[itemBlock.Item.Type].Open(
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

func (i ItemReferenceBlockstore) GetSize(ctx context.Context, c cid.Cid) (int, error) {
	var itemBlocks model.ItemBlock
	err := i.DB.WithContext(ctx).Where("cid = ?", c.String()).First(&itemBlocks).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, format.ErrNotFound{Cid: c}
		}
		return 0, err
	}
	return int(itemBlocks.Length), nil
}

func (i ItemReferenceBlockstore) Put(ctx context.Context, block blocks.Block) error {
	panic("implement me")
}

func (i ItemReferenceBlockstore) PutMany(ctx context.Context, i2 []blocks.Block) error {
	panic("implement me")
}

func (i ItemReferenceBlockstore) AllKeysChan(ctx context.Context) (<-chan cid.Cid, error) {
	panic("implement me")
}

func (i ItemReferenceBlockstore) HashOnRead(enabled bool) {
	panic("implement me")
}
