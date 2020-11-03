package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingStatisticsEncodingsAPI communicates with '/encoding/statistics/encodings/{encoding_id}' endpoints
type EncodingStatisticsEncodingsAPI struct {
    apiClient *apiclient.APIClient

    // Live communicates with '/encoding/statistics/encodings/live' endpoints
    Live *EncodingStatisticsEncodingsLiveAPI
    // Vod communicates with '/encoding/statistics/encodings/vod' endpoints
    Vod *EncodingStatisticsEncodingsVodAPI
    // LiveStatistics communicates with '/encoding/statistics/encodings/{encoding_id}/live-statistics' endpoints
    LiveStatistics *EncodingStatisticsEncodingsLiveStatisticsAPI

}

// NewEncodingStatisticsEncodingsAPI constructor for EncodingStatisticsEncodingsAPI that takes options as argument
func NewEncodingStatisticsEncodingsAPI(options ...apiclient.APIClientOption) (*EncodingStatisticsEncodingsAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingStatisticsEncodingsAPIWithClient(apiClient), nil
}

// NewEncodingStatisticsEncodingsAPIWithClient constructor for EncodingStatisticsEncodingsAPI that takes an APIClient as argument
func NewEncodingStatisticsEncodingsAPIWithClient(apiClient *apiclient.APIClient) *EncodingStatisticsEncodingsAPI {
    a := &EncodingStatisticsEncodingsAPI{apiClient: apiClient}
    a.Live = NewEncodingStatisticsEncodingsLiveAPIWithClient(apiClient)
    a.Vod = NewEncodingStatisticsEncodingsVodAPIWithClient(apiClient)
    a.LiveStatistics = NewEncodingStatisticsEncodingsLiveStatisticsAPIWithClient(apiClient)

    return a
}

// Get Statistics from an Encoding
func (api *EncodingStatisticsEncodingsAPI) Get(encodingId string) (*model.EncodingStats, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel model.EncodingStats
    err := api.apiClient.Get("/encoding/statistics/encodings/{encoding_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}

