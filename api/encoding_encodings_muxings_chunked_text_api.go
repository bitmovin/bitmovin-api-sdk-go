package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsMuxingsChunkedTextAPI communicates with '/encoding/encodings/{encoding_id}/muxings/chunked-text' endpoints
type EncodingEncodingsMuxingsChunkedTextAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/encodings/{encoding_id}/muxings/chunked-text/{muxing_id}/customData' endpoints
	Customdata *EncodingEncodingsMuxingsChunkedTextCustomdataAPI
}

// NewEncodingEncodingsMuxingsChunkedTextAPI constructor for EncodingEncodingsMuxingsChunkedTextAPI that takes options as argument
func NewEncodingEncodingsMuxingsChunkedTextAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsMuxingsChunkedTextAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsMuxingsChunkedTextAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsMuxingsChunkedTextAPIWithClient constructor for EncodingEncodingsMuxingsChunkedTextAPI that takes an APIClient as argument
func NewEncodingEncodingsMuxingsChunkedTextAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsMuxingsChunkedTextAPI {
	a := &EncodingEncodingsMuxingsChunkedTextAPI{apiClient: apiClient}
	a.Customdata = NewEncodingEncodingsMuxingsChunkedTextCustomdataAPIWithClient(apiClient)

	return a
}

// Create Add Chunked Text muxing
func (api *EncodingEncodingsMuxingsChunkedTextAPI) Create(encodingId string, chunkedTextMuxing model.ChunkedTextMuxing) (*model.ChunkedTextMuxing, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	var responseModel model.ChunkedTextMuxing
	err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/chunked-text", &chunkedTextMuxing, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Chunked Text muxing
func (api *EncodingEncodingsMuxingsChunkedTextAPI) Delete(encodingId string, muxingId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/chunked-text/{muxing_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Chunked Text muxing details
func (api *EncodingEncodingsMuxingsChunkedTextAPI) Get(encodingId string, muxingId string) (*model.ChunkedTextMuxing, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["muxing_id"] = muxingId
	}

	var responseModel model.ChunkedTextMuxing
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/chunked-text/{muxing_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List Chunked Text muxings
func (api *EncodingEncodingsMuxingsChunkedTextAPI) List(encodingId string, queryParams ...func(*EncodingEncodingsMuxingsChunkedTextAPIListQueryParams)) (*pagination.ChunkedTextMuxingsListPagination, error) {
	queryParameters := &EncodingEncodingsMuxingsChunkedTextAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.ChunkedTextMuxingsListPagination
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/chunked-text", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingEncodingsMuxingsChunkedTextAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsMuxingsChunkedTextAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsMuxingsChunkedTextAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
