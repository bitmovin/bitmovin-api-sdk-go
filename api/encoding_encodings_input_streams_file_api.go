package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsInputStreamsFileAPI communicates with '/encoding/encodings/{encoding_id}/input-streams/file' endpoints
type EncodingEncodingsInputStreamsFileAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingEncodingsInputStreamsFileAPI constructor for EncodingEncodingsInputStreamsFileAPI that takes options as argument
func NewEncodingEncodingsInputStreamsFileAPI(options ...apiclient.APIClientOption) (*EncodingEncodingsInputStreamsFileAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsInputStreamsFileAPIWithClient(apiClient), nil
}

// NewEncodingEncodingsInputStreamsFileAPIWithClient constructor for EncodingEncodingsInputStreamsFileAPI that takes an APIClient as argument
func NewEncodingEncodingsInputStreamsFileAPIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsInputStreamsFileAPI {
	a := &EncodingEncodingsInputStreamsFileAPI{apiClient: apiClient}
	return a
}

// Create Add File input stream
func (api *EncodingEncodingsInputStreamsFileAPI) Create(encodingId string, fileInputStream model.FileInputStream) (*model.FileInputStream, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	var responseModel model.FileInputStream
	err := api.apiClient.Post("/encoding/encodings/{encoding_id}/input-streams/file", &fileInputStream, &responseModel, reqParams)
	return &responseModel, err
}

// Delete File stream
func (api *EncodingEncodingsInputStreamsFileAPI) Delete(encodingId string, inputStreamId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["input_stream_id"] = inputStreamId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/input-streams/file/{input_stream_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get File input stream details
func (api *EncodingEncodingsInputStreamsFileAPI) Get(encodingId string, inputStreamId string) (*model.FileInputStream, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["input_stream_id"] = inputStreamId
	}

	var responseModel model.FileInputStream
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/input-streams/file/{input_stream_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List File input stream
func (api *EncodingEncodingsInputStreamsFileAPI) List(encodingId string, queryParams ...func(*EncodingEncodingsInputStreamsFileAPIListQueryParams)) (*pagination.FileInputStreamsListPagination, error) {
	queryParameters := &EncodingEncodingsInputStreamsFileAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.FileInputStreamsListPagination
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/input-streams/file", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingEncodingsInputStreamsFileAPIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsInputStreamsFileAPIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsInputStreamsFileAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
