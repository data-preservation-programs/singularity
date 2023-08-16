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

func (c *Client) GetSourceChunks(ctx context.Context, sourceID uint32, request inspect.GetSourceChunksRequest) ([]model.Chunk, error) {
	return inspect.GetSourceChunksHandler(ctx, c.db.WithContext(ctx), sourceID, request)
}
func (c *Client) GetSourceItems(ctx context.Context, sourceID uint32) ([]model.Item, error) {
	return inspect.GetSourceItemsHandler(ctx, c.db.WithContext(ctx), strconv.FormatUint(uint64(sourceID), 10))
}

func (c *Client) GetItem(ctx context.Context, id uint64) (*model.Item, error) {
	return inspect.GetSourceItemDetailHandler(ctx, c.db.WithContext(ctx), strconv.FormatUint(id, 10))
}

func (c *Client) PushItem(ctx context.Context, sourceID uint32, itemInfo dshandler.ItemInfo) (*model.Item, error) {
	return dshandler.PushItemHandler(ctx, c.db.WithContext(ctx), c.datasourceHandlerResolver, sourceID, itemInfo)
}

func (c *Client) Chunk(ctx context.Context, sourceID uint32, request dshandler.ChunkRequest) (*model.Chunk, error) {
	return dshandler.ChunkHandler(ctx, c.db.WithContext(ctx), strconv.FormatUint(uint64(sourceID), 10), request)
}

func (c *Client) Pack(ctx context.Context, chunkID uint64) ([]model.Car, error) {
	return dshandler.PackHandler(c.db.WithContext(ctx), ctx, c.datasourceHandlerResolver, chunkID)
}

var _ client.Client = (*Client)(nil)
