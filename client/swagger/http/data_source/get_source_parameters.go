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

// NewGetSourceParams creates a new GetSourceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSourceParams() *GetSourceParams {
	return &GetSourceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSourceParamsWithTimeout creates a new GetSourceParams object
// with the ability to set a timeout on a request.
func NewGetSourceParamsWithTimeout(timeout time.Duration) *GetSourceParams {
	return &GetSourceParams{
		timeout: timeout,
	}
}

// NewGetSourceParamsWithContext creates a new GetSourceParams object
// with the ability to set a context for a request.
func NewGetSourceParamsWithContext(ctx context.Context) *GetSourceParams {
	return &GetSourceParams{
		Context: ctx,
	}
}

// NewGetSourceParamsWithHTTPClient creates a new GetSourceParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSourceParamsWithHTTPClient(client *http.Client) *GetSourceParams {
	return &GetSourceParams{
		HTTPClient: client,
	}
}

/*
GetSourceParams contains all the parameters to send to the API endpoint

	for the get source operation.

	Typically these are written to a http.Request.
*/
type GetSourceParams struct {

	/* Dataset.

	   Dataset name
	*/
	Dataset *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get source params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSourceParams) WithDefaults() *GetSourceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get source params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSourceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get source params
func (o *GetSourceParams) WithTimeout(timeout time.Duration) *GetSourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get source params
func (o *GetSourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get source params
func (o *GetSourceParams) WithContext(ctx context.Context) *GetSourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get source params
func (o *GetSourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get source params
func (o *GetSourceParams) WithHTTPClient(client *http.Client) *GetSourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get source params
func (o *GetSourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDataset adds the dataset to the get source params
func (o *GetSourceParams) WithDataset(dataset *string) *GetSourceParams {
	o.SetDataset(dataset)
	return o
}

// SetDataset adds the dataset to the get source params
func (o *GetSourceParams) SetDataset(dataset *string) {
	o.Dataset = dataset
}

// WriteToRequest writes these params to a swagger request
func (o *GetSourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Dataset != nil {

		// query param dataset
		var qrDataset string

		if o.Dataset != nil {
			qrDataset = *o.Dataset
		}
		qDataset := qrDataset
		if qDataset != "" {

			if err := r.SetQueryParam("dataset", qDataset); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
