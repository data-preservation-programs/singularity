package file

import (
	"context"

	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

type Handler interface {
	PrepareToPackFileHandler(
		ctx context.Context,
		db *gorm.DB,
		fileID uint64) (int64, error)
	
	GetFileDealsHandler(
		ctx context.Context,
		db *gorm.DB,
		id uint64,
	) ([]model.Deal, error)

	GetFileHandler(
		ctx context.Context,
		db *gorm.DB,
		id uint64,
	) (*model.File, error)

	PushFileHandler(
		ctx context.Context,
		db *gorm.DB,
		preparation string,
		source string,
		fileInfo Info,
	) (*model.File, error)
}

type DefaultHandler struct{}

var Default Handler = &DefaultHandler{}
