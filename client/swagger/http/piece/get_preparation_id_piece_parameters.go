// Code generated by go-swagger; DO NOT EDIT.

package piece

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

// NewGetPreparationIDPieceParams creates a new GetPreparationIDPieceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPreparationIDPieceParams() *GetPreparationIDPieceParams {
	return &GetPreparationIDPieceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPreparationIDPieceParamsWithTimeout creates a new GetPreparationIDPieceParams object
// with the ability to set a timeout on a request.
func NewGetPreparationIDPieceParamsWithTimeout(timeout time.Duration) *GetPreparationIDPieceParams {
	return &GetPreparationIDPieceParams{
		timeout: timeout,
	}
}

// NewGetPreparationIDPieceParamsWithContext creates a new GetPreparationIDPieceParams object
// with the ability to set a context for a request.
func NewGetPreparationIDPieceParamsWithContext(ctx context.Context) *GetPreparationIDPieceParams {
	return &GetPreparationIDPieceParams{
		Context: ctx,
	}
}

// NewGetPreparationIDPieceParamsWithHTTPClient creates a new GetPreparationIDPieceParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPreparationIDPieceParamsWithHTTPClient(client *http.Client) *GetPreparationIDPieceParams {
	return &GetPreparationIDPieceParams{
		HTTPClient: client,
	}
}

/*
GetPreparationIDPieceParams contains all the parameters to send to the API endpoint

	for the get preparation ID piece operation.

	Typically these are written to a http.Request.
*/
type GetPreparationIDPieceParams struct {

	/* ID.

	   Preparation ID
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get preparation ID piece params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPreparationIDPieceParams) WithDefaults() *GetPreparationIDPieceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get preparation ID piece params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPreparationIDPieceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get preparation ID piece params
func (o *GetPreparationIDPieceParams) WithTimeout(timeout time.Duration) *GetPreparationIDPieceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get preparation ID piece params
func (o *GetPreparationIDPieceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get preparation ID piece params
func (o *GetPreparationIDPieceParams) WithContext(ctx context.Context) *GetPreparationIDPieceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get preparation ID piece params
func (o *GetPreparationIDPieceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get preparation ID piece params
func (o *GetPreparationIDPieceParams) WithHTTPClient(client *http.Client) *GetPreparationIDPieceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get preparation ID piece params
func (o *GetPreparationIDPieceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get preparation ID piece params
func (o *GetPreparationIDPieceParams) WithID(id int64) *GetPreparationIDPieceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get preparation ID piece params
func (o *GetPreparationIDPieceParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetPreparationIDPieceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
