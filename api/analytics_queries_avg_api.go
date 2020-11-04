package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// AnalyticsQueriesAvgAPI communicates with '/analytics/queries/avg' endpoints
type AnalyticsQueriesAvgAPI struct {
	apiClient *apiclient.APIClient
}

// NewAnalyticsQueriesAvgAPI constructor for AnalyticsQueriesAvgAPI that takes options as argument
func NewAnalyticsQueriesAvgAPI(options ...apiclient.APIClientOption) (*AnalyticsQueriesAvgAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewAnalyticsQueriesAvgAPIWithClient(apiClient), nil
}

// NewAnalyticsQueriesAvgAPIWithClient constructor for AnalyticsQueriesAvgAPI that takes an APIClient as argument
func NewAnalyticsQueriesAvgAPIWithClient(apiClient *apiclient.APIClient) *AnalyticsQueriesAvgAPI {
	a := &AnalyticsQueriesAvgAPI{apiClient: apiClient}
	return a
}

// Create Avg
func (api *AnalyticsQueriesAvgAPI) Create(analyticsAvgQueryRequest model.AnalyticsAvgQueryRequest) (*model.AnalyticsResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.AnalyticsResponse
	err := api.apiClient.Post("/analytics/queries/avg", &analyticsAvgQueryRequest, &responseModel, reqParams)
	return &responseModel, err
}
