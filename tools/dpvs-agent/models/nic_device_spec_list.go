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

// NicDeviceSpecList nic device spec list
//
// swagger:model NicDeviceSpecList
type NicDeviceSpecList struct {

	// items
	Items []*NicDeviceSpec `json:"Items"`
}

// Validate validates this nic device spec list
func (m *NicDeviceSpecList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NicDeviceSpecList) validateItems(formats strfmt.Registry) error {
	if swag.IsZero(m.Items) { // not required
		return nil
	}

	for i := 0; i < len(m.Items); i++ {
		if swag.IsZero(m.Items[i]) { // not required
			continue
		}

		if m.Items[i] != nil {
			if err := m.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this nic device spec list based on the context it is used
func (m *NicDeviceSpecList) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NicDeviceSpecList) contextValidateItems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Items); i++ {

		if m.Items[i] != nil {
			if err := m.Items[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NicDeviceSpecList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NicDeviceSpecList) UnmarshalBinary(b []byte) error {
	var res NicDeviceSpecList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
