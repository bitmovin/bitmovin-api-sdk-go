package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsStreamsBurnInSubtitlesSrtAPI communicates with '/encoding/encodings/{encoding_id}/streams/{stream_id}/burn-in-subtitles/srt' endpoints
type EncodingEncodingsStreamsBurnInSubtitlesSrtAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingEncodingsStreamsBurnInSubtitlesSrtAPI constructor for EncodingEncodingsStreamsBurnInSubtitlesSrtAPI that takes options as argument
func NewEncodingEncodingsStreamsBurnInSubtitlesSrtAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsStreamsBurnInSubtitlesSrtAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsStreamsBurnInSubtitlesSrtAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsStreamsBurnInSubtitlesSrtAPIWithClient constructor for EncodingEncodingsStreamsBurnInSubtitlesSrtAPI that takes an APIClient as argument
func NewEncodingEncodingsStreamsBurnInSubtitlesSrtAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsStreamsBurnInSubtitlesSrtAPI {
	a := &EncodingEncodingsStreamsBurnInSubtitlesSrtAPI{apiClient: apiClient}
	return a
}

// Create Burn-In SRT Subtitle into Stream
func (api *EncodingEncodingsStreamsBurnInSubtitlesSrtAPI) Create(encodingId string, streamId string, burnInSubtitleSrt model.BurnInSubtitleSrt) (*model.BurnInSubtitleSrt, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["stream_id"] = streamId
	}

	var responseModel model.BurnInSubtitleSrt
	err := api.apiClient.Post("/encoding/encodings/{encoding_id}/streams/{stream_id}/burn-in-subtitles/srt", &burnInSubtitleSrt, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Burn-In SRT Subtitle from Stream
func (api *EncodingEncodingsStreamsBurnInSubtitlesSrtAPI) Delete(encodingId string, streamId string, subtitleId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["stream_id"] = streamId
		params.PathParams["subtitle_id"] = subtitleId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/streams/{stream_id}/burn-in-subtitles/srt/{subtitle_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Burn-In SRT Subtitle Details
func (api *EncodingEncodingsStreamsBurnInSubtitlesSrtAPI) Get(encodingId string, streamId string, subtitleId string) (*model.BurnInSubtitleSrt, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["stream_id"] = streamId
		params.PathParams["subtitle_id"] = subtitleId
	}

	var responseModel model.BurnInSubtitleSrt
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/streams/{stream_id}/burn-in-subtitles/srt/{subtitle_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List the Burn-In SRT subtitles of a stream
func (api *EncodingEncodingsStreamsBurnInSubtitlesSrtAPI) List(encodingId string, streamId string, queryParams ...func(*EncodingEncodingsStreamsBurnInSubtitlesSrtAPIListQueryParams)) (*pagination.BurnInSubtitleSrtsListPagination, error) {
	queryParameters := &EncodingEncodingsStreamsBurnInSubtitlesSrtAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["stream_id"] = streamId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.BurnInSubtitleSrtsListPagination
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/streams/{stream_id}/burn-in-subtitles/srt", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingEncodingsStreamsBurnInSubtitlesSrtAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsStreamsBurnInSubtitlesSrtAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsStreamsBurnInSubtitlesSrtAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
