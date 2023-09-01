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

// NewPostStorageHidriveParams creates a new PostStorageHidriveParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostStorageHidriveParams() *PostStorageHidriveParams {
	return &PostStorageHidriveParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostStorageHidriveParamsWithTimeout creates a new PostStorageHidriveParams object
// with the ability to set a timeout on a request.
func NewPostStorageHidriveParamsWithTimeout(timeout time.Duration) *PostStorageHidriveParams {
	return &PostStorageHidriveParams{
		timeout: timeout,
	}
}

// NewPostStorageHidriveParamsWithContext creates a new PostStorageHidriveParams object
// with the ability to set a context for a request.
func NewPostStorageHidriveParamsWithContext(ctx context.Context) *PostStorageHidriveParams {
	return &PostStorageHidriveParams{
		Context: ctx,
	}
}

// NewPostStorageHidriveParamsWithHTTPClient creates a new PostStorageHidriveParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostStorageHidriveParamsWithHTTPClient(client *http.Client) *PostStorageHidriveParams {
	return &PostStorageHidriveParams{
		HTTPClient: client,
	}
}

/*
PostStorageHidriveParams contains all the parameters to send to the API endpoint

	for the post storage hidrive operation.

	Typically these are written to a http.Request.
*/
type PostStorageHidriveParams struct {

	/* Request.

	   Request body
	*/
	Request *models.StorageCreateHidriveStorageRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post storage hidrive params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostStorageHidriveParams) WithDefaults() *PostStorageHidriveParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post storage hidrive params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostStorageHidriveParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post storage hidrive params
func (o *PostStorageHidriveParams) WithTimeout(timeout time.Duration) *PostStorageHidriveParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post storage hidrive params
func (o *PostStorageHidriveParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post storage hidrive params
func (o *PostStorageHidriveParams) WithContext(ctx context.Context) *PostStorageHidriveParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post storage hidrive params
func (o *PostStorageHidriveParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post storage hidrive params
func (o *PostStorageHidriveParams) WithHTTPClient(client *http.Client) *PostStorageHidriveParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post storage hidrive params
func (o *PostStorageHidriveParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the post storage hidrive params
func (o *PostStorageHidriveParams) WithRequest(request *models.StorageCreateHidriveStorageRequest) *PostStorageHidriveParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the post storage hidrive params
func (o *PostStorageHidriveParams) SetRequest(request *models.StorageCreateHidriveStorageRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *PostStorageHidriveParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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