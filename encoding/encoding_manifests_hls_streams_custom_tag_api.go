package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingManifestsHlsStreamsCustomTagApi struct {
    apiClient *common.ApiClient
}

func NewEncodingManifestsHlsStreamsCustomTagApi(configs ...func(*common.ApiClient)) (*EncodingManifestsHlsStreamsCustomTagApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingManifestsHlsStreamsCustomTagApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingManifestsHlsStreamsCustomTagApi) List(manifestId string, streamId string, queryParams ...func(*query.CustomTagListQueryParams)) (*pagination.CustomTagsListPagination, error) {
    queryParameters := &query.CustomTagListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.CustomTagsListPagination
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["stream_id"] = streamId
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/manifests/hls/{manifest_id}/streams/{stream_id}/custom-tag", &resp, reqParams)
    return resp, err
}
func (api *EncodingManifestsHlsStreamsCustomTagApi) Create(manifestId string, streamId string, customTag model.CustomTag) (*model.CustomTag, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["stream_id"] = streamId
    }
    payload := model.CustomTag(customTag)
    
    err := api.apiClient.Post("/encoding/manifests/hls/{manifest_id}/streams/{stream_id}/custom-tag", &payload, reqParams)
    return &payload, err
}
func (api *EncodingManifestsHlsStreamsCustomTagApi) Get(manifestId string, streamId string, customTagId string) (*model.CustomTag, error) {
    var resp *model.CustomTag
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["stream_id"] = streamId
        params.PathParams["custom_tag_id"] = customTagId
	}
    err := api.apiClient.Get("/encoding/manifests/hls/{manifest_id}/streams/{stream_id}/custom-tag/{custom_tag_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingManifestsHlsStreamsCustomTagApi) Delete(manifestId string, streamId string, customTagId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["stream_id"] = streamId
        params.PathParams["custom_tag_id"] = customTagId
	}
    err := api.apiClient.Delete("/encoding/manifests/hls/{manifest_id}/streams/{stream_id}/custom-tag/{custom_tag_id}", &resp, reqParams)
    return resp, err
}
