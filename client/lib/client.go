package libclient

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/data-preservation-programs/singularity/client"
	"github.com/data-preservation-programs/singularity/datasource"
	"github.com/data-preservation-programs/singularity/handler/dataset"
	dshandler "github.com/data-preservation-programs/singularity/handler/datasource"
	"github.com/data-preservation-programs/singularity/handler/datasource/inspect"
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

type Client struct {
	db                        *gorm.DB
	datasourceHandlerResolver datasource.HandlerResolver
}

func NewClient(db *gorm.DB) (*Client, error) {
	if err := model.AutoMigrate(db); err != nil {
		return nil, err
	}
	return &Client{
		db:                        db,
		datasourceHandlerResolver: &datasource.DefaultHandlerResolver{},
	}, nil
}

func (c *Client) CreateDataset(ctx context.Context, request dataset.CreateRequest) (*model.Dataset, error) {
	return dataset.CreateHandler(ctx, c.db.WithContext(ctx), request)
}

func (c *Client) ListSourcesByDataset(ctx context.Context, datasetName string) ([]model.Source, error) {
	return dshandler.ListSourcesByDatasetHandler(ctx, c.db, datasetName)
}

func (c *Client) CreateLocalSource(ctx context.Context, datasetName string, params dshandler.LocalRequest) (*model.Source, error) {
	paramJSON, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	paramsMap := map[string]interface{}{}
	err = json.Unmarshal(paramJSON, &paramsMap)
	if err != nil {
		return nil, err
	}
	return dshandler.CreateDatasourceHandler(ctx, c.db.WithContext(ctx), "local", datasetName, paramsMap)
}

func (c *Client) GetSourcePackJobs(ctx context.Context, sourceID uint32, request inspect.GetSourcePackJobsRequest) ([]model.PackJob, error) {
	return inspect.GetSourcePackJobsHandler(ctx, c.db.WithContext(ctx), sourceID, request)
}
func (c *Client) GetSourceFiles(ctx context.Context, sourceID uint32) ([]model.File, error) {
	return inspect.GetSourceFilesHandler(ctx, c.db.WithContext(ctx), strconv.FormatUint(uint64(sourceID), 10))
}

func (c *Client) GetFile(ctx context.Context, id uint64) (*model.File, error) {
	return inspect.GetSourceFileDetailHandler(ctx, c.db.WithContext(ctx), strconv.FormatUint(id, 10))
}

func (c *Client) GetFileDeals(ctx context.Context, id uint64) ([]model.Deal, error) {
	return inspect.GetFileDealsHandler(c.db.WithContext(ctx), id)
}

func (c *Client) PushFile(ctx context.Context, sourceID uint32, fileInfo dshandler.FileInfo) (*model.File, error) {
	return dshandler.PushFileHandler(ctx, c.db.WithContext(ctx), c.datasourceHandlerResolver, sourceID, fileInfo)
}

func (c *Client) CreatePackJob(ctx context.Context, sourceID uint32, request dshandler.CreatePackJobRequest) (*model.PackJob, error) {
	return dshandler.CreatePackJobHandler(ctx, c.db.WithContext(ctx), strconv.FormatUint(uint64(sourceID), 10), request)
}

func (c *Client) Pack(ctx context.Context, packJobID uint32) ([]model.Car, error) {
	return dshandler.PackHandler(c.db.WithContext(ctx), ctx, c.datasourceHandlerResolver, packJobID)
}

var _ client.Client = (*Client)(nil)
