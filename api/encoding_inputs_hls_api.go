package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingInputsHlsAPI communicates with '/encoding/inputs/hls' endpoints
type EncodingInputsHlsAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingInputsHlsAPI constructor for EncodingInputsHlsAPI that takes options as argument
func NewEncodingInputsHlsAPI(options ...apiclient.APIClientOption) (*EncodingInputsHlsAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingInputsHlsAPIWithClient(apiClient), nil
}

// NewEncodingInputsHlsAPIWithClient constructor for EncodingInputsHlsAPI that takes an APIClient as argument
func NewEncodingInputsHlsAPIWithClient(apiClient *apiclient.APIClient) *EncodingInputsHlsAPI {
	a := &EncodingInputsHlsAPI{apiClient: apiClient}
	return a
}

// Create HLS input
func (api *EncodingInputsHlsAPI) Create(hlsInput model.HlsInput) (*model.HlsInput, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.HlsInput
	err := api.apiClient.Post("/encoding/inputs/hls", &hlsInput, &responseModel, reqParams)
	return &responseModel, err
}

// Delete HLS Input
func (api *EncodingInputsHlsAPI) Delete(inputId string) (*model.HlsInput, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["input_id"] = inputId
	}

	var responseModel model.HlsInput
	err := api.apiClient.Delete("/encoding/inputs/hls/{input_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get HLS Input Details
func (api *EncodingInputsHlsAPI) Get(inputId string) (*model.HlsInput, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["input_id"] = inputId
	}

	var responseModel model.HlsInput
	err := api.apiClient.Get("/encoding/inputs/hls/{input_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List HLS inputs
func (api *EncodingInputsHlsAPI) List(queryParams ...func(*EncodingInputsHlsAPIListQueryParams)) (*pagination.HlsInputsListPagination, error) {
	queryParameters := &EncodingInputsHlsAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.HlsInputsListPagination
	err := api.apiClient.Get("/encoding/inputs/hls", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingInputsHlsAPIListQueryParams contains all query parameters for the List endpoint
type EncodingInputsHlsAPIListQueryParams struct {
	Offset int32  `query:"offset"`
	Limit  int32  `query:"limit"`
	Name   string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingInputsHlsAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
