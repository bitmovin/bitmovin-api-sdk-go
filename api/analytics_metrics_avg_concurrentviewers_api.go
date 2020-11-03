package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// AnalyticsMetricsAvgConcurrentviewersAPI communicates with '/analytics/metrics/avg-concurrentviewers' endpoints
type AnalyticsMetricsAvgConcurrentviewersAPI struct {
    apiClient *apiclient.APIClient
}

// NewAnalyticsMetricsAvgConcurrentviewersAPI constructor for AnalyticsMetricsAvgConcurrentviewersAPI that takes options as argument
func NewAnalyticsMetricsAvgConcurrentviewersAPI(options ...apiclient.APIClientOption) (*AnalyticsMetricsAvgConcurrentviewersAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewAnalyticsMetricsAvgConcurrentviewersAPIWithClient(apiClient), nil
}

// NewAnalyticsMetricsAvgConcurrentviewersAPIWithClient constructor for AnalyticsMetricsAvgConcurrentviewersAPI that takes an APIClient as argument
func NewAnalyticsMetricsAvgConcurrentviewersAPIWithClient(apiClient *apiclient.APIClient) *AnalyticsMetricsAvgConcurrentviewersAPI {
    a := &AnalyticsMetricsAvgConcurrentviewersAPI{apiClient: apiClient}
    return a
}

// Create Get metrics data
func (api *AnalyticsMetricsAvgConcurrentviewersAPI) Create(analyticsMetricsQueryRequest model.AnalyticsMetricsQueryRequest) (*model.AnalyticsAvgConcurrentViewersResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
    }

    var responseModel model.AnalyticsAvgConcurrentViewersResponse
    err := api.apiClient.Post("/analytics/metrics/avg-concurrentviewers", &analyticsMetricsQueryRequest, &responseModel, reqParams)
    return &responseModel, err
}

