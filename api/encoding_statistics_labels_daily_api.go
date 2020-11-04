package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingStatisticsLabelsDailyAPI communicates with '/encoding/statistics/labels/daily' endpoints
type EncodingStatisticsLabelsDailyAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingStatisticsLabelsDailyAPI constructor for EncodingStatisticsLabelsDailyAPI that takes options as argument
func NewEncodingStatisticsLabelsDailyAPI(options ...apiclient.APIClientOption) (*EncodingStatisticsLabelsDailyAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingStatisticsLabelsDailyAPIWithClient(apiClient), nil
}

// NewEncodingStatisticsLabelsDailyAPIWithClient constructor for EncodingStatisticsLabelsDailyAPI that takes an APIClient as argument
func NewEncodingStatisticsLabelsDailyAPIWithClient(apiClient *apiclient.APIClient) *EncodingStatisticsLabelsDailyAPI {
	a := &EncodingStatisticsLabelsDailyAPI{apiClient: apiClient}
	return a
}

// List Get Daily Statistics per Label
func (api *EncodingStatisticsLabelsDailyAPI) List(queryParams ...func(*EncodingStatisticsLabelsDailyAPIListQueryParams)) (*pagination.DailyStatisticsPerLabelsListPagination, error) {
	queryParameters := &EncodingStatisticsLabelsDailyAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.DailyStatisticsPerLabelsListPagination
	err := api.apiClient.Get("/encoding/statistics/labels/daily", nil, &responseModel, reqParams)
	return &responseModel, err
}

// ListByDateRange Get daily statistics per label within specific dates
func (api *EncodingStatisticsLabelsDailyAPI) ListByDateRange(from model.Date, to model.Date, queryParams ...func(*EncodingStatisticsLabelsDailyAPIListByDateRangeQueryParams)) (*pagination.DailyStatisticsPerLabelsListByDateRangePagination, error) {
	queryParameters := &EncodingStatisticsLabelsDailyAPIListByDateRangeQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["from"] = from.String()
		params.PathParams["to"] = to.String()
		params.QueryParams = queryParameters
	}

	var responseModel pagination.DailyStatisticsPerLabelsListByDateRangePagination
	err := api.apiClient.Get("/encoding/statistics/labels/daily/{from}/{to}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingStatisticsLabelsDailyAPIListQueryParams contains all query parameters for the List endpoint
type EncodingStatisticsLabelsDailyAPIListQueryParams struct {
	Offset int64  `query:"offset"`
	Limit  int64  `query:"limit"`
	Labels string `query:"labels"`
}

// Params will return a map of query parameters
func (q *EncodingStatisticsLabelsDailyAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}

// EncodingStatisticsLabelsDailyAPIListByDateRangeQueryParams contains all query parameters for the ListByDateRange endpoint
type EncodingStatisticsLabelsDailyAPIListByDateRangeQueryParams struct {
	Offset int64  `query:"offset"`
	Limit  int64  `query:"limit"`
	Labels string `query:"labels"`
}

// Params will return a map of query parameters
func (q *EncodingStatisticsLabelsDailyAPIListByDateRangeQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
