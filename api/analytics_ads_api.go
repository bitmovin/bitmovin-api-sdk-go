package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
)

// AnalyticsAdsAPI intermediary API object with no endpoints
type AnalyticsAdsAPI struct {
    apiClient *apiclient.APIClient

    // Queries intermediary API object with no endpoints
    Queries *AnalyticsAdsQueriesAPI

}

// NewAnalyticsAdsAPI constructor for AnalyticsAdsAPI that takes options as argument
func NewAnalyticsAdsAPI(options ...apiclient.APIClientOption) (*AnalyticsAdsAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewAnalyticsAdsAPIWithClient(apiClient), nil
}

// NewAnalyticsAdsAPIWithClient constructor for AnalyticsAdsAPI that takes an APIClient as argument
func NewAnalyticsAdsAPIWithClient(apiClient *apiclient.APIClient) *AnalyticsAdsAPI {
    a := &AnalyticsAdsAPI{apiClient: apiClient}
    a.Queries = NewAnalyticsAdsQueriesAPIWithClient(apiClient)

    return a
}


