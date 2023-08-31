// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// FsDuration fs duration
//
// swagger:model fs.Duration
type FsDuration int64

// for schema
var fsDurationEnum []interface{}

func init() {
	var res []FsDuration
	if err := json.Unmarshal([]byte(`[-9223372036854776000,9223372036854776000,1,1000,1000000,1000000000,60000000000,3600000000000,3153600000000000000]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fsDurationEnum = append(fsDurationEnum, v)
	}
}

func (m FsDuration) validateFsDurationEnum(path, location string, value FsDuration) error {
	if err := validate.EnumCase(path, location, value, fsDurationEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this fs duration
func (m FsDuration) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateFsDurationEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this fs duration based on context it is used
func (m FsDuration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
