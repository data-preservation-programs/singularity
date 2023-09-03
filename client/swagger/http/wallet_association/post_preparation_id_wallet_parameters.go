// Code generated by go-swagger; DO NOT EDIT.

package wallet_association

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
	"github.com/go-openapi/swag"
)

// NewPostPreparationIDWalletParams creates a new PostPreparationIDWalletParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostPreparationIDWalletParams() *PostPreparationIDWalletParams {
	return &PostPreparationIDWalletParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostPreparationIDWalletParamsWithTimeout creates a new PostPreparationIDWalletParams object
// with the ability to set a timeout on a request.
func NewPostPreparationIDWalletParamsWithTimeout(timeout time.Duration) *PostPreparationIDWalletParams {
	return &PostPreparationIDWalletParams{
		timeout: timeout,
	}
}

// NewPostPreparationIDWalletParamsWithContext creates a new PostPreparationIDWalletParams object
// with the ability to set a context for a request.
func NewPostPreparationIDWalletParamsWithContext(ctx context.Context) *PostPreparationIDWalletParams {
	return &PostPreparationIDWalletParams{
		Context: ctx,
	}
}

// NewPostPreparationIDWalletParamsWithHTTPClient creates a new PostPreparationIDWalletParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostPreparationIDWalletParamsWithHTTPClient(client *http.Client) *PostPreparationIDWalletParams {
	return &PostPreparationIDWalletParams{
		HTTPClient: client,
	}
}

/*
PostPreparationIDWalletParams contains all the parameters to send to the API endpoint

	for the post preparation ID wallet operation.

	Typically these are written to a http.Request.
*/
type PostPreparationIDWalletParams struct {

	/* ID.

	   Preparation ID or name
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post preparation ID wallet params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostPreparationIDWalletParams) WithDefaults() *PostPreparationIDWalletParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post preparation ID wallet params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostPreparationIDWalletParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post preparation ID wallet params
func (o *PostPreparationIDWalletParams) WithTimeout(timeout time.Duration) *PostPreparationIDWalletParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post preparation ID wallet params
func (o *PostPreparationIDWalletParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post preparation ID wallet params
func (o *PostPreparationIDWalletParams) WithContext(ctx context.Context) *PostPreparationIDWalletParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post preparation ID wallet params
func (o *PostPreparationIDWalletParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post preparation ID wallet params
func (o *PostPreparationIDWalletParams) WithHTTPClient(client *http.Client) *PostPreparationIDWalletParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post preparation ID wallet params
func (o *PostPreparationIDWalletParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the post preparation ID wallet params
func (o *PostPreparationIDWalletParams) WithID(id int64) *PostPreparationIDWalletParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post preparation ID wallet params
func (o *PostPreparationIDWalletParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostPreparationIDWalletParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
