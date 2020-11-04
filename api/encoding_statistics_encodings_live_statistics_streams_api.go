package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingStatisticsEncodingsLiveStatisticsStreamsAPI communicates with '/encoding/statistics/encodings/{encoding_id}/live-statistics/streams' endpoints
type EncodingStatisticsEncodingsLiveStatisticsStreamsAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingStatisticsEncodingsLiveStatisticsStreamsAPI constructor for EncodingStatisticsEncodingsLiveStatisticsStreamsAPI that takes options as argument
func NewEncodingStatisticsEncodingsLiveStatisticsStreamsAPI(options ...apiclient.APIClientOption) (*EncodingStatisticsEncodingsLiveStatisticsStreamsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingStatisticsEncodingsLiveStatisticsStreamsAPIWithClient(apiClient), nil
}

// NewEncodingStatisticsEncodingsLiveStatisticsStreamsAPIWithClient constructor for EncodingStatisticsEncodingsLiveStatisticsStreamsAPI that takes an APIClient as argument
func NewEncodingStatisticsEncodingsLiveStatisticsStreamsAPIWithClient(apiClient *apiclient.APIClient) *EncodingStatisticsEncodingsLiveStatisticsStreamsAPI {
	a := &EncodingStatisticsEncodingsLiveStatisticsStreamsAPI{apiClient: apiClient}
	return a
}

// List Stream Infos of Live Statistics from an Encoding
func (api *EncodingStatisticsEncodingsLiveStatisticsStreamsAPI) List(encodingId string, queryParams ...func(*EncodingStatisticsEncodingsLiveStatisticsStreamsAPIListQueryParams)) (*pagination.StreamInfossListPagination, error) {
	queryParameters := &EncodingStatisticsEncodingsLiveStatisticsStreamsAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.StreamInfossListPagination
	err := api.apiClient.Get("/encoding/statistics/encodings/{encoding_id}/live-statistics/streams", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingStatisticsEncodingsLiveStatisticsStreamsAPIListQueryParams contains all query parameters for the List endpoint
type EncodingStatisticsEncodingsLiveStatisticsStreamsAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingStatisticsEncodingsLiveStatisticsStreamsAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
