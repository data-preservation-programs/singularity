// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StorageCreateOosInstancePrincipalAuthStorageRequest storage create oos instance principal auth storage request
//
// swagger:model storage.CreateOosInstance_principal_authStorageRequest
type StorageCreateOosInstancePrincipalAuthStorageRequest struct {

	// config
	Config *StorageOosInstancePrincipalAuthConfig `json:"config,omitempty"`

	// Name of the storage, must be unique
	// Example: my-storage
	Name string `json:"name,omitempty"`

	// Path of the storage
	Path string `json:"path,omitempty"`
}

// Validate validates this storage create oos instance principal auth storage request
func (m *StorageCreateOosInstancePrincipalAuthStorageRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StorageCreateOosInstancePrincipalAuthStorageRequest) validateConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.Config) { // not required
		return nil
	}

	if m.Config != nil {
		if err := m.Config.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("config")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this storage create oos instance principal auth storage request based on the context it is used
func (m *StorageCreateOosInstancePrincipalAuthStorageRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StorageCreateOosInstancePrincipalAuthStorageRequest) contextValidateConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.Config != nil {

		if swag.IsZero(m.Config) { // not required
			return nil
		}

		if err := m.Config.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("config")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StorageCreateOosInstancePrincipalAuthStorageRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageCreateOosInstancePrincipalAuthStorageRequest) UnmarshalBinary(b []byte) error {
	var res StorageCreateOosInstancePrincipalAuthStorageRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
