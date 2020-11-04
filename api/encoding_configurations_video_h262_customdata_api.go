package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingConfigurationsVideoH262CustomdataAPI communicates with '/encoding/configurations/video/h262/{configuration_id}/customData' endpoints
type EncodingConfigurationsVideoH262CustomdataAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingConfigurationsVideoH262CustomdataAPI constructor for EncodingConfigurationsVideoH262CustomdataAPI that takes options as argument
func NewEncodingConfigurationsVideoH262CustomdataAPI(options ...apiclient.APIClientOption) (*EncodingConfigurationsVideoH262CustomdataAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingConfigurationsVideoH262CustomdataAPIWithClient(apiClient), nil
}

// NewEncodingConfigurationsVideoH262CustomdataAPIWithClient constructor for EncodingConfigurationsVideoH262CustomdataAPI that takes an APIClient as argument
func NewEncodingConfigurationsVideoH262CustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingConfigurationsVideoH262CustomdataAPI {
	a := &EncodingConfigurationsVideoH262CustomdataAPI{apiClient: apiClient}
	return a
}

// Get H262 Codec Configuration Custom Data
func (api *EncodingConfigurationsVideoH262CustomdataAPI) Get(configurationId string) (*model.CustomData, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["configuration_id"] = configurationId
	}

	var responseModel model.CustomData
	err := api.apiClient.Get("/encoding/configurations/video/h262/{configuration_id}/customData", nil, &responseModel, reqParams)
	return &responseModel, err
}
