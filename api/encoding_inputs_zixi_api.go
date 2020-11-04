package api

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

// EncodingInputsZixiAPI communicates with '/encoding/inputs/zixi' endpoints
type EncodingInputsZixiAPI struct {
	apiClient *apiclient.APIClient

	// Customdata communicates with '/encoding/inputs/zixi/{input_id}/customData' endpoints
	Customdata *EncodingInputsZixiCustomdataAPI
}

// NewEncodingInputsZixiAPI constructor for EncodingInputsZixiAPI that takes options as argument
func NewEncodingInputsZixiAPI(options ...apiclient.APIClientOption) (*EncodingInputsZixiAPI, error) {
	apiClient, err := apiclient.NewAPIClient(options...)
	if err != nil {
		return nil, err
	}

	return NewEncodingInputsZixiAPIWithClient(apiClient), nil
}

// NewEncodingInputsZixiAPIWithClient constructor for EncodingInputsZixiAPI that takes an APIClient as argument
func NewEncodingInputsZixiAPIWithClient(apiClient *apiclient.APIClient) *EncodingInputsZixiAPI {
	a := &EncodingInputsZixiAPI{apiClient: apiClient}
	a.Customdata = NewEncodingInputsZixiCustomdataAPIWithClient(apiClient)

	return a
}

// Create Zixi input
func (api *EncodingInputsZixiAPI) Create(zixiInput model.ZixiInput) (*model.ZixiInput, error) {
	reqParams := func(params *apiclient.RequestParams) {
	}

	var responseModel model.ZixiInput
	err := api.apiClient.Post("/encoding/inputs/zixi", &zixiInput, &responseModel, reqParams)
	return &responseModel, err
}

// Delete Zixi input
func (api *EncodingInputsZixiAPI) Delete(inputId string) (*model.ZixiInput, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["input_id"] = inputId
	}

	var responseModel model.ZixiInput
	err := api.apiClient.Delete("/encoding/inputs/zixi/{input_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// Get Zixi Input Details
func (api *EncodingInputsZixiAPI) Get(inputId string) (*model.ZixiInput, error) {
	reqParams := func(params *apiclient.RequestParams) {
		params.PathParams["input_id"] = inputId
	}

	var responseModel model.ZixiInput
	err := api.apiClient.Get("/encoding/inputs/zixi/{input_id}", nil, &responseModel, reqParams)
	return &responseModel, err
}

// List Zixi inputs
func (api *EncodingInputsZixiAPI) List(queryParams ...func(*EncodingInputsZixiAPIListQueryParams)) (*pagination.ZixiInputsListPagination, error) {
	queryParameters := &EncodingInputsZixiAPIListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
	}

	reqParams := func(params *apiclient.RequestParams) {
		params.QueryParams = queryParameters
	}

	var responseModel pagination.ZixiInputsListPagination
	err := api.apiClient.Get("/encoding/inputs/zixi", nil, &responseModel, reqParams)
	return &responseModel, err
}

// EncodingInputsZixiAPIListQueryParams contains all query parameters for the List endpoint
type EncodingInputsZixiAPIListQueryParams struct {
	Offset int32  `query:"offset"`
	Limit  int32  `query:"limit"`
	Name   string `query:"name"`
}

// Params will return a map of query parameters
func (q *EncodingInputsZixiAPIListQueryParams) Params() map[string]string {
	return apiclient.GetParamsMap(q)
}
