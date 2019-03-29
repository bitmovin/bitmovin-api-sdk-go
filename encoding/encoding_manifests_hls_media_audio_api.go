package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingManifestsHlsMediaAudioApi struct {
    apiClient *common.ApiClient
}

func NewEncodingManifestsHlsMediaAudioApi(configs ...func(*common.ApiClient)) (*EncodingManifestsHlsMediaAudioApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingManifestsHlsMediaAudioApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingManifestsHlsMediaAudioApi) Get(manifestId string, mediaId string) (*model.AudioMediaInfo, error) {
    var resp *model.AudioMediaInfo
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["media_id"] = mediaId
	}
    err := api.apiClient.Get("/encoding/manifests/hls/{manifest_id}/media/audio/{media_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingManifestsHlsMediaAudioApi) Create(manifestId string, audioMediaInfo model.AudioMediaInfo) (*model.AudioMediaInfo, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
    }
    payload := model.AudioMediaInfo(audioMediaInfo)
    
    err := api.apiClient.Post("/encoding/manifests/hls/{manifest_id}/media/audio", &payload, reqParams)
    return &payload, err
}
func (api *EncodingManifestsHlsMediaAudioApi) List(manifestId string, queryParams ...func(*query.AudioMediaInfoListQueryParams)) (*pagination.AudioMediaInfosListPagination, error) {
    queryParameters := &query.AudioMediaInfoListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.AudioMediaInfosListPagination
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/manifests/hls/{manifest_id}/media/audio", &resp, reqParams)
    return resp, err
}
func (api *EncodingManifestsHlsMediaAudioApi) Delete(manifestId string, mediaId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["media_id"] = mediaId
	}
    err := api.apiClient.Delete("/encoding/manifests/hls/{manifest_id}/media/audio/{media_id}", &resp, reqParams)
    return resp, err
}
