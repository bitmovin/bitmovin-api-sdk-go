package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingInputsHttpAPI communicates with '/encoding/inputs/http' endpoints
type EncodingInputsHttpAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/inputs/http/{input_id}/customData' endpoints
	Customdata *EncodingInputsHttpCustomdataAPI
}

// NewEncodingInputsHttpAPI constructor for EncodingInputsHttpAPI that takes options as argument
func NewEncodingInputsHttpAPI(options ...apiclient.APIClientOption) (*EncodingInputsHttpAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingInputsHttpAPIWithClient(apiClient), nil
}

// NewEncodingInputsHttpAPIWithClient constructor for EncodingInputsHttpAPI that takes an APIClient as argument
func NewEncodingInputsHttpAPIWithClient(apiClient *apiclient.APIClient) *EncodingInputsHttpAPI {
	a := &EncodingInputsHttpAPI{apiClient: apiClient}
	a.Customdata = NewEncodingInputsHttpCustomdataAPIWithClient(apiClient)

	return a
}

// Create HTTP Input
func (api *EncodingInputsHttpAPI) Create(httpInput model.HttpInput) (*model.HttpInput, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.HttpInput
	err := api.apiClient.Post("/encoding/inputs/http", &httpInput, &responseModel, reqParams)
	return &responseModel, err
}

// Delete HTTP Input
func (api *EncodingInputsHttpAPI) Delete(inputId string) (*model.HttpInput, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["input_id"] = inputId
	}

	var responseModel model.HttpInput
	err := api.apiClient.Delete("/encoding/inputs/http/{input_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get HTTP Input Details
func (api *EncodingInputsHttpAPI) Get(inputId string) (*model.HttpInput, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["input_id"] = inputId
	}

	var responseModel model.HttpInput
	err := api.apiClient.Get("/encoding/inputs/http/{input_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List HTTP Inputs
func (api *EncodingInputsHttpAPI) List(queryParams ...func(*EncodingInputsHttpAPIListQueryParams)) (*pagination.HttpInputsListPagination, error) {
	queryParameters := &EncodingInputsHttpAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.HttpInputsListPagination
	err := api.apiClient.Get("/encoding/inputs/http", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingInputsHttpAPIListQueryParams contains all query parameters for the List endpoint
type EncodingInputsHttpAPIListQueryParams struct {
	Offset int32  `query:"offset"`
	Limit  int32  `query:"limit"`
	Name   string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingInputsHttpAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
