package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// AnalyticsAdsQueriesMinAPI communicates with '/analytics/ads/queries/min' endpoints
type AnalyticsAdsQueriesMinAPI struct {
	apiClient *apiclient.APIClient
}

// NewAnalyticsAdsQueriesMinAPI constructor for AnalyticsAdsQueriesMinAPI that takes options as argument
func NewAnalyticsAdsQueriesMinAPI(options ...apiclient.APIClientOption) (*AnalyticsAdsQueriesMinAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewAnalyticsAdsQueriesMinAPIWithClient(apiClient), nil
}

// NewAnalyticsAdsQueriesMinAPIWithClient constructor for AnalyticsAdsQueriesMinAPI that takes an APIClient as argument
func NewAnalyticsAdsQueriesMinAPIWithClient(apiClient *apiclient.APIClient) *AnalyticsAdsQueriesMinAPI {
	a := &AnalyticsAdsQueriesMinAPI{apiClient: apiClient}
	return a
}

// Create Min
func (api *AnalyticsAdsQueriesMinAPI) Create(adAnalyticsMinQueryRequest model.AdAnalyticsMinQueryRequest) (*model.AnalyticsResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.AnalyticsResponse
	err := api.apiClient.Post("/analytics/ads/queries/min", &adAnalyticsMinQueryRequest, &responseModel, reqParams)
	return &responseModel, err
}
