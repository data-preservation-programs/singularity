package client

import (
	"context"

	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/handler/dataset"
	"github.com/data-preservation-programs/singularity/handler/datasource"
	"github.com/data-preservation-programs/singularity/handler/datasource/inspect"
	"github.com/data-preservation-programs/singularity/handler/deal/schedule"
	"github.com/data-preservation-programs/singularity/handler/wallet"
	"github.com/data-preservation-programs/singularity/model"
)

type DuplicateRecordError = handler.DuplicateRecordError
type InvalidParameterError = handler.InvalidParameterError
type NotFoundError = handler.NotFoundError

//nolint:interfacebloat
type Client interface {
	CreateDataset(ctx context.Context, request dataset.CreateRequest) (*model.Dataset, error)
	CreateLocalSource(ctx context.Context, datasetName string, params datasource.LocalRequest) (*model.Source, error)
	ListSourcesByDataset(ctx context.Context, datasetName string) ([]model.Source, error)
	GetFile(ctx context.Context, id uint64) (*model.File, error)
	PushFile(ctx context.Context, sourceID uint32, fileInfo datasource.FileInfo) (*model.File, error)
	GetSourcePackJobs(ctx context.Context, sourceID uint32, request inspect.GetSourcePackJobsRequest) ([]model.PackJob, error)
	PrepareToPackFile(ctx context.Context, fileID uint64) (int64, error)
	PrepareToPackSource(ctx context.Context, sourceID uint32) error
	Pack(ctx context.Context, packJobID uint32) ([]model.Car, error)
	GetFileDeals(ctx context.Context, id uint64) ([]model.Deal, error)
	ImportWallet(ctx context.Context, request wallet.ImportRequest) (*model.Wallet, error)
	AddWalletToDataset(ctx context.Context, datasetName string, wallet string) (*model.WalletAssignment, error)
	ListWallets(ctx context.Context) ([]model.Wallet, error)
	ListWalletsByDataset(ctx context.Context, datasetName string) ([]model.Wallet, error)
	CreateSchedule(ctx context.Context, request schedule.CreateRequest) (*model.Schedule, error)
	ListSchedules(ctx context.Context) ([]model.Schedule, error)
}
