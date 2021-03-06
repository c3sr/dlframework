// Code generated by go-swagger; DO NOT EDIT.

package registry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewModelManifestsParams creates a new ModelManifestsParams object
// no default values defined in spec.
func NewModelManifestsParams() ModelManifestsParams {

	return ModelManifestsParams{}
}

// ModelManifestsParams contains all the bound params for the model manifests operation
// typically these are obtained from a http.Request
//
// swagger:parameters ModelManifests
type ModelManifestsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: query
	*/
	FrameworkName *string
	/*
	  In: query
	*/
	FrameworkVersion *string
	/*
	  In: query
	*/
	ModelName *string
	/*
	  In: query
	*/
	ModelVersion *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewModelManifestsParams() beforehand.
func (o *ModelManifestsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qFrameworkName, qhkFrameworkName, _ := qs.GetOK("framework_name")
	if err := o.bindFrameworkName(qFrameworkName, qhkFrameworkName, route.Formats); err != nil {
		res = append(res, err)
	}

	qFrameworkVersion, qhkFrameworkVersion, _ := qs.GetOK("framework_version")
	if err := o.bindFrameworkVersion(qFrameworkVersion, qhkFrameworkVersion, route.Formats); err != nil {
		res = append(res, err)
	}

	qModelName, qhkModelName, _ := qs.GetOK("model_name")
	if err := o.bindModelName(qModelName, qhkModelName, route.Formats); err != nil {
		res = append(res, err)
	}

	qModelVersion, qhkModelVersion, _ := qs.GetOK("model_version")
	if err := o.bindModelVersion(qModelVersion, qhkModelVersion, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindFrameworkName binds and validates parameter FrameworkName from query.
func (o *ModelManifestsParams) bindFrameworkName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.FrameworkName = &raw

	return nil
}

// bindFrameworkVersion binds and validates parameter FrameworkVersion from query.
func (o *ModelManifestsParams) bindFrameworkVersion(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.FrameworkVersion = &raw

	return nil
}

// bindModelName binds and validates parameter ModelName from query.
func (o *ModelManifestsParams) bindModelName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.ModelName = &raw

	return nil
}

// bindModelVersion binds and validates parameter ModelVersion from query.
func (o *ModelManifestsParams) bindModelVersion(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.ModelVersion = &raw

	return nil
}
