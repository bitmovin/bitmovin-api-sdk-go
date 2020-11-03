package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsStreamsHdrDolbyVisionAPI communicates with '/encoding/encodings/{encoding_id}/streams/{stream_id}/hdr/dolby-vision' endpoints
type EncodingEncodingsStreamsHdrDolbyVisionAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingEncodingsStreamsHdrDolbyVisionAPI constructor for EncodingEncodingsStreamsHdrDolbyVisionAPI that takes options as argument
func NewEncodingEncodingsStreamsHdrDolbyVisionAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsStreamsHdrDolbyVisionAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsStreamsHdrDolbyVisionAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsStreamsHdrDolbyVisionAPIWithClient constructor for EncodingEncodingsStreamsHdrDolbyVisionAPI that takes an APIClient as argument
func NewEncodingEncodingsStreamsHdrDolbyVisionAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsStreamsHdrDolbyVisionAPI {
    a := &EncodingEncodingsStreamsHdrDolbyVisionAPI{apiClient: apiClient}
    return a
}

// Create Add Dolby Vision Metadata
func (api *EncodingEncodingsStreamsHdrDolbyVisionAPI) Create(encodingId string, streamId string, dolbyVisionMetadata model.DolbyVisionMetadata) (*model.DolbyVisionMetadata, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
    }

    var responseModel model.DolbyVisionMetadata
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/streams/{stream_id}/hdr/dolby-vision", &dolbyVisionMetadata, &responseModel, reqParams)
    return &responseModel, err
}
// Delete Dolby Vision Metadata
func (api *EncodingEncodingsStreamsHdrDolbyVisionAPI) Delete(encodingId string, streamId string, hdrId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
        params.PathParams["hdr_id"] = hdrId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/streams/{stream_id}/hdr/dolby-vision/{hdr_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get Dolby Vision Metadata Details
func (api *EncodingEncodingsStreamsHdrDolbyVisionAPI) Get(encodingId string, streamId string, hdrId string) (*model.DolbyVisionMetadata, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
        params.PathParams["hdr_id"] = hdrId
    }

    var responseModel model.DolbyVisionMetadata
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/streams/{stream_id}/hdr/dolby-vision/{hdr_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List Dolby Vision Metadata
func (api *EncodingEncodingsStreamsHdrDolbyVisionAPI) List(encodingId string, streamId string, queryParams ...func(*EncodingEncodingsStreamsHdrDolbyVisionAPIListQueryParams)) (*pagination.DolbyVisionMetadatasListPagination, error) {
    queryParameters := &EncodingEncodingsStreamsHdrDolbyVisionAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["stream_id"] = streamId
        params.QueryParams = queryParameters
    }

    var responseModel pagination.DolbyVisionMetadatasListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/streams/{stream_id}/hdr/dolby-vision", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingEncodingsStreamsHdrDolbyVisionAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsStreamsHdrDolbyVisionAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsStreamsHdrDolbyVisionAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


