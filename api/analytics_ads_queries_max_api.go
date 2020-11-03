package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// AnalyticsAdsQueriesMaxAPI communicates with '/analytics/ads/queries/max' endpoints
type AnalyticsAdsQueriesMaxAPI struct {
    apiClient *apiclient.APIClient
}

// NewAnalyticsAdsQueriesMaxAPI constructor for AnalyticsAdsQueriesMaxAPI that takes options as argument
func NewAnalyticsAdsQueriesMaxAPI(options ...apiclient.APIClientOption) (*AnalyticsAdsQueriesMaxAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewAnalyticsAdsQueriesMaxAPIWithClient(apiClient), nil
}

// NewAnalyticsAdsQueriesMaxAPIWithClient constructor for AnalyticsAdsQueriesMaxAPI that takes an APIClient as argument
func NewAnalyticsAdsQueriesMaxAPIWithClient(apiClient *apiclient.APIClient) *AnalyticsAdsQueriesMaxAPI {
    a := &AnalyticsAdsQueriesMaxAPI{apiClient: apiClient}
    return a
}

// Create Max
func (api *AnalyticsAdsQueriesMaxAPI) Create(adAnalyticsMaxQueryRequest model.AdAnalyticsMaxQueryRequest) (*model.AnalyticsResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
    }

    var responseModel model.AnalyticsResponse
    err := api.apiClient.Post("/analytics/ads/queries/max", &adAnalyticsMaxQueryRequest, &responseModel, reqParams)
    return &responseModel, err
}

