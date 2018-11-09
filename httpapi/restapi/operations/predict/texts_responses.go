// Code generated by go-swagger; DO NOT EDIT.

package predict

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/rai-project/dlframework/httpapi/models"
)

// TextsOKCode is the HTTP code returned for type TextsOK
const TextsOKCode int = 200

/*TextsOK A successful response.

swagger:response textsOK
*/
type TextsOK struct {

	/*
	  In: Body
	*/
	Payload *models.DlframeworkFeaturesResponse `json:"body,omitempty"`
}

// NewTextsOK creates TextsOK with default headers values
func NewTextsOK() *TextsOK {

	return &TextsOK{}
}

// WithPayload adds the payload to the texts o k response
func (o *TextsOK) WithPayload(payload *models.DlframeworkFeaturesResponse) *TextsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the texts o k response
func (o *TextsOK) SetPayload(payload *models.DlframeworkFeaturesResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TextsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
