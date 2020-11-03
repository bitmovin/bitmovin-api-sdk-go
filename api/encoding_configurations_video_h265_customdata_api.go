package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingConfigurationsVideoH265CustomdataAPI communicates with '/encoding/configurations/video/h265/{configuration_id}/customData' endpoints
type EncodingConfigurationsVideoH265CustomdataAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingConfigurationsVideoH265CustomdataAPI constructor for EncodingConfigurationsVideoH265CustomdataAPI that takes options as argument
func NewEncodingConfigurationsVideoH265CustomdataAPI(options ...apiclient.APIClientOption) (*EncodingConfigurationsVideoH265CustomdataAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingConfigurationsVideoH265CustomdataAPIWithClient(apiClient), nil
}

// NewEncodingConfigurationsVideoH265CustomdataAPIWithClient constructor for EncodingConfigurationsVideoH265CustomdataAPI that takes an APIClient as argument
func NewEncodingConfigurationsVideoH265CustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingConfigurationsVideoH265CustomdataAPI {
    a := &EncodingConfigurationsVideoH265CustomdataAPI{apiClient: apiClient}
    return a
}

// Get H265/HEVC Codec Configuration Custom Data
func (api *EncodingConfigurationsVideoH265CustomdataAPI) Get(configurationId string) (*model.CustomData, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel model.CustomData
    err := api.apiClient.Get("/encoding/configurations/video/h265/{configuration_id}/customData", nil, &responseModel, reqParams)
    return &responseModel, err
}

