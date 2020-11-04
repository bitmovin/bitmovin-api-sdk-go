package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// PlayerLicensesAnalyticsAPI communicates with '/player/licenses/{license_id}/analytics' endpoints
type PlayerLicensesAnalyticsAPI struct {
	apiClient *apiclient.APIClient
}

// NewPlayerLicensesAnalyticsAPI constructor for PlayerLicensesAnalyticsAPI that takes options as argument
func NewPlayerLicensesAnalyticsAPI(options ...apiclient.APIClientOption) (*PlayerLicensesAnalyticsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewPlayerLicensesAnalyticsAPIWithClient(apiClient), nil
}

// NewPlayerLicensesAnalyticsAPIWithClient constructor for PlayerLicensesAnalyticsAPI that takes an APIClient as argument
func NewPlayerLicensesAnalyticsAPIWithClient(apiClient *apiclient.APIClient) *PlayerLicensesAnalyticsAPI {
	a := &PlayerLicensesAnalyticsAPI{apiClient: apiClient}
	return a
}

// Create Activate Analytics
func (api *PlayerLicensesAnalyticsAPI) Create(licenseId string, playerLicenseAnalytics model.PlayerLicenseAnalytics) (*model.PlayerLicenseAnalytics, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["license_id"] = licenseId
	}

	var responseModel model.PlayerLicenseAnalytics
	err := api.apiClient.Post("/player/licenses/{license_id}/analytics", &playerLicenseAnalytics, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Deactivate Analytics
func (api *PlayerLicensesAnalyticsAPI) Delete(licenseId string) (*model.PlayerLicenseAnalytics, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["license_id"] = licenseId
	}

	var responseModel model.PlayerLicenseAnalytics
	err := api.apiClient.Delete("/player/licenses/{license_id}/analytics", nil, &responseModel, reqParams)
	return &responseModel, err
}
