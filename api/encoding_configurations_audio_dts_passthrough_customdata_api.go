package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingConfigurationsAudioDtsPassthroughCustomdataAPI communicates with '/encoding/configurations/audio/dts-passthrough/{configuration_id}/customData' endpoints
type EncodingConfigurationsAudioDtsPassthroughCustomdataAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingConfigurationsAudioDtsPassthroughCustomdataAPI constructor for EncodingConfigurationsAudioDtsPassthroughCustomdataAPI that takes options as argument
func NewEncodingConfigurationsAudioDtsPassthroughCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingConfigurationsAudioDtsPassthroughCustomdataAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingConfigurationsAudioDtsPassthroughCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingConfigurationsAudioDtsPassthroughCustomdataAPIWithClient constructor for EncodingConfigurationsAudioDtsPassthroughCustomdataAPI that takes an APIClient as argument
func NewEncodingConfigurationsAudioDtsPassthroughCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingConfigurationsAudioDtsPassthroughCustomdataAPI {
    a := &EncodingConfigurationsAudioDtsPassthroughCustomdataAPI{apiClient: apiClient}
    return a
}

// Get DTS Passthrough Codec Configuration Custom Data
func (api *EncodingConfigurationsAudioDtsPassthroughCustomdataAPI) Get(configurationId string) (*model.CustomData, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel model.CustomData
    err := api.apiClient.Get("/encoding/configurations/audio/dts-passthrough/{configuration_id}/customData", nil, &responseModel, reqParams)
    return &responseModel, err
}

