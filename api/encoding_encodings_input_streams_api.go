package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsInputStreamsAPI communicates with '/encoding/encodings/{encoding_id}/input-streams' endpoints
type EncodingEncodingsInputStreamsAPI struct {
	apiClient *apiclient.APIClient

	// Type communicates with '/encoding/encodings/{encoding_id}/input-streams/{input_stream_id}/type' endpoints
	Type *EncodingEncodingsInputStreamsTypeAPI
	// AudioMix communicates with '/encoding/encodings/{encoding_id}/input-streams/audio-mix' endpoints
	AudioMix *EncodingEncodingsInputStreamsAudioMixAPI
	// Ingest communicates with '/encoding/encodings/{encoding_id}/input-streams/ingest' endpoints
	Ingest *EncodingEncodingsInputStreamsIngestAPI
	// Sidecar intermediary API object with no endpoints
	Sidecar *EncodingEncodingsInputStreamsSidecarAPI
	// Concatenation communicates with '/encoding/encodings/{encoding_id}/input-streams/concatenation' endpoints
	Concatenation *EncodingEncodingsInputStreamsConcatenationAPI
	// File communicates with '/encoding/encodings/{encoding_id}/input-streams/file' endpoints
	File *EncodingEncodingsInputStreamsFileAPI
	// Trimming intermediary API object with no endpoints
	Trimming *EncodingEncodingsInputStreamsTrimmingAPI
	// Subtitles intermediary API object with no endpoints
	Subtitles *EncodingEncodingsInputStreamsSubtitlesAPI
	// Captions intermediary API object with no endpoints
	Captions *EncodingEncodingsInputStreamsCaptionsAPI
	// DolbyAtmos communicates with '/encoding/encodings/{encoding_id}/input-streams/dolby-atmos' endpoints
	DolbyAtmos *EncodingEncodingsInputStreamsDolbyAtmosAPI
	// DolbyVision communicates with '/encoding/encodings/{encoding_id}/input-streams/dolby-vision' endpoints
	DolbyVision *EncodingEncodingsInputStreamsDolbyVisionAPI
}

// NewEncodingEncodingsInputStreamsAPI constructor for EncodingEncodingsInputStreamsAPI that takes options as argument
func NewEncodingEncodingsInputStreamsAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsInputStreamsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsInputStreamsAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsInputStreamsAPIWithClient constructor for EncodingEncodingsInputStreamsAPI that takes an APIClient as argument
func NewEncodingEncodingsInputStreamsAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsInputStreamsAPI {
	a := &EncodingEncodingsInputStreamsAPI{apiClient: apiClient}
	a.Type = NewEncodingEncodingsInputStreamsTypeAPIWithClient(apiClient)
	a.AudioMix = NewEncodingEncodingsInputStreamsAudioMixAPIWithClient(apiClient)
	a.Ingest = NewEncodingEncodingsInputStreamsIngestAPIWithClient(apiClient)
	a.Sidecar = NewEncodingEncodingsInputStreamsSidecarAPIWithClient(apiClient)
	a.Concatenation = NewEncodingEncodingsInputStreamsConcatenationAPIWithClient(apiClient)
	a.File = NewEncodingEncodingsInputStreamsFileAPIWithClient(apiClient)
	a.Trimming = NewEncodingEncodingsInputStreamsTrimmingAPIWithClient(apiClient)
	a.Subtitles = NewEncodingEncodingsInputStreamsSubtitlesAPIWithClient(apiClient)
	a.Captions = NewEncodingEncodingsInputStreamsCaptionsAPIWithClient(apiClient)
	a.DolbyAtmos = NewEncodingEncodingsInputStreamsDolbyAtmosAPIWithClient(apiClient)
	a.DolbyVision = NewEncodingEncodingsInputStreamsDolbyVisionAPIWithClient(apiClient)

	return a
}

// Get Input Stream Details
func (api *EncodingEncodingsInputStreamsAPI) Get(encodingId string, inputStreamId string) (*model.InputStream, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["input_stream_id"] = inputStreamId
	}

	var responseModel model.InputStream
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/input-streams/{input_stream_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List All Input Streams
func (api *EncodingEncodingsInputStreamsAPI) List(encodingId string, queryParams ...func(*EncodingEncodingsInputStreamsAPIListQueryParams)) (*pagination.InputStreamsListPagination, error) {
	queryParameters := &EncodingEncodingsInputStreamsAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.InputStreamsListPagination
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/input-streams", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingEncodingsInputStreamsAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsInputStreamsAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsInputStreamsAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
