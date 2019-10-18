// Code generated by go-swagger; DO NOT EDIT.

package endpoint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new endpoint API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for endpoint API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteEndpointID deletes endpoint

Deletes the endpoint specified by the ID. Deletion is imminent and
atomic, if the deletion request is valid and the endpoint exists,
deletion will occur even if errors are encountered in the process. If
errors have been encountered, the code 202 will be returned, otherwise
200 on success.

All resources associated with the endpoint will be freed and the
workload represented by the endpoint will be disconnected.It will no
longer be able to initiate or receive communications of any sort.

*/
func (a *Client) DeleteEndpointID(params *DeleteEndpointIDParams) (*DeleteEndpointIDOK, *DeleteEndpointIDErrors, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteEndpointIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteEndpointID",
		Method:             "DELETE",
		PathPattern:        "/endpoint/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteEndpointIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *DeleteEndpointIDOK:
		return value, nil, nil
	case *DeleteEndpointIDErrors:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for endpoint: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetEndpoint retrieves a list of endpoints that have metadata matching the provided parameters

Retrieves a list of endpoints that have metadata matching the provided parameters, or all endpoints if no parameters provided.

*/
func (a *Client) GetEndpoint(params *GetEndpointParams) (*GetEndpointOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEndpointParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetEndpoint",
		Method:             "GET",
		PathPattern:        "/endpoint",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetEndpointReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetEndpointOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetEndpoint: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetEndpointID gets endpoint by endpoint ID

Returns endpoint information

*/
func (a *Client) GetEndpointID(params *GetEndpointIDParams) (*GetEndpointIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEndpointIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetEndpointID",
		Method:             "GET",
		PathPattern:        "/endpoint/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetEndpointIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetEndpointIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetEndpointID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetEndpointIDConfig retrieves endpoint configuration

Retrieves the configuration of the specified endpoint.

*/
func (a *Client) GetEndpointIDConfig(params *GetEndpointIDConfigParams) (*GetEndpointIDConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEndpointIDConfigParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetEndpointIDConfig",
		Method:             "GET",
		PathPattern:        "/endpoint/{id}/config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetEndpointIDConfigReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetEndpointIDConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetEndpointIDConfig: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetEndpointIDHealthz retrieves the status logs associated with this endpoint
*/
func (a *Client) GetEndpointIDHealthz(params *GetEndpointIDHealthzParams) (*GetEndpointIDHealthzOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEndpointIDHealthzParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetEndpointIDHealthz",
		Method:             "GET",
		PathPattern:        "/endpoint/{id}/healthz",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetEndpointIDHealthzReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetEndpointIDHealthzOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetEndpointIDHealthz: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetEndpointIDLabels retrieves the list of labels associated with an endpoint
*/
func (a *Client) GetEndpointIDLabels(params *GetEndpointIDLabelsParams) (*GetEndpointIDLabelsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEndpointIDLabelsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetEndpointIDLabels",
		Method:             "GET",
		PathPattern:        "/endpoint/{id}/labels",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetEndpointIDLabelsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetEndpointIDLabelsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetEndpointIDLabels: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetEndpointIDLog retrieves the status logs associated with this endpoint
*/
func (a *Client) GetEndpointIDLog(params *GetEndpointIDLogParams) (*GetEndpointIDLogOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEndpointIDLogParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetEndpointIDLog",
		Method:             "GET",
		PathPattern:        "/endpoint/{id}/log",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetEndpointIDLogReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetEndpointIDLogOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetEndpointIDLog: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PatchEndpointID modifies existing endpoint

Applies the endpoint change request to an existing endpoint

*/
func (a *Client) PatchEndpointID(params *PatchEndpointIDParams) (*PatchEndpointIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchEndpointIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchEndpointID",
		Method:             "PATCH",
		PathPattern:        "/endpoint/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchEndpointIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchEndpointIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PatchEndpointID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PatchEndpointIDConfig modifies mutable endpoint configuration

Update the configuration of an existing endpoint and regenerates &
recompiles the corresponding programs automatically.

*/
func (a *Client) PatchEndpointIDConfig(params *PatchEndpointIDConfigParams) (*PatchEndpointIDConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchEndpointIDConfigParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchEndpointIDConfig",
		Method:             "PATCH",
		PathPattern:        "/endpoint/{id}/config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchEndpointIDConfigReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchEndpointIDConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PatchEndpointIDConfig: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PatchEndpointIDLabels sets label configuration of endpoint

Sets labels associated with an endpoint. These can be user provided or
derived from the orchestration system.

*/
func (a *Client) PatchEndpointIDLabels(params *PatchEndpointIDLabelsParams) (*PatchEndpointIDLabelsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchEndpointIDLabelsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchEndpointIDLabels",
		Method:             "PATCH",
		PathPattern:        "/endpoint/{id}/labels",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchEndpointIDLabelsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchEndpointIDLabelsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PatchEndpointIDLabels: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PutEndpointID creates endpoint

Creates a new endpoint

*/
func (a *Client) PutEndpointID(params *PutEndpointIDParams) (*PutEndpointIDCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutEndpointIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutEndpointID",
		Method:             "PUT",
		PathPattern:        "/endpoint/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutEndpointIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutEndpointIDCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PutEndpointID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
