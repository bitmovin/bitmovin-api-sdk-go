package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// AnalyticsImpressionsErrorsAPI communicates with '/analytics/impressions/{impression_id}/errors' endpoints
type AnalyticsImpressionsErrorsAPI struct {
	apiClient *apiclient.APIClient
}

// NewAnalyticsImpressionsErrorsAPI constructor for AnalyticsImpressionsErrorsAPI that takes options as argument
func NewAnalyticsImpressionsErrorsAPI(options ...apiclient.APIClientOption) (*AnalyticsImpressionsErrorsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewAnalyticsImpressionsErrorsAPIWithClient(apiClient), nil
}

// NewAnalyticsImpressionsErrorsAPIWithClient constructor for AnalyticsImpressionsErrorsAPI that takes an APIClient as argument
func NewAnalyticsImpressionsErrorsAPIWithClient(apiClient *apiclient.APIClient) *AnalyticsImpressionsErrorsAPI {
	a := &AnalyticsImpressionsErrorsAPI{apiClient: apiClient}
	return a
}

// Create Impression Error Details
func (api *AnalyticsImpressionsErrorsAPI) Create(impressionId string, analyticsLicenseKey model.AnalyticsLicenseKey) (*model.AnalyticsErrorDetailsResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["impression_id"] = impressionId
	}

	var responseModel model.AnalyticsErrorDetailsResponse
	err := api.apiClient.Post("/analytics/impressions/{impression_id}/errors", &analyticsLicenseKey, &responseModel, reqParams)
	return &responseModel, err
}
