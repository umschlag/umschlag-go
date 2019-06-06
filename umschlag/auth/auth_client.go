// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new auth API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for auth API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
LoginUser authenticates an user by credentials
*/
func (a *Client) LoginUser(params *LoginUserParams) (*LoginUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLoginUserParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "LoginUser",
		Method:             "POST",
		PathPattern:        "/auth/login",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &LoginUserReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*LoginUserOK), nil

}

/*
RefreshAuth refreshes an auth token before it expires
*/
func (a *Client) RefreshAuth(params *RefreshAuthParams) (*RefreshAuthOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRefreshAuthParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RefreshAuth",
		Method:             "GET",
		PathPattern:        "/auth/refresh",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RefreshAuthReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RefreshAuthOK), nil

}

/*
VerifyAuth verifies validity for an authentication token
*/
func (a *Client) VerifyAuth(params *VerifyAuthParams) (*VerifyAuthOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVerifyAuthParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "VerifyAuth",
		Method:             "GET",
		PathPattern:        "/auth/verify/{token}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &VerifyAuthReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*VerifyAuthOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}