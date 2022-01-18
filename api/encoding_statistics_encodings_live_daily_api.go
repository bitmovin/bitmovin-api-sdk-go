package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingStatisticsEncodingsLiveDailyAPI communicates with '/encoding/statistics/encodings/live/daily/{from}/{to}' endpoints
type EncodingStatisticsEncodingsLiveDailyAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingStatisticsEncodingsLiveDailyAPI constructor for EncodingStatisticsEncodingsLiveDailyAPI that takes options as argument
func NewEncodingStatisticsEncodingsLiveDailyAPI(options ...apiclient.APIClientOption) (*EncodingStatisticsEncodingsLiveDailyAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingStatisticsEncodingsLiveDailyAPIWithClient(apiClient), nil
}

// NewEncodingStatisticsEncodingsLiveDailyAPIWithClient constructor for EncodingStatisticsEncodingsLiveDailyAPI that takes an APIClient as argument
func NewEncodingStatisticsEncodingsLiveDailyAPIWithClient(apiClient *apiclient.APIClient) *EncodingStatisticsEncodingsLiveDailyAPI {
	a := &EncodingStatisticsEncodingsLiveDailyAPI{apiClient: apiClient}
	return a
}

// ListByDateRange List daily live encoding statistics within specific dates
func (api *EncodingStatisticsEncodingsLiveDailyAPI) ListByDateRange(from model.Date, to model.Date) (*pagination.EncodingStatisticsLivesListByDateRangePagination, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["from"] = from.String()
		params.PathParams["to"] = to.String()
	}

	var responseModel pagination.EncodingStatisticsLivesListByDateRangePagination
	err := api.apiClient.Get("/encoding/statistics/encodings/live/daily/{from}/{to}", nil, &responseModel, reqParams)
	return &responseModel, err
}
