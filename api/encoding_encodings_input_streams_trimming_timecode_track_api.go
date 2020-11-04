package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsInputStreamsTrimmingTimecodeTrackAPI communicates with '/encoding/encodings/{encoding_id}/input-streams/trimming/timecode-track' endpoints
type EncodingEncodingsInputStreamsTrimmingTimecodeTrackAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingEncodingsInputStreamsTrimmingTimecodeTrackAPI constructor for EncodingEncodingsInputStreamsTrimmingTimecodeTrackAPI that takes options as argument
func NewEncodingEncodingsInputStreamsTrimmingTimecodeTrackAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsInputStreamsTrimmingTimecodeTrackAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsInputStreamsTrimmingTimecodeTrackAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsInputStreamsTrimmingTimecodeTrackAPIWithClient constructor for EncodingEncodingsInputStreamsTrimmingTimecodeTrackAPI that takes an APIClient as argument
func NewEncodingEncodingsInputStreamsTrimmingTimecodeTrackAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsInputStreamsTrimmingTimecodeTrackAPI {
	a := &EncodingEncodingsInputStreamsTrimmingTimecodeTrackAPI{apiClient: apiClient}
	return a
}

// Create Add Timecode Track Trimming Input Stream
func (api *EncodingEncodingsInputStreamsTrimmingTimecodeTrackAPI) Create(encodingId string, timecodeTrackTrimmingInputStream model.TimecodeTrackTrimmingInputStream) (*model.TimecodeTrackTrimmingInputStream, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	var responseModel model.TimecodeTrackTrimmingInputStream
	err := api.apiClient.Post("/encoding/encodings/{encoding_id}/input-streams/trimming/timecode-track", &timecodeTrackTrimmingInputStream, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Timecode Track Trimming Input Stream
func (api *EncodingEncodingsInputStreamsTrimmingTimecodeTrackAPI) Delete(encodingId string, inputStreamId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["input_stream_id"] = inputStreamId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/input-streams/trimming/timecode-track/{input_stream_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Timecode Track Trimming Input Stream Details
func (api *EncodingEncodingsInputStreamsTrimmingTimecodeTrackAPI) Get(encodingId string, inputStreamId string) (*model.TimecodeTrackTrimmingInputStream, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["input_stream_id"] = inputStreamId
	}

	var responseModel model.TimecodeTrackTrimmingInputStream
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/input-streams/trimming/timecode-track/{input_stream_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List Timecode Track Trimming Input Streams
func (api *EncodingEncodingsInputStreamsTrimmingTimecodeTrackAPI) List(encodingId string, queryParams ...func(*EncodingEncodingsInputStreamsTrimmingTimecodeTrackAPIListQueryParams)) (*pagination.TimecodeTrackTrimmingInputStreamsListPagination, error) {
	queryParameters := &EncodingEncodingsInputStreamsTrimmingTimecodeTrackAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.TimecodeTrackTrimmingInputStreamsListPagination
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/input-streams/trimming/timecode-track", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingEncodingsInputStreamsTrimmingTimecodeTrackAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsInputStreamsTrimmingTimecodeTrackAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsInputStreamsTrimmingTimecodeTrackAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
