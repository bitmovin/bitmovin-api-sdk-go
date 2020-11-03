package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
)

// AnalyticsAPI intermediary API object with no endpoints
type AnalyticsAPI struct {
    apiClient *apiclient.APIClient

    // Exports communicates with '/analytics/exports' endpoints
    Exports *AnalyticsExportsAPI
    // Impressions communicates with '/analytics/impressions' endpoints
    Impressions *AnalyticsImpressionsAPI
    // Insights intermediary API object with no endpoints
    Insights *AnalyticsInsightsAPI
    // Metrics intermediary API object with no endpoints
    Metrics *AnalyticsMetricsAPI
    // Ads intermediary API object with no endpoints
    Ads *AnalyticsAdsAPI
    // Queries intermediary API object with no endpoints
    Queries *AnalyticsQueriesAPI
    // Licenses communicates with '/analytics/licenses' endpoints
    Licenses *AnalyticsLicensesAPI
    // Outputs intermediary API object with no endpoints
    Outputs *AnalyticsOutputsAPI
    // Alerting intermediary API object with no endpoints
    Alerting *AnalyticsAlertingAPI

}

// NewAnalyticsAPI constructor for AnalyticsAPI that takes options as argument
func NewAnalyticsAPI(options ...apiclient.APIClientOption) (*AnalyticsAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewAnalyticsAPIWithClient(apiClient), nil
}

// NewAnalyticsAPIWithClient constructor for AnalyticsAPI that takes an APIClient as argument
func NewAnalyticsAPIWithClient(apiClient *apiclient.APIClient) *AnalyticsAPI {
    a := &AnalyticsAPI{apiClient: apiClient}
    a.Exports = NewAnalyticsExportsAPIWithClient(apiClient)
    a.Impressions = NewAnalyticsImpressionsAPIWithClient(apiClient)
    a.Insights = NewAnalyticsInsightsAPIWithClient(apiClient)
    a.Metrics = NewAnalyticsMetricsAPIWithClient(apiClient)
    a.Ads = NewAnalyticsAdsAPIWithClient(apiClient)
    a.Queries = NewAnalyticsQueriesAPIWithClient(apiClient)
    a.Licenses = NewAnalyticsLicensesAPIWithClient(apiClient)
    a.Outputs = NewAnalyticsOutputsAPIWithClient(apiClient)
    a.Alerting = NewAnalyticsAlertingAPIWithClient(apiClient)

    return a
}


