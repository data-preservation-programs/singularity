//nolint:forcetypeassert
package dataprep

import (
	"context"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type Handler interface {
	CreatePreparationHandler(
		ctx context.Context,
		db *gorm.DB,
		request CreateRequest,
	) (*model.Preparation, error)

	ExploreHandler(
		ctx context.Context,
		db *gorm.DB,
		id string,
		name string,
		path string,
	) (*ExploreResult, error)

	ListHandler(ctx context.Context, db *gorm.DB) ([]model.Preparation, error)

	AddOutputStorageHandler(ctx context.Context, db *gorm.DB, id string, output string) (*model.Preparation, error)

	RemoveOutputStorageHandler(ctx context.Context, db *gorm.DB, id string, output string) (*model.Preparation, error)

	ListPiecesHandler(
		ctx context.Context,
		db *gorm.DB,
		id string,
	) ([]PieceList, error)

	AddPieceHandler(
		ctx context.Context,
		db *gorm.DB,
		id string,
		request AddPieceRequest,
	) (*model.Car, error)

	AddSourceStorageHandler(ctx context.Context, db *gorm.DB, id string, source string) (*model.Preparation, error)
	ListSchedulesHandler(
		ctx context.Context,
		db *gorm.DB,
		id string) ([]model.Schedule, error)
}

type DefaultHandler struct{}

var Default Handler = &DefaultHandler{}

type MockDataPrep struct {
	mock.Mock
}

func (m *MockDataPrep) ListSchedulesHandler(ctx context.Context, db *gorm.DB, id string) ([]model.Schedule, error) {
	args := m.Called(ctx, db, id)
	return args.Get(0).([]model.Schedule), args.Error(1)
}

func (m *MockDataPrep) CreatePreparationHandler(ctx context.Context, db *gorm.DB, request CreateRequest) (*model.Preparation, error) {
	args := m.Called(ctx, db, request)
	return args.Get(0).(*model.Preparation), args.Error(1)
}

func (m *MockDataPrep) ExploreHandler(ctx context.Context, db *gorm.DB, id string, name string, path string) (*ExploreResult, error) {
	args := m.Called(ctx, db, id, name, path)
	return args.Get(0).(*ExploreResult), args.Error(1)
}

func (m *MockDataPrep) ListHandler(ctx context.Context, db *gorm.DB) ([]model.Preparation, error) {
	args := m.Called(ctx, db)
	return args.Get(0).([]model.Preparation), args.Error(1)
}

func (m *MockDataPrep) AddOutputStorageHandler(ctx context.Context, db *gorm.DB, id string, output string) (*model.Preparation, error) {
	args := m.Called(ctx, db, id, output)
	return args.Get(0).(*model.Preparation), args.Error(1)
}

func (m *MockDataPrep) RemoveOutputStorageHandler(ctx context.Context, db *gorm.DB, id string, output string) (*model.Preparation, error) {
	args := m.Called(ctx, db, id, output)
	return args.Get(0).(*model.Preparation), args.Error(1)
}

func (m *MockDataPrep) ListPiecesHandler(ctx context.Context, db *gorm.DB, id string) ([]PieceList, error) {
	args := m.Called(ctx, db, id)
	return args.Get(0).([]PieceList), args.Error(1)
}

func (m *MockDataPrep) AddPieceHandler(ctx context.Context, db *gorm.DB, id string, request AddPieceRequest) (*model.Car, error) {
	args := m.Called(ctx, db, id, request)
	return args.Get(0).(*model.Car), args.Error(1)
}

func (m *MockDataPrep) AddSourceStorageHandler(ctx context.Context, db *gorm.DB, id string, source string) (*model.Preparation, error) {
	args := m.Called(ctx, db, id, source)
	return args.Get(0).(*model.Preparation), args.Error(1)
}

var _ Handler = &MockDataPrep{}
