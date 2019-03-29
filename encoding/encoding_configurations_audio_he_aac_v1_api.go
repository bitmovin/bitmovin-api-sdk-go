package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingConfigurationsAudioHeAacV1Api struct {
    apiClient *common.ApiClient
    Customdata *EncodingConfigurationsAudioHeAacV1CustomdataApi
}

func NewEncodingConfigurationsAudioHeAacV1Api(configs ...func(*common.ApiClient)) (*EncodingConfigurationsAudioHeAacV1Api, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingConfigurationsAudioHeAacV1Api{apiClient: apiClient}

    customdataApi, err := NewEncodingConfigurationsAudioHeAacV1CustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingConfigurationsAudioHeAacV1Api) Get(configurationId string) (*model.HeAacV1AudioConfiguration, error) {
    var resp *model.HeAacV1AudioConfiguration
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
	}
    err := api.apiClient.Get("/encoding/configurations/audio/he-aac-v1/{configuration_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingConfigurationsAudioHeAacV1Api) Delete(configurationId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
	}
    err := api.apiClient.Delete("/encoding/configurations/audio/he-aac-v1/{configuration_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingConfigurationsAudioHeAacV1Api) List(queryParams ...func(*query.HeAacV1AudioConfigurationListQueryParams)) (*pagination.HeAacV1AudioConfigurationsListPagination, error) {
    queryParameters := &query.HeAacV1AudioConfigurationListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.HeAacV1AudioConfigurationsListPagination
    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/configurations/audio/he-aac-v1", &resp, reqParams)
    return resp, err
}
func (api *EncodingConfigurationsAudioHeAacV1Api) Create(heAacV1AudioConfiguration model.HeAacV1AudioConfiguration) (*model.HeAacV1AudioConfiguration, error) {
    payload := model.HeAacV1AudioConfiguration(heAacV1AudioConfiguration)
    
    err := api.apiClient.Post("/encoding/configurations/audio/he-aac-v1", &payload)
    return &payload, err
}
