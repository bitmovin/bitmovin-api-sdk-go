package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingInputsSrtCustomdataAPI communicates with '/encoding/inputs/srt/{input_id}/customData' endpoints
type EncodingInputsSrtCustomdataAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingInputsSrtCustomdataAPI constructor for EncodingInputsSrtCustomdataAPI that takes options as argument
func NewEncodingInputsSrtCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingInputsSrtCustomdataAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingInputsSrtCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingInputsSrtCustomdataAPIWithClient constructor for EncodingInputsSrtCustomdataAPI that takes an APIClient as argument
func NewEncodingInputsSrtCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingInputsSrtCustomdataAPI {
    a := &EncodingInputsSrtCustomdataAPI{apiClient: apiClient}
    return a
}

// Get SRT input Custom Data
func (api *EncodingInputsSrtCustomdataAPI) Get(inputId string) (*model.CustomData, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel model.CustomData
    err := api.apiClient.Get("/encoding/inputs/srt/{input_id}/customData", nil, &responseModel, reqParams)
    return &responseModel, err
}

