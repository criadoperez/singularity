// Code generated by go-swagger; DO NOT EDIT.

package preparation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new preparation API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for preparation API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AddOutputStorage(params *AddOutputStorageParams, opts ...ClientOption) (*AddOutputStorageOK, error)

	AddSourceStorage(params *AddSourceStorageParams, opts ...ClientOption) (*AddSourceStorageOK, error)

	CreatePreparation(params *CreatePreparationParams, opts ...ClientOption) (*CreatePreparationOK, error)

	ExplorePreparation(params *ExplorePreparationParams, opts ...ClientOption) (*ExplorePreparationOK, error)

	GetPreparationStatus(params *GetPreparationStatusParams, opts ...ClientOption) (*GetPreparationStatusOK, error)

	ListPreparations(params *ListPreparationsParams, opts ...ClientOption) (*ListPreparationsOK, error)

	RemoveOutputStorage(params *RemoveOutputStorageParams, opts ...ClientOption) (*RemoveOutputStorageOK, error)

	RenamePreparation(params *RenamePreparationParams, opts ...ClientOption) (*RenamePreparationOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
AddOutputStorage attaches an output storage with a preparation
*/
func (a *Client) AddOutputStorage(params *AddOutputStorageParams, opts ...ClientOption) (*AddOutputStorageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddOutputStorageParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AddOutputStorage",
		Method:             "POST",
		PathPattern:        "/preparation/{id}/output/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AddOutputStorageReader{formats: a.formats},
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
	success, ok := result.(*AddOutputStorageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for AddOutputStorage: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AddSourceStorage attaches a source storage with a preparation
*/
func (a *Client) AddSourceStorage(params *AddSourceStorageParams, opts ...ClientOption) (*AddSourceStorageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddSourceStorageParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AddSourceStorage",
		Method:             "POST",
		PathPattern:        "/preparation/{id}/source/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AddSourceStorageReader{formats: a.formats},
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
	success, ok := result.(*AddSourceStorageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for AddSourceStorage: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CreatePreparation creates a new preparation
*/
func (a *Client) CreatePreparation(params *CreatePreparationParams, opts ...ClientOption) (*CreatePreparationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreatePreparationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreatePreparation",
		Method:             "POST",
		PathPattern:        "/preparation",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreatePreparationReader{formats: a.formats},
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
	success, ok := result.(*CreatePreparationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreatePreparation: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ExplorePreparation explores a directory in a prepared source storage
*/
func (a *Client) ExplorePreparation(params *ExplorePreparationParams, opts ...ClientOption) (*ExplorePreparationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExplorePreparationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ExplorePreparation",
		Method:             "GET",
		PathPattern:        "/preparation/{id}/source/{name}/explore/{path}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ExplorePreparationReader{formats: a.formats},
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
	success, ok := result.(*ExplorePreparationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ExplorePreparation: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetPreparationStatus gets the status of a preparation
*/
func (a *Client) GetPreparationStatus(params *GetPreparationStatusParams, opts ...ClientOption) (*GetPreparationStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPreparationStatusParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetPreparationStatus",
		Method:             "GET",
		PathPattern:        "/preparation/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPreparationStatusReader{formats: a.formats},
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
	success, ok := result.(*GetPreparationStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetPreparationStatus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListPreparations lists all preparations
*/
func (a *Client) ListPreparations(params *ListPreparationsParams, opts ...ClientOption) (*ListPreparationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListPreparationsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListPreparations",
		Method:             "GET",
		PathPattern:        "/preparation",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListPreparationsReader{formats: a.formats},
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
	success, ok := result.(*ListPreparationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListPreparations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
RemoveOutputStorage detaches an output storage from a preparation
*/
func (a *Client) RemoveOutputStorage(params *RemoveOutputStorageParams, opts ...ClientOption) (*RemoveOutputStorageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveOutputStorageParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RemoveOutputStorage",
		Method:             "DELETE",
		PathPattern:        "/preparation/{id}/output/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RemoveOutputStorageReader{formats: a.formats},
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
	success, ok := result.(*RemoveOutputStorageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RemoveOutputStorage: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
RenamePreparation renames a preparation
*/
func (a *Client) RenamePreparation(params *RenamePreparationParams, opts ...ClientOption) (*RenamePreparationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRenamePreparationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RenamePreparation",
		Method:             "PATCH",
		PathPattern:        "/preparation/{name}/rename",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RenamePreparationReader{formats: a.formats},
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
	success, ok := result.(*RenamePreparationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RenamePreparation: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
