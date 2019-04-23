package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingManifestsHlsMediaVideoApi struct {
    apiClient *common.ApiClient
}

func NewEncodingManifestsHlsMediaVideoApi(configs ...func(*common.ApiClient)) (*EncodingManifestsHlsMediaVideoApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingManifestsHlsMediaVideoApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingManifestsHlsMediaVideoApi) List(manifestId string, queryParams ...func(*query.VideoMediaInfoListQueryParams)) (*pagination.VideoMediaInfosListPagination, error) {
    queryParameters := &query.VideoMediaInfoListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.VideoMediaInfosListPagination
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/manifests/hls/{manifest_id}/media/video", &resp, reqParams)
    return resp, err
}
func (api *EncodingManifestsHlsMediaVideoApi) Delete(manifestId string, mediaId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["media_id"] = mediaId
	}
    err := api.apiClient.Delete("/encoding/manifests/hls/{manifest_id}/media/video/{media_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingManifestsHlsMediaVideoApi) Create(manifestId string, videoMediaInfo model.VideoMediaInfo) (*model.VideoMediaInfo, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
    }
    payload := model.VideoMediaInfo(videoMediaInfo)
    
    err := api.apiClient.Post("/encoding/manifests/hls/{manifest_id}/media/video", &payload, reqParams)
    return &payload, err
}
func (api *EncodingManifestsHlsMediaVideoApi) Get(manifestId string, mediaId string) (*model.VideoMediaInfo, error) {
    var resp *model.VideoMediaInfo
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["media_id"] = mediaId
	}
    err := api.apiClient.Get("/encoding/manifests/hls/{manifest_id}/media/video/{media_id}", &resp, reqParams)
    return resp, err
}
