// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DlframeworkFeature dlframework feature
// swagger:model dlframeworkFeature
type DlframeworkFeature struct {

	// audio
	Audio *DlframeworkAudio `json:"audio,omitempty"`

	// bounding box
	BoundingBox *DlframeworkBoundingBox `json:"bounding_box,omitempty"`

	// classification
	Classification *DlframeworkClassification `json:"classification,omitempty"`

	// geolocation
	Geolocation *DlframeworkGeoLocation `json:"geolocation,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// image
	Image *DlframeworkImage `json:"image,omitempty"`

	// metadata
	Metadata map[string]string `json:"metadata,omitempty"`

	// probability
	Probability float32 `json:"probability,omitempty"`

	// raw
	Raw *DlframeworkRaw `json:"raw,omitempty"`

	// region
	Region *DlframeworkRegion `json:"region,omitempty"`

	// text
	Text *DlframeworkText `json:"text,omitempty"`

	// type
	Type DlframeworkFeatureType `json:"type,omitempty"`
}

// Validate validates this dlframework feature
func (m *DlframeworkFeature) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAudio(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBoundingBox(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClassification(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGeolocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRaw(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateText(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DlframeworkFeature) validateAudio(formats strfmt.Registry) error {

	if swag.IsZero(m.Audio) { // not required
		return nil
	}

	if m.Audio != nil {
		if err := m.Audio.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("audio")
			}
			return err
		}
	}

	return nil
}

func (m *DlframeworkFeature) validateBoundingBox(formats strfmt.Registry) error {

	if swag.IsZero(m.BoundingBox) { // not required
		return nil
	}

	if m.BoundingBox != nil {
		if err := m.BoundingBox.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bounding_box")
			}
			return err
		}
	}

	return nil
}

func (m *DlframeworkFeature) validateClassification(formats strfmt.Registry) error {

	if swag.IsZero(m.Classification) { // not required
		return nil
	}

	if m.Classification != nil {
		if err := m.Classification.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("classification")
			}
			return err
		}
	}

	return nil
}

func (m *DlframeworkFeature) validateGeolocation(formats strfmt.Registry) error {

	if swag.IsZero(m.Geolocation) { // not required
		return nil
	}

	if m.Geolocation != nil {
		if err := m.Geolocation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("geolocation")
			}
			return err
		}
	}

	return nil
}

func (m *DlframeworkFeature) validateImage(formats strfmt.Registry) error {

	if swag.IsZero(m.Image) { // not required
		return nil
	}

	if m.Image != nil {
		if err := m.Image.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("image")
			}
			return err
		}
	}

	return nil
}

func (m *DlframeworkFeature) validateRaw(formats strfmt.Registry) error {

	if swag.IsZero(m.Raw) { // not required
		return nil
	}

	if m.Raw != nil {
		if err := m.Raw.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("raw")
			}
			return err
		}
	}

	return nil
}

func (m *DlframeworkFeature) validateRegion(formats strfmt.Registry) error {

	if swag.IsZero(m.Region) { // not required
		return nil
	}

	if m.Region != nil {
		if err := m.Region.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("region")
			}
			return err
		}
	}

	return nil
}

func (m *DlframeworkFeature) validateText(formats strfmt.Registry) error {

	if swag.IsZero(m.Text) { // not required
		return nil
	}

	if m.Text != nil {
		if err := m.Text.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("text")
			}
			return err
		}
	}

	return nil
}

func (m *DlframeworkFeature) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DlframeworkFeature) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DlframeworkFeature) UnmarshalBinary(b []byte) error {
	var res DlframeworkFeature
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
