package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingStatisticsEncodingsLiveAPI communicates with '/encoding/statistics/encodings/live' endpoints
type EncodingStatisticsEncodingsLiveAPI struct {
    apiClient *apiclient.APIClient

    // Daily communicates with '/encoding/statistics/encodings/live/daily/{from}/{to}' endpoints
    Daily *EncodingStatisticsEncodingsLiveDailyAPI

}

// NewEncodingStatisticsEncodingsLiveAPI constructor for EncodingStatisticsEncodingsLiveAPI that takes options as argument
func NewEncodingStatisticsEncodingsLiveAPI(options ...apiclient.APIClientOption) (*EncodingStatisticsEncodingsLiveAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingStatisticsEncodingsLiveAPIWithClient(apiClient), nil
}

// NewEncodingStatisticsEncodingsLiveAPIWithClient constructor for EncodingStatisticsEncodingsLiveAPI that takes an APIClient as argument
func NewEncodingStatisticsEncodingsLiveAPIWithClient(apiClient *apiclient.APIClient) *EncodingStatisticsEncodingsLiveAPI {
    a := &EncodingStatisticsEncodingsLiveAPI{apiClient: apiClient}
    a.Daily = NewEncodingStatisticsEncodingsLiveDailyAPIWithClient(apiClient)

    return a
}

// List Live Encoding Statistics
func (api *EncodingStatisticsEncodingsLiveAPI) List(queryParams ...func(*EncodingStatisticsEncodingsLiveAPIListQueryParams)) (*pagination.EncodingStatisticsLivesListPagination, error) {
    queryParameters := &EncodingStatisticsEncodingsLiveAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel pagination.EncodingStatisticsLivesListPagination
    err := api.apiClient.Get("/encoding/statistics/encodings/live", nil, &responseModel, reqParams)
    return &responseModel, err
}
// ListByDateRange List live encoding statistics within specific dates
func (api *EncodingStatisticsEncodingsLiveAPI) ListByDateRange(from model.Date, to model.Date, queryParams ...func(*EncodingStatisticsEncodingsLiveAPIListByDateRangeQueryParams)) (*pagination.EncodingStatisticsLivesListByDateRangePagination, error) {
    queryParameters := &EncodingStatisticsEncodingsLiveAPIListByDateRangeQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["from"] = from.String()
        params.PathParams["to"] = to.String()
        params.QueryParams = queryParameters
    }

    var responseModel pagination.EncodingStatisticsLivesListByDateRangePagination
    err := api.apiClient.Get("/encoding/statistics/encodings/live/{from}/{to}", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingStatisticsEncodingsLiveAPIListQueryParams contains all query parameters for the List endpoint
type EncodingStatisticsEncodingsLiveAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingStatisticsEncodingsLiveAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


// EncodingStatisticsEncodingsLiveAPIListByDateRangeQueryParams contains all query parameters for the ListByDateRange endpoint
type EncodingStatisticsEncodingsLiveAPIListByDateRangeQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingStatisticsEncodingsLiveAPIListByDateRangeQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


