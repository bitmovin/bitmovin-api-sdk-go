package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// AnalyticsAdsQueriesMedianAPI communicates with '/analytics/ads/queries/median' endpoints
type AnalyticsAdsQueriesMedianAPI struct {
	apiClient *apiclient.APIClient
}

// NewAnalyticsAdsQueriesMedianAPI constructor for AnalyticsAdsQueriesMedianAPI that takes options as argument
func NewAnalyticsAdsQueriesMedianAPI(options ...apiclient.APIClientOption) (*AnalyticsAdsQueriesMedianAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewAnalyticsAdsQueriesMedianAPIWithClient(apiClient), nil
}

// NewAnalyticsAdsQueriesMedianAPIWithClient constructor for AnalyticsAdsQueriesMedianAPI that takes an APIClient as argument
func NewAnalyticsAdsQueriesMedianAPIWithClient(apiClient *apiclient.APIClient) *AnalyticsAdsQueriesMedianAPI {
	a := &AnalyticsAdsQueriesMedianAPI{apiClient: apiClient}
	return a
}

// Create Median
func (api *AnalyticsAdsQueriesMedianAPI) Create(adAnalyticsMedianQueryRequest model.AdAnalyticsMedianQueryRequest) (*model.AnalyticsResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.AnalyticsResponse
	err := api.apiClient.Post("/analytics/ads/queries/median", &adAnalyticsMedianQueryRequest, &responseModel, reqParams)
	return &responseModel, err
}
