package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingFiltersDeinterlaceAPI communicates with '/encoding/filters/deinterlace' endpoints
type EncodingFiltersDeinterlaceAPI struct {
    apiClient *apiclient.APIClient

    // Customdata communicates with '/encoding/filters/deinterlace/{filter_id}/customData' endpoints
    Customdata *EncodingFiltersDeinterlaceCustomdataAPI

}

// NewEncodingFiltersDeinterlaceAPI constructor for EncodingFiltersDeinterlaceAPI that takes options as argument
func NewEncodingFiltersDeinterlaceAPI(options ...apiclient.APIClientOption) (*EncodingFiltersDeinterlaceAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingFiltersDeinterlaceAPIWithClient(apiClient), nil
}

// NewEncodingFiltersDeinterlaceAPIWithClient constructor for EncodingFiltersDeinterlaceAPI that takes an APIClient as argument
func NewEncodingFiltersDeinterlaceAPIWithClient(apiClient *apiclient.APIClient) *EncodingFiltersDeinterlaceAPI {
    a := &EncodingFiltersDeinterlaceAPI{apiClient: apiClient}
    a.Customdata = NewEncodingFiltersDeinterlaceCustomdataAPIWithClient(apiClient)

    return a
}

// Create Deinterlace Filter
func (api *EncodingFiltersDeinterlaceAPI) Create(deinterlaceFilter model.DeinterlaceFilter) (*model.DeinterlaceFilter, error) {
    reqParams := func(params *apiclient.RequestParams) {
    }

    var responseModel model.DeinterlaceFilter
    err := api.apiClient.Post("/encoding/filters/deinterlace", &deinterlaceFilter, &responseModel, reqParams)
    return &responseModel, err
}
// Delete Deinterlace Filter
func (api *EncodingFiltersDeinterlaceAPI) Delete(filterId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["filter_id"] = filterId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/filters/deinterlace/{filter_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get Deinterlace Filter Details
func (api *EncodingFiltersDeinterlaceAPI) Get(filterId string) (*model.DeinterlaceFilter, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["filter_id"] = filterId
    }

    var responseModel model.DeinterlaceFilter
    err := api.apiClient.Get("/encoding/filters/deinterlace/{filter_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List Deinterlace Filters
func (api *EncodingFiltersDeinterlaceAPI) List(queryParams ...func(*EncodingFiltersDeinterlaceAPIListQueryParams)) (*pagination.DeinterlaceFiltersListPagination, error) {
    queryParameters := &EncodingFiltersDeinterlaceAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel pagination.DeinterlaceFiltersListPagination
    err := api.apiClient.Get("/encoding/filters/deinterlace", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingFiltersDeinterlaceAPIListQueryParams contains all query parameters for the List endpoint
type EncodingFiltersDeinterlaceAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
    Name string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingFiltersDeinterlaceAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


