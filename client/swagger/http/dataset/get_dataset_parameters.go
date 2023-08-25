// Code generated by go-swagger; DO NOT EDIT.

package dataset

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
)

// NewGetDatasetParams creates a new GetDatasetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDatasetParams() *GetDatasetParams {
	return &GetDatasetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDatasetParamsWithTimeout creates a new GetDatasetParams object
// with the ability to set a timeout on a request.
func NewGetDatasetParamsWithTimeout(timeout time.Duration) *GetDatasetParams {
	return &GetDatasetParams{
		timeout: timeout,
	}
}

// NewGetDatasetParamsWithContext creates a new GetDatasetParams object
// with the ability to set a context for a request.
func NewGetDatasetParamsWithContext(ctx context.Context) *GetDatasetParams {
	return &GetDatasetParams{
		Context: ctx,
	}
}

// NewGetDatasetParamsWithHTTPClient creates a new GetDatasetParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDatasetParamsWithHTTPClient(client *http.Client) *GetDatasetParams {
	return &GetDatasetParams{
		HTTPClient: client,
	}
}

/*
GetDatasetParams contains all the parameters to send to the API endpoint

	for the get dataset operation.

	Typically these are written to a http.Request.
*/
type GetDatasetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get dataset params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDatasetParams) WithDefaults() *GetDatasetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get dataset params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDatasetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get dataset params
func (o *GetDatasetParams) WithTimeout(timeout time.Duration) *GetDatasetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get dataset params
func (o *GetDatasetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get dataset params
func (o *GetDatasetParams) WithContext(ctx context.Context) *GetDatasetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get dataset params
func (o *GetDatasetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get dataset params
func (o *GetDatasetParams) WithHTTPClient(client *http.Client) *GetDatasetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get dataset params
func (o *GetDatasetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetDatasetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
