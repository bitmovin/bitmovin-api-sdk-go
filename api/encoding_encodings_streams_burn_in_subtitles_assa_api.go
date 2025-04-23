package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsStreamsBurnInSubtitlesAssaAPI communicates with '/encoding/encodings/{encoding_id}/streams/{stream_id}/burn-in-subtitles/assa' endpoints
type EncodingEncodingsStreamsBurnInSubtitlesAssaAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingEncodingsStreamsBurnInSubtitlesAssaAPI constructor for EncodingEncodingsStreamsBurnInSubtitlesAssaAPI that takes options as argument
func NewEncodingEncodingsStreamsBurnInSubtitlesAssaAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsStreamsBurnInSubtitlesAssaAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsStreamsBurnInSubtitlesAssaAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsStreamsBurnInSubtitlesAssaAPIWithClient constructor for EncodingEncodingsStreamsBurnInSubtitlesAssaAPI that takes an APIClient as argument
func NewEncodingEncodingsStreamsBurnInSubtitlesAssaAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsStreamsBurnInSubtitlesAssaAPI {
	a := &EncodingEncodingsStreamsBurnInSubtitlesAssaAPI{apiClient: apiClient}
	return a
}

// Create Burn-In ASSA Subtitle into Stream
func (api *EncodingEncodingsStreamsBurnInSubtitlesAssaAPI) Create(encodingId string, streamId string, burnInSubtitleAssa model.BurnInSubtitleAssa) (*model.BurnInSubtitleAssa, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["stream_id"] = streamId
	}

	var responseModel model.BurnInSubtitleAssa
	err := api.apiClient.Post("/encoding/encodings/{encoding_id}/streams/{stream_id}/burn-in-subtitles/assa", &burnInSubtitleAssa, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Burn-In ASSA Subtitle from Stream
func (api *EncodingEncodingsStreamsBurnInSubtitlesAssaAPI) Delete(encodingId string, streamId string, subtitleId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["stream_id"] = streamId
		params.PathParams["subtitle_id"] = subtitleId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/streams/{stream_id}/burn-in-subtitles/assa/{subtitle_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Burn-In ASSA Subtitle Details
func (api *EncodingEncodingsStreamsBurnInSubtitlesAssaAPI) Get(encodingId string, streamId string, subtitleId string) (*model.BurnInSubtitleAssa, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["stream_id"] = streamId
		params.PathParams["subtitle_id"] = subtitleId
	}

	var responseModel model.BurnInSubtitleAssa
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/streams/{stream_id}/burn-in-subtitles/assa/{subtitle_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List the Burn-In ASSA subtitles of a stream
func (api *EncodingEncodingsStreamsBurnInSubtitlesAssaAPI) List(encodingId string, streamId string, queryParams ...func(*EncodingEncodingsStreamsBurnInSubtitlesAssaAPIListQueryParams)) (*pagination.BurnInSubtitleAssasListPagination, error) {
	queryParameters := &EncodingEncodingsStreamsBurnInSubtitlesAssaAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["stream_id"] = streamId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.BurnInSubtitleAssasListPagination
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/streams/{stream_id}/burn-in-subtitles/assa", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingEncodingsStreamsBurnInSubtitlesAssaAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsStreamsBurnInSubtitlesAssaAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsStreamsBurnInSubtitlesAssaAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
