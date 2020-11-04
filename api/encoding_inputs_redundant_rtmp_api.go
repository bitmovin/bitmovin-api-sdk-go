package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingInputsRedundantRtmpAPI communicates with '/encoding/inputs/redundant-rtmp' endpoints
type EncodingInputsRedundantRtmpAPI struct {
	apiClient *apiclient.APIClient
}

// NewEncodingInputsRedundantRtmpAPI constructor for EncodingInputsRedundantRtmpAPI that takes options as argument
func NewEncodingInputsRedundantRtmpAPI(options ...apiclient.APIClientOption) (*EncodingInputsRedundantRtmpAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingInputsRedundantRtmpAPIWithClient(apiClient), nil
}

// NewEncodingInputsRedundantRtmpAPIWithClient constructor for EncodingInputsRedundantRtmpAPI that takes an APIClient as argument
func NewEncodingInputsRedundantRtmpAPIWithClient(apiClient *apiclient.APIClient) *EncodingInputsRedundantRtmpAPI {
	a := &EncodingInputsRedundantRtmpAPI{apiClient: apiClient}
	return a
}

// Create Redundant RTMP Input
func (api *EncodingInputsRedundantRtmpAPI) Create(redundantRtmpInput model.RedundantRtmpInput) (*model.RedundantRtmpInput, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.RedundantRtmpInput
	err := api.apiClient.Post("/encoding/inputs/redundant-rtmp", &redundantRtmpInput, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Redundant RTMP Input
func (api *EncodingInputsRedundantRtmpAPI) Delete(inputId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["input_id"] = inputId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/inputs/redundant-rtmp/{input_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Redundant RTMP Input Details
func (api *EncodingInputsRedundantRtmpAPI) Get(inputId string) (*model.RedundantRtmpInput, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["input_id"] = inputId
	}

	var responseModel model.RedundantRtmpInput
	err := api.apiClient.Get("/encoding/inputs/redundant-rtmp/{input_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List Redundant RTMP Inputs
func (api *EncodingInputsRedundantRtmpAPI) List(queryParams ...func(*EncodingInputsRedundantRtmpAPIListQueryParams)) (*pagination.RedundantRtmpInputsListPagination, error) {
	queryParameters := &EncodingInputsRedundantRtmpAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.RedundantRtmpInputsListPagination
	err := api.apiClient.Get("/encoding/inputs/redundant-rtmp", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingInputsRedundantRtmpAPIListQueryParams contains all query parameters for the List endpoint
type EncodingInputsRedundantRtmpAPIListQueryParams struct {
	Offset int32  `query:"offset"`
	Limit  int32  `query:"limit"`
	Name   string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingInputsRedundantRtmpAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
