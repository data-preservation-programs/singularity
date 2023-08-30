//go:build exclude

package client

import (
	"context"
	"fmt"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/model"
)

type Client interface {
	CreateDataset(ctx context.Context, request dataset.CreateRequest) (*model.Preparation, error)
	CreateLocalSource(ctx context.Context, datasetName string, params datasource.LocalRequest) (*model.Source, error)
	ListSourcesByDataset(ctx context.Context, datasetName string) ([]model.Source, error)
	GetFile(ctx context.Context, id uint64) (*model.File, error)
	PushFile(ctx context.Context, sourceID uint32, fileInfo datasource.FileInfo) (*model.File, error)
	GetSourcePackJobs(ctx context.Context, sourceID uint32, request inspect.GetSourcePackJobsRequest) ([]model.PackJob, error)
	CreatePackJob(ctx context.Context, sourceID uint32, request datasource.CreatePackJobRequest) (*model.PackJob, error)
}

type InvalidParameterError struct {
	Err error
}

func (e InvalidParameterError) Unwrap() error {
	return e.Err
}
func (e InvalidParameterError) Error() string {
	return fmt.Sprintf("invalid parameter: %s", e.Err.Error())
}

func NewInvalidParameterErr(err string) InvalidParameterError {
	return InvalidParameterError{
		Err: errors.New(err),
	}
}

type NotFoundError struct {
	Err error
}

func (e NotFoundError) Unwrap() error {
	return e.Err
}
func (e NotFoundError) Error() string {
	return fmt.Sprintf("not found: %s", e.Err.Error())
}

type DuplicateRecordError struct {
	Err error
}

func (e DuplicateRecordError) Unwrap() error {
	return e.Err
}

func (e DuplicateRecordError) Error() string {
	return fmt.Sprintf("duplicate record: %s", e.Err.Error())
}

func NewDuplicateRecordError(err string) DuplicateRecordError {
	return DuplicateRecordError{
		Err: errors.New(err),
	}
}
