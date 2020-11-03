package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingConfigurationsAudioPcmCustomdataAPI communicates with '/encoding/configurations/audio/pcm/{configuration_id}/customData' endpoints
type EncodingConfigurationsAudioPcmCustomdataAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingConfigurationsAudioPcmCustomdataAPI constructor for EncodingConfigurationsAudioPcmCustomdataAPI that takes options as argument
func NewEncodingConfigurationsAudioPcmCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingConfigurationsAudioPcmCustomdataAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingConfigurationsAudioPcmCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingConfigurationsAudioPcmCustomdataAPIWithClient constructor for EncodingConfigurationsAudioPcmCustomdataAPI that takes an APIClient as argument
func NewEncodingConfigurationsAudioPcmCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingConfigurationsAudioPcmCustomdataAPI {
    a := &EncodingConfigurationsAudioPcmCustomdataAPI{apiClient: apiClient}
    return a
}

// Get PCM Codec Configuration Custom Data
func (api *EncodingConfigurationsAudioPcmCustomdataAPI) Get(configurationId string) (*model.CustomData, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel model.CustomData
    err := api.apiClient.Get("/encoding/configurations/audio/pcm/{configuration_id}/customData", nil, &responseModel, reqParams)
    return &responseModel, err
}

