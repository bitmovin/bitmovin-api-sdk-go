package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// AnalyticsAdsQueriesVarianceAPI communicates with '/analytics/ads/queries/variance' endpoints
type AnalyticsAdsQueriesVarianceAPI struct {
	apiClient *apiclient.APIClient
}

// NewAnalyticsAdsQueriesVarianceAPI constructor for AnalyticsAdsQueriesVarianceAPI that takes options as argument
func NewAnalyticsAdsQueriesVarianceAPI(options ...apiclient.APIClientOption) (*AnalyticsAdsQueriesVarianceAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewAnalyticsAdsQueriesVarianceAPIWithClient(apiClient), nil
}

// NewAnalyticsAdsQueriesVarianceAPIWithClient constructor for AnalyticsAdsQueriesVarianceAPI that takes an APIClient as argument
func NewAnalyticsAdsQueriesVarianceAPIWithClient(apiClient *apiclient.APIClient) *AnalyticsAdsQueriesVarianceAPI {
	a := &AnalyticsAdsQueriesVarianceAPI{apiClient: apiClient}
	return a
}

// Create Variance
func (api *AnalyticsAdsQueriesVarianceAPI) Create(adAnalyticsVarianceQueryRequest model.AdAnalyticsVarianceQueryRequest) (*model.AnalyticsResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.AnalyticsResponse
	err := api.apiClient.Post("/analytics/ads/queries/variance", &adAnalyticsVarianceQueryRequest, &responseModel, reqParams)
	return &responseModel, err
}
