//nolint:forcetypeassert
package statechange

import (
	"context"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

// Handler defines the interface for state change operations
type Handler interface {
	// GetDealStateChangesHandler retrieves state changes for a specific deal
	GetDealStateChangesHandler(ctx context.Context, db *gorm.DB, dealID model.DealID) ([]model.DealStateChange, error)
	
	// ListStateChangesHandler retrieves state changes with filtering and pagination
	ListStateChangesHandler(ctx context.Context, db *gorm.DB, query model.DealStateChangeQuery) (StateChangeResponse, error)
	
	// GetStateChangeStatsHandler returns statistics about state changes
	GetStateChangeStatsHandler(ctx context.Context, db *gorm.DB) (map[string]interface{}, error)
}

// StateChangeResponse represents the response structure for state change listings
type StateChangeResponse struct {
	StateChanges []model.DealStateChange `json:"stateChanges"`
	Total        int64                   `json:"total"`
	Offset       *int                    `json:"offset,omitempty"`
	Limit        *int                    `json:"limit,omitempty"`
}

// DefaultHandler implements the Handler interface
type DefaultHandler struct{}

var Default Handler = &DefaultHandler{}

var _ Handler = &MockStateChange{}

// MockStateChange provides mock implementation for testing
type MockStateChange struct {
	mock.Mock
}

func (m *MockStateChange) GetDealStateChangesHandler(ctx context.Context, db *gorm.DB, dealID model.DealID) ([]model.DealStateChange, error) {
	args := m.Called(ctx, db, dealID)
	return args.Get(0).([]model.DealStateChange), args.Error(1)
}

func (m *MockStateChange) ListStateChangesHandler(ctx context.Context, db *gorm.DB, query model.DealStateChangeQuery) (StateChangeResponse, error) {
	args := m.Called(ctx, db, query)
	return args.Get(0).(StateChangeResponse), args.Error(1)
}

func (m *MockStateChange) GetStateChangeStatsHandler(ctx context.Context, db *gorm.DB) (map[string]interface{}, error) {
	args := m.Called(ctx, db)
	return args.Get(0).(map[string]interface{}), args.Error(1)
}