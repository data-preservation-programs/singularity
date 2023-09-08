// Code generated by go-swagger; DO NOT EDIT.

package deal_schedule

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
)

// NewListPreparationSchedulesParams creates a new ListPreparationSchedulesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListPreparationSchedulesParams() *ListPreparationSchedulesParams {
	return &ListPreparationSchedulesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListPreparationSchedulesParamsWithTimeout creates a new ListPreparationSchedulesParams object
// with the ability to set a timeout on a request.
func NewListPreparationSchedulesParamsWithTimeout(timeout time.Duration) *ListPreparationSchedulesParams {
	return &ListPreparationSchedulesParams{
		timeout: timeout,
	}
}

// NewListPreparationSchedulesParamsWithContext creates a new ListPreparationSchedulesParams object
// with the ability to set a context for a request.
func NewListPreparationSchedulesParamsWithContext(ctx context.Context) *ListPreparationSchedulesParams {
	return &ListPreparationSchedulesParams{
		Context: ctx,
	}
}

// NewListPreparationSchedulesParamsWithHTTPClient creates a new ListPreparationSchedulesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListPreparationSchedulesParamsWithHTTPClient(client *http.Client) *ListPreparationSchedulesParams {
	return &ListPreparationSchedulesParams{
		HTTPClient: client,
	}
}

/*
ListPreparationSchedulesParams contains all the parameters to send to the API endpoint

	for the list preparation schedules operation.

	Typically these are written to a http.Request.
*/
type ListPreparationSchedulesParams struct {

	/* ID.

	   Preparation ID or name
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list preparation schedules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListPreparationSchedulesParams) WithDefaults() *ListPreparationSchedulesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list preparation schedules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListPreparationSchedulesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list preparation schedules params
func (o *ListPreparationSchedulesParams) WithTimeout(timeout time.Duration) *ListPreparationSchedulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list preparation schedules params
func (o *ListPreparationSchedulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list preparation schedules params
func (o *ListPreparationSchedulesParams) WithContext(ctx context.Context) *ListPreparationSchedulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list preparation schedules params
func (o *ListPreparationSchedulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list preparation schedules params
func (o *ListPreparationSchedulesParams) WithHTTPClient(client *http.Client) *ListPreparationSchedulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list preparation schedules params
func (o *ListPreparationSchedulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the list preparation schedules params
func (o *ListPreparationSchedulesParams) WithID(id string) *ListPreparationSchedulesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the list preparation schedules params
func (o *ListPreparationSchedulesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ListPreparationSchedulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
