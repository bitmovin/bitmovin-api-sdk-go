package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
)

// AnalyticsInsightsOrganizationsAPI intermediary API object with no endpoints
type AnalyticsInsightsOrganizationsAPI struct {
    apiClient *apiclient.APIClient

    // Settings communicates with '/analytics/insights/organizations/{organization_id}/settings' endpoints
    Settings *AnalyticsInsightsOrganizationsSettingsAPI

}

// NewAnalyticsInsightsOrganizationsAPI constructor for AnalyticsInsightsOrganizationsAPI that takes options as argument
func NewAnalyticsInsightsOrganizationsAPI(options ...apiclient.APIClientOption) (*AnalyticsInsightsOrganizationsAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewAnalyticsInsightsOrganizationsAPIWithClient(apiClient), nil
}

// NewAnalyticsInsightsOrganizationsAPIWithClient constructor for AnalyticsInsightsOrganizationsAPI that takes an APIClient as argument
func NewAnalyticsInsightsOrganizationsAPIWithClient(apiClient *apiclient.APIClient) *AnalyticsInsightsOrganizationsAPI {
    a := &AnalyticsInsightsOrganizationsAPI{apiClient: apiClient}
    a.Settings = NewAnalyticsInsightsOrganizationsSettingsAPIWithClient(apiClient)

    return a
}


