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

// NewPostStorageZohoParams creates a new PostStorageZohoParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostStorageZohoParams() *PostStorageZohoParams {
	return &PostStorageZohoParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostStorageZohoParamsWithTimeout creates a new PostStorageZohoParams object
// with the ability to set a timeout on a request.
func NewPostStorageZohoParamsWithTimeout(timeout time.Duration) *PostStorageZohoParams {
	return &PostStorageZohoParams{
		timeout: timeout,
	}
}

// NewPostStorageZohoParamsWithContext creates a new PostStorageZohoParams object
// with the ability to set a context for a request.
func NewPostStorageZohoParamsWithContext(ctx context.Context) *PostStorageZohoParams {
	return &PostStorageZohoParams{
		Context: ctx,
	}
}

// NewPostStorageZohoParamsWithHTTPClient creates a new PostStorageZohoParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostStorageZohoParamsWithHTTPClient(client *http.Client) *PostStorageZohoParams {
	return &PostStorageZohoParams{
		HTTPClient: client,
	}
}

/*
PostStorageZohoParams contains all the parameters to send to the API endpoint

	for the post storage zoho operation.

	Typically these are written to a http.Request.
*/
type PostStorageZohoParams struct {

	/* Request.

	   Request body
	*/
	Request *models.StorageCreateZohoStorageRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post storage zoho params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostStorageZohoParams) WithDefaults() *PostStorageZohoParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post storage zoho params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostStorageZohoParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post storage zoho params
func (o *PostStorageZohoParams) WithTimeout(timeout time.Duration) *PostStorageZohoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post storage zoho params
func (o *PostStorageZohoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post storage zoho params
func (o *PostStorageZohoParams) WithContext(ctx context.Context) *PostStorageZohoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post storage zoho params
func (o *PostStorageZohoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post storage zoho params
func (o *PostStorageZohoParams) WithHTTPClient(client *http.Client) *PostStorageZohoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post storage zoho params
func (o *PostStorageZohoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the post storage zoho params
func (o *PostStorageZohoParams) WithRequest(request *models.StorageCreateZohoStorageRequest) *PostStorageZohoParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the post storage zoho params
func (o *PostStorageZohoParams) SetRequest(request *models.StorageCreateZohoStorageRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *PostStorageZohoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
