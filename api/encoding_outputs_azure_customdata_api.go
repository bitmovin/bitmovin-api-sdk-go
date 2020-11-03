package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingOutputsAzureCustomdataAPI communicates with '/encoding/outputs/azure/{output_id}/customData' endpoints
type EncodingOutputsAzureCustomdataAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingOutputsAzureCustomdataAPI constructor for EncodingOutputsAzureCustomdataAPI that takes options as argument
func NewEncodingOutputsAzureCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingOutputsAzureCustomdataAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingOutputsAzureCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingOutputsAzureCustomdataAPIWithClient constructor for EncodingOutputsAzureCustomdataAPI that takes an APIClient as argument
func NewEncodingOutputsAzureCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingOutputsAzureCustomdataAPI {
    a := &EncodingOutputsAzureCustomdataAPI{apiClient: apiClient}
    return a
}

// Get Azure Output Custom Data
func (api *EncodingOutputsAzureCustomdataAPI) Get(outputId string) (*model.CustomData, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["output_id"] = outputId
    }

    var responseModel model.CustomData
    err := api.apiClient.Get("/encoding/outputs/azure/{output_id}/customData", nil, &responseModel, reqParams)
    return &responseModel, err
}

