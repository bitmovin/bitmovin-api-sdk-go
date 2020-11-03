package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsMuxingsProgressiveTsId3PlainTextAPI communicates with '/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/id3/plain-text' endpoints
type EncodingEncodingsMuxingsProgressiveTsId3PlainTextAPI struct {
    apiClient *apiclient.APIClient

    // Customdata communicates with '/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/id3/plain-text/{id3_tag_id}/customData' endpoints
    Customdata *EncodingEncodingsMuxingsProgressiveTsId3PlainTextCustomdataAPI

}

// NewEncodingEncodingsMuxingsProgressiveTsId3PlainTextAPI constructor for EncodingEncodingsMuxingsProgressiveTsId3PlainTextAPI that takes options as argument
func NewEncodingEncodingsMuxingsProgressiveTsId3PlainTextAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsProgressiveTsId3PlainTextAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsMuxingsProgressiveTsId3PlainTextAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsProgressiveTsId3PlainTextAPIWithClient constructor for EncodingEncodingsMuxingsProgressiveTsId3PlainTextAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsProgressiveTsId3PlainTextAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsProgressiveTsId3PlainTextAPI {
    a := &EncodingEncodingsMuxingsProgressiveTsId3PlainTextAPI{apiClient: apiClient}
    a.Customdata = NewEncodingEncodingsMuxingsProgressiveTsId3PlainTextCustomdataAPIWithClient(apiClient)

    return a
}

// Create Add Plain Text ID3 Tag to a Progressive TS muxing
func (api *EncodingEncodingsMuxingsProgressiveTsId3PlainTextAPI) Create(encodingId string, muxingId string, plaintextId3Tag model.PlaintextId3Tag) (*model.PlaintextId3Tag, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel model.PlaintextId3Tag
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/id3/plain-text", &plaintextId3Tag, &responseModel, reqParams)
    return &responseModel, err
}
// Delete Plain Text ID3 Tag of a Progressive TS muxing
func (api *EncodingEncodingsMuxingsProgressiveTsId3PlainTextAPI) Delete(encodingId string, muxingId string, id3TagId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["id3_tag_id"] = id3TagId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/id3/plain-text/{id3_tag_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get Plain Text ID3 Tag Details of a Progressive TS muxing
func (api *EncodingEncodingsMuxingsProgressiveTsId3PlainTextAPI) Get(encodingId string, muxingId string, id3TagId string) (*model.PlaintextId3Tag, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["id3_tag_id"] = id3TagId
    }

    var responseModel model.PlaintextId3Tag
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/id3/plain-text/{id3_tag_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List Plain Text ID3 Tags of a Progressive TS muxing
func (api *EncodingEncodingsMuxingsProgressiveTsId3PlainTextAPI) List(encodingId string, muxingId string, queryParams ...func(*EncodingEncodingsMuxingsProgressiveTsId3PlainTextAPIListQueryParams)) (*pagination.PlaintextId3TagsListPagination, error) {
    queryParameters := &EncodingEncodingsMuxingsProgressiveTsId3PlainTextAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.QueryParams = queryParameters
    }

    var responseModel pagination.PlaintextId3TagsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/id3/plain-text", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingEncodingsMuxingsProgressiveTsId3PlainTextAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsMuxingsProgressiveTsId3PlainTextAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsMuxingsProgressiveTsId3PlainTextAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


