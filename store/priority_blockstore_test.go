package store

import (
	"context"
	"errors"
	"testing"

	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockBlockstore struct {
	mock.Mock
}

func (m *mockBlockstore) DeleteBlock(ctx context.Context, c cid.Cid) error {
	//TODO implement me
	panic("implement me")
}

func (m *mockBlockstore) Put(ctx context.Context, block blocks.Block) error {
	//TODO implement me
	panic("implement me")
}

func (m *mockBlockstore) PutMany(ctx context.Context, i []blocks.Block) error {
	//TODO implement me
	panic("implement me")
}

func (m *mockBlockstore) AllKeysChan(ctx context.Context) (<-chan cid.Cid, error) {
	//TODO implement me
	panic("implement me")
}

func (m *mockBlockstore) HashOnRead(enabled bool) {
	//TODO implement me
	panic("implement me")
}

func (m *mockBlockstore) Has(ctx context.Context, c cid.Cid) (bool, error) {
	args := m.Called(ctx, c)
	return args.Bool(0), args.Error(1)
}

func (m *mockBlockstore) Get(ctx context.Context, c cid.Cid) (blocks.Block, error) {
	args := m.Called(ctx, c)
	return args.Get(0).(blocks.Block), args.Error(1)
}

func (m *mockBlockstore) GetSize(ctx context.Context, c cid.Cid) (int, error) {
	args := m.Called(ctx, c)
	return args.Int(0), args.Error(1)
}

var nullBlock = (*blocks.BasicBlock)(nil)

func TestHas(t *testing.T) {
	ctx := context.Background()
	cid1 := cid.Cid{}

	mockStore1 := &mockBlockstore{}
	mockStore1.On("Has", ctx, cid1).Return(false, nil)

	mockStore2 := &mockBlockstore{}
	mockStore2.On("Has", ctx, cid1).Return(true, nil)

	priorityStore := NewPriorityBlockStore(mockStore1, mockStore2)
	has, err := priorityStore.Has(ctx, cid1)

	assert.NoError(t, err)
	assert.True(t, has)
}

func TestGet(t *testing.T) {
	ctx := context.Background()
	cid1 := cid.Cid{}
	block1 := blocks.NewBlock([]byte("block1"))

	mockStore1 := &mockBlockstore{}
	mockStore1.On("Get", ctx, cid1).Return(nullBlock, errors.New("not found"))

	mockStore2 := &mockBlockstore{}
	mockStore2.On("Get", ctx, cid1).Return(block1, nil)

	priorityStore := NewPriorityBlockStore(mockStore1, mockStore2)
	block, err := priorityStore.Get(ctx, cid1)

	assert.NoError(t, err)
	assert.Equal(t, block1, block)
}

func TestGetSize(t *testing.T) {
	ctx := context.Background()
	cid1 := cid.Cid{}

	mockStore1 := &mockBlockstore{}
	mockStore1.On("GetSize", ctx, cid1).Return(0, errors.New("not found"))

	mockStore2 := &mockBlockstore{}
	mockStore2.On("GetSize", ctx, cid1).Return(100, nil)

	priorityStore := NewPriorityBlockStore(mockStore1, mockStore2)
	size, err := priorityStore.GetSize(ctx, cid1)

	assert.NoError(t, err)
	assert.Equal(t, 100, size)
}
