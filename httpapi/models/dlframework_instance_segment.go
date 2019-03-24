// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DlframeworkInstanceSegment dlframework instance segment
// swagger:model dlframeworkInstanceSegment
type DlframeworkInstanceSegment struct {

	// data
	// Format: byte
	Data strfmt.Base64 `json:"data,omitempty"`

	// height
	Height int32 `json:"height,omitempty"`

	// index
	Index int32 `json:"index,omitempty"`

	// label
	Label string `json:"label,omitempty"`

	// width
	Width int32 `json:"width,omitempty"`

	// xmax
	Xmax float32 `json:"xmax,omitempty"`

	// xmin
	Xmin float32 `json:"xmin,omitempty"`

	// ymax
	Ymax float32 `json:"ymax,omitempty"`

	// ymin
	Ymin float32 `json:"ymin,omitempty"`
}

// Validate validates this dlframework instance segment
func (m *DlframeworkInstanceSegment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DlframeworkInstanceSegment) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	// Format "byte" (base64 string) is already validated when unmarshalled

	return nil
}

// MarshalBinary interface implementation
func (m *DlframeworkInstanceSegment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DlframeworkInstanceSegment) UnmarshalBinary(b []byte) error {
	var res DlframeworkInstanceSegment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
