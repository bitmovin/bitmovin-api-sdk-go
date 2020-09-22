package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// AnalyticsQueriesVarianceAPI communicates with '/analytics/queries/variance' endpoints
type AnalyticsQueriesVarianceAPI struct {
	apiClient *apiclient.APIClient
}

// NewAnalyticsQueriesVarianceAPI constructor for AnalyticsQueriesVarianceAPI that takes options as argument
func NewAnalyticsQueriesVarianceAPI(options ...apiclient.APIClientOption) (*AnalyticsQueriesVarianceAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewAnalyticsQueriesVarianceAPIWithClient(apiClient), nil
}

// NewAnalyticsQueriesVarianceAPIWithClient constructor for AnalyticsQueriesVarianceAPI that takes an APIClient as argument
func NewAnalyticsQueriesVarianceAPIWithClient(apiClient *apiclient.APIClient) *AnalyticsQueriesVarianceAPI {
	a := &AnalyticsQueriesVarianceAPI{apiClient: apiClient}
	return a
}

// Create Variance
func (api *AnalyticsQueriesVarianceAPI) Create(analyticsVarianceQueryRequest model.AnalyticsVarianceQueryRequest) (*model.AnalyticsResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.AnalyticsResponse
	err := api.apiClient.Post("/analytics/queries/variance", &analyticsVarianceQueryRequest, &responseModel, reqParams)
	return &responseModel, err
}
