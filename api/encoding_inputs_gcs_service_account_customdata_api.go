package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingInputsGcsServiceAccountCustomdataAPI communicates with '/encoding/inputs/gcs-service-account/{input_id}/customData' endpoints
type EncodingInputsGcsServiceAccountCustomdataAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingInputsGcsServiceAccountCustomdataAPI constructor for EncodingInputsGcsServiceAccountCustomdataAPI that takes options as argument
func NewEncodingInputsGcsServiceAccountCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingInputsGcsServiceAccountCustomdataAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingInputsGcsServiceAccountCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingInputsGcsServiceAccountCustomdataAPIWithClient constructor for EncodingInputsGcsServiceAccountCustomdataAPI that takes an APIClient as argument
func NewEncodingInputsGcsServiceAccountCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingInputsGcsServiceAccountCustomdataAPI {
    a := &EncodingInputsGcsServiceAccountCustomdataAPI{apiClient: apiClient}
    return a
}

// Get GCS input Custom Data
func (api *EncodingInputsGcsServiceAccountCustomdataAPI) Get(inputId string) (*model.CustomData, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel model.CustomData
    err := api.apiClient.Get("/encoding/inputs/gcs-service-account/{input_id}/customData", nil, &responseModel, reqParams)
    return &responseModel, err
}

