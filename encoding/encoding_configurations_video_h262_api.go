package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingConfigurationsVideoH262Api struct {
    apiClient *common.ApiClient
    Customdata *EncodingConfigurationsVideoH262CustomdataApi
}

func NewEncodingConfigurationsVideoH262Api(configs ...func(*common.ApiClient)) (*EncodingConfigurationsVideoH262Api, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingConfigurationsVideoH262Api{apiClient: apiClient}

    customdataApi, err := NewEncodingConfigurationsVideoH262CustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingConfigurationsVideoH262Api) Create(h262VideoConfiguration model.H262VideoConfiguration) (*model.H262VideoConfiguration, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.H262VideoConfiguration
    err := api.apiClient.Post("/encoding/configurations/video/h262", &h262VideoConfiguration, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingConfigurationsVideoH262Api) Delete(configurationId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/configurations/video/h262/{configuration_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingConfigurationsVideoH262Api) Get(configurationId string) (*model.H262VideoConfiguration, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel *model.H262VideoConfiguration
    err := api.apiClient.Get("/encoding/configurations/video/h262/{configuration_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingConfigurationsVideoH262Api) List(queryParams ...func(*query.H262VideoConfigurationListQueryParams)) (*pagination.H262VideoConfigurationsListPagination, error) {
    queryParameters := &query.H262VideoConfigurationListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.H262VideoConfigurationsListPagination
    err := api.apiClient.Get("/encoding/configurations/video/h262", nil, &responseModel, reqParams)
    return responseModel, err
}

