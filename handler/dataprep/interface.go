package dataprep

import (
	"context"

	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

type Handler interface {
	CreatePreparationHandler(
		ctx context.Context,
		request handler.Request[CreateRequest],
		dep handler.Dependency,
	) (*model.Preparation, error)

	ExploreHandler(
		ctx context.Context,
		db *gorm.DB,
		id string,
		name string,
		path string,
	) (*ExploreResult, error)

	ListHandler(ctx context.Context, db *gorm.DB) ([]model.Preparation, error)

	AddOutputStorageHandler(ctx context.Context, db *gorm.DB, id string, output string) (*model.Preparation, error)

	RemoveOutputStorageHandler(ctx context.Context, db *gorm.DB, id string, output string) (*model.Preparation, error)

	ListPiecesHandler(
		ctx context.Context,
		db *gorm.DB,
		id string,
	) ([]PieceList, error)

	AddPieceHandler(
		ctx context.Context,
		db *gorm.DB,
		id string,
		request AddPieceRequest,
	) (*model.Car, error)

	AddSourceStorageHandler(ctx context.Context, db *gorm.DB, id string, source string) (*model.Preparation, error)
	ListSchedulesHandler(
		ctx context.Context,
		db *gorm.DB,
		id string) ([]model.Schedule, error)
}

type DefaultHandler struct{}

var Default Handler = &DefaultHandler{}
