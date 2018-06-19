// Code generated by go-swagger; DO NOT EDIT.

package predict

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/rai-project/dlframework/httpapi/models"
)

// NewOpenParams creates a new OpenParams object
// with the default values initialized.
func NewOpenParams() *OpenParams {
	var ()
	return &OpenParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOpenParamsWithTimeout creates a new OpenParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOpenParamsWithTimeout(timeout time.Duration) *OpenParams {
	var ()
	return &OpenParams{

		timeout: timeout,
	}
}

// NewOpenParamsWithContext creates a new OpenParams object
// with the default values initialized, and the ability to set a context for a request
func NewOpenParamsWithContext(ctx context.Context) *OpenParams {
	var ()
	return &OpenParams{

		Context: ctx,
	}
}

// NewOpenParamsWithHTTPClient creates a new OpenParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewOpenParamsWithHTTPClient(client *http.Client) *OpenParams {
	var ()
	return &OpenParams{
		HTTPClient: client,
	}
}

/*OpenParams contains all the parameters to send to the API endpoint
for the open operation typically these are written to a http.Request
*/
type OpenParams struct {

	/*Body*/
	Body *models.DlframeworkPredictorOpenRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the open params
func (o *OpenParams) WithTimeout(timeout time.Duration) *OpenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the open params
func (o *OpenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the open params
func (o *OpenParams) WithContext(ctx context.Context) *OpenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the open params
func (o *OpenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the open params
func (o *OpenParams) WithHTTPClient(client *http.Client) *OpenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the open params
func (o *OpenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the open params
func (o *OpenParams) WithBody(body *models.DlframeworkPredictorOpenRequest) *OpenParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the open params
func (o *OpenParams) SetBody(body *models.DlframeworkPredictorOpenRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *OpenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
