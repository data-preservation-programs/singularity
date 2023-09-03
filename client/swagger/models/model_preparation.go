// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ModelPreparation model preparation
//
// swagger:model model.Preparation
type ModelPreparation struct {

	// created at
	CreatedAt string `json:"createdAt,omitempty"`

	// DeleteAfterExport is a flag that indicates whether the source files should be deleted after export.
	DeleteAfterExport bool `json:"deleteAfterExport,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// max size
	MaxSize int64 `json:"maxSize,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// piece size
	PieceSize int64 `json:"pieceSize,omitempty"`

	// updated at
	UpdatedAt string `json:"updatedAt,omitempty"`
}

// Validate validates this model preparation
func (m *ModelPreparation) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this model preparation based on context it is used
func (m *ModelPreparation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ModelPreparation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelPreparation) UnmarshalBinary(b []byte) error {
	var res ModelPreparation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
