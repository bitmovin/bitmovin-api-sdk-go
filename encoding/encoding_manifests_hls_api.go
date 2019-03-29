package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingManifestsHlsApi struct {
    apiClient *common.ApiClient
    Default *EncodingManifestsHlsDefaultApi
    Customdata *EncodingManifestsHlsCustomdataApi
    Streams *EncodingManifestsHlsStreamsApi
    Media *EncodingManifestsHlsMediaApi
}

func NewEncodingManifestsHlsApi(configs ...func(*common.ApiClient)) (*EncodingManifestsHlsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingManifestsHlsApi{apiClient: apiClient}

    defaultApi, err := NewEncodingManifestsHlsDefaultApi(configs...)
    api.Default = defaultApi
    customdataApi, err := NewEncodingManifestsHlsCustomdataApi(configs...)
    api.Customdata = customdataApi
    streamsApi, err := NewEncodingManifestsHlsStreamsApi(configs...)
    api.Streams = streamsApi
    mediaApi, err := NewEncodingManifestsHlsMediaApi(configs...)
    api.Media = mediaApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingManifestsHlsApi) Status(manifestId string) (*model.ModelTask, error) {
    var resp *model.ModelTask
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
	}
    err := api.apiClient.Get("/encoding/manifests/hls/{manifest_id}/status", &resp, reqParams)
    return resp, err
}
func (api *EncodingManifestsHlsApi) Create(hlsManifest model.HlsManifest) (*model.HlsManifest, error) {
    payload := model.HlsManifest(hlsManifest)
    
    err := api.apiClient.Post("/encoding/manifests/hls", &payload)
    return &payload, err
}
func (api *EncodingManifestsHlsApi) Delete(manifestId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
	}
    err := api.apiClient.Delete("/encoding/manifests/hls/{manifest_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingManifestsHlsApi) Start(manifestId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
    }
    
    var payload model.BitmovinResponse
    err := api.apiClient.Post("/encoding/manifests/hls/{manifest_id}/start", &payload, reqParams)
    return &payload, err
}
func (api *EncodingManifestsHlsApi) Stop(manifestId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
    }
    
    var payload model.BitmovinResponse
    err := api.apiClient.Post("/encoding/manifests/hls/{manifest_id}/stop", &payload, reqParams)
    return &payload, err
}
func (api *EncodingManifestsHlsApi) Get(manifestId string) (*model.HlsManifest, error) {
    var resp *model.HlsManifest
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
	}
    err := api.apiClient.Get("/encoding/manifests/hls/{manifest_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingManifestsHlsApi) List(queryParams ...func(*query.HlsManifestListQueryParams)) (*pagination.HlsManifestsListPagination, error) {
    queryParameters := &query.HlsManifestListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.HlsManifestsListPagination
    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/manifests/hls", &resp, reqParams)
    return resp, err
}
