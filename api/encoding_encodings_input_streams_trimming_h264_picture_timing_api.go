package api

import (
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsInputStreamsTrimmingH264PictureTimingAPI communicates with '/encoding/encodings/{encoding_id}/input-streams/trimming/h264-picture-timing' endpoints
type EncodingEncodingsInputStreamsTrimmingH264PictureTimingAPI struct {
    apiClient *apiclient.APIClient
}

// NewEncodingEncodingsInputStreamsTrimmingH264PictureTimingAPI constructor for EncodingEncodingsInputStreamsTrimmingH264PictureTimingAPI that takes options as argument
func NewEncodingEncodingsInputStreamsTrimmingH264PictureTimingAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsInputStreamsTrimmingH264PictureTimingAPI, error) {
    apiClient, err := apiclient.NewAPIClient(options...)
    if err != nil {
        return nil, err
    }

    return NewEncodingEncodingsInputStreamsTrimmingH264PictureTimingAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsInputStreamsTrimmingH264PictureTimingAPIWithClient constructor for EncodingEncodingsInputStreamsTrimmingH264PictureTimingAPI that takes an APIClient as argument
func NewEncodingEncodingsInputStreamsTrimmingH264PictureTimingAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsInputStreamsTrimmingH264PictureTimingAPI {
    a := &EncodingEncodingsInputStreamsTrimmingH264PictureTimingAPI{apiClient: apiClient}
    return a
}

// Create Add H264 Picture Timing Trimming Input Stream
func (api *EncodingEncodingsInputStreamsTrimmingH264PictureTimingAPI) Create(encodingId string, h264PictureTimingTrimmingInputStream model.H264PictureTimingTrimmingInputStream) (*model.H264PictureTimingTrimmingInputStream, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel model.H264PictureTimingTrimmingInputStream
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/input-streams/trimming/h264-picture-timing", &h264PictureTimingTrimmingInputStream, &responseModel, reqParams)
    return &responseModel, err
}
// Delete H264 Picture Timing Trimming Input Stream
func (api *EncodingEncodingsInputStreamsTrimmingH264PictureTimingAPI) Delete(encodingId string, inputStreamId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["input_stream_id"] = inputStreamId
    }

    var responseModel model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/input-streams/trimming/h264-picture-timing/{input_stream_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// Get H264 Picture Timing Trimming Input Stream Details
func (api *EncodingEncodingsInputStreamsTrimmingH264PictureTimingAPI) Get(encodingId string, inputStreamId string) (*model.H264PictureTimingTrimmingInputStream, error) {
    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["input_stream_id"] = inputStreamId
    }

    var responseModel model.H264PictureTimingTrimmingInputStream
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/input-streams/trimming/h264-picture-timing/{input_stream_id}", nil, &responseModel, reqParams)
    return &responseModel, err
}
// List H264 Picture Timing Trimming Input Streams
func (api *EncodingEncodingsInputStreamsTrimmingH264PictureTimingAPI) List(encodingId string, queryParams ...func(*EncodingEncodingsInputStreamsTrimmingH264PictureTimingAPIListQueryParams)) (*pagination.H264PictureTimingTrimmingInputStreamsListPagination, error) {
    queryParameters := &EncodingEncodingsInputStreamsTrimmingH264PictureTimingAPIListQueryParams{}
    for _, queryParam := range queryParams {
        queryParam(queryParameters)
    }

    reqParams := func(params *apiclient.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
    }

    var responseModel pagination.H264PictureTimingTrimmingInputStreamsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/input-streams/trimming/h264-picture-timing", nil, &responseModel, reqParams)
    return &responseModel, err
}

// EncodingEncodingsInputStreamsTrimmingH264PictureTimingAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsInputStreamsTrimmingH264PictureTimingAPIListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsInputStreamsTrimmingH264PictureTimingAPIListQueryParams) Params() map[string]string {
    return apiclient.GetParamsMap(q)
}


