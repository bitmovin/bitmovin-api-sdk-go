package analytics
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type AnalyticsOutputsS3RoleBasedApi struct {
    apiClient *common.ApiClient
    Customdata *AnalyticsOutputsS3RoleBasedCustomdataApi
}

func NewAnalyticsOutputsS3RoleBasedApi(configs ...func(*common.ApiClient)) (*AnalyticsOutputsS3RoleBasedApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &AnalyticsOutputsS3RoleBasedApi{apiClient: apiClient}

    customdataApi, err := NewAnalyticsOutputsS3RoleBasedCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *AnalyticsOutputsS3RoleBasedApi) Delete(outputId string) (*model.S3RoleBasedOutput, error) {
    var resp *model.S3RoleBasedOutput
    reqParams := func(params *common.RequestParams) {
        params.PathParams["output_id"] = outputId
	}
    err := api.apiClient.Delete("/analytics/outputs/s3-role-based/{output_id}", &resp, reqParams)
    return resp, err
}
func (api *AnalyticsOutputsS3RoleBasedApi) List(queryParams ...func(*query.S3RoleBasedOutputListQueryParams)) (*pagination.S3RoleBasedOutputsListPagination, error) {
    queryParameters := &query.S3RoleBasedOutputListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.S3RoleBasedOutputsListPagination
    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/analytics/outputs/s3-role-based", &resp, reqParams)
    return resp, err
}
func (api *AnalyticsOutputsS3RoleBasedApi) Get(outputId string) (*model.S3RoleBasedOutput, error) {
    var resp *model.S3RoleBasedOutput
    reqParams := func(params *common.RequestParams) {
        params.PathParams["output_id"] = outputId
	}
    err := api.apiClient.Get("/analytics/outputs/s3-role-based/{output_id}", &resp, reqParams)
    return resp, err
}
func (api *AnalyticsOutputsS3RoleBasedApi) Create(s3RoleBasedOutput model.S3RoleBasedOutput) (*model.S3RoleBasedOutput, error) {
    payload := model.S3RoleBasedOutput(s3RoleBasedOutput)
    
    err := api.apiClient.Post("/analytics/outputs/s3-role-based", &payload)
    return &payload, err
}
