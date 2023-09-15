//nolint:forcetypeassert
package storage

import (
	"context"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type Handler interface {
	CreateStorageHandler(
		ctx context.Context,
		db *gorm.DB,
		storageType string,
		request CreateRequest,
	) (*model.Storage, error)
	ExploreHandler(
		ctx context.Context,
		db *gorm.DB,
		name string,
		path string,
	) ([]DirEntry, error)
	ListStoragesHandler(
		ctx context.Context,
		db *gorm.DB) ([]model.Storage, error)
	RemoveHandler(
		ctx context.Context,
		db *gorm.DB,
		name string) error
	UpdateStorageHandler(
		ctx context.Context,
		db *gorm.DB,
		name string,
		request UpdateRequest,
	) (*model.Storage, error)
	RenameStorageHandler(
		ctx context.Context,
		db *gorm.DB,
		name string,
		request RenameRequest,
	) (*model.Storage, error)
}

type DefaultHandler struct{}

var Default Handler = &DefaultHandler{}

var _ Handler = &MockStorage{}

type MockStorage struct {
	mock.Mock
}

func (m *MockStorage) RenameStorageHandler(ctx context.Context, db *gorm.DB, name string, request RenameRequest) (*model.Storage, error) {
	args := m.Called(ctx, db, name, request)
	return args.Get(0).(*model.Storage), args.Error(1)
}

func (m *MockStorage) CreateStorageHandler(ctx context.Context, db *gorm.DB, storageType string, request CreateRequest) (*model.Storage, error) {
	args := m.Called(ctx, db, storageType, request)
	return args.Get(0).(*model.Storage), args.Error(1)
}

func (m *MockStorage) ExploreHandler(ctx context.Context, db *gorm.DB, name string, path string) ([]DirEntry, error) {
	args := m.Called(ctx, db, name, path)
	return args.Get(0).([]DirEntry), args.Error(1)
}

func (m *MockStorage) ListStoragesHandler(ctx context.Context, db *gorm.DB) ([]model.Storage, error) {
	args := m.Called(ctx, db)
	return args.Get(0).([]model.Storage), args.Error(1)
}

func (m *MockStorage) RemoveHandler(ctx context.Context, db *gorm.DB, name string) error {
	args := m.Called(ctx, db, name)
	return args.Error(0)
}

func (m *MockStorage) UpdateStorageHandler(ctx context.Context, db *gorm.DB, name string, request UpdateRequest) (*model.Storage, error) {
	args := m.Called(ctx, db, name, request)
	return args.Get(0).(*model.Storage), args.Error(1)
}
