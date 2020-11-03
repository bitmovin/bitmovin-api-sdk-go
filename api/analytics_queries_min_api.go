package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// AnalyticsQueriesMinAPI communicates with '/analytics/queries/min' endpoints
type AnalyticsQueriesMinAPI struct {
    apiClient *apiclient.APIClient
}

// NewAnalyticsQueriesMinAPI constructor for AnalyticsQueriesMinAPI that takes options as argument
func NewAnalyticsQueriesMinAPI(options ...apiclient.APIClientOption) (*AnalyticsQueriesMinAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewAnalyticsQueriesMinAPIWithClient(apiClient), nil
}

// NewAnalyticsQueriesMinAPIWithClient constructor for AnalyticsQueriesMinAPI that takes an APIClient as argument
func NewAnalyticsQueriesMinAPIWithClient(apiClient *apiclient.APIClient) *AnalyticsQueriesMinAPI {
    a := &AnalyticsQueriesMinAPI{apiClient: apiClient}
    return a
}

// Create Min
func (api *AnalyticsQueriesMinAPI) Create(analyticsMinQueryRequest model.AnalyticsMinQueryRequest) (*model.AnalyticsResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
    }

    var responseModel model.AnalyticsResponse
    err := api.apiClient.Post("/analytics/queries/min", &analyticsMinQueryRequest, &responseModel, reqParams)
    return &responseModel, err
}

