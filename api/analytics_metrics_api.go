package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
)

// AnalyticsMetricsAPI intermediary API object with no endpoints
type AnalyticsMetricsAPI struct {
	apiClient *apiclient.APIClient

	// MaxConcurrentviewers communicates with '/analytics/metrics/max-concurrentviewers' endpoints
	MaxConcurrentviewers *AnalyticsMetricsMaxConcurrentviewersAPI
	// AvgConcurrentviewers communicates with '/analytics/metrics/avg-concurrentviewers' endpoints
	AvgConcurrentviewers *AnalyticsMetricsAvgConcurrentviewersAPI
	// AvgDroppedFrames communicates with '/analytics/metrics/avg-dropped-frames' endpoints
	AvgDroppedFrames *AnalyticsMetricsAvgDroppedFramesAPI
}

// NewAnalyticsMetricsAPI constructor for AnalyticsMetricsAPI that takes options as argument
func NewAnalyticsMetricsAPI(options ...apiclient.APIClientOption) (*AnalyticsMetricsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewAnalyticsMetricsAPIWithClient(apiClient), nil
}

// NewAnalyticsMetricsAPIWithClient constructor for AnalyticsMetricsAPI that takes an APIClient as argument
func NewAnalyticsMetricsAPIWithClient(apiClient *apiclient.APIClient) *AnalyticsMetricsAPI {
	a := &AnalyticsMetricsAPI{apiClient: apiClient}
	a.MaxConcurrentviewers = NewAnalyticsMetricsMaxConcurrentviewersAPIWithClient(apiClient)
	a.AvgConcurrentviewers = NewAnalyticsMetricsAvgConcurrentviewersAPIWithClient(apiClient)
	a.AvgDroppedFrames = NewAnalyticsMetricsAvgDroppedFramesAPIWithClient(apiClient)

	return a
}
