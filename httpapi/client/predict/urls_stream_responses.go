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

// UrlsStreamReader is a Reader for the UrlsStream structure.
type UrlsStreamReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UrlsStreamReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUrlsStreamOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUrlsStreamOK creates a UrlsStreamOK with default headers values
func NewUrlsStreamOK() *UrlsStreamOK {
	return &UrlsStreamOK{}
}

/*UrlsStreamOK handles this case with default header values.

(streaming responses)
*/
type UrlsStreamOK struct {
	Payload *models.DlframeworkFeatureResponse
}

func (o *UrlsStreamOK) Error() string {
	return fmt.Sprintf("[POST /predict/stream/urls][%d] urlsStreamOK  %+v", 200, o.Payload)
}

func (o *UrlsStreamOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DlframeworkFeatureResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
