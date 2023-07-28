// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DatasourceQingstorRequest datasource qingstor request
//
// swagger:model datasource.QingstorRequest
type DatasourceQingstorRequest struct {

	// QingStor Access Key ID.
	AccessKeyID string `json:"accessKeyId,omitempty"`

	// Chunk size to use for uploading.
	ChunkSize *string `json:"chunkSize,omitempty"`

	// Number of connection retries.
	ConnectionRetries *string `json:"connectionRetries,omitempty"`

	// Delete the source after exporting to CAR files
	// Required: true
	DeleteAfterExport *bool `json:"deleteAfterExport"`

	// The encoding for the backend.
	Encoding *string `json:"encoding,omitempty"`

	// Enter an endpoint URL to connection QingStor API.
	Endpoint string `json:"endpoint,omitempty"`

	// Get QingStor credentials from runtime.
	EnvAuth *string `json:"envAuth,omitempty"`

	// Automatically rescan the source directory when this interval has passed from last successful scan
	// Required: true
	RescanInterval *string `json:"rescanInterval"`

	// Starting state for scanning
	// Required: true
	ScanningState struct {
		ModelWorkState
	} `json:"scanningState"`

	// QingStor Secret Access Key (password).
	SecretAccessKey string `json:"secretAccessKey,omitempty"`

	// The path of the source to scan items
	// Required: true
	SourcePath *string `json:"sourcePath"`

	// Concurrency for multipart uploads.
	UploadConcurrency *string `json:"uploadConcurrency,omitempty"`

	// Cutoff for switching to chunked upload.
	UploadCutoff *string `json:"uploadCutoff,omitempty"`

	// Zone to connect to.
	Zone string `json:"zone,omitempty"`
}

// Validate validates this datasource qingstor request
func (m *DatasourceQingstorRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeleteAfterExport(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRescanInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScanningState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourcePath(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatasourceQingstorRequest) validateDeleteAfterExport(formats strfmt.Registry) error {

	if err := validate.Required("deleteAfterExport", "body", m.DeleteAfterExport); err != nil {
		return err
	}

	return nil
}

func (m *DatasourceQingstorRequest) validateRescanInterval(formats strfmt.Registry) error {

	if err := validate.Required("rescanInterval", "body", m.RescanInterval); err != nil {
		return err
	}

	return nil
}

func (m *DatasourceQingstorRequest) validateScanningState(formats strfmt.Registry) error {

	return nil
}

func (m *DatasourceQingstorRequest) validateSourcePath(formats strfmt.Registry) error {

	if err := validate.Required("sourcePath", "body", m.SourcePath); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this datasource qingstor request based on the context it is used
func (m *DatasourceQingstorRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateScanningState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatasourceQingstorRequest) contextValidateScanningState(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *DatasourceQingstorRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatasourceQingstorRequest) UnmarshalBinary(b []byte) error {
	var res DatasourceQingstorRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
