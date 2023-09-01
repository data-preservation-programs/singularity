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

// NewPostStorageOosNoAuthParams creates a new PostStorageOosNoAuthParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostStorageOosNoAuthParams() *PostStorageOosNoAuthParams {
	return &PostStorageOosNoAuthParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostStorageOosNoAuthParamsWithTimeout creates a new PostStorageOosNoAuthParams object
// with the ability to set a timeout on a request.
func NewPostStorageOosNoAuthParamsWithTimeout(timeout time.Duration) *PostStorageOosNoAuthParams {
	return &PostStorageOosNoAuthParams{
		timeout: timeout,
	}
}

// NewPostStorageOosNoAuthParamsWithContext creates a new PostStorageOosNoAuthParams object
// with the ability to set a context for a request.
func NewPostStorageOosNoAuthParamsWithContext(ctx context.Context) *PostStorageOosNoAuthParams {
	return &PostStorageOosNoAuthParams{
		Context: ctx,
	}
}

// NewPostStorageOosNoAuthParamsWithHTTPClient creates a new PostStorageOosNoAuthParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostStorageOosNoAuthParamsWithHTTPClient(client *http.Client) *PostStorageOosNoAuthParams {
	return &PostStorageOosNoAuthParams{
		HTTPClient: client,
	}
}

/*
PostStorageOosNoAuthParams contains all the parameters to send to the API endpoint

	for the post storage oos no auth operation.

	Typically these are written to a http.Request.
*/
type PostStorageOosNoAuthParams struct {

	/* Request.

	   Request body
	*/
	Request *models.StorageCreateOosNoAuthStorageRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post storage oos no auth params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostStorageOosNoAuthParams) WithDefaults() *PostStorageOosNoAuthParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post storage oos no auth params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostStorageOosNoAuthParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post storage oos no auth params
func (o *PostStorageOosNoAuthParams) WithTimeout(timeout time.Duration) *PostStorageOosNoAuthParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post storage oos no auth params
func (o *PostStorageOosNoAuthParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post storage oos no auth params
func (o *PostStorageOosNoAuthParams) WithContext(ctx context.Context) *PostStorageOosNoAuthParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post storage oos no auth params
func (o *PostStorageOosNoAuthParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post storage oos no auth params
func (o *PostStorageOosNoAuthParams) WithHTTPClient(client *http.Client) *PostStorageOosNoAuthParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post storage oos no auth params
func (o *PostStorageOosNoAuthParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the post storage oos no auth params
func (o *PostStorageOosNoAuthParams) WithRequest(request *models.StorageCreateOosNoAuthStorageRequest) *PostStorageOosNoAuthParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the post storage oos no auth params
func (o *PostStorageOosNoAuthParams) SetRequest(request *models.StorageCreateOosNoAuthStorageRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *PostStorageOosNoAuthParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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