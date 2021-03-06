// Code generated by go-swagger; DO NOT EDIT.

package prospects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new prospects API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for prospects API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetProspects(params *GetProspectsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetProspectsOK, *GetProspectsNoContent, error)

	PostAddProspectsCampaign(params *PostAddProspectsCampaignParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostAddProspectsCampaignOK, error)

	PostAddProspectsList(params *PostAddProspectsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostAddProspectsListOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetProspects gets the list of prospects

  get the list of prospects

*/
func (a *Client) GetProspects(params *GetProspectsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetProspectsOK, *GetProspectsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProspectsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetProspects",
		Method:             "GET",
		PathPattern:        "/prospects",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetProspectsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *GetProspectsOK:
		return value, nil, nil
	case *GetProspectsNoContent:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetProspectsDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PostAddProspectsCampaign adds prospects to the campaign

  get the list of campaigns

*/
func (a *Client) PostAddProspectsCampaign(params *PostAddProspectsCampaignParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostAddProspectsCampaignOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAddProspectsCampaignParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostAddProspectsCampaign",
		Method:             "POST",
		PathPattern:        "/add_prospects_campaign",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostAddProspectsCampaignReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*PostAddProspectsCampaignOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostAddProspectsCampaignDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PostAddProspectsList adds prospects to the global prospects list

  get the list of campaigns

*/
func (a *Client) PostAddProspectsList(params *PostAddProspectsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostAddProspectsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAddProspectsListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostAddProspectsList",
		Method:             "POST",
		PathPattern:        "/add_prospects_list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostAddProspectsListReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*PostAddProspectsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostAddProspectsListDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
