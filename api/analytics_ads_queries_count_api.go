package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// AnalyticsAdsQueriesCountAPI communicates with '/analytics/ads/queries/count' endpoints
type AnalyticsAdsQueriesCountAPI struct {
	apiClient *apiclient.APIClient
}

// NewAnalyticsAdsQueriesCountAPI constructor for AnalyticsAdsQueriesCountAPI that takes options as argument
func NewAnalyticsAdsQueriesCountAPI(options ...apiclient.APIClientOption) (*AnalyticsAdsQueriesCountAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewAnalyticsAdsQueriesCountAPIWithClient(apiClient), nil
}

// NewAnalyticsAdsQueriesCountAPIWithClient constructor for AnalyticsAdsQueriesCountAPI that takes an APIClient as argument
func NewAnalyticsAdsQueriesCountAPIWithClient(apiClient *apiclient.APIClient) *AnalyticsAdsQueriesCountAPI {
	a := &AnalyticsAdsQueriesCountAPI{apiClient: apiClient}
	return a
}

// Create Count
func (api *AnalyticsAdsQueriesCountAPI) Create(adAnalyticsCountQueryRequest model.AdAnalyticsCountQueryRequest) (*model.AnalyticsResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.AnalyticsResponse
	err := api.apiClient.Post("/analytics/ads/queries/count", &adAnalyticsCountQueryRequest, &responseModel, reqParams)
	return &responseModel, err
}
