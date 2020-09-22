package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingStatisticsEncodingsLiveStatisticsEventsAPI communicates with '/encoding/statistics/encodings/{encoding_id}/live-statistics/events' endpoints
type EncodingStatisticsEncodingsLiveStatisticsEventsAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingStatisticsEncodingsLiveStatisticsEventsAPI constructor for EncodingStatisticsEncodingsLiveStatisticsEventsAPI that takes options as argument
func NewEncodingStatisticsEncodingsLiveStatisticsEventsAPI(options ...apiclient.APIClientOption) (*EncodingStatisticsEncodingsLiveStatisticsEventsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingStatisticsEncodingsLiveStatisticsEventsAPIWithClient(apiClient), nil
}

// NewEncodingStatisticsEncodingsLiveStatisticsEventsAPIWithClient constructor for EncodingStatisticsEncodingsLiveStatisticsEventsAPI that takes an APIClient as argument
func NewEncodingStatisticsEncodingsLiveStatisticsEventsAPIWithClient(apiClient *apiclient.APIClient) *EncodingStatisticsEncodingsLiveStatisticsEventsAPI {
	a := &EncodingStatisticsEncodingsLiveStatisticsEventsAPI{apiClient: apiClient}
	return a
}

// List Events of Live Statistics from an Encoding
func (api *EncodingStatisticsEncodingsLiveStatisticsEventsAPI) List(encodingId string, queryParams ...func(*EncodingStatisticsEncodingsLiveStatisticsEventsAPIListQueryParams)) (*pagination.LiveEncodingStatsEventsListPagination, error) {
	queryParameters := &EncodingStatisticsEncodingsLiveStatisticsEventsAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.LiveEncodingStatsEventsListPagination
	err := api.apiClient.Get("/encoding/statistics/encodings/{encoding_id}/live-statistics/events", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingStatisticsEncodingsLiveStatisticsEventsAPIListQueryParams contains all query parameters for the List endpoint
type EncodingStatisticsEncodingsLiveStatisticsEventsAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingStatisticsEncodingsLiveStatisticsEventsAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
