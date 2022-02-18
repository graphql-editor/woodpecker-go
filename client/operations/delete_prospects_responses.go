// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/graphql-editor/woodpecker/models"
)

// DeleteProspectsReader is a Reader for the DeleteProspects structure.
type DeleteProspectsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteProspectsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteProspectsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteProspectsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteProspectsOK creates a DeleteProspectsOK with default headers values
func NewDeleteProspectsOK() *DeleteProspectsOK {
	return &DeleteProspectsOK{}
}

/* DeleteProspectsOK describes a response with status code 200, with default header values.

Empty response
*/
type DeleteProspectsOK struct {
}

func (o *DeleteProspectsOK) Error() string {
	return fmt.Sprintf("[DELETE /prospects][%d] deleteProspectsOK ", 200)
}

func (o *DeleteProspectsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProspectsDefault creates a DeleteProspectsDefault with default headers values
func NewDeleteProspectsDefault(code int) *DeleteProspectsDefault {
	return &DeleteProspectsDefault{
		_statusCode: code,
	}
}

/* DeleteProspectsDefault describes a response with status code -1, with default header values.

Unexpected error
*/
type DeleteProspectsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete prospects default response
func (o *DeleteProspectsDefault) Code() int {
	return o._statusCode
}

func (o *DeleteProspectsDefault) Error() string {
	return fmt.Sprintf("[DELETE /prospects][%d] DeleteProspects default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteProspectsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteProspectsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}