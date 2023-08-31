// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StorageOosUserPrincipalAuthConfig storage oos user principal auth config
//
// swagger:model storage.OosUser_principal_authConfig
type StorageOosUserPrincipalAuthConfig struct {

	// Chunk size to use for uploading.
	ChunkSize *string `json:"chunkSize,omitempty"`

	// Object storage compartment OCID
	Compartment string `json:"compartment,omitempty"`

	// Path to OCI config file
	// Example: ~/.oci/config
	ConfigFile *string `json:"configFile,omitempty"`

	// Profile name inside the oci config file
	// Example: Default
	ConfigProfile *string `json:"configProfile,omitempty"`

	// Cutoff for switching to multipart copy.
	CopyCutoff *string `json:"copyCutoff,omitempty"`

	// Timeout for copy.
	CopyTimeout *string `json:"copyTimeout,omitempty"`

	// Don't store MD5 checksum with object metadata.
	DisableChecksum *bool `json:"disableChecksum,omitempty"`

	// The encoding for the backend.
	Encoding *string `json:"encoding,omitempty"`

	// Endpoint for Object storage API.
	Endpoint string `json:"endpoint,omitempty"`

	// If true avoid calling abort upload on a failure, leaving all successfully uploaded parts on S3 for manual recovery.
	LeavePartsOnError *bool `json:"leavePartsOnError,omitempty"`

	// Object storage namespace
	Namespace string `json:"namespace,omitempty"`

	// If set, don't attempt to check the bucket exists or create it.
	NoCheckBucket *bool `json:"noCheckBucket,omitempty"`

	// Object storage Region
	Region string `json:"region,omitempty"`

	// If using SSE-C, the optional header that specifies "AES256" as the encryption algorithm.
	SseCustomerAlgorithm string `json:"sseCustomerAlgorithm,omitempty"`

	// To use SSE-C, the optional header that specifies the base64-encoded 256-bit encryption key to use to
	SseCustomerKey string `json:"sseCustomerKey,omitempty"`

	// To use SSE-C, a file containing the base64-encoded string of the AES-256 encryption key associated
	SseCustomerKeyFile string `json:"sseCustomerKeyFile,omitempty"`

	// If using SSE-C, The optional header that specifies the base64-encoded SHA256 hash of the encryption
	SseCustomerKeySha256 string `json:"sseCustomerKeySha256,omitempty"`

	// if using using your own master key in vault, this header specifies the
	SseKmsKeyID string `json:"sseKmsKeyId,omitempty"`

	// The storage class to use when storing new objects in storage. https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm
	// Example: Standard
	StorageTier *string `json:"storageTier,omitempty"`

	// Concurrency for multipart uploads.
	UploadConcurrency *int64 `json:"uploadConcurrency,omitempty"`

	// Cutoff for switching to chunked upload.
	UploadCutoff *string `json:"uploadCutoff,omitempty"`
}

// Validate validates this storage oos user principal auth config
func (m *StorageOosUserPrincipalAuthConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this storage oos user principal auth config based on context it is used
func (m *StorageOosUserPrincipalAuthConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StorageOosUserPrincipalAuthConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageOosUserPrincipalAuthConfig) UnmarshalBinary(b []byte) error {
	var res StorageOosUserPrincipalAuthConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
