package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingManifestsHlsMediaCustomTagsAPI communicates with '/encoding/manifests/hls/{manifest_id}/media/{media_id}/custom-tags' endpoints
type EncodingManifestsHlsMediaCustomTagsAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingManifestsHlsMediaCustomTagsAPI constructor for EncodingManifestsHlsMediaCustomTagsAPI that takes options as argument
func NewEncodingManifestsHlsMediaCustomTagsAPI(options ...apiclient.APIClientOption) (*EncodingManifestsHlsMediaCustomTagsAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingManifestsHlsMediaCustomTagsAPIWithClient(apiClient), nil
}

// NewEncodingManifestsHlsMediaCustomTagsAPIWithClient constructor for EncodingManifestsHlsMediaCustomTagsAPI that takes an APIClient as argument
func NewEncodingManifestsHlsMediaCustomTagsAPIWithClient(apiClient *apiclient.APIClient) *EncodingManifestsHlsMediaCustomTagsAPI {
    a := &EncodingManifestsHlsMediaCustomTagsAPI{apiClient: apiClient}
    return a
}

// Create Add Custom Tag to Audio Media
func (api *EncodingManifestsHlsMediaCustomTagsAPI) Create(manifestId string, mediaId string, customTag model.CustomTag) (*model.CustomTag, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["media_id"] = mediaId
    }

    var responseModel model.CustomTag
    err := api.apiClient.Post("/encoding/manifests/hls/{manifest_id}/media/{media_id}/custom-tags", &customTag, &responseModel, reqParams)
    return &responseModel, err
}
// Delete Custom Tag
func (api *EncodingManifestsHlsMediaCustomTagsAPI) Delete(manifestId string, mediaId string, customTagId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["media_id"] = mediaId
        params.PathParams["custom_tag_id"] = customTagId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/manifests/hls/{manifest_id}/media/{media_id}/custom-tags/{custom_tag_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get Custom Tag Details
func (api *EncodingManifestsHlsMediaCustomTagsAPI) Get(manifestId string, mediaId string, customTagId string) (*model.CustomTag, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["media_id"] = mediaId
        params.PathParams["custom_tag_id"] = customTagId
    }

    var responseModel model.CustomTag
    err := api.apiClient.Get("/encoding/manifests/hls/{manifest_id}/media/{media_id}/custom-tags/{custom_tag_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List all Custom Tags of a Audio media
func (api *EncodingManifestsHlsMediaCustomTagsAPI) List(manifestId string, mediaId string, queryParams ...func(*EncodingManifestsHlsMediaCustomTagsAPIListQueryParams)) (*pagination.CustomTagsListPagination, error) {
    queryParameters := &EncodingManifestsHlsMediaCustomTagsAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["media_id"] = mediaId
        params.QueryParams = queryParameters
    }

    var responseModel pagination.CustomTagsListPagination
    err := api.apiClient.Get("/encoding/manifests/hls/{manifest_id}/media/{media_id}/custom-tags", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingManifestsHlsMediaCustomTagsAPIListQueryParams contains all query parameters for the List endpoint
type EncodingManifestsHlsMediaCustomTagsAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingManifestsHlsMediaCustomTagsAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


