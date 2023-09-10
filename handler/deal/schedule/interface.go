package schedule

import (
	"context"

	"github.com/data-preservation-programs/singularity/model"
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
		scheduleID string,
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
	ResumeHandler(
		ctx context.Context,
		db *gorm.DB,
		scheduleID uint32,
	) (*model.Schedule, error)
}

type DefaultHandler struct{}

var Default Handler = &DefaultHandler{}
