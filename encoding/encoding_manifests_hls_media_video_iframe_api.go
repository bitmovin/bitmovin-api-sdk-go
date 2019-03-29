package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingManifestsHlsMediaVideoIframeApi struct {
    apiClient *common.ApiClient
}

func NewEncodingManifestsHlsMediaVideoIframeApi(configs ...func(*common.ApiClient)) (*EncodingManifestsHlsMediaVideoIframeApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingManifestsHlsMediaVideoIframeApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingManifestsHlsMediaVideoIframeApi) Create(manifestId string, mediaId string, iFramePlaylist model.IFramePlaylist) (*model.IFramePlaylist, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["media_id"] = mediaId
    }
    payload := model.IFramePlaylist(iFramePlaylist)
    
    err := api.apiClient.Post("/encoding/manifests/hls/{manifest_id}/media/video/{media_id}/iframe", &payload, reqParams)
    return &payload, err
}
func (api *EncodingManifestsHlsMediaVideoIframeApi) List(manifestId string, mediaId string, queryParams ...func(*query.IFramePlaylistListQueryParams)) (*pagination.IFramePlaylistsListPagination, error) {
    queryParameters := &query.IFramePlaylistListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.IFramePlaylistsListPagination
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["media_id"] = mediaId
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/manifests/hls/{manifest_id}/media/video/{media_id}/iframe", &resp, reqParams)
    return resp, err
}
func (api *EncodingManifestsHlsMediaVideoIframeApi) Get(manifestId string, mediaId string, iframeId string) (*model.IFramePlaylist, error) {
    var resp *model.IFramePlaylist
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["media_id"] = mediaId
        params.PathParams["iframe_id"] = iframeId
	}
    err := api.apiClient.Get("/encoding/manifests/hls/{manifest_id}/media/video/{media_id}/iframe/{iframe_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingManifestsHlsMediaVideoIframeApi) Delete(manifestId string, mediaId string, iframeId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["media_id"] = mediaId
        params.PathParams["iframe_id"] = iframeId
	}
    err := api.apiClient.Delete("/encoding/manifests/hls/{manifest_id}/media/video/{media_id}/iframe/{iframe_id}", &resp, reqParams)
    return resp, err
}
