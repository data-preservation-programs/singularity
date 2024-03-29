// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ModelFile model file
//
// swagger:model model.File
type ModelFile struct {

	// Associations
	AttachmentID int64 `json:"attachmentId,omitempty"`

	// CID is the CID of the file.
	Cid string `json:"cid,omitempty"`

	// directory Id
	DirectoryID int64 `json:"directoryId,omitempty"`

	// file ranges
	FileRanges []*ModelFileRange `json:"fileRanges"`

	// Hash is the hash of the file.
	Hash string `json:"hash,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// last modified nano
	LastModifiedNano int64 `json:"lastModifiedNano,omitempty"`

	// Path is the relative path to the file inside the storage.
	Path string `json:"path,omitempty"`

	// Size is the size of the file in bytes.
	Size int64 `json:"size,omitempty"`
}

// Validate validates this model file
func (m *ModelFile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFileRanges(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelFile) validateFileRanges(formats strfmt.Registry) error {
	if swag.IsZero(m.FileRanges) { // not required
		return nil
	}

	for i := 0; i < len(m.FileRanges); i++ {
		if swag.IsZero(m.FileRanges[i]) { // not required
			continue
		}

		if m.FileRanges[i] != nil {
			if err := m.FileRanges[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("fileRanges" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("fileRanges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this model file based on the context it is used
func (m *ModelFile) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFileRanges(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelFile) contextValidateFileRanges(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.FileRanges); i++ {

		if m.FileRanges[i] != nil {

			if swag.IsZero(m.FileRanges[i]) { // not required
				return nil
			}

			if err := m.FileRanges[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("fileRanges" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("fileRanges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelFile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelFile) UnmarshalBinary(b []byte) error {
	var res ModelFile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
