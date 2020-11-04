package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// EncodingEncodingsStreamsFiltersAPI communicates with '/encoding/encodings/{encoding_id}/streams/{stream_id}/filters' endpoints
type EncodingEncodingsStreamsFiltersAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingEncodingsStreamsFiltersAPI constructor for EncodingEncodingsStreamsFiltersAPI that takes options as argument
func NewEncodingEncodingsStreamsFiltersAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsStreamsFiltersAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsStreamsFiltersAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsStreamsFiltersAPIWithClient constructor for EncodingEncodingsStreamsFiltersAPI that takes an APIClient as argument
func NewEncodingEncodingsStreamsFiltersAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsStreamsFiltersAPI {
	a := &EncodingEncodingsStreamsFiltersAPI{apiClient: apiClient}
	return a
}

// Create Add Filters to Stream
func (api *EncodingEncodingsStreamsFiltersAPI) Create(encodingId string, streamId string, streamFilter []model.StreamFilter) (*model.StreamFilterList, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["stream_id"] = streamId
	}

	var responseModel model.StreamFilterList
	err := api.apiClient.Post("/encoding/encodings/{encoding_id}/streams/{stream_id}/filters", &streamFilter, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Specific Filter from Stream
func (api *EncodingEncodingsStreamsFiltersAPI) Delete(encodingId string, streamId string, filterId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["stream_id"] = streamId
		params.PathParams["filter_id"] = filterId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/streams/{stream_id}/filters/{filter_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// DeleteAll Delete All Filters from Stream
func (api *EncodingEncodingsStreamsFiltersAPI) DeleteAll(encodingId string, streamId string) (*model.BitmovinResponseList, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["stream_id"] = streamId
	}

	var responseModel model.BitmovinResponseList
	err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/streams/{stream_id}/filters", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List the filters of a stream
func (api *EncodingEncodingsStreamsFiltersAPI) List(encodingId string, streamId string, queryParams ...func(*EncodingEncodingsStreamsFiltersAPIListQueryParams)) (*model.StreamFilterList, error) {
	queryParameters := &EncodingEncodingsStreamsFiltersAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["stream_id"] = streamId
		params.QueryParams = queryParameters
	}

	var responseModel model.StreamFilterList
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/streams/{stream_id}/filters", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingEncodingsStreamsFiltersAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsStreamsFiltersAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsStreamsFiltersAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
