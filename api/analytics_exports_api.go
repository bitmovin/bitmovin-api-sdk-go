package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// AnalyticsExportsAPI communicates with '/analytics/exports' endpoints
type AnalyticsExportsAPI struct {
	apiClient *apiclient.APIClient
}

// NewAnalyticsExportsAPI constructor for AnalyticsExportsAPI that takes options as argument
func NewAnalyticsExportsAPI(options ...apiclient.APIClientOption) (*AnalyticsExportsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewAnalyticsExportsAPIWithClient(apiClient), nil
}

// NewAnalyticsExportsAPIWithClient constructor for AnalyticsExportsAPI that takes an APIClient as argument
func NewAnalyticsExportsAPIWithClient(apiClient *apiclient.APIClient) *AnalyticsExportsAPI {
	a := &AnalyticsExportsAPI{apiClient: apiClient}
	return a
}

// Create Export Task
func (api *AnalyticsExportsAPI) Create(analyticsExportTask model.AnalyticsExportTask) (*model.AnalyticsExportTask, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.AnalyticsExportTask
	err := api.apiClient.Post("/analytics/exports", &analyticsExportTask, &responseModel, reqParams)
	return &responseModel, err
}

// Get export task
func (api *AnalyticsExportsAPI) Get(exportTaskId string) (*model.AnalyticsExportTask, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["export_task_id"] = exportTaskId
	}

	var responseModel model.AnalyticsExportTask
	err := api.apiClient.Get("/analytics/exports/{export_task_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List Export Tasks
func (api *AnalyticsExportsAPI) List(queryParams ...func(*AnalyticsExportsAPIListQueryParams)) (*pagination.AnalyticsExportTasksListPagination, error) {
	queryParameters := &AnalyticsExportsAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.AnalyticsExportTasksListPagination
	err := api.apiClient.Get("/analytics/exports", nil, &responseModel, reqParams)
	return &responseModel, err
}

// AnalyticsExportsAPIListQueryParams contains all query parameters for the List endpoint
type AnalyticsExportsAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *AnalyticsExportsAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
