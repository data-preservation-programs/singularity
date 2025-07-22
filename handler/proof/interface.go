package proof

import (
	"context"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/stretchr/testify/mock"
	"github.com/ybbus/jsonrpc/v3"
	"gorm.io/gorm"
)

type Handler interface {
	ListHandler(ctx context.Context, db *gorm.DB, request ListProofRequest) ([]model.Proof, error)
	SyncHandler(ctx context.Context, db *gorm.DB, lotusClient jsonrpc.RPCClient, request SyncProofRequest) error
}

type DefaultHandler struct{}

var Default Handler = &DefaultHandler{}

var _ Handler = &MockProof{}

type MockProof struct {
	mock.Mock
}

func (m *MockProof) ListHandler(ctx context.Context, db *gorm.DB, request ListProofRequest) ([]model.Proof, error) {
	args := m.Called(ctx, db, request)
	return args.Get(0).([]model.Proof), args.Error(1)
}

func (m *MockProof) SyncHandler(ctx context.Context, db *gorm.DB, lotusClient jsonrpc.RPCClient, request SyncProofRequest) error {
	args := m.Called(ctx, db, lotusClient, request)
	return args.Error(0)
}
