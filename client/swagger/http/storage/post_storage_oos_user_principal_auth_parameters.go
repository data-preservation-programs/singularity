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

// NewPostStorageOosUserPrincipalAuthParams creates a new PostStorageOosUserPrincipalAuthParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostStorageOosUserPrincipalAuthParams() *PostStorageOosUserPrincipalAuthParams {
	return &PostStorageOosUserPrincipalAuthParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostStorageOosUserPrincipalAuthParamsWithTimeout creates a new PostStorageOosUserPrincipalAuthParams object
// with the ability to set a timeout on a request.
func NewPostStorageOosUserPrincipalAuthParamsWithTimeout(timeout time.Duration) *PostStorageOosUserPrincipalAuthParams {
	return &PostStorageOosUserPrincipalAuthParams{
		timeout: timeout,
	}
}

// NewPostStorageOosUserPrincipalAuthParamsWithContext creates a new PostStorageOosUserPrincipalAuthParams object
// with the ability to set a context for a request.
func NewPostStorageOosUserPrincipalAuthParamsWithContext(ctx context.Context) *PostStorageOosUserPrincipalAuthParams {
	return &PostStorageOosUserPrincipalAuthParams{
		Context: ctx,
	}
}

// NewPostStorageOosUserPrincipalAuthParamsWithHTTPClient creates a new PostStorageOosUserPrincipalAuthParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostStorageOosUserPrincipalAuthParamsWithHTTPClient(client *http.Client) *PostStorageOosUserPrincipalAuthParams {
	return &PostStorageOosUserPrincipalAuthParams{
		HTTPClient: client,
	}
}

/*
PostStorageOosUserPrincipalAuthParams contains all the parameters to send to the API endpoint

	for the post storage oos user principal auth operation.

	Typically these are written to a http.Request.
*/
type PostStorageOosUserPrincipalAuthParams struct {

	/* Request.

	   Request body
	*/
	Request models.StorageCreateOosUserPrincipalAuthStorageRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post storage oos user principal auth params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostStorageOosUserPrincipalAuthParams) WithDefaults() *PostStorageOosUserPrincipalAuthParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post storage oos user principal auth params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostStorageOosUserPrincipalAuthParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post storage oos user principal auth params
func (o *PostStorageOosUserPrincipalAuthParams) WithTimeout(timeout time.Duration) *PostStorageOosUserPrincipalAuthParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post storage oos user principal auth params
func (o *PostStorageOosUserPrincipalAuthParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post storage oos user principal auth params
func (o *PostStorageOosUserPrincipalAuthParams) WithContext(ctx context.Context) *PostStorageOosUserPrincipalAuthParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post storage oos user principal auth params
func (o *PostStorageOosUserPrincipalAuthParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post storage oos user principal auth params
func (o *PostStorageOosUserPrincipalAuthParams) WithHTTPClient(client *http.Client) *PostStorageOosUserPrincipalAuthParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post storage oos user principal auth params
func (o *PostStorageOosUserPrincipalAuthParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the post storage oos user principal auth params
func (o *PostStorageOosUserPrincipalAuthParams) WithRequest(request models.StorageCreateOosUserPrincipalAuthStorageRequest) *PostStorageOosUserPrincipalAuthParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the post storage oos user principal auth params
func (o *PostStorageOosUserPrincipalAuthParams) SetRequest(request models.StorageCreateOosUserPrincipalAuthStorageRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *PostStorageOosUserPrincipalAuthParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
