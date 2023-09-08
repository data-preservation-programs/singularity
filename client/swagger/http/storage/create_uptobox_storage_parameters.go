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

// NewCreateUptoboxStorageParams creates a new CreateUptoboxStorageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateUptoboxStorageParams() *CreateUptoboxStorageParams {
	return &CreateUptoboxStorageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateUptoboxStorageParamsWithTimeout creates a new CreateUptoboxStorageParams object
// with the ability to set a timeout on a request.
func NewCreateUptoboxStorageParamsWithTimeout(timeout time.Duration) *CreateUptoboxStorageParams {
	return &CreateUptoboxStorageParams{
		timeout: timeout,
	}
}

// NewCreateUptoboxStorageParamsWithContext creates a new CreateUptoboxStorageParams object
// with the ability to set a context for a request.
func NewCreateUptoboxStorageParamsWithContext(ctx context.Context) *CreateUptoboxStorageParams {
	return &CreateUptoboxStorageParams{
		Context: ctx,
	}
}

// NewCreateUptoboxStorageParamsWithHTTPClient creates a new CreateUptoboxStorageParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateUptoboxStorageParamsWithHTTPClient(client *http.Client) *CreateUptoboxStorageParams {
	return &CreateUptoboxStorageParams{
		HTTPClient: client,
	}
}

/*
CreateUptoboxStorageParams contains all the parameters to send to the API endpoint

	for the create uptobox storage operation.

	Typically these are written to a http.Request.
*/
type CreateUptoboxStorageParams struct {

	/* Request.

	   Request body
	*/
	Request *models.StorageCreateUptoboxStorageRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create uptobox storage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateUptoboxStorageParams) WithDefaults() *CreateUptoboxStorageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create uptobox storage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateUptoboxStorageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create uptobox storage params
func (o *CreateUptoboxStorageParams) WithTimeout(timeout time.Duration) *CreateUptoboxStorageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create uptobox storage params
func (o *CreateUptoboxStorageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create uptobox storage params
func (o *CreateUptoboxStorageParams) WithContext(ctx context.Context) *CreateUptoboxStorageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create uptobox storage params
func (o *CreateUptoboxStorageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create uptobox storage params
func (o *CreateUptoboxStorageParams) WithHTTPClient(client *http.Client) *CreateUptoboxStorageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create uptobox storage params
func (o *CreateUptoboxStorageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the create uptobox storage params
func (o *CreateUptoboxStorageParams) WithRequest(request *models.StorageCreateUptoboxStorageRequest) *CreateUptoboxStorageParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the create uptobox storage params
func (o *CreateUptoboxStorageParams) SetRequest(request *models.StorageCreateUptoboxStorageRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *CreateUptoboxStorageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
