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

// DealListDealRequest deal list deal request
//
// swagger:model deal.ListDealRequest
type DealListDealRequest struct {

	// preparation ID or name filter
	Preparations []string `json:"preparations"`

	// provider filter
	Providers []string `json:"providers"`

	// schedule id filter
	Schedules []int64 `json:"schedules"`

	// source ID or name filter
	Sources []string `json:"sources"`

	// state filter
	States []ModelDealState `json:"states"`
}

// Validate validates this deal list deal request
func (m *DealListDealRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStates(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DealListDealRequest) validateStates(formats strfmt.Registry) error {
	if swag.IsZero(m.States) { // not required
		return nil
	}

	for i := 0; i < len(m.States); i++ {

		if err := m.States[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("states" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("states" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// ContextValidate validate this deal list deal request based on the context it is used
func (m *DealListDealRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DealListDealRequest) contextValidateStates(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.States); i++ {

		if swag.IsZero(m.States[i]) { // not required
			return nil
		}

		if err := m.States[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("states" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("states" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DealListDealRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DealListDealRequest) UnmarshalBinary(b []byte) error {
	var res DealListDealRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
