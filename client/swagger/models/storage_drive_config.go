// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StorageDriveConfig storage drive config
//
// swagger:model storage.driveConfig
type StorageDriveConfig struct {

	// Set to allow files which return cannotDownloadAbusiveFile to be downloaded.
	AcknowledgeAbuse *bool `json:"acknowledgeAbuse,omitempty"`

	// Allow the filetype to change when uploading Google docs.
	AllowImportNameChange *bool `json:"allowImportNameChange,omitempty"`

	// Deprecated: No longer needed.
	AlternateExport *bool `json:"alternateExport,omitempty"`

	// Only consider files owned by the authenticated user.
	AuthOwnerOnly *bool `json:"authOwnerOnly,omitempty"`

	// Auth server URL.
	AuthURL string `json:"authUrl,omitempty"`

	// Upload chunk size.
	ChunkSize *string `json:"chunkSize,omitempty"`

	// Google Application Client Id
	ClientID string `json:"clientId,omitempty"`

	// OAuth Client Secret.
	ClientSecret string `json:"clientSecret,omitempty"`

	// Server side copy contents of shortcuts instead of the shortcut.
	CopyShortcutContent *bool `json:"copyShortcutContent,omitempty"`

	// Disable drive using http2.
	DisableHttp2 *bool `json:"disableHttp2,omitempty"`

	// The encoding for the backend.
	Encoding *string `json:"encoding,omitempty"`

	// Comma separated list of preferred formats for downloading Google docs.
	ExportFormats *string `json:"exportFormats,omitempty"`

	// Deprecated: See export_formats.
	Formats string `json:"formats,omitempty"`

	// Impersonate this user when using a service account.
	Impersonate string `json:"impersonate,omitempty"`

	// Comma separated list of preferred formats for uploading Google docs.
	ImportFormats string `json:"importFormats,omitempty"`

	// Keep new head revision of each file forever.
	KeepRevisionForever *bool `json:"keepRevisionForever,omitempty"`

	// Size of listing chunk 100-1000, 0 to disable.
	ListChunk *int64 `json:"listChunk,omitempty"`

	// Number of API calls to allow without sleeping.
	PacerBurst *int64 `json:"pacerBurst,omitempty"`

	// Minimum time to sleep between API calls.
	PacerMinSleep *string `json:"pacerMinSleep,omitempty"`

	// Resource key for accessing a link-shared file.
	ResourceKey string `json:"resourceKey,omitempty"`

	// ID of the root folder.
	RootFolderID string `json:"rootFolderId,omitempty"`

	// Scope that rclone should use when requesting access from drive.
	// Example: drive
	Scope string `json:"scope,omitempty"`

	// Allow server-side operations (e.g. copy) to work across different drive configs.
	ServerSideAcrossConfigs *bool `json:"serverSideAcrossConfigs,omitempty"`

	// Service Account Credentials JSON blob.
	ServiceAccountCredentials string `json:"serviceAccountCredentials,omitempty"`

	// Service Account Credentials JSON file path.
	ServiceAccountFile string `json:"serviceAccountFile,omitempty"`

	// Only show files that are shared with me.
	SharedWithMe *bool `json:"sharedWithMe,omitempty"`

	// Show sizes as storage quota usage, not actual size.
	SizeAsQuota *bool `json:"sizeAsQuota,omitempty"`

	// Skip MD5 checksum on Google photos and videos only.
	SkipChecksumGphotos *bool `json:"skipChecksumGphotos,omitempty"`

	// If set skip dangling shortcut files.
	SkipDanglingShortcuts *bool `json:"skipDanglingShortcuts,omitempty"`

	// Skip google documents in all listings.
	SkipGdocs *bool `json:"skipGdocs,omitempty"`

	// If set skip shortcut files.
	SkipShortcuts *bool `json:"skipShortcuts,omitempty"`

	// Only show files that are starred.
	StarredOnly *bool `json:"starredOnly,omitempty"`

	// Make download limit errors be fatal.
	StopOnDownloadLimit *bool `json:"stopOnDownloadLimit,omitempty"`

	// Make upload limit errors be fatal.
	StopOnUploadLimit *bool `json:"stopOnUploadLimit,omitempty"`

	// ID of the Shared Drive (Team Drive).
	TeamDrive string `json:"teamDrive,omitempty"`

	// OAuth Access Token as a JSON blob.
	Token string `json:"token,omitempty"`

	// Token server url.
	TokenURL string `json:"tokenUrl,omitempty"`

	// Only show files that are in the trash.
	TrashedOnly *bool `json:"trashedOnly,omitempty"`

	// Cutoff for switching to chunked upload.
	UploadCutoff *string `json:"uploadCutoff,omitempty"`

	// Use file created date instead of modified date.
	UseCreatedDate *bool `json:"useCreatedDate,omitempty"`

	// Use date file was shared instead of modified date.
	UseSharedDate *bool `json:"useSharedDate,omitempty"`

	// Send files to the trash instead of deleting permanently.
	UseTrash *bool `json:"useTrash,omitempty"`

	// If Object's are greater, use drive v2 API to download.
	V2DownloadMinSize *string `json:"v2DownloadMinSize,omitempty"`
}

// Validate validates this storage drive config
func (m *StorageDriveConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this storage drive config based on context it is used
func (m *StorageDriveConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StorageDriveConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageDriveConfig) UnmarshalBinary(b []byte) error {
	var res StorageDriveConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
