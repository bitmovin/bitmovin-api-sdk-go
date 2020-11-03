package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// AnalyticsAdsQueriesAvgAPI communicates with '/analytics/ads/queries/avg' endpoints
type AnalyticsAdsQueriesAvgAPI struct {
    apiClient *apiclient.APIClient
}

// NewAnalyticsAdsQueriesAvgAPI constructor for AnalyticsAdsQueriesAvgAPI that takes options as argument
func NewAnalyticsAdsQueriesAvgAPI(options ...apiclient.APIClientOption) (*AnalyticsAdsQueriesAvgAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewAnalyticsAdsQueriesAvgAPIWithClient(apiClient), nil
}

// NewAnalyticsAdsQueriesAvgAPIWithClient constructor for AnalyticsAdsQueriesAvgAPI that takes an APIClient as argument
func NewAnalyticsAdsQueriesAvgAPIWithClient(apiClient *apiclient.APIClient) *AnalyticsAdsQueriesAvgAPI {
    a := &AnalyticsAdsQueriesAvgAPI{apiClient: apiClient}
    return a
}

// Create Avg
func (api *AnalyticsAdsQueriesAvgAPI) Create(adAnalyticsAvgQueryRequest model.AdAnalyticsAvgQueryRequest) (*model.AnalyticsResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
    }

    var responseModel model.AnalyticsResponse
    err := api.apiClient.Post("/analytics/ads/queries/avg", &adAnalyticsAvgQueryRequest, &responseModel, reqParams)
    return &responseModel, err
}

