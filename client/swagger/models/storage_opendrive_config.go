// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StorageOpendriveConfig storage opendrive config
//
// swagger:model storage.opendriveConfig
type StorageOpendriveConfig struct {

	// Files will be uploaded in chunks this size.
	ChunkSize *string `json:"chunkSize,omitempty"`

	// The encoding for the backend.
	Encoding *string `json:"encoding,omitempty"`

	// Password.
	Password string `json:"password,omitempty"`

	// Username.
	Username string `json:"username,omitempty"`
}

// Validate validates this storage opendrive config
func (m *StorageOpendriveConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this storage opendrive config based on context it is used
func (m *StorageOpendriveConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StorageOpendriveConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageOpendriveConfig) UnmarshalBinary(b []byte) error {
	var res StorageOpendriveConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
