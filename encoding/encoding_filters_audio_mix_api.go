package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingFiltersAudioMixApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingFiltersAudioMixCustomdataApi
}

func NewEncodingFiltersAudioMixApi(configs ...func(*common.ApiClient)) (*EncodingFiltersAudioMixApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingFiltersAudioMixApi{apiClient: apiClient}

    customdataApi, err := NewEncodingFiltersAudioMixCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingFiltersAudioMixApi) Create(audioMixFilter model.AudioMixFilter) (*model.AudioMixFilter, error) {
    payload := model.AudioMixFilter(audioMixFilter)
    
    err := api.apiClient.Post("/encoding/filters/audio-mix", &payload)
    return &payload, err
}
func (api *EncodingFiltersAudioMixApi) Get(filterId string) (*model.AudioMixFilter, error) {
    var resp *model.AudioMixFilter
    reqParams := func(params *common.RequestParams) {
        params.PathParams["filter_id"] = filterId
	}
    err := api.apiClient.Get("/encoding/filters/audio-mix/{filter_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingFiltersAudioMixApi) List(queryParams ...func(*query.AudioMixFilterListQueryParams)) (*pagination.AudioMixFiltersListPagination, error) {
    queryParameters := &query.AudioMixFilterListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.AudioMixFiltersListPagination
    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/filters/audio-mix", &resp, reqParams)
    return resp, err
}
func (api *EncodingFiltersAudioMixApi) Delete(filterId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["filter_id"] = filterId
	}
    err := api.apiClient.Delete("/encoding/filters/audio-mix/{filter_id}", &resp, reqParams)
    return resp, err
}
