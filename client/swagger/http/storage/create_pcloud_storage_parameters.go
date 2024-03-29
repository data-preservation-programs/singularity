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

// NewCreatePcloudStorageParams creates a new CreatePcloudStorageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreatePcloudStorageParams() *CreatePcloudStorageParams {
	return &CreatePcloudStorageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreatePcloudStorageParamsWithTimeout creates a new CreatePcloudStorageParams object
// with the ability to set a timeout on a request.
func NewCreatePcloudStorageParamsWithTimeout(timeout time.Duration) *CreatePcloudStorageParams {
	return &CreatePcloudStorageParams{
		timeout: timeout,
	}
}

// NewCreatePcloudStorageParamsWithContext creates a new CreatePcloudStorageParams object
// with the ability to set a context for a request.
func NewCreatePcloudStorageParamsWithContext(ctx context.Context) *CreatePcloudStorageParams {
	return &CreatePcloudStorageParams{
		Context: ctx,
	}
}

// NewCreatePcloudStorageParamsWithHTTPClient creates a new CreatePcloudStorageParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreatePcloudStorageParamsWithHTTPClient(client *http.Client) *CreatePcloudStorageParams {
	return &CreatePcloudStorageParams{
		HTTPClient: client,
	}
}

/*
CreatePcloudStorageParams contains all the parameters to send to the API endpoint

	for the create pcloud storage operation.

	Typically these are written to a http.Request.
*/
type CreatePcloudStorageParams struct {

	/* Request.

	   Request body
	*/
	Request *models.StorageCreatePcloudStorageRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create pcloud storage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePcloudStorageParams) WithDefaults() *CreatePcloudStorageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create pcloud storage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePcloudStorageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create pcloud storage params
func (o *CreatePcloudStorageParams) WithTimeout(timeout time.Duration) *CreatePcloudStorageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create pcloud storage params
func (o *CreatePcloudStorageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create pcloud storage params
func (o *CreatePcloudStorageParams) WithContext(ctx context.Context) *CreatePcloudStorageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create pcloud storage params
func (o *CreatePcloudStorageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create pcloud storage params
func (o *CreatePcloudStorageParams) WithHTTPClient(client *http.Client) *CreatePcloudStorageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create pcloud storage params
func (o *CreatePcloudStorageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the create pcloud storage params
func (o *CreatePcloudStorageParams) WithRequest(request *models.StorageCreatePcloudStorageRequest) *CreatePcloudStorageParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the create pcloud storage params
func (o *CreatePcloudStorageParams) SetRequest(request *models.StorageCreatePcloudStorageRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *CreatePcloudStorageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
