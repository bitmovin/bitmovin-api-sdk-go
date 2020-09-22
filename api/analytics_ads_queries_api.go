package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
)

// AnalyticsAdsQueriesAPI intermediary API object with no endpoints
type AnalyticsAdsQueriesAPI struct {
	apiClient *apiclient.APIClient

	// Count communicates with '/analytics/ads/queries/count' endpoints
	Count *AnalyticsAdsQueriesCountAPI
	// Sum communicates with '/analytics/ads/queries/sum' endpoints
	Sum *AnalyticsAdsQueriesSumAPI
	// Avg communicates with '/analytics/ads/queries/avg' endpoints
	Avg *AnalyticsAdsQueriesAvgAPI
	// Min communicates with '/analytics/ads/queries/min' endpoints
	Min *AnalyticsAdsQueriesMinAPI
	// Max communicates with '/analytics/ads/queries/max' endpoints
	Max *AnalyticsAdsQueriesMaxAPI
	// Stddev communicates with '/analytics/ads/queries/stddev' endpoints
	Stddev *AnalyticsAdsQueriesStddevAPI
	// Percentile communicates with '/analytics/ads/queries/percentile' endpoints
	Percentile *AnalyticsAdsQueriesPercentileAPI
	// Variance communicates with '/analytics/ads/queries/variance' endpoints
	Variance *AnalyticsAdsQueriesVarianceAPI
	// Median communicates with '/analytics/ads/queries/median' endpoints
	Median *AnalyticsAdsQueriesMedianAPI
}

// NewAnalyticsAdsQueriesAPI constructor for AnalyticsAdsQueriesAPI that takes options as argument
func NewAnalyticsAdsQueriesAPI(options ...apiclient.APIClientOption) (*AnalyticsAdsQueriesAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewAnalyticsAdsQueriesAPIWithClient(apiClient), nil
}

// NewAnalyticsAdsQueriesAPIWithClient constructor for AnalyticsAdsQueriesAPI that takes an APIClient as argument
func NewAnalyticsAdsQueriesAPIWithClient(apiClient *apiclient.APIClient) *AnalyticsAdsQueriesAPI {
	a := &AnalyticsAdsQueriesAPI{apiClient: apiClient}
	a.Count = NewAnalyticsAdsQueriesCountAPIWithClient(apiClient)
	a.Sum = NewAnalyticsAdsQueriesSumAPIWithClient(apiClient)
	a.Avg = NewAnalyticsAdsQueriesAvgAPIWithClient(apiClient)
	a.Min = NewAnalyticsAdsQueriesMinAPIWithClient(apiClient)
	a.Max = NewAnalyticsAdsQueriesMaxAPIWithClient(apiClient)
	a.Stddev = NewAnalyticsAdsQueriesStddevAPIWithClient(apiClient)
	a.Percentile = NewAnalyticsAdsQueriesPercentileAPIWithClient(apiClient)
	a.Variance = NewAnalyticsAdsQueriesVarianceAPIWithClient(apiClient)
	a.Median = NewAnalyticsAdsQueriesMedianAPIWithClient(apiClient)

	return a
}
