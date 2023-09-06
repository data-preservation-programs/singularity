// Code generated by go-swagger; DO NOT EDIT.

package wallet

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

// NewGetDatasetDatasetNameWalletParams creates a new GetDatasetDatasetNameWalletParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDatasetDatasetNameWalletParams() *GetDatasetDatasetNameWalletParams {
	return &GetDatasetDatasetNameWalletParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDatasetDatasetNameWalletParamsWithTimeout creates a new GetDatasetDatasetNameWalletParams object
// with the ability to set a timeout on a request.
func NewGetDatasetDatasetNameWalletParamsWithTimeout(timeout time.Duration) *GetDatasetDatasetNameWalletParams {
	return &GetDatasetDatasetNameWalletParams{
		timeout: timeout,
	}
}

// NewGetDatasetDatasetNameWalletParamsWithContext creates a new GetDatasetDatasetNameWalletParams object
// with the ability to set a context for a request.
func NewGetDatasetDatasetNameWalletParamsWithContext(ctx context.Context) *GetDatasetDatasetNameWalletParams {
	return &GetDatasetDatasetNameWalletParams{
		Context: ctx,
	}
}

// NewGetDatasetDatasetNameWalletParamsWithHTTPClient creates a new GetDatasetDatasetNameWalletParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDatasetDatasetNameWalletParamsWithHTTPClient(client *http.Client) *GetDatasetDatasetNameWalletParams {
	return &GetDatasetDatasetNameWalletParams{
		HTTPClient: client,
	}
}

/*
GetDatasetDatasetNameWalletParams contains all the parameters to send to the API endpoint

	for the get dataset dataset name wallet operation.

	Typically these are written to a http.Request.
*/
type GetDatasetDatasetNameWalletParams struct {

	/* DatasetName.

	   Dataset name
	*/
	DatasetName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get dataset dataset name wallet params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDatasetDatasetNameWalletParams) WithDefaults() *GetDatasetDatasetNameWalletParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get dataset dataset name wallet params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDatasetDatasetNameWalletParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get dataset dataset name wallet params
func (o *GetDatasetDatasetNameWalletParams) WithTimeout(timeout time.Duration) *GetDatasetDatasetNameWalletParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get dataset dataset name wallet params
func (o *GetDatasetDatasetNameWalletParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get dataset dataset name wallet params
func (o *GetDatasetDatasetNameWalletParams) WithContext(ctx context.Context) *GetDatasetDatasetNameWalletParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get dataset dataset name wallet params
func (o *GetDatasetDatasetNameWalletParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get dataset dataset name wallet params
func (o *GetDatasetDatasetNameWalletParams) WithHTTPClient(client *http.Client) *GetDatasetDatasetNameWalletParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get dataset dataset name wallet params
func (o *GetDatasetDatasetNameWalletParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDatasetName adds the datasetName to the get dataset dataset name wallet params
func (o *GetDatasetDatasetNameWalletParams) WithDatasetName(datasetName string) *GetDatasetDatasetNameWalletParams {
	o.SetDatasetName(datasetName)
	return o
}

// SetDatasetName adds the datasetName to the get dataset dataset name wallet params
func (o *GetDatasetDatasetNameWalletParams) SetDatasetName(datasetName string) {
	o.DatasetName = datasetName
}

// WriteToRequest writes these params to a swagger request
func (o *GetDatasetDatasetNameWalletParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param datasetName
	if err := r.SetPathParam("datasetName", o.DatasetName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}