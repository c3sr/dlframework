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

// UserInfoReader is a Reader for the UserInfo structure.
type UserInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUserInfoOK creates a UserInfoOK with default headers values
func NewUserInfoOK() *UserInfoOK {
	return &UserInfoOK{}
}

/*UserInfoOK handles this case with default header values.

UserInfoOK user info o k
*/
type UserInfoOK struct {
	Payload *models.DlframeworkUserInfoResponse
}

func (o *UserInfoOK) Error() string {
	return fmt.Sprintf("[GET /auth/userinfo][%d] userInfoOK  %+v", 200, o.Payload)
}

func (o *UserInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DlframeworkUserInfoResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
