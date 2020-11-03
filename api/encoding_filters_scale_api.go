package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingFiltersScaleAPI communicates with '/encoding/filters/scale' endpoints
type EncodingFiltersScaleAPI struct {
    apiClient *apiclient.APIClient

    // Customdata communicates with '/encoding/filters/scale/{filter_id}/customData' endpoints
    Customdata *EncodingFiltersScaleCustomdataAPI

}

// NewEncodingFiltersScaleAPI constructor for EncodingFiltersScaleAPI that takes options as argument
func NewEncodingFiltersScaleAPI(options ...apiclient.APIClientOption) (*EncodingFiltersScaleAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingFiltersScaleAPIWithClient(apiClient), nil
}

// NewEncodingFiltersScaleAPIWithClient constructor for EncodingFiltersScaleAPI that takes an APIClient as argument
func NewEncodingFiltersScaleAPIWithClient(apiClient *apiclient.APIClient) *EncodingFiltersScaleAPI {
    a := &EncodingFiltersScaleAPI{apiClient: apiClient}
    a.Customdata = NewEncodingFiltersScaleCustomdataAPIWithClient(apiClient)

    return a
}

// Create Scale Filter
func (api *EncodingFiltersScaleAPI) Create(scaleFilter model.ScaleFilter) (*model.ScaleFilter, error) {
    reqParams := func(params *apiclient.RequestParams) {
    }

    var responseModel model.ScaleFilter
    err := api.apiClient.Post("/encoding/filters/scale", &scaleFilter, &responseModel, reqParams)
    return &responseModel, err
}
// Delete Scale Filter
func (api *EncodingFiltersScaleAPI) Delete(filterId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["filter_id"] = filterId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/filters/scale/{filter_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get Scale Filter Details
func (api *EncodingFiltersScaleAPI) Get(filterId string) (*model.ScaleFilter, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["filter_id"] = filterId
    }

    var responseModel model.ScaleFilter
    err := api.apiClient.Get("/encoding/filters/scale/{filter_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List Scale Filters
func (api *EncodingFiltersScaleAPI) List(queryParams ...func(*EncodingFiltersScaleAPIListQueryParams)) (*pagination.ScaleFiltersListPagination, error) {
    queryParameters := &EncodingFiltersScaleAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel pagination.ScaleFiltersListPagination
    err := api.apiClient.Get("/encoding/filters/scale", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingFiltersScaleAPIListQueryParams contains all query parameters for the List endpoint
type EncodingFiltersScaleAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
    Name string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingFiltersScaleAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


