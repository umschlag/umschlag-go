// Code generated by go-swagger; DO NOT EDIT.

package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/umschlag/umschlag-go/models"
)

// DelteTeamFromUserReader is a Reader for the DelteTeamFromUser structure.
type DelteTeamFromUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DelteTeamFromUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDelteTeamFromUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewDelteTeamFromUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 412:
		result := NewDelteTeamFromUserPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewDelteTeamFromUserUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDelteTeamFromUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDelteTeamFromUserOK creates a DelteTeamFromUserOK with default headers values
func NewDelteTeamFromUserOK() *DelteTeamFromUserOK {
	return &DelteTeamFromUserOK{}
}

/*DelteTeamFromUserOK handles this case with default header values.

Plain success message
*/
type DelteTeamFromUserOK struct {
	Payload *models.GeneralError
}

func (o *DelteTeamFromUserOK) Error() string {
	return fmt.Sprintf("[DELETE /teams/{team_id}/users][%d] delteTeamFromUserOK  %+v", 200, o.Payload)
}

func (o *DelteTeamFromUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDelteTeamFromUserForbidden creates a DelteTeamFromUserForbidden with default headers values
func NewDelteTeamFromUserForbidden() *DelteTeamFromUserForbidden {
	return &DelteTeamFromUserForbidden{}
}

/*DelteTeamFromUserForbidden handles this case with default header values.

User is not authorized
*/
type DelteTeamFromUserForbidden struct {
	Payload *models.GeneralError
}

func (o *DelteTeamFromUserForbidden) Error() string {
	return fmt.Sprintf("[DELETE /teams/{team_id}/users][%d] delteTeamFromUserForbidden  %+v", 403, o.Payload)
}

func (o *DelteTeamFromUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDelteTeamFromUserPreconditionFailed creates a DelteTeamFromUserPreconditionFailed with default headers values
func NewDelteTeamFromUserPreconditionFailed() *DelteTeamFromUserPreconditionFailed {
	return &DelteTeamFromUserPreconditionFailed{}
}

/*DelteTeamFromUserPreconditionFailed handles this case with default header values.

Failed to parse request body
*/
type DelteTeamFromUserPreconditionFailed struct {
	Payload *models.GeneralError
}

func (o *DelteTeamFromUserPreconditionFailed) Error() string {
	return fmt.Sprintf("[DELETE /teams/{team_id}/users][%d] delteTeamFromUserPreconditionFailed  %+v", 412, o.Payload)
}

func (o *DelteTeamFromUserPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDelteTeamFromUserUnprocessableEntity creates a DelteTeamFromUserUnprocessableEntity with default headers values
func NewDelteTeamFromUserUnprocessableEntity() *DelteTeamFromUserUnprocessableEntity {
	return &DelteTeamFromUserUnprocessableEntity{}
}

/*DelteTeamFromUserUnprocessableEntity handles this case with default header values.

User is not assigned
*/
type DelteTeamFromUserUnprocessableEntity struct {
	Payload *models.GeneralError
}

func (o *DelteTeamFromUserUnprocessableEntity) Error() string {
	return fmt.Sprintf("[DELETE /teams/{team_id}/users][%d] delteTeamFromUserUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *DelteTeamFromUserUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDelteTeamFromUserDefault creates a DelteTeamFromUserDefault with default headers values
func NewDelteTeamFromUserDefault(code int) *DelteTeamFromUserDefault {
	return &DelteTeamFromUserDefault{
		_statusCode: code,
	}
}

/*DelteTeamFromUserDefault handles this case with default header values.

Some error unrelated to the handler
*/
type DelteTeamFromUserDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// Code gets the status code for the delte team from user default response
func (o *DelteTeamFromUserDefault) Code() int {
	return o._statusCode
}

func (o *DelteTeamFromUserDefault) Error() string {
	return fmt.Sprintf("[DELETE /teams/{team_id}/users][%d] DelteTeamFromUser default  %+v", o._statusCode, o.Payload)
}

func (o *DelteTeamFromUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
