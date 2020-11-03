package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingFiltersDenoiseHqdn3dAPI communicates with '/encoding/filters/denoise-hqdn3d' endpoints
type EncodingFiltersDenoiseHqdn3dAPI struct {
    apiClient *apiclient.APIClient

    // Customdata communicates with '/encoding/filters/denoise-hqdn3d/{filter_id}/customData' endpoints
    Customdata *EncodingFiltersDenoiseHqdn3dCustomdataAPI

}

// NewEncodingFiltersDenoiseHqdn3dAPI constructor for EncodingFiltersDenoiseHqdn3dAPI that takes options as argument
func NewEncodingFiltersDenoiseHqdn3dAPI(options ...apiclient.APIClientOption) (*EncodingFiltersDenoiseHqdn3dAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingFiltersDenoiseHqdn3dAPIWithClient(apiClient), nil
}

// NewEncodingFiltersDenoiseHqdn3dAPIWithClient constructor for EncodingFiltersDenoiseHqdn3dAPI that takes an APIClient as argument
func NewEncodingFiltersDenoiseHqdn3dAPIWithClient(apiClient *apiclient.APIClient) *EncodingFiltersDenoiseHqdn3dAPI {
    a := &EncodingFiltersDenoiseHqdn3dAPI{apiClient: apiClient}
    a.Customdata = NewEncodingFiltersDenoiseHqdn3dCustomdataAPIWithClient(apiClient)

    return a
}

// Create Denoise hqdn3d Filter
func (api *EncodingFiltersDenoiseHqdn3dAPI) Create(denoiseHqdn3dFilter model.DenoiseHqdn3dFilter) (*model.DenoiseHqdn3dFilter, error) {
    reqParams := func(params *apiclient.RequestParams) {
    }

    var responseModel model.DenoiseHqdn3dFilter
    err := api.apiClient.Post("/encoding/filters/denoise-hqdn3d", &denoiseHqdn3dFilter, &responseModel, reqParams)
    return &responseModel, err
}
// Delete Denoise hqdn3d Filter
func (api *EncodingFiltersDenoiseHqdn3dAPI) Delete(filterId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["filter_id"] = filterId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/filters/denoise-hqdn3d/{filter_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get Denoise hqdn3d Filter Details
func (api *EncodingFiltersDenoiseHqdn3dAPI) Get(filterId string) (*model.DenoiseHqdn3dFilter, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["filter_id"] = filterId
    }

    var responseModel model.DenoiseHqdn3dFilter
    err := api.apiClient.Get("/encoding/filters/denoise-hqdn3d/{filter_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List Denoise hqdn3d Filters
func (api *EncodingFiltersDenoiseHqdn3dAPI) List(queryParams ...func(*EncodingFiltersDenoiseHqdn3dAPIListQueryParams)) (*pagination.DenoiseHqdn3dFiltersListPagination, error) {
    queryParameters := &EncodingFiltersDenoiseHqdn3dAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel pagination.DenoiseHqdn3dFiltersListPagination
    err := api.apiClient.Get("/encoding/filters/denoise-hqdn3d", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingFiltersDenoiseHqdn3dAPIListQueryParams contains all query parameters for the List endpoint
type EncodingFiltersDenoiseHqdn3dAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
    Name string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingFiltersDenoiseHqdn3dAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


