package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// AnalyticsAdsQueriesPercentileAPI communicates with '/analytics/ads/queries/percentile' endpoints
type AnalyticsAdsQueriesPercentileAPI struct {
    apiClient *apiclient.APIClient
}

// NewAnalyticsAdsQueriesPercentileAPI constructor for AnalyticsAdsQueriesPercentileAPI that takes options as argument
func NewAnalyticsAdsQueriesPercentileAPI(options ...apiclient.APIClientOption) (*AnalyticsAdsQueriesPercentileAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewAnalyticsAdsQueriesPercentileAPIWithClient(apiClient), nil
}

// NewAnalyticsAdsQueriesPercentileAPIWithClient constructor for AnalyticsAdsQueriesPercentileAPI that takes an APIClient as argument
func NewAnalyticsAdsQueriesPercentileAPIWithClient(apiClient *apiclient.APIClient) *AnalyticsAdsQueriesPercentileAPI {
    a := &AnalyticsAdsQueriesPercentileAPI{apiClient: apiClient}
    return a
}

// Create Percentile
func (api *AnalyticsAdsQueriesPercentileAPI) Create(adAnalyticsPercentileQueryRequest model.AdAnalyticsPercentileQueryRequest) (*model.AnalyticsResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
    }

    var responseModel model.AnalyticsResponse
    err := api.apiClient.Post("/analytics/ads/queries/percentile", &adAnalyticsPercentileQueryRequest, &responseModel, reqParams)
    return &responseModel, err
}

