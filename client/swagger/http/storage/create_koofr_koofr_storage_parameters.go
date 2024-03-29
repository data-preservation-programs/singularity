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

// NewCreateKoofrKoofrStorageParams creates a new CreateKoofrKoofrStorageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateKoofrKoofrStorageParams() *CreateKoofrKoofrStorageParams {
	return &CreateKoofrKoofrStorageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateKoofrKoofrStorageParamsWithTimeout creates a new CreateKoofrKoofrStorageParams object
// with the ability to set a timeout on a request.
func NewCreateKoofrKoofrStorageParamsWithTimeout(timeout time.Duration) *CreateKoofrKoofrStorageParams {
	return &CreateKoofrKoofrStorageParams{
		timeout: timeout,
	}
}

// NewCreateKoofrKoofrStorageParamsWithContext creates a new CreateKoofrKoofrStorageParams object
// with the ability to set a context for a request.
func NewCreateKoofrKoofrStorageParamsWithContext(ctx context.Context) *CreateKoofrKoofrStorageParams {
	return &CreateKoofrKoofrStorageParams{
		Context: ctx,
	}
}

// NewCreateKoofrKoofrStorageParamsWithHTTPClient creates a new CreateKoofrKoofrStorageParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateKoofrKoofrStorageParamsWithHTTPClient(client *http.Client) *CreateKoofrKoofrStorageParams {
	return &CreateKoofrKoofrStorageParams{
		HTTPClient: client,
	}
}

/*
CreateKoofrKoofrStorageParams contains all the parameters to send to the API endpoint

	for the create koofr koofr storage operation.

	Typically these are written to a http.Request.
*/
type CreateKoofrKoofrStorageParams struct {

	/* Request.

	   Request body
	*/
	Request *models.StorageCreateKoofrKoofrStorageRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create koofr koofr storage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateKoofrKoofrStorageParams) WithDefaults() *CreateKoofrKoofrStorageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create koofr koofr storage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateKoofrKoofrStorageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create koofr koofr storage params
func (o *CreateKoofrKoofrStorageParams) WithTimeout(timeout time.Duration) *CreateKoofrKoofrStorageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create koofr koofr storage params
func (o *CreateKoofrKoofrStorageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create koofr koofr storage params
func (o *CreateKoofrKoofrStorageParams) WithContext(ctx context.Context) *CreateKoofrKoofrStorageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create koofr koofr storage params
func (o *CreateKoofrKoofrStorageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create koofr koofr storage params
func (o *CreateKoofrKoofrStorageParams) WithHTTPClient(client *http.Client) *CreateKoofrKoofrStorageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create koofr koofr storage params
func (o *CreateKoofrKoofrStorageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the create koofr koofr storage params
func (o *CreateKoofrKoofrStorageParams) WithRequest(request *models.StorageCreateKoofrKoofrStorageRequest) *CreateKoofrKoofrStorageParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the create koofr koofr storage params
func (o *CreateKoofrKoofrStorageParams) SetRequest(request *models.StorageCreateKoofrKoofrStorageRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *CreateKoofrKoofrStorageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
