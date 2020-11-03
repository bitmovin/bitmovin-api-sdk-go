package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingStatisticsEncodingsLiveStatisticsSrtAPI communicates with '/encoding/statistics/encodings/{encoding_id}/live-statistics/srt' endpoints
type EncodingStatisticsEncodingsLiveStatisticsSrtAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingStatisticsEncodingsLiveStatisticsSrtAPI constructor for EncodingStatisticsEncodingsLiveStatisticsSrtAPI that takes options as argument
func NewEncodingStatisticsEncodingsLiveStatisticsSrtAPI(options ...apiclient.APIClientOption) (*EncodingStatisticsEncodingsLiveStatisticsSrtAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingStatisticsEncodingsLiveStatisticsSrtAPIWithClient(apiClient), nil
}

// NewEncodingStatisticsEncodingsLiveStatisticsSrtAPIWithClient constructor for EncodingStatisticsEncodingsLiveStatisticsSrtAPI that takes an APIClient as argument
func NewEncodingStatisticsEncodingsLiveStatisticsSrtAPIWithClient(apiClient *apiclient.APIClient) *EncodingStatisticsEncodingsLiveStatisticsSrtAPI {
    a := &EncodingStatisticsEncodingsLiveStatisticsSrtAPI{apiClient: apiClient}
    return a
}

// List Stream Infos of Live Statistics from an Encoding
func (api *EncodingStatisticsEncodingsLiveStatisticsSrtAPI) List(encodingId string, queryParams ...func(*EncodingStatisticsEncodingsLiveStatisticsSrtAPIListQueryParams)) (*pagination.SrtStatisticssListPagination, error) {
    queryParameters := &EncodingStatisticsEncodingsLiveStatisticsSrtAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
    }

    var responseModel pagination.SrtStatisticssListPagination
    err := api.apiClient.Get("/encoding/statistics/encodings/{encoding_id}/live-statistics/srt", nil, &responseModel, reqParams)
    return &responseModel, err
}
// ListBySrtInputId List Statistics For SRT Live Stream Input
func (api *EncodingStatisticsEncodingsLiveStatisticsSrtAPI) ListBySrtInputId(encodingId string, srtInputId string, queryParams ...func(*EncodingStatisticsEncodingsLiveStatisticsSrtAPIListBySrtInputIdQueryParams)) (*pagination.SrtStatisticssListBySrtInputIdPagination, error) {
    queryParameters := &EncodingStatisticsEncodingsLiveStatisticsSrtAPIListBySrtInputIdQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["srt_input_id"] = srtInputId
        params.QueryParams = queryParameters
    }

    var responseModel pagination.SrtStatisticssListBySrtInputIdPagination
    err := api.apiClient.Get("/encoding/statistics/encodings/{encoding_id}/live-statistics/srt/{srt_input_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingStatisticsEncodingsLiveStatisticsSrtAPIListQueryParams contains all query parameters for the List endpoint
type EncodingStatisticsEncodingsLiveStatisticsSrtAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingStatisticsEncodingsLiveStatisticsSrtAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


// EncodingStatisticsEncodingsLiveStatisticsSrtAPIListBySrtInputIdQueryParams contains all query parameters for the ListBySrtInputId endpoint
type EncodingStatisticsEncodingsLiveStatisticsSrtAPIListBySrtInputIdQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingStatisticsEncodingsLiveStatisticsSrtAPIListBySrtInputIdQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


