package wallet

import (
	"context"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/ybbus/jsonrpc/v3"
	"gorm.io/gorm"
)

type Handler interface {
	AttachHandler(
		ctx context.Context,
		db *gorm.DB,
		preparationID uint32,
		wallet string,
	) (*model.Preparation, error)
	DetachHandler(
		ctx context.Context,
		db *gorm.DB,
		preparationID uint32,
		wallet string,
	) (*model.Preparation, error)
	ImportHandler(
		ctx context.Context,
		db *gorm.DB,
		lotusClient jsonrpc.RPCClient,
		request ImportRequest,
	) (*model.Wallet, error)
	ListHandler(
		ctx context.Context,
		db *gorm.DB,
	) ([]model.Wallet, error)
	ListAttachedHandler(
		ctx context.Context,
		db *gorm.DB,
		preparationID uint32,
	) ([]model.Wallet, error)
	RemoveHandler(
		ctx context.Context,
		db *gorm.DB,
		address string,
	) error
}

type DefaultHandler struct{}

var Default Handler = &DefaultHandler{}
