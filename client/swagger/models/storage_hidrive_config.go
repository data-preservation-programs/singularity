// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StorageHidriveConfig storage hidrive config
//
// swagger:model storage.hidriveConfig
type StorageHidriveConfig struct {

	// Auth server URL.
	AuthURL string `json:"authUrl,omitempty"`

	// Chunksize for chunked uploads.
	ChunkSize *string `json:"chunkSize,omitempty"`

	// OAuth Client Id.
	ClientID string `json:"clientId,omitempty"`

	// OAuth Client Secret.
	ClientSecret string `json:"clientSecret,omitempty"`

	// Do not fetch number of objects in directories unless it is absolutely necessary.
	DisableFetchingMemberCount *bool `json:"disableFetchingMemberCount,omitempty"`

	// The encoding for the backend.
	Encoding *string `json:"encoding,omitempty"`

	// Endpoint for the service.
	Endpoint *string `json:"endpoint,omitempty"`

	// The root/parent folder for all paths.
	// Example: /
	RootPrefix *string `json:"rootPrefix,omitempty"`

	// Access permissions that rclone should use when requesting access from HiDrive.
	// Example: rw
	ScopeAccess *string `json:"scopeAccess,omitempty"`

	// User-level that rclone should use when requesting access from HiDrive.
	// Example: user
	ScopeRole *string `json:"scopeRole,omitempty"`

	// OAuth Access Token as a JSON blob.
	Token string `json:"token,omitempty"`

	// Token server url.
	TokenURL string `json:"tokenUrl,omitempty"`

	// Concurrency for chunked uploads.
	UploadConcurrency *int64 `json:"uploadConcurrency,omitempty"`

	// Cutoff/Threshold for chunked uploads.
	UploadCutoff *string `json:"uploadCutoff,omitempty"`
}

// Validate validates this storage hidrive config
func (m *StorageHidriveConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this storage hidrive config based on context it is used
func (m *StorageHidriveConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StorageHidriveConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageHidriveConfig) UnmarshalBinary(b []byte) error {
	var res StorageHidriveConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
