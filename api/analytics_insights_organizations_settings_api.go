package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// AnalyticsInsightsOrganizationsSettingsAPI communicates with '/analytics/insights/organizations/{organization_id}/settings' endpoints
type AnalyticsInsightsOrganizationsSettingsAPI struct {
	apiClient *apiclient.APIClient
}

// NewAnalyticsInsightsOrganizationsSettingsAPI constructor for AnalyticsInsightsOrganizationsSettingsAPI that takes options as argument
func NewAnalyticsInsightsOrganizationsSettingsAPI(options ...apiclient.APIClientOption) (*AnalyticsInsightsOrganizationsSettingsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewAnalyticsInsightsOrganizationsSettingsAPIWithClient(apiClient), nil
}

// NewAnalyticsInsightsOrganizationsSettingsAPIWithClient constructor for AnalyticsInsightsOrganizationsSettingsAPI that takes an APIClient as argument
func NewAnalyticsInsightsOrganizationsSettingsAPIWithClient(apiClient *apiclient.APIClient) *AnalyticsInsightsOrganizationsSettingsAPI {
	a := &AnalyticsInsightsOrganizationsSettingsAPI{apiClient: apiClient}
	return a
}

// Get the current organization settings for industry insights
func (api *AnalyticsInsightsOrganizationsSettingsAPI) Get(organizationId string) (*model.AnalyticsInsightsOrganizationSettings, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["organization_id"] = organizationId
	}

	var responseModel model.AnalyticsInsightsOrganizationSettings
	err := api.apiClient.Get("/analytics/insights/organizations/{organization_id}/settings", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Update the organization settings for industry insights
func (api *AnalyticsInsightsOrganizationsSettingsAPI) Update(organizationId string, analyticsInsightsOrganizationSettingsRequest model.AnalyticsInsightsOrganizationSettingsRequest) (*model.AnalyticsInsightsOrganizationSettings, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["organization_id"] = organizationId
	}

	var responseModel model.AnalyticsInsightsOrganizationSettings
	err := api.apiClient.Put("/analytics/insights/organizations/{organization_id}/settings", &analyticsInsightsOrganizationSettingsRequest, &responseModel, reqParams)
	return &responseModel, err
}
