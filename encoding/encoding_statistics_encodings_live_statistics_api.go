package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type EncodingStatisticsEncodingsLiveStatisticsApi struct {
    apiClient *common.ApiClient
    Events *EncodingStatisticsEncodingsLiveStatisticsEventsApi
    Streams *EncodingStatisticsEncodingsLiveStatisticsStreamsApi
}

func NewEncodingStatisticsEncodingsLiveStatisticsApi(configs ...func(*common.ApiClient)) (*EncodingStatisticsEncodingsLiveStatisticsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingStatisticsEncodingsLiveStatisticsApi{apiClient: apiClient}

    eventsApi, err := NewEncodingStatisticsEncodingsLiveStatisticsEventsApi(configs...)
    api.Events = eventsApi
    streamsApi, err := NewEncodingStatisticsEncodingsLiveStatisticsStreamsApi(configs...)
    api.Streams = streamsApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingStatisticsEncodingsLiveStatisticsApi) Get(encodingId string) (*model.LiveEncodingStats, error) {
    var resp *model.LiveEncodingStats
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
	}
    err := api.apiClient.Get("/encoding/statistics/encodings/{encoding_id}/live-statistics", &resp, reqParams)
    return resp, err
}
