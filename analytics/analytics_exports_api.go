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

func (api *AnalyticsExportsApi) Create(analyticsExportTask model.AnalyticsExportTaskDetails) (*model.AnalyticsExportTaskDetails, error) {
    payload := model.AnalyticsExportTaskDetails(analyticsExportTask)
    
    err := api.apiClient.Post("/analytics/exports", &payload)
    return &payload, err
}
func (api *AnalyticsExportsApi) List(queryParams ...func(*query.AnalyticsExportTaskDetailsListQueryParams)) (*pagination.AnalyticsExportTaskDetailssListPagination, error) {
    queryParameters := &query.AnalyticsExportTaskDetailsListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.AnalyticsExportTaskDetailssListPagination
    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/analytics/exports", &resp, reqParams)
    return resp, err
}
func (api *AnalyticsExportsApi) Get(exportTaskId string) (*model.AnalyticsExportTaskDetails, error) {
    var resp *model.AnalyticsExportTaskDetails
    reqParams := func(params *common.RequestParams) {
        params.PathParams["export_task_id"] = exportTaskId
	}
    err := api.apiClient.Get("/analytics/exports/{export_task_id}", &resp, reqParams)
    return resp, err
}
