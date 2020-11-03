package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingOutputsGcsServiceAccountCustomdataAPI communicates with '/encoding/outputs/gcs-service-account/{output_id}/customData' endpoints
type EncodingOutputsGcsServiceAccountCustomdataAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingOutputsGcsServiceAccountCustomdataAPI constructor for EncodingOutputsGcsServiceAccountCustomdataAPI that takes options as argument
func NewEncodingOutputsGcsServiceAccountCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingOutputsGcsServiceAccountCustomdataAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingOutputsGcsServiceAccountCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingOutputsGcsServiceAccountCustomdataAPIWithClient constructor for EncodingOutputsGcsServiceAccountCustomdataAPI that takes an APIClient as argument
func NewEncodingOutputsGcsServiceAccountCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingOutputsGcsServiceAccountCustomdataAPI {
    a := &EncodingOutputsGcsServiceAccountCustomdataAPI{apiClient: apiClient}
    return a
}

// Get Service Account based GCS Output Custom Data
func (api *EncodingOutputsGcsServiceAccountCustomdataAPI) Get(outputId string) (*model.CustomData, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["output_id"] = outputId
    }

    var responseModel model.CustomData
    err := api.apiClient.Get("/encoding/outputs/gcs-service-account/{output_id}/customData", nil, &responseModel, reqParams)
    return &responseModel, err
}

