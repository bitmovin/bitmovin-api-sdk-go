package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsMuxingsProgressiveTsId3FrameIdAPI communicates with '/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/id3/frame-id' endpoints
type EncodingEncodingsMuxingsProgressiveTsId3FrameIdAPI struct {
    apiClient *apiclient.APIClient

    // Customdata communicates with '/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/id3/frame-id/{id3_tag_id}/customData' endpoints
    Customdata *EncodingEncodingsMuxingsProgressiveTsId3FrameIdCustomdataAPI

}

// NewEncodingEncodingsMuxingsProgressiveTsId3FrameIdAPI constructor for EncodingEncodingsMuxingsProgressiveTsId3FrameIdAPI that takes options as argument
func NewEncodingEncodingsMuxingsProgressiveTsId3FrameIdAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsProgressiveTsId3FrameIdAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsMuxingsProgressiveTsId3FrameIdAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsProgressiveTsId3FrameIdAPIWithClient constructor for EncodingEncodingsMuxingsProgressiveTsId3FrameIdAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsProgressiveTsId3FrameIdAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsProgressiveTsId3FrameIdAPI {
    a := &EncodingEncodingsMuxingsProgressiveTsId3FrameIdAPI{apiClient: apiClient}
    a.Customdata = NewEncodingEncodingsMuxingsProgressiveTsId3FrameIdCustomdataAPIWithClient(apiClient)

    return a
}

// Create Add Frame ID ID3 Tag to a Progressive TS muxing
func (api *EncodingEncodingsMuxingsProgressiveTsId3FrameIdAPI) Create(encodingId string, muxingId string, frameIdId3Tag model.FrameIdId3Tag) (*model.FrameIdId3Tag, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel model.FrameIdId3Tag
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/id3/frame-id", &frameIdId3Tag, &responseModel, reqParams)
    return &responseModel, err
}
// Delete Frame ID ID3 Tag of a Progressive TS muxing
func (api *EncodingEncodingsMuxingsProgressiveTsId3FrameIdAPI) Delete(encodingId string, muxingId string, id3TagId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["id3_tag_id"] = id3TagId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/id3/frame-id/{id3_tag_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get Frame ID ID3 Tag Details of a Progressive TS muxing
func (api *EncodingEncodingsMuxingsProgressiveTsId3FrameIdAPI) Get(encodingId string, muxingId string, id3TagId string) (*model.FrameIdId3Tag, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["id3_tag_id"] = id3TagId
    }

    var responseModel model.FrameIdId3Tag
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/id3/frame-id/{id3_tag_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List Frame ID ID3 Tags of a Progressive TS muxing
func (api *EncodingEncodingsMuxingsProgressiveTsId3FrameIdAPI) List(encodingId string, muxingId string, queryParams ...func(*EncodingEncodingsMuxingsProgressiveTsId3FrameIdAPIListQueryParams)) (*pagination.FrameIdId3TagsListPagination, error) {
    queryParameters := &EncodingEncodingsMuxingsProgressiveTsId3FrameIdAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.QueryParams = queryParameters
    }

    var responseModel pagination.FrameIdId3TagsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/id3/frame-id", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingEncodingsMuxingsProgressiveTsId3FrameIdAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsMuxingsProgressiveTsId3FrameIdAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsMuxingsProgressiveTsId3FrameIdAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


