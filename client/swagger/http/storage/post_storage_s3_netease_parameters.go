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

// NewPostStorageS3NeteaseParams creates a new PostStorageS3NeteaseParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostStorageS3NeteaseParams() *PostStorageS3NeteaseParams {
	return &PostStorageS3NeteaseParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostStorageS3NeteaseParamsWithTimeout creates a new PostStorageS3NeteaseParams object
// with the ability to set a timeout on a request.
func NewPostStorageS3NeteaseParamsWithTimeout(timeout time.Duration) *PostStorageS3NeteaseParams {
	return &PostStorageS3NeteaseParams{
		timeout: timeout,
	}
}

// NewPostStorageS3NeteaseParamsWithContext creates a new PostStorageS3NeteaseParams object
// with the ability to set a context for a request.
func NewPostStorageS3NeteaseParamsWithContext(ctx context.Context) *PostStorageS3NeteaseParams {
	return &PostStorageS3NeteaseParams{
		Context: ctx,
	}
}

// NewPostStorageS3NeteaseParamsWithHTTPClient creates a new PostStorageS3NeteaseParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostStorageS3NeteaseParamsWithHTTPClient(client *http.Client) *PostStorageS3NeteaseParams {
	return &PostStorageS3NeteaseParams{
		HTTPClient: client,
	}
}

/*
PostStorageS3NeteaseParams contains all the parameters to send to the API endpoint

	for the post storage s3 netease operation.

	Typically these are written to a http.Request.
*/
type PostStorageS3NeteaseParams struct {

	/* Request.

	   Request body
	*/
	Request models.StorageCreateS3NeteaseStorageRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post storage s3 netease params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostStorageS3NeteaseParams) WithDefaults() *PostStorageS3NeteaseParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post storage s3 netease params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostStorageS3NeteaseParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post storage s3 netease params
func (o *PostStorageS3NeteaseParams) WithTimeout(timeout time.Duration) *PostStorageS3NeteaseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post storage s3 netease params
func (o *PostStorageS3NeteaseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post storage s3 netease params
func (o *PostStorageS3NeteaseParams) WithContext(ctx context.Context) *PostStorageS3NeteaseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post storage s3 netease params
func (o *PostStorageS3NeteaseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post storage s3 netease params
func (o *PostStorageS3NeteaseParams) WithHTTPClient(client *http.Client) *PostStorageS3NeteaseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post storage s3 netease params
func (o *PostStorageS3NeteaseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the post storage s3 netease params
func (o *PostStorageS3NeteaseParams) WithRequest(request models.StorageCreateS3NeteaseStorageRequest) *PostStorageS3NeteaseParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the post storage s3 netease params
func (o *PostStorageS3NeteaseParams) SetRequest(request models.StorageCreateS3NeteaseStorageRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *PostStorageS3NeteaseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
