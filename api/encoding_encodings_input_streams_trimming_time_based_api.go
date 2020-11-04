package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsInputStreamsTrimmingTimeBasedAPI communicates with '/encoding/encodings/{encoding_id}/input-streams/trimming/time-based' endpoints
type EncodingEncodingsInputStreamsTrimmingTimeBasedAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingEncodingsInputStreamsTrimmingTimeBasedAPI constructor for EncodingEncodingsInputStreamsTrimmingTimeBasedAPI that takes options as argument
func NewEncodingEncodingsInputStreamsTrimmingTimeBasedAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsInputStreamsTrimmingTimeBasedAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsInputStreamsTrimmingTimeBasedAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsInputStreamsTrimmingTimeBasedAPIWithClient constructor for EncodingEncodingsInputStreamsTrimmingTimeBasedAPI that takes an APIClient as argument
func NewEncodingEncodingsInputStreamsTrimmingTimeBasedAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsInputStreamsTrimmingTimeBasedAPI {
	a := &EncodingEncodingsInputStreamsTrimmingTimeBasedAPI{apiClient: apiClient}
	return a
}

// Create Add Time-Based Trimming Input Stream
func (api *EncodingEncodingsInputStreamsTrimmingTimeBasedAPI) Create(encodingId string, timeBasedTrimmingInputStream model.TimeBasedTrimmingInputStream) (*model.TimeBasedTrimmingInputStream, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	var responseModel model.TimeBasedTrimmingInputStream
	err := api.apiClient.Post("/encoding/encodings/{encoding_id}/input-streams/trimming/time-based", &timeBasedTrimmingInputStream, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Time-Based Trimming Input Stream
func (api *EncodingEncodingsInputStreamsTrimmingTimeBasedAPI) Delete(encodingId string, inputStreamId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["input_stream_id"] = inputStreamId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/input-streams/trimming/time-based/{input_stream_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Time-Based Trimming Input Stream Details
func (api *EncodingEncodingsInputStreamsTrimmingTimeBasedAPI) Get(encodingId string, inputStreamId string) (*model.TimeBasedTrimmingInputStream, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["input_stream_id"] = inputStreamId
	}

	var responseModel model.TimeBasedTrimmingInputStream
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/input-streams/trimming/time-based/{input_stream_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List Time-Based Trimming Input Streams
func (api *EncodingEncodingsInputStreamsTrimmingTimeBasedAPI) List(encodingId string, queryParams ...func(*EncodingEncodingsInputStreamsTrimmingTimeBasedAPIListQueryParams)) (*pagination.TimeBasedTrimmingInputStreamsListPagination, error) {
	queryParameters := &EncodingEncodingsInputStreamsTrimmingTimeBasedAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.TimeBasedTrimmingInputStreamsListPagination
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/input-streams/trimming/time-based", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingEncodingsInputStreamsTrimmingTimeBasedAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsInputStreamsTrimmingTimeBasedAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsInputStreamsTrimmingTimeBasedAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
