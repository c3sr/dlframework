// Code generated by go-swagger; DO NOT EDIT.

package registry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/rai-project/dlframework/httpapi/models"
)

// FrameworkManifestsReader is a Reader for the FrameworkManifests structure.
type FrameworkManifestsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FrameworkManifestsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFrameworkManifestsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewFrameworkManifestsOK creates a FrameworkManifestsOK with default headers values
func NewFrameworkManifestsOK() *FrameworkManifestsOK {
	return &FrameworkManifestsOK{}
}

/*FrameworkManifestsOK handles this case with default header values.

A successful response.
*/
type FrameworkManifestsOK struct {
	Payload *models.DlframeworkFrameworkManifestsResponse
}

func (o *FrameworkManifestsOK) Error() string {
	return fmt.Sprintf("[GET /registry/frameworks/manifest][%d] frameworkManifestsOK  %+v", 200, o.Payload)
}

func (o *FrameworkManifestsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DlframeworkFrameworkManifestsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
