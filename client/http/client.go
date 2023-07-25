package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/handler/dataset"
	"github.com/data-preservation-programs/singularity/handler/datasource"
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

func (c *Client) CreateDataset(ctx context.Context, request dataset.CreateRequest) (*model.Dataset, error) {
	jsonParams, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	httpRequest, err := http.NewRequestWithContext(ctx, "POST", c.serverURL+"/api/dataset", bytes.NewReader(jsonParams))
	if err != nil {
		return nil, err
	}
	response, err := c.client.Do(httpRequest)
	if err != nil {
		return nil, err
	}
	if response.StatusCode < 200 && response.StatusCode >= 300 {
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
	jsonParams, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	httpRequest, err := http.NewRequestWithContext(ctx, "POST", c.serverURL+"/api/source/local/dataset/"+datasetName, bytes.NewReader(jsonParams))
	if err != nil {
		return nil, err
	}
	response, err := c.client.Do(httpRequest)
	if err != nil {
		return nil, err
	}
	if response.StatusCode < 200 && response.StatusCode >= 300 {
		return nil, parseHTTPError(response)
	}
	var source model.Source
	err = json.NewDecoder(response.Body).Decode(&source)
	if err != nil {
		return nil, err
	}
	return &source, nil
}

func (c *Client) GetItem(ctx context.Context, id uint64) (*model.Item, error) {

	httpRequest, err := http.NewRequestWithContext(ctx, "GET", c.serverURL+"/api/item/"+strconv.FormatUint(id, 10), nil)
	if err != nil {
		return nil, err
	}
	response, err := c.client.Do(httpRequest)
	if err != nil {
		return nil, err
	}
	if response.StatusCode < 200 && response.StatusCode >= 300 {
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
	jsonParams, err := json.Marshal(itemInfo)
	if err != nil {
		return nil, err
	}
	httpRequest, err := http.NewRequestWithContext(ctx, "POST", c.serverURL+"/api/source/"+strconv.FormatUint(uint64(sourceID), 10)+"/push", bytes.NewReader(jsonParams))
	if err != nil {
		return nil, err
	}
	response, err := c.client.Do(httpRequest)
	if err != nil {
		return nil, err
	}
	if response.StatusCode < 200 && response.StatusCode >= 300 {
		return nil, parseHTTPError(response)
	}
	var item model.Item
	err = json.NewDecoder(response.Body).Decode(&item)
	if err != nil {
		return nil, err
	}
	return &item, nil
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
		err = errors.New(httpError.Err)
	} else {
		err = errors.New(string(bodyData))
	}

	if response.StatusCode == http.StatusBadRequest {
		return handler.ErrInvalidParameter{
			Err: err,
		}
	}

	if response.StatusCode == http.StatusBadRequest {
		return handler.ErrNotFound{
			Err: err,
		}
	}

	if response.StatusCode == http.StatusConflict {
		return handler.ErrDuplicateRecord{
			Err: err,
		}
	}

	return err
}
