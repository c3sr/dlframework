// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// DlframeworkUserInfoResponse dlframework user info response
// swagger:model dlframeworkUserInfoResponse
type DlframeworkUserInfoResponse struct {

	// affiliation
	Affiliation string `json:"affiliation,omitempty"`

	// email
	Email string `json:"email,omitempty"`

	// first name
	FirstName string `json:"first_name,omitempty"`

	// last name
	LastName string `json:"last_name,omitempty"`

	// outcome
	Outcome string `json:"outcome,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this dlframework user info response
func (m *DlframeworkUserInfoResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DlframeworkUserInfoResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DlframeworkUserInfoResponse) UnmarshalBinary(b []byte) error {
	var res DlframeworkUserInfoResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
