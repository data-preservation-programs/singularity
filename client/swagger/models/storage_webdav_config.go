// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StorageWebdavConfig storage webdav config
//
// swagger:model storage.webdavConfig
type StorageWebdavConfig struct {

	// Bearer token instead of user/pass (e.g. a Macaroon).
	BearerToken string `json:"bearerToken,omitempty"`

	// Command to run to get a bearer token.
	BearerTokenCommand string `json:"bearerTokenCommand,omitempty"`

	// The encoding for the backend.
	Encoding string `json:"encoding,omitempty"`

	// Set HTTP headers for all transactions.
	Headers string `json:"headers,omitempty"`

	// Password.
	Pass string `json:"pass,omitempty"`

	// URL of http host to connect to.
	URL string `json:"url,omitempty"`

	// User name.
	User string `json:"user,omitempty"`

	// Name of the WebDAV site/service/software you are using.
	// Example: nextcloud
	Vendor string `json:"vendor,omitempty"`
}

// Validate validates this storage webdav config
func (m *StorageWebdavConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this storage webdav config based on context it is used
func (m *StorageWebdavConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StorageWebdavConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageWebdavConfig) UnmarshalBinary(b []byte) error {
	var res StorageWebdavConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
