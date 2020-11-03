package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsStreamsSpritesAPI communicates with '/encoding/encodings/{encoding_id}/streams/{stream_id}/sprites' endpoints
type EncodingEncodingsStreamsSpritesAPI struct {
    apiClient *apiclient.APIClient

    // Customdata communicates with '/encoding/encodings/{encoding_id}/streams/{stream_id}/sprites/{sprite_id}/customData' endpoints
    Customdata *EncodingEncodingsStreamsSpritesCustomdataAPI

}

// NewEncodingEncodingsStreamsSpritesAPI constructor for EncodingEncodingsStreamsSpritesAPI that takes options as argument
func NewEncodingEncodingsStreamsSpritesAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsStreamsSpritesAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsStreamsSpritesAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsStreamsSpritesAPIWithClient constructor for EncodingEncodingsStreamsSpritesAPI that takes an APIClient as argument
func NewEncodingEncodingsStreamsSpritesAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsStreamsSpritesAPI {
    a := &EncodingEncodingsStreamsSpritesAPI{apiClient: apiClient}
    a.Customdata = NewEncodingEncodingsStreamsSpritesCustomdataAPIWithClient(apiClient)

    return a
}

// Create Add Sprite
func (api *EncodingEncodingsStreamsSpritesAPI) Create(encodingId string, streamId string, sprite model.Sprite) (*model.Sprite, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
    }

    var responseModel model.Sprite
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/streams/{stream_id}/sprites", &sprite, &responseModel, reqParams)
    return &responseModel, err
}
// Delete Sprite
func (api *EncodingEncodingsStreamsSpritesAPI) Delete(encodingId string, streamId string, spriteId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
        params.PathParams["sprite_id"] = spriteId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/streams/{stream_id}/sprites/{sprite_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get Sprite Details
func (api *EncodingEncodingsStreamsSpritesAPI) Get(encodingId string, streamId string, spriteId string) (*model.Sprite, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
        params.PathParams["sprite_id"] = spriteId
    }

    var responseModel model.Sprite
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/streams/{stream_id}/sprites/{sprite_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List Sprites
func (api *EncodingEncodingsStreamsSpritesAPI) List(encodingId string, streamId string, queryParams ...func(*EncodingEncodingsStreamsSpritesAPIListQueryParams)) (*pagination.SpritesListPagination, error) {
    queryParameters := &EncodingEncodingsStreamsSpritesAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
        params.QueryParams = queryParameters
    }

    var responseModel pagination.SpritesListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/streams/{stream_id}/sprites", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingEncodingsStreamsSpritesAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsStreamsSpritesAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsStreamsSpritesAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


