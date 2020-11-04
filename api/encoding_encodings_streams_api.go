package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsStreamsAPI communicates with '/encoding/encodings/{encoding_id}/streams' endpoints
type EncodingEncodingsStreamsAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/encodings/{encoding_id}/streams/{stream_id}/customData' endpoints
	Customdata *EncodingEncodingsStreamsCustomdataAPI
	// Input communicates with '/encoding/encodings/{encoding_id}/streams/{stream_id}/input' endpoints
	Input *EncodingEncodingsStreamsInputAPI
	// Inputs communicates with '/encoding/encodings/{encoding_id}/streams/{stream_id}/inputs' endpoints
	Inputs *EncodingEncodingsStreamsInputsAPI
	// Filters communicates with '/encoding/encodings/{encoding_id}/streams/{stream_id}/filters' endpoints
	Filters *EncodingEncodingsStreamsFiltersAPI
	// BurnInSubtitles intermediary API object with no endpoints
	BurnInSubtitles *EncodingEncodingsStreamsBurnInSubtitlesAPI
	// Captions intermediary API object with no endpoints
	Captions *EncodingEncodingsStreamsCaptionsAPI
	// Bifs communicates with '/encoding/encodings/{encoding_id}/streams/{stream_id}/bifs' endpoints
	Bifs *EncodingEncodingsStreamsBifsAPI
	// Hdr intermediary API object with no endpoints
	Hdr *EncodingEncodingsStreamsHdrAPI
	// Thumbnails communicates with '/encoding/encodings/{encoding_id}/streams/{stream_id}/thumbnails' endpoints
	Thumbnails *EncodingEncodingsStreamsThumbnailsAPI
	// Sprites communicates with '/encoding/encodings/{encoding_id}/streams/{stream_id}/sprites' endpoints
	Sprites *EncodingEncodingsStreamsSpritesAPI
	// Qc intermediary API object with no endpoints
	Qc *EncodingEncodingsStreamsQcAPI
}

// NewEncodingEncodingsStreamsAPI constructor for EncodingEncodingsStreamsAPI that takes options as argument
func NewEncodingEncodingsStreamsAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsStreamsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsStreamsAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsStreamsAPIWithClient constructor for EncodingEncodingsStreamsAPI that takes an APIClient as argument
func NewEncodingEncodingsStreamsAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsStreamsAPI {
	a := &EncodingEncodingsStreamsAPI{apiClient: apiClient}
	a.Customdata = NewEncodingEncodingsStreamsCustomdataAPIWithClient(apiClient)
	a.Input = NewEncodingEncodingsStreamsInputAPIWithClient(apiClient)
	a.Inputs = NewEncodingEncodingsStreamsInputsAPIWithClient(apiClient)
	a.Filters = NewEncodingEncodingsStreamsFiltersAPIWithClient(apiClient)
	a.BurnInSubtitles = NewEncodingEncodingsStreamsBurnInSubtitlesAPIWithClient(apiClient)
	a.Captions = NewEncodingEncodingsStreamsCaptionsAPIWithClient(apiClient)
	a.Bifs = NewEncodingEncodingsStreamsBifsAPIWithClient(apiClient)
	a.Hdr = NewEncodingEncodingsStreamsHdrAPIWithClient(apiClient)
	a.Thumbnails = NewEncodingEncodingsStreamsThumbnailsAPIWithClient(apiClient)
	a.Sprites = NewEncodingEncodingsStreamsSpritesAPIWithClient(apiClient)
	a.Qc = NewEncodingEncodingsStreamsQcAPIWithClient(apiClient)

	return a
}

// Create Add Stream
func (api *EncodingEncodingsStreamsAPI) Create(encodingId string, stream model.Stream) (*model.Stream, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	var responseModel model.Stream
	err := api.apiClient.Post("/encoding/encodings/{encoding_id}/streams", &stream, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Stream
func (api *EncodingEncodingsStreamsAPI) Delete(encodingId string, streamId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["stream_id"] = streamId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/streams/{stream_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Stream Details
func (api *EncodingEncodingsStreamsAPI) Get(encodingId string, streamId string) (*model.Stream, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["stream_id"] = streamId
	}

	var responseModel model.Stream
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/streams/{stream_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List Streams
func (api *EncodingEncodingsStreamsAPI) List(encodingId string, queryParams ...func(*EncodingEncodingsStreamsAPIListQueryParams)) (*pagination.StreamsListPagination, error) {
	queryParameters := &EncodingEncodingsStreamsAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.StreamsListPagination
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/streams", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingEncodingsStreamsAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsStreamsAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsStreamsAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
