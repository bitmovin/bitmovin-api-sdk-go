package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingManifestsHlsMediaVideoAPI communicates with '/encoding/manifests/hls/{manifest_id}/media/video' endpoints
type EncodingManifestsHlsMediaVideoAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingManifestsHlsMediaVideoAPI constructor for EncodingManifestsHlsMediaVideoAPI that takes options as argument
func NewEncodingManifestsHlsMediaVideoAPI(options ...apiclient.APIClientOption) (*EncodingManifestsHlsMediaVideoAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingManifestsHlsMediaVideoAPIWithClient(apiClient), nil
}

// NewEncodingManifestsHlsMediaVideoAPIWithClient constructor for EncodingManifestsHlsMediaVideoAPI that takes an APIClient as argument
func NewEncodingManifestsHlsMediaVideoAPIWithClient(apiClient *apiclient.APIClient) *EncodingManifestsHlsMediaVideoAPI {
	a := &EncodingManifestsHlsMediaVideoAPI{apiClient: apiClient}
	return a
}

// Create Add Video Media
func (api *EncodingManifestsHlsMediaVideoAPI) Create(manifestId string, videoMediaInfo model.VideoMediaInfo) (*model.VideoMediaInfo, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
	}

	var responseModel model.VideoMediaInfo
	err := api.apiClient.Post("/encoding/manifests/hls/{manifest_id}/media/video", &videoMediaInfo, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Video Media
func (api *EncodingManifestsHlsMediaVideoAPI) Delete(manifestId string, mediaId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["media_id"] = mediaId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/manifests/hls/{manifest_id}/media/video/{media_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Video Media Details
func (api *EncodingManifestsHlsMediaVideoAPI) Get(manifestId string, mediaId string) (*model.VideoMediaInfo, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["media_id"] = mediaId
	}

	var responseModel model.VideoMediaInfo
	err := api.apiClient.Get("/encoding/manifests/hls/{manifest_id}/media/video/{media_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List all Video Media
func (api *EncodingManifestsHlsMediaVideoAPI) List(manifestId string, queryParams ...func(*EncodingManifestsHlsMediaVideoAPIListQueryParams)) (*pagination.VideoMediaInfosListPagination, error) {
	queryParameters := &EncodingManifestsHlsMediaVideoAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.VideoMediaInfosListPagination
	err := api.apiClient.Get("/encoding/manifests/hls/{manifest_id}/media/video", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingManifestsHlsMediaVideoAPIListQueryParams contains all query parameters for the List endpoint
type EncodingManifestsHlsMediaVideoAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingManifestsHlsMediaVideoAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
