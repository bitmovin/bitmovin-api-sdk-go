package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingConfigurationsAudioHeAacV2Api struct {
    apiClient *common.ApiClient
    Customdata *EncodingConfigurationsAudioHeAacV2CustomdataApi
}

func NewEncodingConfigurationsAudioHeAacV2Api(configs ...func(*common.ApiClient)) (*EncodingConfigurationsAudioHeAacV2Api, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingConfigurationsAudioHeAacV2Api{apiClient: apiClient}

    customdataApi, err := NewEncodingConfigurationsAudioHeAacV2CustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingConfigurationsAudioHeAacV2Api) Delete(configurationId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
	}
    err := api.apiClient.Delete("/encoding/configurations/audio/he-aac-v2/{configuration_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingConfigurationsAudioHeAacV2Api) List(queryParams ...func(*query.HeAacV2AudioConfigurationListQueryParams)) (*pagination.HeAacV2AudioConfigurationsListPagination, error) {
    queryParameters := &query.HeAacV2AudioConfigurationListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.HeAacV2AudioConfigurationsListPagination
    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/configurations/audio/he-aac-v2", &resp, reqParams)
    return resp, err
}
func (api *EncodingConfigurationsAudioHeAacV2Api) Get(configurationId string) (*model.HeAacV2AudioConfiguration, error) {
    var resp *model.HeAacV2AudioConfiguration
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
	}
    err := api.apiClient.Get("/encoding/configurations/audio/he-aac-v2/{configuration_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingConfigurationsAudioHeAacV2Api) Create(heAacV2AudioConfiguration model.HeAacV2AudioConfiguration) (*model.HeAacV2AudioConfiguration, error) {
    payload := model.HeAacV2AudioConfiguration(heAacV2AudioConfiguration)
    
    err := api.apiClient.Post("/encoding/configurations/audio/he-aac-v2", &payload)
    return &payload, err
}
