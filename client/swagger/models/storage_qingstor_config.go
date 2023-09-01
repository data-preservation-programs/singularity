// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StorageQingstorConfig storage qingstor config
//
// swagger:model storage.QingstorConfig
type StorageQingstorConfig struct {

	// QingStor Access Key ID.
	AccessKeyID string `json:"accessKeyId,omitempty"`

	// Chunk size to use for uploading.
	ChunkSize *string `json:"chunkSize,omitempty"`

	// Number of connection retries.
	ConnectionRetries *int64 `json:"connectionRetries,omitempty"`

	// The encoding for the backend.
	Encoding *string `json:"encoding,omitempty"`

	// Enter an endpoint URL to connection QingStor API.
	Endpoint string `json:"endpoint,omitempty"`

	// Get QingStor credentials from runtime.
	// Example: false
	EnvAuth *bool `json:"envAuth,omitempty"`

	// QingStor Secret Access Key (password).
	SecretAccessKey string `json:"secretAccessKey,omitempty"`

	// Concurrency for multipart uploads.
	UploadConcurrency *int64 `json:"uploadConcurrency,omitempty"`

	// Cutoff for switching to chunked upload.
	UploadCutoff *string `json:"uploadCutoff,omitempty"`

	// Zone to connect to.
	// Example: pek3a
	Zone string `json:"zone,omitempty"`
}

// Validate validates this storage qingstor config
func (m *StorageQingstorConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this storage qingstor config based on context it is used
func (m *StorageQingstorConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StorageQingstorConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageQingstorConfig) UnmarshalBinary(b []byte) error {
	var res StorageQingstorConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}