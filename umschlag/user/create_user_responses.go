// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/umschlag/umschlag-go/models"
)

// CreateUserReader is a Reader for the CreateUser structure.
type CreateUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewCreateUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 412:
		result := NewCreateUserPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewCreateUserUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCreateUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateUserOK creates a CreateUserOK with default headers values
func NewCreateUserOK() *CreateUserOK {
	return &CreateUserOK{}
}

/*CreateUserOK handles this case with default header values.

The created user data
*/
type CreateUserOK struct {
	Payload *models.User
}

func (o *CreateUserOK) Error() string {
	return fmt.Sprintf("[POST /users][%d] createUserOK  %+v", 200, o.Payload)
}

func (o *CreateUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserForbidden creates a CreateUserForbidden with default headers values
func NewCreateUserForbidden() *CreateUserForbidden {
	return &CreateUserForbidden{}
}

/*CreateUserForbidden handles this case with default header values.

User is not authorized
*/
type CreateUserForbidden struct {
	Payload *models.GeneralError
}

func (o *CreateUserForbidden) Error() string {
	return fmt.Sprintf("[POST /users][%d] createUserForbidden  %+v", 403, o.Payload)
}

func (o *CreateUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserPreconditionFailed creates a CreateUserPreconditionFailed with default headers values
func NewCreateUserPreconditionFailed() *CreateUserPreconditionFailed {
	return &CreateUserPreconditionFailed{}
}

/*CreateUserPreconditionFailed handles this case with default header values.

Failed to parse request body
*/
type CreateUserPreconditionFailed struct {
	Payload *models.GeneralError
}

func (o *CreateUserPreconditionFailed) Error() string {
	return fmt.Sprintf("[POST /users][%d] createUserPreconditionFailed  %+v", 412, o.Payload)
}

func (o *CreateUserPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserUnprocessableEntity creates a CreateUserUnprocessableEntity with default headers values
func NewCreateUserUnprocessableEntity() *CreateUserUnprocessableEntity {
	return &CreateUserUnprocessableEntity{}
}

/*CreateUserUnprocessableEntity handles this case with default header values.

Failed to validate request
*/
type CreateUserUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *CreateUserUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /users][%d] createUserUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateUserUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserDefault creates a CreateUserDefault with default headers values
func NewCreateUserDefault(code int) *CreateUserDefault {
	return &CreateUserDefault{
		_statusCode: code,
	}
}

/*CreateUserDefault handles this case with default header values.

Some error unrelated to the handler
*/
type CreateUserDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// Code gets the status code for the create user default response
func (o *CreateUserDefault) Code() int {
	return o._statusCode
}

func (o *CreateUserDefault) Error() string {
	return fmt.Sprintf("[POST /users][%d] CreateUser default  %+v", o._statusCode, o.Payload)
}

func (o *CreateUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
