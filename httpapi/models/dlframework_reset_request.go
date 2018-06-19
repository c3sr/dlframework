// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DlframeworkResetRequest dlframework reset request
// swagger:model dlframeworkResetRequest
type DlframeworkResetRequest struct {

	// id
	ID string `json:"id,omitempty"`

	// predictor
	Predictor *DlframeworkPredictor `json:"predictor,omitempty"`
}

// Validate validates this dlframework reset request
func (m *DlframeworkResetRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePredictor(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DlframeworkResetRequest) validatePredictor(formats strfmt.Registry) error {

	if swag.IsZero(m.Predictor) { // not required
		return nil
	}

	if m.Predictor != nil {
		if err := m.Predictor.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("predictor")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DlframeworkResetRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DlframeworkResetRequest) UnmarshalBinary(b []byte) error {
	var res DlframeworkResetRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
