package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingManifestsHlsStreamsAPI communicates with '/encoding/manifests/hls/{manifest_id}/streams' endpoints
type EncodingManifestsHlsStreamsAPI struct {
    apiClient *apiclient.APIClient

    // CustomTags communicates with '/encoding/manifests/hls/{manifest_id}/streams/{stream_id}/custom-tags' endpoints
    CustomTags *EncodingManifestsHlsStreamsCustomTagsAPI
    // Iframe communicates with '/encoding/manifests/hls/{manifest_id}/streams/{stream_id}/iframe' endpoints
    Iframe *EncodingManifestsHlsStreamsIframeAPI

}

// NewEncodingManifestsHlsStreamsAPI constructor for EncodingManifestsHlsStreamsAPI that takes options as argument
func NewEncodingManifestsHlsStreamsAPI(options ...apiclient.APIClientOption) (*EncodingManifestsHlsStreamsAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingManifestsHlsStreamsAPIWithClient(apiClient), nil
}

// NewEncodingManifestsHlsStreamsAPIWithClient constructor for EncodingManifestsHlsStreamsAPI that takes an APIClient as argument
func NewEncodingManifestsHlsStreamsAPIWithClient(apiClient *apiclient.APIClient) *EncodingManifestsHlsStreamsAPI {
    a := &EncodingManifestsHlsStreamsAPI{apiClient: apiClient}
    a.CustomTags = NewEncodingManifestsHlsStreamsCustomTagsAPIWithClient(apiClient)
    a.Iframe = NewEncodingManifestsHlsStreamsIframeAPIWithClient(apiClient)

    return a
}

// Create Add Variant Stream
func (api *EncodingManifestsHlsStreamsAPI) Create(manifestId string, streamInfo model.StreamInfo) (*model.StreamInfo, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
    }

    var responseModel model.StreamInfo
    err := api.apiClient.Post("/encoding/manifests/hls/{manifest_id}/streams", &streamInfo, &responseModel, reqParams)
    return &responseModel, err
}
// Delete Variant Stream
func (api *EncodingManifestsHlsStreamsAPI) Delete(manifestId string, streamId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["stream_id"] = streamId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/manifests/hls/{manifest_id}/streams/{stream_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get Variant Stream Details
func (api *EncodingManifestsHlsStreamsAPI) Get(manifestId string, streamId string) (*model.StreamInfo, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["stream_id"] = streamId
    }

    var responseModel model.StreamInfo
    err := api.apiClient.Get("/encoding/manifests/hls/{manifest_id}/streams/{stream_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List all Variant Streams
func (api *EncodingManifestsHlsStreamsAPI) List(manifestId string, queryParams ...func(*EncodingManifestsHlsStreamsAPIListQueryParams)) (*pagination.StreamInfosListPagination, error) {
    queryParameters := &EncodingManifestsHlsStreamsAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.QueryParams = queryParameters
    }

    var responseModel pagination.StreamInfosListPagination
    err := api.apiClient.Get("/encoding/manifests/hls/{manifest_id}/streams", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingManifestsHlsStreamsAPIListQueryParams contains all query parameters for the List endpoint
type EncodingManifestsHlsStreamsAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingManifestsHlsStreamsAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


