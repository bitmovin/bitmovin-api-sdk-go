package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingConfigurationsTypeAPI communicates with '/encoding/configurations/{configuration_id}/type' endpoints
type EncodingConfigurationsTypeAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingConfigurationsTypeAPI constructor for EncodingConfigurationsTypeAPI that takes options as argument
func NewEncodingConfigurationsTypeAPI(options ...apiclient.APIClientOption) (*EncodingConfigurationsTypeAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingConfigurationsTypeAPIWithClient(apiClient), nil
}

// NewEncodingConfigurationsTypeAPIWithClient constructor for EncodingConfigurationsTypeAPI that takes an APIClient as argument
func NewEncodingConfigurationsTypeAPIWithClient(apiClient *apiclient.APIClient) *EncodingConfigurationsTypeAPI {
    a := &EncodingConfigurationsTypeAPI{apiClient: apiClient}
    return a
}

// Get Codec Configuration Type
func (api *EncodingConfigurationsTypeAPI) Get(configurationId string) (*model.CodecConfigTypeResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel model.CodecConfigTypeResponse
    err := api.apiClient.Get("/encoding/configurations/{configuration_id}/type", nil, &responseModel, reqParams)
    return &responseModel, err
}

