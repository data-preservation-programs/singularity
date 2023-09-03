// Code generated by go-swagger; DO NOT EDIT.

package preparation

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
	"github.com/go-openapi/swag"
)

// NewDeletePreparationIDOutputNameParams creates a new DeletePreparationIDOutputNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeletePreparationIDOutputNameParams() *DeletePreparationIDOutputNameParams {
	return &DeletePreparationIDOutputNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePreparationIDOutputNameParamsWithTimeout creates a new DeletePreparationIDOutputNameParams object
// with the ability to set a timeout on a request.
func NewDeletePreparationIDOutputNameParamsWithTimeout(timeout time.Duration) *DeletePreparationIDOutputNameParams {
	return &DeletePreparationIDOutputNameParams{
		timeout: timeout,
	}
}

// NewDeletePreparationIDOutputNameParamsWithContext creates a new DeletePreparationIDOutputNameParams object
// with the ability to set a context for a request.
func NewDeletePreparationIDOutputNameParamsWithContext(ctx context.Context) *DeletePreparationIDOutputNameParams {
	return &DeletePreparationIDOutputNameParams{
		Context: ctx,
	}
}

// NewDeletePreparationIDOutputNameParamsWithHTTPClient creates a new DeletePreparationIDOutputNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeletePreparationIDOutputNameParamsWithHTTPClient(client *http.Client) *DeletePreparationIDOutputNameParams {
	return &DeletePreparationIDOutputNameParams{
		HTTPClient: client,
	}
}

/*
DeletePreparationIDOutputNameParams contains all the parameters to send to the API endpoint

	for the delete preparation ID output name operation.

	Typically these are written to a http.Request.
*/
type DeletePreparationIDOutputNameParams struct {

	/* ID.

	   Preparation ID or name
	*/
	ID int64

	/* Name.

	   Output storage ID or name
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete preparation ID output name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeletePreparationIDOutputNameParams) WithDefaults() *DeletePreparationIDOutputNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete preparation ID output name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeletePreparationIDOutputNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete preparation ID output name params
func (o *DeletePreparationIDOutputNameParams) WithTimeout(timeout time.Duration) *DeletePreparationIDOutputNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete preparation ID output name params
func (o *DeletePreparationIDOutputNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete preparation ID output name params
func (o *DeletePreparationIDOutputNameParams) WithContext(ctx context.Context) *DeletePreparationIDOutputNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete preparation ID output name params
func (o *DeletePreparationIDOutputNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete preparation ID output name params
func (o *DeletePreparationIDOutputNameParams) WithHTTPClient(client *http.Client) *DeletePreparationIDOutputNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete preparation ID output name params
func (o *DeletePreparationIDOutputNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete preparation ID output name params
func (o *DeletePreparationIDOutputNameParams) WithID(id int64) *DeletePreparationIDOutputNameParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete preparation ID output name params
func (o *DeletePreparationIDOutputNameParams) SetID(id int64) {
	o.ID = id
}

// WithName adds the name to the delete preparation ID output name params
func (o *DeletePreparationIDOutputNameParams) WithName(name string) *DeletePreparationIDOutputNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete preparation ID output name params
func (o *DeletePreparationIDOutputNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePreparationIDOutputNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
