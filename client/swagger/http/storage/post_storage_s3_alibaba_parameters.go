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

// NewPostStorageS3AlibabaParams creates a new PostStorageS3AlibabaParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostStorageS3AlibabaParams() *PostStorageS3AlibabaParams {
	return &PostStorageS3AlibabaParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostStorageS3AlibabaParamsWithTimeout creates a new PostStorageS3AlibabaParams object
// with the ability to set a timeout on a request.
func NewPostStorageS3AlibabaParamsWithTimeout(timeout time.Duration) *PostStorageS3AlibabaParams {
	return &PostStorageS3AlibabaParams{
		timeout: timeout,
	}
}

// NewPostStorageS3AlibabaParamsWithContext creates a new PostStorageS3AlibabaParams object
// with the ability to set a context for a request.
func NewPostStorageS3AlibabaParamsWithContext(ctx context.Context) *PostStorageS3AlibabaParams {
	return &PostStorageS3AlibabaParams{
		Context: ctx,
	}
}

// NewPostStorageS3AlibabaParamsWithHTTPClient creates a new PostStorageS3AlibabaParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostStorageS3AlibabaParamsWithHTTPClient(client *http.Client) *PostStorageS3AlibabaParams {
	return &PostStorageS3AlibabaParams{
		HTTPClient: client,
	}
}

/*
PostStorageS3AlibabaParams contains all the parameters to send to the API endpoint

	for the post storage s3 alibaba operation.

	Typically these are written to a http.Request.
*/
type PostStorageS3AlibabaParams struct {

	/* Request.

	   Request body
	*/
	Request *models.StorageCreateS3AlibabaStorageRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post storage s3 alibaba params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostStorageS3AlibabaParams) WithDefaults() *PostStorageS3AlibabaParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post storage s3 alibaba params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostStorageS3AlibabaParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post storage s3 alibaba params
func (o *PostStorageS3AlibabaParams) WithTimeout(timeout time.Duration) *PostStorageS3AlibabaParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post storage s3 alibaba params
func (o *PostStorageS3AlibabaParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post storage s3 alibaba params
func (o *PostStorageS3AlibabaParams) WithContext(ctx context.Context) *PostStorageS3AlibabaParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post storage s3 alibaba params
func (o *PostStorageS3AlibabaParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post storage s3 alibaba params
func (o *PostStorageS3AlibabaParams) WithHTTPClient(client *http.Client) *PostStorageS3AlibabaParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post storage s3 alibaba params
func (o *PostStorageS3AlibabaParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the post storage s3 alibaba params
func (o *PostStorageS3AlibabaParams) WithRequest(request *models.StorageCreateS3AlibabaStorageRequest) *PostStorageS3AlibabaParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the post storage s3 alibaba params
func (o *PostStorageS3AlibabaParams) SetRequest(request *models.StorageCreateS3AlibabaStorageRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *PostStorageS3AlibabaParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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