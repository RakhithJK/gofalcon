// Code generated by go-swagger; DO NOT EDIT.

package identity_protection

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new identity protection API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for identity protection API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	APIPreemptProxyPostGraphql(params *APIPreemptProxyPostGraphqlParams, opts ...ClientOption) (*APIPreemptProxyPostGraphqlOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  APIPreemptProxyPostGraphql identities protection graph q l API allows to retrieve entities timeline activities identity based incidents and security assessment allows to perform actions on entities and identity based incidents
*/
func (a *Client) APIPreemptProxyPostGraphql(params *APIPreemptProxyPostGraphqlParams, opts ...ClientOption) (*APIPreemptProxyPostGraphqlOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAPIPreemptProxyPostGraphqlParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "api.preempt.proxy.post.graphql",
		Method:             "POST",
		PathPattern:        "/identity-protection/combined/graphql/v1",
		ProducesMediaTypes: []string{"application/json", "application/json; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &APIPreemptProxyPostGraphqlReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*APIPreemptProxyPostGraphqlOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for api.preempt.proxy.post.graphql: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}