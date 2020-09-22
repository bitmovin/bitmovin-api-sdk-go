package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// AnalyticsMetricsAvgDroppedFramesAPI communicates with '/analytics/metrics/avg-dropped-frames' endpoints
type AnalyticsMetricsAvgDroppedFramesAPI struct {
	apiClient *apiclient.APIClient
}

// NewAnalyticsMetricsAvgDroppedFramesAPI constructor for AnalyticsMetricsAvgDroppedFramesAPI that takes options as argument
func NewAnalyticsMetricsAvgDroppedFramesAPI(options ...apiclient.APIClientOption) (*AnalyticsMetricsAvgDroppedFramesAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewAnalyticsMetricsAvgDroppedFramesAPIWithClient(apiClient), nil
}

// NewAnalyticsMetricsAvgDroppedFramesAPIWithClient constructor for AnalyticsMetricsAvgDroppedFramesAPI that takes an APIClient as argument
func NewAnalyticsMetricsAvgDroppedFramesAPIWithClient(apiClient *apiclient.APIClient) *AnalyticsMetricsAvgDroppedFramesAPI {
	a := &AnalyticsMetricsAvgDroppedFramesAPI{apiClient: apiClient}
	return a
}

// Create Get metrics data
func (api *AnalyticsMetricsAvgDroppedFramesAPI) Create(analyticsMetricsQueryRequest model.AnalyticsMetricsQueryRequest) (*model.AnalyticsAvgDroppedFramesResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.AnalyticsAvgDroppedFramesResponse
	err := api.apiClient.Post("/analytics/metrics/avg-dropped-frames", &analyticsMetricsQueryRequest, &responseModel, reqParams)
	return &responseModel, err
}
