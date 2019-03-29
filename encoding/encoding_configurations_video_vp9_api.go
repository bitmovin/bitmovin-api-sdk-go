package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingConfigurationsVideoVp9Api struct {
    apiClient *common.ApiClient
    Customdata *EncodingConfigurationsVideoVp9CustomdataApi
}

func NewEncodingConfigurationsVideoVp9Api(configs ...func(*common.ApiClient)) (*EncodingConfigurationsVideoVp9Api, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingConfigurationsVideoVp9Api{apiClient: apiClient}

    customdataApi, err := NewEncodingConfigurationsVideoVp9CustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingConfigurationsVideoVp9Api) Create(vp9VideoConfiguration model.Vp9VideoConfiguration) (*model.Vp9VideoConfiguration, error) {
    payload := model.Vp9VideoConfiguration(vp9VideoConfiguration)
    
    err := api.apiClient.Post("/encoding/configurations/video/vp9", &payload)
    return &payload, err
}
func (api *EncodingConfigurationsVideoVp9Api) Delete(configurationId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
	}
    err := api.apiClient.Delete("/encoding/configurations/video/vp9/{configuration_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingConfigurationsVideoVp9Api) Get(configurationId string) (*model.Vp9VideoConfiguration, error) {
    var resp *model.Vp9VideoConfiguration
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
	}
    err := api.apiClient.Get("/encoding/configurations/video/vp9/{configuration_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingConfigurationsVideoVp9Api) List(queryParams ...func(*query.Vp9VideoConfigurationListQueryParams)) (*pagination.Vp9VideoConfigurationsListPagination, error) {
    queryParameters := &query.Vp9VideoConfigurationListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.Vp9VideoConfigurationsListPagination
    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/configurations/video/vp9", &resp, reqParams)
    return resp, err
}
