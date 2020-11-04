package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingInputsLocalAPI communicates with '/encoding/inputs/local' endpoints
type EncodingInputsLocalAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/inputs/local/{input_id}/customData' endpoints
	Customdata *EncodingInputsLocalCustomdataAPI
}

// NewEncodingInputsLocalAPI constructor for EncodingInputsLocalAPI that takes options as argument
func NewEncodingInputsLocalAPI(options ...apiclient.APIClientOption) (*EncodingInputsLocalAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingInputsLocalAPIWithClient(apiClient), nil
}

// NewEncodingInputsLocalAPIWithClient constructor for EncodingInputsLocalAPI that takes an APIClient as argument
func NewEncodingInputsLocalAPIWithClient(apiClient *apiclient.APIClient) *EncodingInputsLocalAPI {
	a := &EncodingInputsLocalAPI{apiClient: apiClient}
	a.Customdata = NewEncodingInputsLocalCustomdataAPIWithClient(apiClient)

	return a
}

// Create Local Input
func (api *EncodingInputsLocalAPI) Create(localInput model.LocalInput) (*model.LocalInput, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.LocalInput
	err := api.apiClient.Post("/encoding/inputs/local", &localInput, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Local Input
func (api *EncodingInputsLocalAPI) Delete(inputId string) (*model.BitmovinResponse, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["input_id"] = inputId
	}

	var responseModel model.BitmovinResponse
	err := api.apiClient.Delete("/encoding/inputs/local/{input_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Local Input Details
func (api *EncodingInputsLocalAPI) Get(inputId string) (*model.LocalInput, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["input_id"] = inputId
	}

	var responseModel model.LocalInput
	err := api.apiClient.Get("/encoding/inputs/local/{input_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List Local Inputs
func (api *EncodingInputsLocalAPI) List(queryParams ...func(*EncodingInputsLocalAPIListQueryParams)) (*pagination.LocalInputsListPagination, error) {
	queryParameters := &EncodingInputsLocalAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.LocalInputsListPagination
	err := api.apiClient.Get("/encoding/inputs/local", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingInputsLocalAPIListQueryParams contains all query parameters for the List endpoint
type EncodingInputsLocalAPIListQueryParams struct {
	Offset int32  `query:"offset"`
	Limit  int32  `query:"limit"`
	Name   string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingInputsLocalAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
