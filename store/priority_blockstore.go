package store

import (
	"context"
	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-ipfs-blockstore"
)

type PriorityBlockStore struct {
	stores []blockstore.Blockstore
}

func (p PriorityBlockStore) Has(ctx context.Context, cid cid.Cid) (bool, error) {
	errors := make([]error, 0)
	for _, store := range p.stores {
		has, err := store.Has(ctx, cid)
		if err != nil {
			errors = append(errors, err)
			continue
		}
		if has {
			return true, nil
		}
	}
	return false, AggregateError{Errors: errors}
}

func (p PriorityBlockStore) Get(ctx context.Context, cid cid.Cid) (blocks.Block, error) {
	errors := make([]error, 0)
	for _, store := range p.stores {
		block, err := store.Get(ctx, cid)
		if err != nil {
			errors = append(errors, err)
			continue
		}
		return block, nil
	}
	return nil, AggregateError{Errors: errors}
}

func (p PriorityBlockStore) GetSize(ctx context.Context, cid cid.Cid) (int, error) {
	errors := make([]error, 0)
	for _, store := range p.stores {
		size, err := store.GetSize(ctx, cid)
		if err != nil {
			errors = append(errors, err)
			continue
		}
		return size, nil
	}
	return 0, AggregateError{Errors: errors}
}

func (p PriorityBlockStore) Put(ctx context.Context, block blocks.Block) error {
	//TODO implement me
	panic("implement me")
}

func (p PriorityBlockStore) PutMany(ctx context.Context, blocks []blocks.Block) error {
	//TODO implement me
	panic("implement me")
}

func (p PriorityBlockStore) AllKeysChan(ctx context.Context) (<-chan cid.Cid, error) {
	//TODO implement me
	panic("implement me")
}

func (p PriorityBlockStore) HashOnRead(enabled bool) {
	//TODO implement me
	panic("implement me")
}

func (p PriorityBlockStore) DeleteBlock(ctx context.Context, cid cid.Cid) error {
	//TODO implement me
	panic("implement me")
}

func NewPriorityBlockStore(stores ...blockstore.Blockstore) PriorityBlockStore {
	return PriorityBlockStore{stores}
}
