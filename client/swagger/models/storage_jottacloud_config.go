// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StorageJottacloudConfig storage jottacloud config
//
// swagger:model storage.JottacloudConfig
type StorageJottacloudConfig struct {

	// The encoding for the backend.
	Encoding *string `json:"encoding,omitempty"`

	// Delete files permanently rather than putting them into the trash.
	HardDelete *bool `json:"hardDelete,omitempty"`

	// Files bigger than this will be cached on disk to calculate the MD5 if required.
	Md5MemoryLimit *string `json:"md5MemoryLimit,omitempty"`

	// Avoid server side versioning by deleting files and recreating files instead of overwriting them.
	NoVersions *bool `json:"noVersions,omitempty"`

	// Only show files that are in the trash.
	TrashedOnly *bool `json:"trashedOnly,omitempty"`

	// Files bigger than this can be resumed if the upload fail's.
	UploadResumeLimit *string `json:"uploadResumeLimit,omitempty"`
}

// Validate validates this storage jottacloud config
func (m *StorageJottacloudConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this storage jottacloud config based on context it is used
func (m *StorageJottacloudConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StorageJottacloudConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageJottacloudConfig) UnmarshalBinary(b []byte) error {
	var res StorageJottacloudConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}