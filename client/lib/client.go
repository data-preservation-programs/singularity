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
	return dataset.CreateHandler(c.db.WithContext(ctx), request)
}

func (c *Client) ListSourcesByDataset(ctx context.Context, datasetName string) ([]model.Source, error) {
	return dshandler.ListSourcesByDatasetHandler(c.db, datasetName)
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
	return dshandler.CreateDatasourceHandler(c.db.WithContext(ctx), ctx, c.datasourceHandlerResolver, "local", datasetName, paramsMap)
}

func (c *Client) GetSourceChunks(ctx context.Context, sourceID uint32) ([]model.Chunk, error) {
	return inspect.GetSourceChunksHandler(c.db.WithContext(ctx), strconv.FormatUint(uint64(sourceID), 10))
}

func (c *Client) GetSourceItems(ctx context.Context, sourceID uint32) ([]model.Item, error) {
	return inspect.GetSourceItemsHandler(c.db.WithContext(ctx), strconv.FormatUint(uint64(sourceID), 10))
}

func (c *Client) GetItem(ctx context.Context, id uint64) (*model.Item, error) {
	return inspect.GetSourceItemDetailHandler(c.db.WithContext(ctx), strconv.FormatUint(id, 10))
}

func (c *Client) PushItem(ctx context.Context, sourceID uint32, itemInfo dshandler.ItemInfo) (*model.Item, error) {
	return dshandler.PushItemHandler(c.db.WithContext(ctx), ctx, c.datasourceHandlerResolver, sourceID, itemInfo)
}

func (c *Client) Chunk(ctx context.Context, sourceID uint32, request dshandler.ChunkRequest) (*model.Chunk, error) {
	return dshandler.ChunkHandler(c.db.WithContext(ctx), strconv.FormatUint(uint64(sourceID), 10), request)
}

var _ client.Client = (*Client)(nil)
