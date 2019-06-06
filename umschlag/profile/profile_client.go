// Code generated by go-swagger; DO NOT EDIT.

package profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new profile API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for profile API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
ShowProfile retrieves an unlimited auth token
*/
func (a *Client) ShowProfile(params *ShowProfileParams) (*ShowProfileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewShowProfileParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ShowProfile",
		Method:             "GET",
		PathPattern:        "/profile/self",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ShowProfileReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ShowProfileOK), nil

}

/*
TokenProfile retrieves an unlimited auth token
*/
func (a *Client) TokenProfile(params *TokenProfileParams) (*TokenProfileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTokenProfileParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TokenProfile",
		Method:             "GET",
		PathPattern:        "/profile/token",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &TokenProfileReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TokenProfileOK), nil

}

/*
UpdateProfile retrieves an unlimited auth token
*/
func (a *Client) UpdateProfile(params *UpdateProfileParams) (*UpdateProfileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateProfileParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateProfile",
		Method:             "PUT",
		PathPattern:        "/profile/self",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateProfileReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateProfileOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
