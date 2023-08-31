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

// NewPostStorageS3OtherParams creates a new PostStorageS3OtherParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostStorageS3OtherParams() *PostStorageS3OtherParams {
	return &PostStorageS3OtherParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostStorageS3OtherParamsWithTimeout creates a new PostStorageS3OtherParams object
// with the ability to set a timeout on a request.
func NewPostStorageS3OtherParamsWithTimeout(timeout time.Duration) *PostStorageS3OtherParams {
	return &PostStorageS3OtherParams{
		timeout: timeout,
	}
}

// NewPostStorageS3OtherParamsWithContext creates a new PostStorageS3OtherParams object
// with the ability to set a context for a request.
func NewPostStorageS3OtherParamsWithContext(ctx context.Context) *PostStorageS3OtherParams {
	return &PostStorageS3OtherParams{
		Context: ctx,
	}
}

// NewPostStorageS3OtherParamsWithHTTPClient creates a new PostStorageS3OtherParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostStorageS3OtherParamsWithHTTPClient(client *http.Client) *PostStorageS3OtherParams {
	return &PostStorageS3OtherParams{
		HTTPClient: client,
	}
}

/*
PostStorageS3OtherParams contains all the parameters to send to the API endpoint

	for the post storage s3 other operation.

	Typically these are written to a http.Request.
*/
type PostStorageS3OtherParams struct {

	/* Request.

	   Request body
	*/
	Request models.StorageCreateS3OtherStorageRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post storage s3 other params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostStorageS3OtherParams) WithDefaults() *PostStorageS3OtherParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post storage s3 other params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostStorageS3OtherParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post storage s3 other params
func (o *PostStorageS3OtherParams) WithTimeout(timeout time.Duration) *PostStorageS3OtherParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post storage s3 other params
func (o *PostStorageS3OtherParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post storage s3 other params
func (o *PostStorageS3OtherParams) WithContext(ctx context.Context) *PostStorageS3OtherParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post storage s3 other params
func (o *PostStorageS3OtherParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post storage s3 other params
func (o *PostStorageS3OtherParams) WithHTTPClient(client *http.Client) *PostStorageS3OtherParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post storage s3 other params
func (o *PostStorageS3OtherParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the post storage s3 other params
func (o *PostStorageS3OtherParams) WithRequest(request models.StorageCreateS3OtherStorageRequest) *PostStorageS3OtherParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the post storage s3 other params
func (o *PostStorageS3OtherParams) SetRequest(request models.StorageCreateS3OtherStorageRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *PostStorageS3OtherParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
