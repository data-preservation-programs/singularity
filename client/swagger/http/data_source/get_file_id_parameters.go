// Code generated by go-swagger; DO NOT EDIT.

package data_source

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

// NewGetFileIDParams creates a new GetFileIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetFileIDParams() *GetFileIDParams {
	return &GetFileIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetFileIDParamsWithTimeout creates a new GetFileIDParams object
// with the ability to set a timeout on a request.
func NewGetFileIDParamsWithTimeout(timeout time.Duration) *GetFileIDParams {
	return &GetFileIDParams{
		timeout: timeout,
	}
}

// NewGetFileIDParamsWithContext creates a new GetFileIDParams object
// with the ability to set a context for a request.
func NewGetFileIDParamsWithContext(ctx context.Context) *GetFileIDParams {
	return &GetFileIDParams{
		Context: ctx,
	}
}

// NewGetFileIDParamsWithHTTPClient creates a new GetFileIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetFileIDParamsWithHTTPClient(client *http.Client) *GetFileIDParams {
	return &GetFileIDParams{
		HTTPClient: client,
	}
}

/*
GetFileIDParams contains all the parameters to send to the API endpoint

	for the get file ID operation.

	Typically these are written to a http.Request.
*/
type GetFileIDParams struct {

	/* ID.

	   File ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get file ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFileIDParams) WithDefaults() *GetFileIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get file ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFileIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get file ID params
func (o *GetFileIDParams) WithTimeout(timeout time.Duration) *GetFileIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get file ID params
func (o *GetFileIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get file ID params
func (o *GetFileIDParams) WithContext(ctx context.Context) *GetFileIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get file ID params
func (o *GetFileIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get file ID params
func (o *GetFileIDParams) WithHTTPClient(client *http.Client) *GetFileIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get file ID params
func (o *GetFileIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get file ID params
func (o *GetFileIDParams) WithID(id string) *GetFileIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get file ID params
func (o *GetFileIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetFileIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}