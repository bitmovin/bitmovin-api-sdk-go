package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingStatisticsEncodingsVodDailyAPI communicates with '/encoding/statistics/encodings/vod/daily/{from}/{to}' endpoints
type EncodingStatisticsEncodingsVodDailyAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingStatisticsEncodingsVodDailyAPI constructor for EncodingStatisticsEncodingsVodDailyAPI that takes options as argument
func NewEncodingStatisticsEncodingsVodDailyAPI(options ...apiclient.APIClientOption) (*EncodingStatisticsEncodingsVodDailyAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingStatisticsEncodingsVodDailyAPIWithClient(apiClient), nil
}

// NewEncodingStatisticsEncodingsVodDailyAPIWithClient constructor for EncodingStatisticsEncodingsVodDailyAPI that takes an APIClient as argument
func NewEncodingStatisticsEncodingsVodDailyAPIWithClient(apiClient *apiclient.APIClient) *EncodingStatisticsEncodingsVodDailyAPI {
	a := &EncodingStatisticsEncodingsVodDailyAPI{apiClient: apiClient}
	return a
}

// ListByDateRange List daily VoD encoding statistics within specific dates
func (api *EncodingStatisticsEncodingsVodDailyAPI) ListByDateRange(from model.Date, to model.Date) (*pagination.EncodingStatisticssListByDateRangePagination, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["from"] = from.String()
		params.PathParams["to"] = to.String()
	}

	var responseModel pagination.EncodingStatisticssListByDateRangePagination
	err := api.apiClient.Get("/encoding/statistics/encodings/vod/daily/{from}/{to}", nil, &responseModel, reqParams)
	return &responseModel, err
}
