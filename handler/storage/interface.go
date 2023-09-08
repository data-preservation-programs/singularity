package storage

import (
	"context"

	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

type Handler interface {
	CreateStorageHandler(
		ctx context.Context,
		db *gorm.DB,
		storageType string,
		request CreateRequest,
	) (*model.Storage, error)
	ExploreHandler(
		ctx context.Context,
		db *gorm.DB,
		name string,
		path string,
	) ([]DirEntry, error)
	ListStoragesHandler(
		ctx context.Context,
		db *gorm.DB) ([]model.Storage, error)
	RemoveHandler(
		ctx context.Context,
		db *gorm.DB,
		name string) error
	UpdateStorageHandler(
		ctx context.Context,
		db *gorm.DB,
		name string,
		config map[string]string,
	) (*model.Storage, error)
}

type DefaultHandler struct{}

var Default Handler = &DefaultHandler{}
