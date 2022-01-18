package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingStatisticsEncodingsVodAPI communicates with '/encoding/statistics/encodings/vod' endpoints
type EncodingStatisticsEncodingsVodAPI struct {
	apiClient *apiclient.APIClient

	// Daily communicates with '/encoding/statistics/encodings/vod/daily/{from}/{to}' endpoints
	Daily *EncodingStatisticsEncodingsVodDailyAPI
}

// NewEncodingStatisticsEncodingsVodAPI constructor for EncodingStatisticsEncodingsVodAPI that takes options as argument
func NewEncodingStatisticsEncodingsVodAPI(options ...apiclient.APIClientOption) (*EncodingStatisticsEncodingsVodAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingStatisticsEncodingsVodAPIWithClient(apiClient), nil
}

// NewEncodingStatisticsEncodingsVodAPIWithClient constructor for EncodingStatisticsEncodingsVodAPI that takes an APIClient as argument
func NewEncodingStatisticsEncodingsVodAPIWithClient(apiClient *apiclient.APIClient) *EncodingStatisticsEncodingsVodAPI {
	a := &EncodingStatisticsEncodingsVodAPI{apiClient: apiClient}
	a.Daily = NewEncodingStatisticsEncodingsVodDailyAPIWithClient(apiClient)

	return a
}

// List VOD Encoding Statistics
func (api *EncodingStatisticsEncodingsVodAPI) List(queryParams ...func(*EncodingStatisticsEncodingsVodAPIListQueryParams)) (*pagination.EncodingStatisticsVodsListPagination, error) {
	queryParameters := &EncodingStatisticsEncodingsVodAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.EncodingStatisticsVodsListPagination
	err := api.apiClient.Get("/encoding/statistics/encodings/vod", nil, &responseModel, reqParams)
	return &responseModel, err
}

// ListByDateRange List VOD Encoding Statistics Within Specific Dates
func (api *EncodingStatisticsEncodingsVodAPI) ListByDateRange(from model.Date, to model.Date, queryParams ...func(*EncodingStatisticsEncodingsVodAPIListByDateRangeQueryParams)) (*pagination.EncodingStatisticsVodsListByDateRangePagination, error) {
	queryParameters := &EncodingStatisticsEncodingsVodAPIListByDateRangeQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["from"] = from.String()
		params.PathParams["to"] = to.String()
		params.QueryParams = queryParameters
	}

	var responseModel pagination.EncodingStatisticsVodsListByDateRangePagination
	err := api.apiClient.Get("/encoding/statistics/encodings/vod/{from}/{to}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingStatisticsEncodingsVodAPIListQueryParams contains all query parameters for the List endpoint
type EncodingStatisticsEncodingsVodAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingStatisticsEncodingsVodAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}

// EncodingStatisticsEncodingsVodAPIListByDateRangeQueryParams contains all query parameters for the ListByDateRange endpoint
type EncodingStatisticsEncodingsVodAPIListByDateRangeQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingStatisticsEncodingsVodAPIListByDateRangeQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
