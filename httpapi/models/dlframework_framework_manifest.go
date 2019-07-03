// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DlframeworkFrameworkManifest dlframework framework manifest
// swagger:model dlframeworkFrameworkManifest
type DlframeworkFrameworkManifest struct {

	// container
	Container map[string]DlframeworkContainerHardware `json:"container,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this dlframework framework manifest
func (m *DlframeworkFrameworkManifest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContainer(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DlframeworkFrameworkManifest) validateContainer(formats strfmt.Registry) error {

	if swag.IsZero(m.Container) { // not required
		return nil
	}

	for k := range m.Container {

		if err := validate.Required("container"+"."+k, "body", m.Container[k]); err != nil {
			return err
		}
		if val, ok := m.Container[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DlframeworkFrameworkManifest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DlframeworkFrameworkManifest) UnmarshalBinary(b []byte) error {
	var res DlframeworkFrameworkManifest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
