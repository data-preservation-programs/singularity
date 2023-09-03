package dataprep

import (
	"context"

	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

type Handler interface {
	CreatePreparationHandler(
		ctx context.Context,
		db *gorm.DB,
		request CreateRequest,
	) (*model.Preparation, error)

	StartDagGenHandler(
		ctx context.Context,
		db *gorm.DB,
		id string,
		name string) (*model.Job, error)

	PauseDagGenHandler(
		ctx context.Context,
		db *gorm.DB,
		id string,
		name string) (*model.Job, error)

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

	StartPackHandler(
		ctx context.Context,
		db *gorm.DB,
		id string,
		name string,
		jobID int64) ([]model.Job, error)

	PausePackHandler(
		ctx context.Context,
		db *gorm.DB,
		id string,
		name string,
		jobID int64) ([]model.Job, error)

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

	StartScanHandler(
		ctx context.Context,
		db *gorm.DB,
		id string,
		name string) (*model.Job, error)

	PauseScanHandler(
		ctx context.Context,
		db *gorm.DB,
		id string,
		name string) (*model.Job, error)

	AddSourceStorageHandler(ctx context.Context, db *gorm.DB, id string, source string) (*model.Preparation, error)

	GetStatusHandler(ctx context.Context, db *gorm.DB, id string) ([]SourceStatus, error)
}

type DefaultHandler struct{}

var Default Handler = &DefaultHandler{}
