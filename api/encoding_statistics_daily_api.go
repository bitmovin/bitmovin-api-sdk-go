package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingStatisticsDailyAPI communicates with '/encoding/statistics/daily' endpoints
type EncodingStatisticsDailyAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingStatisticsDailyAPI constructor for EncodingStatisticsDailyAPI that takes options as argument
func NewEncodingStatisticsDailyAPI(options ...apiclient.APIClientOption) (*EncodingStatisticsDailyAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingStatisticsDailyAPIWithClient(apiClient), nil
}

// NewEncodingStatisticsDailyAPIWithClient constructor for EncodingStatisticsDailyAPI that takes an APIClient as argument
func NewEncodingStatisticsDailyAPIWithClient(apiClient *apiclient.APIClient) *EncodingStatisticsDailyAPI {
    a := &EncodingStatisticsDailyAPI{apiClient: apiClient}
    return a
}

// List Daily Statistics
func (api *EncodingStatisticsDailyAPI) List(queryParams ...func(*EncodingStatisticsDailyAPIListQueryParams)) (*pagination.DailyStatisticssListPagination, error) {
    queryParameters := &EncodingStatisticsDailyAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel pagination.DailyStatisticssListPagination
    err := api.apiClient.Get("/encoding/statistics/daily", nil, &responseModel, reqParams)
    return &responseModel, err
}
// ListByDateRange List daily statistics within specific dates
func (api *EncodingStatisticsDailyAPI) ListByDateRange(from model.Date, to model.Date, queryParams ...func(*EncodingStatisticsDailyAPIListByDateRangeQueryParams)) (*pagination.DailyStatisticssListByDateRangePagination, error) {
    queryParameters := &EncodingStatisticsDailyAPIListByDateRangeQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["from"] = from.String()
        params.PathParams["to"] = to.String()
        params.QueryParams = queryParameters
    }

    var responseModel pagination.DailyStatisticssListByDateRangePagination
    err := api.apiClient.Get("/encoding/statistics/daily/{from}/{to}", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingStatisticsDailyAPIListQueryParams contains all query parameters for the List endpoint
type EncodingStatisticsDailyAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingStatisticsDailyAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


// EncodingStatisticsDailyAPIListByDateRangeQueryParams contains all query parameters for the ListByDateRange endpoint
type EncodingStatisticsDailyAPIListByDateRangeQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingStatisticsDailyAPIListByDateRangeQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


