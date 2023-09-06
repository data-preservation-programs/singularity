// Code generated by go-swagger; DO NOT EDIT.

package data_source

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

// NewPostSourceIDRepackParams creates a new PostSourceIDRepackParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostSourceIDRepackParams() *PostSourceIDRepackParams {
	return &PostSourceIDRepackParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostSourceIDRepackParamsWithTimeout creates a new PostSourceIDRepackParams object
// with the ability to set a timeout on a request.
func NewPostSourceIDRepackParamsWithTimeout(timeout time.Duration) *PostSourceIDRepackParams {
	return &PostSourceIDRepackParams{
		timeout: timeout,
	}
}

// NewPostSourceIDRepackParamsWithContext creates a new PostSourceIDRepackParams object
// with the ability to set a context for a request.
func NewPostSourceIDRepackParamsWithContext(ctx context.Context) *PostSourceIDRepackParams {
	return &PostSourceIDRepackParams{
		Context: ctx,
	}
}

// NewPostSourceIDRepackParamsWithHTTPClient creates a new PostSourceIDRepackParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostSourceIDRepackParamsWithHTTPClient(client *http.Client) *PostSourceIDRepackParams {
	return &PostSourceIDRepackParams{
		HTTPClient: client,
	}
}

/*
PostSourceIDRepackParams contains all the parameters to send to the API endpoint

	for the post source ID repack operation.

	Typically these are written to a http.Request.
*/
type PostSourceIDRepackParams struct {

	/* ID.

	   Source ID
	*/
	ID string

	/* Request.

	   Request body
	*/
	Request *models.DatasourceRepackRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post source ID repack params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostSourceIDRepackParams) WithDefaults() *PostSourceIDRepackParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post source ID repack params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostSourceIDRepackParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post source ID repack params
func (o *PostSourceIDRepackParams) WithTimeout(timeout time.Duration) *PostSourceIDRepackParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post source ID repack params
func (o *PostSourceIDRepackParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post source ID repack params
func (o *PostSourceIDRepackParams) WithContext(ctx context.Context) *PostSourceIDRepackParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post source ID repack params
func (o *PostSourceIDRepackParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post source ID repack params
func (o *PostSourceIDRepackParams) WithHTTPClient(client *http.Client) *PostSourceIDRepackParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post source ID repack params
func (o *PostSourceIDRepackParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the post source ID repack params
func (o *PostSourceIDRepackParams) WithID(id string) *PostSourceIDRepackParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post source ID repack params
func (o *PostSourceIDRepackParams) SetID(id string) {
	o.ID = id
}

// WithRequest adds the request to the post source ID repack params
func (o *PostSourceIDRepackParams) WithRequest(request *models.DatasourceRepackRequest) *PostSourceIDRepackParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the post source ID repack params
func (o *PostSourceIDRepackParams) SetRequest(request *models.DatasourceRepackRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *PostSourceIDRepackParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}
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