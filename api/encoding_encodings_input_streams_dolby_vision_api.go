package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsInputStreamsDolbyVisionAPI communicates with '/encoding/encodings/{encoding_id}/input-streams/dolby-vision' endpoints
type EncodingEncodingsInputStreamsDolbyVisionAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingEncodingsInputStreamsDolbyVisionAPI constructor for EncodingEncodingsInputStreamsDolbyVisionAPI that takes options as argument
func NewEncodingEncodingsInputStreamsDolbyVisionAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsInputStreamsDolbyVisionAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsInputStreamsDolbyVisionAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsInputStreamsDolbyVisionAPIWithClient constructor for EncodingEncodingsInputStreamsDolbyVisionAPI that takes an APIClient as argument
func NewEncodingEncodingsInputStreamsDolbyVisionAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsInputStreamsDolbyVisionAPI {
	a := &EncodingEncodingsInputStreamsDolbyVisionAPI{apiClient: apiClient}
	return a
}

// Create Add Dolby Vision input stream
func (api *EncodingEncodingsInputStreamsDolbyVisionAPI) Create(encodingId string, dolbyVisionInputStream model.DolbyVisionInputStream) (*model.DolbyVisionInputStream, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	var responseModel model.DolbyVisionInputStream
	err := api.apiClient.Post("/encoding/encodings/{encoding_id}/input-streams/dolby-vision", &dolbyVisionInputStream, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Dolby Vision input stream
func (api *EncodingEncodingsInputStreamsDolbyVisionAPI) Delete(encodingId string, dolbyVisionInputStreamId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["dolby_vision_input_stream_id"] = dolbyVisionInputStreamId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/input-streams/dolby-vision/{dolby_vision_input_stream_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Dolby Vision input stream details
func (api *EncodingEncodingsInputStreamsDolbyVisionAPI) Get(encodingId string, dolbyVisionInputStreamId string) (*model.DolbyVisionInputStream, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["dolby_vision_input_stream_id"] = dolbyVisionInputStreamId
	}

	var responseModel model.DolbyVisionInputStream
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/input-streams/dolby-vision/{dolby_vision_input_stream_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List Dolby Vision input stream
func (api *EncodingEncodingsInputStreamsDolbyVisionAPI) List(encodingId string, queryParams ...func(*EncodingEncodingsInputStreamsDolbyVisionAPIListQueryParams)) (*pagination.DolbyVisionInputStreamsListPagination, error) {
	queryParameters := &EncodingEncodingsInputStreamsDolbyVisionAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.DolbyVisionInputStreamsListPagination
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/input-streams/dolby-vision", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingEncodingsInputStreamsDolbyVisionAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsInputStreamsDolbyVisionAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsInputStreamsDolbyVisionAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
