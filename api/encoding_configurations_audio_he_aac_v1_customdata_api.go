package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingConfigurationsAudioHeAacV1CustomdataAPI communicates with '/encoding/configurations/audio/he-aac-v1/{configuration_id}/customData' endpoints
type EncodingConfigurationsAudioHeAacV1CustomdataAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingConfigurationsAudioHeAacV1CustomdataAPI constructor for EncodingConfigurationsAudioHeAacV1CustomdataAPI that takes options as argument
func NewEncodingConfigurationsAudioHeAacV1CustomdataAPI(options ...apiclient.APIClientOption) (*EncodingConfigurationsAudioHeAacV1CustomdataAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingConfigurationsAudioHeAacV1CustomdataAPIWithClient(apiClient), nil
}

// NewEncodingConfigurationsAudioHeAacV1CustomdataAPIWithClient constructor for EncodingConfigurationsAudioHeAacV1CustomdataAPI that takes an APIClient as argument
func NewEncodingConfigurationsAudioHeAacV1CustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingConfigurationsAudioHeAacV1CustomdataAPI {
    a := &EncodingConfigurationsAudioHeAacV1CustomdataAPI{apiClient: apiClient}
    return a
}

// Get HE-AAC v1 Codec Configuration Custom Data
func (api *EncodingConfigurationsAudioHeAacV1CustomdataAPI) Get(configurationId string) (*model.CustomData, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel model.CustomData
    err := api.apiClient.Get("/encoding/configurations/audio/he-aac-v1/{configuration_id}/customData", nil, &responseModel, reqParams)
    return &responseModel, err
}

