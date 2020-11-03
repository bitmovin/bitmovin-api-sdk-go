package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingConfigurationsVideoH264CustomdataAPI communicates with '/encoding/configurations/video/h264/{configuration_id}/customData' endpoints
type EncodingConfigurationsVideoH264CustomdataAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingConfigurationsVideoH264CustomdataAPI constructor for EncodingConfigurationsVideoH264CustomdataAPI that takes options as argument
func NewEncodingConfigurationsVideoH264CustomdataAPI(options ...apiclient.APIClientOption) (*EncodingConfigurationsVideoH264CustomdataAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingConfigurationsVideoH264CustomdataAPIWithClient(apiClient), nil
}

// NewEncodingConfigurationsVideoH264CustomdataAPIWithClient constructor for EncodingConfigurationsVideoH264CustomdataAPI that takes an APIClient as argument
func NewEncodingConfigurationsVideoH264CustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingConfigurationsVideoH264CustomdataAPI {
    a := &EncodingConfigurationsVideoH264CustomdataAPI{apiClient: apiClient}
    return a
}

// Get H264/AVC Codec Configuration Custom Data
func (api *EncodingConfigurationsVideoH264CustomdataAPI) Get(configurationId string) (*model.CustomData, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel model.CustomData
    err := api.apiClient.Get("/encoding/configurations/video/h264/{configuration_id}/customData", nil, &responseModel, reqParams)
    return &responseModel, err
}

