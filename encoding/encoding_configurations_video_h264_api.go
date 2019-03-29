package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingConfigurationsVideoH264Api struct {
    apiClient *common.ApiClient
    Customdata *EncodingConfigurationsVideoH264CustomdataApi
}

func NewEncodingConfigurationsVideoH264Api(configs ...func(*common.ApiClient)) (*EncodingConfigurationsVideoH264Api, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingConfigurationsVideoH264Api{apiClient: apiClient}

    customdataApi, err := NewEncodingConfigurationsVideoH264CustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingConfigurationsVideoH264Api) Delete(configurationId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
	}
    err := api.apiClient.Delete("/encoding/configurations/video/h264/{configuration_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingConfigurationsVideoH264Api) List(queryParams ...func(*query.H264VideoConfigurationListQueryParams)) (*pagination.H264VideoConfigurationsListPagination, error) {
    queryParameters := &query.H264VideoConfigurationListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.H264VideoConfigurationsListPagination
    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/configurations/video/h264", &resp, reqParams)
    return resp, err
}
func (api *EncodingConfigurationsVideoH264Api) Get(configurationId string) (*model.H264VideoConfiguration, error) {
    var resp *model.H264VideoConfiguration
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
	}
    err := api.apiClient.Get("/encoding/configurations/video/h264/{configuration_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingConfigurationsVideoH264Api) Create(h264VideoConfiguration model.H264VideoConfiguration) (*model.H264VideoConfiguration, error) {
    payload := model.H264VideoConfiguration(h264VideoConfiguration)
    
    err := api.apiClient.Post("/encoding/configurations/video/h264", &payload)
    return &payload, err
}
