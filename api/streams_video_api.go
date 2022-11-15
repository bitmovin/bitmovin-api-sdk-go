package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// StreamsVideoAPI communicates with '/streams/video' endpoints
type StreamsVideoAPI struct {
	apiClient *apiclient.APIClient
}

// NewStreamsVideoAPI constructor for StreamsVideoAPI that takes options as argument
func NewStreamsVideoAPI(options ...apiclient.APIClientOption) (*StreamsVideoAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewStreamsVideoAPIWithClient(apiClient), nil
}

// NewStreamsVideoAPIWithClient constructor for StreamsVideoAPI that takes an APIClient as argument
func NewStreamsVideoAPIWithClient(apiClient *apiclient.APIClient) *StreamsVideoAPI {
	a := &StreamsVideoAPI{apiClient: apiClient}
	return a
}

// PatchStreamsVideo Update stream by id
func (api *StreamsVideoAPI) PatchStreamsVideo(streamId string, streamsVideoUpdateRequest model.StreamsVideoUpdateRequest) (*model.StreamsVideoResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["stream_id"] = streamId
	}

	var responseModel model.StreamsVideoResponse
	err := api.apiClient.Patch("/streams/video/{stream_id}", &streamsVideoUpdateRequest, &responseModel, reqParams)
	return &responseModel, err
}

// Create new Stream
func (api *StreamsVideoAPI) Create(streamsVideoCreateRequest model.StreamsVideoCreateRequest) (*model.StreamsVideoResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.StreamsVideoResponse
	err := api.apiClient.Post("/streams/video", &streamsVideoCreateRequest, &responseModel, reqParams)
	return &responseModel, err
}

// Get stream by id
func (api *StreamsVideoAPI) Get(streamId string) (*model.StreamsVideoResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["stream_id"] = streamId
	}

	var responseModel model.StreamsVideoResponse
	err := api.apiClient.Get("/streams/video/{stream_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List Get paginated list of streams
func (api *StreamsVideoAPI) List(queryParams ...func(*StreamsVideoAPIListQueryParams)) (*pagination.StreamsVideoResponsesListPagination, error) {
	queryParameters := &StreamsVideoAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.StreamsVideoResponsesListPagination
	err := api.apiClient.Get("/streams/video", nil, &responseModel, reqParams)
	return &responseModel, err
}

// StreamsVideoAPIListQueryParams contains all query parameters for the List endpoint
type StreamsVideoAPIListQueryParams struct {
	Offset int32  `query:"offset"`
	Limit  int32  `query:"limit"`
	Sort   string `query:"sort"`
}

// Params will return a map of query parameters
func (q *StreamsVideoAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
