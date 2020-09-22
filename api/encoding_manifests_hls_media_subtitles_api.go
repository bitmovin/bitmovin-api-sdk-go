package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingManifestsHlsMediaSubtitlesAPI communicates with '/encoding/manifests/hls/{manifest_id}/media/subtitles' endpoints
type EncodingManifestsHlsMediaSubtitlesAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingManifestsHlsMediaSubtitlesAPI constructor for EncodingManifestsHlsMediaSubtitlesAPI that takes options as argument
func NewEncodingManifestsHlsMediaSubtitlesAPI(options ...apiclient.APIClientOption) (*EncodingManifestsHlsMediaSubtitlesAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingManifestsHlsMediaSubtitlesAPIWithClient(apiClient), nil
}

// NewEncodingManifestsHlsMediaSubtitlesAPIWithClient constructor for EncodingManifestsHlsMediaSubtitlesAPI that takes an APIClient as argument
func NewEncodingManifestsHlsMediaSubtitlesAPIWithClient(apiClient *apiclient.APIClient) *EncodingManifestsHlsMediaSubtitlesAPI {
	a := &EncodingManifestsHlsMediaSubtitlesAPI{apiClient: apiClient}
	return a
}

// Create Add Subtitles Media
func (api *EncodingManifestsHlsMediaSubtitlesAPI) Create(manifestId string, subtitlesMediaInfo model.SubtitlesMediaInfo) (*model.SubtitlesMediaInfo, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
	}

	var responseModel model.SubtitlesMediaInfo
	err := api.apiClient.Post("/encoding/manifests/hls/{manifest_id}/media/subtitles", &subtitlesMediaInfo, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Subtitles Media
func (api *EncodingManifestsHlsMediaSubtitlesAPI) Delete(manifestId string, mediaId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["media_id"] = mediaId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/manifests/hls/{manifest_id}/media/subtitles/{media_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Subtitles Media Details
func (api *EncodingManifestsHlsMediaSubtitlesAPI) Get(manifestId string, mediaId string) (*model.SubtitlesMediaInfo, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["media_id"] = mediaId
	}

	var responseModel model.SubtitlesMediaInfo
	err := api.apiClient.Get("/encoding/manifests/hls/{manifest_id}/media/subtitles/{media_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List all Subtitles Media
func (api *EncodingManifestsHlsMediaSubtitlesAPI) List(manifestId string, queryParams ...func(*EncodingManifestsHlsMediaSubtitlesAPIListQueryParams)) (*pagination.SubtitlesMediaInfosListPagination, error) {
	queryParameters := &EncodingManifestsHlsMediaSubtitlesAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.SubtitlesMediaInfosListPagination
	err := api.apiClient.Get("/encoding/manifests/hls/{manifest_id}/media/subtitles", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingManifestsHlsMediaSubtitlesAPIListQueryParams contains all query parameters for the List endpoint
type EncodingManifestsHlsMediaSubtitlesAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingManifestsHlsMediaSubtitlesAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
