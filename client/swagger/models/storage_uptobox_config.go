// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StorageUptoboxConfig storage uptobox config
//
// swagger:model storage.UptoboxConfig
type StorageUptoboxConfig struct {

	// Your access token.
	AccessToken string `json:"accessToken,omitempty"`

	// The encoding for the backend.
	Encoding *string `json:"encoding,omitempty"`
}

// Validate validates this storage uptobox config
func (m *StorageUptoboxConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this storage uptobox config based on context it is used
func (m *StorageUptoboxConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StorageUptoboxConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageUptoboxConfig) UnmarshalBinary(b []byte) error {
	var res StorageUptoboxConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
