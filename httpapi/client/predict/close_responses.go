// Code generated by go-swagger; DO NOT EDIT.

package predict

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/rai-project/dlframework/httpapi/models"
)

// CloseReader is a Reader for the Close structure.
type CloseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CloseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCloseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCloseOK creates a CloseOK with default headers values
func NewCloseOK() *CloseOK {
	return &CloseOK{}
}

/*CloseOK handles this case with default header values.

A successful response.
*/
type CloseOK struct {
	Payload *models.DlframeworkPredictorCloseResponse
}

func (o *CloseOK) Error() string {
	return fmt.Sprintf("[POST /predict/close][%d] closeOK  %+v", 200, o.Payload)
}

func (o *CloseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DlframeworkPredictorCloseResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
