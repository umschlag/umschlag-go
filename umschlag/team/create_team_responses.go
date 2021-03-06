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

// CreateTeamReader is a Reader for the CreateTeam structure.
type CreateTeamReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateTeamReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateTeamOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewCreateTeamForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 412:
		result := NewCreateTeamPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewCreateTeamUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCreateTeamDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateTeamOK creates a CreateTeamOK with default headers values
func NewCreateTeamOK() *CreateTeamOK {
	return &CreateTeamOK{}
}

/*CreateTeamOK handles this case with default header values.

The created team data
*/
type CreateTeamOK struct {
	Payload *models.Team
}

func (o *CreateTeamOK) Error() string {
	return fmt.Sprintf("[POST /teams][%d] createTeamOK  %+v", 200, o.Payload)
}

func (o *CreateTeamOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Team)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTeamForbidden creates a CreateTeamForbidden with default headers values
func NewCreateTeamForbidden() *CreateTeamForbidden {
	return &CreateTeamForbidden{}
}

/*CreateTeamForbidden handles this case with default header values.

User is not authorized
*/
type CreateTeamForbidden struct {
	Payload *models.GeneralError
}

func (o *CreateTeamForbidden) Error() string {
	return fmt.Sprintf("[POST /teams][%d] createTeamForbidden  %+v", 403, o.Payload)
}

func (o *CreateTeamForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTeamPreconditionFailed creates a CreateTeamPreconditionFailed with default headers values
func NewCreateTeamPreconditionFailed() *CreateTeamPreconditionFailed {
	return &CreateTeamPreconditionFailed{}
}

/*CreateTeamPreconditionFailed handles this case with default header values.

Failed to parse request body
*/
type CreateTeamPreconditionFailed struct {
	Payload *models.GeneralError
}

func (o *CreateTeamPreconditionFailed) Error() string {
	return fmt.Sprintf("[POST /teams][%d] createTeamPreconditionFailed  %+v", 412, o.Payload)
}

func (o *CreateTeamPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTeamUnprocessableEntity creates a CreateTeamUnprocessableEntity with default headers values
func NewCreateTeamUnprocessableEntity() *CreateTeamUnprocessableEntity {
	return &CreateTeamUnprocessableEntity{}
}

/*CreateTeamUnprocessableEntity handles this case with default header values.

Failed to validate request
*/
type CreateTeamUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *CreateTeamUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /teams][%d] createTeamUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateTeamUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTeamDefault creates a CreateTeamDefault with default headers values
func NewCreateTeamDefault(code int) *CreateTeamDefault {
	return &CreateTeamDefault{
		_statusCode: code,
	}
}

/*CreateTeamDefault handles this case with default header values.

Some error unrelated to the handler
*/
type CreateTeamDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// Code gets the status code for the create team default response
func (o *CreateTeamDefault) Code() int {
	return o._statusCode
}

func (o *CreateTeamDefault) Error() string {
	return fmt.Sprintf("[POST /teams][%d] CreateTeam default  %+v", o._statusCode, o.Payload)
}

func (o *CreateTeamDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
