package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingOutputsLocalCustomdataAPI communicates with '/encoding/outputs/local/{output_id}/customData' endpoints
type EncodingOutputsLocalCustomdataAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingOutputsLocalCustomdataAPI constructor for EncodingOutputsLocalCustomdataAPI that takes options as argument
func NewEncodingOutputsLocalCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingOutputsLocalCustomdataAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingOutputsLocalCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingOutputsLocalCustomdataAPIWithClient constructor for EncodingOutputsLocalCustomdataAPI that takes an APIClient as argument
func NewEncodingOutputsLocalCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingOutputsLocalCustomdataAPI {
    a := &EncodingOutputsLocalCustomdataAPI{apiClient: apiClient}
    return a
}

// Get Local Output Custom Data
func (api *EncodingOutputsLocalCustomdataAPI) Get(outputId string) (*model.CustomData, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["output_id"] = outputId
    }

    var responseModel model.CustomData
    err := api.apiClient.Get("/encoding/outputs/local/{output_id}/customData", nil, &responseModel, reqParams)
    return &responseModel, err
}

