package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingConfigurationsAudioPcmApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingConfigurationsAudioPcmCustomdataApi
}

func NewEncodingConfigurationsAudioPcmApi(configs ...func(*common.ApiClient)) (*EncodingConfigurationsAudioPcmApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingConfigurationsAudioPcmApi{apiClient: apiClient}

    customdataApi, err := NewEncodingConfigurationsAudioPcmCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingConfigurationsAudioPcmApi) Create(pcmAudioConfiguration model.PcmAudioConfiguration) (*model.PcmAudioConfiguration, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.PcmAudioConfiguration
    err := api.apiClient.Post("/encoding/configurations/audio/pcm", &pcmAudioConfiguration, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingConfigurationsAudioPcmApi) Delete(configurationId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/configurations/audio/pcm/{configuration_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingConfigurationsAudioPcmApi) Get(configurationId string) (*model.PcmAudioConfiguration, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel *model.PcmAudioConfiguration
    err := api.apiClient.Get("/encoding/configurations/audio/pcm/{configuration_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingConfigurationsAudioPcmApi) List(queryParams ...func(*query.PcmAudioConfigurationListQueryParams)) (*pagination.PcmAudioConfigurationsListPagination, error) {
    queryParameters := &query.PcmAudioConfigurationListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.PcmAudioConfigurationsListPagination
    err := api.apiClient.Get("/encoding/configurations/audio/pcm", nil, &responseModel, reqParams)
    return responseModel, err
}

