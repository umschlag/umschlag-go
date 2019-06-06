// Code generated by go-swagger; DO NOT EDIT.

package umschlag

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/umschlag/umschlag-go/umschlag/auth"
	"github.com/umschlag/umschlag-go/umschlag/profile"
	"github.com/umschlag/umschlag-go/umschlag/team"
	"github.com/umschlag/umschlag-go/umschlag/user"
)

// Default umschlag open HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "try.umschlag.tech"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/api/v1"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http", "https"}

// NewHTTPClient creates a new umschlag open HTTP client.
func NewHTTPClient(formats strfmt.Registry) *UmschlagOpen {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new umschlag open HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *UmschlagOpen {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new umschlag open client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *UmschlagOpen {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(UmschlagOpen)
	cli.Transport = transport

	cli.Auth = auth.New(transport, formats)

	cli.Profile = profile.New(transport, formats)

	cli.Team = team.New(transport, formats)

	cli.User = user.New(transport, formats)

	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// UmschlagOpen is a client for umschlag open
type UmschlagOpen struct {
	Auth *auth.Client

	Profile *profile.Client

	Team *team.Client

	User *user.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *UmschlagOpen) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.Auth.SetTransport(transport)

	c.Profile.SetTransport(transport)

	c.Team.SetTransport(transport)

	c.User.SetTransport(transport)

}
