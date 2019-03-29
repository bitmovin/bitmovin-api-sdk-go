package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingManifestsHlsMediaClosedCaptionsApi struct {
    apiClient *common.ApiClient
}

func NewEncodingManifestsHlsMediaClosedCaptionsApi(configs ...func(*common.ApiClient)) (*EncodingManifestsHlsMediaClosedCaptionsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingManifestsHlsMediaClosedCaptionsApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingManifestsHlsMediaClosedCaptionsApi) Create(manifestId string, closedCaptionsMediaInfo model.ClosedCaptionsMediaInfo) (*model.ClosedCaptionsMediaInfo, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
    }
    payload := model.ClosedCaptionsMediaInfo(closedCaptionsMediaInfo)
    
    err := api.apiClient.Post("/encoding/manifests/hls/{manifest_id}/media/closed-captions", &payload, reqParams)
    return &payload, err
}
func (api *EncodingManifestsHlsMediaClosedCaptionsApi) Get(manifestId string, mediaId string) (*model.ClosedCaptionsMediaInfo, error) {
    var resp *model.ClosedCaptionsMediaInfo
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["media_id"] = mediaId
	}
    err := api.apiClient.Get("/encoding/manifests/hls/{manifest_id}/media/closed-captions/{media_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingManifestsHlsMediaClosedCaptionsApi) Delete(manifestId string, mediaId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["media_id"] = mediaId
	}
    err := api.apiClient.Delete("/encoding/manifests/hls/{manifest_id}/media/closed-captions/{media_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingManifestsHlsMediaClosedCaptionsApi) List(manifestId string, queryParams ...func(*query.ClosedCaptionsMediaInfoListQueryParams)) (*pagination.ClosedCaptionsMediaInfosListPagination, error) {
    queryParameters := &query.ClosedCaptionsMediaInfoListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.ClosedCaptionsMediaInfosListPagination
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/manifests/hls/{manifest_id}/media/closed-captions", &resp, reqParams)
    return resp, err
}
