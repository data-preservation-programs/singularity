// Code generated by go-swagger; DO NOT EDIT.

package wallet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new wallet API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for wallet API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteDatasetDatasetNameWalletWallet(params *DeleteDatasetDatasetNameWalletWalletParams, opts ...ClientOption) (*DeleteDatasetDatasetNameWalletWalletNoContent, error)

	DeleteWalletAddress(params *DeleteWalletAddressParams, opts ...ClientOption) (*DeleteWalletAddressNoContent, error)

	GetDatasetDatasetNameWallet(params *GetDatasetDatasetNameWalletParams, opts ...ClientOption) (*GetDatasetDatasetNameWalletOK, error)

	GetWallet(params *GetWalletParams, opts ...ClientOption) (*GetWalletOK, error)

	PostWallet(params *PostWalletParams, opts ...ClientOption) (*PostWalletOK, error)

	PostWalletRemote(params *PostWalletRemoteParams, opts ...ClientOption) (*PostWalletRemoteOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DeleteDatasetDatasetNameWalletWallet removes an associated wallet from a dataset
*/
func (a *Client) DeleteDatasetDatasetNameWalletWallet(params *DeleteDatasetDatasetNameWalletWalletParams, opts ...ClientOption) (*DeleteDatasetDatasetNameWalletWalletNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteDatasetDatasetNameWalletWalletParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteDatasetDatasetNameWalletWallet",
		Method:             "DELETE",
		PathPattern:        "/dataset/{datasetName}/wallet/{wallet}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteDatasetDatasetNameWalletWalletReader{formats: a.formats},
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
	success, ok := result.(*DeleteDatasetDatasetNameWalletWalletNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteDatasetDatasetNameWalletWallet: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteWalletAddress removes a wallet
*/
func (a *Client) DeleteWalletAddress(params *DeleteWalletAddressParams, opts ...ClientOption) (*DeleteWalletAddressNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteWalletAddressParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteWalletAddress",
		Method:             "DELETE",
		PathPattern:        "/wallet/{address}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteWalletAddressReader{formats: a.formats},
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
	success, ok := result.(*DeleteWalletAddressNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteWalletAddress: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetDatasetDatasetNameWallet lists all wallets of a dataset
*/
func (a *Client) GetDatasetDatasetNameWallet(params *GetDatasetDatasetNameWalletParams, opts ...ClientOption) (*GetDatasetDatasetNameWalletOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDatasetDatasetNameWalletParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetDatasetDatasetNameWallet",
		Method:             "GET",
		PathPattern:        "/dataset/{datasetName}/wallet",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetDatasetDatasetNameWalletReader{formats: a.formats},
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
	success, ok := result.(*GetDatasetDatasetNameWalletOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetDatasetDatasetNameWallet: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetWallet lists all imported wallets
*/
func (a *Client) GetWallet(params *GetWalletParams, opts ...ClientOption) (*GetWalletOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWalletParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetWallet",
		Method:             "GET",
		PathPattern:        "/wallet",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetWalletReader{formats: a.formats},
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
	success, ok := result.(*GetWalletOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetWallet: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostWallet imports a private key
*/
func (a *Client) PostWallet(params *PostWalletParams, opts ...ClientOption) (*PostWalletOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostWalletParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostWallet",
		Method:             "POST",
		PathPattern:        "/wallet",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostWalletReader{formats: a.formats},
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
	success, ok := result.(*PostWalletOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostWallet: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostWalletRemote adds a remote wallet
*/
func (a *Client) PostWalletRemote(params *PostWalletRemoteParams, opts ...ClientOption) (*PostWalletRemoteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostWalletRemoteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostWalletRemote",
		Method:             "POST",
		PathPattern:        "/wallet/remote",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostWalletRemoteReader{formats: a.formats},
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
	success, ok := result.(*PostWalletRemoteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostWalletRemote: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}