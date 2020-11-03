package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingStatisticsEncodingsLiveStatisticsAPI communicates with '/encoding/statistics/encodings/{encoding_id}/live-statistics' endpoints
type EncodingStatisticsEncodingsLiveStatisticsAPI struct {
    apiClient *apiclient.APIClient

    // Events communicates with '/encoding/statistics/encodings/{encoding_id}/live-statistics/events' endpoints
    Events *EncodingStatisticsEncodingsLiveStatisticsEventsAPI
    // Streams communicates with '/encoding/statistics/encodings/{encoding_id}/live-statistics/streams' endpoints
    Streams *EncodingStatisticsEncodingsLiveStatisticsStreamsAPI
    // Srt communicates with '/encoding/statistics/encodings/{encoding_id}/live-statistics/srt' endpoints
    Srt *EncodingStatisticsEncodingsLiveStatisticsSrtAPI

}

// NewEncodingStatisticsEncodingsLiveStatisticsAPI constructor for EncodingStatisticsEncodingsLiveStatisticsAPI that takes options as argument
func NewEncodingStatisticsEncodingsLiveStatisticsAPI(options ...apiclient.APIClientOption) (*EncodingStatisticsEncodingsLiveStatisticsAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingStatisticsEncodingsLiveStatisticsAPIWithClient(apiClient), nil
}

// NewEncodingStatisticsEncodingsLiveStatisticsAPIWithClient constructor for EncodingStatisticsEncodingsLiveStatisticsAPI that takes an APIClient as argument
func NewEncodingStatisticsEncodingsLiveStatisticsAPIWithClient(apiClient *apiclient.APIClient) *EncodingStatisticsEncodingsLiveStatisticsAPI {
    a := &EncodingStatisticsEncodingsLiveStatisticsAPI{apiClient: apiClient}
    a.Events = NewEncodingStatisticsEncodingsLiveStatisticsEventsAPIWithClient(apiClient)
    a.Streams = NewEncodingStatisticsEncodingsLiveStatisticsStreamsAPIWithClient(apiClient)
    a.Srt = NewEncodingStatisticsEncodingsLiveStatisticsSrtAPIWithClient(apiClient)

    return a
}

// Get List Live Statistics from an Encoding
func (api *EncodingStatisticsEncodingsLiveStatisticsAPI) Get(encodingId string) (*model.LiveEncodingStats, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel model.LiveEncodingStats
    err := api.apiClient.Get("/encoding/statistics/encodings/{encoding_id}/live-statistics", nil, &responseModel, reqParams)
    return &responseModel, err
}

