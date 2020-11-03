package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingConfigurationsAudioOpusCustomdataAPI communicates with '/encoding/configurations/audio/opus/{configuration_id}/customData' endpoints
type EncodingConfigurationsAudioOpusCustomdataAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingConfigurationsAudioOpusCustomdataAPI constructor for EncodingConfigurationsAudioOpusCustomdataAPI that takes options as argument
func NewEncodingConfigurationsAudioOpusCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingConfigurationsAudioOpusCustomdataAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingConfigurationsAudioOpusCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingConfigurationsAudioOpusCustomdataAPIWithClient constructor for EncodingConfigurationsAudioOpusCustomdataAPI that takes an APIClient as argument
func NewEncodingConfigurationsAudioOpusCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingConfigurationsAudioOpusCustomdataAPI {
    a := &EncodingConfigurationsAudioOpusCustomdataAPI{apiClient: apiClient}
    return a
}

// Get Opus Codec Configuration Custom Data
func (api *EncodingConfigurationsAudioOpusCustomdataAPI) Get(configurationId string) (*model.CustomData, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel model.CustomData
    err := api.apiClient.Get("/encoding/configurations/audio/opus/{configuration_id}/customData", nil, &responseModel, reqParams)
    return &responseModel, err
}

