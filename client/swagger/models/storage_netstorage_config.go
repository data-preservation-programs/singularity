// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StorageNetstorageConfig storage netstorage config
//
// swagger:model storage.NetstorageConfig
type StorageNetstorageConfig struct {

	// Set the NetStorage account name
	Account string `json:"account,omitempty"`

	// Domain+path of NetStorage host to connect to.
	Host string `json:"host,omitempty"`

	// Select between HTTP or HTTPS protocol.
	// Example: http
	Protocol *string `json:"protocol,omitempty"`

	// Set the NetStorage account secret/G2O key for authentication.
	Secret string `json:"secret,omitempty"`
}

// Validate validates this storage netstorage config
func (m *StorageNetstorageConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this storage netstorage config based on context it is used
func (m *StorageNetstorageConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StorageNetstorageConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageNetstorageConfig) UnmarshalBinary(b []byte) error {
	var res StorageNetstorageConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
