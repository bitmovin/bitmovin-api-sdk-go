package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingInputsAzureCustomdataAPI communicates with '/encoding/inputs/azure/{input_id}/customData' endpoints
type EncodingInputsAzureCustomdataAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingInputsAzureCustomdataAPI constructor for EncodingInputsAzureCustomdataAPI that takes options as argument
func NewEncodingInputsAzureCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingInputsAzureCustomdataAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingInputsAzureCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingInputsAzureCustomdataAPIWithClient constructor for EncodingInputsAzureCustomdataAPI that takes an APIClient as argument
func NewEncodingInputsAzureCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingInputsAzureCustomdataAPI {
    a := &EncodingInputsAzureCustomdataAPI{apiClient: apiClient}
    return a
}

// Get Azure Custom Data
func (api *EncodingInputsAzureCustomdataAPI) Get(inputId string) (*model.CustomData, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel model.CustomData
    err := api.apiClient.Get("/encoding/inputs/azure/{input_id}/customData", nil, &responseModel, reqParams)
    return &responseModel, err
}

