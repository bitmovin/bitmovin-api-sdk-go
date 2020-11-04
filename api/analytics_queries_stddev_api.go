package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// AnalyticsQueriesStddevAPI communicates with '/analytics/queries/stddev' endpoints
type AnalyticsQueriesStddevAPI struct {
	apiClient *apiclient.APIClient
}

// NewAnalyticsQueriesStddevAPI constructor for AnalyticsQueriesStddevAPI that takes options as argument
func NewAnalyticsQueriesStddevAPI(options ...apiclient.APIClientOption) (*AnalyticsQueriesStddevAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewAnalyticsQueriesStddevAPIWithClient(apiClient), nil
}

// NewAnalyticsQueriesStddevAPIWithClient constructor for AnalyticsQueriesStddevAPI that takes an APIClient as argument
func NewAnalyticsQueriesStddevAPIWithClient(apiClient *apiclient.APIClient) *AnalyticsQueriesStddevAPI {
	a := &AnalyticsQueriesStddevAPI{apiClient: apiClient}
	return a
}

// Create Stddev
func (api *AnalyticsQueriesStddevAPI) Create(analyticsStddevQueryRequest model.AnalyticsStddevQueryRequest) (*model.AnalyticsResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.AnalyticsResponse
	err := api.apiClient.Post("/analytics/queries/stddev", &analyticsStddevQueryRequest, &responseModel, reqParams)
	return &responseModel, err
}
