package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsStreamsBurnInSubtitlesDvbsubAPI communicates with '/encoding/encodings/{encoding_id}/streams/{stream_id}/burn-in-subtitles/dvbsub' endpoints
type EncodingEncodingsStreamsBurnInSubtitlesDvbsubAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingEncodingsStreamsBurnInSubtitlesDvbsubAPI constructor for EncodingEncodingsStreamsBurnInSubtitlesDvbsubAPI that takes options as argument
func NewEncodingEncodingsStreamsBurnInSubtitlesDvbsubAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsStreamsBurnInSubtitlesDvbsubAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsStreamsBurnInSubtitlesDvbsubAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsStreamsBurnInSubtitlesDvbsubAPIWithClient constructor for EncodingEncodingsStreamsBurnInSubtitlesDvbsubAPI that takes an APIClient as argument
func NewEncodingEncodingsStreamsBurnInSubtitlesDvbsubAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsStreamsBurnInSubtitlesDvbsubAPI {
	a := &EncodingEncodingsStreamsBurnInSubtitlesDvbsubAPI{apiClient: apiClient}
	return a
}

// Create Burn-In DVB-SUB Subtitle into Stream
func (api *EncodingEncodingsStreamsBurnInSubtitlesDvbsubAPI) Create(encodingId string, streamId string, burnInSubtitleDvbSub model.BurnInSubtitleDvbSub) (*model.BurnInSubtitleDvbSub, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["stream_id"] = streamId
	}

	var responseModel model.BurnInSubtitleDvbSub
	err := api.apiClient.Post("/encoding/encodings/{encoding_id}/streams/{stream_id}/burn-in-subtitles/dvbsub", &burnInSubtitleDvbSub, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Burn-In DVB-SUB Subtitle from Stream
func (api *EncodingEncodingsStreamsBurnInSubtitlesDvbsubAPI) Delete(encodingId string, streamId string, subtitleId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["stream_id"] = streamId
		params.PathParams["subtitle_id"] = subtitleId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/streams/{stream_id}/burn-in-subtitles/dvbsub/{subtitle_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Burn-In DVB-SUB Subtitle Details
func (api *EncodingEncodingsStreamsBurnInSubtitlesDvbsubAPI) Get(encodingId string, streamId string, subtitleId string) (*model.BurnInSubtitleDvbSub, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["stream_id"] = streamId
		params.PathParams["subtitle_id"] = subtitleId
	}

	var responseModel model.BurnInSubtitleDvbSub
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/streams/{stream_id}/burn-in-subtitles/dvbsub/{subtitle_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List the Burn-In DVB-SUB subtitles of a stream
func (api *EncodingEncodingsStreamsBurnInSubtitlesDvbsubAPI) List(encodingId string, streamId string, queryParams ...func(*EncodingEncodingsStreamsBurnInSubtitlesDvbsubAPIListQueryParams)) (*pagination.BurnInSubtitleDvbSubsListPagination, error) {
	queryParameters := &EncodingEncodingsStreamsBurnInSubtitlesDvbsubAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["stream_id"] = streamId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.BurnInSubtitleDvbSubsListPagination
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/streams/{stream_id}/burn-in-subtitles/dvbsub", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingEncodingsStreamsBurnInSubtitlesDvbsubAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsStreamsBurnInSubtitlesDvbsubAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsStreamsBurnInSubtitlesDvbsubAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
