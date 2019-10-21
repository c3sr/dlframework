// Code generated by go-swagger; DO NOT EDIT.

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/rai-project/dlframework/httpapi/models"
)

// SignupReader is a Reader for the Signup structure.
type SignupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SignupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSignupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSignupOK creates a SignupOK with default headers values
func NewSignupOK() *SignupOK {
	return &SignupOK{}
}

/*SignupOK handles this case with default header values.

SignupOK signup o k
*/
type SignupOK struct {
	Payload *models.DlframeworkSignupResponse
}

func (o *SignupOK) Error() string {
	return fmt.Sprintf("[POST /auth/signup][%d] signupOK  %+v", 200, o.Payload)
}

func (o *SignupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DlframeworkSignupResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
