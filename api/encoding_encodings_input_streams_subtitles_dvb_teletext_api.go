package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsInputStreamsSubtitlesDvbTeletextAPI communicates with '/encoding/encodings/{encoding_id}/input-streams/subtitles/dvb-teletext' endpoints
type EncodingEncodingsInputStreamsSubtitlesDvbTeletextAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingEncodingsInputStreamsSubtitlesDvbTeletextAPI constructor for EncodingEncodingsInputStreamsSubtitlesDvbTeletextAPI that takes options as argument
func NewEncodingEncodingsInputStreamsSubtitlesDvbTeletextAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsInputStreamsSubtitlesDvbTeletextAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsInputStreamsSubtitlesDvbTeletextAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsInputStreamsSubtitlesDvbTeletextAPIWithClient constructor for EncodingEncodingsInputStreamsSubtitlesDvbTeletextAPI that takes an APIClient as argument
func NewEncodingEncodingsInputStreamsSubtitlesDvbTeletextAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsInputStreamsSubtitlesDvbTeletextAPI {
    a := &EncodingEncodingsInputStreamsSubtitlesDvbTeletextAPI{apiClient: apiClient}
    return a
}

// Create Add DVB-Teletext Input Stream
func (api *EncodingEncodingsInputStreamsSubtitlesDvbTeletextAPI) Create(encodingId string, dvbTeletextInputStream model.DvbTeletextInputStream) (*model.DvbTeletextInputStream, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel model.DvbTeletextInputStream
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/input-streams/subtitles/dvb-teletext", &dvbTeletextInputStream, &responseModel, reqParams)
    return &responseModel, err
}
// Delete DVB-Teletext Input Stream
func (api *EncodingEncodingsInputStreamsSubtitlesDvbTeletextAPI) Delete(encodingId string, inputStreamId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["input_stream_id"] = inputStreamId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/input-streams/subtitles/dvb-teletext/{input_stream_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get DVB-Teletext Input Stream Details
func (api *EncodingEncodingsInputStreamsSubtitlesDvbTeletextAPI) Get(encodingId string, inputStreamId string) (*model.DvbTeletextInputStream, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["input_stream_id"] = inputStreamId
    }

    var responseModel model.DvbTeletextInputStream
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/input-streams/subtitles/dvb-teletext/{input_stream_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List DVB-Teletext Input Streams
func (api *EncodingEncodingsInputStreamsSubtitlesDvbTeletextAPI) List(encodingId string, queryParams ...func(*EncodingEncodingsInputStreamsSubtitlesDvbTeletextAPIListQueryParams)) (*pagination.DvbTeletextInputStreamsListPagination, error) {
    queryParameters := &EncodingEncodingsInputStreamsSubtitlesDvbTeletextAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
    }

    var responseModel pagination.DvbTeletextInputStreamsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/input-streams/subtitles/dvb-teletext", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingEncodingsInputStreamsSubtitlesDvbTeletextAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsInputStreamsSubtitlesDvbTeletextAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsInputStreamsSubtitlesDvbTeletextAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


