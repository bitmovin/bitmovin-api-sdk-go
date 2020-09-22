package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// AnalyticsMetricsMaxConcurrentviewersAPI communicates with '/analytics/metrics/max-concurrentviewers' endpoints
type AnalyticsMetricsMaxConcurrentviewersAPI struct {
	apiClient *apiclient.APIClient
}

// NewAnalyticsMetricsMaxConcurrentviewersAPI constructor for AnalyticsMetricsMaxConcurrentviewersAPI that takes options as argument
func NewAnalyticsMetricsMaxConcurrentviewersAPI(options ...apiclient.APIClientOption) (*AnalyticsMetricsMaxConcurrentviewersAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewAnalyticsMetricsMaxConcurrentviewersAPIWithClient(apiClient), nil
}

// NewAnalyticsMetricsMaxConcurrentviewersAPIWithClient constructor for AnalyticsMetricsMaxConcurrentviewersAPI that takes an APIClient as argument
func NewAnalyticsMetricsMaxConcurrentviewersAPIWithClient(apiClient *apiclient.APIClient) *AnalyticsMetricsMaxConcurrentviewersAPI {
	a := &AnalyticsMetricsMaxConcurrentviewersAPI{apiClient: apiClient}
	return a
}

// Create Get metrics data
func (api *AnalyticsMetricsMaxConcurrentviewersAPI) Create(analyticsMetricsQueryRequest model.AnalyticsMetricsQueryRequest) (*model.AnalyticsMaxConcurrentViewersResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.AnalyticsMaxConcurrentViewersResponse
	err := api.apiClient.Post("/analytics/metrics/max-concurrentviewers", &analyticsMetricsQueryRequest, &responseModel, reqParams)
	return &responseModel, err
}
