package store

import (
	"context"
	"github.com/data-preservation-programs/go-singularity/model"
	"github.com/ipfs/go-cid"
	blockstore "github.com/ipfs/go-ipfs-blockstore"
	"github.com/multiformats/go-varint"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

type PieceReader struct {
	db  *gorm.DB
	car *model.Car
	blockStore blockstore.Blockstore
}

const CidLength = 32

func NewPieceReader(db *gorm.DB, car *model.Car) PieceReader {
	rawBlockStore := RawBlockStore{DB: db}
	itemReferenceBlockStore := ItemReferenceBlockStore{DB: db}
	priorityStore := NewPriorityBlockStore(rawBlockStore, itemReferenceBlockStore)
	return PieceReader{db: db, car: car, blockStore: priorityStore}
}

func (p PieceReader) ReadAt(buf []byte, offset int64) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	if offset < int64(len(p.car.Header)) {
		return copy(buf, p.car.Header[offset:]), nil
	}

	var carBlock model.CarBlock
	err := p.db.WithContext(ctx).Where("car_id = ? AND offset <= ? AND offset + length > ?", p.car.ID, offset, offset).
		First(&carBlock).Error
	if err != nil {
		return 0, errors.Wrap(err, "failed to query for car block")
	}

	written := 0
	varintSize := varint.UvarintSize(carBlock.Varint)
	if offset < int64(carBlock.Offset) + int64(varintSize) {
		varintBytes := varint.ToUvarint(carBlock.Varint)
		written += copy(buf, varintBytes[offset - int64(carBlock.Offset):])
		if written == len(buf) {
			return written, nil
		}
	}

	blockCID, err := cid.Parse(carBlock.CID)
	if err != nil {
		return 0, errors.Wrap(err, "failed to parse CID")
	}
	if offset < int64(carBlock.Offset) + int64(varintSize) + CidLength {
		written += copy(buf[written:], blockCID.Bytes()[offset - int64(carBlock.Offset) - int64(varintSize):])
		if written == len(buf) {
			return written, nil
		}
	}

	// Try to find the block from the blockstore
	block, err := p.blockStore.Get(ctx, blockCID)
	if err != nil {
		return 0, errors.Wrap(err, "failed to get block from blockstore")
	}

	blockBytes := block.RawData()
	written += copy(buf[written:], blockBytes[offset - int64(carBlock.Offset) - int64(varintSize) - CidLength:])
	return written, nil
}
