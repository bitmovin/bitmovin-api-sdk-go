package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingInputsAsperaCustomdataAPI communicates with '/encoding/inputs/aspera/{input_id}/customData' endpoints
type EncodingInputsAsperaCustomdataAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingInputsAsperaCustomdataAPI constructor for EncodingInputsAsperaCustomdataAPI that takes options as argument
func NewEncodingInputsAsperaCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingInputsAsperaCustomdataAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingInputsAsperaCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingInputsAsperaCustomdataAPIWithClient constructor for EncodingInputsAsperaCustomdataAPI that takes an APIClient as argument
func NewEncodingInputsAsperaCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingInputsAsperaCustomdataAPI {
    a := &EncodingInputsAsperaCustomdataAPI{apiClient: apiClient}
    return a
}

// Get Aspera Custom Data
func (api *EncodingInputsAsperaCustomdataAPI) Get(inputId string) (*model.CustomData, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel model.CustomData
    err := api.apiClient.Get("/encoding/inputs/aspera/{input_id}/customData", nil, &responseModel, reqParams)
    return &responseModel, err
}

