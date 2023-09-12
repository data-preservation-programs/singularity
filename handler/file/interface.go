//nolint:forcetypeassert
package file

import (
	"context"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type Handler interface {
	PrepareToPackFileHandler(
		ctx context.Context,
		db *gorm.DB,
		fileID uint64) (int64, error)

	GetFileDealsHandler(
		ctx context.Context,
		db *gorm.DB,
		id uint64,
	) ([]model.Deal, error)

	GetFileHandler(
		ctx context.Context,
		db *gorm.DB,
		id uint64,
	) (*model.File, error)

	PushFileHandler(
		ctx context.Context,
		db *gorm.DB,
		preparation string,
		source string,
		fileInfo Info,
	) (*model.File, error)
}

type DefaultHandler struct{}

var Default Handler = &DefaultHandler{}

var _ Handler = &MockFile{}

type MockFile struct {
	mock.Mock
}

func (m *MockFile) PrepareToPackFileHandler(ctx context.Context, db *gorm.DB, fileID uint64) (int64, error) {
	args := m.Called(ctx, db, fileID)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockFile) PushFileHandler(ctx context.Context, db *gorm.DB, preparation string, source string, fileInfo Info) (*model.File, error) {
	args := m.Called(ctx, db, preparation, source, fileInfo)
	return args.Get(0).(*model.File), args.Error(1)
}

func (m *MockFile) GetFileHandler(ctx context.Context, db *gorm.DB, id uint64) (*model.File, error) {
	args := m.Called(ctx, db, id)
	return args.Get(0).(*model.File), args.Error(1)
}

func (m *MockFile) GetFileDealsHandler(
	ctx context.Context,
	db *gorm.DB,
	id uint64,
) ([]model.Deal, error) {
	args := m.Called(ctx, db, id)
	return args.Get(0).([]model.Deal), args.Error(1)
}
