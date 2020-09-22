package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// AnalyticsQueriesPercentileAPI communicates with '/analytics/queries/percentile' endpoints
type AnalyticsQueriesPercentileAPI struct {
	apiClient *apiclient.APIClient
}

// NewAnalyticsQueriesPercentileAPI constructor for AnalyticsQueriesPercentileAPI that takes options as argument
func NewAnalyticsQueriesPercentileAPI(options ...apiclient.APIClientOption) (*AnalyticsQueriesPercentileAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewAnalyticsQueriesPercentileAPIWithClient(apiClient), nil
}

// NewAnalyticsQueriesPercentileAPIWithClient constructor for AnalyticsQueriesPercentileAPI that takes an APIClient as argument
func NewAnalyticsQueriesPercentileAPIWithClient(apiClient *apiclient.APIClient) *AnalyticsQueriesPercentileAPI {
	a := &AnalyticsQueriesPercentileAPI{apiClient: apiClient}
	return a
}

// Create Percentile
func (api *AnalyticsQueriesPercentileAPI) Create(analyticsPercentileQueryRequest model.AnalyticsPercentileQueryRequest) (*model.AnalyticsResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.AnalyticsResponse
	err := api.apiClient.Post("/analytics/queries/percentile", &analyticsPercentileQueryRequest, &responseModel, reqParams)
	return &responseModel, err
}
