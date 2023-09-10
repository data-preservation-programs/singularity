package deal

import (
	"context"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/replication"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type Handler interface {
	ListHandler(ctx context.Context, db *gorm.DB, request ListDealRequest) ([]model.Deal, error)
	SendManualHandler(
		ctx context.Context,
		db *gorm.DB,
		dealMaker replication.DealMaker,
		request Proposal,
	) (*model.Deal, error)
}

type DefaultHandler struct{}

var Default Handler = &DefaultHandler{}

var _ Handler = &MockDeal{}

type MockDeal struct {
	mock.Mock
}

func (m *MockDeal) ListHandler(ctx context.Context, db *gorm.DB, request ListDealRequest) ([]model.Deal, error) {
	args := m.Called(ctx, db, request)
	return args.Get(0).([]model.Deal), args.Error(1)
}

func (m *MockDeal) SendManualHandler(ctx context.Context, db *gorm.DB, dealMaker replication.DealMaker, request Proposal) (*model.Deal, error) {
	args := m.Called(ctx, db, dealMaker, request)
	return args.Get(0).(*model.Deal), args.Error(1)
}
