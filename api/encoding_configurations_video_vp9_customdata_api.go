package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingConfigurationsVideoVp9CustomdataAPI communicates with '/encoding/configurations/video/vp9/{configuration_id}/customData' endpoints
type EncodingConfigurationsVideoVp9CustomdataAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingConfigurationsVideoVp9CustomdataAPI constructor for EncodingConfigurationsVideoVp9CustomdataAPI that takes options as argument
func NewEncodingConfigurationsVideoVp9CustomdataAPI(options ...apiclient.APIClientOption) (*EncodingConfigurationsVideoVp9CustomdataAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingConfigurationsVideoVp9CustomdataAPIWithClient(apiClient), nil
}

// NewEncodingConfigurationsVideoVp9CustomdataAPIWithClient constructor for EncodingConfigurationsVideoVp9CustomdataAPI that takes an APIClient as argument
func NewEncodingConfigurationsVideoVp9CustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingConfigurationsVideoVp9CustomdataAPI {
    a := &EncodingConfigurationsVideoVp9CustomdataAPI{apiClient: apiClient}
    return a
}

// Get VP9 Codec Configuration Custom Data
func (api *EncodingConfigurationsVideoVp9CustomdataAPI) Get(configurationId string) (*model.CustomData, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel model.CustomData
    err := api.apiClient.Get("/encoding/configurations/video/vp9/{configuration_id}/customData", nil, &responseModel, reqParams)
    return &responseModel, err
}

