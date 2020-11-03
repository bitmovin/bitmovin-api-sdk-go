package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsInputStreamsCaptionsCea608API communicates with '/encoding/encodings/{encoding_id}/input-streams/captions/cea608' endpoints
type EncodingEncodingsInputStreamsCaptionsCea608API struct {
    apiClient *apiclient.APIClient
}

// NewEncodingEncodingsInputStreamsCaptionsCea608API constructor for EncodingEncodingsInputStreamsCaptionsCea608API that takes options as argument
func NewEncodingEncodingsInputStreamsCaptionsCea608API(options ...apiclient.APIClientOption) (*EncodingEncodingsInputStreamsCaptionsCea608API, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsInputStreamsCaptionsCea608APIWithClient(apiClient), nil
}

// NewEncodingEncodingsInputStreamsCaptionsCea608APIWithClient constructor for EncodingEncodingsInputStreamsCaptionsCea608API that takes an APIClient as argument
func NewEncodingEncodingsInputStreamsCaptionsCea608APIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsInputStreamsCaptionsCea608API {
    a := &EncodingEncodingsInputStreamsCaptionsCea608API{apiClient: apiClient}
    return a
}

// Create Add CEA 608 Input Stream
func (api *EncodingEncodingsInputStreamsCaptionsCea608API) Create(encodingId string, cea608CaptionInputStream model.Cea608CaptionInputStream) (*model.Cea608CaptionInputStream, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel model.Cea608CaptionInputStream
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/input-streams/captions/cea608", &cea608CaptionInputStream, &responseModel, reqParams)
    return &responseModel, err
}
// Delete CEA 608 Input Stream
func (api *EncodingEncodingsInputStreamsCaptionsCea608API) Delete(encodingId string, inputStreamId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["input_stream_id"] = inputStreamId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/input-streams/captions/cea608/{input_stream_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get CEA 608 Input Stream Details
func (api *EncodingEncodingsInputStreamsCaptionsCea608API) Get(encodingId string, inputStreamId string) (*model.Cea608CaptionInputStream, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["input_stream_id"] = inputStreamId
    }

    var responseModel model.Cea608CaptionInputStream
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/input-streams/captions/cea608/{input_stream_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List CEA 608 Input Streams
func (api *EncodingEncodingsInputStreamsCaptionsCea608API) List(encodingId string, queryParams ...func(*EncodingEncodingsInputStreamsCaptionsCea608APIListQueryParams)) (*pagination.Cea608CaptionInputStreamsListPagination, error) {
    queryParameters := &EncodingEncodingsInputStreamsCaptionsCea608APIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
    }

    var responseModel pagination.Cea608CaptionInputStreamsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/input-streams/captions/cea608", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingEncodingsInputStreamsCaptionsCea608APIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsInputStreamsCaptionsCea608APIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsInputStreamsCaptionsCea608APIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


