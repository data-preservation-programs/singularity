package job

import (
	"context"

	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

type Handler interface {
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

	GetStatusHandler(ctx context.Context, db *gorm.DB, id string) ([]SourceStatus, error)

	PackHandler(
		ctx context.Context,
		db *gorm.DB,
		jobID int64) (*model.Car, error)

	PrepareToPackSourceHandler(
		ctx context.Context,
		db *gorm.DB,
		id string,
		name string,
	) error
}

type DefaultHandler struct{}

var Default Handler = &DefaultHandler{}
