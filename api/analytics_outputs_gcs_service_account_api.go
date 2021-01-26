package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// AnalyticsOutputsGcsServiceAccountAPI communicates with '/analytics/outputs/gcs-service-account' endpoints
type AnalyticsOutputsGcsServiceAccountAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/analytics/outputs/gcs-service-account/{output_id}/customData' endpoints
	Customdata *AnalyticsOutputsGcsServiceAccountCustomdataAPI
}

// NewAnalyticsOutputsGcsServiceAccountAPI constructor for AnalyticsOutputsGcsServiceAccountAPI that takes options as argument
func NewAnalyticsOutputsGcsServiceAccountAPI(options ...apiclient.APIClientOption) (*AnalyticsOutputsGcsServiceAccountAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewAnalyticsOutputsGcsServiceAccountAPIWithClient(apiClient), nil
}

// NewAnalyticsOutputsGcsServiceAccountAPIWithClient constructor for AnalyticsOutputsGcsServiceAccountAPI that takes an APIClient as argument
func NewAnalyticsOutputsGcsServiceAccountAPIWithClient(apiClient *apiclient.APIClient) *AnalyticsOutputsGcsServiceAccountAPI {
	a := &AnalyticsOutputsGcsServiceAccountAPI{apiClient: apiClient}
	a.Customdata = NewAnalyticsOutputsGcsServiceAccountCustomdataAPIWithClient(apiClient)

	return a
}

// Create Service Account based GCS Output
func (api *AnalyticsOutputsGcsServiceAccountAPI) Create(gcsServiceAccountOutput model.GcsServiceAccountOutput) (*model.GcsServiceAccountOutput, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.GcsServiceAccountOutput
	err := api.apiClient.Post("/analytics/outputs/gcs-service-account", &gcsServiceAccountOutput, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Service Account based GCS Output
func (api *AnalyticsOutputsGcsServiceAccountAPI) Delete(outputId string) (*model.GcsServiceAccountOutput, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["output_id"] = outputId
	}

	var responseModel model.GcsServiceAccountOutput
	err := api.apiClient.Delete("/analytics/outputs/gcs-service-account/{output_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Service Account based GCS Output Details
func (api *AnalyticsOutputsGcsServiceAccountAPI) Get(outputId string) (*model.GcsServiceAccountOutput, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["output_id"] = outputId
	}

	var responseModel model.GcsServiceAccountOutput
	err := api.apiClient.Get("/analytics/outputs/gcs-service-account/{output_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List Service Account based GCS Outputs
func (api *AnalyticsOutputsGcsServiceAccountAPI) List(queryParams ...func(*AnalyticsOutputsGcsServiceAccountAPIListQueryParams)) (*pagination.GcsServiceAccountOutputsListPagination, error) {
	queryParameters := &AnalyticsOutputsGcsServiceAccountAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.GcsServiceAccountOutputsListPagination
	err := api.apiClient.Get("/analytics/outputs/gcs-service-account", nil, &responseModel, reqParams)
	return &responseModel, err
}

// AnalyticsOutputsGcsServiceAccountAPIListQueryParams contains all query parameters for the List endpoint
type AnalyticsOutputsGcsServiceAccountAPIListQueryParams struct {
	Offset int32  `query:"offset"`
	Limit  int32  `query:"limit"`
	Name   string `query:"name"`
}

// Params will return a map of query parameters
func (q *AnalyticsOutputsGcsServiceAccountAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
