// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/data-preservation-programs/singularity/client/swagger/models"
)

// NewCreateHdfsStorageParams creates a new CreateHdfsStorageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateHdfsStorageParams() *CreateHdfsStorageParams {
	return &CreateHdfsStorageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateHdfsStorageParamsWithTimeout creates a new CreateHdfsStorageParams object
// with the ability to set a timeout on a request.
func NewCreateHdfsStorageParamsWithTimeout(timeout time.Duration) *CreateHdfsStorageParams {
	return &CreateHdfsStorageParams{
		timeout: timeout,
	}
}

// NewCreateHdfsStorageParamsWithContext creates a new CreateHdfsStorageParams object
// with the ability to set a context for a request.
func NewCreateHdfsStorageParamsWithContext(ctx context.Context) *CreateHdfsStorageParams {
	return &CreateHdfsStorageParams{
		Context: ctx,
	}
}

// NewCreateHdfsStorageParamsWithHTTPClient creates a new CreateHdfsStorageParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateHdfsStorageParamsWithHTTPClient(client *http.Client) *CreateHdfsStorageParams {
	return &CreateHdfsStorageParams{
		HTTPClient: client,
	}
}

/*
CreateHdfsStorageParams contains all the parameters to send to the API endpoint

	for the create hdfs storage operation.

	Typically these are written to a http.Request.
*/
type CreateHdfsStorageParams struct {

	/* Request.

	   Request body
	*/
	Request *models.StorageCreateHdfsStorageRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create hdfs storage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateHdfsStorageParams) WithDefaults() *CreateHdfsStorageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create hdfs storage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateHdfsStorageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create hdfs storage params
func (o *CreateHdfsStorageParams) WithTimeout(timeout time.Duration) *CreateHdfsStorageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create hdfs storage params
func (o *CreateHdfsStorageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create hdfs storage params
func (o *CreateHdfsStorageParams) WithContext(ctx context.Context) *CreateHdfsStorageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create hdfs storage params
func (o *CreateHdfsStorageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create hdfs storage params
func (o *CreateHdfsStorageParams) WithHTTPClient(client *http.Client) *CreateHdfsStorageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create hdfs storage params
func (o *CreateHdfsStorageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the create hdfs storage params
func (o *CreateHdfsStorageParams) WithRequest(request *models.StorageCreateHdfsStorageRequest) *CreateHdfsStorageParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the create hdfs storage params
func (o *CreateHdfsStorageParams) SetRequest(request *models.StorageCreateHdfsStorageRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *CreateHdfsStorageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
