// Code generated by go-swagger; DO NOT EDIT.

package registry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new registry API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for registry API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
FrameworkAgents framework agents API
*/
func (a *Client) FrameworkAgents(params *FrameworkAgentsParams) (*FrameworkAgentsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFrameworkAgentsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "FrameworkAgents",
		Method:             "GET",
		PathPattern:        "/registry/frameworks/agent",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &FrameworkAgentsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*FrameworkAgentsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for FrameworkAgents: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FrameworkManifests framework manifests API
*/
func (a *Client) FrameworkManifests(params *FrameworkManifestsParams) (*FrameworkManifestsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFrameworkManifestsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "FrameworkManifests",
		Method:             "GET",
		PathPattern:        "/registry/frameworks/manifest",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &FrameworkManifestsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*FrameworkManifestsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for FrameworkManifests: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ModelAgents model agents API
*/
func (a *Client) ModelAgents(params *ModelAgentsParams) (*ModelAgentsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewModelAgentsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ModelAgents",
		Method:             "GET",
		PathPattern:        "/registry/models/agent",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ModelAgentsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ModelAgentsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ModelAgents: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ModelManifests model manifests API
*/
func (a *Client) ModelManifests(params *ModelManifestsParams) (*ModelManifestsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewModelManifestsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ModelManifests",
		Method:             "GET",
		PathPattern:        "/registry/models/manifest",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ModelManifestsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ModelManifestsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ModelManifests: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
