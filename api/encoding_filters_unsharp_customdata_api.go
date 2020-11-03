package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingFiltersUnsharpCustomdataAPI communicates with '/encoding/filters/unsharp/{filter_id}/customData' endpoints
type EncodingFiltersUnsharpCustomdataAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingFiltersUnsharpCustomdataAPI constructor for EncodingFiltersUnsharpCustomdataAPI that takes options as argument
func NewEncodingFiltersUnsharpCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingFiltersUnsharpCustomdataAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingFiltersUnsharpCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingFiltersUnsharpCustomdataAPIWithClient constructor for EncodingFiltersUnsharpCustomdataAPI that takes an APIClient as argument
func NewEncodingFiltersUnsharpCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingFiltersUnsharpCustomdataAPI {
    a := &EncodingFiltersUnsharpCustomdataAPI{apiClient: apiClient}
    return a
}

// Get Unsharp Filter Custom Data
func (api *EncodingFiltersUnsharpCustomdataAPI) Get(filterId string) (*model.CustomData, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["filter_id"] = filterId
    }

    var responseModel model.CustomData
    err := api.apiClient.Get("/encoding/filters/unsharp/{filter_id}/customData", nil, &responseModel, reqParams)
    return &responseModel, err
}

