// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewListUserTeamsParams creates a new ListUserTeamsParams object
// with the default values initialized.
func NewListUserTeamsParams() *ListUserTeamsParams {
	var ()
	return &ListUserTeamsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListUserTeamsParamsWithTimeout creates a new ListUserTeamsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListUserTeamsParamsWithTimeout(timeout time.Duration) *ListUserTeamsParams {
	var ()
	return &ListUserTeamsParams{

		timeout: timeout,
	}
}

// NewListUserTeamsParamsWithContext creates a new ListUserTeamsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListUserTeamsParamsWithContext(ctx context.Context) *ListUserTeamsParams {
	var ()
	return &ListUserTeamsParams{

		Context: ctx,
	}
}

// NewListUserTeamsParamsWithHTTPClient creates a new ListUserTeamsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListUserTeamsParamsWithHTTPClient(client *http.Client) *ListUserTeamsParams {
	var ()
	return &ListUserTeamsParams{
		HTTPClient: client,
	}
}

/*ListUserTeamsParams contains all the parameters to send to the API endpoint
for the list user teams operation typically these are written to a http.Request
*/
type ListUserTeamsParams struct {

	/*UserID
	  A user UUID or slug

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list user teams params
func (o *ListUserTeamsParams) WithTimeout(timeout time.Duration) *ListUserTeamsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list user teams params
func (o *ListUserTeamsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list user teams params
func (o *ListUserTeamsParams) WithContext(ctx context.Context) *ListUserTeamsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list user teams params
func (o *ListUserTeamsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list user teams params
func (o *ListUserTeamsParams) WithHTTPClient(client *http.Client) *ListUserTeamsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list user teams params
func (o *ListUserTeamsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserID adds the userID to the list user teams params
func (o *ListUserTeamsParams) WithUserID(userID string) *ListUserTeamsParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the list user teams params
func (o *ListUserTeamsParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *ListUserTeamsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param user_id
	if err := r.SetPathParam("user_id", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
