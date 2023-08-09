package client

import (
	"context"

	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/handler/dataset"
	"github.com/data-preservation-programs/singularity/handler/datasource"
	"github.com/data-preservation-programs/singularity/handler/datasource/inspect"
	"github.com/data-preservation-programs/singularity/model"
)

type DuplicateRecordError = handler.DuplicateRecordError
type InvalidParameterError = handler.InvalidParameterError
type NotFoundError = handler.NotFoundError

type Client interface {
	CreateDataset(ctx context.Context, request dataset.CreateRequest) (*model.Dataset, error)
	CreateLocalSource(ctx context.Context, datasetName string, params datasource.LocalRequest) (*model.Source, error)
	ListSourcesByDataset(ctx context.Context, datasetName string) ([]model.Source, error)
	GetItem(ctx context.Context, id uint64) (*model.Item, error)
	PushItem(ctx context.Context, sourceID uint32, itemInfo datasource.ItemInfo) (*model.Item, error)
	GetSourceChunks(ctx context.Context, sourceID uint32, request inspect.GetSourceChunksRequest) ([]model.Chunk, error)
	Chunk(ctx context.Context, sourceID uint32, request datasource.ChunkRequest) (*model.Chunk, error)
}
