// Code generated by go-swagger; DO NOT EDIT.

package predict

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/c3sr/dlframework/httpapi/models"
)

// DatasetOKCode is the HTTP code returned for type DatasetOK
const DatasetOKCode int = 200

/*DatasetOK A successful response.

swagger:response datasetOK
*/
type DatasetOK struct {

	/*
	  In: Body
	*/
	Payload *models.DlframeworkFeaturesResponse `json:"body,omitempty"`
}

// NewDatasetOK creates DatasetOK with default headers values
func NewDatasetOK() *DatasetOK {

	return &DatasetOK{}
}

// WithPayload adds the payload to the dataset o k response
func (o *DatasetOK) WithPayload(payload *models.DlframeworkFeaturesResponse) *DatasetOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the dataset o k response
func (o *DatasetOK) SetPayload(payload *models.DlframeworkFeaturesResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DatasetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
