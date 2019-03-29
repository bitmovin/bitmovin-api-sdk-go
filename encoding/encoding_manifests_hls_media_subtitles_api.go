package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingManifestsHlsMediaSubtitlesApi struct {
    apiClient *common.ApiClient
}

func NewEncodingManifestsHlsMediaSubtitlesApi(configs ...func(*common.ApiClient)) (*EncodingManifestsHlsMediaSubtitlesApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingManifestsHlsMediaSubtitlesApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingManifestsHlsMediaSubtitlesApi) Delete(manifestId string, mediaId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["media_id"] = mediaId
	}
    err := api.apiClient.Delete("/encoding/manifests/hls/{manifest_id}/media/subtitles/{media_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingManifestsHlsMediaSubtitlesApi) Get(manifestId string, mediaId string) (*model.SubtitlesMediaInfo, error) {
    var resp *model.SubtitlesMediaInfo
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["media_id"] = mediaId
	}
    err := api.apiClient.Get("/encoding/manifests/hls/{manifest_id}/media/subtitles/{media_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingManifestsHlsMediaSubtitlesApi) List(manifestId string, queryParams ...func(*query.SubtitlesMediaInfoListQueryParams)) (*pagination.SubtitlesMediaInfosListPagination, error) {
    queryParameters := &query.SubtitlesMediaInfoListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.SubtitlesMediaInfosListPagination
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/manifests/hls/{manifest_id}/media/subtitles", &resp, reqParams)
    return resp, err
}
func (api *EncodingManifestsHlsMediaSubtitlesApi) Create(manifestId string, subtitlesMediaInfo model.SubtitlesMediaInfo) (*model.SubtitlesMediaInfo, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
    }
    payload := model.SubtitlesMediaInfo(subtitlesMediaInfo)
    
    err := api.apiClient.Post("/encoding/manifests/hls/{manifest_id}/media/subtitles", &payload, reqParams)
    return &payload, err
}
