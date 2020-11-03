package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsInputStreamsSubtitlesDvbSubtitleAPI communicates with '/encoding/encodings/{encoding_id}/input-streams/subtitles/dvb-subtitle' endpoints
type EncodingEncodingsInputStreamsSubtitlesDvbSubtitleAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingEncodingsInputStreamsSubtitlesDvbSubtitleAPI constructor for EncodingEncodingsInputStreamsSubtitlesDvbSubtitleAPI that takes options as argument
func NewEncodingEncodingsInputStreamsSubtitlesDvbSubtitleAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsInputStreamsSubtitlesDvbSubtitleAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsInputStreamsSubtitlesDvbSubtitleAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsInputStreamsSubtitlesDvbSubtitleAPIWithClient constructor for EncodingEncodingsInputStreamsSubtitlesDvbSubtitleAPI that takes an APIClient as argument
func NewEncodingEncodingsInputStreamsSubtitlesDvbSubtitleAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsInputStreamsSubtitlesDvbSubtitleAPI {
    a := &EncodingEncodingsInputStreamsSubtitlesDvbSubtitleAPI{apiClient: apiClient}
    return a
}

// Create Add DVB Subtitle Input Stream
func (api *EncodingEncodingsInputStreamsSubtitlesDvbSubtitleAPI) Create(encodingId string, dvbSubtitleInputStream model.DvbSubtitleInputStream) (*model.DvbSubtitleInputStream, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel model.DvbSubtitleInputStream
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/input-streams/subtitles/dvb-subtitle", &dvbSubtitleInputStream, &responseModel, reqParams)
    return &responseModel, err
}
// Delete DVB Subtitle Input Stream
func (api *EncodingEncodingsInputStreamsSubtitlesDvbSubtitleAPI) Delete(encodingId string, inputStreamId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["input_stream_id"] = inputStreamId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/input-streams/subtitles/dvb-subtitle/{input_stream_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get DVB Subtitle Input Stream Details
func (api *EncodingEncodingsInputStreamsSubtitlesDvbSubtitleAPI) Get(encodingId string, inputStreamId string) (*model.DvbSubtitleInputStream, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["input_stream_id"] = inputStreamId
    }

    var responseModel model.DvbSubtitleInputStream
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/input-streams/subtitles/dvb-subtitle/{input_stream_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List DVB Subtitle Input Streams
func (api *EncodingEncodingsInputStreamsSubtitlesDvbSubtitleAPI) List(encodingId string, queryParams ...func(*EncodingEncodingsInputStreamsSubtitlesDvbSubtitleAPIListQueryParams)) (*pagination.DvbSubtitleInputStreamsListPagination, error) {
    queryParameters := &EncodingEncodingsInputStreamsSubtitlesDvbSubtitleAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
    }

    var responseModel pagination.DvbSubtitleInputStreamsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/input-streams/subtitles/dvb-subtitle", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingEncodingsInputStreamsSubtitlesDvbSubtitleAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsInputStreamsSubtitlesDvbSubtitleAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsInputStreamsSubtitlesDvbSubtitleAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


