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

// NewCreateJottacloudStorageParams creates a new CreateJottacloudStorageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateJottacloudStorageParams() *CreateJottacloudStorageParams {
	return &CreateJottacloudStorageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateJottacloudStorageParamsWithTimeout creates a new CreateJottacloudStorageParams object
// with the ability to set a timeout on a request.
func NewCreateJottacloudStorageParamsWithTimeout(timeout time.Duration) *CreateJottacloudStorageParams {
	return &CreateJottacloudStorageParams{
		timeout: timeout,
	}
}

// NewCreateJottacloudStorageParamsWithContext creates a new CreateJottacloudStorageParams object
// with the ability to set a context for a request.
func NewCreateJottacloudStorageParamsWithContext(ctx context.Context) *CreateJottacloudStorageParams {
	return &CreateJottacloudStorageParams{
		Context: ctx,
	}
}

// NewCreateJottacloudStorageParamsWithHTTPClient creates a new CreateJottacloudStorageParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateJottacloudStorageParamsWithHTTPClient(client *http.Client) *CreateJottacloudStorageParams {
	return &CreateJottacloudStorageParams{
		HTTPClient: client,
	}
}

/*
CreateJottacloudStorageParams contains all the parameters to send to the API endpoint

	for the create jottacloud storage operation.

	Typically these are written to a http.Request.
*/
type CreateJottacloudStorageParams struct {

	/* Request.

	   Request body
	*/
	Request *models.StorageCreateJottacloudStorageRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create jottacloud storage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateJottacloudStorageParams) WithDefaults() *CreateJottacloudStorageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create jottacloud storage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateJottacloudStorageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create jottacloud storage params
func (o *CreateJottacloudStorageParams) WithTimeout(timeout time.Duration) *CreateJottacloudStorageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create jottacloud storage params
func (o *CreateJottacloudStorageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create jottacloud storage params
func (o *CreateJottacloudStorageParams) WithContext(ctx context.Context) *CreateJottacloudStorageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create jottacloud storage params
func (o *CreateJottacloudStorageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create jottacloud storage params
func (o *CreateJottacloudStorageParams) WithHTTPClient(client *http.Client) *CreateJottacloudStorageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create jottacloud storage params
func (o *CreateJottacloudStorageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the create jottacloud storage params
func (o *CreateJottacloudStorageParams) WithRequest(request *models.StorageCreateJottacloudStorageRequest) *CreateJottacloudStorageParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the create jottacloud storage params
func (o *CreateJottacloudStorageParams) SetRequest(request *models.StorageCreateJottacloudStorageRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *CreateJottacloudStorageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
