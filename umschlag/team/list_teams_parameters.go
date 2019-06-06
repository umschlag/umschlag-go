// Code generated by go-swagger; DO NOT EDIT.

package team

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

// NewListTeamsParams creates a new ListTeamsParams object
// with the default values initialized.
func NewListTeamsParams() *ListTeamsParams {

	return &ListTeamsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListTeamsParamsWithTimeout creates a new ListTeamsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListTeamsParamsWithTimeout(timeout time.Duration) *ListTeamsParams {

	return &ListTeamsParams{

		timeout: timeout,
	}
}

// NewListTeamsParamsWithContext creates a new ListTeamsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListTeamsParamsWithContext(ctx context.Context) *ListTeamsParams {

	return &ListTeamsParams{

		Context: ctx,
	}
}

// NewListTeamsParamsWithHTTPClient creates a new ListTeamsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListTeamsParamsWithHTTPClient(client *http.Client) *ListTeamsParams {

	return &ListTeamsParams{
		HTTPClient: client,
	}
}

/*ListTeamsParams contains all the parameters to send to the API endpoint
for the list teams operation typically these are written to a http.Request
*/
type ListTeamsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list teams params
func (o *ListTeamsParams) WithTimeout(timeout time.Duration) *ListTeamsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list teams params
func (o *ListTeamsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list teams params
func (o *ListTeamsParams) WithContext(ctx context.Context) *ListTeamsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list teams params
func (o *ListTeamsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list teams params
func (o *ListTeamsParams) WithHTTPClient(client *http.Client) *ListTeamsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list teams params
func (o *ListTeamsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListTeamsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
