// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StorageHTTPConfig storage http config
//
// swagger:model storage.httpConfig
type StorageHTTPConfig struct {

	// Set HTTP headers for all transactions.
	Headers string `json:"headers,omitempty"`

	// Don't use HEAD requests.
	NoHead *bool `json:"noHead,omitempty"`

	// Set this if the site doesn't end directories with /.
	NoSlash *bool `json:"noSlash,omitempty"`

	// URL of HTTP host to connect to.
	URL string `json:"url,omitempty"`
}

// Validate validates this storage http config
func (m *StorageHTTPConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this storage http config based on context it is used
func (m *StorageHTTPConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StorageHTTPConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageHTTPConfig) UnmarshalBinary(b []byte) error {
	var res StorageHTTPConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
