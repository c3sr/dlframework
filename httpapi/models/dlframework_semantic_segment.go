// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// DlframeworkSemanticSegment dlframework semantic segment
// swagger:model dlframeworkSemanticSegment
type DlframeworkSemanticSegment struct {

	// height
	Height int32 `json:"height,omitempty"`

	// int mask
	IntMask []int32 `json:"int_mask"`

	// width
	Width int32 `json:"width,omitempty"`
}

// Validate validates this dlframework semantic segment
func (m *DlframeworkSemanticSegment) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DlframeworkSemanticSegment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DlframeworkSemanticSegment) UnmarshalBinary(b []byte) error {
	var res DlframeworkSemanticSegment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
