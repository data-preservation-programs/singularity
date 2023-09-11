//nolint:forcetypeassert
package job

import (
	"context"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type Handler interface {
	StartDagGenHandler(
		ctx context.Context,
		db *gorm.DB,
		id string,
		name string) (*model.Job, error)

	PauseDagGenHandler(
		ctx context.Context,
		db *gorm.DB,
		id string,
		name string) (*model.Job, error)

	StartPackHandler(
		ctx context.Context,
		db *gorm.DB,
		id string,
		name string,
		jobID int64) ([]model.Job, error)

	PausePackHandler(
		ctx context.Context,
		db *gorm.DB,
		id string,
		name string,
		jobID int64) ([]model.Job, error)

	StartScanHandler(
		ctx context.Context,
		db *gorm.DB,
		id string,
		name string) (*model.Job, error)

	PauseScanHandler(
		ctx context.Context,
		db *gorm.DB,
		id string,
		name string) (*model.Job, error)

	GetStatusHandler(ctx context.Context, db *gorm.DB, id string) ([]SourceStatus, error)

	PackHandler(
		ctx context.Context,
		db *gorm.DB,
		jobID uint64) (*model.Car, error)

	PrepareToPackSourceHandler(
		ctx context.Context,
		db *gorm.DB,
		id string,
		name string,
	) error
}

type DefaultHandler struct{}

var Default Handler = &DefaultHandler{}

type MockJob struct {
	mock.Mock
}

var _ Handler = &MockJob{}

func (m *MockJob) PackHandler(ctx context.Context, db *gorm.DB, jobID uint64) (*model.Car, error) {
	args := m.Called(ctx, db, jobID)
	return args.Get(0).(*model.Car), args.Error(1)
}

func (m *MockJob) PrepareToPackSourceHandler(ctx context.Context, db *gorm.DB, id string, name string) error {
	args := m.Called(ctx, db, id, name)
	return args.Error(0)
}

func (m *MockJob) StartScanHandler(ctx context.Context, db *gorm.DB, id string, name string) (*model.Job, error) {
	args := m.Called(ctx, db, id, name)
	return args.Get(0).(*model.Job), args.Error(1)
}

func (m *MockJob) PauseScanHandler(ctx context.Context, db *gorm.DB, id string, name string) (*model.Job, error) {
	args := m.Called(ctx, db, id, name)
	return args.Get(0).(*model.Job), args.Error(1)
}

func (m *MockJob) GetStatusHandler(ctx context.Context, db *gorm.DB, id string) ([]SourceStatus, error) {
	args := m.Called(ctx, db, id)
	return args.Get(0).([]SourceStatus), args.Error(1)
}

func (m *MockJob) StartPackHandler(ctx context.Context, db *gorm.DB, id string, name string, jobID int64) ([]model.Job, error) {
	args := m.Called(ctx, db, id, name, jobID)
	return args.Get(0).([]model.Job), args.Error(1)
}

func (m *MockJob) PausePackHandler(ctx context.Context, db *gorm.DB, id string, name string, jobID int64) ([]model.Job, error) {
	args := m.Called(ctx, db, id, name, jobID)
	return args.Get(0).([]model.Job), args.Error(1)
}

func (m *MockJob) StartDagGenHandler(ctx context.Context, db *gorm.DB, id string, name string) (*model.Job, error) {
	args := m.Called(ctx, db, id, name)
	return args.Get(0).(*model.Job), args.Error(1)
}

func (m *MockJob) PauseDagGenHandler(ctx context.Context, db *gorm.DB, id string, name string) (*model.Job, error) {
	args := m.Called(ctx, db, id, name)
	return args.Get(0).(*model.Job), args.Error(1)
}
