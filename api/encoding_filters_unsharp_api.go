package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingFiltersUnsharpAPI communicates with '/encoding/filters/unsharp' endpoints
type EncodingFiltersUnsharpAPI struct {
    apiClient *apiclient.APIClient

    // Customdata communicates with '/encoding/filters/unsharp/{filter_id}/customData' endpoints
    Customdata *EncodingFiltersUnsharpCustomdataAPI

}

// NewEncodingFiltersUnsharpAPI constructor for EncodingFiltersUnsharpAPI that takes options as argument
func NewEncodingFiltersUnsharpAPI(options ...apiclient.APIClientOption) (*EncodingFiltersUnsharpAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingFiltersUnsharpAPIWithClient(apiClient), nil
}

// NewEncodingFiltersUnsharpAPIWithClient constructor for EncodingFiltersUnsharpAPI that takes an APIClient as argument
func NewEncodingFiltersUnsharpAPIWithClient(apiClient *apiclient.APIClient) *EncodingFiltersUnsharpAPI {
    a := &EncodingFiltersUnsharpAPI{apiClient: apiClient}
    a.Customdata = NewEncodingFiltersUnsharpCustomdataAPIWithClient(apiClient)

    return a
}

// Create Unsharp Filter
func (api *EncodingFiltersUnsharpAPI) Create(unsharpFilter model.UnsharpFilter) (*model.UnsharpFilter, error) {
    reqParams := func(params *apiclient.RequestParams) {
    }

    var responseModel model.UnsharpFilter
    err := api.apiClient.Post("/encoding/filters/unsharp", &unsharpFilter, &responseModel, reqParams)
    return &responseModel, err
}
// Delete Unsharp Filter
func (api *EncodingFiltersUnsharpAPI) Delete(filterId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["filter_id"] = filterId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/filters/unsharp/{filter_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get Unsharp Filter Details
func (api *EncodingFiltersUnsharpAPI) Get(filterId string) (*model.UnsharpFilter, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["filter_id"] = filterId
    }

    var responseModel model.UnsharpFilter
    err := api.apiClient.Get("/encoding/filters/unsharp/{filter_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List Unsharp Filters
func (api *EncodingFiltersUnsharpAPI) List(queryParams ...func(*EncodingFiltersUnsharpAPIListQueryParams)) (*pagination.UnsharpFiltersListPagination, error) {
    queryParameters := &EncodingFiltersUnsharpAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel pagination.UnsharpFiltersListPagination
    err := api.apiClient.Get("/encoding/filters/unsharp", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingFiltersUnsharpAPIListQueryParams contains all query parameters for the List endpoint
type EncodingFiltersUnsharpAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
    Name string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingFiltersUnsharpAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


