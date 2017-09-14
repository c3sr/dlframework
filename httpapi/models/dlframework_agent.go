// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DlframeworkAgent dlframework agent
// swagger:model dlframeworkAgent

type DlframeworkAgent struct {

	// host
	Host string `json:"host,omitempty"`

	// port
	Port string `json:"port,omitempty"`

	// specification
	Specification string `json:"specification,omitempty"`
}

/* polymorph dlframeworkAgent host false */

/* polymorph dlframeworkAgent port false */

/* polymorph dlframeworkAgent specification false */

// Validate validates this dlframework agent
func (m *DlframeworkAgent) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DlframeworkAgent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DlframeworkAgent) UnmarshalBinary(b []byte) error {
	var res DlframeworkAgent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
