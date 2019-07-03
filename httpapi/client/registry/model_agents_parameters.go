// Code generated by go-swagger; DO NOT EDIT.

package registry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewModelAgentsParams creates a new ModelAgentsParams object
// with the default values initialized.
func NewModelAgentsParams() *ModelAgentsParams {
	var ()
	return &ModelAgentsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewModelAgentsParamsWithTimeout creates a new ModelAgentsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewModelAgentsParamsWithTimeout(timeout time.Duration) *ModelAgentsParams {
	var ()
	return &ModelAgentsParams{

		timeout: timeout,
	}
}

// NewModelAgentsParamsWithContext creates a new ModelAgentsParams object
// with the default values initialized, and the ability to set a context for a request
func NewModelAgentsParamsWithContext(ctx context.Context) *ModelAgentsParams {
	var ()
	return &ModelAgentsParams{

		Context: ctx,
	}
}

// NewModelAgentsParamsWithHTTPClient creates a new ModelAgentsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewModelAgentsParamsWithHTTPClient(client *http.Client) *ModelAgentsParams {
	var ()
	return &ModelAgentsParams{
		HTTPClient: client,
	}
}

/*ModelAgentsParams contains all the parameters to send to the API endpoint
for the model agents operation typically these are written to a http.Request
*/
type ModelAgentsParams struct {

	/*FrameworkName*/
	FrameworkName *string
	/*FrameworkVersion*/
	FrameworkVersion *string
	/*ModelName*/
	ModelName *string
	/*ModelVersion*/
	ModelVersion *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the model agents params
func (o *ModelAgentsParams) WithTimeout(timeout time.Duration) *ModelAgentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the model agents params
func (o *ModelAgentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the model agents params
func (o *ModelAgentsParams) WithContext(ctx context.Context) *ModelAgentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the model agents params
func (o *ModelAgentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the model agents params
func (o *ModelAgentsParams) WithHTTPClient(client *http.Client) *ModelAgentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the model agents params
func (o *ModelAgentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFrameworkName adds the frameworkName to the model agents params
func (o *ModelAgentsParams) WithFrameworkName(frameworkName *string) *ModelAgentsParams {
	o.SetFrameworkName(frameworkName)
	return o
}

// SetFrameworkName adds the frameworkName to the model agents params
func (o *ModelAgentsParams) SetFrameworkName(frameworkName *string) {
	o.FrameworkName = frameworkName
}

// WithFrameworkVersion adds the frameworkVersion to the model agents params
func (o *ModelAgentsParams) WithFrameworkVersion(frameworkVersion *string) *ModelAgentsParams {
	o.SetFrameworkVersion(frameworkVersion)
	return o
}

// SetFrameworkVersion adds the frameworkVersion to the model agents params
func (o *ModelAgentsParams) SetFrameworkVersion(frameworkVersion *string) {
	o.FrameworkVersion = frameworkVersion
}

// WithModelName adds the modelName to the model agents params
func (o *ModelAgentsParams) WithModelName(modelName *string) *ModelAgentsParams {
	o.SetModelName(modelName)
	return o
}

// SetModelName adds the modelName to the model agents params
func (o *ModelAgentsParams) SetModelName(modelName *string) {
	o.ModelName = modelName
}

// WithModelVersion adds the modelVersion to the model agents params
func (o *ModelAgentsParams) WithModelVersion(modelVersion *string) *ModelAgentsParams {
	o.SetModelVersion(modelVersion)
	return o
}

// SetModelVersion adds the modelVersion to the model agents params
func (o *ModelAgentsParams) SetModelVersion(modelVersion *string) {
	o.ModelVersion = modelVersion
}

// WriteToRequest writes these params to a swagger request
func (o *ModelAgentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FrameworkName != nil {

		// query param framework_name
		var qrFrameworkName string
		if o.FrameworkName != nil {
			qrFrameworkName = *o.FrameworkName
		}
		qFrameworkName := qrFrameworkName
		if qFrameworkName != "" {
			if err := r.SetQueryParam("framework_name", qFrameworkName); err != nil {
				return err
			}
		}

	}

	if o.FrameworkVersion != nil {

		// query param framework_version
		var qrFrameworkVersion string
		if o.FrameworkVersion != nil {
			qrFrameworkVersion = *o.FrameworkVersion
		}
		qFrameworkVersion := qrFrameworkVersion
		if qFrameworkVersion != "" {
			if err := r.SetQueryParam("framework_version", qFrameworkVersion); err != nil {
				return err
			}
		}

	}

	if o.ModelName != nil {

		// query param model_name
		var qrModelName string
		if o.ModelName != nil {
			qrModelName = *o.ModelName
		}
		qModelName := qrModelName
		if qModelName != "" {
			if err := r.SetQueryParam("model_name", qModelName); err != nil {
				return err
			}
		}

	}

	if o.ModelVersion != nil {

		// query param model_version
		var qrModelVersion string
		if o.ModelVersion != nil {
			qrModelVersion = *o.ModelVersion
		}
		qModelVersion := qrModelVersion
		if qModelVersion != "" {
			if err := r.SetQueryParam("model_version", qModelVersion); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
