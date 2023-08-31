// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StorageStorjNewConfig storage storj new config
//
// swagger:model storage.StorjNewConfig
type StorageStorjNewConfig struct {

	// API key.
	APIKey string `json:"apiKey,omitempty"`

	// Encryption passphrase.
	Passphrase string `json:"passphrase,omitempty"`

	// Satellite address.
	// Example: us1.storj.io
	SatelliteAddress *string `json:"satelliteAddress,omitempty"`
}

// Validate validates this storage storj new config
func (m *StorageStorjNewConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this storage storj new config based on context it is used
func (m *StorageStorjNewConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StorageStorjNewConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageStorjNewConfig) UnmarshalBinary(b []byte) error {
	var res StorageStorjNewConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
