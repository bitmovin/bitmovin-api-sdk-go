package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingManifestsHlsMediaCustomTagApi struct {
    apiClient *common.ApiClient
}

func NewEncodingManifestsHlsMediaCustomTagApi(configs ...func(*common.ApiClient)) (*EncodingManifestsHlsMediaCustomTagApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingManifestsHlsMediaCustomTagApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingManifestsHlsMediaCustomTagApi) List(manifestId string, mediaId string, queryParams ...func(*query.CustomTagListQueryParams)) (*pagination.CustomTagsListPagination, error) {
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
    err := api.apiClient.Get("/encoding/manifests/hls/{manifest_id}/media/{media_id}/custom-tag", &resp, reqParams)
    return resp, err
}
func (api *EncodingManifestsHlsMediaCustomTagApi) Get(manifestId string, mediaId string, customTagId string) (*model.CustomTag, error) {
    var resp *model.CustomTag
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["media_id"] = mediaId
        params.PathParams["custom_tag_id"] = customTagId
	}
    err := api.apiClient.Get("/encoding/manifests/hls/{manifest_id}/media/{media_id}/custom-tag/{custom_tag_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingManifestsHlsMediaCustomTagApi) Create(manifestId string, mediaId string, customTag model.CustomTag) (*model.CustomTag, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["media_id"] = mediaId
    }
    payload := model.CustomTag(customTag)
    
    err := api.apiClient.Post("/encoding/manifests/hls/{manifest_id}/media/{media_id}/custom-tag", &payload, reqParams)
    return &payload, err
}
func (api *EncodingManifestsHlsMediaCustomTagApi) Delete(manifestId string, mediaId string, customTagId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["media_id"] = mediaId
        params.PathParams["custom_tag_id"] = customTagId
	}
    err := api.apiClient.Delete("/encoding/manifests/hls/{manifest_id}/media/{media_id}/custom-tag/{custom_tag_id}", &resp, reqParams)
    return resp, err
}
