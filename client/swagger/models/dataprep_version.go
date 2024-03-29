// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DataprepVersion dataprep version
//
// swagger:model dataprep.Version
type DataprepVersion struct {

	// cid
	Cid string `json:"cid,omitempty"`

	// hash
	Hash string `json:"hash,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// last modified
	LastModified string `json:"lastModified,omitempty"`

	// size
	Size int64 `json:"size,omitempty"`
}

// Validate validates this dataprep version
func (m *DataprepVersion) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this dataprep version based on context it is used
func (m *DataprepVersion) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DataprepVersion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataprepVersion) UnmarshalBinary(b []byte) error {
	var res DataprepVersion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
