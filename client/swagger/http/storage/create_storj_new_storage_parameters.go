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

// NewCreateStorjNewStorageParams creates a new CreateStorjNewStorageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateStorjNewStorageParams() *CreateStorjNewStorageParams {
	return &CreateStorjNewStorageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateStorjNewStorageParamsWithTimeout creates a new CreateStorjNewStorageParams object
// with the ability to set a timeout on a request.
func NewCreateStorjNewStorageParamsWithTimeout(timeout time.Duration) *CreateStorjNewStorageParams {
	return &CreateStorjNewStorageParams{
		timeout: timeout,
	}
}

// NewCreateStorjNewStorageParamsWithContext creates a new CreateStorjNewStorageParams object
// with the ability to set a context for a request.
func NewCreateStorjNewStorageParamsWithContext(ctx context.Context) *CreateStorjNewStorageParams {
	return &CreateStorjNewStorageParams{
		Context: ctx,
	}
}

// NewCreateStorjNewStorageParamsWithHTTPClient creates a new CreateStorjNewStorageParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateStorjNewStorageParamsWithHTTPClient(client *http.Client) *CreateStorjNewStorageParams {
	return &CreateStorjNewStorageParams{
		HTTPClient: client,
	}
}

/*
CreateStorjNewStorageParams contains all the parameters to send to the API endpoint

	for the create storj new storage operation.

	Typically these are written to a http.Request.
*/
type CreateStorjNewStorageParams struct {

	/* Request.

	   Request body
	*/
	Request *models.StorageCreateStorjNewStorageRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create storj new storage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateStorjNewStorageParams) WithDefaults() *CreateStorjNewStorageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create storj new storage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateStorjNewStorageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create storj new storage params
func (o *CreateStorjNewStorageParams) WithTimeout(timeout time.Duration) *CreateStorjNewStorageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create storj new storage params
func (o *CreateStorjNewStorageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create storj new storage params
func (o *CreateStorjNewStorageParams) WithContext(ctx context.Context) *CreateStorjNewStorageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create storj new storage params
func (o *CreateStorjNewStorageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create storj new storage params
func (o *CreateStorjNewStorageParams) WithHTTPClient(client *http.Client) *CreateStorjNewStorageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create storj new storage params
func (o *CreateStorjNewStorageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the create storj new storage params
func (o *CreateStorjNewStorageParams) WithRequest(request *models.StorageCreateStorjNewStorageRequest) *CreateStorjNewStorageParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the create storj new storage params
func (o *CreateStorjNewStorageParams) SetRequest(request *models.StorageCreateStorjNewStorageRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *CreateStorjNewStorageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
