package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// StreamsLiveAPI communicates with '/streams/live' endpoints
type StreamsLiveAPI struct {
	apiClient *apiclient.APIClient

	// Stop communicates with '/streams/live/{stream_id}/stop' endpoints
	Stop *StreamsLiveStopAPI
	// Start communicates with '/streams/live/{stream_id}/start' endpoints
	Start *StreamsLiveStartAPI
}

// NewStreamsLiveAPI constructor for StreamsLiveAPI that takes options as argument
func NewStreamsLiveAPI(options ...apiclient.APIClientOption) (*StreamsLiveAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewStreamsLiveAPIWithClient(apiClient), nil
}

// NewStreamsLiveAPIWithClient constructor for StreamsLiveAPI that takes an APIClient as argument
func NewStreamsLiveAPIWithClient(apiClient *apiclient.APIClient) *StreamsLiveAPI {
	a := &StreamsLiveAPI{apiClient: apiClient}
	a.Stop = NewStreamsLiveStopAPIWithClient(apiClient)
	a.Start = NewStreamsLiveStartAPIWithClient(apiClient)

	return a
}

// PatchStreamsLive Partially update live stream by id
func (api *StreamsLiveAPI) PatchStreamsLive(streamId string, streamsLiveUpdateRequest model.StreamsLiveUpdateRequest) (*model.StreamsLiveUpdateRequest, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["stream_id"] = streamId
	}

	var responseModel model.StreamsLiveUpdateRequest
	err := api.apiClient.Patch("/streams/live/{stream_id}", &streamsLiveUpdateRequest, &responseModel, reqParams)
	return &responseModel, err
}

// Create new live stream
func (api *StreamsLiveAPI) Create(streamsLiveCreateRequest model.StreamsLiveCreateRequest) (*model.StreamsLiveResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.StreamsLiveResponse
	err := api.apiClient.Post("/streams/live", &streamsLiveCreateRequest, &responseModel, reqParams)
	return &responseModel, err
}

// Get live stream by id
func (api *StreamsLiveAPI) Get(streamId string) (*model.StreamsLiveResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["stream_id"] = streamId
	}

	var responseModel model.StreamsLiveResponse
	err := api.apiClient.Get("/streams/live/{stream_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List Get paginated list of live streams
func (api *StreamsLiveAPI) List(queryParams ...func(*StreamsLiveAPIListQueryParams)) (*pagination.StreamsLiveResponsesListPagination, error) {
	queryParameters := &StreamsLiveAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.StreamsLiveResponsesListPagination
	err := api.apiClient.Get("/streams/live", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Update live stream by id
func (api *StreamsLiveAPI) Update(streamId string, streamsLiveUpdateRequest model.StreamsLiveUpdateRequest) (*model.StreamsLiveUpdateRequest, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["stream_id"] = streamId
	}

	var responseModel model.StreamsLiveUpdateRequest
	err := api.apiClient.Put("/streams/live/{stream_id}", &streamsLiveUpdateRequest, &responseModel, reqParams)
	return &responseModel, err
}

// StreamsLiveAPIListQueryParams contains all query parameters for the List endpoint
type StreamsLiveAPIListQueryParams struct {
	Offset int32  `query:"offset"`
	Limit  int32  `query:"limit"`
	Sort   string `query:"sort"`
}

// Params will return a map of query parameters
func (q *StreamsLiveAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
