package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// AnalyticsImpressionsAPI communicates with '/analytics/impressions' endpoints
type AnalyticsImpressionsAPI struct {
	apiClient *apiclient.APIClient

	// Errors communicates with '/analytics/impressions/{impression_id}/errors' endpoints
	Errors *AnalyticsImpressionsErrorsAPI
}

// NewAnalyticsImpressionsAPI constructor for AnalyticsImpressionsAPI that takes options as argument
func NewAnalyticsImpressionsAPI(options ...apiclient.APIClientOption) (*AnalyticsImpressionsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewAnalyticsImpressionsAPIWithClient(apiClient), nil
}

// NewAnalyticsImpressionsAPIWithClient constructor for AnalyticsImpressionsAPI that takes an APIClient as argument
func NewAnalyticsImpressionsAPIWithClient(apiClient *apiclient.APIClient) *AnalyticsImpressionsAPI {
	a := &AnalyticsImpressionsAPI{apiClient: apiClient}
	a.Errors = NewAnalyticsImpressionsErrorsAPIWithClient(apiClient)

	return a
}

// GetImpressions List impressions
func (api *AnalyticsImpressionsAPI) GetImpressions(analyticsImpressionsQuery model.AnalyticsImpressionsQuery) (*model.AnalyticsImpressionsResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.AnalyticsImpressionsResponse
	err := api.apiClient.Post("/analytics/impressions", &analyticsImpressionsQuery, &responseModel, reqParams)
	return &responseModel, err
}
