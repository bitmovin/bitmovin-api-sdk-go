package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingFiltersTextCustomdataAPI communicates with '/encoding/filters/text/{filter_id}/customData' endpoints
type EncodingFiltersTextCustomdataAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingFiltersTextCustomdataAPI constructor for EncodingFiltersTextCustomdataAPI that takes options as argument
func NewEncodingFiltersTextCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingFiltersTextCustomdataAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingFiltersTextCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingFiltersTextCustomdataAPIWithClient constructor for EncodingFiltersTextCustomdataAPI that takes an APIClient as argument
func NewEncodingFiltersTextCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingFiltersTextCustomdataAPI {
    a := &EncodingFiltersTextCustomdataAPI{apiClient: apiClient}
    return a
}

// Get Text Filter Custom Data
func (api *EncodingFiltersTextCustomdataAPI) Get(filterId string) (*model.CustomData, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["filter_id"] = filterId
    }

    var responseModel model.CustomData
    err := api.apiClient.Get("/encoding/filters/text/{filter_id}/customData", nil, &responseModel, reqParams)
    return &responseModel, err
}

