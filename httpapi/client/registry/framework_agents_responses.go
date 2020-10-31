// Code generated by go-swagger; DO NOT EDIT.

package registry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/c3sr/dlframework/httpapi/models"
)

// FrameworkAgentsReader is a Reader for the FrameworkAgents structure.
type FrameworkAgentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FrameworkAgentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFrameworkAgentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewFrameworkAgentsOK creates a FrameworkAgentsOK with default headers values
func NewFrameworkAgentsOK() *FrameworkAgentsOK {
	return &FrameworkAgentsOK{}
}

/*FrameworkAgentsOK handles this case with default header values.

A successful response.
*/
type FrameworkAgentsOK struct {
	Payload *models.DlframeworkAgents
}

func (o *FrameworkAgentsOK) Error() string {
	return fmt.Sprintf("[GET /registry/frameworks/agent][%d] frameworkAgentsOK  %+v", 200, o.Payload)
}

func (o *FrameworkAgentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DlframeworkAgents)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
