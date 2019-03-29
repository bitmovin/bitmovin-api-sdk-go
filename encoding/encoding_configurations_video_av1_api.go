package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingConfigurationsVideoAv1Api struct {
    apiClient *common.ApiClient
    Customdata *EncodingConfigurationsVideoAv1CustomdataApi
}

func NewEncodingConfigurationsVideoAv1Api(configs ...func(*common.ApiClient)) (*EncodingConfigurationsVideoAv1Api, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingConfigurationsVideoAv1Api{apiClient: apiClient}

    customdataApi, err := NewEncodingConfigurationsVideoAv1CustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingConfigurationsVideoAv1Api) Create(av1VideoConfiguration model.Av1VideoConfiguration) (*model.Av1VideoConfiguration, error) {
    payload := model.Av1VideoConfiguration(av1VideoConfiguration)
    
    err := api.apiClient.Post("/encoding/configurations/video/av1", &payload)
    return &payload, err
}
func (api *EncodingConfigurationsVideoAv1Api) Get(configurationId string) (*model.Av1VideoConfiguration, error) {
    var resp *model.Av1VideoConfiguration
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
	}
    err := api.apiClient.Get("/encoding/configurations/video/av1/{configuration_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingConfigurationsVideoAv1Api) Delete(configurationId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
	}
    err := api.apiClient.Delete("/encoding/configurations/video/av1/{configuration_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingConfigurationsVideoAv1Api) List(queryParams ...func(*query.Av1VideoConfigurationListQueryParams)) (*pagination.Av1VideoConfigurationsListPagination, error) {
    queryParameters := &query.Av1VideoConfigurationListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.Av1VideoConfigurationsListPagination
    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/configurations/video/av1", &resp, reqParams)
    return resp, err
}
