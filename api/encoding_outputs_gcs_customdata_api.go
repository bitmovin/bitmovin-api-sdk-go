package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingOutputsGcsCustomdataAPI communicates with '/encoding/outputs/gcs/{output_id}/customData' endpoints
type EncodingOutputsGcsCustomdataAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingOutputsGcsCustomdataAPI constructor for EncodingOutputsGcsCustomdataAPI that takes options as argument
func NewEncodingOutputsGcsCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingOutputsGcsCustomdataAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingOutputsGcsCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingOutputsGcsCustomdataAPIWithClient constructor for EncodingOutputsGcsCustomdataAPI that takes an APIClient as argument
func NewEncodingOutputsGcsCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingOutputsGcsCustomdataAPI {
    a := &EncodingOutputsGcsCustomdataAPI{apiClient: apiClient}
    return a
}

// Get GCS Output Custom Data
func (api *EncodingOutputsGcsCustomdataAPI) Get(outputId string) (*model.CustomData, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["output_id"] = outputId
    }

    var responseModel model.CustomData
    err := api.apiClient.Get("/encoding/outputs/gcs/{output_id}/customData", nil, &responseModel, reqParams)
    return &responseModel, err
}

