package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// AnalyticsQueriesCountAPI communicates with '/analytics/queries/count' endpoints
type AnalyticsQueriesCountAPI struct {
    apiClient *apiclient.APIClient
}

// NewAnalyticsQueriesCountAPI constructor for AnalyticsQueriesCountAPI that takes options as argument
func NewAnalyticsQueriesCountAPI(options ...apiclient.APIClientOption) (*AnalyticsQueriesCountAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewAnalyticsQueriesCountAPIWithClient(apiClient), nil
}

// NewAnalyticsQueriesCountAPIWithClient constructor for AnalyticsQueriesCountAPI that takes an APIClient as argument
func NewAnalyticsQueriesCountAPIWithClient(apiClient *apiclient.APIClient) *AnalyticsQueriesCountAPI {
    a := &AnalyticsQueriesCountAPI{apiClient: apiClient}
    return a
}

// Create Count
func (api *AnalyticsQueriesCountAPI) Create(analyticsCountQueryRequest model.AnalyticsCountQueryRequest) (*model.AnalyticsResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
    }

    var responseModel model.AnalyticsResponse
    err := api.apiClient.Post("/analytics/queries/count", &analyticsCountQueryRequest, &responseModel, reqParams)
    return &responseModel, err
}

