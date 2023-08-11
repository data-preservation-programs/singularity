package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/pkg/errors"

	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/handler/dataset"
	"github.com/data-preservation-programs/singularity/handler/datasource"
	"github.com/data-preservation-programs/singularity/handler/datasource/inspect"
	"github.com/data-preservation-programs/singularity/model"
)

type Client struct {
	client    *http.Client
	serverURL string
}

func NewHTTPClient(client *http.Client, serverURL string) *Client {
	return &Client{
		client:    client,
		serverURL: serverURL,
	}
}

func (c *Client) jsonRequest(ctx context.Context, method string, endpoint string, request any) (*http.Response, error) {
	jsonParams, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	httpRequest, err := http.NewRequestWithContext(ctx, method, endpoint, bytes.NewReader(jsonParams))
	if err != nil {
		return nil, err
	}
	httpRequest.Header.Add("Content-Type", "application/json")
	return c.client.Do(httpRequest)
}

func (c *Client) CreateDataset(ctx context.Context, request dataset.CreateRequest) (*model.Dataset, error) {
	response, err := c.jsonRequest(ctx, http.MethodPost, c.serverURL+"/api/dataset", request)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = response.Body.Close()
	}()
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return nil, parseHTTPError(response)
	}
	var dataset model.Dataset
	err = json.NewDecoder(response.Body).Decode(&dataset)
	if err != nil {
		return nil, err
	}
	return &dataset, nil
}

func (c *Client) CreateLocalSource(ctx context.Context, datasetName string, params datasource.LocalRequest) (*model.Source, error) {
	response, err := c.jsonRequest(ctx, http.MethodPost, c.serverURL+"/api/source/local/dataset/"+datasetName, params)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = response.Body.Close()
	}()
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return nil, parseHTTPError(response)
	}
	var source model.Source
	err = json.NewDecoder(response.Body).Decode(&source)
	if err != nil {
		return nil, err
	}
	return &source, nil
}

func (c *Client) ListSourcesByDataset(ctx context.Context, datasetName string) ([]model.Source, error) {
	httpRequest, err := http.NewRequestWithContext(ctx, http.MethodGet, c.serverURL+"/api/dataset/"+datasetName+"/sources", nil)
	if err != nil {
		return nil, err
	}
	response, err := c.client.Do(httpRequest)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = response.Body.Close()
	}()
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return nil, parseHTTPError(response)
	}
	var sources []model.Source
	err = json.NewDecoder(response.Body).Decode(&sources)
	if err != nil {
		return nil, err
	}
	return sources, nil
}

func (c *Client) GetItem(ctx context.Context, id uint64) (*model.Item, error) {
	httpRequest, err := http.NewRequestWithContext(ctx, http.MethodGet, c.serverURL+"/api/item/"+strconv.FormatUint(id, 10), nil)
	if err != nil {
		return nil, err
	}
	response, err := c.client.Do(httpRequest)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = response.Body.Close()
	}()
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return nil, parseHTTPError(response)
	}
	var item model.Item
	err = json.NewDecoder(response.Body).Decode(&item)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (c *Client) PushItem(ctx context.Context, sourceID uint32, itemInfo datasource.ItemInfo) (*model.Item, error) {
	response, err := c.jsonRequest(ctx, http.MethodPost, c.serverURL+"/api/source/"+strconv.FormatUint(uint64(sourceID), 10)+"/push", itemInfo)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = response.Body.Close()
	}()
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return nil, parseHTTPError(response)
	}
	var item model.Item
	err = json.NewDecoder(response.Body).Decode(&item)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (c *Client) GetSourceChunks(ctx context.Context, sourceID uint32, request inspect.GetSourceChunksRequest) ([]model.Chunk, error) {
	response, err := c.jsonRequest(ctx, http.MethodGet, c.serverURL+"/api/source/"+strconv.FormatUint(uint64(sourceID), 10)+"/chunks", request)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = response.Body.Close()
	}()
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return nil, parseHTTPError(response)
	}
	var chunks []model.Chunk
	err = json.NewDecoder(response.Body).Decode(&chunks)
	if err != nil {
		return nil, err
	}
	return chunks, nil
}

func (c *Client) Chunk(ctx context.Context, sourceID uint32, request datasource.ChunkRequest) (*model.Chunk, error) {
	response, err := c.jsonRequest(ctx, http.MethodPost, c.serverURL+"/api/source/"+strconv.FormatUint(uint64(sourceID), 10)+"/chunk", request)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = response.Body.Close()
	}()
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return nil, parseHTTPError(response)
	}
	var chunk model.Chunk
	err = json.NewDecoder(response.Body).Decode(&chunk)
	if err != nil {
		return nil, err
	}
	return &chunk, nil
}

func (c *Client) Pack(ctx context.Context, chunkID uint64) ([]model.Car, error) {
	response, err := c.jsonRequest(ctx, http.MethodPost, c.serverURL+"/api/chunk/"+strconv.FormatUint(chunkID, 10)+"/pack", nil)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = response.Body.Close()
	}()
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return nil, parseHTTPError(response)
	}
	var cars []model.Car
	err = json.NewDecoder(response.Body).Decode(&cars)
	if err != nil {
		return nil, err
	}
	return cars, nil
}

type HTTPError struct {
	Err string `json:"err"`
}

func parseHTTPError(response *http.Response) error {
	bodyData, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	var httpError HTTPError
	jsonErr := json.Unmarshal(bodyData, &httpError)
	if jsonErr == nil {
		err = errors.New(httpError.Err) //nolint:goerr113
	} else {
		err = errors.New(string(bodyData)) //nolint:goerr113
	}

	switch response.StatusCode {
	case http.StatusBadRequest:
		return handler.InvalidParameterError{
			Err: err,
		}
	case http.StatusNotFound:
		return handler.NotFoundError{
			Err: err,
		}
	case http.StatusConflict:
		return handler.DuplicateRecordError{
			Err: err,
		}
	default:
		return err
	}
}
