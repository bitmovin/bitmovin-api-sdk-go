package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingManifestsHlsMediaAudioAPI communicates with '/encoding/manifests/hls/{manifest_id}/media/audio' endpoints
type EncodingManifestsHlsMediaAudioAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingManifestsHlsMediaAudioAPI constructor for EncodingManifestsHlsMediaAudioAPI that takes options as argument
func NewEncodingManifestsHlsMediaAudioAPI(options ...apiclient.APIClientOption) (*EncodingManifestsHlsMediaAudioAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingManifestsHlsMediaAudioAPIWithClient(apiClient), nil
}

// NewEncodingManifestsHlsMediaAudioAPIWithClient constructor for EncodingManifestsHlsMediaAudioAPI that takes an APIClient as argument
func NewEncodingManifestsHlsMediaAudioAPIWithClient(apiClient *apiclient.APIClient) *EncodingManifestsHlsMediaAudioAPI {
	a := &EncodingManifestsHlsMediaAudioAPI{apiClient: apiClient}
	return a
}

// Create Add Audio Media
func (api *EncodingManifestsHlsMediaAudioAPI) Create(manifestId string, audioMediaInfo model.AudioMediaInfo) (*model.AudioMediaInfo, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
	}

	var responseModel model.AudioMediaInfo
	err := api.apiClient.Post("/encoding/manifests/hls/{manifest_id}/media/audio", &audioMediaInfo, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Audio Media
func (api *EncodingManifestsHlsMediaAudioAPI) Delete(manifestId string, mediaId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["media_id"] = mediaId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/manifests/hls/{manifest_id}/media/audio/{media_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Audio Media Details
func (api *EncodingManifestsHlsMediaAudioAPI) Get(manifestId string, mediaId string) (*model.AudioMediaInfo, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["media_id"] = mediaId
	}

	var responseModel model.AudioMediaInfo
	err := api.apiClient.Get("/encoding/manifests/hls/{manifest_id}/media/audio/{media_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List all Audio Media
func (api *EncodingManifestsHlsMediaAudioAPI) List(manifestId string, queryParams ...func(*EncodingManifestsHlsMediaAudioAPIListQueryParams)) (*pagination.AudioMediaInfosListPagination, error) {
	queryParameters := &EncodingManifestsHlsMediaAudioAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.AudioMediaInfosListPagination
	err := api.apiClient.Get("/encoding/manifests/hls/{manifest_id}/media/audio", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingManifestsHlsMediaAudioAPIListQueryParams contains all query parameters for the List endpoint
type EncodingManifestsHlsMediaAudioAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingManifestsHlsMediaAudioAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
