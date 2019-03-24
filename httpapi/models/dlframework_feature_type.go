// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// DlframeworkFeatureType dlframework feature type
// swagger:model dlframeworkFeatureType
type DlframeworkFeatureType string

const (

	// DlframeworkFeatureTypeUNKNOWN captures enum value "UNKNOWN"
	DlframeworkFeatureTypeUNKNOWN DlframeworkFeatureType = "UNKNOWN"

	// DlframeworkFeatureTypeIMAGE captures enum value "IMAGE"
	DlframeworkFeatureTypeIMAGE DlframeworkFeatureType = "IMAGE"

	// DlframeworkFeatureTypeCLASSIFICATION captures enum value "CLASSIFICATION"
	DlframeworkFeatureTypeCLASSIFICATION DlframeworkFeatureType = "CLASSIFICATION"

	// DlframeworkFeatureTypeBOUNDINGBOX captures enum value "BOUNDINGBOX"
	DlframeworkFeatureTypeBOUNDINGBOX DlframeworkFeatureType = "BOUNDINGBOX"

	// DlframeworkFeatureTypeSEGMENT captures enum value "SEGMENT"
	DlframeworkFeatureTypeSEGMENT DlframeworkFeatureType = "SEGMENT"

	// DlframeworkFeatureTypeINSTANCESEGMENT captures enum value "INSTANCESEGMENT"
	DlframeworkFeatureTypeINSTANCESEGMENT DlframeworkFeatureType = "INSTANCESEGMENT"

	// DlframeworkFeatureTypeGEOLOCATION captures enum value "GEOLOCATION"
	DlframeworkFeatureTypeGEOLOCATION DlframeworkFeatureType = "GEOLOCATION"

	// DlframeworkFeatureTypeREGION captures enum value "REGION"
	DlframeworkFeatureTypeREGION DlframeworkFeatureType = "REGION"

	// DlframeworkFeatureTypeTEXT captures enum value "TEXT"
	DlframeworkFeatureTypeTEXT DlframeworkFeatureType = "TEXT"

	// DlframeworkFeatureTypeAUDIO captures enum value "AUDIO"
	DlframeworkFeatureTypeAUDIO DlframeworkFeatureType = "AUDIO"

	// DlframeworkFeatureTypeRAW captures enum value "RAW"
	DlframeworkFeatureTypeRAW DlframeworkFeatureType = "RAW"
)

// for schema
var dlframeworkFeatureTypeEnum []interface{}

func init() {
	var res []DlframeworkFeatureType
	if err := json.Unmarshal([]byte(`["UNKNOWN","IMAGE","CLASSIFICATION","BOUNDINGBOX","SEGMENT","INSTANCESEGMENT","GEOLOCATION","REGION","TEXT","AUDIO","RAW"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		dlframeworkFeatureTypeEnum = append(dlframeworkFeatureTypeEnum, v)
	}
}

func (m DlframeworkFeatureType) validateDlframeworkFeatureTypeEnum(path, location string, value DlframeworkFeatureType) error {
	if err := validate.Enum(path, location, value, dlframeworkFeatureTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this dlframework feature type
func (m DlframeworkFeatureType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDlframeworkFeatureTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
