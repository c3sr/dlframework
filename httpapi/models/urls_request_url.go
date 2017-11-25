// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// UrlsRequestURL urls request URL
// swagger:model URLsRequestURL
type UrlsRequestURL struct {

	// data
	Data string `json:"data,omitempty"`

	// An id used to identify the output feature: maps to input_id for output
	ID string `json:"id,omitempty"`
}

// Validate validates this urls request URL
func (m *UrlsRequestURL) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *UrlsRequestURL) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UrlsRequestURL) UnmarshalBinary(b []byte) error {
	var res UrlsRequestURL
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
