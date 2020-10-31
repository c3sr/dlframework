// Code generated by go-swagger; DO NOT EDIT.

package predict

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	time "time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/c3sr/dlframework/httpapi/models"
)

// NewURLsParams creates a new URLsParams object
// with the default values initialized.
func NewURLsParams() *URLsParams {
	var ()
	return &URLsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewURLsParamsWithTimeout creates a new URLsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewURLsParamsWithTimeout(timeout time.Duration) *URLsParams {
	var ()
	return &URLsParams{

		timeout: timeout,
	}
}

// NewURLsParamsWithContext creates a new URLsParams object
// with the default values initialized, and the ability to set a context for a request
func NewURLsParamsWithContext(ctx context.Context) *URLsParams {
	var ()
	return &URLsParams{

		Context: ctx,
	}
}

// NewURLsParamsWithHTTPClient creates a new URLsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewURLsParamsWithHTTPClient(client *http.Client) *URLsParams {
	var ()
	return &URLsParams{
		HTTPClient: client,
	}
}

/*URLsParams contains all the parameters to send to the API endpoint
for the u r ls operation typically these are written to a http.Request
*/
type URLsParams struct {

	/*Body*/
	Body *models.DlframeworkURLsRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the u r ls params
func (o *URLsParams) WithTimeout(timeout time.Duration) *URLsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the u r ls params
func (o *URLsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the u r ls params
func (o *URLsParams) WithContext(ctx context.Context) *URLsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the u r ls params
func (o *URLsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the u r ls params
func (o *URLsParams) WithHTTPClient(client *http.Client) *URLsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the u r ls params
func (o *URLsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the u r ls params
func (o *URLsParams) WithBody(body *models.DlframeworkURLsRequest) *URLsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the u r ls params
func (o *URLsParams) SetBody(body *models.DlframeworkURLsRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *URLsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
