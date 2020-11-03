package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingInputsGcsCustomdataAPI communicates with '/encoding/inputs/gcs/{input_id}/customData' endpoints
type EncodingInputsGcsCustomdataAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingInputsGcsCustomdataAPI constructor for EncodingInputsGcsCustomdataAPI that takes options as argument
func NewEncodingInputsGcsCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingInputsGcsCustomdataAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingInputsGcsCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingInputsGcsCustomdataAPIWithClient constructor for EncodingInputsGcsCustomdataAPI that takes an APIClient as argument
func NewEncodingInputsGcsCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingInputsGcsCustomdataAPI {
    a := &EncodingInputsGcsCustomdataAPI{apiClient: apiClient}
    return a
}

// Get GCS input Custom Data
func (api *EncodingInputsGcsCustomdataAPI) Get(inputId string) (*model.CustomData, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel model.CustomData
    err := api.apiClient.Get("/encoding/inputs/gcs/{input_id}/customData", nil, &responseModel, reqParams)
    return &responseModel, err
}

