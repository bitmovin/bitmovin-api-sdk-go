package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingManifestsHlsStreamsApi struct {
    apiClient *common.ApiClient
    CustomTag *EncodingManifestsHlsStreamsCustomTagApi
    Iframe *EncodingManifestsHlsStreamsIframeApi
}

func NewEncodingManifestsHlsStreamsApi(configs ...func(*common.ApiClient)) (*EncodingManifestsHlsStreamsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingManifestsHlsStreamsApi{apiClient: apiClient}

    customTagApi, err := NewEncodingManifestsHlsStreamsCustomTagApi(configs...)
    api.CustomTag = customTagApi
    iframeApi, err := NewEncodingManifestsHlsStreamsIframeApi(configs...)
    api.Iframe = iframeApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingManifestsHlsStreamsApi) Get(manifestId string, streamId string) (*model.StreamInfo, error) {
    var resp *model.StreamInfo
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["stream_id"] = streamId
	}
    err := api.apiClient.Get("/encoding/manifests/hls/{manifest_id}/streams/{stream_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingManifestsHlsStreamsApi) Create(manifestId string, streamInfo model.StreamInfo) (*model.StreamInfo, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
    }
    payload := model.StreamInfo(streamInfo)
    
    err := api.apiClient.Post("/encoding/manifests/hls/{manifest_id}/streams", &payload, reqParams)
    return &payload, err
}
func (api *EncodingManifestsHlsStreamsApi) List(manifestId string, queryParams ...func(*query.StreamInfoListQueryParams)) (*pagination.StreamInfosListPagination, error) {
    queryParameters := &query.StreamInfoListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.StreamInfosListPagination
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/manifests/hls/{manifest_id}/streams", &resp, reqParams)
    return resp, err
}
func (api *EncodingManifestsHlsStreamsApi) Delete(manifestId string, streamId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["stream_id"] = streamId
	}
    err := api.apiClient.Delete("/encoding/manifests/hls/{manifest_id}/streams/{stream_id}", &resp, reqParams)
    return resp, err
}
