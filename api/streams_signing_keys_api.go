package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// StreamsSigningKeysAPI communicates with '/streams/signing-keys' endpoints
type StreamsSigningKeysAPI struct {
	apiClient *apiclient.APIClient
}

// NewStreamsSigningKeysAPI constructor for StreamsSigningKeysAPI that takes options as argument
func NewStreamsSigningKeysAPI(options ...apiclient.APIClientOption) (*StreamsSigningKeysAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewStreamsSigningKeysAPIWithClient(apiClient), nil
}

// NewStreamsSigningKeysAPIWithClient constructor for StreamsSigningKeysAPI that takes an APIClient as argument
func NewStreamsSigningKeysAPIWithClient(apiClient *apiclient.APIClient) *StreamsSigningKeysAPI {
	a := &StreamsSigningKeysAPI{apiClient: apiClient}
	return a
}

// Create new signing-key
// Create new signing-key pair. There is a limit of 2 active signing keys per organization.
func (api *StreamsSigningKeysAPI) Create() (*model.StreamsSigningKeyResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.StreamsSigningKeyResponse
	err := api.apiClient.Post("/streams/signing-keys", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Delete signing-key
func (api *StreamsSigningKeysAPI) Delete(keyId string) error {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["key_id"] = keyId
	}

	err := api.apiClient.Delete("/streams/signing-keys/{key_id}", nil, nil, reqParams)
	return err
}

// Get list of public signing key ids
func (api *StreamsSigningKeysAPI) Get() (*model.StreamsPublicSigningKeyResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.StreamsPublicSigningKeyResponse
	err := api.apiClient.Get("/streams/signing-keys", nil, &responseModel, reqParams)
	return &responseModel, err
}
