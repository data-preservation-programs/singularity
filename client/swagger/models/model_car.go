// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ModelCar model car
//
// swagger:model model.Car
type ModelCar struct {

	// created at
	CreatedAt string `json:"createdAt,omitempty"`

	// dataset Id
	DatasetID int64 `json:"datasetId,omitempty"`

	// file path
	FilePath string `json:"filePath,omitempty"`

	// file size
	FileSize int64 `json:"fileSize,omitempty"`

	// header
	Header []int64 `json:"header"`

	// id
	ID int64 `json:"id,omitempty"`

	// pack job Id
	PackJobID int64 `json:"packJobId,omitempty"`

	// piece cid
	PieceCid ModelCID `json:"pieceCid,omitempty"`

	// piece size
	PieceSize int64 `json:"pieceSize,omitempty"`

	// root cid
	RootCid ModelCID `json:"rootCid,omitempty"`

	// source Id
	SourceID int64 `json:"sourceId,omitempty"`
}

// Validate validates this model car
func (m *ModelCar) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this model car based on context it is used
func (m *ModelCar) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ModelCar) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelCar) UnmarshalBinary(b []byte) error {
	var res ModelCar
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
