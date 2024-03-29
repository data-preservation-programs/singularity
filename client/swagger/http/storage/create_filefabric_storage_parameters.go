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

// NewCreateFilefabricStorageParams creates a new CreateFilefabricStorageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateFilefabricStorageParams() *CreateFilefabricStorageParams {
	return &CreateFilefabricStorageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateFilefabricStorageParamsWithTimeout creates a new CreateFilefabricStorageParams object
// with the ability to set a timeout on a request.
func NewCreateFilefabricStorageParamsWithTimeout(timeout time.Duration) *CreateFilefabricStorageParams {
	return &CreateFilefabricStorageParams{
		timeout: timeout,
	}
}

// NewCreateFilefabricStorageParamsWithContext creates a new CreateFilefabricStorageParams object
// with the ability to set a context for a request.
func NewCreateFilefabricStorageParamsWithContext(ctx context.Context) *CreateFilefabricStorageParams {
	return &CreateFilefabricStorageParams{
		Context: ctx,
	}
}

// NewCreateFilefabricStorageParamsWithHTTPClient creates a new CreateFilefabricStorageParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateFilefabricStorageParamsWithHTTPClient(client *http.Client) *CreateFilefabricStorageParams {
	return &CreateFilefabricStorageParams{
		HTTPClient: client,
	}
}

/*
CreateFilefabricStorageParams contains all the parameters to send to the API endpoint

	for the create filefabric storage operation.

	Typically these are written to a http.Request.
*/
type CreateFilefabricStorageParams struct {

	/* Request.

	   Request body
	*/
	Request *models.StorageCreateFilefabricStorageRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create filefabric storage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateFilefabricStorageParams) WithDefaults() *CreateFilefabricStorageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create filefabric storage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateFilefabricStorageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create filefabric storage params
func (o *CreateFilefabricStorageParams) WithTimeout(timeout time.Duration) *CreateFilefabricStorageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create filefabric storage params
func (o *CreateFilefabricStorageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create filefabric storage params
func (o *CreateFilefabricStorageParams) WithContext(ctx context.Context) *CreateFilefabricStorageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create filefabric storage params
func (o *CreateFilefabricStorageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create filefabric storage params
func (o *CreateFilefabricStorageParams) WithHTTPClient(client *http.Client) *CreateFilefabricStorageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create filefabric storage params
func (o *CreateFilefabricStorageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the create filefabric storage params
func (o *CreateFilefabricStorageParams) WithRequest(request *models.StorageCreateFilefabricStorageRequest) *CreateFilefabricStorageParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the create filefabric storage params
func (o *CreateFilefabricStorageParams) SetRequest(request *models.StorageCreateFilefabricStorageRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *CreateFilefabricStorageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
