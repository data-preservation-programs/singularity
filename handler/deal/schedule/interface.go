//nolint:forcetypeassert
package schedule

import (
	"context"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/stretchr/testify/mock"
	"github.com/ybbus/jsonrpc/v3"
	"gorm.io/gorm"
)

type Handler interface {
	CreateHandler(
		ctx context.Context,
		db *gorm.DB,
		lotusClient jsonrpc.RPCClient,
		request CreateRequest,
	) (*model.Schedule, error)
	UpdateHandler(
		ctx context.Context,
		db *gorm.DB,
		scheduleID uint32,
		request UpdateRequest,
	) (*model.Schedule, error)
	ListHandler(
		ctx context.Context,
		db *gorm.DB,
	) ([]model.Schedule, error)
	PauseHandler(
		ctx context.Context,
		db *gorm.DB,
		scheduleID uint32,
	) (*model.Schedule, error)
	RemoveHandler(
		ctx context.Context,
		db *gorm.DB,
		scheduleID uint32,
	) error
	ResumeHandler(
		ctx context.Context,
		db *gorm.DB,
		scheduleID uint32,
	) (*model.Schedule, error)
}

type DefaultHandler struct{}

var Default Handler = &DefaultHandler{}

var _ Handler = &MockSchedule{}

type MockSchedule struct {
	mock.Mock
}

func (m *MockSchedule) RemoveHandler(ctx context.Context, db *gorm.DB, scheduleID uint32) error {
	args := m.Called(ctx, db, scheduleID)
	return args.Error(0)
}

func (m *MockSchedule) UpdateHandler(ctx context.Context, db *gorm.DB, scheduleID uint32, request UpdateRequest) (*model.Schedule, error) {
	args := m.Called(ctx, db, scheduleID, request)
	return args.Get(0).(*model.Schedule), args.Error(1)
}

func (m *MockSchedule) CreateHandler(ctx context.Context, db *gorm.DB, lotusClient jsonrpc.RPCClient, request CreateRequest) (*model.Schedule, error) {
	args := m.Called(ctx, db, lotusClient, request)
	return args.Get(0).(*model.Schedule), args.Error(1)
}

func (m *MockSchedule) ListHandler(ctx context.Context, db *gorm.DB) ([]model.Schedule, error) {
	args := m.Called(ctx, db)
	return args.Get(0).([]model.Schedule), args.Error(1)
}

func (m *MockSchedule) PauseHandler(ctx context.Context, db *gorm.DB, scheduleID uint32) (*model.Schedule, error) {
	args := m.Called(ctx, db, scheduleID)
	return args.Get(0).(*model.Schedule), args.Error(1)
}

func (m *MockSchedule) ResumeHandler(ctx context.Context, db *gorm.DB, scheduleID uint32) (*model.Schedule, error) {
	args := m.Called(ctx, db, scheduleID)
	return args.Get(0).(*model.Schedule), args.Error(1)
}
