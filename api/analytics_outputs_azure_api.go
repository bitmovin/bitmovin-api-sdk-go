package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// AnalyticsOutputsAzureAPI communicates with '/analytics/outputs/azure' endpoints
type AnalyticsOutputsAzureAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/analytics/outputs/azure/{output_id}/customData' endpoints
	Customdata *AnalyticsOutputsAzureCustomdataAPI
}

// NewAnalyticsOutputsAzureAPI constructor for AnalyticsOutputsAzureAPI that takes options as argument
func NewAnalyticsOutputsAzureAPI(options ...apiclient.APIClientOption) (*AnalyticsOutputsAzureAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewAnalyticsOutputsAzureAPIWithClient(apiClient), nil
}

// NewAnalyticsOutputsAzureAPIWithClient constructor for AnalyticsOutputsAzureAPI that takes an APIClient as argument
func NewAnalyticsOutputsAzureAPIWithClient(apiClient *apiclient.APIClient) *AnalyticsOutputsAzureAPI {
	a := &AnalyticsOutputsAzureAPI{apiClient: apiClient}
	a.Customdata = NewAnalyticsOutputsAzureCustomdataAPIWithClient(apiClient)

	return a
}

// Create Microsoft Azure Output
func (api *AnalyticsOutputsAzureAPI) Create(analyticsAzureOutput model.AnalyticsAzureOutput) (*model.AnalyticsAzureOutput, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.AnalyticsAzureOutput
	err := api.apiClient.Post("/analytics/outputs/azure", &analyticsAzureOutput, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Microsoft Azure Output
func (api *AnalyticsOutputsAzureAPI) Delete(outputId string) (*model.AnalyticsAzureOutput, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["output_id"] = outputId
	}

	var responseModel model.AnalyticsAzureOutput
	err := api.apiClient.Delete("/analytics/outputs/azure/{output_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Microsoft Azure Output Details
func (api *AnalyticsOutputsAzureAPI) Get(outputId string) (*model.AnalyticsAzureOutput, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["output_id"] = outputId
	}

	var responseModel model.AnalyticsAzureOutput
	err := api.apiClient.Get("/analytics/outputs/azure/{output_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List Microsoft Azure Outputs
func (api *AnalyticsOutputsAzureAPI) List(queryParams ...func(*AnalyticsOutputsAzureAPIListQueryParams)) (*pagination.AnalyticsAzureOutputsListPagination, error) {
	queryParameters := &AnalyticsOutputsAzureAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.AnalyticsAzureOutputsListPagination
	err := api.apiClient.Get("/analytics/outputs/azure", nil, &responseModel, reqParams)
	return &responseModel, err
}

// AnalyticsOutputsAzureAPIListQueryParams contains all query parameters for the List endpoint
type AnalyticsOutputsAzureAPIListQueryParams struct {
	Offset int32  `query:"offset"`
	Limit  int32  `query:"limit"`
	Name   string `query:"name"`
}

// Params will return a map of query parameters
func (q *AnalyticsOutputsAzureAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
