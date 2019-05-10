package analytics
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type AnalyticsExportsApi struct {
    apiClient *common.ApiClient
}

func NewAnalyticsExportsApi(configs ...func(*common.ApiClient)) (*AnalyticsExportsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &AnalyticsExportsApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *AnalyticsExportsApi) List(queryParams ...func(*query.AnalyticsExportTaskListQueryParams)) (*pagination.AnalyticsExportTasksListPagination, error) {
    queryParameters := &query.AnalyticsExportTaskListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.AnalyticsExportTasksListPagination
    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/analytics/exports", &resp, reqParams)
    return resp, err
}
func (api *AnalyticsExportsApi) Create(analyticsExportTask model.AnalyticsExportTask) (*model.AnalyticsExportTask, error) {
    payload := model.AnalyticsExportTask(analyticsExportTask)
    
    err := api.apiClient.Post("/analytics/exports", &payload)
    return &payload, err
}
func (api *AnalyticsExportsApi) Get(exportTaskId string) (*model.AnalyticsExportTask, error) {
    var resp *model.AnalyticsExportTask
    reqParams := func(params *common.RequestParams) {
        params.PathParams["export_task_id"] = exportTaskId
	}
    err := api.apiClient.Get("/analytics/exports/{export_task_id}", &resp, reqParams)
    return resp, err
}
