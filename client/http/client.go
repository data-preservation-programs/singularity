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

func (c *Client) GetFile(ctx context.Context, id uint64) (*model.File, error) {
	httpRequest, err := http.NewRequestWithContext(ctx, http.MethodGet, c.serverURL+"/api/file/"+strconv.FormatUint(id, 10), nil)
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
	var file model.File
	err = json.NewDecoder(response.Body).Decode(&file)
	if err != nil {
		return nil, err
	}
	return &file, nil
}

func (c *Client) PushFile(ctx context.Context, sourceID uint32, fileInfo datasource.FileInfo) (*model.File, error) {
	response, err := c.jsonRequest(ctx, http.MethodPost, c.serverURL+"/api/source/"+strconv.FormatUint(uint64(sourceID), 10)+"/push", fileInfo)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = response.Body.Close()
	}()
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return nil, parseHTTPError(response)
	}
	var file model.File
	err = json.NewDecoder(response.Body).Decode(&file)
	if err != nil {
		return nil, err
	}
	return &file, nil
}

func (c *Client) GetSourcePackJobs(ctx context.Context, sourceID uint32, request inspect.GetSourcePackJobsRequest) ([]model.PackJob, error) {
	response, err := c.jsonRequest(ctx, http.MethodGet, c.serverURL+"/api/source/"+strconv.FormatUint(uint64(sourceID), 10)+"/packjobs", request)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = response.Body.Close()
	}()
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return nil, parseHTTPError(response)
	}
	var packJobs []model.PackJob
	err = json.NewDecoder(response.Body).Decode(&packJobs)
	if err != nil {
		return nil, err
	}
	return packJobs, nil
}

func (c *Client) PrepareToPackFile(ctx context.Context, fileID uint64) (int64, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodPost, c.serverURL+"/api/file/"+strconv.FormatUint(fileID, 10)+"/prepare_to_pack", nil)
	if err != nil {
		return 0, err
	}
	response, err := c.client.Do(request)
	if err != nil {
		return 0, err
	}
	defer func() {
		_ = response.Body.Close()
	}()
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return 0, parseHTTPError(response)
	}
	var incomplete int64
	err = json.NewDecoder(response.Body).Decode(&incomplete)
	if err != nil {
		return 0, err
	}
	return incomplete, nil
}

func (c *Client) PrepareToPackSource(ctx context.Context, sourceID uint32) error {
	request, err := http.NewRequestWithContext(ctx, http.MethodPost, c.serverURL+"/api/source/"+strconv.FormatUint(uint64(sourceID), 10)+"/prepare_to_pack", nil)
	if err != nil {
		return err
	}
	response, err := c.client.Do(request)
	if err != nil {
		return err
	}
	defer func() {
		_ = response.Body.Close()
	}()
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return parseHTTPError(response)
	}
	return nil
}

func (c *Client) Pack(ctx context.Context, packJobID uint32) ([]model.Car, error) {
	response, err := c.jsonRequest(ctx, http.MethodPost, c.serverURL+"/api/packjob/"+strconv.FormatUint(uint64(packJobID), 10)+"/pack", nil)
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

func (c *Client) GetFileDeals(ctx context.Context, id uint64) ([]model.Deal, error) {
	httpRequest, err := http.NewRequestWithContext(ctx, http.MethodGet, c.serverURL+"/api/file/"+strconv.FormatUint(id, 10)+"/deals", nil)
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
	var deals []model.Deal
	err = json.NewDecoder(response.Body).Decode(&deals)
	if err != nil {
		return nil, err
	}
	return deals, nil
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
