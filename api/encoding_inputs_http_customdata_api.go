package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingInputsHttpCustomdataAPI communicates with '/encoding/inputs/http/{input_id}/customData' endpoints
type EncodingInputsHttpCustomdataAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingInputsHttpCustomdataAPI constructor for EncodingInputsHttpCustomdataAPI that takes options as argument
func NewEncodingInputsHttpCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingInputsHttpCustomdataAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingInputsHttpCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingInputsHttpCustomdataAPIWithClient constructor for EncodingInputsHttpCustomdataAPI that takes an APIClient as argument
func NewEncodingInputsHttpCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingInputsHttpCustomdataAPI {
    a := &EncodingInputsHttpCustomdataAPI{apiClient: apiClient}
    return a
}

// Get HTTP Custom Data
func (api *EncodingInputsHttpCustomdataAPI) Get(inputId string) (*model.CustomData, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel model.CustomData
    err := api.apiClient.Get("/encoding/inputs/http/{input_id}/customData", nil, &responseModel, reqParams)
    return &responseModel, err
}

