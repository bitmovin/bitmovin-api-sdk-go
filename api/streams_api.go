package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// StreamsAPI communicates with '/streams' endpoints
type StreamsAPI struct {
	apiClient *apiclient.APIClient
}

// NewStreamsAPI constructor for StreamsAPI that takes options as argument
func NewStreamsAPI(options ...apiclient.APIClientOption) (*StreamsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewStreamsAPIWithClient(apiClient), nil
}

// NewStreamsAPIWithClient constructor for StreamsAPI that takes an APIClient as argument
func NewStreamsAPIWithClient(apiClient *apiclient.APIClient) *StreamsAPI {
	a := &StreamsAPI{apiClient: apiClient}
	return a
}

// PatchBitmovinStreamsStreamsByStreamId Update Stream by id
func (api *StreamsAPI) PatchBitmovinStreamsStreamsByStreamId(streamId string, updateBitmovinStreamRequest model.UpdateBitmovinStreamRequest) (*model.BitmovinStreamResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["stream_id"] = streamId
	}

	var responseModel model.BitmovinStreamResponse
	err := api.apiClient.Patch("/streams/{stream_id}", &updateBitmovinStreamRequest, &responseModel, reqParams)
	return &responseModel, err
}

// Create new Stream
func (api *StreamsAPI) Create(createBitmovinStreamRequest model.CreateBitmovinStreamRequest) (*model.BitmovinStreamResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.BitmovinStreamResponse
	err := api.apiClient.Post("/streams", &createBitmovinStreamRequest, &responseModel, reqParams)
	return &responseModel, err
}

// Get Stream by id
func (api *StreamsAPI) Get(streamId string) (*model.BitmovinStreamResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["stream_id"] = streamId
	}

	var responseModel model.BitmovinStreamResponse
	err := api.apiClient.Get("/streams/{stream_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List Get paginated list of Streams
func (api *StreamsAPI) List(queryParams ...func(*StreamsAPIListQueryParams)) (*pagination.BitmovinStreamResponsesListPagination, error) {
	queryParameters := &StreamsAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.BitmovinStreamResponsesListPagination
	err := api.apiClient.Get("/streams", nil, &responseModel, reqParams)
	return &responseModel, err
}

// StreamsAPIListQueryParams contains all query parameters for the List endpoint
type StreamsAPIListQueryParams struct {
	Offset int32  `query:"offset"`
	Limit  int32  `query:"limit"`
	Sort   string `query:"sort"`
}

// Params will return a map of query parameters
func (q *StreamsAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
