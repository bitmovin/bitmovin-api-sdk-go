package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// AnalyticsQueriesMedianAPI communicates with '/analytics/queries/median' endpoints
type AnalyticsQueriesMedianAPI struct {
	apiClient *apiclient.APIClient
}

// NewAnalyticsQueriesMedianAPI constructor for AnalyticsQueriesMedianAPI that takes options as argument
func NewAnalyticsQueriesMedianAPI(options ...apiclient.APIClientOption) (*AnalyticsQueriesMedianAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewAnalyticsQueriesMedianAPIWithClient(apiClient), nil
}

// NewAnalyticsQueriesMedianAPIWithClient constructor for AnalyticsQueriesMedianAPI that takes an APIClient as argument
func NewAnalyticsQueriesMedianAPIWithClient(apiClient *apiclient.APIClient) *AnalyticsQueriesMedianAPI {
	a := &AnalyticsQueriesMedianAPI{apiClient: apiClient}
	return a
}

// Create Median
func (api *AnalyticsQueriesMedianAPI) Create(analyticsMedianQueryRequest model.AnalyticsMedianQueryRequest) (*model.AnalyticsResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.AnalyticsResponse
	err := api.apiClient.Post("/analytics/queries/median", &analyticsMedianQueryRequest, &responseModel, reqParams)
	return &responseModel, err
}
