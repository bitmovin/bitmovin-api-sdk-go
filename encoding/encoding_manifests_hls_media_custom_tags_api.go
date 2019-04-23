package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingManifestsHlsMediaCustomTagsApi struct {
    apiClient *common.ApiClient
}

func NewEncodingManifestsHlsMediaCustomTagsApi(configs ...func(*common.ApiClient)) (*EncodingManifestsHlsMediaCustomTagsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingManifestsHlsMediaCustomTagsApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingManifestsHlsMediaCustomTagsApi) Delete(manifestId string, mediaId string, customTagId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["media_id"] = mediaId
        params.PathParams["custom_tag_id"] = customTagId
	}
    err := api.apiClient.Delete("/encoding/manifests/hls/{manifest_id}/media/{media_id}/custom-tags/{custom_tag_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingManifestsHlsMediaCustomTagsApi) Get(manifestId string, mediaId string, customTagId string) (*model.CustomTag, error) {
    var resp *model.CustomTag
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["media_id"] = mediaId
        params.PathParams["custom_tag_id"] = customTagId
	}
    err := api.apiClient.Get("/encoding/manifests/hls/{manifest_id}/media/{media_id}/custom-tags/{custom_tag_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingManifestsHlsMediaCustomTagsApi) List(manifestId string, mediaId string, queryParams ...func(*query.CustomTagListQueryParams)) (*pagination.CustomTagsListPagination, error) {
    queryParameters := &query.CustomTagListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.CustomTagsListPagination
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["media_id"] = mediaId
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/manifests/hls/{manifest_id}/media/{media_id}/custom-tags", &resp, reqParams)
    return resp, err
}
func (api *EncodingManifestsHlsMediaCustomTagsApi) Create(manifestId string, mediaId string, customTag model.CustomTag) (*model.CustomTag, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["media_id"] = mediaId
    }
    payload := model.CustomTag(customTag)
    
    err := api.apiClient.Post("/encoding/manifests/hls/{manifest_id}/media/{media_id}/custom-tags", &payload, reqParams)
    return &payload, err
}
