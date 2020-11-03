package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsInputStreamsIngestAPI communicates with '/encoding/encodings/{encoding_id}/input-streams/ingest' endpoints
type EncodingEncodingsInputStreamsIngestAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingEncodingsInputStreamsIngestAPI constructor for EncodingEncodingsInputStreamsIngestAPI that takes options as argument
func NewEncodingEncodingsInputStreamsIngestAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsInputStreamsIngestAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsInputStreamsIngestAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsInputStreamsIngestAPIWithClient constructor for EncodingEncodingsInputStreamsIngestAPI that takes an APIClient as argument
func NewEncodingEncodingsInputStreamsIngestAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsInputStreamsIngestAPI {
    a := &EncodingEncodingsInputStreamsIngestAPI{apiClient: apiClient}
    return a
}

// Create Add Ingest Input Stream
func (api *EncodingEncodingsInputStreamsIngestAPI) Create(encodingId string, ingestInputStream model.IngestInputStream) (*model.IngestInputStream, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel model.IngestInputStream
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/input-streams/ingest", &ingestInputStream, &responseModel, reqParams)
    return &responseModel, err
}
// Delete Ingest Input Stream
func (api *EncodingEncodingsInputStreamsIngestAPI) Delete(encodingId string, inputStreamId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["input_stream_id"] = inputStreamId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/input-streams/ingest/{input_stream_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get Ingest Input Stream Details
func (api *EncodingEncodingsInputStreamsIngestAPI) Get(encodingId string, inputStreamId string) (*model.IngestInputStream, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["input_stream_id"] = inputStreamId
    }

    var responseModel model.IngestInputStream
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/input-streams/ingest/{input_stream_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List Ingest Input Streams
func (api *EncodingEncodingsInputStreamsIngestAPI) List(encodingId string, queryParams ...func(*EncodingEncodingsInputStreamsIngestAPIListQueryParams)) (*pagination.IngestInputStreamsListPagination, error) {
    queryParameters := &EncodingEncodingsInputStreamsIngestAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
    }

    var responseModel pagination.IngestInputStreamsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/input-streams/ingest", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingEncodingsInputStreamsIngestAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsInputStreamsIngestAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsInputStreamsIngestAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


