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

// NewCreateS3NeteaseStorageParams creates a new CreateS3NeteaseStorageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateS3NeteaseStorageParams() *CreateS3NeteaseStorageParams {
	return &CreateS3NeteaseStorageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateS3NeteaseStorageParamsWithTimeout creates a new CreateS3NeteaseStorageParams object
// with the ability to set a timeout on a request.
func NewCreateS3NeteaseStorageParamsWithTimeout(timeout time.Duration) *CreateS3NeteaseStorageParams {
	return &CreateS3NeteaseStorageParams{
		timeout: timeout,
	}
}

// NewCreateS3NeteaseStorageParamsWithContext creates a new CreateS3NeteaseStorageParams object
// with the ability to set a context for a request.
func NewCreateS3NeteaseStorageParamsWithContext(ctx context.Context) *CreateS3NeteaseStorageParams {
	return &CreateS3NeteaseStorageParams{
		Context: ctx,
	}
}

// NewCreateS3NeteaseStorageParamsWithHTTPClient creates a new CreateS3NeteaseStorageParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateS3NeteaseStorageParamsWithHTTPClient(client *http.Client) *CreateS3NeteaseStorageParams {
	return &CreateS3NeteaseStorageParams{
		HTTPClient: client,
	}
}

/*
CreateS3NeteaseStorageParams contains all the parameters to send to the API endpoint

	for the create s3 netease storage operation.

	Typically these are written to a http.Request.
*/
type CreateS3NeteaseStorageParams struct {

	/* Request.

	   Request body
	*/
	Request *models.StorageCreateS3NeteaseStorageRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create s3 netease storage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateS3NeteaseStorageParams) WithDefaults() *CreateS3NeteaseStorageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create s3 netease storage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateS3NeteaseStorageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create s3 netease storage params
func (o *CreateS3NeteaseStorageParams) WithTimeout(timeout time.Duration) *CreateS3NeteaseStorageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create s3 netease storage params
func (o *CreateS3NeteaseStorageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create s3 netease storage params
func (o *CreateS3NeteaseStorageParams) WithContext(ctx context.Context) *CreateS3NeteaseStorageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create s3 netease storage params
func (o *CreateS3NeteaseStorageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create s3 netease storage params
func (o *CreateS3NeteaseStorageParams) WithHTTPClient(client *http.Client) *CreateS3NeteaseStorageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create s3 netease storage params
func (o *CreateS3NeteaseStorageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the create s3 netease storage params
func (o *CreateS3NeteaseStorageParams) WithRequest(request *models.StorageCreateS3NeteaseStorageRequest) *CreateS3NeteaseStorageParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the create s3 netease storage params
func (o *CreateS3NeteaseStorageParams) SetRequest(request *models.StorageCreateS3NeteaseStorageRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *CreateS3NeteaseStorageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
