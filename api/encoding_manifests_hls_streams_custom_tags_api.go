package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingManifestsHlsStreamsCustomTagsAPI communicates with '/encoding/manifests/hls/{manifest_id}/streams/{stream_id}/custom-tags' endpoints
type EncodingManifestsHlsStreamsCustomTagsAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingManifestsHlsStreamsCustomTagsAPI constructor for EncodingManifestsHlsStreamsCustomTagsAPI that takes options as argument
func NewEncodingManifestsHlsStreamsCustomTagsAPI(options ...apiclient.APIClientOption) (*EncodingManifestsHlsStreamsCustomTagsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingManifestsHlsStreamsCustomTagsAPIWithClient(apiClient), nil
}

// NewEncodingManifestsHlsStreamsCustomTagsAPIWithClient constructor for EncodingManifestsHlsStreamsCustomTagsAPI that takes an APIClient as argument
func NewEncodingManifestsHlsStreamsCustomTagsAPIWithClient(apiClient *apiclient.APIClient) *EncodingManifestsHlsStreamsCustomTagsAPI {
	a := &EncodingManifestsHlsStreamsCustomTagsAPI{apiClient: apiClient}
	return a
}

// Create Add Custom Tag to Variant Stream
func (api *EncodingManifestsHlsStreamsCustomTagsAPI) Create(manifestId string, streamId string, customTag model.CustomTag) (*model.CustomTag, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["stream_id"] = streamId
	}

	var responseModel model.CustomTag
	err := api.apiClient.Post("/encoding/manifests/hls/{manifest_id}/streams/{stream_id}/custom-tags", &customTag, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Custom Tag
func (api *EncodingManifestsHlsStreamsCustomTagsAPI) Delete(manifestId string, streamId string, customTagId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["stream_id"] = streamId
		params.PathParams["custom_tag_id"] = customTagId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/manifests/hls/{manifest_id}/streams/{stream_id}/custom-tags/{custom_tag_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Custom Tag Details
func (api *EncodingManifestsHlsStreamsCustomTagsAPI) Get(manifestId string, streamId string, customTagId string) (*model.CustomTag, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["stream_id"] = streamId
		params.PathParams["custom_tag_id"] = customTagId
	}

	var responseModel model.CustomTag
	err := api.apiClient.Get("/encoding/manifests/hls/{manifest_id}/streams/{stream_id}/custom-tags/{custom_tag_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List all Custom Tags of a Variant Stream
func (api *EncodingManifestsHlsStreamsCustomTagsAPI) List(manifestId string, streamId string, queryParams ...func(*EncodingManifestsHlsStreamsCustomTagsAPIListQueryParams)) (*pagination.CustomTagsListPagination, error) {
	queryParameters := &EncodingManifestsHlsStreamsCustomTagsAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["manifest_id"] = manifestId
		params.PathParams["stream_id"] = streamId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.CustomTagsListPagination
	err := api.apiClient.Get("/encoding/manifests/hls/{manifest_id}/streams/{stream_id}/custom-tags", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingManifestsHlsStreamsCustomTagsAPIListQueryParams contains all query parameters for the List endpoint
type EncodingManifestsHlsStreamsCustomTagsAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingManifestsHlsStreamsCustomTagsAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
