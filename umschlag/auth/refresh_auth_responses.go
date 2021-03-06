// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/umschlag/umschlag-go/models"
)

// RefreshAuthReader is a Reader for the RefreshAuth structure.
type RefreshAuthReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RefreshAuthReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRefreshAuthOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewRefreshAuthUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewRefreshAuthDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRefreshAuthOK creates a RefreshAuthOK with default headers values
func NewRefreshAuthOK() *RefreshAuthOK {
	return &RefreshAuthOK{}
}

/*RefreshAuthOK handles this case with default header values.

A refreshed token with expire
*/
type RefreshAuthOK struct {
	Payload *models.AuthToken
}

func (o *RefreshAuthOK) Error() string {
	return fmt.Sprintf("[GET /auth/refresh][%d] refreshAuthOK  %+v", 200, o.Payload)
}

func (o *RefreshAuthOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthToken)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRefreshAuthUnauthorized creates a RefreshAuthUnauthorized with default headers values
func NewRefreshAuthUnauthorized() *RefreshAuthUnauthorized {
	return &RefreshAuthUnauthorized{}
}

/*RefreshAuthUnauthorized handles this case with default header values.

Unauthorized if failed to generate
*/
type RefreshAuthUnauthorized struct {
	Payload *models.GeneralError
}

func (o *RefreshAuthUnauthorized) Error() string {
	return fmt.Sprintf("[GET /auth/refresh][%d] refreshAuthUnauthorized  %+v", 401, o.Payload)
}

func (o *RefreshAuthUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRefreshAuthDefault creates a RefreshAuthDefault with default headers values
func NewRefreshAuthDefault(code int) *RefreshAuthDefault {
	return &RefreshAuthDefault{
		_statusCode: code,
	}
}

/*RefreshAuthDefault handles this case with default header values.

Some error unrelated to the handler
*/
type RefreshAuthDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// Code gets the status code for the refresh auth default response
func (o *RefreshAuthDefault) Code() int {
	return o._statusCode
}

func (o *RefreshAuthDefault) Error() string {
	return fmt.Sprintf("[GET /auth/refresh][%d] RefreshAuth default  %+v", o._statusCode, o.Payload)
}

func (o *RefreshAuthDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
