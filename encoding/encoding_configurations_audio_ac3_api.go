package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingConfigurationsAudioAc3Api struct {
    apiClient *common.ApiClient
    Customdata *EncodingConfigurationsAudioAc3CustomdataApi
}

func NewEncodingConfigurationsAudioAc3Api(configs ...func(*common.ApiClient)) (*EncodingConfigurationsAudioAc3Api, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingConfigurationsAudioAc3Api{apiClient: apiClient}

    customdataApi, err := NewEncodingConfigurationsAudioAc3CustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingConfigurationsAudioAc3Api) Delete(configurationId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
	}
    err := api.apiClient.Delete("/encoding/configurations/audio/ac3/{configuration_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingConfigurationsAudioAc3Api) List(queryParams ...func(*query.Ac3AudioConfigurationListQueryParams)) (*pagination.Ac3AudioConfigurationsListPagination, error) {
    queryParameters := &query.Ac3AudioConfigurationListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.Ac3AudioConfigurationsListPagination
    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/configurations/audio/ac3", &resp, reqParams)
    return resp, err
}
func (api *EncodingConfigurationsAudioAc3Api) Get(configurationId string) (*model.Ac3AudioConfiguration, error) {
    var resp *model.Ac3AudioConfiguration
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
	}
    err := api.apiClient.Get("/encoding/configurations/audio/ac3/{configuration_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingConfigurationsAudioAc3Api) Create(ac3AudioConfiguration model.Ac3AudioConfiguration) (*model.Ac3AudioConfiguration, error) {
    payload := model.Ac3AudioConfiguration(ac3AudioConfiguration)
    
    err := api.apiClient.Post("/encoding/configurations/audio/ac3", &payload)
    return &payload, err
}
