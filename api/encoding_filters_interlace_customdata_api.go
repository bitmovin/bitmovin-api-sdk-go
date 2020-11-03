package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingFiltersInterlaceCustomdataAPI communicates with '/encoding/filters/interlace/{filter_id}/customData' endpoints
type EncodingFiltersInterlaceCustomdataAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingFiltersInterlaceCustomdataAPI constructor for EncodingFiltersInterlaceCustomdataAPI that takes options as argument
func NewEncodingFiltersInterlaceCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingFiltersInterlaceCustomdataAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingFiltersInterlaceCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingFiltersInterlaceCustomdataAPIWithClient constructor for EncodingFiltersInterlaceCustomdataAPI that takes an APIClient as argument
func NewEncodingFiltersInterlaceCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingFiltersInterlaceCustomdataAPI {
    a := &EncodingFiltersInterlaceCustomdataAPI{apiClient: apiClient}
    return a
}

// Get Interlace Filter Custom Data
func (api *EncodingFiltersInterlaceCustomdataAPI) Get(filterId string) (*model.CustomData, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["filter_id"] = filterId
    }

    var responseModel model.CustomData
    err := api.apiClient.Get("/encoding/filters/interlace/{filter_id}/customData", nil, &responseModel, reqParams)
    return &responseModel, err
}

