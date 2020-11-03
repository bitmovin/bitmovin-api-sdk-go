package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingConfigurationsVideoAv1CustomdataAPI communicates with '/encoding/configurations/video/av1/{configuration_id}/customData' endpoints
type EncodingConfigurationsVideoAv1CustomdataAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingConfigurationsVideoAv1CustomdataAPI constructor for EncodingConfigurationsVideoAv1CustomdataAPI that takes options as argument
func NewEncodingConfigurationsVideoAv1CustomdataAPI(options ...apiclient.APIClientOption) (*EncodingConfigurationsVideoAv1CustomdataAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingConfigurationsVideoAv1CustomdataAPIWithClient(apiClient), nil
}

// NewEncodingConfigurationsVideoAv1CustomdataAPIWithClient constructor for EncodingConfigurationsVideoAv1CustomdataAPI that takes an APIClient as argument
func NewEncodingConfigurationsVideoAv1CustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingConfigurationsVideoAv1CustomdataAPI {
    a := &EncodingConfigurationsVideoAv1CustomdataAPI{apiClient: apiClient}
    return a
}

// Get AV1 Codec Configuration Custom Data
func (api *EncodingConfigurationsVideoAv1CustomdataAPI) Get(configurationId string) (*model.CustomData, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel model.CustomData
    err := api.apiClient.Get("/encoding/configurations/video/av1/{configuration_id}/customData", nil, &responseModel, reqParams)
    return &responseModel, err
}

