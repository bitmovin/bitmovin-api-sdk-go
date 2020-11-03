package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingFiltersAudioMixCustomdataAPI communicates with '/encoding/filters/audio-mix/{filter_id}/customData' endpoints
type EncodingFiltersAudioMixCustomdataAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingFiltersAudioMixCustomdataAPI constructor for EncodingFiltersAudioMixCustomdataAPI that takes options as argument
func NewEncodingFiltersAudioMixCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingFiltersAudioMixCustomdataAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingFiltersAudioMixCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingFiltersAudioMixCustomdataAPIWithClient constructor for EncodingFiltersAudioMixCustomdataAPI that takes an APIClient as argument
func NewEncodingFiltersAudioMixCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingFiltersAudioMixCustomdataAPI {
    a := &EncodingFiltersAudioMixCustomdataAPI{apiClient: apiClient}
    return a
}

// Get Audio Mix Filter Custom Data
func (api *EncodingFiltersAudioMixCustomdataAPI) Get(filterId string) (*model.CustomData, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["filter_id"] = filterId
    }

    var responseModel model.CustomData
    err := api.apiClient.Get("/encoding/filters/audio-mix/{filter_id}/customData", nil, &responseModel, reqParams)
    return &responseModel, err
}

