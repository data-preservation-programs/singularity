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

// NewCreateS3WasabiStorageParams creates a new CreateS3WasabiStorageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateS3WasabiStorageParams() *CreateS3WasabiStorageParams {
	return &CreateS3WasabiStorageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateS3WasabiStorageParamsWithTimeout creates a new CreateS3WasabiStorageParams object
// with the ability to set a timeout on a request.
func NewCreateS3WasabiStorageParamsWithTimeout(timeout time.Duration) *CreateS3WasabiStorageParams {
	return &CreateS3WasabiStorageParams{
		timeout: timeout,
	}
}

// NewCreateS3WasabiStorageParamsWithContext creates a new CreateS3WasabiStorageParams object
// with the ability to set a context for a request.
func NewCreateS3WasabiStorageParamsWithContext(ctx context.Context) *CreateS3WasabiStorageParams {
	return &CreateS3WasabiStorageParams{
		Context: ctx,
	}
}

// NewCreateS3WasabiStorageParamsWithHTTPClient creates a new CreateS3WasabiStorageParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateS3WasabiStorageParamsWithHTTPClient(client *http.Client) *CreateS3WasabiStorageParams {
	return &CreateS3WasabiStorageParams{
		HTTPClient: client,
	}
}

/*
CreateS3WasabiStorageParams contains all the parameters to send to the API endpoint

	for the create s3 wasabi storage operation.

	Typically these are written to a http.Request.
*/
type CreateS3WasabiStorageParams struct {

	/* Request.

	   Request body
	*/
	Request *models.StorageCreateS3WasabiStorageRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create s3 wasabi storage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateS3WasabiStorageParams) WithDefaults() *CreateS3WasabiStorageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create s3 wasabi storage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateS3WasabiStorageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create s3 wasabi storage params
func (o *CreateS3WasabiStorageParams) WithTimeout(timeout time.Duration) *CreateS3WasabiStorageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create s3 wasabi storage params
func (o *CreateS3WasabiStorageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create s3 wasabi storage params
func (o *CreateS3WasabiStorageParams) WithContext(ctx context.Context) *CreateS3WasabiStorageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create s3 wasabi storage params
func (o *CreateS3WasabiStorageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create s3 wasabi storage params
func (o *CreateS3WasabiStorageParams) WithHTTPClient(client *http.Client) *CreateS3WasabiStorageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create s3 wasabi storage params
func (o *CreateS3WasabiStorageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the create s3 wasabi storage params
func (o *CreateS3WasabiStorageParams) WithRequest(request *models.StorageCreateS3WasabiStorageRequest) *CreateS3WasabiStorageParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the create s3 wasabi storage params
func (o *CreateS3WasabiStorageParams) SetRequest(request *models.StorageCreateS3WasabiStorageRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *CreateS3WasabiStorageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
