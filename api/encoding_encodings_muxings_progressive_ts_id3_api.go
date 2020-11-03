package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsMuxingsProgressiveTsId3API communicates with '/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/id3' endpoints
type EncodingEncodingsMuxingsProgressiveTsId3API struct {
    apiClient *apiclient.APIClient

    // Raw communicates with '/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/id3/raw' endpoints
    Raw *EncodingEncodingsMuxingsProgressiveTsId3RawAPI
    // FrameId communicates with '/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/id3/frame-id' endpoints
    FrameId *EncodingEncodingsMuxingsProgressiveTsId3FrameIdAPI
    // PlainText communicates with '/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/id3/plain-text' endpoints
    PlainText *EncodingEncodingsMuxingsProgressiveTsId3PlainTextAPI

}

// NewEncodingEncodingsMuxingsProgressiveTsId3API constructor for EncodingEncodingsMuxingsProgressiveTsId3API that takes options as argument
func NewEncodingEncodingsMuxingsProgressiveTsId3API(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsProgressiveTsId3API, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsMuxingsProgressiveTsId3APIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsProgressiveTsId3APIWithClient constructor for EncodingEncodingsMuxingsProgressiveTsId3API that takes an APIClient as argument
func NewEncodingEncodingsMuxingsProgressiveTsId3APIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsProgressiveTsId3API {
    a := &EncodingEncodingsMuxingsProgressiveTsId3API{apiClient: apiClient}
    a.Raw = NewEncodingEncodingsMuxingsProgressiveTsId3RawAPIWithClient(apiClient)
    a.FrameId = NewEncodingEncodingsMuxingsProgressiveTsId3FrameIdAPIWithClient(apiClient)
    a.PlainText = NewEncodingEncodingsMuxingsProgressiveTsId3PlainTextAPIWithClient(apiClient)

    return a
}

// List all ID3 Tags of a Progressive TS muxing
func (api *EncodingEncodingsMuxingsProgressiveTsId3API) List(encodingId string, muxingId string, queryParams ...func(*EncodingEncodingsMuxingsProgressiveTsId3APIListQueryParams)) (*pagination.Id3TagsListPagination, error) {
    queryParameters := &EncodingEncodingsMuxingsProgressiveTsId3APIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.QueryParams = queryParameters
    }

    var responseModel pagination.Id3TagsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/id3", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingEncodingsMuxingsProgressiveTsId3APIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsMuxingsProgressiveTsId3APIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsMuxingsProgressiveTsId3APIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


