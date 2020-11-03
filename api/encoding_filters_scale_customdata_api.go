package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingFiltersScaleCustomdataAPI communicates with '/encoding/filters/scale/{filter_id}/customData' endpoints
type EncodingFiltersScaleCustomdataAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingFiltersScaleCustomdataAPI constructor for EncodingFiltersScaleCustomdataAPI that takes options as argument
func NewEncodingFiltersScaleCustomdataAPI(options ...apiclient.APIClientOption) (*EncodingFiltersScaleCustomdataAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingFiltersScaleCustomdataAPIWithClient(apiClient), nil
}

// NewEncodingFiltersScaleCustomdataAPIWithClient constructor for EncodingFiltersScaleCustomdataAPI that takes an APIClient as argument
func NewEncodingFiltersScaleCustomdataAPIWithClient(apiClient *apiclient.APIClient) *EncodingFiltersScaleCustomdataAPI {
    a := &EncodingFiltersScaleCustomdataAPI{apiClient: apiClient}
    return a
}

// Get Scale Filter Custom Data
func (api *EncodingFiltersScaleCustomdataAPI) Get(filterId string) (*model.CustomData, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["filter_id"] = filterId
    }

    var responseModel model.CustomData
    err := api.apiClient.Get("/encoding/filters/scale/{filter_id}/customData", nil, &responseModel, reqParams)
    return &responseModel, err
}

