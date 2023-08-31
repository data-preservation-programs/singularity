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

// NewPostStorageWebdavParams creates a new PostStorageWebdavParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostStorageWebdavParams() *PostStorageWebdavParams {
	return &PostStorageWebdavParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostStorageWebdavParamsWithTimeout creates a new PostStorageWebdavParams object
// with the ability to set a timeout on a request.
func NewPostStorageWebdavParamsWithTimeout(timeout time.Duration) *PostStorageWebdavParams {
	return &PostStorageWebdavParams{
		timeout: timeout,
	}
}

// NewPostStorageWebdavParamsWithContext creates a new PostStorageWebdavParams object
// with the ability to set a context for a request.
func NewPostStorageWebdavParamsWithContext(ctx context.Context) *PostStorageWebdavParams {
	return &PostStorageWebdavParams{
		Context: ctx,
	}
}

// NewPostStorageWebdavParamsWithHTTPClient creates a new PostStorageWebdavParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostStorageWebdavParamsWithHTTPClient(client *http.Client) *PostStorageWebdavParams {
	return &PostStorageWebdavParams{
		HTTPClient: client,
	}
}

/*
PostStorageWebdavParams contains all the parameters to send to the API endpoint

	for the post storage webdav operation.

	Typically these are written to a http.Request.
*/
type PostStorageWebdavParams struct {

	/* Request.

	   Request body
	*/
	Request *models.StorageCreateWebdavStorageRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post storage webdav params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostStorageWebdavParams) WithDefaults() *PostStorageWebdavParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post storage webdav params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostStorageWebdavParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post storage webdav params
func (o *PostStorageWebdavParams) WithTimeout(timeout time.Duration) *PostStorageWebdavParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post storage webdav params
func (o *PostStorageWebdavParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post storage webdav params
func (o *PostStorageWebdavParams) WithContext(ctx context.Context) *PostStorageWebdavParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post storage webdav params
func (o *PostStorageWebdavParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post storage webdav params
func (o *PostStorageWebdavParams) WithHTTPClient(client *http.Client) *PostStorageWebdavParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post storage webdav params
func (o *PostStorageWebdavParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the post storage webdav params
func (o *PostStorageWebdavParams) WithRequest(request *models.StorageCreateWebdavStorageRequest) *PostStorageWebdavParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the post storage webdav params
func (o *PostStorageWebdavParams) SetRequest(request *models.StorageCreateWebdavStorageRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *PostStorageWebdavParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
