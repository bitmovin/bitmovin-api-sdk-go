package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingConfigurationsAudioDtsPassthroughApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingConfigurationsAudioDtsPassthroughCustomdataApi
}

func NewEncodingConfigurationsAudioDtsPassthroughApi(configs ...func(*common.ApiClient)) (*EncodingConfigurationsAudioDtsPassthroughApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingConfigurationsAudioDtsPassthroughApi{apiClient: apiClient}

    customdataApi, err := NewEncodingConfigurationsAudioDtsPassthroughCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingConfigurationsAudioDtsPassthroughApi) Create(dtsPassthroughAudioConfiguration model.DtsPassthroughAudioConfiguration) (*model.DtsPassthroughAudioConfiguration, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.DtsPassthroughAudioConfiguration
    err := api.apiClient.Post("/encoding/configurations/audio/dts-passthrough", &dtsPassthroughAudioConfiguration, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingConfigurationsAudioDtsPassthroughApi) Delete(configurationId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/configurations/audio/dts-passthrough/{configuration_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingConfigurationsAudioDtsPassthroughApi) Get(configurationId string) (*model.DtsPassthroughAudioConfiguration, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["configuration_id"] = configurationId
    }

    var responseModel *model.DtsPassthroughAudioConfiguration
    err := api.apiClient.Get("/encoding/configurations/audio/dts-passthrough/{configuration_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingConfigurationsAudioDtsPassthroughApi) List(queryParams ...func(*query.DtsPassthroughAudioConfigurationListQueryParams)) (*pagination.DtsPassthroughAudioConfigurationsListPagination, error) {
    queryParameters := &query.DtsPassthroughAudioConfigurationListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.DtsPassthroughAudioConfigurationsListPagination
    err := api.apiClient.Get("/encoding/configurations/audio/dts-passthrough", nil, &responseModel, reqParams)
    return responseModel, err
}

