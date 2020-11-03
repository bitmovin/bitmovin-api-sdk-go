package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
)

// AnalyticsInsightsAPI intermediary API object with no endpoints
type AnalyticsInsightsAPI struct {
    apiClient *apiclient.APIClient

    // Organizations intermediary API object with no endpoints
    Organizations *AnalyticsInsightsOrganizationsAPI

}

// NewAnalyticsInsightsAPI constructor for AnalyticsInsightsAPI that takes options as argument
func NewAnalyticsInsightsAPI(options ...apiclient.APIClientOption) (*AnalyticsInsightsAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewAnalyticsInsightsAPIWithClient(apiClient), nil
}

// NewAnalyticsInsightsAPIWithClient constructor for AnalyticsInsightsAPI that takes an APIClient as argument
func NewAnalyticsInsightsAPIWithClient(apiClient *apiclient.APIClient) *AnalyticsInsightsAPI {
    a := &AnalyticsInsightsAPI{apiClient: apiClient}
    a.Organizations = NewAnalyticsInsightsOrganizationsAPIWithClient(apiClient)

    return a
}


