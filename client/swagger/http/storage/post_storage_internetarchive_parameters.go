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

// NewPostStorageInternetarchiveParams creates a new PostStorageInternetarchiveParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostStorageInternetarchiveParams() *PostStorageInternetarchiveParams {
	return &PostStorageInternetarchiveParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostStorageInternetarchiveParamsWithTimeout creates a new PostStorageInternetarchiveParams object
// with the ability to set a timeout on a request.
func NewPostStorageInternetarchiveParamsWithTimeout(timeout time.Duration) *PostStorageInternetarchiveParams {
	return &PostStorageInternetarchiveParams{
		timeout: timeout,
	}
}

// NewPostStorageInternetarchiveParamsWithContext creates a new PostStorageInternetarchiveParams object
// with the ability to set a context for a request.
func NewPostStorageInternetarchiveParamsWithContext(ctx context.Context) *PostStorageInternetarchiveParams {
	return &PostStorageInternetarchiveParams{
		Context: ctx,
	}
}

// NewPostStorageInternetarchiveParamsWithHTTPClient creates a new PostStorageInternetarchiveParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostStorageInternetarchiveParamsWithHTTPClient(client *http.Client) *PostStorageInternetarchiveParams {
	return &PostStorageInternetarchiveParams{
		HTTPClient: client,
	}
}

/*
PostStorageInternetarchiveParams contains all the parameters to send to the API endpoint

	for the post storage internetarchive operation.

	Typically these are written to a http.Request.
*/
type PostStorageInternetarchiveParams struct {

	/* Request.

	   Request body
	*/
	Request models.StorageCreateInternetarchiveStorageRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post storage internetarchive params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostStorageInternetarchiveParams) WithDefaults() *PostStorageInternetarchiveParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post storage internetarchive params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostStorageInternetarchiveParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post storage internetarchive params
func (o *PostStorageInternetarchiveParams) WithTimeout(timeout time.Duration) *PostStorageInternetarchiveParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post storage internetarchive params
func (o *PostStorageInternetarchiveParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post storage internetarchive params
func (o *PostStorageInternetarchiveParams) WithContext(ctx context.Context) *PostStorageInternetarchiveParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post storage internetarchive params
func (o *PostStorageInternetarchiveParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post storage internetarchive params
func (o *PostStorageInternetarchiveParams) WithHTTPClient(client *http.Client) *PostStorageInternetarchiveParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post storage internetarchive params
func (o *PostStorageInternetarchiveParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the post storage internetarchive params
func (o *PostStorageInternetarchiveParams) WithRequest(request models.StorageCreateInternetarchiveStorageRequest) *PostStorageInternetarchiveParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the post storage internetarchive params
func (o *PostStorageInternetarchiveParams) SetRequest(request models.StorageCreateInternetarchiveStorageRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *PostStorageInternetarchiveParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
