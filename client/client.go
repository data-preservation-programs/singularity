package client

import (
	"context"

	"github.com/data-preservation-programs/singularity/handler/dataset"
	"github.com/data-preservation-programs/singularity/handler/datasource"
	"github.com/data-preservation-programs/singularity/handler/datasource/inspect"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
)

type DuplicateRecordError = handlererror.DuplicateRecordError
type InvalidParameterError = handlererror.InvalidParameterError
type NotFoundError = handlererror.NotFoundError

type Client interface {
	CreateDataset(ctx context.Context, request dataset.CreateRequest) (*model.Preparation, error)
	CreateLocalSource(ctx context.Context, datasetName string, params datasource.LocalRequest) (*model.Source, error)
	ListSourcesByDataset(ctx context.Context, datasetName string) ([]model.Source, error)
	GetFile(ctx context.Context, id uint64) (*model.File, error)
	PushFile(ctx context.Context, sourceID uint32, fileInfo datasource.FileInfo) (*model.File, error)
	GetSourcePackJobs(ctx context.Context, sourceID uint32, request inspect.GetSourcePackJobsRequest) ([]model.PackJob, error)
	CreatePackJob(ctx context.Context, sourceID uint32, request datasource.CreatePackJobRequest) (*model.PackJob, error)
}
