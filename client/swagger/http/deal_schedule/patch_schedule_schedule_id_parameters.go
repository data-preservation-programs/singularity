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

	"github.com/data-preservation-programs/singularity/client/swagger/models"
)

// NewPatchScheduleScheduleIDParams creates a new PatchScheduleScheduleIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchScheduleScheduleIDParams() *PatchScheduleScheduleIDParams {
	return &PatchScheduleScheduleIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchScheduleScheduleIDParamsWithTimeout creates a new PatchScheduleScheduleIDParams object
// with the ability to set a timeout on a request.
func NewPatchScheduleScheduleIDParamsWithTimeout(timeout time.Duration) *PatchScheduleScheduleIDParams {
	return &PatchScheduleScheduleIDParams{
		timeout: timeout,
	}
}

// NewPatchScheduleScheduleIDParamsWithContext creates a new PatchScheduleScheduleIDParams object
// with the ability to set a context for a request.
func NewPatchScheduleScheduleIDParamsWithContext(ctx context.Context) *PatchScheduleScheduleIDParams {
	return &PatchScheduleScheduleIDParams{
		Context: ctx,
	}
}

// NewPatchScheduleScheduleIDParamsWithHTTPClient creates a new PatchScheduleScheduleIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchScheduleScheduleIDParamsWithHTTPClient(client *http.Client) *PatchScheduleScheduleIDParams {
	return &PatchScheduleScheduleIDParams{
		HTTPClient: client,
	}
}

/*
PatchScheduleScheduleIDParams contains all the parameters to send to the API endpoint

	for the patch schedule schedule ID operation.

	Typically these are written to a http.Request.
*/
type PatchScheduleScheduleIDParams struct {

	/* Body.

	   Update request
	*/
	Body *models.ScheduleUpdateRequest

	/* ScheduleID.

	   Schedule ID
	*/
	ScheduleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch schedule schedule ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchScheduleScheduleIDParams) WithDefaults() *PatchScheduleScheduleIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch schedule schedule ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchScheduleScheduleIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch schedule schedule ID params
func (o *PatchScheduleScheduleIDParams) WithTimeout(timeout time.Duration) *PatchScheduleScheduleIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch schedule schedule ID params
func (o *PatchScheduleScheduleIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch schedule schedule ID params
func (o *PatchScheduleScheduleIDParams) WithContext(ctx context.Context) *PatchScheduleScheduleIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch schedule schedule ID params
func (o *PatchScheduleScheduleIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch schedule schedule ID params
func (o *PatchScheduleScheduleIDParams) WithHTTPClient(client *http.Client) *PatchScheduleScheduleIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch schedule schedule ID params
func (o *PatchScheduleScheduleIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch schedule schedule ID params
func (o *PatchScheduleScheduleIDParams) WithBody(body *models.ScheduleUpdateRequest) *PatchScheduleScheduleIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch schedule schedule ID params
func (o *PatchScheduleScheduleIDParams) SetBody(body *models.ScheduleUpdateRequest) {
	o.Body = body
}

// WithScheduleID adds the scheduleID to the patch schedule schedule ID params
func (o *PatchScheduleScheduleIDParams) WithScheduleID(scheduleID string) *PatchScheduleScheduleIDParams {
	o.SetScheduleID(scheduleID)
	return o
}

// SetScheduleID adds the scheduleId to the patch schedule schedule ID params
func (o *PatchScheduleScheduleIDParams) SetScheduleID(scheduleID string) {
	o.ScheduleID = scheduleID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchScheduleScheduleIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param scheduleId
	if err := r.SetPathParam("scheduleId", o.ScheduleID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
