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
		id uint32,
		name string) (*model.Job, error)

	PauseDagGenHandler(
		ctx context.Context,
		db *gorm.DB,
		id uint32,
		name string) (*model.Job, error)

	ExploreHandler(
		ctx context.Context,
		db *gorm.DB,
		id uint32,
		name string,
		path string,
	) (*ExploreResult, error)

	ListHandler(ctx context.Context, db *gorm.DB) ([]model.Preparation, error)

	AddOutputStorageHandler(ctx context.Context, db *gorm.DB, id uint32, output string) (*model.Preparation, error)

	RemoveOutputStorageHandler(ctx context.Context, db *gorm.DB, id uint32, output string) (*model.Preparation, error)

	StartPackHandler(
		ctx context.Context,
		db *gorm.DB,
		id uint32,
		name string,
		jobID int64) ([]model.Job, error)

	PausePackHandler(
		ctx context.Context,
		db *gorm.DB,
		id uint32,
		name string,
		jobID int64) ([]model.Job, error)

	ListPiecesHandler(
		ctx context.Context,
		db *gorm.DB,
		id uint32,
	) ([]PieceList, error)

	AddPieceHandler(
		ctx context.Context,
		db *gorm.DB,
		id uint32,
		request AddPieceRequest,
	) (*model.Car, error)

	StartScanHandler(
		ctx context.Context,
		db *gorm.DB,
		id uint32,
		name string) (*model.Job, error)

	PauseScanHandler(
		ctx context.Context,
		db *gorm.DB,
		id uint32,
		name string) (*model.Job, error)

	AddSourceStorageHandler(ctx context.Context, db *gorm.DB, id uint32, source string) (*model.Preparation, error)

	GetStatusHandler(ctx context.Context, db *gorm.DB, id uint32) ([]SourceStatus, error)
}

type DefaultHandler struct{}

var Default Handler = &DefaultHandler{}
