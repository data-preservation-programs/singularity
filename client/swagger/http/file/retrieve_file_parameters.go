// Code generated by go-swagger; DO NOT EDIT.

package file

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

// NewRetrieveFileParams creates a new RetrieveFileParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRetrieveFileParams() *RetrieveFileParams {
	return &RetrieveFileParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRetrieveFileParamsWithTimeout creates a new RetrieveFileParams object
// with the ability to set a timeout on a request.
func NewRetrieveFileParamsWithTimeout(timeout time.Duration) *RetrieveFileParams {
	return &RetrieveFileParams{
		timeout: timeout,
	}
}

// NewRetrieveFileParamsWithContext creates a new RetrieveFileParams object
// with the ability to set a context for a request.
func NewRetrieveFileParamsWithContext(ctx context.Context) *RetrieveFileParams {
	return &RetrieveFileParams{
		Context: ctx,
	}
}

// NewRetrieveFileParamsWithHTTPClient creates a new RetrieveFileParams object
// with the ability to set a custom HTTPClient for a request.
func NewRetrieveFileParamsWithHTTPClient(client *http.Client) *RetrieveFileParams {
	return &RetrieveFileParams{
		HTTPClient: client,
	}
}

/*
RetrieveFileParams contains all the parameters to send to the API endpoint

	for the retrieve file operation.

	Typically these are written to a http.Request.
*/
type RetrieveFileParams struct {

	/* Range.

	   HTTP Range Header
	*/
	Range *string

	/* ID.

	   File ID
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the retrieve file params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RetrieveFileParams) WithDefaults() *RetrieveFileParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the retrieve file params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RetrieveFileParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the retrieve file params
func (o *RetrieveFileParams) WithTimeout(timeout time.Duration) *RetrieveFileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the retrieve file params
func (o *RetrieveFileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the retrieve file params
func (o *RetrieveFileParams) WithContext(ctx context.Context) *RetrieveFileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the retrieve file params
func (o *RetrieveFileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the retrieve file params
func (o *RetrieveFileParams) WithHTTPClient(client *http.Client) *RetrieveFileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the retrieve file params
func (o *RetrieveFileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRange adds the rangeVar to the retrieve file params
func (o *RetrieveFileParams) WithRange(rangeVar *string) *RetrieveFileParams {
	o.SetRange(rangeVar)
	return o
}

// SetRange adds the range to the retrieve file params
func (o *RetrieveFileParams) SetRange(rangeVar *string) {
	o.Range = rangeVar
}

// WithID adds the id to the retrieve file params
func (o *RetrieveFileParams) WithID(id int64) *RetrieveFileParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the retrieve file params
func (o *RetrieveFileParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *RetrieveFileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Range != nil {

		// header param Range
		if err := r.SetHeaderParam("Range", *o.Range); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
