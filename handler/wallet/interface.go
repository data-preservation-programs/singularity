//nolint:forcetypeassert
package wallet

import (
	"context"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/stretchr/testify/mock"
	"github.com/ybbus/jsonrpc/v3"
	"gorm.io/gorm"
)

type Handler interface {
	AttachHandler(
		ctx context.Context,
		db *gorm.DB,
		preparation string,
		wallet string,
	) (*model.Preparation, error)
	CreateHandler(
		ctx context.Context,
		db *gorm.DB,
		lotusClient jsonrpc.RPCClient,
		request CreateRequest,
	) (*model.Wallet, error)
	DetachHandler(
		ctx context.Context,
		db *gorm.DB,
		preparation string,
		wallet string,
	) (*model.Preparation, error)
	ImportHandler(
		ctx context.Context,
		db *gorm.DB,
		lotusClient jsonrpc.RPCClient,
		request ImportRequest,
	) (*model.Wallet, error)
	InitHandler(
		ctx context.Context,
		db *gorm.DB,
		lotusClient jsonrpc.RPCClient,
		address string,
	) (*model.Wallet, error)
	ListHandler(
		ctx context.Context,
		db *gorm.DB,
	) ([]model.Wallet, error)
	ListAttachedHandler(
		ctx context.Context,
		db *gorm.DB,
		preparation string,
	) ([]model.Wallet, error)
	RemoveHandler(
		ctx context.Context,
		db *gorm.DB,
		address string,
	) error
	UpdateHandler(
		ctx context.Context,
		db *gorm.DB,
		address string,
		request UpdateRequest,
	) (*model.Wallet, error)
}

type DefaultHandler struct{}

var Default Handler = &DefaultHandler{}

var _ Handler = &MockWallet{}

type MockWallet struct {
	mock.Mock
}

func (m *MockWallet) AttachHandler(ctx context.Context, db *gorm.DB, preparation string, wallet string) (*model.Preparation, error) {
	args := m.Called(ctx, db, preparation, wallet)
	return args.Get(0).(*model.Preparation), args.Error(1)
}

func (m *MockWallet) CreateHandler(ctx context.Context, db *gorm.DB, lotusClient jsonrpc.RPCClient, request CreateRequest) (*model.Wallet, error) {
	args := m.Called(ctx, db, request)
	return args.Get(0).(*model.Wallet), args.Error(1)
}

func (m *MockWallet) DetachHandler(ctx context.Context, db *gorm.DB, preparation string, wallet string) (*model.Preparation, error) {
	args := m.Called(ctx, db, preparation, wallet)
	return args.Get(0).(*model.Preparation), args.Error(1)
}

func (m *MockWallet) ImportHandler(ctx context.Context, db *gorm.DB, lotusClient jsonrpc.RPCClient, request ImportRequest) (*model.Wallet, error) {
	args := m.Called(ctx, db, lotusClient, request)
	return args.Get(0).(*model.Wallet), args.Error(1)
}

func (m *MockWallet) InitHandler(ctx context.Context, db *gorm.DB, lotusClient jsonrpc.RPCClient, wallet string) (*model.Wallet, error) {
	args := m.Called(ctx, db, lotusClient, wallet)
	return args.Get(0).(*model.Wallet), args.Error(1)
}

func (m *MockWallet) ListHandler(ctx context.Context, db *gorm.DB) ([]model.Wallet, error) {
	args := m.Called(ctx, db)
	return args.Get(0).([]model.Wallet), args.Error(1)
}

func (m *MockWallet) ListAttachedHandler(ctx context.Context, db *gorm.DB, preparation string) ([]model.Wallet, error) {
	args := m.Called(ctx, db, preparation)
	return args.Get(0).([]model.Wallet), args.Error(1)
}

func (m *MockWallet) RemoveHandler(ctx context.Context, db *gorm.DB, address string) error {
	args := m.Called(ctx, db, address)
	return args.Error(0)
}

func (m *MockWallet) UpdateHandler(ctx context.Context, db *gorm.DB, address string, request UpdateRequest) (*model.Wallet, error) {
	args := m.Called(ctx, db, address, request)
	return args.Get(0).(*model.Wallet), args.Error(1)
}
