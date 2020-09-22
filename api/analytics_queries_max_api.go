package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// AnalyticsQueriesMaxAPI communicates with '/analytics/queries/max' endpoints
type AnalyticsQueriesMaxAPI struct {
	apiClient *apiclient.APIClient
}

// NewAnalyticsQueriesMaxAPI constructor for AnalyticsQueriesMaxAPI that takes options as argument
func NewAnalyticsQueriesMaxAPI(options ...apiclient.APIClientOption) (*AnalyticsQueriesMaxAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewAnalyticsQueriesMaxAPIWithClient(apiClient), nil
}

// NewAnalyticsQueriesMaxAPIWithClient constructor for AnalyticsQueriesMaxAPI that takes an APIClient as argument
func NewAnalyticsQueriesMaxAPIWithClient(apiClient *apiclient.APIClient) *AnalyticsQueriesMaxAPI {
	a := &AnalyticsQueriesMaxAPI{apiClient: apiClient}
	return a
}

// Create Max
func (api *AnalyticsQueriesMaxAPI) Create(analyticsMaxQueryRequest model.AnalyticsMaxQueryRequest) (*model.AnalyticsResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.AnalyticsResponse
	err := api.apiClient.Post("/analytics/queries/max", &analyticsMaxQueryRequest, &responseModel, reqParams)
	return &responseModel, err
}
