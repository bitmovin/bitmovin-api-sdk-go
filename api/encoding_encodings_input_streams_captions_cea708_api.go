package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingEncodingsInputStreamsCaptionsCea708API communicates with '/encoding/encodings/{encoding_id}/input-streams/captions/cea708' endpoints
type EncodingEncodingsInputStreamsCaptionsCea708API struct {
	apiClient *apiclient.APIClient
}

// NewEncodingEncodingsInputStreamsCaptionsCea708API constructor for EncodingEncodingsInputStreamsCaptionsCea708API that takes options as argument
func NewEncodingEncodingsInputStreamsCaptionsCea708API(options ...apiclient.APIClientOption) (*EncodingEncodingsInputStreamsCaptionsCea708API, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingEncodingsInputStreamsCaptionsCea708APIWithClient(apiClient), nil
}

// NewEncodingEncodingsInputStreamsCaptionsCea708APIWithClient constructor for EncodingEncodingsInputStreamsCaptionsCea708API that takes an APIClient as argument
func NewEncodingEncodingsInputStreamsCaptionsCea708APIWithClient(apiClient *apiclient.APIClient) *EncodingEncodingsInputStreamsCaptionsCea708API {
	a := &EncodingEncodingsInputStreamsCaptionsCea708API{apiClient: apiClient}
	return a
}

// Create Add CEA 708 Input Stream
func (api *EncodingEncodingsInputStreamsCaptionsCea708API) Create(encodingId string, cea708CaptionInputStream model.Cea708CaptionInputStream) (*model.Cea708CaptionInputStream, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
	}

	var responseModel model.Cea708CaptionInputStream
	err := api.apiClient.Post("/encoding/encodings/{encoding_id}/input-streams/captions/cea708", &cea708CaptionInputStream, &responseModel, reqParams)
	return &responseModel, err
}

// Delete CEA 708 Input Stream
func (api *EncodingEncodingsInputStreamsCaptionsCea708API) Delete(encodingId string, inputStreamId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["input_stream_id"] = inputStreamId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/input-streams/captions/cea708/{input_stream_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get CEA 708 Input Stream Details
func (api *EncodingEncodingsInputStreamsCaptionsCea708API) Get(encodingId string, inputStreamId string) (*model.Cea708CaptionInputStream, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.PathParams["input_stream_id"] = inputStreamId
	}

	var responseModel model.Cea708CaptionInputStream
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/input-streams/captions/cea708/{input_stream_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List CEA 708 Input Streams
func (api *EncodingEncodingsInputStreamsCaptionsCea708API) List(encodingId string, queryParams ...func(*EncodingEncodingsInputStreamsCaptionsCea708APIListQueryParams)) (*pagination.Cea708CaptionInputStreamsListPagination, error) {
	queryParameters := &EncodingEncodingsInputStreamsCaptionsCea708APIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["encoding_id"] = encodingId
		params.QueryParams = queryParameters
	}

	var responseModel pagination.Cea708CaptionInputStreamsListPagination
	err := api.apiClient.Get("/encoding/encodings/{encoding_id}/input-streams/captions/cea708", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingEncodingsInputStreamsCaptionsCea708APIListQueryParams contains all query parameters for the List endpoint
type EncodingEncodingsInputStreamsCaptionsCea708APIListQueryParams struct {
	Offset int32 `query:"offset"`
	Limit  int32 `query:"limit"`
}

// Params will return a map of query parameters
func (q *EncodingEncodingsInputStreamsCaptionsCea708APIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
