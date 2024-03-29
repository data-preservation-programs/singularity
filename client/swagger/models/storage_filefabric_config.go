// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StorageFilefabricConfig storage filefabric config
//
// swagger:model storage.filefabricConfig
type StorageFilefabricConfig struct {

	// The encoding for the backend.
	Encoding *string `json:"encoding,omitempty"`

	// Permanent Authentication Token.
	PermanentToken string `json:"permanentToken,omitempty"`

	// ID of the root folder.
	RootFolderID string `json:"rootFolderId,omitempty"`

	// Session Token.
	Token string `json:"token,omitempty"`

	// Token expiry time.
	TokenExpiry string `json:"tokenExpiry,omitempty"`

	// URL of the Enterprise File Fabric to connect to.
	// Example: https://storagemadeeasy.com
	URL string `json:"url,omitempty"`

	// Version read from the file fabric.
	Version string `json:"version,omitempty"`
}

// Validate validates this storage filefabric config
func (m *StorageFilefabricConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this storage filefabric config based on context it is used
func (m *StorageFilefabricConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StorageFilefabricConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageFilefabricConfig) UnmarshalBinary(b []byte) error {
	var res StorageFilefabricConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
