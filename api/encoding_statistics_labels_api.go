package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingStatisticsLabelsAPI communicates with '/encoding/statistics/labels/' endpoints
type EncodingStatisticsLabelsAPI struct {
	apiClient *apiclient.APIClient

	// Daily communicates with '/encoding/statistics/labels/daily' endpoints
	Daily *EncodingStatisticsLabelsDailyAPI
}

// NewEncodingStatisticsLabelsAPI constructor for EncodingStatisticsLabelsAPI that takes options as argument
func NewEncodingStatisticsLabelsAPI(options ...apiclient.APIClientOption) (*EncodingStatisticsLabelsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingStatisticsLabelsAPIWithClient(apiClient), nil
}

// NewEncodingStatisticsLabelsAPIWithClient constructor for EncodingStatisticsLabelsAPI that takes an APIClient as argument
func NewEncodingStatisticsLabelsAPIWithClient(apiClient *apiclient.APIClient) *EncodingStatisticsLabelsAPI {
	a := &EncodingStatisticsLabelsAPI{apiClient: apiClient}
	a.Daily = NewEncodingStatisticsLabelsDailyAPIWithClient(apiClient)

	return a
}

// List Get Statistics per Label
func (api *EncodingStatisticsLabelsAPI) List(queryParams ...func(*EncodingStatisticsLabelsAPIListQueryParams)) (*pagination.StatisticsPerLabelsListPagination, error) {
	queryParameters := &EncodingStatisticsLabelsAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.StatisticsPerLabelsListPagination
	err := api.apiClient.Get("/encoding/statistics/labels/", nil, &responseModel, reqParams)
	return &responseModel, err
}

// ListByDateRange Get statistics per label within specific dates
func (api *EncodingStatisticsLabelsAPI) ListByDateRange(from model.Date, to model.Date, queryParams ...func(*EncodingStatisticsLabelsAPIListByDateRangeQueryParams)) (*pagination.StatisticsPerLabelsListByDateRangePagination, error) {
	queryParameters := &EncodingStatisticsLabelsAPIListByDateRangeQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["from"] = from.String()
		params.PathParams["to"] = to.String()
		params.QueryParams = queryParameters
	}

	var responseModel pagination.StatisticsPerLabelsListByDateRangePagination
	err := api.apiClient.Get("/encoding/statistics/labels/{from}/{to}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingStatisticsLabelsAPIListQueryParams contains all query parameters for the List endpoint
type EncodingStatisticsLabelsAPIListQueryParams struct {
	Offset int32  `query:"offset"`
	Limit  int32  `query:"limit"`
	Labels string `query:"labels"`
}

// Params will return a map of query parameters
func (q *EncodingStatisticsLabelsAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}

// EncodingStatisticsLabelsAPIListByDateRangeQueryParams contains all query parameters for the ListByDateRange endpoint
type EncodingStatisticsLabelsAPIListByDateRangeQueryParams struct {
	Offset int32  `query:"offset"`
	Limit  int32  `query:"limit"`
	Labels string `query:"labels"`
}

// Params will return a map of query parameters
func (q *EncodingStatisticsLabelsAPIListByDateRangeQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
