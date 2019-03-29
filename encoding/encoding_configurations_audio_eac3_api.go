package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingConfigurationsAudioEac3Api struct {
    apiClient *common.ApiClient
    Customdata *EncodingConfigurationsAudioEac3CustomdataApi
}

func NewEncodingConfigurationsAudioEac3Api(configs ...func(*common.ApiClient)) (*EncodingConfigurationsAudioEac3Api, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingConfigurationsAudioEac3Api{apiClient: apiClient}

    customdataApi, err := NewEncodingConfigurationsAudioEac3CustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingConfigurationsAudioEac3Api) Delete(configurationId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
	}
    err := api.apiClient.Delete("/encoding/configurations/audio/eac3/{configuration_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingConfigurationsAudioEac3Api) List(queryParams ...func(*query.Eac3AudioConfigurationListQueryParams)) (*pagination.Eac3AudioConfigurationsListPagination, error) {
    queryParameters := &query.Eac3AudioConfigurationListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.Eac3AudioConfigurationsListPagination
    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/configurations/audio/eac3", &resp, reqParams)
    return resp, err
}
func (api *EncodingConfigurationsAudioEac3Api) Get(configurationId string) (*model.Eac3AudioConfiguration, error) {
    var resp *model.Eac3AudioConfiguration
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
	}
    err := api.apiClient.Get("/encoding/configurations/audio/eac3/{configuration_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingConfigurationsAudioEac3Api) Create(eac3AudioConfiguration model.Eac3AudioConfiguration) (*model.Eac3AudioConfiguration, error) {
    payload := model.Eac3AudioConfiguration(eac3AudioConfiguration)
    
    err := api.apiClient.Post("/encoding/configurations/audio/eac3", &payload)
    return &payload, err
}
