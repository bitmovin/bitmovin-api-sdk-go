package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// AccountApiKeysAPI communicates with '/account/api-keys' endpoints
type AccountApiKeysAPI struct {
	apiClient *apiclient.APIClient
}

// NewAccountApiKeysAPI constructor for AccountApiKeysAPI that takes options as argument
func NewAccountApiKeysAPI(options ...apiclient.APIClientOption) (*AccountApiKeysAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewAccountApiKeysAPIWithClient(apiClient), nil
}

// NewAccountApiKeysAPIWithClient constructor for AccountApiKeysAPI that takes an APIClient as argument
func NewAccountApiKeysAPIWithClient(apiClient *apiclient.APIClient) *AccountApiKeysAPI {
	a := &AccountApiKeysAPI{apiClient: apiClient}
	return a
}

// Create Api Key
func (api *AccountApiKeysAPI) Create() (*model.AccountApiKey, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.AccountApiKey
	err := api.apiClient.Post("/account/api-keys", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Api Key
func (api *AccountApiKeysAPI) Delete(apiKeyId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["api_key_id"] = apiKeyId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/account/api-keys/{api_key_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Api Key
func (api *AccountApiKeysAPI) Get(apiKeyId string) (*model.AccountApiKey, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["api_key_id"] = apiKeyId
	}

	var responseModel model.AccountApiKey
	err := api.apiClient.Get("/account/api-keys/{api_key_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List Api Keys
func (api *AccountApiKeysAPI) List() (*pagination.AccountApiKeysListPagination, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel pagination.AccountApiKeysListPagination
	err := api.apiClient.Get("/account/api-keys", nil, &responseModel, reqParams)
	return &responseModel, err
}
