package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
)

// AnalyticsQueriesAPI intermediary API object with no endpoints
type AnalyticsQueriesAPI struct {
    apiClient *apiclient.APIClient

    // Count communicates with '/analytics/queries/count' endpoints
    Count *AnalyticsQueriesCountAPI
    // Sum communicates with '/analytics/queries/sum' endpoints
    Sum *AnalyticsQueriesSumAPI
    // Avg communicates with '/analytics/queries/avg' endpoints
    Avg *AnalyticsQueriesAvgAPI
    // Min communicates with '/analytics/queries/min' endpoints
    Min *AnalyticsQueriesMinAPI
    // Max communicates with '/analytics/queries/max' endpoints
    Max *AnalyticsQueriesMaxAPI
    // Stddev communicates with '/analytics/queries/stddev' endpoints
    Stddev *AnalyticsQueriesStddevAPI
    // Percentile communicates with '/analytics/queries/percentile' endpoints
    Percentile *AnalyticsQueriesPercentileAPI
    // Variance communicates with '/analytics/queries/variance' endpoints
    Variance *AnalyticsQueriesVarianceAPI
    // Median communicates with '/analytics/queries/median' endpoints
    Median *AnalyticsQueriesMedianAPI

}

// NewAnalyticsQueriesAPI constructor for AnalyticsQueriesAPI that takes options as argument
func NewAnalyticsQueriesAPI(options ...apiclient.APIClientOption) (*AnalyticsQueriesAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewAnalyticsQueriesAPIWithClient(apiClient), nil
}

// NewAnalyticsQueriesAPIWithClient constructor for AnalyticsQueriesAPI that takes an APIClient as argument
func NewAnalyticsQueriesAPIWithClient(apiClient *apiclient.APIClient) *AnalyticsQueriesAPI {
    a := &AnalyticsQueriesAPI{apiClient: apiClient}
    a.Count = NewAnalyticsQueriesCountAPIWithClient(apiClient)
    a.Sum = NewAnalyticsQueriesSumAPIWithClient(apiClient)
    a.Avg = NewAnalyticsQueriesAvgAPIWithClient(apiClient)
    a.Min = NewAnalyticsQueriesMinAPIWithClient(apiClient)
    a.Max = NewAnalyticsQueriesMaxAPIWithClient(apiClient)
    a.Stddev = NewAnalyticsQueriesStddevAPIWithClient(apiClient)
    a.Percentile = NewAnalyticsQueriesPercentileAPIWithClient(apiClient)
    a.Variance = NewAnalyticsQueriesVarianceAPIWithClient(apiClient)
    a.Median = NewAnalyticsQueriesMedianAPIWithClient(apiClient)

    return a
}


