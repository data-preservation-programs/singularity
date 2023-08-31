// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StorageFtpConfig storage ftp config
//
// swagger:model storage.FtpConfig
type StorageFtpConfig struct {

	// Allow asking for FTP password when needed.
	AskPassword *bool `json:"askPassword,omitempty"`

	// Maximum time to wait for a response to close.
	CloseTimeout *string `json:"closeTimeout,omitempty"`

	// Maximum number of FTP simultaneous connections, 0 for unlimited.
	Concurrency int64 `json:"concurrency,omitempty"`

	// Disable using EPSV even if server advertises support.
	DisableEpsv *bool `json:"disableEpsv,omitempty"`

	// Disable using MLSD even if server advertises support.
	DisableMlsd *bool `json:"disableMlsd,omitempty"`

	// Disable TLS 1.3 (workaround for FTP servers with buggy TLS)
	DisableTls13 *bool `json:"disableTls13,omitempty"`

	// Disable using UTF-8 even if server advertises support.
	DisableUTF8 *bool `json:"disableUtf8,omitempty"`

	// The encoding for the backend.
	// Example: Asterisk,Ctl,Dot,Slash
	Encoding *string `json:"encoding,omitempty"`

	// Use Explicit FTPS (FTP over TLS).
	ExplicitTLS *bool `json:"explicitTls,omitempty"`

	// Use LIST -a to force listing of hidden files and folders. This will disable the use of MLSD.
	ForceListHidden *bool `json:"forceListHidden,omitempty"`

	// FTP host to connect to.
	Host string `json:"host,omitempty"`

	// Max time before closing idle connections.
	IdleTimeout *string `json:"idleTimeout,omitempty"`

	// Do not verify the TLS certificate of the server.
	NoCheckCertificate *bool `json:"noCheckCertificate,omitempty"`

	// FTP password.
	Pass string `json:"pass,omitempty"`

	// FTP port number.
	Port *int64 `json:"port,omitempty"`

	// Maximum time to wait for data connection closing status.
	ShutTimeout *string `json:"shutTimeout,omitempty"`

	// Use Implicit FTPS (FTP over TLS).
	TLS *bool `json:"tls,omitempty"`

	// Size of TLS session cache for all control and data connections.
	TLSCacheSize *int64 `json:"tlsCacheSize,omitempty"`

	// FTP username.
	User *string `json:"user,omitempty"`

	// Use MDTM to set modification time (VsFtpd quirk)
	WritingMdtm *bool `json:"writingMdtm,omitempty"`
}

// Validate validates this storage ftp config
func (m *StorageFtpConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this storage ftp config based on context it is used
func (m *StorageFtpConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StorageFtpConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageFtpConfig) UnmarshalBinary(b []byte) error {
	var res StorageFtpConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
