package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingFiltersCropAPI communicates with '/encoding/filters/crop' endpoints
type EncodingFiltersCropAPI struct {
    apiClient *apiclient.APIClient

    // Customdata communicates with '/encoding/filters/crop/{filter_id}/customData' endpoints
    Customdata *EncodingFiltersCropCustomdataAPI

}

// NewEncodingFiltersCropAPI constructor for EncodingFiltersCropAPI that takes options as argument
func NewEncodingFiltersCropAPI(options ...apiclient.APIClientOption) (*EncodingFiltersCropAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingFiltersCropAPIWithClient(apiClient), nil
}

// NewEncodingFiltersCropAPIWithClient constructor for EncodingFiltersCropAPI that takes an APIClient as argument
func NewEncodingFiltersCropAPIWithClient(apiClient *apiclient.APIClient) *EncodingFiltersCropAPI {
    a := &EncodingFiltersCropAPI{apiClient: apiClient}
    a.Customdata = NewEncodingFiltersCropCustomdataAPIWithClient(apiClient)

    return a
}

// Create Crop Filter
func (api *EncodingFiltersCropAPI) Create(cropFilter model.CropFilter) (*model.CropFilter, error) {
    reqParams := func(params *apiclient.RequestParams) {
    }

    var responseModel model.CropFilter
    err := api.apiClient.Post("/encoding/filters/crop", &cropFilter, &responseModel, reqParams)
    return &responseModel, err
}
// Delete Crop Filter
func (api *EncodingFiltersCropAPI) Delete(filterId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["filter_id"] = filterId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/filters/crop/{filter_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get Crop Filter Details
func (api *EncodingFiltersCropAPI) Get(filterId string) (*model.CropFilter, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["filter_id"] = filterId
    }

    var responseModel model.CropFilter
    err := api.apiClient.Get("/encoding/filters/crop/{filter_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List Crop Filters
func (api *EncodingFiltersCropAPI) List(queryParams ...func(*EncodingFiltersCropAPIListQueryParams)) (*pagination.CropFiltersListPagination, error) {
    queryParameters := &EncodingFiltersCropAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel pagination.CropFiltersListPagination
    err := api.apiClient.Get("/encoding/filters/crop", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingFiltersCropAPIListQueryParams contains all query parameters for the List endpoint
type EncodingFiltersCropAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
    Name string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingFiltersCropAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


