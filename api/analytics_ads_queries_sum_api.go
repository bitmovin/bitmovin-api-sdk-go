package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// AnalyticsAdsQueriesSumAPI communicates with '/analytics/ads/queries/sum' endpoints
type AnalyticsAdsQueriesSumAPI struct {
    apiClient *apiclient.APIClient
}

// NewAnalyticsAdsQueriesSumAPI constructor for AnalyticsAdsQueriesSumAPI that takes options as argument
func NewAnalyticsAdsQueriesSumAPI(options ...apiclient.APIClientOption) (*AnalyticsAdsQueriesSumAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewAnalyticsAdsQueriesSumAPIWithClient(apiClient), nil
}

// NewAnalyticsAdsQueriesSumAPIWithClient constructor for AnalyticsAdsQueriesSumAPI that takes an APIClient as argument
func NewAnalyticsAdsQueriesSumAPIWithClient(apiClient *apiclient.APIClient) *AnalyticsAdsQueriesSumAPI {
    a := &AnalyticsAdsQueriesSumAPI{apiClient: apiClient}
    return a
}

// Create Sum
func (api *AnalyticsAdsQueriesSumAPI) Create(adAnalyticsSumQueryRequest model.AdAnalyticsSumQueryRequest) (*model.AnalyticsResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
    }

    var responseModel model.AnalyticsResponse
    err := api.apiClient.Post("/analytics/ads/queries/sum", &adAnalyticsSumQueryRequest, &responseModel, reqParams)
    return &responseModel, err
}

