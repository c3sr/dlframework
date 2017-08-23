// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DlframeworkPredictURLRequest dlframework predict URL request
// swagger:model dlframeworkPredictURLRequest
type DlframeworkPredictURLRequest struct {

	// framework name
	FrameworkName string `json:"framework_name,omitempty"`

	// framework version
	FrameworkVersion string `json:"framework_version,omitempty"`

	// input id
	InputID string `json:"input_id,omitempty"`

	// limit
	Limit int32 `json:"limit,omitempty"`

	// model name
	ModelName string `json:"model_name,omitempty"`

	// model version
	ModelVersion string `json:"model_version,omitempty"`

	// request id
	RequestID string `json:"request_id,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this dlframework predict URL request
func (m *DlframeworkPredictURLRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DlframeworkPredictURLRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DlframeworkPredictURLRequest) UnmarshalBinary(b []byte) error {
	var res DlframeworkPredictURLRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
