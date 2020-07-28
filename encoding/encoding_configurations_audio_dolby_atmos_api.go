package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingConfigurationsAudioDolbyAtmosApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingConfigurationsAudioDolbyAtmosCustomdataApi
}

func NewEncodingConfigurationsAudioDolbyAtmosApi(configs ...func(*common.ApiClient)) (*EncodingConfigurationsAudioDolbyAtmosApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingConfigurationsAudioDolbyAtmosApi{apiClient: apiClient}

    customdataApi, err := NewEncodingConfigurationsAudioDolbyAtmosCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingConfigurationsAudioDolbyAtmosApi) Create(dolbyAtmosAudioConfiguration model.DolbyAtmosAudioConfiguration) (*model.DolbyAtmosAudioConfiguration, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.DolbyAtmosAudioConfiguration
    err := api.apiClient.Post("/encoding/configurations/audio/dolby-atmos", &dolbyAtmosAudioConfiguration, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingConfigurationsAudioDolbyAtmosApi) Delete(configurationId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/configurations/audio/dolby-atmos/{configuration_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingConfigurationsAudioDolbyAtmosApi) Get(configurationId string) (*model.DolbyAtmosAudioConfiguration, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel *model.DolbyAtmosAudioConfiguration
    err := api.apiClient.Get("/encoding/configurations/audio/dolby-atmos/{configuration_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingConfigurationsAudioDolbyAtmosApi) List(queryParams ...func(*query.DolbyAtmosAudioConfigurationListQueryParams)) (*pagination.DolbyAtmosAudioConfigurationsListPagination, error) {
    queryParameters := &query.DolbyAtmosAudioConfigurationListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.DolbyAtmosAudioConfigurationsListPagination
    err := api.apiClient.Get("/encoding/configurations/audio/dolby-atmos", nil, &responseModel, reqParams)
    return responseModel, err
}

