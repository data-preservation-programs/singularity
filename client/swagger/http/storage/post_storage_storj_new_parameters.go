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

// NewPostStorageStorjNewParams creates a new PostStorageStorjNewParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostStorageStorjNewParams() *PostStorageStorjNewParams {
	return &PostStorageStorjNewParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostStorageStorjNewParamsWithTimeout creates a new PostStorageStorjNewParams object
// with the ability to set a timeout on a request.
func NewPostStorageStorjNewParamsWithTimeout(timeout time.Duration) *PostStorageStorjNewParams {
	return &PostStorageStorjNewParams{
		timeout: timeout,
	}
}

// NewPostStorageStorjNewParamsWithContext creates a new PostStorageStorjNewParams object
// with the ability to set a context for a request.
func NewPostStorageStorjNewParamsWithContext(ctx context.Context) *PostStorageStorjNewParams {
	return &PostStorageStorjNewParams{
		Context: ctx,
	}
}

// NewPostStorageStorjNewParamsWithHTTPClient creates a new PostStorageStorjNewParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostStorageStorjNewParamsWithHTTPClient(client *http.Client) *PostStorageStorjNewParams {
	return &PostStorageStorjNewParams{
		HTTPClient: client,
	}
}

/*
PostStorageStorjNewParams contains all the parameters to send to the API endpoint

	for the post storage storj new operation.

	Typically these are written to a http.Request.
*/
type PostStorageStorjNewParams struct {

	/* Request.

	   Request body
	*/
	Request *models.StorageCreateStorjNewStorageRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post storage storj new params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostStorageStorjNewParams) WithDefaults() *PostStorageStorjNewParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post storage storj new params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostStorageStorjNewParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post storage storj new params
func (o *PostStorageStorjNewParams) WithTimeout(timeout time.Duration) *PostStorageStorjNewParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post storage storj new params
func (o *PostStorageStorjNewParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post storage storj new params
func (o *PostStorageStorjNewParams) WithContext(ctx context.Context) *PostStorageStorjNewParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post storage storj new params
func (o *PostStorageStorjNewParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post storage storj new params
func (o *PostStorageStorjNewParams) WithHTTPClient(client *http.Client) *PostStorageStorjNewParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post storage storj new params
func (o *PostStorageStorjNewParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the post storage storj new params
func (o *PostStorageStorjNewParams) WithRequest(request *models.StorageCreateStorjNewStorageRequest) *PostStorageStorjNewParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the post storage storj new params
func (o *PostStorageStorjNewParams) SetRequest(request *models.StorageCreateStorjNewStorageRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *PostStorageStorjNewParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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