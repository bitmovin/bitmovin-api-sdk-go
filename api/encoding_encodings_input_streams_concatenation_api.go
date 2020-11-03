package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsInputStreamsConcatenationAPI communicates with '/encoding/encodings/{encoding_id}/input-streams/concatenation' endpoints
type EncodingEncodingsInputStreamsConcatenationAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingEncodingsInputStreamsConcatenationAPI constructor for EncodingEncodingsInputStreamsConcatenationAPI that takes options as argument
func NewEncodingEncodingsInputStreamsConcatenationAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsInputStreamsConcatenationAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsInputStreamsConcatenationAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsInputStreamsConcatenationAPIWithClient constructor for EncodingEncodingsInputStreamsConcatenationAPI that takes an APIClient as argument
func NewEncodingEncodingsInputStreamsConcatenationAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsInputStreamsConcatenationAPI {
    a := &EncodingEncodingsInputStreamsConcatenationAPI{apiClient: apiClient}
    return a
}

// Create Add Concatenation Input Stream
func (api *EncodingEncodingsInputStreamsConcatenationAPI) Create(encodingId string, concatenationInputStream model.ConcatenationInputStream) (*model.ConcatenationInputStream, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel model.ConcatenationInputStream
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/input-streams/concatenation", &concatenationInputStream, &responseModel, reqParams)
    return &responseModel, err
}
// Delete Concatenation Input Stream
func (api *EncodingEncodingsInputStreamsConcatenationAPI) Delete(encodingId string, inputStreamId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["input_stream_id"] = inputStreamId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/input-streams/concatenation/{input_stream_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get Concatenation Input Stream Details
func (api *EncodingEncodingsInputStreamsConcatenationAPI) Get(encodingId string, inputStreamId string) (*model.ConcatenationInputStream, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["input_stream_id"] = inputStreamId
    }

    var responseModel model.ConcatenationInputStream
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/input-streams/concatenation/{input_stream_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List Concatenation Input Streams
func (api *EncodingEncodingsInputStreamsConcatenationAPI) List(encodingId string, queryParams ...func(*EncodingEncodingsInputStreamsConcatenationAPIListQueryParams)) (*pagination.ConcatenationInputStreamsListPagination, error) {
    queryParameters := &EncodingEncodingsInputStreamsConcatenationAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
    }

    var responseModel pagination.ConcatenationInputStreamsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/input-streams/concatenation", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingEncodingsInputStreamsConcatenationAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsInputStreamsConcatenationAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsInputStreamsConcatenationAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


