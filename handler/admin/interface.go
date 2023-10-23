//nolint:forcetypeassert
package admin

import (
	"context"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type Handler interface {
	InitHandler(ctx context.Context, db *gorm.DB) error
	ResetHandler(ctx context.Context, db *gorm.DB) error
	SetIdentityHandler(ctx context.Context, db *gorm.DB, request SetIdentityRequest) error
}

type DefaultHandler struct{}

var Default Handler = &DefaultHandler{}

var _ Handler = &MockAdmin{}

type MockAdmin struct {
	mock.Mock
}

func (m *MockAdmin) SetIdentityHandler(ctx context.Context, db *gorm.DB, request SetIdentityRequest) error {
	args := m.Called(ctx, db, request)
	return args.Error(0)
}

func (m *MockAdmin) InitHandler(ctx context.Context, db *gorm.DB) error {
	args := m.Called(ctx, db)
	return args.Error(0)
}

func (m *MockAdmin) ResetHandler(ctx context.Context, db *gorm.DB) error {
	args := m.Called(ctx, db)
	return args.Error(0)
}
