package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingInfrastructureGceRegionsApi struct {
    apiClient *common.ApiClient
}

func NewEncodingInfrastructureGceRegionsApi(configs ...func(*common.ApiClient)) (*EncodingInfrastructureGceRegionsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingInfrastructureGceRegionsApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingInfrastructureGceRegionsApi) Create(infrastructureId string, region model.GoogleCloudRegion, gceAccountRegionSettings model.GceAccountRegionSettings) (*model.GceAccountRegionSettings, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["infrastructure_id"] = infrastructureId
        params.PathParams["region"] = string(region)
    }

    var responseModel *model.GceAccountRegionSettings
    err := api.apiClient.Post("/encoding/infrastructure/gce/{infrastructure_id}/regions/{region}", &gceAccountRegionSettings, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingInfrastructureGceRegionsApi) Delete(infrastructureId string, region model.GoogleCloudRegion) (*model.GceAccountRegionSettings, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["infrastructure_id"] = infrastructureId
        params.PathParams["region"] = string(region)
    }

    var responseModel *model.GceAccountRegionSettings
    err := api.apiClient.Delete("/encoding/infrastructure/gce/{infrastructure_id}/regions/{region}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingInfrastructureGceRegionsApi) Get(infrastructureId string, region model.GoogleCloudRegion) (*model.GceAccountRegionSettings, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["infrastructure_id"] = infrastructureId
        params.PathParams["region"] = string(region)
    }

    var responseModel *model.GceAccountRegionSettings
    err := api.apiClient.Get("/encoding/infrastructure/gce/{infrastructure_id}/regions/{region}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingInfrastructureGceRegionsApi) List(infrastructureId string, queryParams ...func(*query.GceAccountRegionSettingsListQueryParams)) (*pagination.GceAccountRegionSettingssListPagination, error) {
    queryParameters := &query.GceAccountRegionSettingsListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["infrastructure_id"] = infrastructureId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.GceAccountRegionSettingssListPagination
    err := api.apiClient.Get("/encoding/infrastructure/gce/{infrastructure_id}/regions", nil, &responseModel, reqParams)
    return responseModel, err
}

