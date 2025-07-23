//nolint:forcetypeassert
package dealtemplate

import (
	"context"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type Handler interface {
	CreateHandler(ctx context.Context, db *gorm.DB, request CreateRequest) (*model.DealTemplate, error)
	ListHandler(ctx context.Context, db *gorm.DB) ([]model.DealTemplate, error)
	GetHandler(ctx context.Context, db *gorm.DB, idOrName string) (*model.DealTemplate, error)
	UpdateHandler(ctx context.Context, db *gorm.DB, idOrName string, request UpdateRequest) (*model.DealTemplate, error)
	DeleteHandler(ctx context.Context, db *gorm.DB, idOrName string) error
	ApplyTemplateToPreparation(template *model.DealTemplate, prep *model.Preparation)
}

type DefaultHandler struct{}

var Default Handler = &DefaultHandler{}

var _ Handler = &MockDealTemplate{}

type MockDealTemplate struct {
	mock.Mock
}

func (m *MockDealTemplate) CreateHandler(ctx context.Context, db *gorm.DB, request CreateRequest) (*model.DealTemplate, error) {
	args := m.Called(ctx, db, request)
	return args.Get(0).(*model.DealTemplate), args.Error(1)
}

func (m *MockDealTemplate) ListHandler(ctx context.Context, db *gorm.DB) ([]model.DealTemplate, error) {
	args := m.Called(ctx, db)
	return args.Get(0).([]model.DealTemplate), args.Error(1)
}

func (m *MockDealTemplate) GetHandler(ctx context.Context, db *gorm.DB, idOrName string) (*model.DealTemplate, error) {
	args := m.Called(ctx, db, idOrName)
	return args.Get(0).(*model.DealTemplate), args.Error(1)
}

func (m *MockDealTemplate) UpdateHandler(ctx context.Context, db *gorm.DB, idOrName string, request UpdateRequest) (*model.DealTemplate, error) {
	args := m.Called(ctx, db, idOrName, request)
	return args.Get(0).(*model.DealTemplate), args.Error(1)
}

func (m *MockDealTemplate) DeleteHandler(ctx context.Context, db *gorm.DB, idOrName string) error {
	args := m.Called(ctx, db, idOrName)
	return args.Error(0)
}

func (m *MockDealTemplate) ApplyTemplateToPreparation(template *model.DealTemplate, prep *model.Preparation) {
	m.Called(template, prep)
}
