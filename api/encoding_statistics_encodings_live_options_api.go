package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingStatisticsEncodingsLiveOptionsAPI communicates with '/encoding/statistics/encodings/live/options' endpoints
type EncodingStatisticsEncodingsLiveOptionsAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingStatisticsEncodingsLiveOptionsAPI constructor for EncodingStatisticsEncodingsLiveOptionsAPI that takes options as argument
func NewEncodingStatisticsEncodingsLiveOptionsAPI(options ...apiclient.APIClientOption) (*EncodingStatisticsEncodingsLiveOptionsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingStatisticsEncodingsLiveOptionsAPIWithClient(apiClient), nil
}

// NewEncodingStatisticsEncodingsLiveOptionsAPIWithClient constructor for EncodingStatisticsEncodingsLiveOptionsAPI that takes an APIClient as argument
func NewEncodingStatisticsEncodingsLiveOptionsAPIWithClient(apiClient *apiclient.APIClient) *EncodingStatisticsEncodingsLiveOptionsAPI {
	a := &EncodingStatisticsEncodingsLiveOptionsAPI{apiClient: apiClient}
	return a
}

// Get List live options encoding statistics for a given encoding
func (api *EncodingStatisticsEncodingsLiveOptionsAPI) Get(encodingId string) (*model.LiveEncodingOptionsStatistics, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	var responseModel model.LiveEncodingOptionsStatistics
	err := api.apiClient.Get("/encoding/statistics/encodings/live/{encoding_id}/options", nil, &responseModel, reqParams)
	return &responseModel, err
}

// ListByDateRange List live options encoding statistics within specific dates
func (api *EncodingStatisticsEncodingsLiveOptionsAPI) ListByDateRange(queryParams ...func(*EncodingStatisticsEncodingsLiveOptionsAPIListByDateRangeQueryParams)) (*model.LiveOptionsStatistics, error) {
	queryParameters := &EncodingStatisticsEncodingsLiveOptionsAPIListByDateRangeQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel model.LiveOptionsStatistics
	err := api.apiClient.Get("/encoding/statistics/encodings/live/options", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingStatisticsEncodingsLiveOptionsAPIListByDateRangeQueryParams contains all query parameters for the ListByDateRange endpoint
type EncodingStatisticsEncodingsLiveOptionsAPIListByDateRangeQueryParams struct {
	From   model.Date `query:"from"`
	To     model.Date `query:"to"`
	Offset int32      `query:"offset"`
	Limit  int32      `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingStatisticsEncodingsLiveOptionsAPIListByDateRangeQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
