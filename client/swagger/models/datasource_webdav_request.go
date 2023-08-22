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

// DatasourceWebdavRequest datasource webdav request
//
// swagger:model datasource.WebdavRequest
type DatasourceWebdavRequest struct {

	// Bearer token instead of user/pass (e.g. a Macaroon).
	BearerToken string `json:"bearerToken,omitempty"`

	// Command to run to get a bearer token.
	BearerTokenCommand string `json:"bearerTokenCommand,omitempty"`

	// Delete the source after exporting to CAR files
	// Required: true
	DeleteAfterExport *bool `json:"deleteAfterExport"`

	// The encoding for the backend.
	Encoding string `json:"encoding,omitempty"`

	// Set HTTP headers for all transactions.
	Headers string `json:"headers,omitempty"`

	// Password.
	Pass string `json:"pass,omitempty"`

	// Automatically rescan the source directory when this interval has passed from last successful scan
	// Required: true
	RescanInterval *string `json:"rescanInterval"`

	// Starting state for scanning
	// Required: true
	ScanningState struct {
		ModelWorkState
	} `json:"scanningState"`

	// The path of the source to scan files
	// Required: true
	SourcePath *string `json:"sourcePath"`

	// URL of http host to connect to.
	URL string `json:"url,omitempty"`

	// User name.
	User string `json:"user,omitempty"`

	// Name of the WebDAV site/service/software you are using.
	Vendor string `json:"vendor,omitempty"`
}

// Validate validates this datasource webdav request
func (m *DatasourceWebdavRequest) Validate(formats strfmt.Registry) error {
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

func (m *DatasourceWebdavRequest) validateDeleteAfterExport(formats strfmt.Registry) error {

	if err := validate.Required("deleteAfterExport", "body", m.DeleteAfterExport); err != nil {
		return err
	}

	return nil
}

func (m *DatasourceWebdavRequest) validateRescanInterval(formats strfmt.Registry) error {

	if err := validate.Required("rescanInterval", "body", m.RescanInterval); err != nil {
		return err
	}

	return nil
}

func (m *DatasourceWebdavRequest) validateScanningState(formats strfmt.Registry) error {

	return nil
}

func (m *DatasourceWebdavRequest) validateSourcePath(formats strfmt.Registry) error {

	if err := validate.Required("sourcePath", "body", m.SourcePath); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this datasource webdav request based on the context it is used
func (m *DatasourceWebdavRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateScanningState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatasourceWebdavRequest) contextValidateScanningState(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *DatasourceWebdavRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatasourceWebdavRequest) UnmarshalBinary(b []byte) error {
	var res DatasourceWebdavRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
