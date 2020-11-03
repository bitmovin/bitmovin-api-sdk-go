package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingManifestsHlsMediaVttAPI communicates with '/encoding/manifests/hls/{manifest_id}/media/vtt' endpoints
type EncodingManifestsHlsMediaVttAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingManifestsHlsMediaVttAPI constructor for EncodingManifestsHlsMediaVttAPI that takes options as argument
func NewEncodingManifestsHlsMediaVttAPI(options ...apiclient.APIClientOption) (*EncodingManifestsHlsMediaVttAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingManifestsHlsMediaVttAPIWithClient(apiClient), nil
}

// NewEncodingManifestsHlsMediaVttAPIWithClient constructor for EncodingManifestsHlsMediaVttAPI that takes an APIClient as argument
func NewEncodingManifestsHlsMediaVttAPIWithClient(apiClient *apiclient.APIClient) *EncodingManifestsHlsMediaVttAPI {
    a := &EncodingManifestsHlsMediaVttAPI{apiClient: apiClient}
    return a
}

// Create Add VTT Media
func (api *EncodingManifestsHlsMediaVttAPI) Create(manifestId string, vttMediaInfo model.VttMediaInfo) (*model.VttMediaInfo, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
    }

    var responseModel model.VttMediaInfo
    err := api.apiClient.Post("/encoding/manifests/hls/{manifest_id}/media/vtt", &vttMediaInfo, &responseModel, reqParams)
    return &responseModel, err
}
// Delete VTT Media
func (api *EncodingManifestsHlsMediaVttAPI) Delete(manifestId string, mediaId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["media_id"] = mediaId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/manifests/hls/{manifest_id}/media/vtt/{media_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get VTT Media Details
func (api *EncodingManifestsHlsMediaVttAPI) Get(manifestId string, mediaId string) (*model.VttMediaInfo, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.PathParams["media_id"] = mediaId
    }

    var responseModel model.VttMediaInfo
    err := api.apiClient.Get("/encoding/manifests/hls/{manifest_id}/media/vtt/{media_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List all VTT Media
func (api *EncodingManifestsHlsMediaVttAPI) List(manifestId string, queryParams ...func(*EncodingManifestsHlsMediaVttAPIListQueryParams)) (*pagination.VttMediaInfosListPagination, error) {
    queryParameters := &EncodingManifestsHlsMediaVttAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
        params.QueryParams = queryParameters
    }

    var responseModel pagination.VttMediaInfosListPagination
    err := api.apiClient.Get("/encoding/manifests/hls/{manifest_id}/media/vtt", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingManifestsHlsMediaVttAPIListQueryParams contains all query parameters for the List endpoint
type EncodingManifestsHlsMediaVttAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingManifestsHlsMediaVttAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


