package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingStatisticsAPI communicates with '/encoding/statistics' endpoints
type EncodingStatisticsAPI struct {
    apiClient *apiclient.APIClient

    // Daily communicates with '/encoding/statistics/daily' endpoints
    Daily *EncodingStatisticsDailyAPI
    // Encodings communicates with '/encoding/statistics/encodings/{encoding_id}' endpoints
    Encodings *EncodingStatisticsEncodingsAPI
    // Labels communicates with '/encoding/statistics/labels/' endpoints
    Labels *EncodingStatisticsLabelsAPI

}

// NewEncodingStatisticsAPI constructor for EncodingStatisticsAPI that takes options as argument
func NewEncodingStatisticsAPI(options ...apiclient.APIClientOption) (*EncodingStatisticsAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingStatisticsAPIWithClient(apiClient), nil
}

// NewEncodingStatisticsAPIWithClient constructor for EncodingStatisticsAPI that takes an APIClient as argument
func NewEncodingStatisticsAPIWithClient(apiClient *apiclient.APIClient) *EncodingStatisticsAPI {
    a := &EncodingStatisticsAPI{apiClient: apiClient}
    a.Daily = NewEncodingStatisticsDailyAPIWithClient(apiClient)
    a.Encodings = NewEncodingStatisticsEncodingsAPIWithClient(apiClient)
    a.Labels = NewEncodingStatisticsLabelsAPIWithClient(apiClient)

    return a
}

// Get Show Overall Statistics
func (api *EncodingStatisticsAPI) Get() (*model.Statistics, error) {
    reqParams := func(params *apiclient.RequestParams) {
    }

    var responseModel model.Statistics
    err := api.apiClient.Get("/encoding/statistics", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List Show Overall Statistics Within Specific Dates
func (api *EncodingStatisticsAPI) List(from model.Date, to model.Date, queryParams ...func(*EncodingStatisticsAPIListQueryParams)) (*pagination.StatisticssListPagination, error) {
    queryParameters := &EncodingStatisticsAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["from"] = from.String()
        params.PathParams["to"] = to.String()
        params.QueryParams = queryParameters
    }

    var responseModel pagination.StatisticssListPagination
    err := api.apiClient.Get("/encoding/statistics/{from}/{to}", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingStatisticsAPIListQueryParams contains all query parameters for the List endpoint
type EncodingStatisticsAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingStatisticsAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


