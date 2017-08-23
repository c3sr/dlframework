// Code generated by go-swagger; DO NOT EDIT.

package predictor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new predictor API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for predictor API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
Dataset echos method receives a simple message and returns it

The message posted as the id parameter will also be
returned.
*/
func (a *Client) Dataset(params *DatasetParams) (*DatasetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDatasetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Dataset",
		Method:             "POST",
		PathPattern:        "/v1/predict/dataset",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DatasetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DatasetOK), nil

}

/*
Images images API
*/
func (a *Client) Images(params *ImagesParams) (*ImagesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImagesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Images",
		Method:             "POST",
		PathPattern:        "/v1/predict/images",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ImagesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ImagesOK), nil

}

/*
Urls urls API
*/
func (a *Client) Urls(params *UrlsParams) (*UrlsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUrlsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "URLs",
		Method:             "POST",
		PathPattern:        "/v1/predict/urls",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UrlsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UrlsOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}