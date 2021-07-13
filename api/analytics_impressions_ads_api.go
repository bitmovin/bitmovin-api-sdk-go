package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// AnalyticsImpressionsAdsAPI communicates with '/analytics/impressions/{impression_id}/ads' endpoints
type AnalyticsImpressionsAdsAPI struct {
	apiClient *apiclient.APIClient
}

// NewAnalyticsImpressionsAdsAPI constructor for AnalyticsImpressionsAdsAPI that takes options as argument
func NewAnalyticsImpressionsAdsAPI(options ...apiclient.APIClientOption) (*AnalyticsImpressionsAdsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewAnalyticsImpressionsAdsAPIWithClient(apiClient), nil
}

// NewAnalyticsImpressionsAdsAPIWithClient constructor for AnalyticsImpressionsAdsAPI that takes an APIClient as argument
func NewAnalyticsImpressionsAdsAPIWithClient(apiClient *apiclient.APIClient) *AnalyticsImpressionsAdsAPI {
	a := &AnalyticsImpressionsAdsAPI{apiClient: apiClient}
	return a
}

// Create Ads Impressions
func (api *AnalyticsImpressionsAdsAPI) Create(impressionId string, analyticsLicenseKey model.AnalyticsLicenseKey) (*model.AnalyticsAdsImpressionsResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["impression_id"] = impressionId
	}

	var responseModel model.AnalyticsAdsImpressionsResponse
	err := api.apiClient.Post("/analytics/impressions/{impression_id}/ads", &analyticsLicenseKey, &responseModel, reqParams)
	return &responseModel, err
}
