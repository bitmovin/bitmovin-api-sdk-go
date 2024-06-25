package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsMuxingsAPI communicates with '/encoding/encodings/{encoding_id}/muxings' endpoints
type EncodingEncodingsMuxingsAPI struct {
	apiClient *apiclient.APIClient

	// Type communicates with '/encoding/encodings/{encoding_id}/muxings/{muxing_id}/type' endpoints
	Type *EncodingEncodingsMuxingsTypeAPI
	// Fmp4 communicates with '/encoding/encodings/{encoding_id}/muxings/fmp4' endpoints
	Fmp4 *EncodingEncodingsMuxingsFmp4API
	// ChunkedText communicates with '/encoding/encodings/{encoding_id}/muxings/chunked-text' endpoints
	ChunkedText *EncodingEncodingsMuxingsChunkedTextAPI
	// Cmaf communicates with '/encoding/encodings/{encoding_id}/muxings/cmaf' endpoints
	Cmaf *EncodingEncodingsMuxingsCmafAPI
	// SegmentedRaw communicates with '/encoding/encodings/{encoding_id}/muxings/segmented-raw' endpoints
	SegmentedRaw *EncodingEncodingsMuxingsSegmentedRawAPI
	// PackedAudio communicates with '/encoding/encodings/{encoding_id}/muxings/packed-audio' endpoints
	PackedAudio *EncodingEncodingsMuxingsPackedAudioAPI
	// Text communicates with '/encoding/encodings/{encoding_id}/muxings/text' endpoints
	Text *EncodingEncodingsMuxingsTextAPI
	// Ts communicates with '/encoding/encodings/{encoding_id}/muxings/ts' endpoints
	Ts *EncodingEncodingsMuxingsTsAPI
	// Webm communicates with '/encoding/encodings/{encoding_id}/muxings/webm' endpoints
	Webm *EncodingEncodingsMuxingsWebmAPI
	// Mp3 communicates with '/encoding/encodings/{encoding_id}/muxings/mp3' endpoints
	Mp3 *EncodingEncodingsMuxingsMp3API
	// Mp4 communicates with '/encoding/encodings/{encoding_id}/muxings/mp4' endpoints
	Mp4 *EncodingEncodingsMuxingsMp4API
	// Mxf communicates with '/encoding/encodings/{encoding_id}/muxings/mxf' endpoints
	Mxf *EncodingEncodingsMuxingsMxfAPI
	// ProgressiveTs communicates with '/encoding/encodings/{encoding_id}/muxings/progressive-ts' endpoints
	ProgressiveTs *EncodingEncodingsMuxingsProgressiveTsAPI
	// BroadcastTs communicates with '/encoding/encodings/{encoding_id}/muxings/broadcast-ts' endpoints
	BroadcastTs *EncodingEncodingsMuxingsBroadcastTsAPI
	// ProgressiveWav communicates with '/encoding/encodings/{encoding_id}/muxings/progressive-wav' endpoints
	ProgressiveWav *EncodingEncodingsMuxingsProgressiveWavAPI
	// ProgressiveWebm communicates with '/encoding/encodings/{encoding_id}/muxings/progressive-webm' endpoints
	ProgressiveWebm *EncodingEncodingsMuxingsProgressiveWebmAPI
	// ProgressiveMov communicates with '/encoding/encodings/{encoding_id}/muxings/progressive-mov' endpoints
	ProgressiveMov *EncodingEncodingsMuxingsProgressiveMovAPI
}

// NewEncodingEncodingsMuxingsAPI constructor for EncodingEncodingsMuxingsAPI that takes options as argument
func NewEncodingEncodingsMuxingsAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMuxingsAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsAPIWithClient constructor for EncodingEncodingsMuxingsAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsAPI {
	a := &EncodingEncodingsMuxingsAPI{apiClient: apiClient}
	a.Type = NewEncodingEncodingsMuxingsTypeAPIWithClient(apiClient)
	a.Fmp4 = NewEncodingEncodingsMuxingsFmp4APIWithClient(apiClient)
	a.ChunkedText = NewEncodingEncodingsMuxingsChunkedTextAPIWithClient(apiClient)
	a.Cmaf = NewEncodingEncodingsMuxingsCmafAPIWithClient(apiClient)
	a.SegmentedRaw = NewEncodingEncodingsMuxingsSegmentedRawAPIWithClient(apiClient)
	a.PackedAudio = NewEncodingEncodingsMuxingsPackedAudioAPIWithClient(apiClient)
	a.Text = NewEncodingEncodingsMuxingsTextAPIWithClient(apiClient)
	a.Ts = NewEncodingEncodingsMuxingsTsAPIWithClient(apiClient)
	a.Webm = NewEncodingEncodingsMuxingsWebmAPIWithClient(apiClient)
	a.Mp3 = NewEncodingEncodingsMuxingsMp3APIWithClient(apiClient)
	a.Mp4 = NewEncodingEncodingsMuxingsMp4APIWithClient(apiClient)
	a.Mxf = NewEncodingEncodingsMuxingsMxfAPIWithClient(apiClient)
	a.ProgressiveTs = NewEncodingEncodingsMuxingsProgressiveTsAPIWithClient(apiClient)
	a.BroadcastTs = NewEncodingEncodingsMuxingsBroadcastTsAPIWithClient(apiClient)
	a.ProgressiveWav = NewEncodingEncodingsMuxingsProgressiveWavAPIWithClient(apiClient)
	a.ProgressiveWebm = NewEncodingEncodingsMuxingsProgressiveWebmAPIWithClient(apiClient)
	a.ProgressiveMov = NewEncodingEncodingsMuxingsProgressiveMovAPIWithClient(apiClient)

	return a
}

// Get Muxing Details
func (api *EncodingEncodingsMuxingsAPI) Get(encodingId string, muxingId string) (*model.Muxing, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
	}

	var responseModel model.Muxing
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/{muxing_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List All Muxings
func (api *EncodingEncodingsMuxingsAPI) List(encodingId string, queryParams ...func(*EncodingEncodingsMuxingsAPIListQueryParams)) (*pagination.MuxingsListPagination, error) {
	queryParameters := &EncodingEncodingsMuxingsAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.MuxingsListPagination
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingEncodingsMuxingsAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsMuxingsAPIListQueryParams struct {
	Offset     int32            `query:"offset"`
	Limit      int32            `query:"limit"`
	StreamMode model.StreamMode `query:"streamMode"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsMuxingsAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
