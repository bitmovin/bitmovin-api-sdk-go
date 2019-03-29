package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingConfigurationsVideoVp8Api struct {
    apiClient *common.ApiClient
    Customdata *EncodingConfigurationsVideoVp8CustomdataApi
}

func NewEncodingConfigurationsVideoVp8Api(configs ...func(*common.ApiClient)) (*EncodingConfigurationsVideoVp8Api, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingConfigurationsVideoVp8Api{apiClient: apiClient}

    customdataApi, err := NewEncodingConfigurationsVideoVp8CustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingConfigurationsVideoVp8Api) Create(vp8VideoConfiguration model.Vp8VideoConfiguration) (*model.Vp8VideoConfiguration, error) {
    payload := model.Vp8VideoConfiguration(vp8VideoConfiguration)
    
    err := api.apiClient.Post("/encoding/configurations/video/vp8", &payload)
    return &payload, err
}
func (api *EncodingConfigurationsVideoVp8Api) Delete(configurationId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
	}
    err := api.apiClient.Delete("/encoding/configurations/video/vp8/{configuration_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingConfigurationsVideoVp8Api) List(queryParams ...func(*query.Vp8VideoConfigurationListQueryParams)) (*pagination.Vp8VideoConfigurationsListPagination, error) {
    queryParameters := &query.Vp8VideoConfigurationListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.Vp8VideoConfigurationsListPagination
    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/configurations/video/vp8", &resp, reqParams)
    return resp, err
}
func (api *EncodingConfigurationsVideoVp8Api) Get(configurationId string) (*model.Vp8VideoConfiguration, error) {
    var resp *model.Vp8VideoConfiguration
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
	}
    err := api.apiClient.Get("/encoding/configurations/video/vp8/{configuration_id}", &resp, reqParams)
    return resp, err
}
