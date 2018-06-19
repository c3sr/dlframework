// Code generated by go-swagger; DO NOT EDIT.

package predict

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/rai-project/dlframework/httpapi/models"
)

// ResetOKCode is the HTTP code returned for type ResetOK
const ResetOKCode int = 200

/*ResetOK reset o k

swagger:response resetOK
*/
type ResetOK struct {

	/*
	  In: Body
	*/
	Payload *models.DlframeworkResetResponse `json:"body,omitempty"`
}

// NewResetOK creates ResetOK with default headers values
func NewResetOK() *ResetOK {

	return &ResetOK{}
}

// WithPayload adds the payload to the reset o k response
func (o *ResetOK) WithPayload(payload *models.DlframeworkResetResponse) *ResetOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the reset o k response
func (o *ResetOK) SetPayload(payload *models.DlframeworkResetResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ResetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
