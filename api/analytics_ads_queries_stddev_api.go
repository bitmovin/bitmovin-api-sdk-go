package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// AnalyticsAdsQueriesStddevAPI communicates with '/analytics/ads/queries/stddev' endpoints
type AnalyticsAdsQueriesStddevAPI struct {
	apiClient *apiclient.APIClient
}

// NewAnalyticsAdsQueriesStddevAPI constructor for AnalyticsAdsQueriesStddevAPI that takes options as argument
func NewAnalyticsAdsQueriesStddevAPI(options ...apiclient.APIClientOption) (*AnalyticsAdsQueriesStddevAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewAnalyticsAdsQueriesStddevAPIWithClient(apiClient), nil
}

// NewAnalyticsAdsQueriesStddevAPIWithClient constructor for AnalyticsAdsQueriesStddevAPI that takes an APIClient as argument
func NewAnalyticsAdsQueriesStddevAPIWithClient(apiClient *apiclient.APIClient) *AnalyticsAdsQueriesStddevAPI {
	a := &AnalyticsAdsQueriesStddevAPI{apiClient: apiClient}
	return a
}

// Create Stddev
func (api *AnalyticsAdsQueriesStddevAPI) Create(adAnalyticsStddevQueryRequest model.AdAnalyticsStddevQueryRequest) (*model.AnalyticsResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.AnalyticsResponse
	err := api.apiClient.Post("/analytics/ads/queries/stddev", &adAnalyticsStddevQueryRequest, &responseModel, reqParams)
	return &responseModel, err
}
