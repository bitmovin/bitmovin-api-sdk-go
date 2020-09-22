package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingManifestsHlsStreamsIframeAPI communicates with '/encoding/manifests/hls/{manifest_id}/streams/{stream_id}/iframe' endpoints
type EncodingManifestsHlsStreamsIframeAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingManifestsHlsStreamsIframeAPI constructor for EncodingManifestsHlsStreamsIframeAPI that takes options as argument
func NewEncodingManifestsHlsStreamsIframeAPI(options ...apiclient.APIClientOption) (*EncodingManifestsHlsStreamsIframeAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingManifestsHlsStreamsIframeAPIWithClient(apiClient), nil
}

// NewEncodingManifestsHlsStreamsIframeAPIWithClient constructor for EncodingManifestsHlsStreamsIframeAPI that takes an APIClient as argument
func NewEncodingManifestsHlsStreamsIframeAPIWithClient(apiClient *apiclient.APIClient) *EncodingManifestsHlsStreamsIframeAPI {
	a := &EncodingManifestsHlsStreamsIframeAPI{apiClient: apiClient}
	return a
}

// Create Add I-frame playlist to variant stream
func (api *EncodingManifestsHlsStreamsIframeAPI) Create(manifestId string, streamId string, iFramePlaylist model.IFramePlaylist) (*model.IFramePlaylist, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["stream_id"] = streamId
	}

	var responseModel model.IFramePlaylist
	err := api.apiClient.Post("/encoding/manifests/hls/{manifest_id}/streams/{stream_id}/iframe", &iFramePlaylist, &responseModel, reqParams)
	return &responseModel, err
}

// Delete I-frame playlist
func (api *EncodingManifestsHlsStreamsIframeAPI) Delete(manifestId string, streamId string, iframeId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["stream_id"] = streamId
		params.PathParams["iframe_id"] = iframeId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/manifests/hls/{manifest_id}/streams/{stream_id}/iframe/{iframe_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get I-frame playlist Details
func (api *EncodingManifestsHlsStreamsIframeAPI) Get(manifestId string, streamId string, iframeId string) (*model.IFramePlaylist, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["stream_id"] = streamId
		params.PathParams["iframe_id"] = iframeId
	}

	var responseModel model.IFramePlaylist
	err := api.apiClient.Get("/encoding/manifests/hls/{manifest_id}/streams/{stream_id}/iframe/{iframe_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List all I-frame playlists of a variant stream
func (api *EncodingManifestsHlsStreamsIframeAPI) List(manifestId string, streamId string, queryParams ...func(*EncodingManifestsHlsStreamsIframeAPIListQueryParams)) (*pagination.IFramePlaylistsListPagination, error) {
	queryParameters := &EncodingManifestsHlsStreamsIframeAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["stream_id"] = streamId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.IFramePlaylistsListPagination
	err := api.apiClient.Get("/encoding/manifests/hls/{manifest_id}/streams/{stream_id}/iframe", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingManifestsHlsStreamsIframeAPIListQueryParams contains all query parameters for the List endpoint
type EncodingManifestsHlsStreamsIframeAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingManifestsHlsStreamsIframeAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
