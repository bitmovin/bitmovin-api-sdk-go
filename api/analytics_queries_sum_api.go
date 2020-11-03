package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// AnalyticsQueriesSumAPI communicates with '/analytics/queries/sum' endpoints
type AnalyticsQueriesSumAPI struct {
    apiClient *apiclient.APIClient
}

// NewAnalyticsQueriesSumAPI constructor for AnalyticsQueriesSumAPI that takes options as argument
func NewAnalyticsQueriesSumAPI(options ...apiclient.APIClientOption) (*AnalyticsQueriesSumAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewAnalyticsQueriesSumAPIWithClient(apiClient), nil
}

// NewAnalyticsQueriesSumAPIWithClient constructor for AnalyticsQueriesSumAPI that takes an APIClient as argument
func NewAnalyticsQueriesSumAPIWithClient(apiClient *apiclient.APIClient) *AnalyticsQueriesSumAPI {
    a := &AnalyticsQueriesSumAPI{apiClient: apiClient}
    return a
}

// Create Sum
func (api *AnalyticsQueriesSumAPI) Create(analyticsSumQueryRequest model.AnalyticsSumQueryRequest) (*model.AnalyticsResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
    }

    var responseModel model.AnalyticsResponse
    err := api.apiClient.Post("/analytics/queries/sum", &analyticsSumQueryRequest, &responseModel, reqParams)
    return &responseModel, err
}

