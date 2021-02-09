package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// AnalyticsOutputsS3RoleBasedAPI communicates with '/analytics/outputs/s3-role-based' endpoints
type AnalyticsOutputsS3RoleBasedAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/analytics/outputs/s3-role-based/{output_id}/customData' endpoints
	Customdata *AnalyticsOutputsS3RoleBasedCustomdataAPI
}

// NewAnalyticsOutputsS3RoleBasedAPI constructor for AnalyticsOutputsS3RoleBasedAPI that takes options as argument
func NewAnalyticsOutputsS3RoleBasedAPI(options ...apiclient.APIClientOption) (*AnalyticsOutputsS3RoleBasedAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewAnalyticsOutputsS3RoleBasedAPIWithClient(apiClient), nil
}

// NewAnalyticsOutputsS3RoleBasedAPIWithClient constructor for AnalyticsOutputsS3RoleBasedAPI that takes an APIClient as argument
func NewAnalyticsOutputsS3RoleBasedAPIWithClient(apiClient *apiclient.APIClient) *AnalyticsOutputsS3RoleBasedAPI {
	a := &AnalyticsOutputsS3RoleBasedAPI{apiClient: apiClient}
	a.Customdata = NewAnalyticsOutputsS3RoleBasedCustomdataAPIWithClient(apiClient)

	return a
}

// Create S3 Role-based Output
func (api *AnalyticsOutputsS3RoleBasedAPI) Create(analyticsS3RoleBasedOutput model.AnalyticsS3RoleBasedOutput) (*model.AnalyticsS3RoleBasedOutput, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.AnalyticsS3RoleBasedOutput
	err := api.apiClient.Post("/analytics/outputs/s3-role-based", &analyticsS3RoleBasedOutput, &responseModel, reqParams)
	return &responseModel, err
}

// Delete S3 Role-based Output
func (api *AnalyticsOutputsS3RoleBasedAPI) Delete(outputId string) (*model.S3RoleBasedOutput, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["output_id"] = outputId
	}

	var responseModel model.S3RoleBasedOutput
	err := api.apiClient.Delete("/analytics/outputs/s3-role-based/{output_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get S3 Role-based Output Details
func (api *AnalyticsOutputsS3RoleBasedAPI) Get(outputId string) (*model.S3RoleBasedOutput, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["output_id"] = outputId
	}

	var responseModel model.S3RoleBasedOutput
	err := api.apiClient.Get("/analytics/outputs/s3-role-based/{output_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List S3 Role-based Outputs
func (api *AnalyticsOutputsS3RoleBasedAPI) List(queryParams ...func(*AnalyticsOutputsS3RoleBasedAPIListQueryParams)) (*pagination.AnalyticsS3RoleBasedOutputsListPagination, error) {
	queryParameters := &AnalyticsOutputsS3RoleBasedAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.AnalyticsS3RoleBasedOutputsListPagination
	err := api.apiClient.Get("/analytics/outputs/s3-role-based", nil, &responseModel, reqParams)
	return &responseModel, err
}

// AnalyticsOutputsS3RoleBasedAPIListQueryParams contains all query parameters for the List endpoint
type AnalyticsOutputsS3RoleBasedAPIListQueryParams struct {
	Offset int32  `query:"offset"`
	Limit  int32  `query:"limit"`
	Name   string `query:"name"`
}

// Params will return a map of query parameters
func (q *AnalyticsOutputsS3RoleBasedAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
