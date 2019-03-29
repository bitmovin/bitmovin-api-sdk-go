package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingManifestsHlsMediaVttApi struct {
    apiClient *common.ApiClient
}

func NewEncodingManifestsHlsMediaVttApi(configs ...func(*common.ApiClient)) (*EncodingManifestsHlsMediaVttApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingManifestsHlsMediaVttApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingManifestsHlsMediaVttApi) List(manifestId string, queryParams ...func(*query.VttMediaInfoListQueryParams)) (*pagination.VttMediaInfosListPagination, error) {
    queryParameters := &query.VttMediaInfoListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.VttMediaInfosListPagination
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/manifests/hls/{manifest_id}/media/vtt", &resp, reqParams)
    return resp, err
}
func (api *EncodingManifestsHlsMediaVttApi) Create(manifestId string, vttMediaInfo model.VttMediaInfo) (*model.VttMediaInfo, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
    }
    payload := model.VttMediaInfo(vttMediaInfo)
    
    err := api.apiClient.Post("/encoding/manifests/hls/{manifest_id}/media/vtt", &payload, reqParams)
    return &payload, err
}
func (api *EncodingManifestsHlsMediaVttApi) Delete(manifestId string, mediaId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["media_id"] = mediaId
	}
    err := api.apiClient.Delete("/encoding/manifests/hls/{manifest_id}/media/vtt/{media_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingManifestsHlsMediaVttApi) Get(manifestId string, mediaId string) (*model.VttMediaInfo, error) {
    var resp *model.VttMediaInfo
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["media_id"] = mediaId
	}
    err := api.apiClient.Get("/encoding/manifests/hls/{manifest_id}/media/vtt/{media_id}", &resp, reqParams)
    return resp, err
}
