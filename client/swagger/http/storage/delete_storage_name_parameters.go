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
)

// NewDeleteStorageNameParams creates a new DeleteStorageNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteStorageNameParams() *DeleteStorageNameParams {
	return &DeleteStorageNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteStorageNameParamsWithTimeout creates a new DeleteStorageNameParams object
// with the ability to set a timeout on a request.
func NewDeleteStorageNameParamsWithTimeout(timeout time.Duration) *DeleteStorageNameParams {
	return &DeleteStorageNameParams{
		timeout: timeout,
	}
}

// NewDeleteStorageNameParamsWithContext creates a new DeleteStorageNameParams object
// with the ability to set a context for a request.
func NewDeleteStorageNameParamsWithContext(ctx context.Context) *DeleteStorageNameParams {
	return &DeleteStorageNameParams{
		Context: ctx,
	}
}

// NewDeleteStorageNameParamsWithHTTPClient creates a new DeleteStorageNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteStorageNameParamsWithHTTPClient(client *http.Client) *DeleteStorageNameParams {
	return &DeleteStorageNameParams{
		HTTPClient: client,
	}
}

/*
DeleteStorageNameParams contains all the parameters to send to the API endpoint

	for the delete storage name operation.

	Typically these are written to a http.Request.
*/
type DeleteStorageNameParams struct {

	/* Name.

	   Name
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete storage name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteStorageNameParams) WithDefaults() *DeleteStorageNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete storage name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteStorageNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete storage name params
func (o *DeleteStorageNameParams) WithTimeout(timeout time.Duration) *DeleteStorageNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete storage name params
func (o *DeleteStorageNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete storage name params
func (o *DeleteStorageNameParams) WithContext(ctx context.Context) *DeleteStorageNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete storage name params
func (o *DeleteStorageNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete storage name params
func (o *DeleteStorageNameParams) WithHTTPClient(client *http.Client) *DeleteStorageNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete storage name params
func (o *DeleteStorageNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the delete storage name params
func (o *DeleteStorageNameParams) WithName(name string) *DeleteStorageNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete storage name params
func (o *DeleteStorageNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteStorageNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
