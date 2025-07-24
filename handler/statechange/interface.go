//nolint:forcetypeassert
package statechange

import (
	"context"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

// StateChangeResponse represents the response structure for state change API calls
type StateChangeResponse struct {
	StateChanges []model.DealStateChange `json:"stateChanges"`
	Total        int64                   `json:"total"`
	Offset       *int                    `json:"offset,omitempty"`
	Limit        *int                    `json:"limit,omitempty"`
}

type Handler interface {
	ListStateChangesHandler(
		ctx context.Context,
		db *gorm.DB,
		query model.DealStateChangeQuery,
	) (StateChangeResponse, error)

	GetDealStateChangesHandler(
		ctx context.Context,
		db *gorm.DB,
		dealID model.DealID,
	) ([]model.DealStateChange, error)

	GetStateChangeStatsHandler(
		ctx context.Context,
		db *gorm.DB,
	) (map[string]interface{}, error)
}

type DefaultHandler struct{}

var Default Handler = &DefaultHandler{}

type MockStateChange struct {
	mock.Mock
}

func (m *MockStateChange) ListStateChangesHandler(
	ctx context.Context,
	db *gorm.DB,
	query model.DealStateChangeQuery,
) (StateChangeResponse, error) {
	args := m.Called(ctx, db, query)
	return args.Get(0).(StateChangeResponse), args.Error(1)
}

func (m *MockStateChange) GetDealStateChangesHandler(
	ctx context.Context,
	db *gorm.DB,
	dealID model.DealID,
) ([]model.DealStateChange, error) {
	args := m.Called(ctx, db, dealID)
	return args.Get(0).([]model.DealStateChange), args.Error(1)
}

func (m *MockStateChange) GetStateChangeStatsHandler(
	ctx context.Context,
	db *gorm.DB,
) (map[string]interface{}, error) {
	args := m.Called(ctx, db)
	return args.Get(0).(map[string]interface{}), args.Error(1)
}

var _ Handler = &MockStateChange{}
