package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsInputStreamsAudioMixAPI communicates with '/encoding/encodings/{encoding_id}/input-streams/audio-mix' endpoints
type EncodingEncodingsInputStreamsAudioMixAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingEncodingsInputStreamsAudioMixAPI constructor for EncodingEncodingsInputStreamsAudioMixAPI that takes options as argument
func NewEncodingEncodingsInputStreamsAudioMixAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsInputStreamsAudioMixAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsInputStreamsAudioMixAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsInputStreamsAudioMixAPIWithClient constructor for EncodingEncodingsInputStreamsAudioMixAPI that takes an APIClient as argument
func NewEncodingEncodingsInputStreamsAudioMixAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsInputStreamsAudioMixAPI {
    a := &EncodingEncodingsInputStreamsAudioMixAPI{apiClient: apiClient}
    return a
}

// Create Add audio mix input stream
func (api *EncodingEncodingsInputStreamsAudioMixAPI) Create(encodingId string, audioMixInputStream model.AudioMixInputStream) (*model.AudioMixInputStream, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel model.AudioMixInputStream
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/input-streams/audio-mix", &audioMixInputStream, &responseModel, reqParams)
    return &responseModel, err
}
// Delete audio mix input stream
func (api *EncodingEncodingsInputStreamsAudioMixAPI) Delete(encodingId string, inputStreamId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["input_stream_id"] = inputStreamId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/input-streams/audio-mix/{input_stream_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get Audio mix input stream details
func (api *EncodingEncodingsInputStreamsAudioMixAPI) Get(encodingId string, inputStreamId string) (*model.AudioMixInputStream, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["input_stream_id"] = inputStreamId
    }

    var responseModel model.AudioMixInputStream
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/input-streams/audio-mix/{input_stream_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List audio mix input stream
func (api *EncodingEncodingsInputStreamsAudioMixAPI) List(encodingId string, queryParams ...func(*EncodingEncodingsInputStreamsAudioMixAPIListQueryParams)) (*pagination.AudioMixInputStreamsListPagination, error) {
    queryParameters := &EncodingEncodingsInputStreamsAudioMixAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
    }

    var responseModel pagination.AudioMixInputStreamsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/input-streams/audio-mix", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingEncodingsInputStreamsAudioMixAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsInputStreamsAudioMixAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsInputStreamsAudioMixAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


