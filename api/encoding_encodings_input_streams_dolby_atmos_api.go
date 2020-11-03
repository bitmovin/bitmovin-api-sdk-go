package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsInputStreamsDolbyAtmosAPI communicates with '/encoding/encodings/{encoding_id}/input-streams/dolby-atmos' endpoints
type EncodingEncodingsInputStreamsDolbyAtmosAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingEncodingsInputStreamsDolbyAtmosAPI constructor for EncodingEncodingsInputStreamsDolbyAtmosAPI that takes options as argument
func NewEncodingEncodingsInputStreamsDolbyAtmosAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsInputStreamsDolbyAtmosAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsInputStreamsDolbyAtmosAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsInputStreamsDolbyAtmosAPIWithClient constructor for EncodingEncodingsInputStreamsDolbyAtmosAPI that takes an APIClient as argument
func NewEncodingEncodingsInputStreamsDolbyAtmosAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsInputStreamsDolbyAtmosAPI {
    a := &EncodingEncodingsInputStreamsDolbyAtmosAPI{apiClient: apiClient}
    return a
}

// Create Add Dolby Atmos input stream
func (api *EncodingEncodingsInputStreamsDolbyAtmosAPI) Create(encodingId string, dolbyAtmosIngestInputStream model.DolbyAtmosIngestInputStream) (*model.DolbyAtmosIngestInputStream, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel model.DolbyAtmosIngestInputStream
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/input-streams/dolby-atmos", &dolbyAtmosIngestInputStream, &responseModel, reqParams)
    return &responseModel, err
}
// Delete Dolby Atmos input stream
func (api *EncodingEncodingsInputStreamsDolbyAtmosAPI) Delete(encodingId string, inputStreamId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["input_stream_id"] = inputStreamId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/input-streams/dolby-atmos/{input_stream_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get Dolby Atmos input stream details
func (api *EncodingEncodingsInputStreamsDolbyAtmosAPI) Get(encodingId string, inputStreamId string) (*model.DolbyAtmosIngestInputStream, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["input_stream_id"] = inputStreamId
    }

    var responseModel model.DolbyAtmosIngestInputStream
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/input-streams/dolby-atmos/{input_stream_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List Dolby Atmos input streams
func (api *EncodingEncodingsInputStreamsDolbyAtmosAPI) List(encodingId string, queryParams ...func(*EncodingEncodingsInputStreamsDolbyAtmosAPIListQueryParams)) (*pagination.DolbyAtmosIngestInputStreamsListPagination, error) {
    queryParameters := &EncodingEncodingsInputStreamsDolbyAtmosAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
    }

    var responseModel pagination.DolbyAtmosIngestInputStreamsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/input-streams/dolby-atmos", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingEncodingsInputStreamsDolbyAtmosAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsInputStreamsDolbyAtmosAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsInputStreamsDolbyAtmosAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


