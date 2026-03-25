//nolint:forcetypeassert
package sppool

import (
	"context"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/stretchr/testify/mock"
	"github.com/ybbus/jsonrpc/v3"
	"gorm.io/gorm"
)

type Handler interface {
	CreateHandler(ctx context.Context, db *gorm.DB, request CreateRequest) (*model.SPPool, error)
	ListHandler(ctx context.Context, db *gorm.DB) ([]model.SPPool, error)
	GetHandler(ctx context.Context, db *gorm.DB, id uint32) (*model.SPPool, error)
	UpdateHandler(ctx context.Context, db *gorm.DB, id uint32, request UpdateRequest) (*model.SPPool, error)
	RemoveHandler(ctx context.Context, db *gorm.DB, id uint32) error
	PauseHandler(ctx context.Context, db *gorm.DB, id uint32) (*model.SPPool, error)
	ResumeHandler(ctx context.Context, db *gorm.DB, id uint32) (*model.SPPool, error)
	AddProviderHandler(ctx context.Context, db *gorm.DB, lotusClient jsonrpc.RPCClient, id uint32, request AddProviderRequest) (*model.SPPoolProvider, error)
	RemoveProviderHandler(ctx context.Context, db *gorm.DB, id uint32, providerID uint32) error
	AddPreparationHandler(ctx context.Context, db *gorm.DB, id uint32, request AddPreparationRequest) (*model.SPPoolPreparation, error)
	RemovePreparationHandler(ctx context.Context, db *gorm.DB, id uint32, preparationID uint32) error
}

type DefaultHandler struct{}

var Default Handler = &DefaultHandler{}

var _ Handler = &MockSPPool{}

type MockSPPool struct {
	mock.Mock
}

func (m *MockSPPool) CreateHandler(ctx context.Context, db *gorm.DB, request CreateRequest) (*model.SPPool, error) {
	args := m.Called(ctx, db, request)
	return args.Get(0).(*model.SPPool), args.Error(1)
}

func (m *MockSPPool) ListHandler(ctx context.Context, db *gorm.DB) ([]model.SPPool, error) {
	args := m.Called(ctx, db)
	return args.Get(0).([]model.SPPool), args.Error(1)
}

func (m *MockSPPool) GetHandler(ctx context.Context, db *gorm.DB, id uint32) (*model.SPPool, error) {
	args := m.Called(ctx, db, id)
	return args.Get(0).(*model.SPPool), args.Error(1)
}

func (m *MockSPPool) UpdateHandler(ctx context.Context, db *gorm.DB, id uint32, request UpdateRequest) (*model.SPPool, error) {
	args := m.Called(ctx, db, id, request)
	return args.Get(0).(*model.SPPool), args.Error(1)
}

func (m *MockSPPool) RemoveHandler(ctx context.Context, db *gorm.DB, id uint32) error {
	args := m.Called(ctx, db, id)
	return args.Error(0)
}

func (m *MockSPPool) PauseHandler(ctx context.Context, db *gorm.DB, id uint32) (*model.SPPool, error) {
	args := m.Called(ctx, db, id)
	return args.Get(0).(*model.SPPool), args.Error(1)
}

func (m *MockSPPool) ResumeHandler(ctx context.Context, db *gorm.DB, id uint32) (*model.SPPool, error) {
	args := m.Called(ctx, db, id)
	return args.Get(0).(*model.SPPool), args.Error(1)
}

func (m *MockSPPool) AddProviderHandler(ctx context.Context, db *gorm.DB, lotusClient jsonrpc.RPCClient, id uint32, request AddProviderRequest) (*model.SPPoolProvider, error) {
	args := m.Called(ctx, db, lotusClient, id, request)
	return args.Get(0).(*model.SPPoolProvider), args.Error(1)
}

func (m *MockSPPool) RemoveProviderHandler(ctx context.Context, db *gorm.DB, id uint32, providerID uint32) error {
	args := m.Called(ctx, db, id, providerID)
	return args.Error(0)
}

func (m *MockSPPool) AddPreparationHandler(ctx context.Context, db *gorm.DB, id uint32, request AddPreparationRequest) (*model.SPPoolPreparation, error) {
	args := m.Called(ctx, db, id, request)
	return args.Get(0).(*model.SPPoolPreparation), args.Error(1)
}

func (m *MockSPPool) RemovePreparationHandler(ctx context.Context, db *gorm.DB, id uint32, preparationID uint32) error {
	args := m.Called(ctx, db, id, preparationID)
	return args.Error(0)
}
