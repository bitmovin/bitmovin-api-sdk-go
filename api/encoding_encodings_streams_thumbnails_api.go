package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsStreamsThumbnailsAPI communicates with '/encoding/encodings/{encoding_id}/streams/{stream_id}/thumbnails' endpoints
type EncodingEncodingsStreamsThumbnailsAPI struct {
    apiClient *apiclient.APIClient

    // Customdata communicates with '/encoding/encodings/{encoding_id}/streams/{stream_id}/thumbnails/{thumbnail_id}/customData' endpoints
    Customdata *EncodingEncodingsStreamsThumbnailsCustomdataAPI

}

// NewEncodingEncodingsStreamsThumbnailsAPI constructor for EncodingEncodingsStreamsThumbnailsAPI that takes options as argument
func NewEncodingEncodingsStreamsThumbnailsAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsStreamsThumbnailsAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsStreamsThumbnailsAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsStreamsThumbnailsAPIWithClient constructor for EncodingEncodingsStreamsThumbnailsAPI that takes an APIClient as argument
func NewEncodingEncodingsStreamsThumbnailsAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsStreamsThumbnailsAPI {
    a := &EncodingEncodingsStreamsThumbnailsAPI{apiClient: apiClient}
    a.Customdata = NewEncodingEncodingsStreamsThumbnailsCustomdataAPIWithClient(apiClient)

    return a
}

// Create Add Thumbnail
func (api *EncodingEncodingsStreamsThumbnailsAPI) Create(encodingId string, streamId string, thumbnail model.Thumbnail) (*model.Thumbnail, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
    }

    var responseModel model.Thumbnail
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/streams/{stream_id}/thumbnails", &thumbnail, &responseModel, reqParams)
    return &responseModel, err
}
// Delete Thumbnail
func (api *EncodingEncodingsStreamsThumbnailsAPI) Delete(encodingId string, streamId string, thumbnailId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
        params.PathParams["thumbnail_id"] = thumbnailId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/streams/{stream_id}/thumbnails/{thumbnail_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get Thumbnail Details
func (api *EncodingEncodingsStreamsThumbnailsAPI) Get(encodingId string, streamId string, thumbnailId string) (*model.Thumbnail, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
        params.PathParams["thumbnail_id"] = thumbnailId
    }

    var responseModel model.Thumbnail
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/streams/{stream_id}/thumbnails/{thumbnail_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List Thumbnails
func (api *EncodingEncodingsStreamsThumbnailsAPI) List(encodingId string, streamId string, queryParams ...func(*EncodingEncodingsStreamsThumbnailsAPIListQueryParams)) (*pagination.ThumbnailsListPagination, error) {
    queryParameters := &EncodingEncodingsStreamsThumbnailsAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
        params.QueryParams = queryParameters
    }

    var responseModel pagination.ThumbnailsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/streams/{stream_id}/thumbnails", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingEncodingsStreamsThumbnailsAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsStreamsThumbnailsAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsStreamsThumbnailsAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


