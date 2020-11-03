package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingFiltersConformAPI communicates with '/encoding/filters/conform' endpoints
type EncodingFiltersConformAPI struct {
    apiClient *apiclient.APIClient

    // Customdata communicates with '/encoding/filters/conform/{filter_id}/customData' endpoints
    Customdata *EncodingFiltersConformCustomdataAPI

}

// NewEncodingFiltersConformAPI constructor for EncodingFiltersConformAPI that takes options as argument
func NewEncodingFiltersConformAPI(options ...apiclient.APIClientOption) (*EncodingFiltersConformAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingFiltersConformAPIWithClient(apiClient), nil
}

// NewEncodingFiltersConformAPIWithClient constructor for EncodingFiltersConformAPI that takes an APIClient as argument
func NewEncodingFiltersConformAPIWithClient(apiClient *apiclient.APIClient) *EncodingFiltersConformAPI {
    a := &EncodingFiltersConformAPI{apiClient: apiClient}
    a.Customdata = NewEncodingFiltersConformCustomdataAPIWithClient(apiClient)

    return a
}

// Create Conform Filter. Keeps all the frames of the input. The playback time of the output will be slower or faster.
func (api *EncodingFiltersConformAPI) Create(conformFilter model.ConformFilter) (*model.ConformFilter, error) {
    reqParams := func(params *apiclient.RequestParams) {
    }

    var responseModel model.ConformFilter
    err := api.apiClient.Post("/encoding/filters/conform", &conformFilter, &responseModel, reqParams)
    return &responseModel, err
}
// Delete Conform Filter
func (api *EncodingFiltersConformAPI) Delete(filterId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["filter_id"] = filterId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/filters/conform/{filter_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get Conform Filter Details
func (api *EncodingFiltersConformAPI) Get(filterId string) (*model.ConformFilter, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["filter_id"] = filterId
    }

    var responseModel model.ConformFilter
    err := api.apiClient.Get("/encoding/filters/conform/{filter_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List Conform Filters
func (api *EncodingFiltersConformAPI) List(queryParams ...func(*EncodingFiltersConformAPIListQueryParams)) (*pagination.ConformFiltersListPagination, error) {
    queryParameters := &EncodingFiltersConformAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel pagination.ConformFiltersListPagination
    err := api.apiClient.Get("/encoding/filters/conform", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingFiltersConformAPIListQueryParams contains all query parameters for the List endpoint
type EncodingFiltersConformAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
    Name string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingFiltersConformAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


